package util

func PanicWhenError(err error, msg string) {
	if err != nil {
		panic(msg + "遇到错误:" + err.Error())
	}
}
