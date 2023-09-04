package main

import (
	"fmt"
)

func main() {

	i := Item{1, 0}
	fmt.Printf("%#v\n", i)

	i1 := Item{
		X: 1,
		Y: 2,
	}

	fmt.Printf("%#v\n", i1)

	i2 := Item{
		X: 20,
	}
	fmt.Printf("%#v\n", i2)

	fmt.Println(NewItem(2, 3))

	fmt.Println(NewItem(-2, 1000))

	i2.Move(20, 20)
	fmt.Printf("i2's New position: %#v\n", i2)

	p1 := Player{
		Name: "Nithya",
		Item: i2,
		Keys: []Key{1},
	}

	fmt.Println(p1.X, p1.Y)

	p1.Move(300, 400)

	fmt.Println(p1.FoundKey(Jade))
	fmt.Println(p1.FoundKey(Copper))
	fmt.Println(p1.FoundKey(Crystal))
	fmt.Println(p1.FoundKey(25))

	//fmt.Printf("Player's latest position: %#v", p1.Item)

	ms := []mover{
		&i1,
		&p1,
	}

	moveAll(ms, 0, 0)
	for _, m := range ms {
		fmt.Printf("%#v\n", m)
	}

	k := Jade
	fmt.Println(k)

}

func (k Key) String() string {
	switch k {
	case Jade:
		return "Jade"
	case Copper:
		return "Copper"
	case Crystal:
		return "Crystal"

	}
	return fmt.Sprintf("<Key>:%d", k)
}

type Item struct {
	X int
	Y int
}

type Player struct {
	Name string
	Item
	Keys []Key
}

const (
	maxX = 1000
	maxY = 800
)

func NewItem(x, y int) (*Item, error) {

	if x < 0 || x > maxX || y < 0 || y > maxY {
		return nil, fmt.Errorf("%d,%d are out of bounds %d,%d", x, y, maxX, maxY)
	}

	i := Item{
		X: x,
		Y: y,
	}
	return &i, nil
}

func (i *Item) Move(x, y int) {
	i.X = x
	i.Y = y
}

func (p *Player) FoundKey(k Key) error {

	if k < Jade || k > invalidKey {
		return fmt.Errorf("not a valid key")
	}

	if containsKey(p.Keys, k) {
		return fmt.Errorf("key already exists")
	} else {
		p.Keys = append(p.Keys, k)
		return nil
	}
}

func containsKey(keys []Key, k Key) bool {
	for _, k2 := range keys {
		if k2 == k {
			return true
		}
	}
	return false
}

/*
func containsKey(keys []Key, k Key) bool {
	for _, k2 := range keys {
		if k2 == k {
			return true
		}
	}
	return false
}
*/

type mover interface {
	Move(x, y int)
}

func moveAll(ms []mover, x, y int) {
	for _, m := range ms {
		m.Move(x, y)

	}
}

type Key byte

const (
	Jade Key = iota + 1
	Copper
	Crystal
	invalidKey
)
