from .mysql_utils import MysqlConnection, MYSQL_CONFIG
from models import Appointment, TimeSlot


class AppointmentsRepository:
    def __init__(self, config: MYSQL_CONFIG=None):
        if not config:
            config = MYSQL_CONFIG.createFromEnv()
        self.config = config
    
    def insertAppointment(self, appointment: Appointment) -> tuple[int, Exception]:
        last_id = None
        with MysqlConnection(self.config) as conn:
            cursor = conn.cursor(prepared=True)
            sql = "INSERT INTO appointments (expert, customer_email, start, duration, status, created_at) VALUES (?, ?, ?, ?, ?, ?)"
            params = (appointment.expert, appointment.customer_email, appointment.utc_start, appointment.duration.total_seconds() * 1000, appointment.status, appointment.created_at)
            cursor.execute(sql, params)
            last_id = cursor.lastrowid
            conn.commit()
        
        return last_id, None
    
    def getAppointmentByCustomerExpert(self, customer_email: str, expert: int) -> tuple[Appointment, None]:
        with MysqlConnection(self.config) as conn:
            cursor = conn.cursor(prepared=True)
            sql = "SELECT * FROM appointments WHERE customer_email = ? AND expert = ?"
            params = (customer_email, expert)
            cursor.execute(sql, params)
            result = cursor.fetchone()
            if result:
                return Appointment(*result), None
            return None, None
        
    def verifyAvailability(self, appointment: Appointment) -> bool:
        with MysqlConnection(self.config) as conn:
            cursor = conn.cursor(prepared=True)
            sql = "SELECT COUNT(*) AS `overlaps` FROM `appointments` WHERE `expert`=? AND `start` <= DATE_ADD(?, INTERVAL ? SECOND) AND DATE_ADD(`start`, INTERVAL `duration` SECOND) >= ?"
            params = (appointment.expert, appointment.utc_start, appointment.duration.total_seconds(), appointment.utc_start)
            cursor.execute(sql, params)
            result = cursor.fetchone()
            
        return result[0] == 0
        
    