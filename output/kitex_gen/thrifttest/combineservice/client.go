// Code generated by Kitex v0.0.8. DO NOT EDIT.

package combineservice

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/output/kitex_gen/thrifttest"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	TestVoid(ctx context.Context, callOptions ...callopt.Option) (err error)
	TestString(ctx context.Context, thing string, callOptions ...callopt.Option) (r string, err error)
	TestBool(ctx context.Context, thing bool, callOptions ...callopt.Option) (r bool, err error)
	TestByte(ctx context.Context, thing int8, callOptions ...callopt.Option) (r int8, err error)
	TestI32(ctx context.Context, thing int32, callOptions ...callopt.Option) (r int32, err error)
	TestI64(ctx context.Context, thing int64, callOptions ...callopt.Option) (r int64, err error)
	TestDouble(ctx context.Context, thing float64, callOptions ...callopt.Option) (r float64, err error)
	TestBinary(ctx context.Context, thing []byte, callOptions ...callopt.Option) (r []byte, err error)
	TestStruct(ctx context.Context, thing *thrifttest.Xtruct, callOptions ...callopt.Option) (r *thrifttest.Xtruct, err error)
	TestNest(ctx context.Context, thing *thrifttest.Xtruct2, callOptions ...callopt.Option) (r *thrifttest.Xtruct2, err error)
	TestMap(ctx context.Context, thing map[int32]int32, callOptions ...callopt.Option) (r map[int32]int32, err error)
	TestStringMap(ctx context.Context, thing map[string]string, callOptions ...callopt.Option) (r map[string]string, err error)
	TestSet(ctx context.Context, thing []int32, callOptions ...callopt.Option) (r []int32, err error)
	TestList(ctx context.Context, thing []int32, callOptions ...callopt.Option) (r []int32, err error)
	TestEnum(ctx context.Context, thing thrifttest.Numberz, callOptions ...callopt.Option) (r thrifttest.Numberz, err error)
	TestTypedef(ctx context.Context, thing thrifttest.UserId, callOptions ...callopt.Option) (r thrifttest.UserId, err error)
	TestMapMap(ctx context.Context, hello int32, callOptions ...callopt.Option) (r map[int32]map[int32]int32, err error)
	TestInsanity(ctx context.Context, argument *thrifttest.Insanity, callOptions ...callopt.Option) (r map[thrifttest.UserId]map[thrifttest.Numberz]*thrifttest.Insanity, err error)
	TestMulti(ctx context.Context, arg0 int8, arg1 int32, arg2 int64, arg3 map[int16]string, arg4 thrifttest.Numberz, arg5 thrifttest.UserId, callOptions ...callopt.Option) (r *thrifttest.Xtruct, err error)
	TestException(ctx context.Context, arg string, callOptions ...callopt.Option) (err error)
	TestMultiException(ctx context.Context, arg0 string, arg1 string, callOptions ...callopt.Option) (r *thrifttest.Xtruct, err error)
	TestOneway(ctx context.Context, secondsToSleep int32, callOptions ...callopt.Option) (err error)
	SecondtestString(ctx context.Context, thing string, callOptions ...callopt.Option) (r string, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kCombineServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kCombineServiceClient struct {
	*kClient
}

func (p *kCombineServiceClient) TestVoid(ctx context.Context, callOptions ...callopt.Option) (err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.TestVoid(ctx)
}

func (p *kCombineServiceClient) TestString(ctx context.Context, thing string, callOptions ...callopt.Option) (r string, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.TestString(ctx, thing)
}

func (p *kCombineServiceClient) TestBool(ctx context.Context, thing bool, callOptions ...callopt.Option) (r bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.TestBool(ctx, thing)
}

func (p *kCombineServiceClient) TestByte(ctx context.Context, thing int8, callOptions ...callopt.Option) (r int8, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.TestByte(ctx, thing)
}

func (p *kCombineServiceClient) TestI32(ctx context.Context, thing int32, callOptions ...callopt.Option) (r int32, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.TestI32(ctx, thing)
}

func (p *kCombineServiceClient) TestI64(ctx context.Context, thing int64, callOptions ...callopt.Option) (r int64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.TestI64(ctx, thing)
}

func (p *kCombineServiceClient) TestDouble(ctx context.Context, thing float64, callOptions ...callopt.Option) (r float64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.TestDouble(ctx, thing)
}

func (p *kCombineServiceClient) TestBinary(ctx context.Context, thing []byte, callOptions ...callopt.Option) (r []byte, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.TestBinary(ctx, thing)
}

func (p *kCombineServiceClient) TestStruct(ctx context.Context, thing *thrifttest.Xtruct, callOptions ...callopt.Option) (r *thrifttest.Xtruct, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.TestStruct(ctx, thing)
}

func (p *kCombineServiceClient) TestNest(ctx context.Context, thing *thrifttest.Xtruct2, callOptions ...callopt.Option) (r *thrifttest.Xtruct2, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.TestNest(ctx, thing)
}

func (p *kCombineServiceClient) TestMap(ctx context.Context, thing map[int32]int32, callOptions ...callopt.Option) (r map[int32]int32, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.TestMap(ctx, thing)
}

func (p *kCombineServiceClient) TestStringMap(ctx context.Context, thing map[string]string, callOptions ...callopt.Option) (r map[string]string, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.TestStringMap(ctx, thing)
}

func (p *kCombineServiceClient) TestSet(ctx context.Context, thing []int32, callOptions ...callopt.Option) (r []int32, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.TestSet(ctx, thing)
}

func (p *kCombineServiceClient) TestList(ctx context.Context, thing []int32, callOptions ...callopt.Option) (r []int32, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.TestList(ctx, thing)
}

func (p *kCombineServiceClient) TestEnum(ctx context.Context, thing thrifttest.Numberz, callOptions ...callopt.Option) (r thrifttest.Numberz, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.TestEnum(ctx, thing)
}

func (p *kCombineServiceClient) TestTypedef(ctx context.Context, thing thrifttest.UserId, callOptions ...callopt.Option) (r thrifttest.UserId, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.TestTypedef(ctx, thing)
}

func (p *kCombineServiceClient) TestMapMap(ctx context.Context, hello int32, callOptions ...callopt.Option) (r map[int32]map[int32]int32, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.TestMapMap(ctx, hello)
}

func (p *kCombineServiceClient) TestInsanity(ctx context.Context, argument *thrifttest.Insanity, callOptions ...callopt.Option) (r map[thrifttest.UserId]map[thrifttest.Numberz]*thrifttest.Insanity, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.TestInsanity(ctx, argument)
}

func (p *kCombineServiceClient) TestMulti(ctx context.Context, arg0 int8, arg1 int32, arg2 int64, arg3 map[int16]string, arg4 thrifttest.Numberz, arg5 thrifttest.UserId, callOptions ...callopt.Option) (r *thrifttest.Xtruct, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.TestMulti(ctx, arg0, arg1, arg2, arg3, arg4, arg5)
}

func (p *kCombineServiceClient) TestException(ctx context.Context, arg string, callOptions ...callopt.Option) (err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.TestException(ctx, arg)
}

func (p *kCombineServiceClient) TestMultiException(ctx context.Context, arg0 string, arg1 string, callOptions ...callopt.Option) (r *thrifttest.Xtruct, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.TestMultiException(ctx, arg0, arg1)
}

func (p *kCombineServiceClient) TestOneway(ctx context.Context, secondsToSleep int32, callOptions ...callopt.Option) (err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.TestOneway(ctx, secondsToSleep)
}

func (p *kCombineServiceClient) SecondtestString(ctx context.Context, thing string, callOptions ...callopt.Option) (r string, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SecondtestString(ctx, thing)
}
