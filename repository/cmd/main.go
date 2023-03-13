package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ltvco/go-design-patterns/repository"
)

var (
	logger = log.New(
		os.Stdout,
		fmt.Sprintf("%s: ", "repository pattern"),
		log.LstdFlags|log.Lmicroseconds|log.LUTC,
	)
)

func main() {
	mysqlRepo, err := repository.NewMysqlRepository(conf.MysqlDB)
	if err != nil {
		logger.Printf("ERROR: %+v", err)
		return
	}
	ctx := context.Background()

	logger.Println("Processing data with MySQL")
	err = processRepoData(ctx, mysqlRepo)
	if err != nil {
		logger.Printf("ERROR: %+v", err)
	}

}

func processRepoData(ctx context.Context, repo repository.IRepository) error {
	logger.Println("Creating a couple of new users called Ellie and Joel")
	err := repo.CreateUser(ctx, "Ellie", 18)
	if err != nil {
		return fmt.Errorf("failed to create user Ellie: %+v", err)
	}

	err = repo.CreateUser(ctx, "Joel", 57)
	if err != nil {
		return fmt.Errorf("failed to create user Joel: %+v", err)
	}

	logger.Println("===============================================")
	logger.Println("Obtaining data from Ellie")
	user, err := repo.GetUserByName(ctx, "Ellie")
	if err != nil {
		return fmt.Errorf("failed to retrieve Ellie's data: %+v", err)
	}
	logger.Printf("User name: %v", user.Name)
	logger.Printf("User age: %v", user.Age)

	logger.Println("===============================================")
	logger.Println("Updating Ellie's age...")
	attributes := map[string]interface{}{"age": 19}
	err = repo.UpdateUser(ctx, "Ellie", attributes)
	if err != nil {
		return fmt.Errorf("failed to update Ellie's data: %+v", err)
	}

	logger.Println("===============================================")
	logger.Println("Obtaining updated data from Ellie")
	user, err = repo.GetUserByName(ctx, "Ellie")
	if err != nil {
		return fmt.Errorf("failed to retrieve Ellie's data: %+v", err)
	}
	logger.Printf("User name: %v", user.Name)
	logger.Printf("User age: %v", user.Age)
	logger.Println("Ellie's age was updated!")

	logger.Println("===============================================")
	logger.Println("Deleting created users")
	err = repo.DeleteUser(ctx, "Joel")
	if err != nil {
		return fmt.Errorf("failed to delete Joel's data: %+v", err)
	}

	err = repo.DeleteUser(ctx, "Ellie")
	if err != nil {
		return fmt.Errorf("failed to delete Ellie's data: %+v", err)
	}

	return nil
}
