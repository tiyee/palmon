package schema

type ISchema interface {
	Valid() error
	Hook()
}
