package intermediate

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	var strFlag string
	flag.StringVar(&strFlag, "user", "John", "Name of user")

	subcommand1 := flag.NewFlagSet("firstSub", flag.ExitOnError)
	subcommand2 := flag.NewFlagSet("secondtSub", flag.ExitOnError)
	firstFlag := subcommand1.Bool("processing", false, "Command processing status")
	secondFlag := subcommand1.Int("bytes", 1024, "Byte lenth of result")

	flagSc2 := subcommand2.String("language", "go", "enter your language")
	if len(os.Args) < 2 {
		fmt.Println("This program needs additional commands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "firstSub":
		subcommand1.Parse(os.Args[2:])
		fmt.Println("subcommand1:")
		fmt.Println("processing:", *firstFlag)
		fmt.Println("bytes:", *secondFlag)
	case "secondSub":
		subcommand2.Parse(os.Args[2:])
		fmt.Println("subcommand2:")
		fmt.Println("language:", *flagSc2)
	default:
		fmt.Println("no subcommand enterd!")
		os.Exit(1)
	}
}
