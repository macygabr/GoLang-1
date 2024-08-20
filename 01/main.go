package main

type Human struct {
	yers int
	name string
}

func (*Human) move() {
	println("Human move")
}

type Action struct {
	man Human
	act bool
}

func (a *Action) move() {
	a.man.move()
}

func main() {
	man := Human{}
	man.move()
	act := Action{}
	act.move()
}
