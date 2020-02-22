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
	ErrCode int
	UserID  int
}
