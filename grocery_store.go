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
const (
	stringTypeOfGoodsUnknown = "UNKNOWN"
	stringTypeGoodsEatable   = "EATABLE"
	stringTypeGoodsClothes   = "CLOTHES"
)

var resolveTypeGoodsToString = map[TypeGoods]string{
	TypeGoodsEatable: stringTypeGoodsEatable,
	TypeGoodsClothes: stringTypeGoodsClothes,
}

type TypeGoods uint

func (t TypeGoods) String() string {
	str, exist := resolveTypeGoodsToString[t]
	if exist {
		return str
	}
	return stringTypeOfGoodsUnknown
}

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
	BankAccount       uint
}

func (s *GroceryStore) removeGoods(gc *GoodsCard) error {
	gs, exists := s.Warehouse[gc.Type]
	if !exists {
		return fmt.Errorf("type of good %s is not exist", gc.Type)
	}
	for i := range gs {
		if gs[i] == gc {
			s.Warehouse[gc.Type] = append(gs[:i], gs[i+1:]...) // gs[2 : 10] взять из массива срез  от второго инджкса включительно по 10 индекс не включиттельно

			return nil
		}
	}

	return fmt.Errorf("card for goods %s is not found", gc.Brand)
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
	for _, gs := range basket {
		if err := human.BuyGoods(gs.Price, gs.Goods); err == nil {
			s.BankAccount += gs.Price
			if err := s.removeGoods(gs); err != nil {
				// show logs
				fmt.Printf("failed to calc goods at warehouse: %v\n", err)
			}
		}
	}

	fmt.Printf("Bye bye %s\n", human.Name)

	return nil
}
