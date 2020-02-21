package main

import (
  "os"
  "fmt"
  "net/http"
  "log"
  "encoding/json"
  // go moduleで管理
  "app-standards/go/src/db"
  _ "github.com/go-sql-driver/mysql"
)

//type Task struct{
	//Id	     int
	//Title    string
	//Body     string
	//CreateAt string
	//UpdateAt string
	//IsDelete bool
//
//}

func getAllEvents(w http.ResponseWriter, r *http.Request) {
//  task := Task{}

  dbInfo, err := db.Init()

  if err != nil{
    log.Printf("db init failed: %v", err)
		os.Exit(1)
  }
  fmt.Println(*dbInfo)
  task, err := db.GetTasks(dbInfo)
  if err != nil{
    log.Printf("db get row failed: %v", err)
    os.Exit(1)
  }
  res,_ := json.Marshal(task)

  //以下CORSの設定
  //許可するアドレスの指定
  w.Header().Set( "Access-Control-Allow-Origin", "*")
  //許可するメソッドの指定
  w.Header().Set( "Access-Control-Allow-Methods","GET, POST, PUT, DELETE, OPTIONS" )
  //使用可能なリクエストヘッダ
  w.Header().Set( "Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization" )
  w.Header().Set("Content-Type", "application/json")
  w.Write(res)
}

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", getAllEvents)
  mux.HandleFunc("/api/v1/events", getAllEvents)

  http.ListenAndServe(":8080", mux)
}