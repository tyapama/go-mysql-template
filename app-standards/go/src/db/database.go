package db

import(
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Task struct{
	ID	     int `json:"ID"`
	Title    string `json:"Title"`
	Body     string `json:"Body"`
	CreateAt string `json:"CreateAt"`
	UpdateAt string `json:"UpdateAt"`
	IsDelete bool `json:"IsDelete"`

}

func Init()(*sql.DB, error) {
	db, err := sql.Open("mysql","root:root@tcp(db-container:3306)/react_todo_app")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return db, nil
}

func getConnectionString() string {
    host := getParamString("MYSQL_DB_HOST", "127.0.0.1")
    port := getParamString("MYSQL_PORT", "3306")
    user := getParamString("MYSQL_USER", "root")
    pass := getParamString("MYSQL_PASSWORD", "root")
    dbname := getParamString("MYSQL_DB", "react_todo_app")
    protocol := getParamString("MYSQL_PROTOCOL", "tcp")

    return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s",
        user, pass, protocol, host, port, dbname,)
}

func getParamString(param string, defaultValue string) string {
	env := os.Getenv(param)
	if env != "" {
		return env
	}
	return defaultValue
}

func GetTasks(Db *sql.DB) (task Task, err error) {
	task = Task{}
	err = Db.QueryRow("select * from tasks").Scan(&task.ID, &task.Title, &task.Body, &task.CreateAt, &task.UpdateAt, &task.IsDelete)
	return 
}

