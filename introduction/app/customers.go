package main
type(
	Customer struct {
		FirstName string
		LastName string
		FullName string
	}
)

func GetCustomers()(customers []Customer){
	marcel := Customer{ FirstName: "Marcel", LastName: "Dempers" }
	customers = append(customers, marcel, 
	Customer{FirstName:"Ben", LastName:"Spain"},
	Customer{FirstName:"Aleen", LastName:"Janmohamed"},
	Customer{FirstName:"Jamie", LastName:"le Notre"},
	Customer{FirstName:"Victor", LastName:"Savkov"},
	Customer{FirstName:"P", LastName:"The Admin"},
	Customer{FirstName:"Adrian", LastName:"Opera"},
	Customer{FirstName:"Jonathan", LastName:"D"},
)

return customers
}