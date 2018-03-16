package api

import "gopkg.in/mgo.v2/bson"
import "time"

// Transaction - Represent a transaction by a user
type Transaction struct {
	ID   bson.ObjectId `json:"id" bson:"_id,omitempty"`
	User string        `json:"user" bson:"user"`
	Amount int         `json:"amount" bson:"amount"`
	Date time.Time     `json:"date" bson:"date"`
}

// Transactions - Array of Transaction
type Transactions []Transaction
