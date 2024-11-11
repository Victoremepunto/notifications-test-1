// main.go
package main

import (
	"log"
	"notifications-test-1/database"
	"notifications-test-1/models"
)

func main() {
	// Initialize the database
	database.Init()

	// Run migrations
	err := database.DB.AutoMigrate(
		&models.Team{},
		&models.User{},
		&models.Product{},
		&models.Incident{},
		&models.Subscription{},
		&models.Message{},
	)
	if err != nil {
		log.Fatal("failed to run migrations:", err)
	}

	log.Println("Database migration completed successfully.")

	productA := &models.Product{Name: "Product A"}
	productB := &models.Product{Name: "Product B"}
	userA := &models.User{Name: "User A", Email: "usera@foobar.com"}
	userB := &models.User{Name: "User B", Email: "userb@foobar.com"}
	teamA := &models.Team{Name: "Team A"}
	teamB := &models.Team{Name: "Team B"}
	incidentA := &models.Incident{Name: "Incident A"}
	incidentB := &models.Incident{Name: "Incident B"}

	database.DB.Create(productB)
	database.DB.Create(userA)
	database.DB.Create(userB)
	database.DB.Create(teamA)
	database.DB.Create(teamB)
	database.DB.Create(incidentA)
	database.DB.Create(incidentB)

	subscriptionA := &models.Subscription{
		TeamID:     nil,
		UserID:     &userA.ID,
		ProductID:  &productA.ID,
		IncidentID: nil,
		Type:       "slack",
	}
	subscriptionB := &models.Subscription{
		TeamID:     nil,
		UserID:     &userA.ID,
		ProductID:  &productA.ID,
		IncidentID: nil,
		Type:       "email",
	}
	subscriptionC := &models.Subscription{
		TeamID:     nil,
		UserID:     &userA.ID,
		ProductID:  nil,
		IncidentID: &incidentA.ID,
		Type:       "slack",
	}
	subscriptionD := &models.Subscription{
		TeamID:     &teamA.ID,
		UserID:     nil,
		ProductID:  nil,
		IncidentID: &incidentB.ID,
		Type:       "email",
	}

	createSubscription(subscriptionA)
	createSubscription(subscriptionB)
	createSubscription(subscriptionC)
	createSubscription(subscriptionD)

	subscriptionA_1 := &models.Subscription{
		TeamID:     nil,
		UserID:     &userA.ID,
		ProductID:  &productA.ID,
		IncidentID: nil,
		Type:       "slack",
	}

	createSubscription(subscriptionA_1)

	database.DB.Delete(productA)
	database.DB.Delete(productB)
	database.DB.Delete(userA)
	database.DB.Delete(userB)
	database.DB.Delete(teamA)
	database.DB.Delete(teamB)
	database.DB.Delete(incidentA)
	database.DB.Delete(incidentB)

}

func createSubscription(subscription *models.Subscription) {
	if err := database.DB.Create(subscription).Error; err != nil {
		log.Fatal(err)
	} else {
		log.Println("Subscription created successfully.")
	}
}
