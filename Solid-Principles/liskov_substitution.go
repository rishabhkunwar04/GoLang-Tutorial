package main

//**** Before LSP

/*
type Bird interface {
	Fly()
}

type Sparrow struct{}

func (s Sparrow) Fly() {
	fmt.Println("Sparrow flying")
}

type Ostrich struct{}

func (o Ostrich) Fly() {
	panic("Ostriches can't fly!") //
}

// problem: You cannot substitute Ostrich where Bird is expected without risking a panic — it violates LSP.


*/

/*
// ***** After LSP
type Flyer interface {
	Fly()
}
type Sparrow struct{}

func (s Sparrow) Fly() {
	fmt.Println("Sparrow flying")
}
type Ostrich struct{}

func (o Ostrich) Walk() {
	fmt.Println("Ostrich walking")
}
func LetItFly(f Flyer) {
	f.Fly()
}

func main() {
	var s Sparrow
	LetItFly(s)

	var o Ostrich
	// LetItFly(o) // ❌ Compile error – and that's good!
}

*/
