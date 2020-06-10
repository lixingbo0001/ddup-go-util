package util

func RegMust(f func() error) {
	if err := f(); err != nil {
		panic("注册失败:" + err.Error())
	}
}
