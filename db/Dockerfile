FROM postgres
COPY init.sql /docker-entrypoint-initdb.d
VOLUME [ "/backup" ]
EXPOSE 5432