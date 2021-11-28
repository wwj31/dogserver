package main

import (
	"errors"
	"math/rand"
)

// Live 活着
func Live(money1, money2, money3 int64) error {
	if err := GoodGoodStudy(money1); err != nil {
		return err
	}
	if err := BuyHouse(money2); err != nil {
		return err
	}
	if err := Marry(money3); err != nil {
		return err
	}
	return nil
}

// GoodGoodStudy 好好学习
func GoodGoodStudy(money int64) error {
	if rand.Intn(100) > 0 {
		return errors.New("error")
	}
	_ = money
	return nil
}

// BuyHouse 买房
func BuyHouse(money int64) error {
	if rand.Intn(100) > 0 {
		return errors.New("error")
	}
	_ = money
	return nil
}

// Marry 结婚
func Marry(money int64) error {
	if rand.Intn(100) > 0 {
		return errors.New("error")
	}
	_ = money
	return nil
}
