# VoxCakeGo
VoxCakeGo is just the toy where you can render voxel models... Just at this moment.
## Features
- [X] Volume component;
- [X] Camera component;
- [X] Mesh component:
- [ ] Math wrapper;

## Components
You can create your own component which will be executed by engine.
```go
package main

import (
	"fmt"

	. "./voxcake"
	. "github.com/go-gl/mathgl/mgl32"
)

type CameraController struct {
}

func (cameraController *CameraController) Start() {

}

func (cameraController *CameraController) Update() {
	nFront := Vec3{0, 0, 0}.Sub(Camera.Front)

	if Input.GetKeyDown(KeyCode.W) {
		Camera.Translate(Camera.Front)
	}
	if Input.GetKeyDown(KeyCode.S) {
		Camera.Translate(nFront)
	}
	if Input.GetKeyDown(KeyCode.A) {
		Camera.Translate(Vec3.Normalize(Vec3.Cross(nFront, Camera.Up)))
	}
	if Input.GetKeyDown(KeyCode.D) {
		Camera.Translate(Vec3.Normalize(Vec3.Cross(Camera.Front, Camera.Up)))
	}

	Camera.Rotate(Input.GetAxis("X"), Input.GetAxis("Y"))

	if Input.GetButtonDown(0) {
		//vec := RaycastVolume(&Camera.Position, &Camera.Direction, Volume("Map.vcmap"))
		//vec := RaycastVolume(Camera.Position.X(), Camera.Position.Y(), Camera.Position.Z(), Camera.Front.X(), Camera.Front.Y(), Camera.Front.Z(), uint8(1), Volume("Map.vcmap"))
		//fmt.Printf("Hitted: Vec3(%[1]d, %[2]d, %[3]d)\n", int(vec.X()), int(vec.Y()), int(vec.Z()))
	}
	if Input.GetButtonDown(1) {
		fmt.Printf("%s", "WHEEL\n")
	}
	if Input.GetButtonDown(2) {
		fmt.Printf("%s", "RIGHT\n")
	}

	if Input.GetKeyDown(KeyCode.Escape) {
		Engine.Quit()
	}

	if Input.GetKeyDown(KeyCode.F) {
		Window.SetFullscreen(true)
	}
}

func (cameraController *CameraController) Name() string {
	return "CameraController"
}

func (cameraController *CameraController) Return() interface{} {
	return cameraController
}
```
