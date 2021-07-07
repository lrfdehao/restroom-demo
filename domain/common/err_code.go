package common

import "demo/domain/context"

var (
	ErrToiletNotFound = context.NewError(404, "找不到坑")

	ErrorVoInit = context.NewError(400, "数据校验失败")

	ErrRestroomNotFound = context.NewError(404, "找不到厕所")
)
