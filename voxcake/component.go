package voxcake

type IComponent interface {
	Name() string
	Start()
	Update()
	Return() interface{}
}
