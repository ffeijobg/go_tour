package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

var ma = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

var mb = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	fmt.Println(ma)

	fmt.Println(mb)

	mc := make(map[string]int)

	mc["Answer"] = 42
	fmt.Println("The value:", mc["Answer"])

	mc["Answer"] = 48
	fmt.Println("The value:", mc["Answer"])

	delete(mc, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := mc["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
