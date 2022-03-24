package any

// the keyword "any" can now be used in place of "interface{}"

type AType any

type BType interface{}

type MyGeneric interface {
	AType | BType
}

func CreateSlice[T MyGeneric](args ...T) []T {
	return args
}
