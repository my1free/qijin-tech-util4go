package strings

import "strings"

/**
 * 判断一个字符串是否为Blank
 */
func IsBlank(s string) bool {
	return len(strings.Trim(s, " ")) == 0
}
