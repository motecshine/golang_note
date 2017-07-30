package extends

import (
	"fmt"
)
// interface
type Printer interface {
	Print()
}

type user struct {
	name string
	age byte
}



type manager struct {
	user
	title string
}



func (u user) ToString() string {
	return fmt.Sprintf("%+v", u)
}

func (u user) Print(){
	fmt.Printf("%+v", u)
}

func Extends() {
	var m manager
	m.name = "Tom"
	m.age = 14
	println(m.ToString())
	
	var u user
	u.name = "ssds"
	u.age = 15
	var p Printer = u 
	p.Print()
}

