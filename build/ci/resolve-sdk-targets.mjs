import { appendFileSync } from "node:fs";
import { execFileSync } from "node:child_process";

function readMakeList(args, description) {
  const value = execFileSync("make", ["--no-print-directory", "-s", ...args], { encoding: "utf8" }).trim();
  if (!value) throw new Error(`${description} returned no values`);
  return value.split(/\s+/);
}

// Keep CI aligned with the same BFF and generator declarations used by Make.
const bffs = readMakeList(["ci-bff-servers"], "BFF server list");
const languages = readMakeList(["-C", `app/${bffs[0]}`, "ci-sdk-languages"], "SDK language list");
const full = process.env.EVENT_NAME === "workflow_dispatch" || /^0+$/.test(process.env.BASE_SHA ?? "");
const paths = full ? [] : execFileSync("git", ["diff", "--name-only", `${process.env.BASE_SHA}..${process.env.HEAD_SHA}`], { encoding: "utf8" })
  .trim().split("\n").filter(Boolean);

console.log(full ? "A full SDK generation was requested." : `Changed paths:\n${paths.join("\n")}`);

const selected = new Set(full ? bffs : []);
for (const path of paths) {
  const parts = path.split("/");
  if ([".github/workflows/sdk.yml", "build/ci/resolve-sdk-targets.mjs", "build/make/doc.mk", "common/proto/buf/gen.openapi.yaml"].includes(path) || path.startsWith("common/proto/sdk/")) {
    bffs.forEach((name) => selected.add(name));
  } else if (path.startsWith("common/proto/app/")) {
    if (bffs.includes(parts[3])) selected.add(parts[3]);
    else bffs.forEach((name) => selected.add(name));
  }
}

const rows = [...selected].sort().flatMap((service) => languages.map((language) => ({ service, language })));
appendFileSync(process.env.GITHUB_OUTPUT, `sdk=${Boolean(rows.length)}\n`);
appendFileSync(process.env.GITHUB_OUTPUT, `matrix=${JSON.stringify({ include: rows })}\n`);
