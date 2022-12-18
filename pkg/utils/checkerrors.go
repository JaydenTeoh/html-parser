package checkerrors

func CheckErrors(err error) {
	if err != nil {
		panic(err)
	}
}
