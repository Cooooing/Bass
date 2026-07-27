package template

import "embed"

//go:embed email/** station/**
var templateFS embed.FS

func ReadTemplate(path string) (string, error) {
	b, err := templateFS.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func MustReadTemplate(path string) string {
	content, err := ReadTemplate(path)
	if err != nil {
		panic(err)
	}
	return content
}
