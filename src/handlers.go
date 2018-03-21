package api

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var connectedClients = make(map[*websocket.Conn]bool)
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Welcome Handler
func Welcome(writer http.ResponseWriter, request *http.Request) {
	startTime := time.Now()
	writer.Header().Set("content-type", "application/json")
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(`{"message": "Welcome to CCAC API !"}`))
	stopTime := time.Now()
	printLog(request, startTime, stopTime)
}

// GetTransactionsHandler - Return all transactions
func GetTransactionsHandler(writer http.ResponseWriter, request *http.Request) {
	startTime := time.Now()
	writer.Header().Set("content-type", "application/json")
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	transactions, err := GetTransactions()
	if err != nil {
		log.Fatal(err)
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(`{"error": "Read transactions on database failed"}`)
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(transactions)
	stopTime := time.Now()
	printLog(request, startTime, stopTime)
}

// CreateTransactionHandler - Create a transaction
func CreateTransactionHandler(writer http.ResponseWriter, request *http.Request) {
	startTime := time.Now()
	// Delay for simulate loading webapp
	time.Sleep(3 * time.Second)
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.Header().Add("Access-Control-Allow-Origin", "*")

	var transaction Transaction
	body, errRead := ioutil.ReadAll(io.LimitReader(request.Body, 1048576))
	errClose := request.Body.Close()
	errJSON := json.Unmarshal(body, &transaction)

	if errRead != nil || errClose != nil || errJSON != nil {
		log.Fatal("Error bad data transaction")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(`{"error": "bad data transaction"}`)
	}

	errCreate := CreateTransaction(transaction)
	if errCreate != nil {
		log.Fatal(errCreate)
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(`{"error": "Create transaction on database failed"}`)
	}
	// Send transactions at all clients
	SendTransactions(connectedClients)

	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(transaction)
	stopTime := time.Now()
	printLog(request, startTime, stopTime)
}

// GetUsersHandler - Return all users
func GetUsersHandler(writer http.ResponseWriter, request *http.Request) {
	startTime := time.Now()
	// Delay for simulate loading webapp
	time.Sleep(3 * time.Second)
	writer.Header().Set("content-type", "application/json")
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	users, err := GetUsers()
	if err != nil {
		log.Fatal(err)
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(`{"error": "Read users on database failed"}`)
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(users)
	stopTime := time.Now()
	printLog(request, startTime, stopTime)
}

// WebSockectConnection - New WebSocket connection
func WebSockectConnection(writer http.ResponseWriter, request *http.Request) {
	startTime := time.Now()
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Register our new client
	connectedClients[ws] = true
	// Send transaction our new client
	SendTransactions(map[*websocket.Conn]bool{ws: true})
	stopTime := time.Now()
	printLog(request, startTime, stopTime)
}

// SendTransactions - send transactions at clients
func SendTransactions(clients map[*websocket.Conn]bool) {
	transactions, err := GetTransactions()
	if err != nil {
		log.Fatal(err)
	}
	for client := range clients {
		err := client.WriteJSON(transactions)
		if err != nil {
			client.Close()
			delete(clients, client)
		}
	}
}

func printLog(request *http.Request, start time.Time, stop time.Time) {
	log.Printf("[%s] %q %v\n", request.Method, request.URL.String(), stop.Sub(start))
}
