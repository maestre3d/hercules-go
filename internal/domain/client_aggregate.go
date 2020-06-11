package domain

import "time"

type ClientAggregate struct {
	Name          string `json:"name"`
	LastName      string `json:"last_name"`
	Age           string `json:"age"`
	Weight        string `json:"weight"`
	Height        string `json:"height"`
	BMI           string `json:"bmi"`
	ActivityLevel string `json:"activity_level"`
	Gender        string `json:"gender"`
	DietType      string `json:"diet_type"`
}

type ClientRoot struct {
	ID             string          `json:"id"`
	Name           string          `json:"name"`
	LastName       string          `json:"last_name"`
	CreateTime     time.Time       `json:"create_time"`
	UpdateTime     time.Time       `json:"update_time"`
	Active         bool            `json:"active"`
	HealthInfo     *HealthInfo     `json:"health_info"`
	HealthInsights *HealthInsights `json:"health_insights"`
}
