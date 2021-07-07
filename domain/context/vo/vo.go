package vo

import (
	"demo/domain/common"
	"demo/domain/context"
)

var _ context.IValueObject = new(Id)
var _ context.IValueObject = new(GenId)

// ========== 自增Id值对象 ==========
type GenId struct {
	Value string
}

func (g GenId) Valid() *context.ServerError {
	if len(g.Value) != 32 {
		return common.ErrorVoInit
	}

	return nil
}

// ========== Id 值对象 ==========
type Id struct {
	Value int64
}

func (i Id) Valid() *context.ServerError {
	if i.Value <= 0 {
		return common.ErrorVoInit
	}
	return nil
}
