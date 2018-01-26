# MySQL on docker
### Some usefull commands
```
docker pull mysql
docker run --name first -p 3306:3306 -e MYSQL\_ROOT\_PASSWORD=123456 -d mysql
mysql -h127.0.0.1 -P3306 -uroot -p123456
docker exec -it first bash
```
```
# using dir on host system
docker run --name first -p 3306:3306 -v ./data:/var/lib/mysql -e MYSQL\_ROOT\_PASSWORD=123456 -d mysql
```
