FROM mysql/mysql-server:latest

COPY /scripts/favourites.sql /docker-entrypoint-initdb.d

ENV MYSQL_ROOT_PASSWORD=sameer123
ENV HOSTNAME=docker-database
ENV MYSQL_ROOT_HOST=%

#Docker MySQL by default doesn't allow connection from outside