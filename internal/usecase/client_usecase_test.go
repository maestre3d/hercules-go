package usecase

import (
	"testing"

	"github.com/maestre3d/hercules-go/internal/domain"
)

func TestClientUseCase_Create(t *testing.T) {
	usecase := new(ClientUseCase)

	aggregate := &domain.ClientAggregate{
		Name:          "Elon",
		LastName:      "Musk",
		Age:           "27",
		Weight:        "68",
		Height:        "170",
		BMI:           "29.10",
		ActivityLevel: "2",
		Gender:        domain.Female,
		DietType:      "0",
	}

	c, err := usecase.Create(aggregate)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%+v", c)
}
