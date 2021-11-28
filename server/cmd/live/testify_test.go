package main

import (
	"errors"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Live1(t *testing.T) {
	patches := gomonkey.NewPatches()
	// GoodGoodStudy error
	patches.ApplyFunc(GoodGoodStudy, func(int64) error {
		return errors.New("error")
	})
	assert.Error(t, Live(100, 100, 100))
	patches.Reset()
	// BuyHouse error
	patches.ApplyFunc(GoodGoodStudy, func(int64) error {
		return nil
	})
	patches.ApplyFunc(BuyHouse, func(int64) error {
		return errors.New("error")
	})
	assert.Error(t, Live(100, 100, 100))
	patches.Reset()
	// Marry error
	patches.ApplyFunc(GoodGoodStudy, func(int64) error {
		return nil
	})
	patches.ApplyFunc(BuyHouse, func(int64) error {
		return nil
	})
	patches.ApplyFunc(Marry, func(int64) error {
		return errors.New("error")
	})
	assert.Error(t, Live(100, 100, 100))
	patches.Reset()
	// ok
	patches.ApplyFunc(GoodGoodStudy, func(int64) error {
		return nil
	})
	patches.ApplyFunc(BuyHouse, func(int64) error {
		return nil
	})
	patches.ApplyFunc(Marry, func(int64) error {
		return nil
	})
	assert.NoError(t, Live(100, 100, 100))
	patches.Reset()
}

func Test_Live2(t *testing.T) {
	patches := gomonkey.NewPatches()
	defer patches.Reset()
	output := []gomonkey.OutputCell{
		{Values: gomonkey.Params{errors.New("error")}, Times: 1},
		{Values: gomonkey.Params{nil}, Times: 3},
	}
	patches.ApplyFuncSeq(GoodGoodStudy, output)
	output = []gomonkey.OutputCell{
		{Values: gomonkey.Params{errors.New("error")}, Times: 1},
		{Values: gomonkey.Params{nil}, Times: 2},
	}
	patches.ApplyFuncSeq(BuyHouse, output)
	output = []gomonkey.OutputCell{
		{Values: gomonkey.Params{errors.New("error")}, Times: 1},
		{Values: gomonkey.Params{nil}, Times: 1},
	}
	patches.ApplyFuncSeq(Marry, output)
	// GoodGoodStudy error
	assert.Error(t, Live(100, 100, 100))
	// BuyHouse error
	assert.Error(t, Live(100, 100, 100))
	// Marry error
	assert.Error(t, Live(100, 100, 100))
	// ok
	assert.NoError(t, Live(100, 100, 100))
}
