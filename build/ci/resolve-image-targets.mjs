import { appendFileSync, readFileSync } from "node:fs";
import { execFileSync } from "node:child_process";

function readMakeList(target) {
  const value = execFileSync("make", ["--no-print-directory", "-s", target], { encoding: "utf8" }).trim();
  if (!value) throw new Error(`make target ${target} returned no modules`);
  return value.split(/\s+/);
}

function readMonolithModules() {
  const catalog = readFileSync("monolith/internal/catalog/catalog.go", "utf8");
  const modules = [...catalog.matchAll(/\bNamed\("([^"]+)"/g)].map((match) => match[1]);
  if (!modules.length) throw new Error("monolith catalog contains no module descriptors");
  return new Set(modules);
}

// Services and BFFs are defined by Make; the monolith catalog is its own
// composition source. CI only reads these existing declarations.
const services = readMakeList("ci-services");
const monolithModules = readMonolithModules();
const full = process.env.EVENT_NAME === "workflow_dispatch" || /^0+$/.test(process.env.BASE_SHA ?? "");
const paths = full ? [] : execFileSync("git", ["diff", "--name-only", `${process.env.BASE_SHA}..${process.env.HEAD_SHA}`], { encoding: "utf8" })
  .trim().split("\n").filter(Boolean);

console.log(full ? "A full image build was requested." : `Changed paths:\n${paths.join("\n")}`);

const selectedServices = new Set();
let monolith = full;
if (full) services.forEach((name) => selectedServices.add(name));

for (const path of paths) {
  const parts = path.split("/");
  if (["build/docker/toolchain.Dockerfile", "build/make/toolchain.mk", "build/ci/resolve-image-targets.mjs"].includes(path)) {
    services.forEach((name) => selectedServices.add(name));
    monolith = true;
  } else if (path === ".github/workflows/image-build.yml" || path.startsWith("build/")) {
    services.forEach((name) => selectedServices.add(name));
    monolith = true;
  } else if (path.startsWith("common/pkg/") || ["common/go.mod", "common/go.sum"].includes(path) || (path.startsWith("common/proto/buf/") && path !== "common/proto/buf/gen.openapi.yaml")) {
    services.forEach((name) => selectedServices.add(name));
    monolith = true;
  } else if (path.startsWith("common/proto/app/")) {
    services.forEach((name) => selectedServices.add(name));
    monolith = true;
  } else if (parts[0] === "app" && services.includes(parts[1])) {
    selectedServices.add(parts[1]);
    monolith ||= monolithModules.has(parts[1]);
  } else if (path.startsWith("monolith/")) {
    monolith = true;
  }
}

const rows = [...selectedServices].sort().map((name) => ({ name }));
const output = {
  images: String(Boolean(rows.length || monolith)),
  services: JSON.stringify({ include: rows }),
  monolith: String(monolith),
};
for (const [key, value] of Object.entries(output)) appendFileSync(process.env.GITHUB_OUTPUT, `${key}=${value}\n`);
