package main

import (
	"fmt"
)

type employee struct {
	name            string
	vacationBalance int
	bonus           float64
}

type manager struct {
	employee
}

func NewEmployee(name string, vacationBalance int, bonus float64) *employee {
	return &employee{name, vacationBalance, bonus}
}

type Speak interface {
	speak() string
}

func (emp *employee) setNextYearBonus(rate float64) {
	emp.bonus = emp.bonus * rate
}

func (emp *employee) speak() string {
	return "ok"
}

func (mgr *manager) setNextYearBonus(rate float64) {
	mgr.bonus = mgr.bonus * rate
}

func (mgr *manager) speak() string {
	return "cool down!"
}
func talk(head Speak) {
	fmt.Println(head)
	fmt.Println(head.speak())
}

func totalBonus(list []*employee) float64 {
	var total float64
	for _, emp := range list {
		total = total + emp.bonus
	}
	return total
}

func main() {
	empA := NewEmployee("John L", 30, 20000.00)
	fmt.Println(*empA)
	mgrB := NewEmployee("Steve J", 55, 50000.00)
	fmt.Println(*mgrB)
	empA.setNextYearBonus(1.2)
	mgrB.setNextYearBonus(1.5)
	fmt.Println(*empA)
	fmt.Println(*mgrB)

	empList := make([]*employee, 0)
	empList = append(empList, empA)
	empList = append(empList, mgrB)
	fmt.Println("Total bonus next year:", totalBonus(empList))
	talk(empA)
	talk(mgrB)
}

/*
$ go run oop.go
{John L 30 20000}
{Steve J 55 50000}
{John L 30 24000}
{Steve J 55 75000}
Total bonus next year: 99000
&{John L 30 24000}
ok
&{Steve J 55 75000}
ok
*/
