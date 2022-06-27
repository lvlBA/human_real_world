package hrw

import "testing"

func TestHuman_checkMoney(t *testing.T) {
	type fields struct {
		Name     string
		Age      uint8
		Height   uint8
		Weight   uint8
		Parents  map[ParentType]*Human
		Wallets  map[WalletType]uint
		HP       uint8
		Wardrobe []*Clothes
	}
	type args struct {
		count uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Success",
			fields: fields{
				Name:    "Andrey",
				Age:     40,
				Height:  135,
				Weight:  182,
				Parents: nil,
				Wallets: map[WalletType]uint{
					WalletTypeCash:       1000,
					WalletTypeCreditCard: 10000,
				},
				HP:       8,
				Wardrobe: nil,
			},
			args: args{
				count: 100,
			},
			want: true,
		},
		{
			name: "Failed",
			fields: fields{
				Name:    "Andrey",
				Age:     40,
				Height:  135,
				Weight:  182,
				Parents: nil,
				Wallets: map[WalletType]uint{
					WalletTypeCash:       1,
					WalletTypeCreditCard: 1,
				},
				HP:       8,
				Wardrobe: nil,
			},
			args: args{
				count: 100,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Human{
				Name:     tt.fields.Name,
				Age:      tt.fields.Age,
				Height:   tt.fields.Height,
				Weight:   tt.fields.Weight,
				Parents:  tt.fields.Parents,
				Wallets:  tt.fields.Wallets,
				HP:       tt.fields.HP,
				Wardrobe: tt.fields.Wardrobe,
			}
			if got := h.checkMoney(tt.args.count); got != tt.want {
				t.Errorf("checkMoney() = %v, want %v", got, tt.want)
			}
		})
	}
}
