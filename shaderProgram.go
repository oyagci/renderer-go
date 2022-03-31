package main

type IShaderProgram interface {
	Bind() error
	SetupLayout(mesh OpenGLMesh)
}

type IShaderProgramBuilder interface {
	AddVertexShaderFromSource(source string) error
	AddFragmentShaderFromSource(source string) error
	Build() (IShaderProgram, error)
}
