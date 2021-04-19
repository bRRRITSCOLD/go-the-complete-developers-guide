// package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

// func main() {
// 	source, err := os.Open(os.Args[1])
// 	if err != nil {
// 		fmt.Println("Error:", err.Error())
// 		os.Exit(1)
// 	}
// 	defer source.Close()
// 	io.Copy(os.Stdout, source)
// }
package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	filename := os.Args[1] // Taking filename argument from command line

	file, err := os.Open(filename) // Opening the file

	if err != nil {
		log.Fatal(err)
	}

	file_data := make([]byte, 99999) // Making byte slice

	file.Read(file_data) //Reading the contents of file into byte slice //created above

	if err != nil {
		log.Fatal(err)
	}

	output_data := strings.Split(string(file_data), "\n") // Converting //byte slice to string and creating slice by splitting using new lines

	for _, ele := range output_data {
		fmt.Println(ele)
	}
}
