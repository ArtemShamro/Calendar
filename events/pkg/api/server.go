package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ArtemShamro/Calendar/db"
	"github.com/ArtemShamro/Calendar/events/app"
	"github.com/ArtemShamro/Calendar/events/domain"
	"github.com/ArtemShamro/Calendar/events/pkg/repository"
)

type EventServer struct {
	service *app.Service
}

func NewEventServer() *EventServer {
	return &EventServer{
	}
}

func (server *EventServer) Init(ctx context.Context) error {  // ADD CONFIG
	database_url := os.Getenv("DATABASE_URL")
	
	//инициализация grpc, http, роутинг, адаптеров, репозиториев, кафка, коннекторов к другим микросервисам,
	conn, err := db.NewPostgres(database_url)
	// conn, err := db.NewPostgres("postgres", "12345", database_url, "13500", "postgres")
    if err != nil {
        log.Fatalf("Error : %v", err)
    }
	repo := repository.NewRepository(conn)
	server.service = app.NewService(repo)

	return nil
}

func (server *EventServer) GetEventsList (w http.ResponseWriter, _ *http.Request) {
	fmt.Println("SERVER : GetEventsList")
	m, err := server.service.GetEventsList(context.Background())
	if err != nil {
		fmt.Printf("ERROR")
	}
	fmt.Println(m)
	err = json.NewEncoder(w).Encode(m)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func (server *EventServer) CreateNewEvent (w http.ResponseWriter, r *http.Request) {
	fmt.Println("SERVER : CREATE")
	var m CreateNewEventJSONRequestBody
	
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, "Bad argument", http.StatusBadRequest)
	}
	fmt.Println(m)
	
	h := domain.Event{
		Owner :			*m.Owner,
		CreatedAt :		m.CreatedAt.Time,
		Name : 			*m.Name,
		Description : 	*m.Description,
		WhenOccurs :	m.WhenOccurs.Time,
		WhenRemind :	m.WhenRemind.Time,
	}
	res, err := server.service.CreateNewEvent(context.Background(), h)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}


func (server *EventServer) EditEvent (w http.ResponseWriter, r *http.Request) {
	fmt.Println("SERVER : EDIT")

	var m EditEventJSONRequestBody
	
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, "Bad argument", http.StatusBadRequest)
	}
	
	h := domain.Event{
		Id :			*m.Id,
		Owner :			*m.Owner,
		CreatedAt :		m.CreatedAt.Time,
		Name : 			*m.Name,
		Description : 	*m.Description,
		WhenOccurs :	m.WhenOccurs.Time,
		WhenRemind :	m.WhenRemind.Time,
	}
	
	res, err := server.service.EditEvent(context.Background(), &h)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func (server *EventServer) DeleteEvent (w http.ResponseWriter, r *http.Request) {
	fmt.Println("SERVER : DELETE")

	var m DeleteEventJSONRequestBody
	
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, "Bad argument", http.StatusBadRequest)
	}
	
	event, err := server.service.DeleteEvent(context.Background(), *m.Id)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(event)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}