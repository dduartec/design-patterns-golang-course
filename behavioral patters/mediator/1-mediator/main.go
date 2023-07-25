package main

import "fmt"

type ChatRoom struct {
	people []*Person
}

func (c *ChatRoom) Broadcast(source, message string) {
	for _, p := range c.people {
		if p.Name != source {
			p.Receive(source, message)
		}
	}
}

func (c *ChatRoom) Message(source, destination, message string) {
	for _, p := range c.people {
		if p.Name == destination {
			p.Receive(source, message)
		}
	}
}

func (c *ChatRoom) Join(p *Person) {
	joinMsg := p.Name + " has joined the chat"
	c.Broadcast("Room", joinMsg)

	p.Room = c
	c.people = append(c.people, p)
}

type Person struct {
	Name    string
	Room    *ChatRoom // mediator
	chatlog []string
}

func NewPerson(name string) *Person {
	return &Person{Name: name}
}

func (p *Person) Receive(sender, message string) {
	s := fmt.Sprintf("%s: %s", sender, message)
	fmt.Printf("[%s's chat session]: %s \n", p.Name, s)
	p.chatlog = append(p.chatlog, s)
}

func (p *Person) Say(message string) {
	p.Room.Broadcast(p.Name, message)
}

func (p *Person) PrivateMessage(receiver, message string) {
	p.Room.Message(p.Name, receiver, message)
}

func main() {

	room := ChatRoom{}

	p1 := NewPerson("Diego")
	p2 := NewPerson("Dan")

	room.Join(p1)
	room.Join(p2)

	p1.Say("Hola!!")
	p2.PrivateMessage("Diego", "Hola Diego :D")
}
