package msg

type C_A_Login struct {
	AccountName string `json:"account_name"`
	Password    string
	Channel     int
	LoginType   int
	HeadIcon    string
	Sex         int
	NickName    string
}

type A_C_Login struct {
	ErrorStr string
	UserID   int
}

type C_S_Login struct {
	UserID int
}

type S_C_Login struct {
	ErrorStr    string
	UserID      int
	Gold        int64
	Coin        int64
	Money       int64
	HeadIconUrl string
	UnderWrite  string
	Sex         int
	ActiveFlag  int64
	NickName    string
}
