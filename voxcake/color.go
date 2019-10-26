package voxcake

import "fmt"

func RGBAToUint(r uint32, g uint32, b uint32, a uint32) uint32 {
	return uint32((a << 24) | (r << 16) | (g << 8) | (b << 0))
}

func UintToRGBA(value uint32) string {
	a := uint8(value >> 24)
	r := uint8(value >> 16)
	g := uint8(value >> 8)
	b := uint8(value >> 0)

	return fmt.Sprintf("%[1]d, %[2]d, %[3]d, %[4]d", r, g, b, a)
}

func UintToR(value uint32) uint32 {
	return uint32(uint8(value >> 0))
}

func UintToG(value uint32) uint32 {
	return uint32(uint8(value >> 8))
}

func UintToB(value uint32) uint32 {
	return uint32(uint8(value >> 16))
}

func UintToA(value uint32) uint32 {
	return uint32(uint8(value >> 24))
}

func CalculateAmbient(r uint32, g uint32, b uint32, ao float32) uint32 {
	fr := float32(r) - float32(r)/ao
	fg := float32(g) - float32(g)/ao
	fb := float32(b) - float32(b)/ao

	if fr > 255 {
		fr = 255
	}
	if fr < 0 {
		fr = 0
	}
	if fg > 255 {
		fg = 255
	}
	if fg < 0 {
		fg = 0
	}
	if fb > 255 {
		fb = 255
	}
	if fb < 0 {
		fb = 0
	}

	return RGBAToUint(uint32(fr), uint32(fg), uint32(fb), 255)
}

func CalculateLight(r uint32, g uint32, b uint32, ao float32) uint32 {
	fr := float32(r) + float32(r)/ao
	fg := float32(g) + float32(g)/ao
	fb := float32(b) + float32(b)/ao

	if fr > 255 {
		fr = 255
	}
	if fr < 0 {
		fr = 0
	}
	if fg > 255 {
		fg = 255
	}
	if fg < 0 {
		fg = 0
	}
	if fb > 255 {
		fb = 255
	}
	if fb < 0 {
		fb = 0
	}

	return RGBAToUint(uint32(fr), uint32(fg), uint32(fb), 255)
}
