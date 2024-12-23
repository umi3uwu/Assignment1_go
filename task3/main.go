package main

import (
	"github.com/umi3uwu/Assignment1_go/company"
)

func main() {
	c := company.NewCompany()

	c.AddEmployee(company.FullTimeEmployee{ID: 1, Name: "Alice", Salary: 5000})
	c.AddEmployee(company.PartTimeEmployee{ID: 2, Name: "Bob", HourlyRate: 20, HoursWorked: 100})

	c.ListEmployees()
}
