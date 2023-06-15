package weather

import "time"

type Weather struct {
	ID          uint   `gorm:"primaryKey"`
	WindStatus  string `gorm:"not null"`
	WaterStatus string `gorm:"not null"`
	Wind        int    `gorm:"not null"`
	Water       int    `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (w *Weather) GenerateStatus() {
	switch {
	case w.Water < 6:
		w.WaterStatus = "aman"
	case w.Water <= 8:
		w.WaterStatus = "siaga"
	default:
		w.WaterStatus = "bahaya"
	}

	switch {
	case w.Wind < 7:
		w.WindStatus = "aman"
	case w.Wind <= 15:
		w.WindStatus = "siaga"
	default:
		w.WindStatus = "bahaya"
	}
}

func (w *Weather) UpdateWeather(water int, wind int) {
	w.Water = water
	w.Wind = wind
	w.UpdatedAt = time.Now()
}
