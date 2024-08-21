package aggregate

import (
	"time"

	commands "github.com/mohammadrahimi/PayMent_Service_Go/src/Domain.Contract/Commands/PayMent/Create"
	"github.com/mohammadrahimi/PayMent_Service_Go/src/Domain/PayMent/ValueObject/OrderId"
	"github.com/mohammadrahimi/PayMent_Service_Go/src/Domain/PayMent/ValueObject/PayMentId"
	"github.com/mohammadrahimi/PayMent_Service_Go/src/Domain/PayMent/ValueObject/Transaction"
	"github.com/mohammadrahimi/PayMent_Service_Go/src/Domain/PayMent/ValueObject/UserId"
)

 

type PayMent struct{
    Id             PayMentId.PayMentId
    CreatedAt      time.Time
    OrderId        OrderId.OrderId
    UserId         UserId.UserId
    Description    string
    Transaction    Transaction.Transaction
}



func NewPayMent(createPayMentCommand *commands.CreatePayMentCommand) (*PayMent,error){

    orderId,err := OrderId.Get(createPayMentCommand.OrderId)
    if(err != nil){
       return &PayMent{},   err
    }
    userId,err := UserId.Get(createPayMentCommand.UserId)
    if(err != nil){
       return &PayMent{},   err
    }
    transaction,err := Transaction.New(createPayMentCommand.Transaction.TransactionCode,createPayMentCommand.Transaction.TansactionState)
    if(err != nil){
       return &PayMent{},   err
    }
   

     return &PayMent{
         Id: PayMentId.New(),
         CreatedAt: time.Now(),
         OrderId: orderId,
         UserId: userId,
         Description: createPayMentCommand.Description,
         Transaction: transaction,

     }, nil

}

func (p *PayMent) GetOrderId() OrderId.OrderId{
     return p.OrderId
}
func (p *PayMent) GetUserId() UserId.UserId{
    return p.UserId
}
func (p *PayMent) GetTransaction() Transaction.Transaction{
    return p.Transaction
}
func (p *PayMent) GetDescription() string{
    return p.Description
}
func (p *PayMent) GetCreatedAt()  time.Time{
    return p.CreatedAt
}
