package main

import (
	"context"
	"fmt"
	"log"

	"github.com/imanaspaul/todoapp/ent"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")

	if err != nil {
		log.Fatalf("failed to open a connection to sqlite3 %v", err)
	}

	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("Failed to creating scheam resources, %v", err)
	}
	user, err := createUser(context.Background(), client)
	if err != nil {
		log.Fatalf("Failed to create the user %v", err)
	}
	fmt.Println(user)
}

func createUser(ctx context.Context, client *ent.Client) (*ent.User, error) {

	u, err := client.User.Create().SetAge(20).SetName("Manas").Save(ctx)

	if err != nil {
		log.Fatalf("Error creatig the user %v", err)
		return nil, err
	}

	return u, nil
}
