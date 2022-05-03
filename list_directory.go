package main


import (
	"fmt"
	"io/ioutil"
	"log"
)

/* log package: og in Golang implements the simple logging package. It defines a type, Logger, with methods for formatting output. Golang Log will be helpful in the critical scenarios in real-time applications

*/


/* in io/ioutil:'ReadDir' reads the directory named by dirname and returns a list of sorted directory entries */


func main(){
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		fmt.Println(f.Name())
	}
}



