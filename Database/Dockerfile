FROM mariadb:10.5 

ENV MARIADB_ROOT_PASSWORD='da!nyblue'
ENV MARIADB_DATABASE=itz_coaching

COPY ./Itz_db_schema.sql /docker-entrypoint-initdb.d

EXPOSE 3306