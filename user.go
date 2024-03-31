package main

import (
	"context"
	"log"

	"github.com/imanaspaul/todoapp/ent"
)

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {

	u, err := client.User.Create().SetAge(230).SetName("a8m").Save(ctx)

	if err != nil {
		log.Fatalf("Error creatig the user %v", err)
		return nil, err
	}

	return u, nil
}

func GetUser(ctx context.Context, client *ent.Client) ([]*ent.User, error) {
	u, err := client.User.Query().All(ctx)
	if err != nil {
		log.Fatalf("User not found %v", err)
		return nil, err
	}
	return u, nil
}
