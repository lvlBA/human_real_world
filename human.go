package hrw

const (
	ParentTypeUnknown ParentType = iota
	ParentTypeFather
	ParentTypeMother
)

type ParentType uint8

const (
	WalletTypeUnknown WalletType = iota
	WalletTypeCash
	WalletTypeCreditCard
)

type WalletType uint8

// Human - describe a person in the real world
type Human struct {
	Name     string                // Name - main name of a person
	Age      uint8                 // Age - how old is a person
	Height   uint8                 // Height - what height of a person
	Weight   uint8                 // Weight - what weight of the person
	Parents  map[ParentType]*Human // Parents - amounts of parents
	Wallets  map[WalletType]uint   // Wallets - storage all wallets of a person
	HP       uint8                 // HP - Health Point
	Wardrobe []*Clothes
}

func (h Human) DoJob() error {
	return nil
}

func (h Human) DemandShopping(goods *GoodsCard) bool {
	switch {
	case !h.checkGoods(goods.Type):
		return false
	case !h.checkMoney(goods.Price):
		return false
	}

	return true
}

func (h Human) checkMoney(count uint) bool {
	for _, v := range h.Wallets {
		if v >= count {
			return true
		}
	}

	return false
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

func (h Human) checkClothes() bool {
	for i := range h.Wardrobe {
		if h.Wardrobe[i].Condition > 0 {
			return true
		}
	}

	return false
}
