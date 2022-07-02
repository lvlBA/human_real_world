package hrw

import "time"

type Employer struct {
	Name              string
	WorkingHoursBegin time.Time
	WorkingHoursEnd   time.Time
	BankAccount       uint
}
