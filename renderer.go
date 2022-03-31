package main

type Renderer struct {
}

type RendererBackend uint

const (
	None RendererBackend = iota + 1
	OpenGLBackend
)

func NewRenderer(backend RendererBackend) Renderer {
	return Renderer{}
}
