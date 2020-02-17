package dbmodel

var (
	Models []interface{}
)

type User struct {
	ID           int    `gorm:"primary_key";AUTO_INCREMENT`
	Account      string `gorm:"type:varchar(255);not null;unique;index:account"`
	Gold         int64
	Coin         int64
	Money        int64
	IpAddr       string `gorm:"type:varchar(40)"`
	HeadIconUrl  string `gorm:"type:varchar(255)"`
	Passwd       string
	Channel      int
	LoginType    int
	UnderWrite   string `gorm:"type:varchar(255)"`
	Sex          int
	ActiveFlag   int64
	NickName     string `gorm:"type:varchar(255)"`
	RegisterTime int64
	LoginTime    int64
	LogoutTime   int64
}

type GameCfg struct {
	ID          int `gorm:"primary_key";AUTO_INCREMENT`
	GameId      int
	GameName    string `gorm:"type:char(64)"`
	ServiceName string `gorm:"type:char(64)"`
	RoomNumber  int
	GameRule    string `gorm:"type:varchar(512)"`
}

func init() {
	Models = append(Models, &User{})
	Models = append(Models, &GameCfg{})
}
