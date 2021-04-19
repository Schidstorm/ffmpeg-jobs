package dependencies

import (
	"github.com/schidstorm/ffmpeg-jobs/api/database"
	"github.com/sirupsen/logrus"
)

var Current *Collection

type Collection struct {
	Database *database.Connection
}

func InitCollection() {
	db := database.NewConnection()
	err := db.Connect()
	if err != nil {
		logrus.Error(err)
	}
	Current = &Collection{Database: db}
}
