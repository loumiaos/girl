module game

go 1.13

require (
	dbmodel v0.0.0-00010101000000-000000000000
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/jinzhu/gorm v1.9.12 // indirect
	github.com/phachon/go-logger v0.0.0-20191215032019-86e4227f71ea // indirect
	github.com/snowyyj001/loumiao v0.0.0-20200407122517-f26276dcd4f4
)

replace dbmodel => ../dbmodel
