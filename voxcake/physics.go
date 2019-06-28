package voxcake

import (
	"fmt"
	"math"

	. "github.com/go-gl/mathgl/mgl32"
)

type Ray struct {
	Origin    Vec3
	Direction Vec3
}

type AABB struct {
	Bounds []Vec3
	Tag    string
}

//func RaycastVolume(origin *Vec3, direction *Vec3, volume *VolumeStruct) Vec3 {
func RaycastVolume(gx0, gy0, gz0, rayX, rayY, rayZ float32, mode uint8, volume *volumeStruct) Vec3 {
	distance := float32(32)
	gx1 := gx0 + rayX*distance
	gy1 := gy0 + rayY*distance
	gz1 := gz0 + rayZ*distance

	gx0idx := FloorI(gx0)
	gy0idx := FloorI(gy0)
	gz0idx := FloorI(gz0)
	gx1idx := FloorI(gx1)
	gy1idx := FloorI(gy1)
	gz1idx := FloorI(gz1)

	sx := 0
	sy := 0
	sz := 0
	if gx1idx > gx0idx {
		sx = 1
	} else {
		if gx1idx < gx0idx {
			sx = -1
		}
	}

	if gy1idx > gy0idx {
		sy = 1
	} else {
		if gy1idx < gy0idx {
			sy = -1
		}
	}

	if gz1idx > gz0idx {
		sz = 1
	} else {
		if gz1idx < gz0idx {
			sz = -1
		}
	}

	gx := gx0idx
	gy := gy0idx
	gz := gz0idx

	gxp := 0
	gyp := 0
	gzp := 0
	if gx1idx > gx0idx {
		gxp = gx0idx + 1
	} else {
		gxp = gx0idx + 0
	}
	if gy1idx > gy0idx {
		gyp = gy0idx + 1
	} else {
		gyp = gy0idx + 0
	}
	if gz1idx > gz0idx {
		gzp = gz0idx + 1
	} else {
		gzp = gz0idx + 0
	}

	vx := float32(0)
	vy := float32(0)
	vz := float32(0)
	if gx1 == gx0 {
		vx = 1
	} else {
		vx = gx1 - gx0
	}
	if gy1 == gy0 {
		vy = 1
	} else {
		vy = gy1 - gy0
	}
	if gz1 == gz0 {
		vz = 1
	} else {
		vz = gz1 - gz0
	}

	vxvy := vx * vy
	vxvz := vx * vz
	vyvz := vy * vz

	errx := float32(gxp) - gx0*vyvz
	erry := float32(gyp) - gy0*vxvz
	errz := float32(gzp) - gz0*vxvy

	derrx := float32(sx) * vyvz
	derry := float32(sy) * vxvz
	derrz := float32(sz) * vxvy

	gxPre := gx
	gyPre := gy
	gzPre := gz

	ret := []int{0, 0, 0, 0, 0, 0}
	for {
		if gx < 0 || gx >= volume.width || gy < 0 || gy >= volume.height || gz < 0 || gz >= volume.depth {
			return Vec3{228, 228, 228}
		}
		switch mode {
		case 0:
			if volume.GetVoxel(gx, gy, gz) != 0 && volume.GetVoxel(gxPre, gyPre, gzPre) == 0 {
				ret[0] = gxPre
				ret[1] = gyPre
				ret[2] = gzPre
				return Vec3{float32(ret[0]), float32(ret[1]), float32(ret[2])}
			}
			break
		case 1:
			if volume.GetVoxel(gx, gy, gz) != 0 {
				ret[0] = gx
				ret[1] = gy
				ret[2] = gz
				ret[3] = gxPre
				ret[4] = gyPre
				ret[5] = gzPre
				return Vec3{float32(ret[0]), float32(ret[1]), float32(ret[2])}
			}
			break
		}
		gxPre = gx
		gyPre = gy
		gzPre = gz

		if gx == gx1idx && gy == gy1idx && gz == gz1idx {
			break
		}

		xr := AbsI(errx)
		yr := AbsI(erry)
		zr := AbsI(errz)

		if sx != 0 && (sy == 0 || xr < yr) && (sz == 0 || xr < zr) {
			gx += sx
			errx += derrx
		} else if sy != 0 && (sz == 0 || yr < zr) {
			gy += sy
			erry += derry
		} else if sz != 0 {
			gz += sz
			errz += derrz
		}
	}
	return Vec3{1337, 1337, 1337}
}

func (ray *Ray) Hit(aabb *AABB, hitPoint Vec3) bool {
	dx := 1.0 / ray.Direction.X()
	dy := 1.0 / ray.Direction.Y()
	dz := 1.0 / ray.Direction.Z()

	t1 := float64((aabb.Bounds[0].X() - ray.Origin.X()) * dx)
	t2 := float64((aabb.Bounds[1].X() - ray.Origin.X()) * dx)
	t3 := float64((aabb.Bounds[0].Y() - ray.Origin.Y()) * dy)
	t4 := float64((aabb.Bounds[1].Y() - ray.Origin.Y()) * dy)
	t5 := float64((aabb.Bounds[0].Z() - ray.Origin.Z()) * dz)
	t6 := float64((aabb.Bounds[1].Z() - ray.Origin.Z()) * dz)

	tmin := float32(math.Max(math.Max(math.Min(t1, t2), math.Min(t3, t4)), math.Min(t5, t6)))
	tmax := float32(math.Min(math.Min(math.Max(t1, t2), math.Max(t3, t4)), math.Max(t5, t6)))

	t := float32(0)
	if tmax < 0 {
		t = tmax
		return false
	}
	if tmin > tmax {
		t = tmax
		return false
	}
	t = tmin
	fmt.Printf("%d", t)

	return true
}

func FloorI(f32 float32) int {
	return int(math.Floor(float64(f32)))
}

func AbsI(f32 float32) int {
	return int(math.Abs(float64(f32)))
}

func SwapF(a *float32, b *float32) {
	tmp := b
	b = a
	a = tmp
}

func SwapI(a *int, b *int) {
	tmp := b
	b = a
	a = tmp
}
