package main

import (
	"fmt"
	"log"
	"os"

	"github.com/asynched/geti/domain/repositories"
	"github.com/asynched/geti/http/controllers"
	"github.com/asynched/geti/infra/db"
	"github.com/asynched/geti/infra/env"
	"github.com/gofiber/fiber/v2"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile | log.Lmsgprefix)
	log.SetPrefix(fmt.Sprintf("[%d] [geti] ", os.Getpid()))
}

func main() {
	app := fiber.New(fiber.Config{DisableStartupMessage: true})

	log.Println("Initializing database")
	db := db.CreateClient(env.GetDatabaseUrl())
	defer db.Close()

	log.Println("Initializing repositories")
	linkRepository := repositories.NewLinkRepository(db)
	visitRepository := repositories.NewVisitRepository(db)

	log.Println("Initializing controllers")
	linkController := controllers.NewLinkController(linkRepository)
	visitController := controllers.NewVisitController(linkRepository, visitRepository)
	healthController := controllers.NewHealthController()

	log.Println("Initializing routes")
	app.Post("/links", linkController.Create)
	app.Get("/links/:slug", visitController.Create)
	app.Get("/links/:slug/visits", visitController.ListAll)
	app.Get("/health", healthController.Get)

	url := env.GetUrl()

	log.Println("Starting server")
	log.Printf("Server is listening on: http://%s\n", url)
	log.Printf("Check health status at: http://%s/health\n", url)
	if err := app.Listen(url); err != nil {
		log.Fatal("Error:", err)
	}
}
