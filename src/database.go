package api

import (
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// ServerAdress - Adresse of Mongo server
const ServerAdress = "localhost:27017"

// DataBase - Name of Mongo database
const DataBase = "ccac-db"

// TransactionCollection - Name of Transaction collection
const TransactionCollection = "transactions"

// GetTransactions - Return all transactions
func GetTransactions() ([]Transaction, error) {
	session := openSession()
	defer closeSession(session)
	transactionCollection := getTransactionCollection(session)
	transactions := Transactions{}
	err := transactionCollection.Find(nil).Sort("-date").All(&transactions)
	return transactions, err
}

// CreateTransaction - Create a transaction
func CreateTransaction(newTransaction Transaction) error {
	session := openSession()
	defer closeSession(session)
	transactionCollection := getTransactionCollection(session)
	newTransaction.ID = bson.NewObjectId()
	err := transactionCollection.Insert(newTransaction)
	return err
}

// GetUsers - Return all users
func GetUsers() ([]string, error) {
	session := openSession()
	defer closeSession(session)
	transactionCollection := getTransactionCollection(session)
	users := []string{}
	err := transactionCollection.Find(nil).Distinct("user", &users)
	return users, err
}

// openSession - return mongo session
func openSession() *mgo.Session {
	session, err := mgo.Dial(ServerAdress)
	if err != nil {
		log.Fatal(err)
	}
	return session
}

// closeConnection - close mongo session
func closeSession(session *mgo.Session) {
	session.Close()
}

// getTransactionCollection - return mongo transaction collection
func getTransactionCollection(session *mgo.Session) *mgo.Collection {
	return session.DB(DataBase).C(TransactionCollection)
}
