package error

func IfErrNonNil(err error, f func(e error)) {
	if err != nil {
		f(err)
	}
}

func IfErrNonNilPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func ErrNonNilDo[T any](one T, err error, f func(e error)) T {
	if err != nil {
		f(err)
	}
	return one
}

func ErrNonNilDo2[T any, E any](one T, two E, err error, f func(e error)) (T, E) {
	if err != nil {
		f(err)
	}
	return one, two
}

func ErrNonNilPanic[T any](one T, err error) T {
	if err != nil {
		panic(err)
	}
	return one
}

func ErrNonNilPanic2[T any, E any](one T, two E, err error) (T, E) {
	if err != nil {
		panic(err)
	}
	return one, two
}
