package main

import "fmt"

type EmployeeContract struct {
	EmpID         int
	Salary        float64
	TaxPercentage float64
}

type FreelancerContract struct {
	EmpID       int
	HourlyRate  float64
	HoursWorked int
}

func (e EmployeeContract) CalculateSalary() float64 {
	return e.Salary - (e.TaxPercentage * e.Salary)
}

func (f FreelancerContract) CalculateSalary() float64 {
	return f.HourlyRate * float64(f.HoursWorked)
}

type SalaryCalculator interface {
	CalculateSalary() float64
}

func main3123() {
	homer := EmployeeContract{EmpID: 1, Salary: 479.60, TaxPercentage: 0.245}
	fmt.Printf("Homer's salary: $%.2f\n", homer.CalculateSalary()) // Homer's salary: $362.10

	deadpool := FreelancerContract{EmpID: 2, HourlyRate: 50_000.00, HoursWorked: 10}
	fmt.Printf("Deadpool's salary: $%.2f\n", deadpool.CalculateSalary()) // Deadpool's salary: $500000.00

	employees := []SalaryCalculator{homer, deadpool}
	fmt.Printf("Monthly sal. expense: $%.2f\n", salaryExpense(employees)) // Monthly sal. expense: "$500362.10
}

// salaryExpense() takes a slice of []SalaryCalculator and returns the sum of all salaries
func salaryExpense(salaries []SalaryCalculator) float64 {
	totalExpense := 0.0
	for _, v := range salaries {
		totalExpense += v.CalculateSalary()
	}
	return totalExpense
}
