FROM mysql/mysql-server:latest

COPY ../scripts/favourites.sql /docker-entrypoint-initdb.d

ENV MYSQL_ROOT_PASSWORD=root


#Docker MySQL by default doesn't allow connection from outside