package cross_cut_test

import (
	"quake-log-reader/cross_cut"
	"testing"
)

func TestNewRegister(t *testing.T) {
	reg := cross_cut.NewRegister()
	if reg == nil {
		panic("fail to create new register")
	}
	if reg.ProcessEventUseCase == nil {
		panic("fail to create process event use case")
	}
}
