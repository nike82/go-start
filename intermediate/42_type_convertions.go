package intermediate

import "fmt"

func main() {
	var a int = 32
	b := int32(a)
	c := float64(b)
	//d:= bool()
	p := 3.14
	f := int(p)
	fmt.Println(f, c)

	g := "Hello RR!"
	var h []byte
	h = []byte(g)
	fmt.Println(h)
	i := []byte{101, 72}
	j:= string(i)
	fmt.Println(j)
}
