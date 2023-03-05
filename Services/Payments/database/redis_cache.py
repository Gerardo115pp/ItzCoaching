from .redis_utils import RedisConnection, REDIS_CONFIG
from models import Appointment
from datetime import datetime, timedelta
import pickle
import hashlib

class RedisCache:
    def __init__(self, config: REDIS_CONFIG=None):
        self.config = config or REDIS_CONFIG.createFromEnv()
    
    def getPendingAppointment(self, appointment_hash: str) -> tuple[Appointment, Exception]:
        err = None
        appointment = None
        
        with RedisConnection(self.config) as conn:
            data = conn.get(appointment_hash)
            
            if not data:
                print(f"no data for {appointment_hash}")
                return None, Exception("no data for appointment")
        try:
            appointment = pickle.loads(data)
        except Exception as e:
            return None, e

        return appointment, err
        
    def setPendingAppointment(self, appointment: Appointment) -> tuple[str, Exception]:
        appointment_bytes = pickle.dumps(appointment)
        appointment_hash = hashlib.sha256(appointment_bytes).hexdigest()
        
        with RedisConnection(self.config) as conn:
            conn.set(appointment_hash, appointment_bytes, ex=240) # 4 minutes
        
        return appointment_hash, None
    
    def getAllPendingAppointments(self) -> tuple[list[Appointment], Exception]:
        all_appointments = []
        err = None
        
        with RedisConnection(self.config) as conn:
            for key in conn.scan_iter("*"):
                
                value = conn.get(key)
                
                if value.startswith(b'\x80\x04\x95'):
                    try:
                        all_appointments.append(pickle.loads(value))
                    except Exception as e:
                        err = e
                        return [], err
        
        return all_appointments, err