package challenge4

import (
	"fmt"
	"log"
	"os"
)

func Example() {
	f, err := os.Open("4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	plaintext, _ := DetectSingleCharacterXOR(f)

	fmt.Printf("%s\n", plaintext)
	// Output: Now that the party is jumping
}
