package voxcake

var Input inputStruct
var KeyCode = initializeKeyCode()

type inputStruct struct {
	mouse    mouseStruct
	keyboard keyboardStruct
}

type mouseStruct struct {
	ButtonPrev [3]uint8
	ButtonNow  [3]uint8
	X          int32
	Y          int32
}

type keyboardStruct struct {
	Key   uint8
	State uint8
}

func (mouse *mouseStruct) SetMouseButton(button uint8, state uint8) {
	mouse.ButtonNow[button-1] = state
	if mouse.ButtonNow[button-1] == 0 {
		mouse.ButtonPrev[button-1] = 0
	}
}

func (mouse *mouseStruct) SetMousePosition(x int32, y int32) {
	mouse.X = x
	mouse.Y = y
}

func (keyboard *keyboardStruct) SetKeyboardInput(state uint8, key uint8) {
	keyboard.State = state
	keyboard.Key = key
}

func (inputHandle *inputStruct) GetKeyDown(keyCode uint8) bool {
	if inputHandle.keyboard.Key == keyCode && inputHandle.keyboard.State != 0 {
		return true
	}
	return false
}

func (inputHandle *inputStruct) GetButtonDown(button uint8) bool {
	if inputHandle.mouse.ButtonNow[button] == 1 && inputHandle.mouse.ButtonPrev[button] != 1 {
		inputHandle.mouse.ButtonPrev[button] = 1
		return true
	}
	return false
}

func (inputHandle *inputStruct) GetAxis(axis string) int32 {
	if axis == "X" {
		return inputHandle.mouse.X
	} else if axis == "Y" {
		return inputHandle.mouse.Y
	}
	return 0
}
