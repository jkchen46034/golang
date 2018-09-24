// This file demonstrates subclassing via embedded typing,
// and polymorphism via interface

package main

import (
	"fmt"
)

type Introduce interface {
	introduce()
}

type CalculateBonus interface {
	calculateBonus() float64
}

type Employee struct {
	name   string
	salary float64
	bonus  float64
}

func NewEmployee(name string, salary float64, bonus float64) *Employee {
	return &Employee{name, salary, bonus}
}

type Manager struct {
	Employee
	office string
}

func (emp *Employee) calculateBonus() float64 {
	emp.bonus = 2000 - emp.salary*0.01
	return emp.bonus
}

func (emp *Employee) introduce() {
	fmt.Println("My name " + emp.name + ", and I am solid")
}

func NewManager(name string, salary float64, bonus float64, office string) *Manager {
	return &Manager{Employee{name, salary, bonus}, office}
}

func (mgr *Manager) calculateBonus() float64 {
	mgr.bonus = 8000 + mgr.salary*0.10
	return mgr.bonus
}

func (mgr *Manager) introduce() {
	for i := 0; i < 2; i++ {
		fmt.Println("Manager speaking here: My name "+mgr.name+", office at", mgr.office)
	}
}

func calculateTotalBonus(emps ...CalculateBonus) float64 {
	var sum float64
	for _, emp := range emps {
		sum = sum + emp.calculateBonus()
	}
	return sum
}

func everyBodyIntroduce(emps ...Introduce) {
	for _, emp := range emps {
		emp.introduce()
	}
}

func main() {
	empA := NewEmployee("John Jarvis", 80000.00, 0)
	mgrB := NewManager("Steve Young", 150000.00, 0, "corner office E")

	fmt.Println(*empA)
	fmt.Println(*mgrB)

	everyBodyIntroduce(empA, mgrB)

	sum := calculateTotalBonus(empA, mgrB)
	fmt.Printf("Total bonus is %d\n", int64(sum))

	empList := []CalculateBonus{empA, mgrB}
	sum = calculateTotalBonus(empList...)
	fmt.Printf("Total bonus is %d\n", int64(sum))

	fmt.Println(*empA)
	fmt.Println(*mgrB)
}

/*
$ go run oop.go
{John Jarvis 80000 0}
{{Steve Young 150000 0} corner office E}
My name John Jarvis, and I am solid
Manager speaking here: My name Steve Young, office at corner office E
Manager speaking here: My name Steve Young, office at corner office E
Total bonus is 24200
{John Jarvis 80000 1200}
{{Steve Young 150000 23000} corner office E}
*/
