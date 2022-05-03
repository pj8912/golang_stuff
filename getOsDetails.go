package main

import(
	"fmt"
	"runtime"
)

func main(){

	// := is declaration + assignment
	// = asusal for assignment
	os := runtime.GOOS
	switch os {
		case "windows":
			fmt.Println("Windows")
		case "linux":
			fmt.Println("Linux")
		default:
			fmt.Printf("%s.\n", os)
	}
}
