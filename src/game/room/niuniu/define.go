package niuniu

const (
	MAX_ROOM_RENSHU = 50 //房间最大人数
	MAX_SEAT        = 13 //座位数
	POKER_NUMBER    = 54 //扑克数量
	MIN_GAMER       = 2  //最小玩家数量
	HANDCARD_NUM    = 5  //手牌数量
)

type RoomState int

const (
	State_Idle RoomState = iota
	State_Ready
	State_Gameing
)

type FSMState int

const (
	FSM_Idle FSMState = iota
	FSM_Fapai
	FSM_QZhuang
	FSM_KaiJiang
	FSM_Bipai
	FSM_Result
)

const (
	Time_Idle     = 20 * 1000
	Time_FaPai    = 5 * 1000
	Time_QZhang   = 5 * 1000
	Time_KaiJiang = 60 * 1000
	Time_BiPai    = 30 * 1000
	Time_Result   = 10 * 1000
)
