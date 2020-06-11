package domain

import (
	"errors"
	"math"
)

type HealthInsights struct {
	BMR               float32 `json:"bmr"`
	CarbohydrateGrams float32 `json:"carbohydrate_grams"`
	FatGrams          float32 `json:"fat_grams"`
	ProteinGrams      float32 `json:"protein_grams"`
	CarbohydrateRate  int     `json:"carbohydrate_rate"`
	FatRate           int     `json:"fat_rate"`
	ProteinRate       int     `json:"protein_rate"`
	RDI               float32 `json:"rdi"`
	Client            string  `json:"client"`
}

func (h *HealthInsights) CalculateBMR(weight, bmi float32) {
	h.BMR = float32(math.Abs(float64(weight*bmi - weight)))
}

func (h *HealthInsights) CalculateRDI(gender string, age, activity uint8, weight float32) {
	switch gender {
	case Male:
		switch {
		case age == 0 || age < 3:
			h.RDI = 60.9*weight - 54
		case age == 3 || age < 10:
			h.RDI = 22.7*weight + 495
		case age == 10 || age <= 18:
			h.RDI = 17.5*weight + 651
		case age == 19 || age <= 30:
			h.RDI = 15.3*weight + 679
		case age == 31 || age <= 60:
			h.RDI = 11.6*weight + 879
		default:
			h.RDI = 13.5*weight + 487
		}
	case Female:
		switch {
		case age == 0 || age < 3:
			h.RDI = 61*weight - 51
		case age == 3 || age < 10:
			h.RDI = 22.5*weight + 499
		case age == 10 || age <= 18:
			h.RDI = 12.2*weight + 746
		case age == 19 || age <= 30:
			h.RDI = 14.7*weight + 496
		case age == 31 || age <= 60:
			h.RDI = 8.7*weight + 829
		default:
			h.RDI = 10.5*weight + 596
		}
	}

	switch activity {
	case Sedentary:
		h.RDI *= 1.2
	case LowActive:
		h.RDI *= 1.375
	case Moderate:
		h.RDI *= 1.55
	case Active:
		h.RDI *= 1.725
	case VeryActive:
		h.RDI *= 1.9
	}
}

func (h *HealthInsights) SetDietType(diet uint8) {
	switch diet {
	case Gain:
		h.RDI += 500
	case Loss:
		h.RDI -= 500
	}
}

func (h *HealthInsights) CalculateProtein(weight float32) {
	h.ProteinGrams = weight * 2
	h.ProteinRate = int(math.Round(float64(h.ProteinGrams*4/h.RDI) * 100))
}

func (h *HealthInsights) CalculateFat() {
	h.FatGrams = 0.7 * h.BMR
	h.FatRate = int(math.Round(float64(h.FatGrams*9/h.RDI) * 100))
}

func (h *HealthInsights) CalculateCarbs() {
	protein := (float32(h.ProteinRate) / 100) * h.RDI
	fat := float32(h.FatRate) / 100 * h.RDI

	h.CarbohydrateGrams = (h.RDI - (protein + fat)) / 4
	h.CarbohydrateRate = int(math.Round(float64(h.CarbohydrateGrams*4/h.RDI) * 100))
}

func (h HealthInsights) IsValid() error {
	totalRate := h.ProteinRate + h.FatRate + h.CarbohydrateRate
	if totalRate < 100 {
		return errors.New("invalid macros")
	}

	return nil
}
