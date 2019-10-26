package voxcake

// Debug used to check application for errors
var Debug debugStruct

type debugStruct struct{}

func (debug *debugStruct) Check(err error) {
	if err != nil {
		panic(err)
	}
}

func (debug *debugStruct) Log(message string) {
	print(message)
}
