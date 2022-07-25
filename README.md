# simple_api
RESTful API with Golang.
## USEAGE
- 移動到資料夾下：
前言：目前沒有特別寫go的設定檔，因此有些參數是先寫死的。
````
$ cd ./simple_api
````
## build mongodb
- 設定連線通道：

````
$ docker network create mongo-network
````
- 使用官方映像檔：

````
$ docker  run -d --network mongo-network -p 27017:27017  --name example-mongo mongo:latest
````

- 設定模組需用到的資料：
	* 進入mongo DB 容器中
	````
	$ docker exec -it example-mongo mongo admin
	````
	* 創造DB
	````
db.createUser({ user:'root',pwd:'zxc12345',roles:[ { role:'userAdminAnyDatabase', db: 'admin'},"readWriteAnyDatabase"]});
	use climate;
	db.rowdata.insert({});
	````

## build go modules

- 製作映像檔：

````
$ docker build -t simple_api:v1 .
````
- 部署：

````
$ docker run  --network mongo-network -p 8080:8080  --name simple_api simple_api:v1
````


