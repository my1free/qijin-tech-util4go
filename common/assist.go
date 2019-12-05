package common

import (
	"time"
)

/**
 * 重试包装方法
 * 使用示例:
 * ```go
 * f := func() (error) {
 *		r, err = Tt()
 *		return err
 * }
 * Retry(f, 3, 5*time.Millisecond)
 * func Tt() (string, error) {
 *		return "s", errors.New("asdfasdf")
 *	}
 * ```
 */
func Retry(f func() (error), maxTimes int, interval time.Duration) error {
	// 如果 maxTimes 参数非法，则只执行一次
	if maxTimes <= 0 {
		return f()
	}
	var err error
	for t := maxTimes; t > 0; t-- {
		if err = f(); err != nil {
			time.Sleep(interval)
			continue
		} else {
			break
		}
	}
	return err
}
