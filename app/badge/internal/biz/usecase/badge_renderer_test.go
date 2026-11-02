package usecase

import (
	"bytes"
	"context"
	"encoding/base64"
	"image"
	"image/color"
	"image/png"
	"strings"
	"testing"

	v1 "common/proto/gen/badge/v1"
)

type badgeRenderTestResourceReader struct {
	content  []byte
	contents map[int64][]byte
}

func (r badgeRenderTestResourceReader) Read(_ context.Context, assetID int64) ([]byte, error) {
	if content, ok := r.contents[assetID]; ok {
		return content, nil
	}
	return r.content, nil
}

func badgeRenderPNG(t *testing.T) []byte {
	t.Helper()
	content := new(bytes.Buffer)
	image := image.NewRGBA(image.Rect(0, 0, 2, 2))
	image.SetRGBA(0, 0, color.RGBA{R: 237, G: 143, B: 37, A: 255})
	if err := png.Encode(content, image); err != nil {
		t.Fatalf("encode PNG fixture: %v", err)
	}
	return content.Bytes()
}

func TestBadgeRendererMatchesFishPiDefaults(t *testing.T) {
	svg, err := NewBadgeRenderUsecase(nil).renderSVG(badgeRenderParams{Text: "FishPi"})
	if err != nil {
		t.Fatalf("render badge: %v", err)
	}
	for _, expected := range []string{
		`width="103px"`,
		`height="35px"`,
		`rgb(10,128,255), rgb(0,108,224)`,
		`rgb(255,255,255)`,
		`animation-duration: 500ms`,
		`font-size: 16px`,
	} {
		if !strings.Contains(svg, expected) {
			t.Fatalf("expected SVG fragment %q in %s", expected, svg)
		}
	}
}

func TestBadgeRendererEmbedsImageAndFontLikeFishPi(t *testing.T) {
	svg, err := NewBadgeRenderUsecase(nil).renderSVG(badgeRenderParams{
		Text:          "摸鱼者事竟成",
		ImageData:     "data:image/webp;base64,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",
		ImageMIME:     "image/png",
		FontData:      "AQID",
		Background:    "ed8f25,ff0000",
		FontColor:     "auto",
		BackgroundWay: "45deg",
		FontWay:       "top-left",
	})
	if err != nil {
		t.Fatalf("render badge: %v", err)
	}
	if !strings.Contains(svg, "data:image/png;base64,") || !strings.Contains(svg, "@font-face") {
		t.Fatalf("rendered SVG did not preserve embedded resources: %s", svg)
	}
}

func TestBadgeRendererEscapesText(t *testing.T) {
	svg, err := NewBadgeRenderUsecase(nil).renderSVG(badgeRenderParams{Text: `<script>alert(1)</script>`})
	if err != nil {
		t.Fatalf("render badge: %v", err)
	}
	if strings.Contains(svg, "<script>") || !strings.Contains(svg, "&lt;script&gt;") {
		t.Fatalf("badge text was not escaped: %s", svg)
	}
}

func TestBadgeRenderUsecaseSupportsPNG(t *testing.T) {
	usecase := NewBadgeRenderUsecase(badgeRenderTestResourceReader{content: badgeRenderPNG(t)})
	svg, err := usecase.RenderSpec(context.Background(), &v1.RenderSpec{
		Text:         "PNG",
		ImageAssetId: 1,
	})
	if err != nil {
		t.Fatalf("render PNG badge: %v", err)
	}
	if !strings.Contains(svg, "data:image/png;base64,") {
		t.Fatalf("rendered SVG does not contain converted PNG image: %s", svg)
	}
}

