package dbconnection

import (
	models "github.com/mohammadrahimi/PayMent_Service_Go/src/Infrastructure/Persistence.Sql/Models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DataBasePostgres interface {
	DBPostgres() *gorm.DB
}

type PostgresConnection struct {
	ConnectionString string
}

func NewPostgresConnection(ConnectionString string) *PostgresConnection{
   return &PostgresConnection{
	     ConnectionString: ConnectionString,
   }
}

func(cn *PostgresConnection) DBPostgres() (*gorm.DB,error){

    var db *gorm.DB 

	db, err := gorm.Open(postgres.Open(cn.ConnectionString), &gorm.Config{})
    if err != nil { 
      panic("failed to connect database") 
    } 

    err = db.Table("PayMent").AutoMigrate(&models.PayMentEntity{}) 
    if err != nil { 
      return  nil,err
    } 
      
    return db,nil
}