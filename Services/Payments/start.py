import Config as service_config
from flask_cors import CORS
from handlers.appointments import appointments_blueprint
from handlers.payments import payments_blueprint
from flask import Flask
import repository
import database
import os

def create_app():
    app = Flask(__name__)
    CORS(app)
    app.register_blueprint(appointments_blueprint, url_prefix='/appointments')
    app.register_blueprint(payments_blueprint, url_prefix='/payments')
    
    print(os.path.dirname(os.path.abspath(__file__)))
    
    # set redis cache
    redis_cache = database.createRedisCache()
    repository.redis_cache.setCache(redis_cache)
    
    # set appointments repository
    appointments_repository = database.createAppointmentsRepository()
    repository.appointments.setAppointmentRepository(appointments_repository)
    
    # set experts repository
    experts_repository = database.createExpertsRepository()
    repository.experts.setExpertsRepository(experts_repository)
    
    # set payments repository
    payments_repository = database.createPaymentsRepository()
    repository.payments.setPaymentsRepository(payments_repository)
    
    return app

if __name__ == '__main__':
    app = create_app()
    app.run(debug=True, host=service_config.SERVER_HOST, port=service_config.SERVER_PORT) 
