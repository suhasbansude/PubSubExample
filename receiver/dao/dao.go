package dao

import (
	"../model"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func failOnError(err error, msg string) {
	if err != nil {
		log.Errorf("%s: %s", msg, err)
	}
}

func InitDB(userName string, password string, host string, port string, databaseName string) {
	dsn := userName + ":" + password + "@tcp(" + host + ":" + port + ")/" + databaseName + "?charset=utf8mb4&parseTime=True&loc=Local"
	dbObj, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	failOnError(err, "Fail To connect database")
	dbObj.AutoMigrate(&model.Amenity{})
	dbObj.AutoMigrate(&model.Room{})
	dbObj.AutoMigrate(&model.CancellationPolicy{})
	dbObj.AutoMigrate(&model.OtherCondition{})
	dbObj.AutoMigrate(&model.RatePlan{})
	dbObj.AutoMigrate(&model.Hotel{})
	db = dbObj
}

func SaveOfferDataDAO(myData model.DataHolder) error {
	db.Create(&myData.Offers[0].Hotel)
	return nil
}
