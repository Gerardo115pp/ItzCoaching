from . import Appointment as ap_module
from . import Payment as pay_module
from . import TimeSlot as ts_module
from . import Expert as ex_module
from inspect import getfullargspec

def _allArgsPresent(obj: object, args: list[str]) -> bool:
    all_args = getfullargspec(obj).args
    are_all_present = True
    for arg in all_args:
        if arg not in args and arg != "self":
            print(f"Missing argument: {arg} in {obj.__name__}")
            are_all_present = False
            break
        
    return are_all_present

ap_module.allArgsPresent = _allArgsPresent
Appointment = ap_module.Appointment
AppointmentStatus = ap_module.AppointmentStatus
AppointmentConfirmation = ap_module.AppointmentConfirmation

pay_module.allArgsPresent = _allArgsPresent
Payment = pay_module.Payment
PaymentStatus = pay_module.PaymentStatus
createEmptyPayment = pay_module.createEmptyPayment

ts_module.allArgsPresent = _allArgsPresent
TimeSlot = ts_module.TimeSlot

ex_module.allArgsPresent = _allArgsPresent
ExpertPersonalData = ex_module.ExpertPersonalData