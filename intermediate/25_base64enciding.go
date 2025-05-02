package intermediate

import (
	"encoding/base64"
	"fmt"
)

func main() {

	data := []byte("He~lo, Base64 Encoding")
	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println(encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error in decoding:", err)
		return
	}
	fmt.Println("Decocded:", string(decoded))

	urlSafeEncoded := base64.URLEncoding.EncodeToString(data)
	fmt.Println("URL Safe encoded:", urlSafeEncoded)


}
