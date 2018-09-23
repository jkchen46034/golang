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
	for i := 0; i < 2; i++ {
		fmt.Println("My name " + emp.name + ", and I am solid")
	}
}

func NewManager(name string, salary float64, bonus float64, office string) *Manager {
	return &Manager{Employee{name, salary, bonus}, office}
}

func (mgr *Manager) calculateBonus() float64 {
	mgr.bonus = 8000 + mgr.salary*0.10
	return mgr.bonus
}

func (mgr *Manager) introduce() {
	for i := 0; i < 3; i++ {
		fmt.Println("Manager speaking here: My name "+mgr.name+", office is at", mgr.office)
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

func totalBonus(emps []*Employee) float64 {
	var total float64
	for _, emp := range emps {
		total = total + emp.bonus
	}
	return total
}

func main() {
	empA := NewEmployee("John L", 80000.00, 0)
	mgrB := NewManager("Steve J", 150000.00, 0, "corner office E")

	fmt.Println(empA)
	fmt.Println(mgrB)

	everyBodyIntroduce(empA, mgrB)

	fmt.Printf("Total bonus is %f\n", calculateTotalBonus(empA, mgrB))
}

/*
$ go run oop.go
&{John L 80000 0}
&{{Steve J 150000 0} corner office E}
My name John L, and I am solid
My name John L, and I am solid
Manager speaking here: My name Steve J, office is at corner office E
Manager speaking here: My name Steve J, office is at corner office E
Manager speaking here: My name Steve J, office is at corner office E
Total bonus is 24200.000000

*/
