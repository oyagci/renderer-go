package glBuffers

type EShaderDataType uint64

const (
	Int32 EShaderDataType = iota + 1
	Int64
	Uint32
	Uint64
	Float32
	Float64
	Vector2f
	Vector3f
	Vector4f
)

type BufferElement struct {
	Name           string
	ShaderDataType EShaderDataType
	Size           uint64
	Normalized     bool
	Offset         uint32
}

type BufferLayout struct {
	elements []BufferElement
	stride   int32
}

func NewBufferLayout(elements []BufferElement) BufferLayout {
	bufferLayout := BufferLayout{
		elements: elements,
		stride:   0,
	}

	bufferLayout.calculateStrideAndOffset()

	return bufferLayout
}

func (bl *BufferLayout) calculateStrideAndOffset() {
	//var offset uint64 = 0
	var stride int32 = 0

	for _, elem := range bl.elements {
		//elem.Offset = offset
		stride += int32(elem.Size)
		//offset += elem.Size
	}

	bl.stride = stride
}

func (bl *BufferLayout) GetStride() int32 {
	return bl.stride
}

func (bl *BufferLayout) GetElements() []BufferElement {
	return bl.elements
}
