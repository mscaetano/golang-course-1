package main

import "fmt"

func main() {
	x := 20.3032

	xs := fmt.Sprint(x)
	fmt.Println("O Mathesu tem: " + xs + " anos")
	fmt.Println("O Matheus tem: ", x, "anos")
	fmt.Printf("O Matheus tem: %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
