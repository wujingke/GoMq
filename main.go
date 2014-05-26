package main

import (
	"github.com/hoisie/redis"
	"net/http"
	"runtime"
)

var client *redis.Client

func init() {
	client = &redis.Client{
		Addr:        "127.0.0.1:6379",
		Db:          0, // default db is 0
		MaxPoolSize: 10000,
	}
	runtime.GOMAXPROCS(runtime.NumCPU())
}
func main() {

	http.HandleFunc("/put", PutHandler)
	http.HandleFunc("/get", GetHandler)
	http.HandleFunc("/", HomeHandler)
	if err := http.ListenAndServe(":1988", nil); err != nil {
		panic("ListenAndServe: " + err.Error())
	}

}

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("welcome to use GoMq!"))
}

func PutHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	if err := client.Lpush("queue", []byte(req.FormValue("data"))); err != nil {
		w.Write([]byte("error"))
		return
	}
	w.Write([]byte("ok"))
	return
}
func GetHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	value, err := client.Rpop("queue")
	if err == nil {
		w.Write([]byte(value))
		return
	}
	w.Write([]byte("end"))
	return
}
