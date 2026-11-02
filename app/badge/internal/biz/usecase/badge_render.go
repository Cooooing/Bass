package usecase

import (
	"badge/internal/biz/repo"
	badgetemplate "badge/template"
	"bytes"
	v1 "common/proto/gen/badge/v1"
	"context"
	"encoding/base64"
	"fmt"
	"html"
	"math"
	"strconv"
	"strings"
	texttemplate "text/template"

	"image"
	_ "image/gif"  // Register the GIF decoder for legacy badge resources.
	_ "image/jpeg" // Register the JPEG decoder for legacy badge resources.
	"image/png"

	imagedraw "golang.org/x/image/draw"

	// Register the WebP decoder used by image.Decode for legacy non-GIF input.
	_ "golang.org/x/image/webp"
)

// BadgeRenderUsecase adapts the legacy /gen query contract to the pure SVG
// renderer. It keeps image fetching and byte transformation outside templates.
type BadgeRenderUsecase struct {
	resources                                   repo.BadgeResourceReader
	defaultSize, defaultBorder, defaultFontSize uint32
	defaultAnimation, defaultShadow             float64
	defaultGradientDirection                    string
}

func NewBadgeRenderUsecase(resources repo.BadgeResourceReader) *BadgeRenderUsecase {
	return &BadgeRenderUsecase{
		resources:                resources,
		defaultSize:              32,
		defaultBorder:            3,
		defaultFontSize:          15,
		defaultAnimation:         .5,
		defaultShadow:            .5,
		defaultGradientDirection: "to bottom",
	}
}

// fontType identifies the uploaded font binary so SVG @font-face receives the
// correct MIME type and CSS format, independent of the opaque asset URL.
func (*BadgeRenderUsecase) fontType(content []byte) (string, string) {
	if len(content) >= 4 {
		switch string(content[:4]) {
		case "wOFF":
			return "font/woff", "woff"
		case "wOF2":
			return "font/woff2", "woff2"
		case "OTTO":
			return "font/otf", "opentype"
		}
	}
	// TrueType starts with 0x00010000. Existing uploads are TTF by default
	// because the accepted-font list is restricted at the BFF boundary.
	return "font/ttf", "truetype"
}

// RenderSpec adapts the typed internal contract to the legacy-equivalent renderer.
func (u *BadgeRenderUsecase) RenderSpec(ctx context.Context, spec *v1.RenderSpec) (string, error) {
	if spec == nil {
		return "", fmt.Errorf("badge image asset is required")
	}
	if spec.GetImageAssetId() <= 0 {
		return "", fmt.Errorf("badge image asset is required")
	}

	// Numeric zero in RenderSpec means "use the renderer's automatic/default
	// value". Do not serialize it as "0": legacy Render treats a non-empty
	// bar length of zero as an explicit fixed width and collapses the SVG.
	params := badgeRenderParams{
		Text:          spec.GetText(),
		Background:    strings.Join(spec.GetBackgroundColors(), ","),
		FontColor:     strings.Join(spec.GetTextColors(), ","),
		BackgroundWay: spec.GetBackgroundDirection(),
		FontWay:       spec.GetTextDirection(),
	}
	if spec.GetScale() != 0 {
		params.Scale = spec.GetScale()
	}
	if spec.GetShadow() != 0 {
		params.Shadow = spec.GetShadow()
	}
	if spec.GetAnimationSeconds() != 0 {
		params.Animation = spec.GetAnimationSeconds()
	}
	if spec.GetSize() != 0 {
		params.Size = spec.GetSize()
	}
	if spec.GetBorder() != 0 {
		params.Border = spec.GetBorder()
	}
	if spec.GetTextBarLength() != 0 {
		params.BarLength = strconv.FormatUint(uint64(spec.GetTextBarLength()), 10)
	}
	if spec.GetTextSize() != 0 {
		params.FontSize = spec.GetTextSize()
	}
	if spec.GetTextBarRadius() != 0 {
		params.BarRadius = spec.GetTextBarRadius()
	}
	if err := u.addImage(ctx, &params, spec.GetImageAssetId()); err != nil {
		return "", err
	}
	if fontAssetID := spec.GetFontAssetId(); fontAssetID != 0 {
		fontBytes, err := u.resources.Read(ctx, fontAssetID)
		if err != nil {
			return "", err
		}
		params.FontData = base64.StdEncoding.EncodeToString(fontBytes)
		params.FontMIME, params.FontFormat = u.fontType(fontBytes)
	}
	return u.renderSVG(params)
}

