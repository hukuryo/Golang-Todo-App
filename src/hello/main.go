package main

import "github.com/labstack/echo/v4"

var i int
i = 1

i := 1

var (
    a int
    b int
    c string
)

// 定数
const TEST = "TEST"

func hoge() {
fmt.Println("hoge!!")
}

func add(x int, y int) int {
	return x + y
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func named(a, b string) (x, y string) {
	x = "x: " + a
	y = "y: " + b
	return
}

func main() {
	e := echo.New()
	e.GET("/hello", func(c echo.Context) error {
		v := Vertex{3, 4}
		fmt.Println(v.X)
		fmt.Println(v.Abs())
		
		v.Scale(10)
		fmt.Println(v.X, v.Y)
		return c.String(200, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}