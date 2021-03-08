package main

type Room struct {
	ID   int
	Name string
	Hub  *Hub
}

func CreateRoom(id int, name string) Room {
	hub := newHub()
	go hub.run()

	return Room{
		Name: name,
		ID:   id,
		Hub:  hub,
	}
}


