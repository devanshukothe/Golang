package main

import "fmt"

//jwtToken := 300000 not allowed

//var jwtToken int := 3000 allowed

const LoginToken string = "xyz" //pubic



func main()  {
	var username string = "Devanshu"
	fmt.Println(username)
	fmt.Printf("Variables is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variables is of type: %T \n", isLoggedIn)
   

	var smallVal int = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.42739842838532984
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)
   

	// default values and aliases
	var anotherVariables int
	fmt.Println(anotherVariables)
	fmt.Printf("Variables is of type: %T \n", anotherVariables)
   

	//implicit type
	var website = "Devanshukothe.in"
	fmt.Println(website)

	// no var style 

	numberOfUser := 30000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variables is of type: %T \n", LoginToken)
}