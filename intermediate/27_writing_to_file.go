package intermediate

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	data := []byte("Hello World!\n\n\n")
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Data has been written to file")

	file1, err := os.Create("writeStr.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file1.Close()
	_, err = file1.WriteString("Hello Nike!\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Data has been written to file")
}
