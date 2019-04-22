package main

import (
	"encoding/base64"
	"github.com/disintegration/imaging" // FauxGL
	fgl "github.com/fogleman/fauxgl"    // FauxGL
	"github.com/nfnt/resize"            // FauxGL
	ic "image/color"                    // FauxGL
	// pt "github.com/fogleman/pt/pt"         // Path Tracer
	"log"
	"os"
)

func render(stl, png string) {

	file, err := os.Create(png)
	if err != nil {
		log.Fatalln(err)
	}
	if stl == "" {
		empty, _ := base64.StdEncoding.DecodeString("iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mP8z/C/HgAGgwJ/lK3Q6wAAAABJRU5ErkJggg==")
		_, _ = file.Write(empty)
		if err != nil {
			log.Fatalln(err)
		}
		return
	}
	file.Close()

	// FauxGL
	mesh, err := fgl.LoadSTL(stl)
	if err != nil {
		log.Fatalln(err)
	}
	const (
		scale  = 1
		width  = 1000
		height = 1000
		fovy   = 30
		near   = 1
		far    = 10
	)
	var (
		eye    = fgl.V(5, 0, 0)
		center = fgl.V(0, -0.07, 0)
		up     = fgl.V(0, 1, 0)
		light  = fgl.V(1, 0, -0.7).Normalize()
		color  = fgl.HexColor("#02f2b4")
	)
	mesh.BiUnitCube()
	mesh.SmoothNormalsThreshold(fgl.Radians(30))
	context := fgl.NewContext(width*scale, height*scale)
	context.ClearColorBufferWith(fgl.HexColor("#000000"))
	aspect := float64(width) / float64(height)
	matrix := fgl.LookAt(eye, center, up).Perspective(fovy, aspect, near, far)
	shader := fgl.NewPhongShader(matrix, light, eye)
	shader.ObjectColor = color
	context.Shader = shader
	context.DrawMesh(mesh)
	image := context.Image()
	image = resize.Resize(width, height, image, resize.Bilinear)
	image = imaging.Rotate(image, 90, ic.RGBA{0, 0, 0, 1})
	fgl.SavePNG(png, image)
	// FauxGL

	// Path Tracer
	// scene := pt.Scene{}
	// material := pt.DiffuseMaterial(pt.White)
	// blue := pt.DiffuseMaterial(pt.Color{0.3, 0.3, 1})
	// plane := pt.NewPlane(pt.V(0, 0, 0), pt.V(0, 0, 1), material)
	// scene.Add(plane)
	// object, _ := pt.LoadSTL(stl, blue)
	// scene.Add(object)
	// light := pt.NewSphere(pt.V(3, -3, 5), 1, pt.LightMaterial(pt.White, 10))
	// scene.Add(light)
	// camera := pt.LookAt(pt.V(2, 1, 2.5), pt.V(0, 0, 1), pt.V(0, 0, 3), 50)
	// sampler := pt.NewSampler(4, 4)
	// renderer := pt.NewRenderer(&scene, &camera, sampler, 900, 900)
	// renderer.AdaptiveSamples = 128
	// renderer.IterativeRender("mesh.png", 3)
	// Path Tracer
}
