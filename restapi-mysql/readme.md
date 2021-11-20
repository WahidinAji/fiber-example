### Go Fiber with MySql.
* Standart Rest Api
* CRUD Employees

* Can't update if all data is same with latest data
* update can't show the current id (its bug)

## header
`Content-Type: application/json`
# Get All `http://localhost:3000/api/employees` Method `GET`
# Create One `http://localhost:3000/api/employees` 
* Method `POST` 
* Body 
```json
{
    "name": "Wahidin",
    "salary": 50.44,
    "age": 22
}
```
# Get By Id `http://localhost:3000/api/employees/1` Method `GET`
# Update One `http://localhost:3000/api/employees/1`
* Method `PUT` 
* Body 
```json
{
    "name": "Aji",
    "salary": 50.55,
    "age": 25
}
```
# Delete One `http://localhost:3000/api/employees/1` Method `DELETE`