package voxcake

import (
	. "github.com/go-gl/mathgl/mgl32"
)

var vertices = make([]float32, 0)
var colors = make([]uint8, 0)
var elements = make([]uint32, 0)
var AmbientOcclusion = []float32{7, 6, 4}
var EdgeLight = []float32{12, 11, 10.7}

func (chunk *chunkStruct) generateMesh() {
	faceIndex := 0
	Scale := float32(1)

	for x := 0; x < ChunkSizeX; x++ {
		for y := 0; y < ChunkSizeY; y++ {
			for z := 0; z < ChunkSizeZ; z++ {
				voxel := chunk.get(x, y, z)
				if voxel != 0 {
					block01 := chunk.get(x-1, y-1, z+1)
					block02 := chunk.get(x, y-1, z+1)
					block03 := chunk.get(x+1, y-1, z+1)
					block04 := chunk.get(x-1, y-1, z)
					block05 := chunk.get(x, y-1, z)
					block06 := chunk.get(x+1, y-1, z)
					block07 := chunk.get(x-1, y-1, z-1)
					block08 := chunk.get(x, y-1, z-1)
					block09 := chunk.get(x+1, y-1, z-1)

					block11 := chunk.get(x-1, y, z+1)
					block12 := chunk.get(x, y, z+1)
					block13 := chunk.get(x+1, y, z+1)
					block14 := chunk.get(x-1, y, z)
					block16 := chunk.get(x+1, y, z)
					block17 := chunk.get(x-1, y, z-1)
					block18 := chunk.get(x, y, z-1)
					block19 := chunk.get(x+1, y, z-1)

					block21 := chunk.get(x-1, y+1, z+1)
					block22 := chunk.get(x, y+1, z+1)
					block23 := chunk.get(x+1, y+1, z+1)
					block24 := chunk.get(x-1, y+1, z)
					block25 := chunk.get(x, y+1, z)
					block26 := chunk.get(x+1, y+1, z)
					block27 := chunk.get(x-1, y+1, z-1)
					block28 := chunk.get(x, y+1, z-1)
					block29 := chunk.get(x+1, y+1, z-1)

					xScale := float32(x) / Scale
					yScale := float32(y)/Scale + 1
					zScale := float32(z) / Scale
					i := float32(1) / Scale

					if block16 == 0 {
						SetFace(Vec3{xScale + i, yScale - i, zScale},
							Vec3{xScale + i, yScale, zScale},
							Vec3{xScale + i, yScale, zScale + i},
							Vec3{xScale + i, yScale - i, zScale + i},
							faceIndex, voxel, 0.20, chunk,

							block09, block06, block03,
							block19, block13,
							block29, block26, block23,

							block08, block05, block02,
							block18, block12,
							block28, block25, block22)
						faceIndex++
					}

					if block14 == 0 {
						SetFace(Vec3{xScale, yScale - i, zScale + i},
							Vec3{xScale, yScale, zScale + i},
							Vec3{xScale, yScale, zScale},
							Vec3{xScale, yScale - i, zScale},
							faceIndex, voxel, 0.20, chunk,

							block01, block04, block07,
							block11, block17,
							block21, block24, block27,

							block02, block05, block08,
							block12, block18,
							block22, block25, block28)
						faceIndex++
					}

					if block25 == 0 {
						SetFace(Vec3{xScale, yScale, zScale + i},
							Vec3{xScale + i, yScale, zScale + i},
							Vec3{xScale + i, yScale, zScale},
							Vec3{xScale, yScale, zScale},
							faceIndex, voxel, 0, chunk,

							block21, block24, block27,
							block22, block28,
							block23, block26, block29,

							block11, block14, block17,
							block12, block18,
							block13, block16, block19)
						faceIndex++
					}

					if block05 == 0 {
						SetFace(Vec3{xScale, yScale - i, zScale},
							Vec3{xScale + i, yScale - i, zScale},
							Vec3{xScale + i, yScale - i, zScale + i},
							Vec3{xScale, yScale - i, zScale + i},
							faceIndex, voxel, 0.5, chunk,

							block07, block04, block01,
							block08, block02,
							block09, block06, block03,

							block17, block14, block11,
							block18, block12,
							block19, block16, block13)
						faceIndex++
					}

					if block12 == 0 {
						SetFace(Vec3{xScale + i, yScale - i, zScale + i},
							Vec3{xScale + i, yScale, zScale + i},
							Vec3{xScale, yScale, zScale + i},
							Vec3{xScale, yScale - i, zScale + i},
							faceIndex, voxel, 0.25, chunk,

							block03, block02, block01,
							block13, block11,
							block23, block22, block21,

							block06, block05, block04,
							block16, block14,
							block26, block25, block24)
						faceIndex++
					}

					if block18 == 0 {
						SetFace(Vec3{xScale, yScale - i, zScale},
							Vec3{xScale, yScale, zScale},
							Vec3{xScale + i, yScale, zScale},
							Vec3{xScale + i, yScale - i, zScale},
							faceIndex, voxel, 0.25, chunk,

							block07, block08, block09,
							block17, block19,
							block27, block28, block29,

							block04, block05, block06,
							block14, block16,
							block24, block25, block26)
						faceIndex++
					}
				}
			}
		}
	}

	if len(vertices) != 0 {
		chunk.mesh.Vertices = vertices
		chunk.mesh.Colors = colors
		chunk.mesh.Indices = elements
		chunk.mesh.Start()
	}
	vertices = nil
	colors = nil
	elements = nil
}

