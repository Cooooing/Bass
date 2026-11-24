package util

import "os"

// RemoveFile 删除单个受调用方路径校验约束的临时文件。
func RemoveFile(path string) error {
	return os.Remove(path)
}
