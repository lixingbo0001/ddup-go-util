package test

import (
	"ddup-go-util/utest"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func init() {

	utest.Name("test1").Test(func(t *testing.T, r utest.V) utest.V {

		Convey("没有依赖", func() {

			So(1, ShouldEqual, 1)
		})

		return 111
	})

	utest.Name("test2").Depends("test1").Test(func(t *testing.T, r utest.V) utest.V {

		Convey("依赖:test1", func() {

			So(r, ShouldEqual, 111)
		})

		return 2
	})
}

func TestTest1(t *testing.T) {
	utest.Test("test1", t, "第一个测试")
}

func TestTest2Depends(t *testing.T) {
	utest.Test("test2", t, "第二个测试")
}

func TestTest3Depends(t *testing.T) {
	// utest.Test("test3", t, "第三个测试不存在")
}
