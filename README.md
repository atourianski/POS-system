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
root@contID# mysql -u root -ppasswd
mysql > CREATE DATABASE banya;
```
Hop in the container, enter passwd on prompt. Start creating! Once ready to share updates...

## Dump & export
```bash
docker exec banya-mysql mysqldump banya -uroot -ppasswd > dump.sql
```
The file `dump.sql` is now on your local machine. Commit it and/or share it.

## Import the dump
On another host, import the changes.

```bash
docker exec -it banya-mysql bash
root@contID# mysql -u root -ppasswd
mysql > CREATE DATABASE banya;
mysql > exit
root@contID# mysql -u root -ppasswd banya < dump.sql
```
The new/updated tables should be loaded!

### Resources
- https://hub.docker.com/_/mysql/
- http://www.luiselizondo.net/a-tutorial-on-how-to-use-mysql-with-docker/
- https://www.digitalocean.com/community/tutorials/a-basic-mysql-tutorial
- http://www.nkode.io/2014/09/12/easymysql.html

