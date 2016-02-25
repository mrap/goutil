package builtin

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}
