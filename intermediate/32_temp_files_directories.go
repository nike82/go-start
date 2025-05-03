package intermediate

import (
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	// tempFileName := "tempFile"

	// tempFile, err := os.CreateTemp("", tempFileName)
	// checkError(err)
	// fmt.Println("Temp file is created:", tempFile.Name())
	// defer tempFile.Close()
	// defer os.Remove(tempFile.Name())

	tempDir, err := os.MkdirTemp("", "NikeTempDir")
	checkError(err)
	defer os.Remove(tempDir)
	fmt.Println("Temp dir created:", tempDir)
}
