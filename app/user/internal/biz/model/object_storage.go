package model

import "user/internal/data/ent/gen"

type ObjectStorage struct {
	*gen.ObjectStorage
}

type UploadToken struct {
	Key   string
	Token string
}
