from flask import Blueprint, request, redirect, make_response, jsonify
import repository
import workflows
import models
import json

payments_blueprint = Blueprint('payments', __name__)

@payments_blueprint.route('/confirm', methods=['POST'])
def POSTstripeConfirmHandler():
    json_body = request.get_json()
    session_id = json_body.get('session_id', None)
    if not session_id:
        print("Missing session_id")
        return make_response("", 400)
    
    checkout_session, err = workflows.stripe_flows.retriveCheckoutSession(session_id)
    if err:
        print(f"Error while retriving checkout session: {err}")
        return make_response("", 400)
    
    if not workflows.stripe_flows.isSessionPaid(checkout_session):
        print("Session not paid")
        return make_response("", 400)
    
    cache_hash = checkout_session.client_reference_id
    
    appointment, err = repository.redis_cache.getPendindAppointment(cache_hash)
    if err:
        print(f"Error while getting appointment from redis cache: {err}")
        # appointment reservation expired, retriving from metadata to check if still available
        appointment = models.Appointment.fromDict(checkout_session.metadata)
        
        is_available = workflows.schedulers.isAvailable(appointment)
        if not is_available:
            # refund payment and send apropiate response
            refound, err = workflows.stripe_flows.refoundSession(checkout_session, "Appointment reservation expired")
            if err:
                print(f"Error while refounding session: {err}. redfound will need to be done manually")
                appointment.status = models.AppointmentStatus.CANCELLED
                new_id, err = repository.appointments.insertAppointment(appointment)
                if err:
                    print(f"Error while inserting cancelled appointment: {err}")
                    
                appointment.id = new_id
                err =workflows.stripe_flows.registerUnpayment(checkout_session, appointment)
                if err:
                    print(f"Error while registering unpayment: {err}")
                
                return make_response("", 400)
            
            amount = checkout_session.amount_total / 100
            error_response = {
                "human_readable": f"El horario seleccionado ya no está disponible, se ha reembolsado el pago de ${amount}",
            }
            
            return make_response(json.dumps(error_response), 409)
    
    
    print(f"Confirming appointment: {appointment.utc_start} of {appointment.durationMinutes()} minutes")
    appointment.status = models.AppointmentStatus.PAID
    
    new_id, err = repository.appointments.insertAppointment(appointment)
    if err:
        print(f"Error while inserting appointment: {err}")
        return make_response("", 500)

    appointment.id = new_id
    
    payment, err = workflows.stripe_flows.registerPayment(checkout_session, appointment)
    if err:
        print(f"Error while registering payment: {err}")
        return make_response("", 500)
    
    # TODO: send email to customer and expert. register appointment in google calendar
    workflows.email.sendSuccessfulPaymentEmail(appointment, payment, checkout_session.id)
    if err:
        print(f"Error while sending email: {err}")
    
    return make_response("", 200)
    
    
    
    
    
    