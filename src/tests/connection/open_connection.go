package connection

import (
    "context"
    "log"
    "os"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func OpenConnection() (*mongo.Database, func()) {
    mongoURL := os.Getenv("MONGODB_URL")
    if mongoURL == "" {
        mongoURL = "mongodb://localhost:27017"
    }

    client, err := mongo.NewClient(options.Client().ApplyURI(mongoURL))
    if err != nil {
        log.Fatal("Error creating mongo client:", err)
    }

    if err := client.Connect(context.Background()); err != nil {
        log.Fatal("Error connecting to mongo:", err)
    }

    dbName := os.Getenv("MONGODB_USER_DB")
    if dbName == "" {
        dbName = "test_user"
    }

    closeFunc := func() {
        _ = client.Disconnect(context.Background())
    }

    return client.Database(dbName), closeFunc
}
