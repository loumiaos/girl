package msg

type SyncPlayers struct {
	ID          int
	Gold        int64
	IpAddr      string
	HeadIconUrl string
	Sex         int
	NickName    string

	Seat  int
	State int
}

type R_C_JoinRoom struct {
	RoomId   int
	Seat     int
	UserInfo SyncPlayers
}

type R_C_SyncPlayers struct {
	Players []SyncPlayers
}

type C_R_SitDown struct {
}

type R_C_PlayerStatus struct {
	UserId int
	State  int
}

type C_R_Ready struct {
}

type R_C_Ready struct {
	UserId int
}
