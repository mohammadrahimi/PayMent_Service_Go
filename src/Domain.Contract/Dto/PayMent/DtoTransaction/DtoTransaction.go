
package DtoTransaction
 

type DtoTransaction struct {
	 
	TransactionCode int32
	TansactionState string   

}

func NewDtoTransaction(transactionCode int32,tansactionState string) DtoTransaction{
      return DtoTransaction{
		TransactionCode: transactionCode,
		TansactionState: tansactionState,
	  }
}
