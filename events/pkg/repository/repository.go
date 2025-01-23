package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/ArtemShamro/Calendar/events/domain"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	pgx *pgx.Conn // сторонняя библиотека для общения с бд
}

func NewRepository(pgx *pgx.Conn) *Repository {
	return&Repository{
		pgx: pgx,
	}
}

func (r *Repository) CreateNewEvent(ctx context.Context, event domain.Event) (*domain.Event, error) {
	fmt.Println("REPO : CREATE")
    query := `
        INSERT INTO events (owner, created_at, name, description, when_occurs, when_remind) 
        VALUES (@owner, @created_at, @name, @description, @when_occurs, @when_remind) 
        RETURNING id;
    `
    fmt.Println(event)
    // Define the named arguments for the query.
    args := pgx.NamedArgs{
        "owner"	:           event.Owner,
        "created_at" :      event.CreatedAt,
        "name"  :           event.Name,
        "description"  :    event.Description,
        "when_occurs"  :    event.WhenOccurs,
        "when_remind"  :    event.WhenRemind,
    }
    
    // Execute the query with named arguments to insert the book details into the database.
    err := r.pgx.QueryRow(context.Background(), query, args).Scan(&event.Id)
    if err != nil {
        log.Println("Error Adding Event")
        return nil, err
    }

	return &event, nil
}


func (r *Repository) GetEventsList(ctx context.Context) (*[]domain.Event, error) {
	fmt.Println("REPO : GetEventsList")
    // Define the SQL query to select all books.
    query := `
        SELECT * FROM events
    `
    // Execute the query to fetch all book details from the database.
    rows, err := r.pgx.Query(context.Background(), query)
    if err != nil {
        log.Printf("Error Querying the Table")
        return nil, err
    }
    defer rows.Close()

    events, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Event])
    if err != nil {
        fmt.Printf("CollectRows error: %v", err)
    }
    
    return &events, nil
}

func (r *Repository) EditEvent(ctx context.Context, event *domain.Event) (*domain.Event, error) {
    query := `
        UPDATE events 
        SET 
            owner = $1, created_at = $2, name = $3, description = $4, when_occurs = $5, when_remind = $6 
        WHERE id = $7
        RETURNING *
    `
    
    // args := pgx.NamedArgs{
    //     "owner"	:          event.Owner,
    //     "created_at" :      event.CreatedAt,
    //     "name"  :           event.Name,
    //     "description"  :    event.Description,
    //     "when_occurs"  :    event.WhenOccurs,
    //     "when_remind"  :    event.WhenRemind,
    // }

    revent := domain.Event{}

    // err := r.pgx.QueryRow(context.Background(), query, event.Owner, args).Scan(&revent) // ADD OTHER FIELDS

    err := r.pgx.QueryRow(context.Background(), 
            query, event.Owner, event.CreatedAt, event.Name, event.Description, event.WhenOccurs, event.WhenRemind, event.Id).Scan(
            &revent.Id, &revent.Owner, &revent.CreatedAt, &revent.Name, &revent.Description, &revent.WhenOccurs, &revent.WhenRemind) // ADD OTHER FIELDS
    if err != nil {
        log.Printf("Error Querying the Table")
        return nil, err
    }
    
    return &revent, nil
}

func (r *Repository) DeleteEvent(ctx context.Context, id int) (*domain.Event, error) {
	fmt.Println("REPO : Delete")

    query := `
        DELETE FROM events
		WHERE id = $1
		RETURNING *;
    `
    event := domain.Event{}
    
    err := r.pgx.QueryRow(context.Background(), query, id).Scan(
        &event.Id, &event.Owner, &event.CreatedAt, &event.Name, &event.Description, &event.WhenOccurs, &event.WhenRemind)
    if err != nil {
        log.Printf("Error Querying the Table")
        return nil, err
    }
    
    return &event, nil
}