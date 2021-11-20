package controllers

import (
	"database/sql"
	_ "database/sql"
	"log"
	"strconv"

	"github.com/WahidinAji/fiber-example/restapi-mysql/database"
	"github.com/WahidinAji/fiber-example/restapi-mysql/helpers"
	"github.com/WahidinAji/fiber-example/restapi-mysql/models"
	"github.com/gofiber/fiber/v2"
)

func GetAll(ctx *fiber.Ctx) error {
	//Get Employee list from database
	rows, err := database.Db.Query("SELECT id, name, salary, age FROM employees order by id")

	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	
	/* response with data: {"employees" : []}
	{
		"code": 200,
		"data": {
			"employees": []
		},
		"message": "Success",
		"status": true
	}
	*/
	result := models.Employees{}
	for rows.Next() {
		employee := models.Employee{}
		err := rows.Scan(&employee.Id, &employee.Name, &employee.Salary, &employee.Age)
		if err != nil {
			return err // Exit if we get an error
		}
		// Append Employee to Employees
		result.Employees = append(result.Employees, employee)
		log.Println(result)
	}
	return helpers.ResponseJson(ctx, 200, "Success", true, &result)
	/* response without data: "employees" or == data: []
	{
		"code": 200,
		"data": [],
		"message": "Success",
		"status": true
	}
	*/
	// result := []models.Employee{}
	// for rows.Next() {
	// 	employee := models.Employee{}
	// 	if err := rows.Scan(&employee.Id,&employee.Name,&employee.Salary,&employee.Age); err != nil {
	// 		return err // Exit if we get an error
	// 	}
	// 	log.Println(employee)
	// 	// Append Employee to Employees
	// 	result = append(result, employee)
	// }
	// return helpers.ResponseJson(ctx, 200, "Success", true, result)
	// return helpers.ResponseJson(ctx, 200, "Success", true, models.EmployeesSlice)
}

func CreateEmployee(ctx *fiber.Ctx) error {
	rows, err := database.Db.Query("SELECT id FROM Employees ORDER BY id DESC LIMIT 1")
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	//new employee struct
	newEmployee := new(models.Employee)

	//parse body into struct
	if err := ctx.BodyParser(newEmployee); err != nil {
		return ctx.Status(503).SendString(err.Error())
	}
	for rows.Next() {
		employee := models.Employee{}
		if err := rows.Scan(&employee.Id); err != nil {
			return err // Exit if we get an error
		}
		// log.Println(employee)
		// Append Employee to Employees
		newEmployee.Id = employee.Id +1
	}

	//update Employee record in db
	response, err := database.Db.Query("INSERT INTO employees (ID, NAME, SALARY, AGE) VALUES (?, ?, ?, ?)", newEmployee.Id, newEmployee.Name, newEmployee.Salary, newEmployee.Age)
	if err != nil {
		return err
	}

	defer response.Close()
	
	//print result
	return helpers.ResponseJson(ctx, 201, "Created Successfuly!!", true, newEmployee)
}

func GetById(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}
	rows, err := database.Db.Query("SELECT id, name, salary, age FROM employees WHERE id = ?", id)
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	result := models.Employees{}
	if rows.Next() { 					//u can for beside if
		employee := models.Employee{}
		err := rows.Scan(&employee.Id, &employee.Name, &employee.Salary, &employee.Age)
		if err != nil {
			return err // Exit if we get an error
		}
		// Append Employee to Employees
		result.Employees = append(result.Employees, employee)
		log.Println(result)
		return helpers.ResponseJson(ctx, 200, "Success", true, result)
	} else {
		return  helpers.ResponseJson(ctx, 404, "Not Found", true, result)
	}
}

//belum
func UpdateEmployee2(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}
	// tx, err := database.Db.Begin()
	tx, err := database.Db.BeginTx(ctx.Context(), &sql.TxOptions{Isolation: sql.LevelLinearizable})
	if err != nil {
		log.Fatal(err)
	}

	response, execErr := tx.Exec("UPDATE employees SET name=?,salary=?,age=? WHERE id = ?", "paid", id)
	if execErr != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil{
			// log.Fatalf("update failed: %v, unable to rollback: %v\n", execErr, rollbackErr)
			return helpers.ResponseJson(ctx,500, rollbackErr.Error(), false, []models.Employees{})
		}
		// log.Fatalf("update failed: %v", execErr)
		return helpers.ResponseJson(ctx,500, execErr.Error(), false, []models.Employees{})
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
	return helpers.ResponseJson(ctx,200,"ok",true,response)
}

func UpdateEmployee(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}
	employee := new(models.Employee)
	// Parse body into struct
	if err := ctx.BodyParser(employee); err != nil {
		return ctx.Status(400).SendString(err.Error())
	}
	response, err := database.Db.Exec("UPDATE employees SET name=?,salary=?,age=? WHERE id = ?",employee.Name,employee.Salary,employee.Age,id)
	if err != nil {
		return err
	}
	rows, err := response.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	if rows != 1 {
		return helpers.ResponseJson(ctx,404, strconv.Itoa(int(rows)) + " Data was not found",true,[]models.Employees{})
	}
	return helpers.ResponseJson(ctx,200,"Updated success!!!",true,employee)
}

func DeleteEmployee(ctx *fiber.Ctx) error {
	id , err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}
	
	res, err := database.Db.Exec("DELETE FROM employees WHERE id=?", id)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	if rows != 1 {
		return helpers.ResponseJson(ctx,404,strconv.Itoa(int(rows)) + " Data was not found",true,[]models.Employees{})
		// log.Fatalf("expected to affect 1 row, affected %d", rows)
	}
	return helpers.ResponseJson(ctx,200,"Deleted success",true,[]models.Employees{})
}