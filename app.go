package main

import (
	"context"
	"database/sql"

	"log"
	"mosaic/database"

	_ "modernc.org/sqlite"
)

// App struct
type App struct {
	ctx context.Context
	db  *database.Queries
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	conn, err := sql.Open("sqlite", "./sql/db.sqlite")

	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	a.db = database.New(conn)
}

// THIS IS A SAMPLE FUNCTION LATER TO DELETE

// Greet returns a greeting for the given name
// func (a *App) Greet(name string) string {
// 	return fmt.Sprintf("Hello %s, It's show time!", name)
// }

func (a *App) OpenFileExplorer() {
	OpenFileExplorer(a)
}
