package repository

import "itz_JD_service/models"

type AppointmentsRepository interface {
	GetActiveAppointments() ([]models.Appointment, error)
	GetActiveAppointmentsByExpertID(id int) ([]models.Appointment, error)
	GetAllAppointments() ([]models.Appointment, error)
	GetAllAppointmentsByCustomerEmail(email string) ([]models.Appointment, error)
	GetAppointmentByID(id int) (models.Appointment, error)
	IsTimeSlotAvailable(expert_id int, start_time string, end_time string) (bool, error)
	Close()
}

var appointment_repo_implentation AppointmentsRepository

func SetAppointmentsRepository(repo AppointmentsRepository) {
	appointment_repo_implentation = repo
}

func GetActiveAppointments() ([]models.Appointment, error) {
	return appointment_repo_implentation.GetActiveAppointments()
}

func GetActiveAppointmentsByExpertID(id int) ([]models.Appointment, error) {
	return appointment_repo_implentation.GetActiveAppointmentsByExpertID(id)
}

func GetAllAppointments() ([]models.Appointment, error) {
	return appointment_repo_implentation.GetAllAppointments()
}

func GetAllAppointmentsByCustomerEmail(email string) ([]models.Appointment, error) {
	return appointment_repo_implentation.GetAllAppointmentsByCustomerEmail(email)
}

func GetAppointmentByID(id int) (models.Appointment, error) {
	return appointment_repo_implentation.GetAppointmentByID(id)
}

func IsTimeSlotAvailable(expert_id int, start_time string, end_time string) (bool, error) {
	return appointment_repo_implentation.IsTimeSlotAvailable(expert_id, start_time, end_time)
}

func Close() {
	appointment_repo_implentation.Close()
}

// Path: repository/appointments.go
