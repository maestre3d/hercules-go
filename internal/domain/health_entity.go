package domain

import (
	"errors"
	"strconv"
	"strings"
)

const (
	Sedentary = iota
	LowActive
	Moderate
	Active
	VeryActive
)

const (
	Male   = "MALE"
	Female = "FEMALE"
)

const (
	Gain = iota
	Maintain
	Loss
)

type HealthInfo struct {
	Age           uint8   `json:"age"`
	Weight        float32 `json:"weight"`
	Height        float32 `json:"height"`
	BMI           float32 `json:"bmi"`
	ActivityLevel uint8   `json:"activity_level"`
	Gender        string  `json:"gender"`
	DietType      uint8   `json:"diet_type"`
	Client        string  `json:"client"`
}

type HealthInfoAggregate struct {
	Age           string `json:"age"`
	Weight        string `json:"weight"`
	Height        string `json:"height"`
	BMI           string `json:"bmi"`
	ActivityLevel string `json:"activity_level"`
	Gender        string `json:"gender"`
	DietType      string `json:"diet_type"`
	Client        string `json:"client"`
}

func NewHealthInfo(aggregate *HealthInfoAggregate) (*HealthInfo, error) {
	age, err := strconv.ParseUint(aggregate.Age, 10, 8)
	if err != nil {
		return nil, err
	}

	weight, err := strconv.ParseFloat(aggregate.Weight, 32)
	if err != nil {
		return nil, err
	}

	height, err := strconv.ParseFloat(aggregate.Height, 32)
	if err != nil {
		return nil, err
	}

	bmi, err := strconv.ParseFloat(aggregate.BMI, 32)
	if err != nil {
		return nil, err
	}

	bmi /= 100

	actLevel, err := strconv.ParseUint(aggregate.ActivityLevel, 10, 8)
	if err != nil {
		return nil, err
	}

	dietType, err := strconv.ParseUint(aggregate.DietType, 10, 8)
	if err != nil {
		return nil, err
	}

	return &HealthInfo{
		Age:           uint8(age),
		Weight:        float32(weight),
		Height:        float32(height),
		BMI:           float32(bmi),
		ActivityLevel: uint8(actLevel),
		Gender:        strings.ToUpper(aggregate.Gender),
		DietType:      uint8(dietType),
		Client:        aggregate.Client,
	}, nil
}

func (h HealthInfo) IsValid() error {
	switch {
	case h.Age < 3 || h.Age > 100:
		return errors.New("invalid age")
	case h.Weight < 20 || h.Weight > 250:
		return errors.New("invalid weight")
	case h.Height < 130 || h.Height > 250:
		return errors.New("invalid height")
	case h.ActivityLevel > 4:
		return errors.New("activity level not valid [0,4)")
	case h.DietType > 2:
		return errors.New("diet type not valid [0,2)")
	case h.Gender != Male && h.Gender != Female:
		return errors.New("invalid gender [male, female)")
	}

	return nil
}
