package intermediate

import (
	"errors"
	"fmt"
)

func main() {

	if err := doSomthing(); err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println("Operation completed!")

}

type customError struct {
	code    int
	message string
	er      error
}

func (e *customError) Error() string {
	return fmt.Sprintf("Error %d: %s, %v\n", e.code, e.message, e.er)
}

func doSmth() error {
	return &customError{
		code:    500,
		message: "Smth went wrong!",
	}
}

func doSomthing() error {
	if err := doSmthElse(); err != nil {
		return &customError{
			code:    500,
			message: "Smth went wrong",
			er:      err,
		}
	}
	return nil
}

func doSmthElse() error {
	return errors.New("Internal error")
}
