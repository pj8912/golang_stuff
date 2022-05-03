package main

import "fmt"


//struct declaration
type User struct{

	//defining struct variables

	user_name string
	user_age int
}


//a function to print the user_details
func (user User) user_details(){
	fmt.Printf("User name:  %s", user.user_name)
	fmt.Printf("\nUser age: %d \n", user.user_age)
}

//main

func main(){
	
	user_1 := User{"Stephen Strange", 43}
	user_1.user_details()


}
