package core

import "time"

// LogEntry is what is expected as input
type LogEntry struct {
	TimeStampEntry time.Time
	TimeStampLabel TimeLabel
	Item           ItemType
	Weight         AmountType
	Water          AmountType
}

// Ingredient ties Item names to the nutritional info
type Ingredient struct {
	Name      string
	Subname   string
	Calories  float32
	Serving   float32
	Carbs     float32
	Fiber     float32
	Protein   float32
	Fat       float32
	Sodium    float32
	Magnesium float32
	Potassium float32
}

// AmountType holds a value and a unit (ie, 500 grams)
type AmountType struct {
	Value float32
	Unit  string
}

// ItemType defines a food, can contain other items
type ItemType struct {
	Name     string
	Servings Amount
	Contains []ItemType
}

// TimeLabel - not sure if I'll use it, but for use case where I want to group food by time of day
type TimeLabel int

const (
	Breakfast = iota,
		Lunch
	SnackAM
	SnackPM
	Supper
	SnackEve
)
