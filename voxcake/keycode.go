package voxcake

type keyCodeStruct struct {
	A uint8
	B uint8
	C uint8
	D uint8
	E uint8
	F uint8
	G uint8
	H uint8
	I uint8
	J uint8
	K uint8
	L uint8
	M uint8
	N uint8
	O uint8
	P uint8
	Q uint8
	R uint8
	S uint8
	T uint8
	U uint8
	V uint8
	W uint8
	X uint8
	Y uint8
	Z uint8

	Alpha0 uint8
	Alpha1 uint8
	Alpha2 uint8
	Alpha3 uint8
	Alpha4 uint8
	Alpha5 uint8
	Alpha6 uint8
	Alpha7 uint8
	Alpha8 uint8
	Alpha9 uint8
	Plus   uint8
	Minus  uint8

	Num0 uint8
	Num1 uint8
	Num2 uint8
	Num3 uint8
	Num4 uint8
	Num5 uint8
	Num6 uint8
	Num7 uint8
	Num8 uint8
	Num9 uint8

	F1  uint8
	F2  uint8
	F3  uint8
	F4  uint8
	F5  uint8
	F6  uint8
	F7  uint8
	F8  uint8
	F9  uint8
	F10 uint8
	F11 uint8
	F12 uint8

	Escape     uint8
	Enter      uint8
	Space      uint8
	AltLeft    uint8
	CtrlLeft   uint8
	ShiftLeft  uint8
	AltRight   uint8
	CtrlRight  uint8
	ShiftRight uint8
	Tab        uint8

	PrintScreen uint8
	Del         uint8
	End         uint8
	Home        uint8
	Hz          uint8

	ArrowLeft  uint8
	ArrowRight uint8
	ArrowUp    uint8
	ArrowDown  uint8
}

func initializeKeyCode() *keyCodeStruct {
	var keyCode keyCodeStruct
	keyCode.A = 97
	keyCode.B = 98
	keyCode.C = 99
	keyCode.D = 100
	keyCode.E = 101
	keyCode.F = 102
	keyCode.G = 103
	keyCode.H = 104
	keyCode.I = 105
	keyCode.J = 106
	keyCode.K = 107
	keyCode.L = 108
	keyCode.M = 109
	keyCode.N = 110
	keyCode.O = 111
	keyCode.P = 112
	keyCode.Q = 113
	keyCode.R = 114
	keyCode.S = 115
	keyCode.T = 116
	keyCode.U = 117
	keyCode.V = 118
	keyCode.W = 119
	keyCode.X = 120
	keyCode.Y = 121
	keyCode.Z = 122

	keyCode.Alpha0 = 48
	keyCode.Alpha1 = 49
	keyCode.Alpha2 = 50
	keyCode.Alpha3 = 51
	keyCode.Alpha4 = 52
	keyCode.Alpha5 = 53
	keyCode.Alpha6 = 54
	keyCode.Alpha7 = 55
	keyCode.Alpha8 = 56
	keyCode.Alpha9 = 57
	keyCode.Minus = 45
	keyCode.Plus = 61

	//keyCode.Num0 = 89
	keyCode.Num1 = 89
	keyCode.Num2 = 90
	keyCode.Num3 = 91
	keyCode.Num4 = 92
	keyCode.Num5 = 93
	keyCode.Num6 = 94
	keyCode.Num7 = 95
	keyCode.Num8 = 96
	keyCode.Num9 = 97

	keyCode.F1 = 58
	keyCode.F2 = 59
	keyCode.F3 = 60
	keyCode.F4 = 61
	keyCode.F5 = 62
	keyCode.F6 = 63
	keyCode.F7 = 64
	keyCode.F8 = 65
	keyCode.F9 = 66
	keyCode.F10 = 67
	keyCode.F11 = 68
	keyCode.F12 = 69

	keyCode.Escape = 27
	keyCode.Enter = 13
	keyCode.Space = 32
	keyCode.AltLeft = 226
	keyCode.CtrlLeft = 224
	keyCode.ShiftLeft = 225
	keyCode.AltRight = 230
	keyCode.CtrlRight = 228
	keyCode.ShiftRight = 229
	keyCode.Tab = 99

	keyCode.PrintScreen = 70
	keyCode.Del = 127
	keyCode.End = 74
	//keyCode.Home =
	keyCode.Hz = 96 //~

	return &keyCode
}
