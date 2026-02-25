package app

import db "github.com/atharvYadavXperate/kapad-app/firestore"

type App struct {
	ProjectId string
	Database  *db.Database
}
