package PayMentId

import (
	"github.com/google/uuid"
	"github.com/mohammadrahimi/PayMent_Service_Go/src/Domain/PayMent/Errors"
)

type PayMentId struct{
    Id        uuid.UUID
}

 
func  New() PayMentId{
	return PayMentId{
		  Id: uuid.New(),
	  } 
}

 
func Get(value string) (PayMentId,error){
	if( value ==  ""){
            return PayMentId{}, Errors.ErrorPayMentId
	}
	return PayMentId{
		  Id:  uuid.MustParse(value),
	  } ,nil
}


