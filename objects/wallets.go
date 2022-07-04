package objects

const (
	WalletTypeUnknown WalletType = iota
	WalletTypeCash
	WalletTypeCreditCard
)

type WalletType uint8
