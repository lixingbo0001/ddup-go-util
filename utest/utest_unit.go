package utest

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var testListPool *testsList

type V interface{}

type TestHandle func(t *testing.T, r V) V

func init() {
	testListPool = getTestsList()
}

//执行测试
func Test(name string, t *testing.T, message string) {

	utest := testListPool.get(name)

	Convey(message, t, func() {
		utest.val.handle(t, testListPool.depends(utest.next, t))
	})
}

type uTest struct {
	name        string
	dependsName string
	handle      TestHandle
}

//设置测试名称
func Name(name string) *uTest {

	return &uTest{
		name: name,
	}
}

//设置测试依赖
func (u *uTest) Depends(name string) *uTest {

	return &uTest{
		name:        u.name,
		dependsName: name,
		handle:      nil,
	}
}

//设置测试处理器
func (u *uTest) Test(handle TestHandle) {

	//追加测试用例
	testListPool.add(&testNode{
		val: &uTest{
			name:        u.name,
			dependsName: u.dependsName,
			handle:      handle,
		}})
}
