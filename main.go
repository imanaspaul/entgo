package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/imanaspaul/todoapp/ent"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open("sqlite3", "db.sqlite3?mode=memory&cache=shared&_fk=1")

	if err != nil {
		log.Fatalf("failed to open a connection to sqlite3 %v", err)
	}

	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("Failed to creating scheam resources, %v", err)
	}
	user, err := GetUser(context.Background(), client)
	if err != nil {
		log.Fatalf("Failed to create the user %v", err)
	}
	res_user, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(res_user))
}
