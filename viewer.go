package main

import(
	"github.com/go-redis/redis"
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/jobs", HandleJobs)
	http.HandleFunc("/results", HandleResults)
	http.HandleFunc("/processing", HandleProcessing)
	fmt.Println("Starting web server on port: 8080")
	http.ListenAndServe(":8080", nil)
}

func HandleProcessing(w http.ResponseWriter, r *http.Request) {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	responseArr, err := client.LRange("job:processing", 0, -1).Result()
	if err != nil{
		http.Error(w, err.Error(), 500)
	}
	response := ""
	for _, element := range responseArr{
		response += element
	}
	w.Write([]byte(response))
}

func HandleJobs(w http.ResponseWriter, r *http.Request) {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	responseArr, err := client.LRange("job", 0, -1).Result()
	if err != nil{
		http.Error(w, err.Error(), 500)
	}
	response := ""
	for _, element := range responseArr{
		response += element
	}
	w.Write([]byte(response))
}


func HandleResults(w http.ResponseWriter, r *http.Request) {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	responseArr, err := client.LRange("job:results", 0, -1).Result()
	if err != nil{
		http.Error(w, err.Error(), 500)
	}
	response := ""
	for _, element := range responseArr{
		response += element
	}
	w.Write([]byte(response))
}
