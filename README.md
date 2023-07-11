# simple_patient_order
simple project includes crud operations by mongoDB

## api
```
GET    /patients                取得patient列表
POST   /patients/:id/orders     新增該patient的order
GET    /patients/:id/orders     取得該patient的orders
PUT    /orders/:id              更新該筆order
DELETE /orders/:id              刪除該筆order
```

## how to run
Must connect to mongodb first then run the server by commands below.
```terminal
$ make run
```
or
```terminal
$ go run main.go
```

## preview config
```terminal
$ go run main.go config
```
---
## layers(folders) description
- cmd/root: Init dependencies like config, db, handler...

- internal/server: Use gin framework. Define middleware, api path....

- internal/handler: Layer to convert input and output to dto.

- internal/controller: Utilize repository and can handle some conversion or a little bussiness logic.

- internal/repository: Only this layer can interact with db.

- script: stores any db script including init, migration....

## how to test
Can take a look to internal/controller/patient_test.go

Different layer dependency should be mocked by add mockgen to Makefile in mock like
```
mockgen --source=./{fileName} --destination ./{fileName} --package {asSamelayerPackage}
```
then run
```terminal
$ make mock
$ make test
```