func (u *BadgeRenderUsecase) addImage(ctx context.Context, params *badgeRenderParams, assetID int64) error {
	size, border := params.Size, params.Border
	if size == 0 {
		size = u.defaultSize
	}
	if border == 0 {
		border = u.defaultBorder
	}
	if size < border*2 {
		return fmt.Errorf("badge size must be at least twice its border")
	}

	content, err := u.resources.Read(ctx, assetID)
	if err != nil {
		return err
	}
	_, format, err := image.DecodeConfig(bytes.NewReader(content))
	if err != nil {
		return err
	}
	decoded, _, err := image.Decode(bytes.NewReader(content))
	if err != nil {
		return err
	}
	params.ImageColor = u.averageColor(decoded)
	if format == "gif" {
		params.ImageData = base64.StdEncoding.EncodeToString(content)
		params.ImageMIME = "image/gif"
		return nil
	}

	limit := int((size - border*2) * 4)
	square := u.square(decoded)
	if square.Bounds().Dx() > limit {
		target := image.NewRGBA(image.Rect(0, 0, limit, limit))
		imagedraw.CatmullRom.Scale(target, target.Bounds(), square, square.Bounds(), imagedraw.Over, nil)
		square = target
	}
	var converted bytes.Buffer
	if err = png.Encode(&converted, square); err != nil {
		return err
	}
	params.ImageData = base64.StdEncoding.EncodeToString(converted.Bytes())
	params.ImageMIME = "image/png"
	return nil
}

func (*BadgeRenderUsecase) square(source image.Image) image.Image {
	bounds := source.Bounds()
	size := bounds.Dx()
	if bounds.Dy() < size {
		size = bounds.Dy()
	}
	return source.(interface {
		SubImage(image.Rectangle) image.Image
	}).SubImage(image.Rect(bounds.Min.X+(bounds.Dx()-size)/2, bounds.Min.Y+(bounds.Dy()-size)/2, bounds.Min.X+(bounds.Dx()+size)/2, bounds.Min.Y+(bounds.Dy()+size)/2))
}
func (u *BadgeRenderUsecase) averageColor(source image.Image) [3]uint8 {
	bounds := source.Bounds()
	var red, green, blue, count uint64
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := source.At(x, y).RGBA()
			red += uint64(r >> 8)
			green += uint64(g >> 8)
			blue += uint64(b >> 8)
			count++
		}
	}
	if count == 0 {
		return [3]uint8{0, 123, 255}
	}
	return [3]uint8{
		uint8(min(red/count, uint64(255))),
		uint8(min(green/count, uint64(255))),
		uint8(min(blue/count, uint64(255))),
	}
}

// badgeRenderParams mirrors fishpi-badge-rust after the data adapter has
// fetched an image and transcoded non-GIF bytes to lossless PNG.
type badgeRenderParams struct {
	Text          string
	ImageData     string
	ImageColor    [3]uint8
	ImageMIME     string
	FontData      string
	FontMIME      string
	FontFormat    string
	Scale         float64
	Size          uint32
	Border        uint32
	BarLength     string
	FontSize      uint32
	BarRadius     uint32
	FontColor     string
	Background    string
	Shadow        float64
	Animation     float64
	BackgroundWay string
	FontWay       string
}

type badgeRenderData struct {
	Text, FontData, FontMIME, FontFormat, FontColor, Background, BackgroundWay, FontWay string
	ImageData                                                                           string
	Size, Border, BarLength, FontSize, BarRadius, AnimationMS                           uint32
	Shadow                                                                              float64
}

