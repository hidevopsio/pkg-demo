package world

type Greeter struct {
	Name string
}

func NewGreeter() *Greeter  {
	return &Greeter{Name: "Hello, Hiboot"}
}