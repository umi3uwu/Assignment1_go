package company

import "fmt"

type Employee interface {
	GetDetails() string
}

type FullTimeEmployee struct {
	ID     uint64
	Name   string
	Salary uint32
}

func (f FullTimeEmployee) GetDetails() string {
	return fmt.Sprintf("Full Time Employee - ID: %d, Name: %s, Salary: %d", f.ID, f.Name, f.Salary)
}

type PartTimeEmployee struct {
	ID          uint64
	Name        string
	HourlyRate  uint64
	HoursWorked float32
}

func (p PartTimeEmployee) GetDetails() string {
	return fmt.Sprintf("Part Tim eEmployee - ID: %d, Name: %s, Hourly Rate: %d, Hours Worked: %.1f", p.ID, p.Name, p.HourlyRate, p.HoursWorked)
}

type Company struct {
	Employees map[string]Employee
}

func NewCompany() *Company {
	return &Company{
		Employees: make(map[string]Employee),
	}
}

func (c *Company) AddEmployee(emp Employee) {
	switch e := emp.(type) {
	case FullTimeEmployee:
		c.Employees[fmt.Sprint(e.ID)] = emp
	case PartTimeEmployee:
		c.Employees[fmt.Sprint(e.ID)] = emp
	default:
		fmt.Println("Unsupported employee type")
	}
}

func (c *Company) ListEmployees() {
	fmt.Println("Employee List:")
	for _, emp := range c.Employees {
		fmt.Println(emp.GetDetails())
	}
}
