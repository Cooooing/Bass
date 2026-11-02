package template

import "embed"

//go:embed svg/**
var templateFS embed.FS

func ReadTemplate(path string) (string, error) {
	content, err := templateFS.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func MustReadTemplate(path string) string {
	content, err := ReadTemplate(path)
	if err != nil {
		panic(err)
	}
	return content
}
