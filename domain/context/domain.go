package context

type IDomain interface {
}

type IValueObject interface {
	Valid() *ServerError
}
