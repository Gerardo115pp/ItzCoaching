from datetime import datetime, timedelta
from models import Appointment
import repository

def cacheAppointment(appointment: Appointment) -> tuple[str, Exception]:
    err: Exception = None
    
    if appointment.duration > timedelta(hours=1):
        return None, Exception("Appointment duration cannot be greater than 1 hour")
    
    first_appointment, err = repository.appointments.getAppointmentByCustomerExpert(appointment.customer_email, appointment.expert)
    if not err and first_appointment:
        return None, Exception("Appointment already exists")
    
   
    return  repository.redis_cache.setPendingAppointment(appointment)

def isAvailableInCachedAppointments(appointment: Appointment) -> bool:
    # Check availability in redis cache
    
    all_cached_appointments, err = repository.redis_cache.getAllPendingAppointments()
    if err:
        print(f"Error getting all cached appointments: {err}")
        return False
    
    is_available = not(any([appointment.overlaps(other) for other in all_cached_appointments]))  # check that is already in cache
    return is_available

def isAvailableInDatabase(appointment: Appointment) -> bool:
    # Check availability in database
    return repository.appointments.verifyAvailability(appointment)

def isAvailable(appointment: Appointment) -> bool:
    # Check availability in redis cache
    is_available = isAvailableInCachedAppointments(appointment)
    if is_available:
        # Check availability in database
        is_available = isAvailableInDatabase(appointment)
    return is_available
