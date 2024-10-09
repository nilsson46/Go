package employee

//Kan inte ha medelmsfunktioner i säg.
type Employee struct {
	Name       string
	Age        int
	Street     string
	PostalCode string
	City       string
	//func salary går ej då alltså
}

func New(name string, age int) Employee {
	return Employee{
		Name: name,
		Age:  age,
	}
}

/*func CalculateSalary(employee Employee) int {
	return employee.Age * 100;
} */

func (employee Employee) CalculateSalary() int { //Detta blir en medelmsfunktion för Employee.
	return employee.Age * 100
}
