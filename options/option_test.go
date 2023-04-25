package options_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/zhuangzhi/busters/options"
	"time"

	"testing"
)

// SomeOption contains all options
type SomeOption struct {
	options.AddressOption[SomeOption]
	options.TimeoutOption[SomeOption]
	options.RetryOption[SomeOption]
	options.PriorityOption[SomeOption]
	options.TLSOption[SomeOption]
	options.PoolSizeOption[SomeOption]
	options.KeepAliveOption[SomeOption]
	options.ContextOption[SomeOption]
	options.CapacityOption[SomeOption]
	options.BufferOption[SomeOption]
	options.OnErrorOption[SomeOption]
	options.OnSuccessOption[SomeOption, int]
	options.ServersOption[SomeOption]
}

func NewSomeOption() *SomeOption {
	option := new(SomeOption)
	option.AddressOption.Of(option)
	option.ServersOption.Of(option)
	option.TimeoutOption.Of(option)
	option.RetryOption.Of(option)
	option.PriorityOption.Of(option)
	option.TLSOption.Of(option)
	option.PoolSizeOption.Of(option)
	option.KeepAliveOption.Of(option)
	option.ContextOption.Of(option)
	option.CapacityOption.Of(option)
	option.BufferOption.Of(option)
	option.OnErrorOption.Of(option)
	option.OnSuccessOption.Of(option)

	return option
}

// Test all options
func Test(t *testing.T) {
	ops := NewSomeOption()
	ops.WithAddress("address", 888).
		WithTimeout(time.Second).
		WithRetry(1, time.Second).
		WithPriority(1).
		WithCapacity(1).
		WithBuffer(1).
		WithServers([]string{"a", "b"}).
		WithContext(context.Background()).
		WithKeepAlive(time.Second).
		WithPoolSize(1).
		WithTLS("cert.pem", "key.pem", "ca.pem").
		WithOnError(func(err error) {
			t.Log(err)
		}).
		WithOnSuccess(func(i int) {
			t.Log(i)
		})

	assert.Equal(t, "address", ops.Host)
	assert.Equal(t, uint16(888), ops.Port)
	assert.Equal(t, time.Second, ops.Timeout)
	assert.Equal(t, 1, ops.MaxRetryTimes)
	assert.Equal(t, time.Second, ops.RetryInterval)
	assert.Equal(t, 1, ops.Priority)
	assert.Equal(t, 1, ops.Capacity)
	assert.Equal(t, 1, ops.BufferSize)
	assert.Equal(t, []string{"a", "b"}, ops.Servers)
	assert.Equal(t, context.Background(), ops.Context)
	assert.Equal(t, time.Second, ops.KeepAlive)
	assert.Equal(t, 1, ops.PoolSize)
	assert.Equal(t, "", ops.CertFile)
	assert.Equal(t, "", ops.KeyFile)
	assert.NotNil(t, ops.OnError)
	assert.NotNil(t, ops.OnSuccess)

}
