# girl
game server
robot cleint

1.消息协议&socket
NET_PROTOCOL："PROTOBUF" or "JSON"，使用pb或者json
NET_WEBSOCKET：true or false，使用websocket或者socket
2.数据库
使用mysql、mongo、redis，自己组合决定
3.开启一个服务
loumiao.Prepare(new(gate.GateServer), "GateServer", false)
4.注册网关监听事件
loumiao.RegisterNetHandler
5.注册rpc监听事件
igo.Register
6.创建一个服务参考GameServer