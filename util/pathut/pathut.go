package pathut

import (
	"os"
	"path/filepath"
	"regexp"
)

var splitRe = regexp.MustCompile(`[\\/]`)

func JoinPath(elem ...string) string {

	if 0 == len(elem) {
		return ""
	}

	tmpSlice := make([]string, 0, len(elem))

	if len(elem[0]) > 0 {
		if elem[0][0] == '/' || elem[0][0] == '\\' {
			tmpSlice = append(tmpSlice, string(os.PathSeparator))

		}
	}

	for _, v := range elem {

		s := SplitPath(v)

		tmpSlice = append(tmpSlice, s...)
	}

	// 使用filepath.Join方法将临时切片中的子路径合并成一个完整的路径字符串，并返回
	return filepath.Join(tmpSlice...)
}

// SplitPath 将给定的路径字符串按照正斜杠("/")或反斜杠("\\")分割成多个部分，
// 并返回一个非空字符串的切片。
//
// 参数：
//
//	path string - 要分割的路径字符串
//
// 返回值：
//
//	[]string - 分割后的非空字符串切片
func SplitPath(path string) []string {
	// 使用正则表达式 splitRe 将路径字符串 path 按照规则分割成多个部分，并将结果保存在 parts 切片中
	parts := splitRe.Split(path, -1)
	var nonEmptyParts []string

	for _, part := range parts {
		// 遍历 parts 切片中的每个部分
		if part != "" {
			// 如果当前部分不为空字符串，则将其添加到 nonEmptyParts 切片中
			nonEmptyParts = append(nonEmptyParts, part)
		}
	}

	// 返回 nonEmptyParts 切片，其中只包含非空字符串
	return nonEmptyParts
}