func TestBadgeRenderSpecKeepsAutomaticTextBarValues(t *testing.T) {
	usecase := NewBadgeRenderUsecase(badgeRenderTestResourceReader{content: badgeRenderPNG(t)})
	svg, err := usecase.RenderSpec(context.Background(), &v1.RenderSpec{
		ImageAssetId:     1,
		Text:             "摸鱼者事竟成",
		BackgroundColors: []string{"ed8f26"},
		TextColors:       []string{"ffffff"},
		Size:             32,
		Border:           3,
		TextSize:         15,
		Scale:            .79,
	})
	if err != nil {
		t.Fatalf("render typed spec: %v", err)
	}
	if strings.Contains(svg, `width="39px"`) {
		t.Fatalf("automatic text bar collapsed to zero width: %s", svg)
	}
	if !strings.Contains(svg, `animation-name: move`) {
		t.Fatalf("typed spec lost text bar animation: %s", svg)
	}
}

func TestBadgeRenderSpecRejectsImageWithInvalidSizeAndBorder(t *testing.T) {
	usecase := NewBadgeRenderUsecase(badgeRenderTestResourceReader{})
	_, err := usecase.RenderSpec(context.Background(), &v1.RenderSpec{
		ImageAssetId: 1,
		Size:         1,
		Border:       3,
	})
	if err == nil || !strings.Contains(err.Error(), "badge size") {
		t.Fatalf("expected invalid badge dimensions error, got %v", err)
	}
}

func TestBadgeRenderUsecaseDetectsUploadedFontFormats(t *testing.T) {
	usecase := NewBadgeRenderUsecase(nil)
	for _, test := range []struct {
		name    string
		content []byte
		mime    string
		format  string
	}{
		{name: "truetype", content: []byte{0, 1, 0, 0}, mime: "font/ttf", format: "truetype"},
		{name: "opentype", content: []byte("OTTO"), mime: "font/otf", format: "opentype"},
		{name: "woff", content: []byte("wOFF"), mime: "font/woff", format: "woff"},
		{name: "woff2", content: []byte("wOF2"), mime: "font/woff2", format: "woff2"},
	} {
		t.Run(test.name, func(t *testing.T) {
			mime, format := usecase.fontType(test.content)
			if mime != test.mime || format != test.format {
				t.Fatalf("font type = %q/%q, want %q/%q", mime, format, test.mime, test.format)
			}
		})
	}
}

func TestBadgeRendererRejectsUnsafeGradientDirection(t *testing.T) {
	usecase := NewBadgeRenderUsecase(nil)
	for _, direction := range []string{"bottom; color:red", "721deg", "url(https://example.com)"} {
		if value := usecase.direction(direction); value != "to bottom" {
			t.Fatalf("direction %q = %q, want fallback", direction, value)
		}
	}
	if value := usecase.direction("45deg"); value != "45deg" {
		t.Fatalf("safe degree direction = %q", value)
	}
}

func TestBadgeRenderSpecEmbedsUploadedFont(t *testing.T) {
	font := []byte("wOFFfont-payload")
	usecase := NewBadgeRenderUsecase(badgeRenderTestResourceReader{contents: map[int64][]byte{
		1: badgeRenderPNG(t),
		2: font,
	}})
	svg, err := usecase.RenderSpec(context.Background(), &v1.RenderSpec{
		ImageAssetId: 1,
		Text:         "FishPi",
		FontAssetId:  2,
	})
	if err != nil {
		t.Fatalf("render badge with uploaded font: %v", err)
	}

	encodedFont := base64.StdEncoding.EncodeToString(font)
	for _, expected := range []string{
		"@font-face",
		"font-family: 'CustomFont'",
		"data:font/woff;base64," + encodedFont,
		"format('woff')",
		"font-family: 'CustomFont',",
	} {
		if !strings.Contains(svg, expected) {
			t.Fatalf("expected embedded font fragment %q in %s", expected, svg)
		}
	}
}

func TestBadgeRenderSpecRequiresPlatformImageAsset(t *testing.T) {
	_, err := NewBadgeRenderUsecase(badgeRenderTestResourceReader{}).RenderSpec(context.Background(), &v1.RenderSpec{})
	if err == nil || !strings.Contains(err.Error(), "image asset") {
		t.Fatalf("expected missing image asset error, got %v", err)
	}
}
