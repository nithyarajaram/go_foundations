package main

import "fmt"

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
	}

	fmt.Println(p1.X, p1.Y)

	p1.Move(300, 400)

	//fmt.Printf("Player's latest position: %#v", p1.Item)

	ms := []mover{
		&i1,
		&p1,
	}

	moveAll(ms, 0, 0)
	for _, m := range ms {
		fmt.Printf("%#v", m)
	}

}

type Item struct {
	X int
	Y int
}

type Player struct {
	Name string
	Item
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

type mover interface {
	Move(x, y int)
}

func moveAll(ms []mover, x, y int) {
	for _, m := range ms {
		m.Move(x, y)

	}
}
