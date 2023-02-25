import Config as service_config
from flask_cors import CORS
from handlers.appointments import appointments_blueprint
from flask import Flask
import repository
import database

def create_app():
    app = Flask(__name__)
    CORS(app)
    app.register_blueprint(appointments_blueprint, url_prefix='/appointments')
    
    # set redis cache
    redis_cache = database.createRedisCache()
    repository.redis_cache.setCache(redis_cache)
    
    return app

if __name__ == '__main__':
    app = create_app()
    app.run(debug=True, host=service_config.SERVER_HOST, port=service_config.SERVER_PORT) 