// renderSVG reproduces fishpi-badge-rust's foreignObject layout. Unlike its
// Tera template, Go escapes text and attribute values for public BFF safety.
func (u *BadgeRenderUsecase) renderSVG(params badgeRenderParams) (string, error) {
	data, err := u.normalize(params)
	if err != nil {
		return "", err
	}
	view, err := texttemplate.New("badge").Funcs(texttemplate.FuncMap{"add": func(a, b uint32) uint32 { return a + b }, "sub": func(a, b uint32) uint32 { return a - b }, "mul": func(a, b uint32) uint32 { return a * b }, "div": func(a, b uint32) uint32 { return a / b }}).Parse(badgetemplate.MustReadTemplate("svg/badge.svg.tmpl"))
	if err != nil {
		return "", err
	}
	var output strings.Builder
	if err = view.Execute(&output, data); err != nil {
		return "", err
	}
	return output.String(), nil
}

func (u *BadgeRenderUsecase) normalize(params badgeRenderParams) (badgeRenderData, error) {
	scale := params.Scale
	if scale == 0 {
		scale = 1
	}
	size := u.scale(params.Size, u.defaultSize, scale)
	border := u.scale(params.Border, u.defaultBorder, scale)
	fontSize := u.scale(params.FontSize, u.defaultFontSize, scale)
	if size < border*2 {
		return badgeRenderData{}, fmt.Errorf("badge size must be at least twice its border")
	}
	barLength := u.barLength(params.BarLength, params.Text, params.FontSize, params.Border)
	barLength = u.scale(barLength, 0, scale)
	barRadius := params.BarRadius
	if barRadius == 0 {
		barRadius = params.Size / 2
		if params.Size == 0 {
			barRadius = u.defaultSize / 2
		}
	}
	barRadius = u.scale(barRadius, 0, scale)
	background, baseColor := u.background(params.Background, params.ImageColor)
	fontColor := u.fontColor(params.FontColor, baseColor)
	shadow := params.Shadow
	if shadow == 0 {
		shadow = u.defaultShadow
	}
	animation := params.Animation
	if animation == 0 {
		animation = u.defaultAnimation
	}
	imageData := ""
	if params.ImageData != "" {
		imageData = "data:" + params.ImageMIME + ";base64," + params.ImageData
	}
	fontMIME, fontFormat := params.FontMIME, params.FontFormat
	if params.FontData != "" && fontMIME == "" {
		fontMIME = "font/ttf"
		fontFormat = "truetype"
	}
	return badgeRenderData{
		Text:          html.EscapeString(params.Text),
		ImageData:     imageData,
		FontData:      strings.TrimPrefix(params.FontData, "data:font/ttf;base64,"),
		FontMIME:      fontMIME,
		FontFormat:    fontFormat,
		Size:          size,
		Border:        border,
		BarLength:     barLength,
		FontSize:      fontSize,
		BarRadius:     barRadius,
		Background:    background,
		FontColor:     fontColor,
		Shadow:        shadow,
		AnimationMS:   uint32(animation * 1000),
		BackgroundWay: u.direction(params.BackgroundWay),
		FontWay:       u.direction(params.FontWay),
	}, nil
}

