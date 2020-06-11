package main

import (
	"log"

	"github.com/maestre3d/hercules-go/internal/domain"
	"github.com/maestre3d/hercules-go/internal/usecase"
)

func main() {
	usecase := new(usecase.ClientUseCase)

	aggregate := &domain.ClientAggregate{
		Name:          "Elon",
		LastName:      "Musk",
		Age:           "27",
		Weight:        "68",
		Height:        "170",
		BMI:           "29.10",
		ActivityLevel: "2",
		Gender:        "FEMALE",
		DietType:      "0",
	}

	c, err := usecase.Create(aggregate)
	if err != nil {
		panic(err)
	}

	log.Printf("%+v", c)
	log.Print(c.HealthInsights.BMR)
	log.Print(c.HealthInsights.RDI)
	log.Print(c.HealthInsights.ProteinRate)
	log.Print(c.HealthInsights.FatRate)
	log.Print(c.HealthInsights.CarbohydrateRate)
}
