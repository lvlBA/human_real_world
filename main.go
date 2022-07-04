package main

import (
	"go.uber.org/zap"
	"hrw/objects"
	"math/rand"
	"time"
)

const sunrise, sunset = 8, 20

func Init() *zap.SugaredLogger {
	cfg := zap.NewProductionConfig()
	cfg.DisableStacktrace = true
	cfg.DisableCaller = true
	cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	return logger.Sugar()
}
func main() {
	logger := Init()
	Daniel := objects.Human{
		Name:    "Daniel",
		Age:     33,
		Height:  180,
		Weight:  82,
		Parents: nil,
		Wallets: map[objects.WalletType]uint{
			objects.WalletTypeCreditCard: 5500,
		},
		HP: 100,
		Wardrobe: []*objects.Clothes{
			{
				Name:      "Working robe",
				Condition: 15,
			},
		},
		JobPlace: nil,
	}

	rand.Seed(time.Now().UnixNano())

	Schlumberger := objects.Employer{
		Name: "Schlumberger",
		WorkingHoursBegin: objects.WorkTime{
			Hours:   8,
			Minutes: 0,
		},
		WorkingHoursEnd: objects.WorkTime{
			Hours:   16,
			Minutes: 0,
		},
		BankAccount: 1000000,
	}
	Daniel.JobPlace = &Schlumberger
	var sol, hour int
	for {
		logger.Infow("Now:", "hour", hour, "sol", sol)

		if hour < sunrise || hour > sunset {
			logger.Info("Time for sleeping and resting")
		} else {
			logger.Warn(Daniel.Depression())
		}
		time.Sleep(10 * time.Millisecond)
		hour++
		if hour >= 24 {
			hour = 0
			sol++
			Daniel.HP -= 10
			logger.Infof("Daniel hp is %d:", Daniel.HP)
			if Daniel.HP <= 0 {
				logger.Errorw("He is Dead", "dead time", time.Now())
				break
			}
		}
	}
}
