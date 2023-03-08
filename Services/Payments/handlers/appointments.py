from flask import Blueprint, request, redirect, make_response, jsonify
from middleware.auth import token_required
import Config as service_config
from datetime import datetime, timedelta
import repository
import workflows
import models
import json

appointments_blueprint = Blueprint('appointments', __name__)

@appointments_blueprint.route('/', methods=['POST'])
def POSTappointmentsHandler():
    json_body = request.get_json()
    if not(all(k in json_body for k in {'utc_start', 'utc_end', 'expert_id', "customer_email"})):
        return make_response(jsonify({'error': 'missing required fields'}), 400)
    
    appointment_ts = models.TimeSlot(json_body['utc_start'], json_body['utc_end'])
    appointment = models.Appointment(None, json_body['expert_id'], json_body['customer_email'], appointment_ts.starttime, appointment_ts.Duration)
    
    if not (workflows.schedulers.isAvailable(appointment)):
        print(f"Appointment not available: {appointment.utc_start}")
        error_response = {
            "human_readable": "El horario seleccionado ya no está disponible",
        }
        
        return make_response(json.dumps(error_response), 409)
    
    cache_hash, err = workflows.schedulers.cacheAppointment(appointment)
    if err:
        print(f"Error while caching appointment: {err}")
        return make_response("", 406)
    
    checkout_session = workflows.stripe_flows.createAppointmentCheckoutSession(appointment, f"{service_config.CUSTOMERS_URL}/#/appointment-success", f"{service_config.CUSTOMERS_URL}/#/appointment-failed", cache_hash)
    if not checkout_session:
        return make_response("", 500)
    
    response = jsonify({
        'session_id': checkout_session.id,
        'session_url': checkout_session.url,
        'deleteme': checkout_session
    })
    
    return response
    # return make_response(cache_hash, 200)

@appointments_blueprint.route('/appointment', methods=['GET'])
def GETappointmentsHandler():
    appointment_hash = request.args.get('hash')
    
    appointment, err = repository.redis_cache.getPendindAppointment(appointment_hash)
    if err:
        print(f"Error while getting appointment from redis cache: {err}")
        return make_response(jsonify({'error': err}), 404)
    
    response = make_response(appointment.toJson(), 200)
    response.headers.add_header("Content-Type", "application/json")
    
    return response