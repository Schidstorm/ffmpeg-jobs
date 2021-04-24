package dependencies

import (
	"github.com/schidstorm/ffmpeg-jobs/api/database"
	"github.com/sirupsen/logrus"
)

var Current *Collection

type Collection struct {
	Database *database.Connection
}

func InitConnection(dialectorName, dsn string) {
	db := database.NewConnection()
	err := db.Connect(dialectorName, dsn)
	if err != nil {
		logrus.Error(err)
		panic(err)
	}
	Current = &Collection{Database: db}
}
