package collections

/**
 * 判断一个数组是否为空 or nil
 */
func IsEmpty(arr []interface{}) bool {
	return arr == nil || len(arr) <= 0
}
