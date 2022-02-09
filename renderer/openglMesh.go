package renderer

import "unsafe"

type Position3 struct {
	X float32
	Y float32
	Z float32
}

type TriangleVertex struct {
	Position Position3
	Color    Position3
}

type OpenGLMeshData struct {
	Vertices []TriangleVertex
	Indices  []uint32
}

type OpenGLMesh struct {
	data OpenGLMeshData

	verticesSize   int
	verticesOffset int
	indicesSize    int
	indicesOffset  int
}

func NewOpenGLMesh(vertices []TriangleVertex, indices []uint32) OpenGLMesh {
	vertexSize := uint64(unsafe.Sizeof(vertices[0]))
	verticesSize := int(vertexSize) * len(vertices)

	mesh := OpenGLMesh{
		data: OpenGLMeshData{
			Vertices: vertices,
			Indices:  indices,
		},
		verticesSize:   verticesSize,
		verticesOffset: 0,
		indicesSize:    4 * len(indices),
		indicesOffset:  verticesSize,
	}

	return mesh
}

func (mesh OpenGLMesh) GetData() OpenGLMeshData {
	return mesh.data
}

func (mesh OpenGLMesh) GetVerticesSize() int {
	return mesh.verticesSize
}

func (mesh OpenGLMesh) GetVerticesOffset() int {
	return mesh.verticesOffset
}

func (mesh OpenGLMesh) GetIndicesSize() int {
	return mesh.indicesSize
}

func (mesh OpenGLMesh) GetIndicesOffset() int {
	return mesh.indicesOffset
}
