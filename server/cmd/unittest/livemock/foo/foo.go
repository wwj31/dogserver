package foo

//go:generate mockgen -package mock -source foo.go -destination=../mock/foo_mock.go

// Life 人生
type Life interface {
	// GoodGoodStudy 好好学习
	GoodGoodStudy(money int64) error
	// BuyHouse 买房
	BuyHouse(money int64) error
	// Marry 结婚
	Marry(money int64) error
}

// Person 普通人
type Person struct {
	life Life
}

func New(lf Life) *Person {
	return &Person{
		life: lf,
	}
}

// Live 活着
func (p *Person) Live(money1, money2, money3 int64) error {
	if err := p.life.GoodGoodStudy(money1); err != nil {
		return err
	}
	if err := p.life.BuyHouse(money2); err != nil {
		return err
	}
	if err := p.life.Marry(money3); err != nil {
		return err
	}
	return nil
}
