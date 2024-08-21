package payment

import (
	"errors"

	queries "github.com/mohammadrahimi/PayMent_Service_Go/src/Domain.Contract/Queries/PayMent/ResultPayMent"
	aggregate "github.com/mohammadrahimi/PayMent_Service_Go/src/Domain/PayMent"
	repository "github.com/mohammadrahimi/PayMent_Service_Go/src/Domain/PayMent/Repository"
	"github.com/mohammadrahimi/PayMent_Service_Go/src/Domain/PayMent/ValueObject/OrderId"
	"github.com/mohammadrahimi/PayMent_Service_Go/src/Domain/PayMent/ValueObject/PayMentId"
	"github.com/mohammadrahimi/PayMent_Service_Go/src/Domain/PayMent/ValueObject/UserId"
	models "github.com/mohammadrahimi/PayMent_Service_Go/src/Infrastructure/Persistence.Sql/Models"
	"gorm.io/gorm"
)

var (

	ErrorFailCreate = errors.New("PayMent Not Create! ") 
	ErrorFailDelete = errors.New("PayMent Not Delete! ")
)

type PayMentRepository struct {
	db *gorm.DB
}

func NewPayMentRepository(db *gorm.DB) repository.IPayMentRepository {
	return &PayMentRepository{db: db}
}
 
func (p *PayMentRepository) Create(payment *aggregate.PayMent) (status string, err error) {

	dbPayMent := fromDBPayMent(payment)

	res := p.db.Table("PayMent").Create(dbPayMent)
	if res.RowsAffected > 0 {
		return "Status.200", nil
	}
	return "Status.400", ErrorFailCreate
}

func (p *PayMentRepository) FindAll() ([]queries.ResultPayMentQuery, error) {

	var dbPayMents []models.PayMentEntity
	p.db.Table("PayMent").Find(&dbPayMents)
	payments := make([]queries.ResultPayMentQuery, len(dbPayMents))
	for i, payment := range dbPayMents {
		payments[i] = fromDbPayMentQuery(&payment)
	}
	return payments, nil
}

func (p *PayMentRepository) FindByPayMentId(id string) (queries.ResultPayMentQuery, error) {

	paymentId, _ := PayMentId.Get(id)
	var payment models.PayMentEntity
	p.db.Table("PayMent").Select("*").Where("id=?", paymentId.Id).Find(&payment)
	return fromDbPayMentQuery(&payment), nil

}
 
func (p *PayMentRepository) FindByOrderId(id string) (queries.ResultPayMentQuery, error) {

	orderId, _ := OrderId.Get(id)
	var payment models.PayMentEntity
	p.db.Table("PayMent").Select("*").Where("OrderId=?", orderId.Id).Find(&payment)
	return fromDbPayMentQuery(&payment), nil

}

func (p *PayMentRepository) FindByUserId(id string) (queries.ResultPayMentQuery, error) {

	userId, _ := UserId.Get(id)
	var payment models.PayMentEntity
	p.db.Table("PayMent").Select("*").Where("UserId=?", userId.Id).Find(&payment)
	return fromDbPayMentQuery(&payment), nil

}
 
