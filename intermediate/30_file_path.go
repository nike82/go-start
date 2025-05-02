package intermediate

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	relativePath := "./data/file.txt"
	absolutePath := "/home/user/docs/file.txt"
	joinedPath := filepath.Join("home", "Docsumens", "downloads", "file.zip")
	fmt.Println("Joined path:", joinedPath)
	normalizedPath := filepath.Clean("./data/../data/file.txt")
	fmt.Println("Normalized Path:", normalizedPath)
	dir, file := filepath.Split("/home/user/docs/file.txt")
	fmt.Println("file:", file)
	fmt.Println("Dir:", dir)
	fmt.Println(filepath.Base("/home/user/docs/file.txt"))

	fmt.Println("Is relative:", filepath.IsAbs(relativePath))
	fmt.Println("Is absolute:", filepath.IsAbs(absolutePath))

	fmt.Println(filepath.Ext(file))
	fmt.Println(strings.TrimSuffix(file, filepath.Ext(file)))

	rel, err := filepath.Rel("a/b", "a/b/t/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/c", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	absPath, err := filepath.Abs(relativePath)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Absolute path:", absPath)
	}

}
