package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/Metalscreame/go-training/day_6/networking-handlers/customerrors"
	"github.com/Metalscreame/go-training/day_6/networking-handlers/entity"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

//go:generate mockgen -destination=repo_mock.go -package=server -build_flags=-mod=mod github.com/Metalscreame/go-training/day_6/networking-handlers/server BookRepository
type BookRepository interface {
	Create(ctx context.Context, b entity.Book) (string, error)
	Update(ctx context.Context, b entity.Book) error
	GetByID(ctx context.Context, id string) (entity.Book, error)
	GetAll(ctx context.Context) ([]entity.Book, error)
	Delete(ctx context.Context, id string) error
}

const idKey = "id"

type Server struct {
	repo   BookRepository
	logger *zap.Logger
}

func NewServer(repo BookRepository, log *zap.Logger) *Server {
	return &Server{repo: repo, logger: log}
}

// GetBooks Get all books
func (s *Server) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := s.repo.GetAll(r.Context())
	if err != nil {
		s.writeErrorResponse(http.StatusInternalServerError, err.Error(), w)
		return
	}

	s.logger.Info("ger all request succeeded")
	s.render(w, books)
}

// GetBook Get single book
func (s *Server) GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Gets params
	// Loop through books and find one with the id from the params
	id := params[idKey]

	book, err := s.repo.GetByID(r.Context(), id)
	if err == nil {
		s.render(w, book)
		return
	}

	s.logger.Error("can't get a book", zap.Error(err))

	// if you wan't you can set content type of the headers directly here
	w.Header().Set("Content-Type", "application/json")
	s.writeErrorResponse(http.StatusNotFound,
		fmt.Sprintf("can't find a book with %v id", id), w)
}

// CreateBook Add new book
func (s *Server) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book entity.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		s.writeErrorResponse(http.StatusBadRequest, "can't parse a book", w)
		return
	}

	id, err := s.repo.Create(r.Context(), book)
	if err != nil {
		s.logger.Error("can't create a book", zap.Error(err))
		s.writeErrorResponse(http.StatusInternalServerError, err.Error(), w)
		return
	}
	book.ID = id
	s.logger.Info("book has been created", zap.String(idKey, id))
	s.render(w, book)
}

// UpdateBook updates a book
func (s *Server) UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var book entity.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		s.writeErrorResponse(http.StatusBadRequest, "can't parse a book", w)
		return
	}

	book.ID = params["id"]
	if err := s.repo.Update(r.Context(), book); err != nil {
		s.logger.Error("can't update a book", zap.Error(err))
		if errors.Is(err, customerrors.NotFound) {
			s.writeErrorResponse(http.StatusNotFound,
				fmt.Sprintf("can't find a book with %v id", book.ID), w)
			return
		}

		s.writeErrorResponse(http.StatusInternalServerError, err.Error(), w)
		return
	}

	s.render(w, book)
}

// DeleteBook deletes a book from storage
func (s *Server) DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params[idKey]
	if err := s.repo.Delete(r.Context(), id); err != nil {
		s.logger.Error("can't delete", zap.Error(err))
		if errors.Is(err, customerrors.NotFound) {
			s.writeErrorResponse(http.StatusNotFound,
				fmt.Sprintf("can't find a book with %v id", id), w)
			return
		}

		s.writeErrorResponse(http.StatusInternalServerError, err.Error(), w)
		return
	}

	s.render(w, Message{"deleted"})
}

type Message struct {
	Msg string
}

func (*Server) writeErrorResponse(code int, msg string, w http.ResponseWriter) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(Message{msg})
}

func (s *Server) render(w http.ResponseWriter, data interface{}) {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "\t")
	enc.Encode(data)
}
