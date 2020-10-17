package test

import (
	"fmt"
	utest "github.com/winkb/ddup-go-util/utestv2"
	"strconv"
	"testing"
)

type userDomainTest struct {
	utest.TestCase
}

var ud = &userDomainTest{}

func init() {

	ud.Name("create").Test(func(t *testing.T, r utest.V) utest.V {
		fmt.Println("创建....")
		return 1
	})

	ud.Name("delete").Depends("create").Test(func(t *testing.T, r utest.V) utest.V {
		id := r.(int)

		fmt.Println("删除" + strconv.Itoa(id) + "....")

		return nil
	})

}

func TestStruct1(t *testing.T) {
	ud.Run("create", t, "创建测试")
}

func TestStruct2(t *testing.T) {
	ud.Run("delete", t, "删除测试")
}
