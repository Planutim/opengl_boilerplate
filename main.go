package main

import (
	"fmt"
	"math"
	"runtime"

	"github.com/Planutim/myopengl"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

func init() {
	runtime.LockOSThread()
}

var (
	WINDOW_WIDTH  = 800
	WINDOW_HEIGHT = 600
)

var (
	offset [2]float32
)
var offsetValue float32 = 0.01
var u_center mgl32.Vec2 = mgl32.Vec2{0.5, 0.5}
var centerIncValue float32 = 0.01
var u_force float32 = 0.2
var forceIncValue float32 = 0.01

var repeatVal int32 = gl.CLAMP_TO_EDGE

var isCircle = false

func main() {
	window := myopengl.InitGlfw(WINDOW_WIDTH, WINDOW_HEIGHT)
	window.SetKeyCallback(keyCallback)
	window.SetCursorPosCallback(mouseCallback)
	window.SetMouseButtonCallback(mousePressCallback)
	window.SetFramebufferSizeCallback(sizeCallback)

	shader := myopengl.NewShader("shaders/shader.vert", "shaders/shader.frag")

	vao := makeVao(false)

	tex := makeTexture()
	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)

		shader.Use()
		shader.SetInt("u_texture", 0)

		shader.SetVec2("u_center", &u_center)
		shader.SetFloat("u_force", u_force)

		ratio := float32(WINDOW_WIDTH) / float32(WINDOW_HEIGHT)
		// fmt.Println(ratio)

		shader.SetBool("u_circle", isCircle)

		// fmt.Println("RATIO: ", ratio)
		shader.SetFloat("u_ratio", ratio)

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

	// gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_BORDER)
	// gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_BORDER)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(img.Rect.Dx()), int32(img.Rect.Dy()), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(img.Pix))

	return tex
}

func keyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyEscape {
		w.SetShouldClose(true)
	}

	if key == glfw.KeyLeft {
		u_center = u_center.Add(mgl32.Vec2{-centerIncValue, 0})
		// offset[0] -= offsetValue
	}
	if key == glfw.KeyRight {
		u_center = u_center.Add(mgl32.Vec2{centerIncValue, 0})
		// offset[0] += offsetValue
	}
	if key == glfw.KeyUp {
		u_center = u_center.Add(mgl32.Vec2{0, -centerIncValue})
		// offset[1] += offsetValue
	}

	if key == glfw.KeyDown {
		u_center = u_center.Add(mgl32.Vec2{0, centerIncValue})
		// offset[1] -= offsetValue
	}

	if key == glfw.KeyKPAdd {
		if math.Abs(float64(u_force)) < 1 {
			u_force += forceIncValue
		} else {
			u_force += forceIncValue * u_force
		}
	}
	if key == glfw.KeyKPSubtract {
		if math.Abs(float64(u_force)) < 1 {
			u_force += -forceIncValue
		} else {
			u_force += -forceIncValue * u_force
		}
	}

	if action == glfw.Press && key == glfw.KeyY {
		isCircle = !isCircle

		fmt.Println("is circle: ", isCircle)
	}

	if action == glfw.Press && key == glfw.KeyT {
		enabled = !enabled
		fmt.Println("Mouse drag enabled: ", enabled)
	}

	if mods == glfw.ModShift {
		changed := false
		switch key {
		case glfw.Key1:
			repeatVal = gl.CLAMP_TO_EDGE
			fmt.Println("CLAMP_TO_EDGE")
			changed = true
		case glfw.Key2:
			repeatVal = gl.CLAMP_TO_BORDER
			fmt.Println("CLAMP_TO_BORDER")
			changed = true

		case glfw.Key3:
			repeatVal = gl.REPEAT
			fmt.Println("REPEAT")
			changed = true

		case glfw.Key4:
			repeatVal = gl.MIRRORED_REPEAT
			fmt.Println("MIRRORED_REPEAT")
			changed = true

		}

		if changed {
			gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, repeatVal)
			gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, repeatVal)

		}

	} else {
		fmt.Printf("center: %v, force: %v\n", u_center, u_force)
	}

}

func mouseCallback(w *glfw.Window, xpos, ypos float64) {

	if enabled {
		xoff := xpos / float64(WINDOW_WIDTH)
		yoff := ypos / float64(WINDOW_HEIGHT)
		u_center = mgl32.Vec2{float32(xoff), float32(yoff)}
		fmt.Println(xoff, "   ", yoff)
	}
}

var enabled bool = true

func mousePressCallback(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
	if action == glfw.Press && button == glfw.MouseButton1 {
		enabled = true
	}
	if action == glfw.Release && button == glfw.MouseButton1 {
		enabled = false
	}
}

func sizeCallback(w *glfw.Window, width, height int) {
	fmt.Println(width, "   w:h    ", height)
	WINDOW_WIDTH = width
	WINDOW_HEIGHT = height
	gl.Viewport(0, 0, int32(width), int32(height))
}
