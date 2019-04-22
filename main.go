package main

func main() {
	vector := "vector.svg"
	flat := "vector.png"
	mesh := "mesh.stl"
	rendered := "mesh.png"
	shape(vector, mesh)
	svgtopng(vector, flat)
	render(mesh, rendered)
}
