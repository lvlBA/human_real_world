package hrw

type Driver struct {
	Humaner
}

func (d *Driver) Drive() bool {
	return true
}

type Healer struct {
	Humaner
}

func (h *Healer) Heal(hp uint8) (uint8, bool) {
	if hp < 5 {
		hp += 15
	}

	return hp, true
}
