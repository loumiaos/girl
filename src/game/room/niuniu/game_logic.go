package niuniu

import (
	"github.com/snowyyj001/loumiao/util"
)

var CARD_POKERS = []int{
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d,
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d,
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d,
	0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d,
	0x4e, 0x4f,
}

const (
	VALUE_MASK    = 0x0f
	COLOR_MASK    = 0xf0
	JOKER_X_INDEX = 52
	JOKER_D_INDEX = 53
	JOKER_X       = 0x4e
	JOKER_D       = 0x4f
)

//--牌型
const (
	TYPE_NONE   = 0  //没牛
	TYPE_NIUNIU = 10 //牛牛
	TYPE_WHN    = 11 //五花牛
	TYPE_HLN    = 12 //葫芦牛
	TYPE_ZDN    = 13 //炸弹牛
	TYPE_WXN    = 14 //五小牛
)

type GameLogic struct {
}

var g_GameLogic *GameLogic

//混乱扑克
func (self *GameLogic) shuffle(cards []int) {
	var rindex1, tmpdata int
	length := len(cards)
	for i := 0; i < length; i++ {
		rindex1 = util.Random(length)
		tmpdata = cards[rindex1]
		cards[rindex1] = cards[i]
		cards[i] = tmpdata
	}
}

func (self *GameLogic) value2index(cards []int) []int {
	indexs := []int{0, 0, 0, 0, 0}
	for k, v := range cards {
		if v == JOKER_X {
			indexs[k] = JOKER_X_INDEX
		} else if v == JOKER_D {
			indexs[k] = JOKER_D_INDEX
		} else {
			indexs[k] = (self.getValue(v)-1)*4 + self.getColor(v)
		}
	}
	return indexs
}

func (self *GameLogic) index2value(cards []int) []int {
	indexs := []int{0, 0, 0, 0, 0}
	for k, v := range cards {
		if v == JOKER_X_INDEX {
			indexs[k] = JOKER_X
		} else if v == JOKER_D_INDEX {
			indexs[k] = JOKER_D
		} else {
			indexs[k] = ((v % 4) << 4) + util.FloorInt(v/4) + 1
		}
	}
	return indexs
}

func (self *GameLogic) getLogicValue(_poker int) int {
	val := _poker & VALUE_MASK
	if val > 10 {
		return 10
	}
	return val
}

func (self *GameLogic) getValue(_poker int) int {
	return _poker & VALUE_MASK
}

func (self *GameLogic) getColor(_poker int) int {
	return (_poker & COLOR_MASK) >> 4
}

func (self *GameLogic) sortByValue(cards []int, l2b bool) {
	sortflag := true
	length := len(cards) - 1
	for sortflag {
		sortflag = false
		for i := 0; i < length; i++ {
			if l2b {
				if self.getValue(cards[i]) > self.getValue(cards[i+1]) {
					cards[i], cards[i+1] = cards[i+1], cards[i]
					sortflag = true
				}
			} else {
				if self.getValue(cards[i]) < self.getValue(cards[i+1]) {
					cards[i], cards[i+1] = cards[i+1], cards[i]
					sortflag = true
				}
			}
		}
		length = length - 1
	}
}

func (self *GameLogic) getSpecialType(cards []int) int {
	vlsmall := 0
	for i := 0; i < HANDCARD_NUM; i++ {
		if cards[i] > 5 || vlsmall > 10 {
			vlsmall = 11
			break
		}
		vlsmall = vlsmall + cards[i]
	}
	if vlsmall <= 10 {
		return TYPE_WXN
	}
	if cards[0] == cards[3] || cards[1] == cards[4] {
		return TYPE_ZDN
	}
	if cards[0] == cards[2] && cards[3] == cards[4] {
		return TYPE_HLN
	}
	if cards[0] == cards[1] && cards[2] == cards[4] {
		return TYPE_HLN
	}
	wuhua := true
	for i := 0; i < HANDCARD_NUM; i++ {
		if cards[i] < 11 {
			wuhua = false
			break
		}
	}
	if wuhua {
		return TYPE_WHN
	}
	return TYPE_NONE
}

var allComposes = [][]int{
	{0, 1, 2, 3, 4},
	{0, 1, 3, 2, 4},
	{0, 1, 4, 2, 3},
	{0, 2, 3, 1, 4},
	{0, 2, 4, 1, 3},
	{0, 3, 4, 1, 2},
	{1, 2, 3, 0, 4},
	{1, 2, 4, 0, 3},
	{1, 3, 4, 0, 2},
	{2, 3, 4, 0, 1},
}

func (self *GameLogic) getBaseType(cards []int) (int, []int) {
	for i := 0; i < 10; i++ {
		if (cards[allComposes[i][1]]+cards[allComposes[i][2]]+cards[allComposes[i][3]])%10 == 0 {
			val := (cards[allComposes[i][4]] + cards[allComposes[i][5]]) % 10
			if val == 0 {
				return TYPE_NIUNIU, allComposes[i]
			} else {
				return val, allComposes[i]
			}
		}
	}
	return TYPE_NONE, allComposes[9]
}

func init() {
	g_GameLogic = &GameLogic{}
}
