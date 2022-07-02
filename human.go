package hrw

import "errors"

type Humaner interface {
	getSalary() uint8
}

const (
	ParentTypeUnknown ParentType = iota
	ParentTypeFather
	ParentTypeMother
)

type ParentType uint8

// Human - describe a person in the real world
type Human struct {
	Name     string                // Name - main name of a person
	Age      uint8                 // Age - how old is a person
	Height   uint8                 // Height - what height of a person
	Weight   uint8                 // Weight - what weight of the person
	Parents  map[ParentType]*Human // Parents - amounts of parents
	Wallets  map[WalletType]uint   // Wallets - storage all wallets of a person
	HP       uint8                 // HP - Health Point
	Wardrobe []*Clothes            // Wardrobe - place where person store clothes
	JobPlace *Employer             //	JobPlace - place where person is working
}

func (h Human) DoJob() error {
	return nil
}

func (h Human) DemandShopping(goods *GoodsCard) bool {
	if !h.checkGoods(goods.Type) {
		return false
	}

	if ok, _ := h.checkMoney(goods.Price); !ok {
		return false
	}

	return true
}

func (h Human) checkMoney(count uint) (bool, WalletType) {
	for k, v := range h.Wallets {
		if v >= count {
			return true, k
		}
	}

	return false, WalletTypeUnknown
}

func (h Human) checkGoods(gt TypeGoods) bool {
	switch {
	case h.HP < 20 && gt != TypeGoodsEatable:
		return false
	case h.HP > 100 && gt == TypeGoodsEatable:
		return false
	case h.checkClothes() && gt == TypeGoodsClothes:
		return false
	}

	return true
}

func (h Human) checkHeals(hp uint8) bool {
	if h.HP < 10 {
		if ok, _ := h.checkMoney(100); !ok {
			return true
		}

	}

	return true

}

func (h Human) BuyGoods(price uint, goods interface{}) error {
	ok, wt := h.checkMoney(price)
	if !ok {
		return errors.New("not enough money")
	}

	switch g := goods.(type) {
	case *Eatable:
		h.HP += g.Calories
	case *Clothes:
		h.Wardrobe = append(h.Wardrobe, g)
	default:
		return errors.New("IDK what it is")
	}
	h.Wallets[wt] -= price

	return nil
}

func (h Human) checkClothes() bool {
	for i := range h.Wardrobe {
		if h.Wardrobe[i].Condition > 0 {
			return true
		}
	}

	return false
}
