# 搭建本地开发环境

## go
略

proxy

## docker
docker desktop
PostgreSQL
```
docker run --name postgres  -e PGDATA=/var/lib/postgresql/data/pgdata -v /Users/seven/data:/var/lib/postgresql/data/pgdata   -e POSTGRES_PASSWORD=link233 -d postgres
```
redis
```shell script
# 
```


## cosmtrek/air

```
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
```


## postman 