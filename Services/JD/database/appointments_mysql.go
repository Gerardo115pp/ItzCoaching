package database

import (
	"database/sql"
	"itz_JD_service/models"
)

type AppointmentsDB struct {
	MysqlRepository
}

func NewAppointmentRepository() (*AppointmentsDB, error) {
	db, err := sql.Open("mysql", createDSN())
	if err != nil {
		return nil, err
	}

	new_appointment_storage := AppointmentsDB{MysqlRepository{db: db}}

	return &new_appointment_storage, nil
}

func (appointment_repo *AppointmentsDB) GetActiveAppointments() ([]models.Appointment, error) {
	rows, err := appointment_repo.db.Query("SELECT id, expert, customer_email, start, duration, status, created_at FROM `appointments` WHERE `start` >= CONVERT_TZ(NOW(), @@session.time_zone, '+00:00');")
	if err != nil {
		return nil, err
	}

	var appointments []models.Appointment
	for rows.Next() {
		var appointment models.Appointment
		err = rows.Scan(&appointment.ID, &appointment.Expert, &appointment.Customer_email, &appointment.Utc_start, &appointment.Duration, &appointment.Status, &appointment.Created_at)
		if err != nil {
			return nil, err
		}

		appointments = append(appointments, appointment)
	}

	return appointments, nil
}

func (appointment_repo *AppointmentsDB) GetActiveAppointmentsByExpertID(id int) ([]models.Appointment, error) {
	stmt, err := appointment_repo.db.Prepare("SELECT id, expert, customer_email, start, duration, status, created_at FROM `appointments` WHERE `start` >= CONVERT_TZ(NOW(), @@session.time_zone, '+00:00') AND expert = ?")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(id)

	var appointments []models.Appointment
	for rows.Next() {
		var appointment models.Appointment
		err = rows.Scan(&appointment.ID, &appointment.Expert, &appointment.Customer_email, &appointment.Utc_start, &appointment.Duration, &appointment.Status, &appointment.Created_at)
		if err != nil {
			return nil, err
		}

		appointments = append(appointments, appointment)
	}

	return appointments, nil
}

func (appointment_repo *AppointmentsDB) GetAllAppointments() ([]models.Appointment, error) {
	rows, err := appointment_repo.db.Query("SELECT id, expert, customer_email, start, duration, status, created_at FROM `appointments`")
	if err != nil {
		return nil, err
	}

	var appointments []models.Appointment
	for rows.Next() {
		var appointment models.Appointment
		err = rows.Scan(&appointment.ID, &appointment.Expert, &appointment.Customer_email, &appointment.Utc_start, &appointment.Duration, &appointment.Status, &appointment.Created_at)
		if err != nil {
			return nil, err
		}

		appointments = append(appointments, appointment)
	}

	return appointments, nil
}

func (appointment_repo *AppointmentsDB) GetAllAppointmentsByCustomerEmail(email string) ([]models.Appointment, error) {
	stmt, err := appointment_repo.db.Prepare("SELECT id, expert, customer_email, start, duration, status, created_at FROM `appointments` WHERE customer_email = ?")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(email)

	var appointments []models.Appointment
	for rows.Next() {
		var appointment models.Appointment
		err = rows.Scan(&appointment.ID, &appointment.Expert, &appointment.Customer_email, &appointment.Utc_start, &appointment.Duration, &appointment.Status, &appointment.Created_at)
		if err != nil {
			return nil, err
		}

		appointments = append(appointments, appointment)
	}

	return appointments, nil
}

func (appointment_repo *AppointmentsDB) GetAppointmentByID(id int) (models.Appointment, error) {
	var appointment models.Appointment
	err := appointment_repo.db.QueryRow("SELECT id, expert, customer_email, start, duration, status, created_at FROM `appointments` WHERE id = ?", id).Scan(&appointment.ID, &appointment.Expert, &appointment.Customer_email, &appointment.Utc_start, &appointment.Duration, &appointment.Status, &appointment.Created_at)
	if err != nil {
		return models.Appointment{}, err
	}

	return appointment, nil
}

func (appointment_repo *AppointmentsDB) IsTimeSlotAvailable(expert_id int, start_time string, end_time string) (bool, error) {
	var count int
	err := appointment_repo.db.QueryRow("SELECT COUNT(*) FROM `appointments` WHERE expert = ? AND start >= ? AND start <= ?", expert_id, start_time, end_time).Scan(&count)
	if err != nil {
		return false, err
	}

	return count == 0, nil
}

func (appointment_repo *AppointmentsDB) Close() {
	appointment_repo.db.Close()
}
