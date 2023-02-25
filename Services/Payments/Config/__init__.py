from os import getenv, path

SERVER_HOST = getenv('SERVER_HOST')
assert SERVER_HOST != "", "SERVER_HOST is not set"

SERVER_PORT = getenv('SERVER_PORT')
assert SERVER_PORT != "", "SERVER_PORT is not set"

JWT_SECRET = getenv('JWT_SECRET')
assert JWT_SECRET != "", "JWT_SECRET is not set" # stupid copilot your becoming rather useless

DOMAIN_SECRET = getenv('DOMAIN_SECRET')
assert DOMAIN_SECRET != "", "DOMAIN_SECRET is not set"

USERS_SERVER = getenv('USERS_SERVER')
assert USERS_SERVER != "", "USERS_SERVER is not set"

AUTH_SERVER = getenv('AUTH_SERVER')
assert AUTH_SERVER != "", "AUTH_SERVER is not set"

JD_SERVER = getenv('JD_SERVER')
assert JD_SERVER != "", "JD_SERVER is not set"

MYSQL_DB = getenv('MYSQL_DB')
assert MYSQL_DB != "", "MYSQL_DB is not set"

MYSQL_PORT = getenv('MYSQL_PORT')
assert MYSQL_PORT != "", "MYSQL_PORT is not set"

MYSQL_HOST = getenv('MYSQL_HOST')
assert MYSQL_HOST != "", "MYSQL_HOST is not set"

MYSQL_USER = getenv('MYSQL_USER')
assert MYSQL_USER != "", "MYSQL_USER is not set"

MYSQL_PASSWORD = getenv('MYSQL_PASSWORD')
assert MYSQL_PASSWORD != "", "MYSQL_PASSWORD is not set"

REDIS_HOST = getenv('REDIS_HOST')
assert REDIS_HOST != "", "REDIS_URL is not set"

REDIS_PORT = getenv('REDIS_PORT')
assert REDIS_PORT != "", "REDIS_PORT is not set"

REDIS_PASSWORD = getenv('REDIS_PASSWORD')

MAIL_USERNAME = getenv('MAIL_USERNAME')
assert MAIL_USERNAME != "", "MAIL_USERNAME is not set"

MAIL_PASSWORD = getenv('MAIL_PASSWORD')
assert MAIL_PASSWORD != "", "MAIL_PASSWORD is not set"

MAIL_SERVER = getenv('MAIL_SERVER')
assert MAIL_SERVER != "", "MAIL_SERVER is not set"

MAIL_PORT = getenv('MAIL_PORT')
assert MAIL_PORT != "", "MAIL_PORT is not set"

MAIL_FROM = getenv('MAIL_FROM')
assert MAIL_FROM != "", "MAIL_FROM is not set"



