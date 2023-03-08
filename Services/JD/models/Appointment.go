package models

import (
	"encoding/json"
	"time"
)

type AppointmentStatus string

const (
	// Appointment types
	appointment_status_UNPAID    AppointmentStatus = "unpaid"
	appointment_status_PAID      AppointmentStatus = "paid"
	appointment_status_FINALIZED AppointmentStatus = "finalized"
	appointment_status_CANCELLED AppointmentStatus = "cancelled"
)

type TimeSlot struct {
	Start time.Time `json:"start_time"`
	End   time.Time `json:"end_time"`
}

func (ts TimeSlot) overlaps(other TimeSlot) bool {
	return ts.Start.Before(other.End) && ts.End.After(other.Start)
}

type Appointment struct {
	ID             int               `json:"id"`
	Expert         int               `json:"expert"`
	Customer_email string            `json:"customer_email"`
	Utc_start      time.Time         `json:"utc_start"`
	Duration       int               `json:"duration"` // milliseconds
	Status         AppointmentStatus `json:"status"`
	Created_at     time.Time         `json:"created_at"`
}

func (appointment *Appointment) toTS() *TimeSlot {
	return &TimeSlot{
		Start: appointment.Utc_start,
		End:   appointment.Utc_start.Add(time.Duration(appointment.Duration) * time.Millisecond),
	}
}

func (appointment *Appointment) overlaps(other *Appointment) bool {
	return appointment.toTS().overlaps(*other.toTS())
}

func (appointment Appointment) Minutes() int {
	return appointment.Duration / 60000
}

func (appointment Appointment) MarshalJSON() ([]byte, error) {
	aux := map[string]interface{}{
		"id":             appointment.ID,
		"expert":         appointment.Expert,
		"customer_email": appointment.Customer_email,
		"utc_start":      appointment.Utc_start.Format("2006-01-02 15:04:05"),
		"duration":       appointment.Duration,
		"status":         appointment.Status,
		"created_at":     appointment.Created_at.Format("2006-01-02 15:04:05"),
	}

	return json.Marshal(aux)
}
