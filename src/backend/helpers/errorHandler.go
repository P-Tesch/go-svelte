package helpers

func HandleError(e error) {
	if e != nil {
		panic(e.Error())
	}
}
