package intermediate

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
)

//go:embed 33_example.txt
var content string

//go:embed basics
var basicsFolder embed.FS

func main() {

	fmt.Println("Embedded content:", content)
	content, err := basicsFolder.ReadFile("basics/hello.go")
	if err != nil {
		return
	}
	fmt.Println("Embeded file content:", string(content))
	err = fs.WalkDir(basicsFolder, "basics", func(path string, d fs.DirEntry, err error) error{
		if err != nil {
			return err
		}
		fmt.Println(path)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
