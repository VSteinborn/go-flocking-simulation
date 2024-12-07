package main

type Vector struct{
	x, y float64
}

func (a Vector) Add(b Vector) Vector{
	a.x += b.x 
	a.y += b.y
	return a
}

func (a *Vector) AddAssign(b Vector) {
	a.x += b.x 
	a.y += b.y
}

func (a Vector) Scale(scalar float64) Vector{
	a.x *= scalar
	a.y *= scalar
	return a
}

func (a *Vector) ScaleAssign(scalar float64){
	a.x *= scalar
	a.y *= scalar
}