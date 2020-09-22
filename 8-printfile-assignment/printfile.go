package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	if len(os.Args) > 2 {
		fmt.Println("Error: The Program only accept 1 file to be read")
		os.Exit(1)
	}
	filename := os.Args[1]

	file, err := os.Open(filename) // For read access.
	if err != nil {
		log.Fatal(err)
	}

	//******without io.Copy******//

	// data := make([]byte, 100)
	// count, err := file.Read(data)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("read %d bytes: %q\n", count, data[:count])
	//******with io.Copy******//

	io.Copy(os.Stdout, file)

}
