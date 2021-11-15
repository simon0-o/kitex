package main

import (
	"context"

	"github.com/cloudwego/kitex/output/kitex_gen/thrifttest"
)

// CombineServiceImpl implements the last service interface defined in the IDL.
type CombineServiceImpl struct{}

// TestVoid implements the ThriftTestImpl interface.
func (s *CombineServiceImpl) TestVoid(ctx context.Context) (err error) {
	// TODO: Your code here...
	return
}

// TestString implements the ThriftTestImpl interface.
func (s *CombineServiceImpl) TestString(ctx context.Context, thing string) (resp string, err error) {
	// TODO: Your code here...
	return
}

// TestBool implements the ThriftTestImpl interface.
func (s *CombineServiceImpl) TestBool(ctx context.Context, thing bool) (resp bool, err error) {
	// TODO: Your code here...
	return
}

// TestByte implements the ThriftTestImpl interface.
func (s *CombineServiceImpl) TestByte(ctx context.Context, thing int8) (resp int8, err error) {
	// TODO: Your code here...
	return
}

// TestI32 implements the ThriftTestImpl interface.
func (s *CombineServiceImpl) TestI32(ctx context.Context, thing int32) (resp int32, err error) {
	// TODO: Your code here...
	return
}

// TestI64 implements the ThriftTestImpl interface.
func (s *CombineServiceImpl) TestI64(ctx context.Context, thing int64) (resp int64, err error) {
	// TODO: Your code here...
	return
}

// TestDouble implements the ThriftTestImpl interface.
func (s *CombineServiceImpl) TestDouble(ctx context.Context, thing float64) (resp float64, err error) {
	// TODO: Your code here...
	return
}

// TestBinary implements the ThriftTestImpl interface.
func (s *CombineServiceImpl) TestBinary(ctx context.Context, thing []byte) (resp []byte, err error) {
	// TODO: Your code here...
	return
}

// TestStruct implements the ThriftTestImpl interface.
func (s *CombineServiceImpl) TestStruct(ctx context.Context, thing *thrifttest.Xtruct) (resp *thrifttest.Xtruct, err error) {
	// TODO: Your code here...
	return
}

// TestNest implements the ThriftTestImpl interface.
func (s *CombineServiceImpl) TestNest(ctx context.Context, thing *thrifttest.Xtruct2) (resp *thrifttest.Xtruct2, err error) {
	// TODO: Your code here...
	return
}

// TestMap implements the ThriftTestImpl interface.
func (s *CombineServiceImpl) TestMap(ctx context.Context, thing map[int32]int32) (resp map[int32]int32, err error) {
	// TODO: Your code here...
	resp = thing
	return
}

// TestStringMap implements the ThriftTestImpl interface.
func (s *CombineServiceImpl) TestStringMap(ctx context.Context, thing map[string]string) (resp map[string]string, err error) {
	// TODO: Your code here...
	resp = thing
	return
}

// TestSet implements the ThriftTestImpl interface.
func (s *CombineServiceImpl) TestSet(ctx context.Context, thing []int32) (resp []int32, err error) {
	// TODO: Your code here...
	return
}

// TestList implements the ThriftTestImpl interface.
func (s *CombineServiceImpl) TestList(ctx context.Context, thing []int32) (resp []int32, err error) {
	// TODO: Your code here...
	return
}

// TestEnum implements the ThriftTestImpl interface.
func (s *CombineServiceImpl) TestEnum(ctx context.Context, thing thrifttest.Numberz) (resp thrifttest.Numberz, err error) {
	// TODO: Your code here...
	return
}

// TestTypedef implements the ThriftTestImpl interface.
func (s *CombineServiceImpl) TestTypedef(ctx context.Context, thing thrifttest.UserId) (resp thrifttest.UserId, err error) {
	// TODO: Your code here...
	return
}

// TestMapMap implements the ThriftTestImpl interface.
func (s *CombineServiceImpl) TestMapMap(ctx context.Context, hello int32) (resp map[int32]map[int32]int32, err error) {
	// TODO: Your code here...
	resp = map[int32]map[int32]int32{
		1:   {1: 1},
		12:  {1: 1, 12: 12},
		123: {1: 1, 12: 12, 123: 123},
	}
	return
}

// TestInsanity implements the ThriftTestImpl interface.
func (s *CombineServiceImpl) TestInsanity(ctx context.Context, argument *thrifttest.Insanity) (resp map[thrifttest.UserId]map[thrifttest.Numberz]*thrifttest.Insanity, err error) {
	// TODO: Your code here...
	return
}

// TestMulti implements the ThriftTestImpl interface.
func (s *CombineServiceImpl) TestMulti(ctx context.Context, arg0 int8, arg1 int32, arg2 int64, arg3 map[int16]string, arg4 thrifttest.Numberz, arg5 thrifttest.UserId) (resp *thrifttest.Xtruct, err error) {
	// TODO: Your code here...
	return
}

// TestException implements the ThriftTestImpl interface.
func (s *CombineServiceImpl) TestException(ctx context.Context, arg string) (err error) {
	// TODO: Your code here...
	return
}

// TestMultiException implements the ThriftTestImpl interface.
func (s *CombineServiceImpl) TestMultiException(ctx context.Context, arg0 string, arg1 string) (resp *thrifttest.Xtruct, err error) {
	// TODO: Your code here...
	return
}

// TestOneway implements the ThriftTestImpl interface.
func (s *CombineServiceImpl) TestOneway(ctx context.Context, secondsToSleep int32) (err error) {
	// TODO: Your code here...
	return
}

// SecondtestString implements the SecondServiceImpl interface.
func (s *CombineServiceImpl) SecondtestString(ctx context.Context, thing string) (resp string, err error) {
	// TODO: Your code here...
	return
}
