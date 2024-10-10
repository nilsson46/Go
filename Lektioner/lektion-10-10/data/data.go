package data

var employees []Employee

func GetAllEmployees() []Employee {
	return employees
}

func CreateNewEmployee(employee Employee) {
	employees = append(employees, employee)
}

func GetEmployeeById(id int) *Employee {
	for _, employee := range employees {
		if employee.Id == id {
			return &employee
		}
	}
	return nil
}

func Init() {
	employees = append(employees, Employee{Id: 1, Name: "John", Age: 25})
	employees = append(employees, Employee{Id: 2, Name: "Paul", Age: 23})
	employees = append(employees, Employee{Id: 3, Name: "Steve", Age: 30})
}
