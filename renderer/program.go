package renderer

type ShaderProgram interface {
}

type ProgramBuilder interface {
	AddVertexShaderSource(source string) (ProgramBuilder, error)
	AddFragmentShaderSource(source string) (ProgramBuilder, error)
	Build() (ShaderProgram, error)
}

type GLShaderProgram struct {
	id int32
}

type GLProgramBuilder struct {
	vertexShaderSrc   string
	fragmentShaderSrc string
}

func NewGLShaderProgramBuilder() GLProgramBuilder {
	return GLProgramBuilder{}
}

func (builder *GLProgramBuilder) AddVertexShaderSource(source string) (GLProgramBuilder, error) {
	builder.vertexShaderSrc = source
	return *builder, nil
}

func (builder *GLProgramBuilder) AddFragmentShaderSource(source string) (GLProgramBuilder, error) {
	builder.fragmentShaderSrc = source
	return *builder, nil
}
