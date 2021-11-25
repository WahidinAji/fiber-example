package database

const (
	user = "root"
	password = ""
	port = "3306"
	url = "127.0.0.1"
	dbname = "fiber_gorm"
)
func Database() string {
	dsn := "root@tcp(127.0.0.1:3306)/fiber_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	// fmt.Println("Connection Opened to Database")
	return dsn
}