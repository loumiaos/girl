package msg

type NN_RC_TableInfo struct {
	State     int
	LeftTime  int
	BaseScore int
}

type NN_RC_FaPai struct {
	Cards    []int
	LeftTime int
}

type NN_RC_NoticeQZhuang struct {
	LeftTime int
}

type NN_CR_QZhuang struct {
	Flag int
}

type NN_RC_QZhuang struct {
	UserId int
	Flag   int
}

type NN_RC_DingZhuang struct {
	UserId int
	Flag   int
}
