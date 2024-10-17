package data

type Employee struct {
	Id   int
	Name string `json:"name"`
	Age  int
	City string
}

func (emp Employee) CalculateSalary() int {
	if emp.Name == "John" {
		return emp.Age * 2000
	}
	return emp.Age * 1000
}
