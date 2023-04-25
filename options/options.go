package options

/*
Option[T] is a generic type, it can be used to create a new type with a generic field.
For example, if you want to create a new type with a field of type T, you can use Option[T] to do it.


*/

import (
	"context"
	"time"

	"github.com/zhuangzhi/busters/function"
)

type Option[T any] struct {
	T *T
}

func (opt *Option[T]) Of(t *T) {
	opt.T = t
}

type ContextOption[T any] struct {
	Option[T]
	Context context.Context
}

func (opt *ContextOption[T]) WithContext(ctx context.Context) *T {
	opt.Context = ctx
	return opt.T
}

type TimeoutOption[T any] struct {
	Option[T]
	Timeout time.Duration
}

func (opt *TimeoutOption[T]) WithTimeout(timeout time.Duration) *T {
	opt.Timeout = timeout
	return opt.T
}

type RetryOption[T any] struct {
	Option[T]
	MaxRetryTimes int
	RetryInterval time.Duration
}

func (opt *RetryOption[T]) WithRetry(maxRetryTimes int, retryInterval time.Duration) *T {
	opt.MaxRetryTimes = maxRetryTimes
	opt.RetryInterval = retryInterval
	return opt.T
}

// KeepAliveOption ... KeepAlive option
type KeepAliveOption[T any] struct {
	Option[T]
	KeepAlive time.Duration
}

// WithKeepAlive ... Set KeepAlive option
func (opt *KeepAliveOption[T]) WithKeepAlive(keepAlive time.Duration) *T {
	opt.KeepAlive = keepAlive
	return opt.T
}

// BufferOption ... Buffer option
type BufferOption[T any] struct {
	Option[T]
	BufferSize int
}

// WithBuffer ... Set Buffer option
func (opt *BufferOption[T]) WithBuffer(bufferSize int) *T {
	opt.BufferSize = bufferSize
	return opt.T
}

// PriorityOption ... Priority option
type PriorityOption[T any] struct {
	Option[T]
	Priority int
}

// WithPriority ... Set priority option 0 is the top priority
func (opt *PriorityOption[T]) WithPriority(priority int) *T {
	opt.Priority = priority
	return opt.T
}

type OnError func(error)

type OnErrorOption[T any] struct {
	Option[T]
	OnError function.Consumer[error]
}

func (opt *OnErrorOption[T]) WithOnError(onError function.Consumer[error]) *T {
	opt.OnError = onError
	return opt.T
}

// OnResultOption ... OnResult option
type OnResultOption[T any, R any] struct {
	Option[T]
	OnResult function.BiConsumer[R, error]
}

// WithOnResult ... Set OnResult option
func (opt *OnResultOption[T, R]) WithOnResult(onResult function.BiConsumer[R, error]) *T {
	opt.OnResult = onResult
	return opt.T
}

// OnSuccessOption ... OnSuccess option
type OnSuccessOption[T any, R any] struct {
	Option[T]
	OnSuccess function.Consumer[R]
}

// WithOnSuccess ... Set OnSuccess option
func (opt *OnSuccessOption[T, R]) WithOnSuccess(onSuccess function.Consumer[R]) *T {
	opt.OnSuccess = onSuccess
	return opt.T
}

// URLOption ... URL option
type URLOption[T any] struct {
	Option[T]
	URL string
}

// WithURL ... Set URL option
func (opt *URLOption[T]) WithURL(url string) *T {
	opt.URL = url
	return opt.T
}

// AddressOption ... Address (hostname, port) option
type AddressOption[T any] struct {
	Option[T]
	Host string
	Port uint16
}

// WithAddress ... Set Address option
func (opt *AddressOption[T]) WithAddress(host string, port uint16) *T {
	opt.Host = host
	opt.Port = port
	return opt.T
}

// ServersOption ... Servers option set server address list
type ServersOption[T any] struct {
	Option[T]
	Servers []string
}

// WithServers ... Set Servers option
func (opt *ServersOption[T]) WithServers(servers []string) *T {
	opt.Servers = servers
	return opt.T
}

type TLSOption[T any] struct {
	Option[T]
	CertFile string
	KeyFile  string
	CaFile   string
	IsTls    bool
}

func (opt *TLSOption[T]) WithTLS(certFile, keyFile, caFile string) *T {
	opt.CertFile = certFile
	opt.KeyFile = keyFile
	opt.CaFile = caFile
	opt.IsTls = true
	return opt.T
}

type UsernamePasswordOption[T any] struct {
	Option[T]
	Username string
	Password string
	HasAuth  bool
}

// WithUsernamePassword ... Set UsernamePassword option
func (opt *UsernamePasswordOption[T]) WithUsernamePassword(username, password string) *T {
	opt.Username = username
	opt.Password = password
	opt.HasAuth = true
	return opt.T
}

// TokenOption ... Token option set token
type TokenOption[T any] struct {
	Option[T]
	Token    string
	HasToken bool
}

// WithToken ... Set Token option
func (opt *TokenOption[T]) WithToken(token string) *T {
	opt.Token = token
	opt.HasToken = true
	return opt.T
}

// CapacityOption ... Capacity option set capacity
type CapacityOption[T any] struct {
	Option[T]
	Capacity int
}

// WithCapacity ... Set Capacity option
func (opt *CapacityOption[T]) WithCapacity(capacity int) *T {
	opt.Capacity = capacity
	return opt.T
}

// QueueSizeOption ... QueueSize option set queue size
type QueueSizeOption[T any] struct {
	Option[T]
	QueueSize int
}

// WithQueueSize ... Set QueueSize option
func (opt *QueueSizeOption[T]) WithQueueSize(queueSize int) *T {
	opt.QueueSize = queueSize
	return opt.T
}

// PoolSizeOption ... PoolSize option set pool size
type PoolSizeOption[T any] struct {
	Option[T]
	PoolSize int
}

// WithPoolSize ... Set PoolSize option
func (opt *PoolSizeOption[T]) WithPoolSize(poolSize int) *T {
	opt.PoolSize = poolSize
	return opt.T
}
