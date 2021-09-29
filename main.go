package main

import (
	"math"

	"github.com/Planutim/myopengl"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func main() {
	window := myopengl.InitGlfw(800, 600)
	window.SetKeyCallback(keyCallback)

	shader := myopengl.NewShader("shaders/shader.vert", "shaders/shader.frag")

	vao := makeVao(false)

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)

		shader.Use()

		gl.BindVertexArray(vao)

		gl.DrawArrays(gl.TRIANGLES, 0, 6)

		glfw.PollEvents()
		window.SwapBuffers()
	}
}

func makeVao(inverted bool) uint32 {

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	defer gl.BindVertexArray(0)

	vertices := []float32{
		-1, -1, 0, 1,
		-1, 1, 0, 0,
		1, -1, 1, 1,

		-1, 1, 0, 0,
		1, -1, 1, 1,
		1, 1, 1, 0,
	}

	if inverted {
		for i := range vertices {
			if i%4 == 3 {
				vertices[i] = float32(math.Mod(float64(vertices[i]+1), 2))
			}
		}
	}
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 4, gl.FLOAT, false, 0, gl.PtrOffset(0))

	return vao
}

func keyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyEscape {
		w.SetShouldClose(true)
	}
}
