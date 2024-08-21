package Errors

import "errors"
 
 
var (
	ErrorTransaction = errors.New("Transaction   is not correct")
	ErrorUserId = errors.New("UserId   is not correct")
	ErrorOrderId= errors.New("OrderId   is not correct")
	ErrorPayMentId= errors.New("PayMentId   is not correct")
	 
)
