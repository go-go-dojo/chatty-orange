package main

type Room struct {
	ID   int
	Name string
	Users map[string]interface{}
	Hub  *Hub
}

func (r *Room) AddUser(user string) {
	r.Users[user] = user
}

func CreateRoom(id int, name string) *Room {
	hub := newHub()
	go hub.run()

	return &Room{
		Name: name,
		ID:   id,
		Hub:  hub,
		Users: make(map[string]interface{}),
	}
}
