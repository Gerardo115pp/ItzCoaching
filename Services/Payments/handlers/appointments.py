from flask import Blueprint, request, redirect, make_response, jsonify
from middleware.auth import token_required
import Config as service_config
from datetime import datetime, timedelta
import repository
import models

appointments_blueprint = Blueprint('appointments', __name__)

@appointments_blueprint.route('/', methods=['POST'])
def POSTappointmentsHandler():
    json_body = request.get_json()
    if not(all(k in json_body for k in {'utc_start', 'utc_end', 'expert_id', "customer_email"})):
        return make_response(jsonify({'error': 'missing required fields'}), 400)
    
    appointment_ts = models.TimeSlot(json_body['utc_start'], json_body['utc_end'])
    appointment = models.Appointment(None, json_body['expert_id'], json_body['customer_email'], appointment_ts.starttime, appointment_ts.Duration)
    
    #TODO: validate that the appointment is no longer than 1 hour
    if appointment.duration > timedelta(hours=1):
        return make_response("", 406)
    
    #TODO: validate that there is no other appointment for the same expert and the same customer_email, because we dont want to allow DDoS attacks
    
    
    repository.redis_cache.setPendingAppointment(appointment)
    
    response = make_response("", 201)
    return response

@appointments_blueprint.route('/appointment', methods=['GET'])
def GETappointmentsHandler():
    appointment_hash = request.args.get('hash')
    
    appointment, err = repository.redis_cache.getPendindAppointment(appointment_hash)
    if err:
        print(f"Error while getting appointment from redis cache: {err}")
        return make_response(jsonify({'error': err}), 404)
    
    return jsonify(appointment.toJson())