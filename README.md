# POS-system

## Start it up
```bash
docker run --name banya-mysql -e MYSQL_ROOT_PASSWORD=passwd -d mysql
docker build -t banya-app .
docker run --name banya-pos-app --link banya-mysql:mysql -d banya-app
docker ps -a
```

## Create the DB
```bash
docker exec -it banya-mysql bash
stuff# mysql -u root -p
mysql > CREATE DATABASE banya;
> exit
stuff# exit
```
Hop in the container, enter passwd on prompt. Start creating!

## Dump & export
```bash
docker exec banya-mysql mysqldump banya -uroot -ppasswd > dump.sql
```

## Import the dump
todo

### Resources
- https://hub.docker.com/_/mysql/
- http://www.luiselizondo.net/a-tutorial-on-how-to-use-mysql-with-docker/
- https://www.digitalocean.com/community/tutorials/a-basic-mysql-tutorial

