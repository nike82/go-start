package intermediate

import (
	"fmt"
	"strconv"
)

func main() {

	numStr := "1234"
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Error parsing the value:", err)
	}
	fmt.Println("Parsed num:", num)
	fmt.Println("Parsed num:", num+1)

	numistr, err := strconv.ParseInt(numStr, 10, 32)
	if err != nil {
		fmt.Println("Error parsing the value:", err)
	}
	fmt.Println("Parsed int:", numistr)

	floatStr := "3.14"
	floatVal, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		fmt.Println("Error parsing the value:", err)
	}
	fmt.Printf("Parsed float: %.2f\n", floatVal)

	binaryStr := "1010"
	decimal, err := strconv.ParseInt(binaryStr, 2, 64)
	if err != nil {
		fmt.Println("Error parsing the value:", err)
		return
	}
	fmt.Println("Parsed binary to decimal:", decimal)

	hexStr := "FF"
	hex, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		fmt.Println("Error parsing the value:", err)
		return
	}
	fmt.Println("Parsed hex to decimal:", hex)

	invalidNum := "6aFF"
	invalidPars, err := strconv.Atoi(invalidNum)
	if err != nil {
		fmt.Println("Error parsing value:", err)
		return
	}
	fmt.Println("Parsed invalid num to decimal:", invalidPars)

}
