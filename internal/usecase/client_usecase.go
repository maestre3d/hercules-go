package usecase

import (
	"github.com/maestre3d/hercules-go/internal/domain"
)

type ClientUseCase struct{}

func (u *ClientUseCase) Create(aggregate *domain.ClientAggregate) (*domain.ClientRoot, error) {
	c := domain.NewClient(aggregate.Name, aggregate.LastName)
	err := c.IsValid()
	if err != nil {
		return nil, err
	}

	h, err := domain.NewHealthInfo(&domain.HealthInfoAggregate{
		Age:           aggregate.Age,
		Weight:        aggregate.Weight,
		Height:        aggregate.Height,
		BMI:           aggregate.BMI,
		ActivityLevel: aggregate.ActivityLevel,
		Gender:        aggregate.Gender,
		DietType:      aggregate.DietType,
		Client:        c.ID,
	})
	err = h.IsValid()
	if err != nil {
		return nil, err
	}

	insight := &domain.HealthInsights{
		Client: c.ID,
	}

	insight.CalculateBMR(h.Weight, h.BMI)
	insight.CalculateRDI(h.Gender, h.Age, h.ActivityLevel, h.Weight)
	insight.SetDietType(h.DietType)
	insight.CalculateProtein(h.Weight)
	insight.CalculateFat()
	insight.CalculateCarbs()

	err = insight.IsValid()
	if err != nil {
		return nil, err
	}

	return &domain.ClientRoot{
		ID:             c.ID,
		Name:           c.Name,
		LastName:       c.LastName,
		CreateTime:     c.CreateTime,
		UpdateTime:     c.UpdateTime,
		Active:         c.Active,
		HealthInfo:     h,
		HealthInsights: insight,
	}, nil
}
