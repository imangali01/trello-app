package handler

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Imangali2002/trello-app/config"
	"github.com/Imangali2002/trello-app/repository"
	"github.com/Imangali2002/trello-app/repository/psql"
)

type App struct {
	db *sql.DB

	userRepo repository.UserInterface
}

func (app *App) Initialize(config *config.Config) {
	db, err := psql.NewPostgresDB(psql.Config{
		Host:     config.Host,
		Port:     config.Port,
		Username: config.Username,
		Password: config.Password,
		DBName:   config.DBName,
	})

	if err != nil {
		log.Fatalf("[-] Failed to initializing configs: %s", err.Error())
	}

	app.db = db
	app.userRepo = repository.UserRepositoryInit(db)
	app.setRouters()
}

func (app *App) setRouters() {
	http.HandleFunc("/users", app.userHandler)
	http.HandleFunc("/users/", app.userHandler)
}

func (app *App) Run(port string) {
	log.Println("[+] Server started at", port)

	log.Fatal(http.ListenAndServe(port, nil))
}
