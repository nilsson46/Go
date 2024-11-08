package data

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

//var employees []Employee

func GetAllEmployees() []Employee {
	var employees []Employee
	db.Find(&employees)
	return employees
}

func CreateNewEmployee(newEmployee Employee) {
	db.Create(&newEmployee)
	//employees = append(employees, employee)
}

func UpdateEmployee(employee Employee) {
	db.Save(&employee)
}

func GetEmployeeById(id int) *Employee {
	var employee Employee
	db.First(&employee, id)
	return &employee
}

func DeleteEmployee(employee Employee) {
	db.Delete(&employee)
}

/*for i, emp := range

func GetEmployeeById(id int) *Employee {
	var employee Employee
	db.First(&employee, id)
	return &employee
	/*for _, employee := range employees {
		if employee.Id == id {
			return &employee
		}
	}
	return nil*/

var db *gorm.DB

func Init(file, server, database, username, password string, port int) {

	db, _ = gorm.Open(sqlite.Open(file), &gorm.Config{})

	db.AutoMigrate(&Employee{})

	var antal int64
	db.Model(&Employee{}).Count(&antal)
	if antal == 0 {
		db.Create(&Employee{Id: 1, Name: "John", Age: 25})
		db.Create(&Employee{Id: 2, Name: "Paul", Age: 23})
		db.Create(&Employee{Id: 3, Name: "Steve", Age: 30})

		/*employees = append(employees, Employee{Id: 1, Name: "John", Age: 25})
		employees = append(employees, Employee{Id: 2, Name: "Paul", Age: 23})
		employees = append(employees, Employee{Id: 3, Name: "Steve", Age: 30}) */
	}
}
