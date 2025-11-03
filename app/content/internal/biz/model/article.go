package model

import "content/internal/data/ent/gen"

type Article gen.Article

// Summary 文章摘要
func (a *Article) Summary() {
	r := []rune(a.Content)
	if len(r) > 20 {
		a.Content = string(r[:20]) + "..."
	}
}
