FROM postgres
RUN apt-get update && \
    apt-get install -y postgis && \
    apt-get install -y postgresql-11-postgis-2.5
COPY init.sql /docker-entrypoint-initdb.d
VOLUME [ "/backup" ]
EXPOSE 5432