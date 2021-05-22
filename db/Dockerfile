FROM postgres 
ENV POSTGRES_PASSWORD 12345 
ENV POSTGRES_DB resume02
ENV POSTGRES_USER postgres
COPY /migration/resume02.sql /docker-entrypoint-initdb.d/
