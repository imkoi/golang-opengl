package voxcake

var Time timeStruct

type timeStruct struct {
	lastTick  uint32
	DeltaTime float32
}

func (time *timeStruct) setDeltaTime(ticks uint32) {
	time.DeltaTime = float32(ticks - time.lastTick)
	time.lastTick = ticks
}
