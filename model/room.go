package model

type Room struct {
	Id       int
	Code     string
	CampusId int
}

type Rooms map[int]Room

func NewRooms() Rooms {
	rooms := make(Rooms)
	return rooms
}
