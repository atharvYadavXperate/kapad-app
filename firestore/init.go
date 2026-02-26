package db

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

func NewDatabase(ctx context.Context, projectId string) (*Database, error) {
	once.Do(func() {
		c, err := firestore.NewClient(ctx, projectId)
		log.Println("New Database Connection")
		if err != nil {
			initErr = err
			return
		}
		instance = &Database{
			Client: c,
		}
	})
	if initErr != nil {
		return nil, initErr
	}
	return instance, nil
}
