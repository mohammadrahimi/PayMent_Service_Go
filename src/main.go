package main

import (
	"encoding/json"
	"log"
	"net/http"

	useCaseCreatePayMent "github.com/mohammadrahimi/PayMent_Service_Go/src/Application/UseCase/PayMent/Commands/Create"
	commandCreatePayMent "github.com/mohammadrahimi/PayMent_Service_Go/src/Domain.Contract/Commands/PayMent/Create"
	 

	useCaseQuery "github.com/mohammadrahimi/PayMent_Service_Go/src/Application/UseCase/PayMent/Queries"

	queryByIdOrder "github.com/mohammadrahimi/PayMent_Service_Go/src/Domain.Contract/Queries/PayMent/ByIdOrder"
	queryByIdPayMent "github.com/mohammadrahimi/PayMent_Service_Go/src/Domain.Contract/Queries/PayMent/ByIdPayMent"
	queryByIdUser "github.com/mohammadrahimi/PayMent_Service_Go/src/Domain.Contract/Queries/PayMent/ByIdUser"
	queryListPayMent "github.com/mohammadrahimi/PayMent_Service_Go/src/Domain.Contract/Queries/PayMent/ListPayMent"

	cqrs "github.com/mohammadrahimi/PayMent_Service_Go/src/Framework.Core/Bus"
	dbconnection "github.com/mohammadrahimi/PayMent_Service_Go/src/Infrastructure/Persistence.Sql/DbConnection"
	payment "github.com/mohammadrahimi/PayMent_Service_Go/src/Infrastructure/Persistence.Sql/Repository/PayMent"
)

var (
	commandBus cqrs.CommandBus
	queryBus   cqrs.QueryBus
)

func init() {

	//postgres://postgres:123123@localhost:5432/paymentDB
	cn := "host=localhost user=postgres password=123123 dbname=paymentDB port=5432"
	DbConnection := dbconnection.NewPostgresConnection(cn)
	DB, err := DbConnection.DBPostgres()
	if err != nil {
		panic(" ConnectionSql is Error !  " + err.Error())
	}

	repo := payment.NewPayMentRepository(DB)

	CreatePayMentCommandHandler := useCaseCreatePayMent.NewCreatePayMentCommandHandler(repo)

	commandBus = *cqrs.NewCommandBus()
	commandBus.RegisterHandler(CreatePayMentCommandHandler, &commandCreatePayMent.CreatePayMentCommand{})

	ByOrderIdQueryHandler := useCaseQuery.NewByOrderIdQueryHandler(repo)
	ByPayMentIdQueryHandler := useCaseQuery.NewByPayMentIdQueryHandler(repo)
	ByUserIdQueryHandler := useCaseQuery.NewByUserIdQueryHandler(repo)
	ListPayMentQueryHandler := useCaseQuery.NewListPayMentQueryHandler(repo)

	queryBus = *cqrs.NewQueryBus()
	queryBus.RegisterHandler(ByOrderIdQueryHandler, &queryByIdOrder.ByIdOrderQuery{})
	queryBus.RegisterHandler(ByPayMentIdQueryHandler, &queryByIdPayMent.ByIdPayMentQuery{})
	queryBus.RegisterHandler(ByUserIdQueryHandler, &queryByIdUser.ByIdUserQuery{})
	queryBus.RegisterHandler(ListPayMentQueryHandler, &queryListPayMent.ListPayMentQuery{})

}

func main() {

	mux := ControllerApi()
	if err := http.ListenAndServe(":5005", mux); err != nil {
		log.Fatal(err)
	}
}

func ControllerApi() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {

			command := &commandCreatePayMent.CreatePayMentCommand{}

			//dtoTransaction := DtoTransaction.NewDtoTransaction(25800, "OkTransaction")
			// command := &commandCreatePayMent.CreatePayMentCommand{
			// 	OrderId:     "48e17033-08ac-4845-bcf3-b925b72fcb33",
			// 	UserId:      "191f4494-d59c-4b96-90cc-c0f21f2090ab",
			// 	Description: " This is test payment ",
			// 	Transaction: dtoTransaction,
			// }

			err := json.NewDecoder(r.Body).Decode(&command)
			if err != nil {
				http.Error(w, " err: "+ err.Error(), http.StatusBadRequest)
				return
			}

			errbus := commandBus.Send(command)
			if errbus != nil {
				http.Error(w, " errbus: "+ errbus.Error(), http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("{State:OkCreate,Status:200}"))
		}

	})

	mux.HandleFunc("/ByPayMentId/{id}", func(w http.ResponseWriter, r *http.Request) {

	  
		if r.Method == http.MethodGet {

			id := r.PathValue("id")
			query := queryByIdPayMent.NewByIdPayMentQuery(id)

			result, err := queryBus.Send(query)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			data, err := json.Marshal(result)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		}

	})

	mux.HandleFunc("/ByOrderId/{id}", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {

			id := r.PathValue("id")
			query := queryByIdOrder.NewByIdOrderQuery(id)

			result, err := queryBus.Send(query)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			data, err := json.Marshal(result)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		}

	})

	mux.HandleFunc("/ByUserId/{id}", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {

			id := r.PathValue("id")
			query := queryByIdUser.NewByIdUserQuery(id)

			result, err := queryBus.Send(query)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			data, err := json.Marshal(result)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		}

	})

	mux.HandleFunc("/ByAll", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {

			query := queryListPayMent.NewListPayMentQuery()

			result, err := queryBus.Send(query)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			data, err := json.Marshal(result)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(data)

		}

	})

	return mux
}
