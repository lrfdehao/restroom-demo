package vo

import "demo/domain/domain"

var (
	ErrorVoInit = domain.NewError("init_vo_error", "数据校验失败")
)

var _ domain.IValueObject = new(Id)
var _ domain.IValueObject = new(GenId)

// ========== 自增Id值对象 ==========
type GenId struct {
	value string
}

func (g GenId) Init(i interface{}) (domain.IValueObject, error) {
	vObj, ok := i.(string)
	if !ok || len(vObj) != 32 {
		return nil, ErrorVoInit
	}
	return GenId{value: vObj}, nil
}

// ========== Id 值对象 ==========
type Id struct {
	value int64
}

func (id Id) Init(i interface{}) (domain.IValueObject, error) {
	vObj, ok := i.(int64)
	if !ok {
		return nil, ErrorVoInit
	}
	return Id{value: vObj}, nil
}
