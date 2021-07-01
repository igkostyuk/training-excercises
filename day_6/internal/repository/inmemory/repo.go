package inmemory

import (
	"context"
	"fmt"

	"github.com/Metalscreame/go-training/day_6/networking-handlers/entity"
	"github.com/google/uuid"
)

func NewRepository() *Repository {
	storage := make(map[string]entity.Book)
	return &Repository{
		storage: storage,
	}
}

type Repository struct {
	storage map[string]entity.Book
}

func (r *Repository) Create(_ context.Context, b entity.Book) (string, error) {
	b.ID = uuid.New().String()
	r.storage[b.ID] = b
	// time.Sleep(time.Millisecond*200) to simulate load
	return b.ID, nil
}

func (r *Repository) Update(_ context.Context, b entity.Book) error {
	if _, ok := r.storage[b.ID]; !ok {
		return fmt.Errorf("no book with %s id", b.ID)
	}

	r.storage[b.ID] = b
	return nil
}

func (r *Repository) Delete(_ context.Context, id string) error {
	if _, ok := r.storage[id]; !ok {
		return fmt.Errorf("no book with %s id", id)
	}

	delete(r.storage, id)
	return nil
}

func (r *Repository) GetByID(_ context.Context, id string) (entity.Book, error) {
	b, ok := r.storage[id]
	if !ok {
		return entity.Book{}, fmt.Errorf("there is no book with %s id", id)
	}
	return b, nil
}

func (r *Repository) GetAll(_ context.Context) (res []entity.Book, err error) {
	for _, v := range r.storage {
		res = append(res, v)
	}
	return
}
