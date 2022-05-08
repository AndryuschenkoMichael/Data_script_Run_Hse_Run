package model

type Vertex struct {
	Id       int
	Type     string
	RoomId   int
	CampusId int
}

type Vertices map[int]Vertex

func NewVertices() Vertices {
	vertices := make(Vertices)
	return vertices
}