func (u *BadgeRenderUsecase) barLength(value, text string, fontSize, border uint32) uint32 {
	if value != "" && value != "auto" {
		parsed, err := strconv.ParseUint(value, 10, 32)
		if err == nil {
			return uint32(parsed)
		}
		return 0
	}
	if text == "" {
		return 0
	}
	if fontSize == 0 {
		fontSize = u.defaultFontSize
	}
	if border == 0 {
		border = u.defaultBorder
	}
	count := float64(len([]rune(text)))
	for _, character := range text {
		if character > 127 {
			count += .84
		}
	}
	return uint32(float64(fontSize)*count*.55 + 2.6*float64(border))
}
func (*BadgeRenderUsecase) scale(value, fallback uint32, scale float64) uint32 {
	if value == 0 {
		value = fallback
	}
	return uint32(float64(value) * scale)
}
func (u *BadgeRenderUsecase) background(value string, image [3]uint8) (string, [3]uint8) {
	if value != "" && !strings.EqualFold(value, "auto") {
		colors := u.colors(value)
		return strings.Join(colors, ", "), u.hex(strings.Split(value, ",")[0])
	}
	if image == [3]uint8{} {
		image = [3]uint8{0, 123, 255}
	}
	return u.autoGradient(image), image
}
func (u *BadgeRenderUsecase) fontColor(value string, background [3]uint8) string {
	if value != "" && !strings.EqualFold(value, "auto") {
		return strings.Join(u.colors(value), ", ")
	}
	if (float64(background[0])+float64(background[1])+float64(background[2]))/3 > 214 {
		return "rgb(33,33,33)"
	}
	return "rgb(255,255,255)"
}
func (u *BadgeRenderUsecase) colors(value string) []string {
	parts := strings.Split(value, ",")
	colors := make([]string, 0, len(parts))
	for _, part := range parts {
		color := u.hex(part)
		colors = append(colors, fmt.Sprintf("rgb(%d,%d,%d)", color[0], color[1], color[2]))
	}
	if len(colors) == 1 {
		colors = append(colors, colors[0])
	}
	return colors
}
func (*BadgeRenderUsecase) hex(value string) [3]uint8 {
	value = strings.TrimPrefix(strings.TrimSpace(value), "#")
	parsed, err := strconv.ParseUint(value, 16, 32)
	if len(value) != 6 || err != nil {
		return [3]uint8{255, 255, 255}
	}
	return [3]uint8{uint8((parsed >> 16) & 0xff), uint8((parsed >> 8) & 0xff), uint8(parsed & 0xff)}
}
func (u *BadgeRenderUsecase) autoGradient(color [3]uint8) string {
	hue, lightness, saturation := u.hsl(color)
	light := u.rgbFromHSL(hue, math.Min(lightness+.02, 1), saturation)
	dark := u.rgbFromHSL(hue, math.Max(lightness-.06, 0), saturation)
	return fmt.Sprintf("rgb(%d,%d,%d), rgb(%d,%d,%d)", light[0], light[1], light[2], dark[0], dark[1], dark[2])
}
func (*BadgeRenderUsecase) hsl(color [3]uint8) (float64, float64, float64) {
	r := float64(color[0]) / 255
	g := float64(color[1]) / 255
	b := float64(color[2]) / 255
	maximum := math.Max(r, math.Max(g, b))
	minimum := math.Min(r, math.Min(g, b))
	l := (maximum + minimum) / 2
	if maximum == minimum {
		return 0, l, 0
	}
	d := maximum - minimum
	s := d / (1 - math.Abs(2*l-1))
	var h float64
	switch maximum {
	case r:
		h = math.Mod((g-b)/d, 6)
	case g:
		h = (b-r)/d + 2
	default:
		h = (r-g)/d + 4
	}
	return h * 60, l, s
}
func (*BadgeRenderUsecase) rgbFromHSL(h, l, s float64) [3]uint8 {
	if s == 0 {
		v := uint8(math.Round(l * 255))
		return [3]uint8{v, v, v}
	}
	c := (1 - math.Abs(2*l-1)) * s
	x := c * (1 - math.Abs(math.Mod(h/60, 2)-1))
	m := l - c/2
	r, g, b := c, x, 0.0
	switch {
	case h < 60:
	case h < 120:
		r, g = x, c
	case h < 180:
		r, g, b = 0, c, x
	case h < 240:
		r, g, b = 0, x, c
	case h < 300:
		r, g, b = x, 0, c
	default:
		r, g, b = c, 0, x
	}
	return [3]uint8{uint8(math.Round((r + m) * 255)), uint8(math.Round((g + m) * 255)), uint8(math.Round((b + m) * 255))}
}
func (u *BadgeRenderUsecase) direction(value string) string {
	switch value {
	case "right":
		return "to right"
	case "left":
		return "to left"
	case "top":
		return "to top"
	case "bottom":
		return u.defaultGradientDirection
	case "top-right":
		return "to top right"
	case "top-left":
		return "to top left"
	case "bottom-right":
		return "to bottom right"
	case "bottom-left":
		return "to bottom left"
	case "":
		return u.defaultGradientDirection
	default:
		if before, ok := strings.CutSuffix(value, "deg"); ok {
			degrees, err := strconv.ParseUint(before, 10, 16)
			if err == nil && degrees <= 360 {
				return strconv.FormatUint(degrees, 10) + "deg"
			}
		}
		return u.defaultGradientDirection
	}
}
