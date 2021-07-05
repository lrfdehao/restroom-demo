package domain

type IDomain interface {
}

type IValueObject interface {
	Init(interface{}) (IValueObject, error)
}
