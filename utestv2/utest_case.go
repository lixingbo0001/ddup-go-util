package utestv2

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

type TestCase struct {
	pool *testsList
}

func (u *TestCase) Name(name string) *uTest {
	if u.pool == nil {
		u.pool = getTestsList()
	}

	return Name(name).SetPool(u.pool)
}

func (u *TestCase) Run(name string, t *testing.T, message string) {
	curTest := u.pool.get(name)

	convey.Convey(message, t, func() {
		curTest.val.handle(t, u.pool.depends(curTest.next, t))
	})
}
