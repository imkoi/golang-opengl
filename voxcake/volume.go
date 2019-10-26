package voxcake

import (
	"io/ioutil"

	. "github.com/go-gl/mathgl/mgl32"
)

type volumeStruct struct {
	width  int
	height int
	depth  int
	wdc    int
	hdc    int
	ddc    int
	data   [][][]uint32
	chunk  [][][]chunkStruct
}

//Volume is constructor method
func Volume(file string) *volumeStruct {
	var volume volumeStruct
	volume.width = 512
	volume.height = 200
	volume.depth = 512
	volume.wdc = volume.width / ChunkSizeX
	volume.hdc = volume.height / ChunkSizeY
	volume.ddc = volume.depth / ChunkSizeZ

	volume.data = make([][][]uint32, volume.width)
	for i := range volume.data {
		volume.data[i] = make([][]uint32, volume.height)
		for j := range volume.data[i] {
			volume.data[i][j] = make([]uint32, volume.depth)
		}
	}

	volume.chunk = make([][][]chunkStruct, volume.wdc)
	for i := range volume.chunk {
		volume.chunk[i] = make([][]chunkStruct, volume.hdc)
		for j := range volume.chunk[i] {
			volume.chunk[i][j] = make([]chunkStruct, volume.ddc)
		}
	}

	return &volume
}

func (volume *volumeStruct) Start() {
	volume.Load("cityofchicago.vxl")
	volume.Rotate(0, 0, 0)
	for x := 0; x < volume.wdc; x++ {
		for y := 0; y < volume.hdc; y++ {
			for z := 0; z < volume.ddc; z++ {
				volume.chunk[x][y][z].position = Vec3{float32(x * ChunkSizeX), float32(y * ChunkSizeY), float32(z * ChunkSizeZ)}
				volume.chunk[x][y][z].volume = volume
				volume.chunk[x][y][z].generateMesh()
			}
		}
	}
}

func (volume *volumeStruct) Update() {
	cxMin, cxMax, cyMin, cyMax, czMin, czMax := volume.getViewChunks()

	for x := cxMin; x < cxMax; x++ {
		for y := cyMin; y < cyMax; y++ {
			for z := czMin; z < czMax; z++ {
				volume.chunk[x][y][z].render()
			}
		}
	}
}

func (volume *volumeStruct) Load(file string) {
	bytes, err := ioutil.ReadFile("./resources/maps/" + file)
	Debug.Check(err)

	pos := 0
	for x := 0; x < volume.width; x++ {
		for z := 0; z < volume.depth; z++ {
			y := 0
			for ; y < volume.height; y++ {
				volume.data[x][y][z] = RGBAToUint(192, 32, 32, 255)
			}
			y = 0

			for {
				number4ByteChunks := int(bytes[pos+0])
				topColorStart := int(bytes[pos+1])
				topColorEnd := int(bytes[pos+2])
				colorPos := int(pos + 4)

				for ; y < topColorStart; y++ {
					volume.data[x][y][z] = 0
				}

				for ; y <= topColorEnd; y++ {
					volume.data[x][y][z] = RGBAToUint(uint32(bytes[colorPos]), uint32(bytes[colorPos+1]), uint32(bytes[colorPos+2]), uint32(bytes[colorPos+3]))
					colorPos += 4
				}
				if topColorEnd == volume.height-2 {
					volume.data[x][y][volume.height-1] = volume.data[x][y][volume.height-2]
				}

				lenBottom := topColorEnd - topColorStart + 1
				if number4ByteChunks == 0 {
					pos += 4 * (lenBottom + 1)
					break
				}

				lenTop := number4ByteChunks - lenBottom - 1
				pos += int(bytes[pos]) * 4
				bottomColorEnd := int(bytes[pos+3])
				bottomColorStart := bottomColorEnd - lenTop

				for y = bottomColorStart; y < bottomColorEnd; y++ {
					volume.data[x][y][z] = RGBAToUint(uint32(bytes[colorPos]), uint32(bytes[colorPos+1]), uint32(bytes[colorPos+2]), uint32(bytes[colorPos+3]))
					colorPos += 4
				}
				if bottomColorEnd == volume.height-1 {
					volume.data[x][y][volume.height-1] = volume.data[x][y][volume.height-2]
				}
			}
		}
	}
}

func (volume *volumeStruct) Rotate(x, y, z int) {
	for i := 0; i < volume.width; i++ {
		for j := 0; j < volume.height; j++ {
			for k := 0; k < volume.depth; k++ {
				voxel := volume.data[i][j][k]
				if voxel != 0 {
					volume.data[i][volume.height-j-1][k] = voxel
					volume.data[i][j][k] = 0
				}
			}
		}
	}
}

func (volume *volumeStruct) GetVoxel(x, y, z int) uint32 {
	if x < 0 || x >= volume.width || y < 0 || y >= volume.height || z < 0 || z >= volume.depth {
		return 0
	}
	return volume.data[x][y][z]
}

func (volume *volumeStruct) SetVoxel(x, y, z int, value uint32) {
	if x < 0 || x >= volume.width || y < 0 || y >= volume.height || z < 0 || z >= volume.depth {
		return
	}
	volume.data[x][y][z] = value
}

func (volume *volumeStruct) getViewChunks() (int, int, int, int, int, int) {
	camX := int(Camera.Position.X() / ChunkSizeX)
	camY := int(Camera.Position.Y() / ChunkSizeY)
	camZ := int(Camera.Position.Z() / ChunkSizeZ)

	maxCx := volume.wdc
	maxCy := volume.hdc
	maxCz := volume.ddc

	cxMin := camX - 32
	cxMax := camX + 32
	cyMin := camY - 32
	cyMax := camY + 32
	czMin := camZ - 32
	czMax := camZ + 32

	if cxMin < 0 {
		cxMin = 0
	}
	if cxMax > maxCx {
		cxMax = maxCx
	}
	if cyMin < 0 {
		cyMin = 0
	}
	if cyMax > maxCy {
		cyMax = maxCy
	}
	if czMin < 0 {
		czMin = 0
	}
	if czMax > maxCz {
		czMax = maxCz
	}

	return cxMin, cxMax, cyMin, cyMax, czMin, czMax
}

func (volume *volumeStruct) Name() string {
	return "Volume"
}

func (volume *volumeStruct) Return() interface{} {
	return volume
}
