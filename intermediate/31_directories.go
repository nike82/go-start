package intermediate

import (
	"fmt"
	"os"
	"path/filepath"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	//checkError(os.Mkdir("subdir", 0755))
	//defer os.RemoveAll("subdir")

	//os.WriteFile("subdir1/file", []byte(""), 0755)
	os.MkdirAll("subdir1/parent/child1", 0755)
	os.MkdirAll("subdir1/parent/child2", 0755)
	os.MkdirAll("subdir1/parent/child3", 0755)
	os.WriteFile("subdir1/parent/file", []byte(""), 0755)
	os.WriteFile("subdir1/parent/child1/file", []byte(""), 0755)

	result, err := os.ReadDir("subdir1/parent")
	checkError(err)
	for _, entry := range result {
		fmt.Println(entry.Name(), entry.IsDir(), entry.Type())
	}
	dir, err := os.Getwd()
	checkError(err)
	fmt.Println(dir)
	checkError(os.Chdir("subdir1/parent/child1"))
	result, err = os.ReadDir(".")
	checkError(err)
	fmt.Println("+++++++++++++++++++++++")
	for _, entry := range result {
		fmt.Println(entry)
	}
	checkError(os.Chdir("../../.."))
	dir1, err := os.Getwd()
	checkError(err)
	fmt.Println(dir1)

	pathFile := "subdir1"
	fmt.Println("------Walking dir------")
	err = filepath.WalkDir(pathFile, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		fmt.Println(path)
		return nil
	})
	checkError(err)

	checkError(os.RemoveAll("./subdir"))
	checkError(os.RemoveAll("./subdir1"))
}
