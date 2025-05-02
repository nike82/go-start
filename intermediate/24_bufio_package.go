package intermediate

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(strings.NewReader("Hello, bufio package!!!\n S'up dude?"))
	data := make([]byte, 20)
	n, err := reader.Read(data)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}
	fmt.Printf("Read %d byts: %s\n", n, data[:n])

	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading string:", err)
		return
	}
	fmt.Println("Read string", line)

	writer := bufio.NewWriter(os.Stdout)
	data1 := []byte("Hello, bufio package!\n")
	n1, err := writer.Write(data1)
	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}
	fmt.Printf("Wrote %d bytes\n", n1)
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing:", err)
		return
	}

	str1:="This is a string\n"
	n, err = writer.WriteString(str1)
	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}
	fmt.Printf("Wrote %d bytes\n", n)
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}
}
