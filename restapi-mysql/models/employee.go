package models

import (
	"encoding/json"
)

type Employee struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"` 
	Salary float64 `json:"salary"`
	Age    int     `json:"age"`
}
// type Employees []Employee

// Employees struct
type Employees struct {
	Employees []Employee `json:"employees"`
}

func (b Employees) MarshalJSON() ([]byte, error)  {
	type Alias Employees

	a := struct {
		Alias
	}{
		Alias: (Alias)(b),
	}

	if a.Employees == nil {
		a.Employees = make([]Employee, 0)
	}
	return json.Marshal(a)
}


var EmployeesSlice = []Employee{
	// {Id: 1, Name: "Aji", Salary: 19.99, Age: 22},
}