func SetFace(v1 Vec3, v2 Vec3, v3 Vec3, v4 Vec3, index int, color uint32, light float32, chunk *chunkStruct,
	ao1 uint32, ao2 uint32, ao3 uint32, ao4 uint32, ao6 uint32, ao7 uint32, ao8 uint32, ao9 uint32,
	le1 uint32, le2 uint32, le3 uint32, le4 uint32, le6 uint32, le7 uint32, le8 uint32, le9 uint32) {
	flip := false
	index = index * 4

	r := float32(UintToR(color))
	g := float32(UintToG(color))
	b := float32(UintToB(color))

	color = RGBAToUint(uint32(r-r*light),
		uint32(g-g*light),
		uint32(b-b*light),
		255)

	color1 := GetLight(color, le2, le1, le4, ao2, ao1, ao4, 1, &flip)
	color2 := GetLight(color, le4, le7, le8, ao4, ao7, ao8, 2, &flip)
	color3 := GetLight(color, le8, le9, le6, ao8, ao9, ao6, 3, &flip)
	color4 := GetLight(color, le6, le3, le2, ao6, ao3, ao2, 4, &flip)

	vertices = append(vertices, v1.X())
	vertices = append(vertices, v1.Y())
	vertices = append(vertices, v1.Z())

	vertices = append(vertices, v2.X())
	vertices = append(vertices, v2.Y())
	vertices = append(vertices, v2.Z())

	vertices = append(vertices, v3.X())
	vertices = append(vertices, v3.Y())
	vertices = append(vertices, v3.Z())

	vertices = append(vertices, v4.X())
	vertices = append(vertices, v4.Y())
	vertices = append(vertices, v4.Z())

	color1 = GetAmbient(color1, ao2, ao1, ao4, 1, &flip)
	colors = append(colors, uint8(UintToR(color1)))
	colors = append(colors, uint8(UintToG(color1)))
	colors = append(colors, uint8(UintToB(color1)))

	color2 = GetAmbient(color2, ao4, ao7, ao8, 2, &flip)
	colors = append(colors, uint8(UintToR(color2)))
	colors = append(colors, uint8(UintToG(color2)))
	colors = append(colors, uint8(UintToB(color2)))

	color3 = GetAmbient(color3, ao8, ao9, ao6, 3, &flip)
	colors = append(colors, uint8(UintToR(color3)))
	colors = append(colors, uint8(UintToG(color3)))
	colors = append(colors, uint8(UintToB(color3)))

	color4 = GetAmbient(color4, ao6, ao3, ao2, 4, &flip)
	colors = append(colors, uint8(UintToR(color4)))
	colors = append(colors, uint8(UintToG(color4)))
	colors = append(colors, uint8(UintToB(color4)))

	face := uint32(index / 4)
	if !flip {
		elements = append(elements, face*4)
		elements = append(elements, face*4+1)
		elements = append(elements, face*4+2)
		elements = append(elements, face*4)
		elements = append(elements, face*4+2)
		elements = append(elements, face*4+3)
	} else {
		elements = append(elements, face*4+3)
		elements = append(elements, face*4)
		elements = append(elements, face*4+1)
		elements = append(elements, face*4+3)
		elements = append(elements, face*4+1)
		elements = append(elements, face*4+2)
	}
	flip = false
}

func GetLight(color uint32, b1 uint32, b2 uint32, b3 uint32, b11 uint32, b12 uint32, b13 uint32, vertex uint8, flip *bool) uint32 {
	/*
		if b1 == 0 && b3 == 0 && b11 == 0 && b13 == 0 {
			if vertex == 1 || vertex == 3 {
				*flip = true
			}
			return CalculateLight(UintToR(color), UintToG(color), UintToB(color), EdgeLight[2])
		} else if (b1 == 0 && b11 == 0 && b2 == 0 && b12 == 0) || (b2 == 0 && b12 == 0 && b3 == 0 && b13 == 0) {
			return CalculateLight(UintToR(color), UintToG(color), UintToB(color), EdgeLight[1])
		} else if (b1 == 0 && b11 == 0) || (b3 == 0 && b13 == 0) {
			return CalculateLight(UintToR(color), UintToG(color), UintToB(color), EdgeLight[0])
		} else if b2 == 0 && b12 == 0 {
			if vertex == 1 || vertex == 3 {
				*flip = true
			}
			return CalculateLight(UintToR(color), UintToG(color), UintToB(color), EdgeLight[0])
		}*/
	return color
}

func GetAmbient(color uint32, b1 uint32, b2 uint32, b3 uint32, vertex uint8, flip *bool) uint32 {
	/*
		if b1 != 0 && b3 != 0 {
			if vertex == 1 || vertex == 3 {
				*flip = true
			}
			return CalculateAmbient(UintToR(color), UintToG(color), UintToB(color), AmbientOcclusion[2])
		} else if (b1 != 0 && b2 != 0) || (b2 != 0 && b3 != 0) {
			return CalculateAmbient(UintToR(color), UintToG(color), UintToB(color), AmbientOcclusion[1])
		} else if b1 != 0 || b3 != 0 {
			return CalculateAmbient(UintToR(color), UintToG(color), UintToB(color), AmbientOcclusion[0])
		} else if b2 != 0 {
			if vertex == 1 || vertex == 3 {
				*flip = true
			}
			return CalculateAmbient(UintToR(color), UintToG(color), UintToB(color), AmbientOcclusion[0])
		}*/
	return color
}
