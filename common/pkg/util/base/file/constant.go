package file

// HashAlgorithm 类型
type HashAlgorithm string

// 可选算法常量
const (
	MD5    HashAlgorithm = "md5"
	SHA1   HashAlgorithm = "sha1"
	SHA256 HashAlgorithm = "sha256"
)
