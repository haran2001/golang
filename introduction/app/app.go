package main

import (
	"fmt"
)


func main(){

	customers := GetCustomers()

	for _, customer := range customers{
		fmt.Println(customer)
	}

}


func getData()(customers []string){
	customers = []string{"Marcel Dempers", "Bob Smith", "John Smith"}
	customers = append(customers, "Ben Spain")
	customers = append(customers, "Aleem Janmohamed")
	customers = append(customers, "Jamie le Norte")
	customers = append(customers, "P The admin")
	customers = append(customers, "Victor Savkov")
	customers = append(customers, "Adrian Opera")
	customers = append(customers, "Jonathan D")
	
	for _, customer := range customers{
		fmt.Println(customer)
	}
	
	return customers
}


