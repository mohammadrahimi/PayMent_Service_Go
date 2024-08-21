package OrderId

import (
	"github.com/google/uuid"
	"github.com/mohammadrahimi/PayMent_Service_Go/src/Domain/PayMent/Errors"
)

type OrderId struct{
    Id        uuid.UUID
}

 
 
func Get(value string) (OrderId,error){
	if( value ==  ""){
            return OrderId{}, Errors.ErrorOrderId
	}
	return OrderId{
		  Id:  uuid.MustParse(value),
	  } ,nil
}


