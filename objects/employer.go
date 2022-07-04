package objects

type WorkTime struct {
	Hours   uint8
	Minutes uint8
}

type Employer struct {
	Name              string
	WorkingHoursBegin WorkTime
	WorkingHoursEnd   WorkTime
	BankAccount       uint
}
