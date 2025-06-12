package grpc

import (
	"context"

	userpb "github.com/Dmitriihub/project-protos/proto/user"
	"github.com/Dmitriihub/users-service/internal/user"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	svc *user.Service
	userpb.UnimplementedUserServiceServer
}

func NewHandler(svc *user.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	u := &user.User{
		Email:    req.Email,
		Password: req.Password,
	}
	if err := h.svc.CreateUser(u); err != nil {
		return nil, err
	}
	return &userpb.CreateUserResponse{
		User: &userpb.User{
			Id:    uint32(u.ID),
			Email: u.Email,
		},
	}, nil
}

func (h *Handler) GetUser(ctx context.Context, req *userpb.User) (*userpb.User, error) {
	u, err := h.svc.GetUserByID(uint(req.Id))
	if err != nil {
		return nil, err
	}
	return &userpb.User{
		Id:    uint32(u.ID),
		Email: u.Email,
	}, nil
}

func (h *Handler) ListUsers(ctx context.Context, _ *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	users, err := h.svc.GetAllUsers()
	if err != nil {
		return nil, err
	}
	var pbUsers []*userpb.User
	for _, u := range users {
		pbUsers = append(pbUsers, &userpb.User{
			Id:    uint32(u.ID),
			Email: u.Email,
		})
	}
	return &userpb.ListUsersResponse{Users: pbUsers}, nil
}

func (h *Handler) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	u := &user.User{
		Email:    req.Email,
		Password: req.Password,
	}
	updated, err := h.svc.UpdateUserByID(uint(req.Id), u)
	if err != nil {
		return nil, err
	}
	return &userpb.UpdateUserResponse{
		User: &userpb.User{
			Id:    uint32(updated.ID),
			Email: updated.Email,
		},
	}, nil
}

func (h *Handler) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*emptypb.Empty, error) {
	if err := h.svc.DeleteUserByID(uint(req.Id)); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
