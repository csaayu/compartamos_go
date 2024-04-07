package helper

func ErrorSync(err error) {
	if err != nil {
		panic(err)
	}
}
