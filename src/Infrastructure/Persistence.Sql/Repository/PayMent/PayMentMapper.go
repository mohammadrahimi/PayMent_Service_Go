package payment

import (
	"log"
	"time"

	queries "github.com/mohammadrahimi/PayMent_Service_Go/src/Domain.Contract/Queries/PayMent/ResultPayMent"
	aggregate "github.com/mohammadrahimi/PayMent_Service_Go/src/Domain/PayMent"
	"github.com/mohammadrahimi/PayMent_Service_Go/src/Domain/PayMent/ValueObject/OrderId"
	"github.com/mohammadrahimi/PayMent_Service_Go/src/Domain/PayMent/ValueObject/PayMentId"
	"github.com/mohammadrahimi/PayMent_Service_Go/src/Domain/PayMent/ValueObject/Transaction"
	"github.com/mohammadrahimi/PayMent_Service_Go/src/Domain/PayMent/ValueObject/UserId"
	models "github.com/mohammadrahimi/PayMent_Service_Go/src/Infrastructure/Persistence.Sql/Models"
)

 

func toDBPayMent(PayMent *models.PayMentEntity) *aggregate.PayMent {

	orderId,err := OrderId.Get(PayMent.OrderId.String())
    if(err != nil){
		log.Fatal("OrderId is false")
    }
    userId,err := UserId.Get(PayMent.UserId.String())
    if(err != nil){
		log.Fatal("UserId is false")
    }
    transaction,err := Transaction.New(PayMent.TransactionCode,PayMent.TansactionState)
    if(err != nil){
		log.Fatal("Transaction is false")
    }
	
 
	var p = &aggregate.PayMent{
		Id: PayMentId.New(),
		CreatedAt: time.Now(),
		OrderId: orderId,
		UserId: userId,
		Description: PayMent.Description,
		Transaction: transaction,
	}

	return p
}

func fromDBPayMent(payment *aggregate.PayMent) *models.PayMentEntity {

	var p = &models.PayMentEntity{
		 Id: payment.Id.Id,
		 Description: payment.Description,
		 CreatedAt: payment.CreatedAt,
		 OrderId: payment.OrderId.Id,
		 UserId: payment.UserId.Id,
		 TransactionCode: payment.Transaction.TransactionCode,
		 TansactionState: payment.Transaction.TansactionState,
	}
	 
	return p
}

func fromDbPayMentQuery(payment *models.PayMentEntity)  queries.ResultPayMentQuery{
	var p =  queries.ResultPayMentQuery{
		 CreatedAt: payment.CreatedAt,
		 OrderId: payment.OrderId.String(),
		 UserId: payment.UserId.String(),
		 Description: payment.Description,
		 TransactionCode: int32(payment.TransactionCode),
		 TansactionState: payment.TansactionState,
		 PayMentId: payment.Id.String(),
	}
	 
	return p
}