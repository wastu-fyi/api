package tests

import (
	"github.com/goravel/framework/testing"

	"wastu/bootstrap"
)

func init() {
	bootstrap.Boot()
}

type TestCase struct {
	testing.TestCase
}
