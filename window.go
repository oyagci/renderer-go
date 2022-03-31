package main

import "github.com/go-gl/glfw/v3.3/glfw"

type WindowProps struct {
	Width  int
	Height int
	Title  string
}

type RendererWindow struct {
	window *glfw.Window
}

func CreateWindow(props WindowProps) (*RendererWindow, error) {
	window, err := glfw.CreateWindow(props.Width, props.Height, props.Title, nil, nil)

	if err != nil {
		return nil, err
	}

	return &RendererWindow{
		window: window,
	}, nil
}

func (w *RendererWindow) GetInternal() *glfw.Window {
	return w.window
}
