package config

//sys
const (
	OFFLINE_TIME int64 = 5 * 60 * 1000 //离线五分钟则删除内存
)

//player
const (
	PLAYER_GOLD  int64 = 1000
	PLAYER_COIN  int64 = 0
	PLAYER_MONEY int64 = 0
)

//net
const (
	NET_PROTOCOL  = "JSON"      //OR"PROTOBUF" 使用JSON协议
	NET_WEBSOCKET = true        //使用websocket
	NET_RPC_IP    = "127.0.0.1" //RPC监听地址
	NET_RPC_PORT  = 5678        //RPC监听端口
	NET_BE_CHILD  = true        //分布式结构
)

//db
const (
	MYSQL_URI     = "117.51.136.136:3306"
	MYSQL_DBNAME  = "loumiao"
	MYSQL_ACCOUNT = "root"
	MYSQL_PASS    = "123456"
)
