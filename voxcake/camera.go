package voxcake

import (
	"math"

	. "github.com/go-gl/mathgl/mgl32"
)

var Camera CameraStruct

type CameraStruct struct {
	Fov       float32
	Position  Vec3
	Rotation  Vec3
	Direction Vec3

	right Vec3
	Up    Vec3
	Front Vec3

	DrawDistance float32

	view       Mat4
	projection Mat4

	lastX float32
	lastY float32

	pitch float32
	yaw   float32
}

func NewCamera() *CameraStruct {
	var camera CameraStruct

	camera.Position = Vec3{0, 0, -5}

	camera.Front = Vec3{0, 0, 1}
	camera.Up = Vec3{0, 1, 0}

	camera.Fov = 60.0
	camera.DrawDistance = 1000.0

	camera.lastX = float32(Window.Width / 2)
	camera.lastY = float32(Window.Height / 2)

	return &camera
}

func (camera *CameraStruct) Update() {
	camera.view = LookAtV(camera.Position, Vec3.Add(camera.Position, camera.Front), camera.Up)
	camera.projection = Perspective(DegToRad(camera.Fov), float32(Window.Width/Window.Height), 0.1, camera.DrawDistance)
}

func (camera *CameraStruct) Translate(direction Vec3) {
	dir := Vec3.Mul(direction.Mul(0.05), Time.DeltaTime)
	camera.Position = Vec3.Add(camera.Position, dir)
}

func (camera *CameraStruct) Rotate(x int32, y int32) {
	xOffset := float32(x) - camera.lastX
	yOffset := camera.lastY - float32(y)

	xOffset *= 0.1
	yOffset *= 0.1

	camera.yaw += xOffset
	camera.pitch += yOffset

	if camera.pitch < -89.0 {
		camera.pitch = -89.0
	}
	if camera.pitch > 89.0 {
		camera.pitch = 89.0
	}

	cPitch := math.Cos(float64(DegToRad(camera.pitch)))
	sPitch := math.Sin(float64(DegToRad(camera.pitch)))
	cYaw := math.Cos(float64(DegToRad(camera.yaw)))
	sYaw := math.Sin(float64(DegToRad(camera.yaw)))

	camera.Front = Vec3{
		float32(cPitch * cYaw),
		float32(sPitch),
		float32(cPitch * sYaw)}
	Camera.Front.Normalize()

	camera.lastX = float32(x)
	camera.lastY = float32(y)
}
