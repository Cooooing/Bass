package constant

type OssType string

func (o OssType) String() string {
	return string(o)
}

const (
	Minio OssType = "minio"
	Qiniu OssType = "qiniu"
)
