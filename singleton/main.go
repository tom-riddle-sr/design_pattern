package main

import (
	"fmt"
	"sync"
)

// Singleton struct
type Database struct {
	connection string
}

var instance *Database
var once sync.Once

// GetInstance returns the single instance of Database (thread-safe)
func GetInstance() *Database {
	once.Do(func() {
		fmt.Println("Creating singleton instance...")
		instance = &Database{
			connection: "Database Connection Established",
		}
	})
	return instance
}

func (db *Database) Query(sql string) string {
	return fmt.Sprintf("[%s] Executing: %s", db.connection, sql)
}

func main() {
	// Get singleton instance multiple times
	db1 := GetInstance()
	db2 := GetInstance()
	db3 := GetInstance()

	// Verify all references point to the same instance
	fmt.Println("\nInstance addresses:")
	fmt.Printf("db1: %p\n", db1)
	fmt.Printf("db2: %p\n", db2)
	fmt.Printf("db3: %p\n", db3)

	// Use the singleton
	fmt.Println("\nExecuting queries:")
	fmt.Println(db1.Query("SELECT * FROM users"))
	fmt.Println(db2.Query("INSERT INTO logs VALUES (...)"))
	fmt.Println(db3.Query("UPDATE config SET value = 'new'"))

	// Confirm they're the same instance
	fmt.Println("\nAre they the same instance?", db1 == db2 && db2 == db3)
}
