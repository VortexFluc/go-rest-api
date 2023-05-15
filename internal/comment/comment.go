package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("Failed to fetch comment by id")
	ErrNotImplemented = errors.New("Not Implemented!")
)

// Comment is a represantation of the comment sturct
// for service
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Store is an interface defines all methods, that service needs in order to operate
type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

// Service - contains logic
type Service struct {
	Store Store
}

// NewService is a "Constructor"
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("Retrieving a comment")
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
	}
	return cmt, nil
}

func (s *Service) UpdateComment(ctx context.Context, cmt Comment) error {
	return ErrNotImplemented
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return ErrNotImplemented
}

func (s *Service) CreateComment(ctx context.Context, cmt Comment) (Comment, error) {
	return Comment{}, ErrNotImplemented
}