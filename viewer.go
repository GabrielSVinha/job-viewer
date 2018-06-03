package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"strconv"

	"github.com/go-redis/redis"
)

type CachedConn []Connection

type Connection struct {
	Client *redis.Client
	Name   string
	Queue  Queue
}
type Queue struct {
	Name string
}

func main() {
	cache := CachedConn{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		uri := strings.Split(r.RequestURI, "/")
		if uri[3] == "count" {
			HandleCount(w, r, cache, uri)
		} else {
			HandleQueue(w, r, cache, uri)
		}
	})
	fmt.Println("Starting web server on port: 8080")
	http.ListenAndServe(":8080", nil)
}

func HandleQueue(w http.ResponseWriter, r *http.Request, cache CachedConn, uri []string) {
	// fmt.Println("Incoming request: /" + conn.Queue.Name)

	var conn Connection
	conn, err := ReturnConn(cache, uri[1])

	if err != nil {
		// Connection not found
		client := redis.NewClient(&redis.Options{
			Addr:     uri[1] + ":6379",
			Password: "",
			DB:       0,
		})
		conn = Connection{
			Client: client,
			Name:   uri[1],
			Queue: Queue{
				uri[2],
			},
		}
		cache = append(cache, conn)
	}
	responseArr, err := conn.Client.LRange(conn.Queue.Name, 0, -1).Result()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	response := ""
	for _, element := range responseArr {
		response += element
		response += "\n"
	}
	w.Write([]byte(response))
}

func HandleCount(w http.ResponseWriter, r *http.Request, cache CachedConn, uri []string) {
	// fmt.Println("Incoming request: /" + conn.Queue.Name + "/count")
	var conn Connection
	conn, err := ReturnConn(cache, uri[1])

	if err != nil {
		// Connection not found
		client := redis.NewClient(&redis.Options{
			Addr:     uri[1] + ":6379",
			Password: "",
			DB:       0,
		})
		conn = Connection{
			Client: client,
			Name:   uri[1],
			Queue: Queue{
				uri[2],
			},
		}
		cache = append(cache, conn)
	}
	response, err := conn.Client.LLen(conn.Queue.Name).Result()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	w.Write([]byte(strconv.Itoa(int(response))))
}

func ReturnConn(cache CachedConn, name string) (Connection, error) {
	var connection Connection

	for _, conn := range cache {
		if conn.Name == name {
			return conn, nil
		}
	}
	return connection, errors.New("Connection not found")
}
