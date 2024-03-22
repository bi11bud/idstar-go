package configs

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	model "idstar.com/app/models"
)

var DB *gorm.DB

const dbIds = "ptids"

func InitDB() {
	// Create connection to postgres
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatal("Error when try connecting to database: ", err.Error())
		panic(err)
	}

	// Active debug mode
	db.Debug()

	// Check availability the database
	res := db.Exec("SELECT 'CREATE DATABASE " + dbIds + "' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = '" + dbIds + "')")

	if res.Error != nil {
		log.Fatal("Error when check availability the database: ", err.Error())
		panic(err)
	}

	if res.RowsAffected > 0 {
		err = db.Exec("CREATE DATABASE " + dbIds).Error
		if err != nil {
			if err != nil {
				log.Fatal("Error when creating the database: ", err.Error())
				panic(err)
			}
		}
		log.Print("Database successfully created..")
	} else {
		log.Print("Database already detected..")
	}

	dsn = "host=localhost user=postgres password=postgres dbname=" + dbIds + " port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatal("Error when open connection to database: ", err.Error())
		panic(err)
	}

	db.AutoMigrate(&model.DetailKaryawanEntity{})
	db.AutoMigrate(&model.KaryawanEntity{})
	db.AutoMigrate(&model.RekeningEntity{})
	db.AutoMigrate(&model.TrainingEntity{})
	db.AutoMigrate(&model.KaryawanTrainingEntity{})
	db.AutoMigrate(&model.UserEntity{})

	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
