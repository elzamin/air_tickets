package user

import (
	"context"
	"log"

	"github.com/elzamin/air_tickets/user/internal/entity"
)

 func TestUserDb(ctx context.Context,userRepository *repository) {

	err := userRepository.TouchTable(ctx)
	if err != nil {
		log.Fatal("Failed to touch table user", err)
	}

	err = userRepository.Create(ctx, entity.User{
		Id: "1", 
		Name: "Elzamin", 
		Age: 28, 
		Address: "Presnya", 
		Work: "Empty"})
	if err != nil {
		log.Fatal("Failed to create user", err)
	}

	user, err := userRepository.Get(ctx, "1")
	if err != nil {
		log.Fatal("Failed to get user", err)
	}
	log.Println(user)

	err = userRepository.Update(ctx, entity.User{
		Id: "1", 
		Name: "Valera", 
		Age: 35, 
		Address: "Homovniki", 
		Work: "Police"})
	if err != nil {
		log.Fatal("Failed to update user", err)
	}

	user, err = userRepository.Get(ctx, "1")
	if err != nil {
		log.Fatal("Failed to get user", err)
	}
	log.Println(user)

	err = userRepository.Create(ctx, entity.User{
		Id: "2", 
		Name: "Elzamin", 
		Age: 28, 
		Address: "Presnya", 
		Work: "Empty"})
	if err != nil {
		log.Fatal("Failed to create user", err)
	}

	users, _ := userRepository.GetAll(ctx)
	log.Println(users)

	err = userRepository.Delete(ctx, "1")
	if err != nil {
		log.Fatal("Failed to delete user", err)
	}

	err = userRepository.Delete(ctx, "2")
	if err != nil {
		log.Fatal("Failed to delete user", err)
	}
}
