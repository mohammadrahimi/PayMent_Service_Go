package UserId

import (
	"github.com/google/uuid"
	"github.com/mohammadrahimi/PayMent_Service_Go/src/Domain/PayMent/Errors"
)

type UserId struct{
    Id        uuid.UUID
}

 

func  New() UserId{
      return UserId{
		    Id: uuid.New(),
		} 
}
 
func Get(value string) (UserId,error){
	if( value ==  ""){
            return UserId{}, Errors.ErrorUserId
	}
	return UserId{
		  Id:  uuid.MustParse(value),
	  } ,nil
}

