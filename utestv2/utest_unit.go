package utestv2

import (
	"testing"
)

type V interface{}

type TestHandle func(t *testing.T, r V) V

type uTest struct {
	name        string
	dependsName string
	handle      TestHandle
	pool        *testsList
}

//设置测试名称
func Name(name string) *uTest {

	return &uTest{
		name: name,
	}
}

//设置测试方法池子
func (u *uTest) SetPool(pool *testsList) *uTest {
	u.pool = pool

	return u
}

//设置测试依赖
func (u *uTest) Depends(name string) *uTest {

	return &uTest{
		name:        u.name,
		dependsName: name,
		handle:      u.handle,
		pool:        u.pool,
	}
}

//设置测试处理器
func (u *uTest) Test(handle TestHandle) {

	//追加测试用例
	u.pool.add(&testNode{
		val: &uTest{
			name:        u.name,
			dependsName: u.dependsName,
			handle:      handle,
		}})
}
