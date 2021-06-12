package status

import (
	"fmt"
)

type ErrDetail struct {
	ErrCode int32
	ErrMsg  string
}

var (
	Success  = ErrDetail{0000, "成功"}
	InnerErr = ErrDetail{0001, "内部错误"}

	// 1000 用户模块错误
	// 2000 文章模块错误
	// 3000 分类模块错误
)

// ErrorPrint 打印错误
func ErrPrint(err ErrDetail) {
	fmt.Println("ErrCode:", err.ErrCode, "ErrMsg:", err.ErrMsg)
	return
}
