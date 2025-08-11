package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type email struct {
	Email string
	Phone string
}

type address struct {
	House string
	City  string
	State string
}

type Employee struct {
	Person_Details  Person
	Email_Details   email
	Address_Details address
}

func main() {
	var person1 Person
	fmt.Println("Person1:", person1)
	person1.Name = "Shahmeer"
	person1.Age = 25
	fmt.Println("Person1:", person1)

	//directly initializing struct
	person2 := Person{
		Name: "Bob",
		Age:  30,
	}
	fmt.Println("Person2:", person2)

	//new keyword to create a pointer to a struct
	person3 := new(Person)
	person3.Name = "Alice"
	person3.Age = 28
	fmt.Println("Person3:", *person3)

	fmt.Println("**************************************")

	//creating a struct with multiple fields
	var employee1 Employee
	employee1.Person_Details = Person{
		Name: "John",
		Age:  35,
	}

	employee1.Email_Details = email{
		Email: "Example@gmail.com",
		Phone: "1234567890",
	}

	employee1.Address_Details.City = "New York"
	employee1.Address_Details.House = "123 Main St"
	employee1.Address_Details.State = "NY"
	fmt.Println("Employee1:", employee1)
	fmt.Println("Employee1 Name:", employee1.Person_Details)
	fmt.Println("Employee1 Email:", employee1.Email_Details)
	fmt.Println("Employee1 Address:", employee1.Address_Details)

}
