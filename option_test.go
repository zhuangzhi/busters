package option_test

import (
	"testing"

	"github.com/zhuangzhi/busters/option"
)

// SomeOption contains all options
type SomeOption struct {
	option.AddressOption[SomeOption]
	option.TimeoutOption[SomeOption]
	option.RetryOption[SomeOption]
	option.PriorityOption[SomeOption]
	option.TLSOption[SomeOption]
	option.PoolSizeOption[SomeOption]
	option.KeepAliveOption[SomeOption]
	option.ContextOption[SomeOption]
	option.CapacityOption[SomeOption]
	option.BufferOption[SomeOption]
	option.OnErrorOption[SomeOption]
	option.OnSuccessOption[SomeOption]
}

func NewSomeOption() *SomeOption {
	option := new(SomeOption)
	option.AddressOption.Of(option)
	option.TimeoutOption.Of(option)
	option.RetryOption.Of(option)
	option.PriorityOption.Of(option)
	return option
}

// Test all options
func Test(t *testing.T) {

}
