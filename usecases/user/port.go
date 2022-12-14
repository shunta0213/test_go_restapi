package user

import (
	"context"

	"github.com/shunta0213/test_go_restapi/entities"
	"github.com/shunta0213/test_go_restapi/models"
)

/*
Implement: Interactor, User: Controller
*/
type UserInputPort interface {
	AddUser(ctx context.Context, user *entities.User) error
	GetUser(ctx context.Context, id int) error
	GetAllUser(ctx context.Context) error
	GetRangeUsers(ctx context.Context, start_id int, end_id int) error
	GetUsers(ctx context.Context, id ...int) error
}

/*
Implement: Presenter, User: Interactor
*/
type UserOutputPort interface {
	AddUser(ctx context.Context, user *models.User, err error) (*entities.User, error)
	OutPutUser(ctx context.Context, user *models.User, err error) (*entities.User, error)
	OutputUsers(ctx context.Context, users []*models.User, err error) (*entities.Users, error)
}

type UserRepository interface {
	AddUser(ctx context.Context, user *entities.User) (*models.User, error)
	GetUser(ctx context.Context, id int) (*models.User, error)
	GetAllUser(ctx context.Context) ([]*models.User, error)
	GetRangeUsers(ctx context.Context, start_id int, end_id int) ([]*models.User, error)
	GetUsers(ctx context.Context, ids ...int) ([]*models.User, error)
}
