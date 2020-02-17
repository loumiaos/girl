package niuniu

func (self *Room) update(dt int64) {
	if self.curState == self.nextState {
		self.curFsmFunc[1](dt)
	} else {
		self.curFsmFunc[2](dt)
		self.changeState()
		self.curFsmFunc[0](dt)
		self.curFsmFunc[1](dt)
	}
}

func (self *Room) onEnterIdle(dt int64) {

}

func (self *Room) onExecIdle(dt int64) {

}

func (self *Room) onExitIdle(dt int64) {

}

func (self *Room) onEnterFaPai(dt int64) {

}

func (self *Room) onExecFaPai(dt int64) {

}

func (self *Room) onExitFaPai(dt int64) {

}

func (self *Room) onEnterBiPai(dt int64) {

}

func (self *Room) onExecBiPai(dt int64) {

}

func (self *Room) onExitBiPai(dt int64) {

}

func (self *Room) onEnterResult(dt int64) {

}

func (self *Room) onExecResult(dt int64) {

}

func (self *Room) onExitResult(dt int64) {

}
