package fsm

// CLOSED 门打开
var CLOSED = 1

// OPENED 门关闭
var OPENED = 0

// Current 当前状态 默认打开
var Current = 1

// DoorFsm 用门举例来实现一个简单的状态机
type DoorFsm struct {
	CURRENT int
	OPENED  int
	CLOSED  int
}

// GetCurrentState 获取当前状态
func (d *DoorFsm) GetCurrentState(current int) {}

// OpeningEvent 开门的事件 将会把门关闭的状态转化为开启
func (d *DoorFsm) OpeningEvent() {
	d.CURRENT = OPENED
}

// ClosingEvent 关门的事件 将会把门开启的状态转化为关闭
func (d *DoorFsm) ClosingEvent() {
	d.CURRENT = CLOSED
}
