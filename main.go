package main

import (
	"math"
	"runtime"

	"github.com/Planutim/myopengl"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func init() {
	runtime.LockOSThread()
}
func main() {
	window := myopengl.InitGlfw(800, 600)
	window.SetKeyCallback(keyCallback)

	shader := myopengl.NewShader("shaders/shader.vert", "shaders/shader.frag")

	vao := makeVao(false)

	tex := makeTexture()
	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)

		shader.Use()
		shader.SetInt("u_texture", 0)

		gl.BindVertexArray(vao)

		gl.ActiveTexture(gl.TEXTURE0)
		gl.BindTexture(gl.TEXTURE_2D, tex)

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

func makeTexture() uint32 {
	var tex uint32
	gl.GenTextures(1, &tex)
	gl.BindTexture(gl.TEXTURE_2D, tex)
	defer gl.BindTexture(gl.TEXTURE_2D, 0)

	img := myopengl.LoadImage("water.jpg")

	// fmt.Println("STRIDE IS: ", img.Stride)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(img.Rect.Dx()), int32(img.Rect.Dy()), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(img.Pix))

	return tex
}

func keyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyEscape {
		w.SetShouldClose(true)
	}
}
