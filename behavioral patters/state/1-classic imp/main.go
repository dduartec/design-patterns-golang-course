package main

import "fmt"

type Switch struct {
	State State
}

func NewSwitch() *Switch {
	return &Switch{NewOnState()}
}

func (s *Switch) On() {
	s.State.On(s)
}

func (s *Switch) Off() {
	s.State.Off(s)
}

type State interface {
	On(s *Switch)
	Off(s *Switch)
}

type BaseState struct {
}

func (b *BaseState) On(s *Switch) {
	fmt.Println("Light is already on")
}

func (b *BaseState) Off(s *Switch) {
	fmt.Println("Light is already off")
}

type OnState struct {
	BaseState
}

func NewOnState() *OnState {
	fmt.Println("Light turned on")
	return &OnState{BaseState{}}
}

func (o *OnState) Off(s *Switch) {
	fmt.Println("turning the light off...")
	s.State = NewOffState()
}

type OffState struct {
	BaseState
}

func NewOffState() *OffState {
	fmt.Println("Light turned on")
	return &OffState{BaseState{}}
}

func (o *OffState) On(s *Switch) {
	fmt.Println("turning the light on...")
	s.State = NewOnState()
}

func main() {
	s := NewSwitch()
	s.Off()
	s.Off()
	s.On()
}
