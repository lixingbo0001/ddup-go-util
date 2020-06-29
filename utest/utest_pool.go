package utest

import (
	"fmt"
	"testing"
)

type testNode struct {
	val  *uTest
	next *testNode
}

type testsList struct {
	testMap map[string]*testNode
}

//获取测试用例池对象
func getTestsList() *testsList {
	return &testsList{
		testMap: make(map[string]*testNode),
	}
}

//测试用例是否存在
func (l *testsList) isExists(name string) bool {
	_, exists := l.testMap[name]

	return exists
}

//获取测试用例的依赖结果
func (l *testsList) depends(node *testNode, t *testing.T) interface{} {

	if node == nil {
		return nil
	}

	return node.val.handle(t, l.depends(node.next, t))
}

//获取测试用例
func (l *testsList) get(name string) *testNode {

	if !l.isExists(name) {
		panic(fmt.Sprintf("测试用例'不'存在[%s]", name))
	}

	return l.testMap[name]
}

//追加测试用例
func (l *testsList) add(tNode *testNode) {

	u := tNode.val

	if l.isExists(u.name) {
		panic(fmt.Sprintf("测试用例'已经'存在[%s]", u.name))
	}

	if u.dependsName != "" && !l.isExists(u.dependsName) {
		panic(fmt.Sprintf("测试用例依赖项'不'存在[%s]", u.dependsName))
	}

	//设置依赖链条
	if u.dependsName != "" {
		tNode.next = l.get(u.dependsName)
	}

	l.testMap[u.name] = tNode
}
