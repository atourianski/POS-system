# POS-system

```bash
docker run --name banya-mysql -e MYSQL_ROOT_PASSWORD=passwd -d mysql
docker build -t banya-app .
docker run --name banya-pos-app --link banya-mysql:mysql -d banya-app
docker ps -a
```
