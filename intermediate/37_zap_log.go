package intermediate

// go get go.uber.org/zap
import (
	"log"

	"go.uber.org/zap"
)

func main() {

	logger, err := zap.NewProduction()
	if err != nil {
		log.Println("Error zap init")
	}
	defer logger.Sync()

	logger.Info("This is an info message.")
	logger.Info("User logged in", zap.String("username", "John Full"),zap.String("method", "GET"))
}
