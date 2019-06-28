package voxcake

import (
	. "github.com/go-gl/mathgl/mgl32"
)

const (
	ChunkSizeX = 16
	ChunkSizeY = 16
	ChunkSizeZ = 16
)

type chunkStruct struct {
	volume   *volumeStruct
	position Vec3
	mesh     MeshStruct
}

//Render method render this Chunk
func (chunk *chunkStruct) render() {
	chunk.mesh.Draw(Translate3D(chunk.position.X(), chunk.position.Y(), chunk.position.Z()))
}

func (chunk *chunkStruct) get(x, y, z int) uint32 {
	return chunk.volume.GetVoxel(x+int(chunk.position.X()), y+int(chunk.position.Y()), z+int(chunk.position.Z()))
}

func (chunk *chunkStruct) update() {

}
