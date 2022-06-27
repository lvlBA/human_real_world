package hrw

import (
	"fmt"
	"time"
)

const (
	TypeGoodsUnknown TypeGoods = iota
	TypeGoodsEatable
	TypeGoodsClothes
)

type TypeGoods uint

type GoodsCard struct {
	Brand string
	Type  TypeGoods
	Price uint
	Goods interface{}
}

type GroceryStore struct {
	Name              string
	WorkingHoursBegin time.Time
	WorkingHoursEnd   time.Time
	Warehouse         map[TypeGoods][]*GoodsCard
}

func (s GroceryStore) GetVisitor(human Human) error {
	now := time.Now()
	if now.Before(s.WorkingHoursBegin) || now.After(s.WorkingHoursEnd) {
		return fmt.Errorf("shop is closed, working hours from %s to %s", s.WorkingHoursBegin, s.WorkingHoursEnd)
	}

	var basket []*GoodsCard
	for _, gs := range s.Warehouse {
		for i := range gs {
			if human.DemandShopping(gs[i]) {
				basket = append(basket, gs[i])
			}
		}
	}

	//Payment issues
	// Say Good bye to the client
	return nil
}
