// Code generated by protoc-gen-saber-asynq. DO NOT EDIT.
// versions:
//   - protoc-gen-saber-asynq v0.1.0
//   - protoc            v3.21.2
// source: asynq.proto

package asynq

import (
	context "context"
	errors "errors"
	asynq "github.com/hibiken/asynq"
	asynq_auxiliary "github.com/things-go/protogen-saber/core/asynq_auxiliary"
	proto "google.golang.org/protobuf/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible.
var _ = errors.New
var _ = context.TODO
var _ = asynq.NewServeMux
var _ = new(emptypb.Empty)
var _ = proto.Reset

const Pattern_User_CreateUser = "user:create"
const Pattern_User_UpdateUser = "user:update"
const CronSpec_User_UpdateUser = "@every 120s"

type UserTaskHandler interface {
	// CreateUser 异步创建用户
	CreateUser(context.Context, *CreateUserPayload) error
	// UpdateUser 异步更新用户
	UpdateUser(context.Context, *UpdateUserPayload) error
}

func RegisterUserTaskHandler(mux *asynq.ServeMux, srv UserTaskHandler, opts ...asynq_auxiliary.HandlerOption) {
	settings := asynq_auxiliary.NewHandlerSettings()
	for _, opt := range opts {
		opt(settings)
	}
	mux.HandleFunc(Pattern_User_CreateUser, _User_CreateUser_Task_Handler(srv, settings))
	mux.HandleFunc(Pattern_User_UpdateUser, _User_UpdateUser_Task_Handler(srv, settings))
}

func _User_CreateUser_Task_Handler(srv UserTaskHandler, settings *asynq_auxiliary.HandlerSettings) func(context.Context, *asynq.Task) error {
	return func(ctx context.Context, task *asynq.Task) error {
		var in CreateUserPayload

		if err := settings.UnmarshalBinary(task.Payload(), &in); err != nil {
			return err
		}
		return srv.CreateUser(ctx, &in)
	}
}

func _User_UpdateUser_Task_Handler(srv UserTaskHandler, settings *asynq_auxiliary.HandlerSettings) func(context.Context, *asynq.Task) error {
	return func(ctx context.Context, task *asynq.Task) error {
		var in UpdateUserPayload

		if err := settings.UnmarshalBinary(task.Payload(), &in); err != nil {
			return err
		}
		return srv.UpdateUser(ctx, &in)
	}
}

type UserTaskClient interface {
	// CreateUser 异步创建用户
	CreateUser(context.Context, *CreateUserPayload, ...asynq.Option) (*asynq.TaskInfo, error)
	// UpdateUser 异步更新用户
	UpdateUser(context.Context, *UpdateUserPayload, ...asynq.Option) (*asynq.TaskInfo, error)
}

type UserTaskClientImpl struct {
	cc       *asynq.Client
	settings *asynq_auxiliary.ClientSettings
}

// NewUserTaskClient new client. use default proto.Marhsal.
func NewUserTaskClient(client *asynq.Client, opts ...asynq_auxiliary.ClientOption) UserTaskClient {
	settings := asynq_auxiliary.NewClientSettings()
	for _, opt := range opts {
		opt(settings)
	}
	return &UserTaskClientImpl{
		cc:       client,
		settings: settings,
	}
}

// CreateUser 异步创建用户
func (c *UserTaskClientImpl) CreateUser(ctx context.Context, in *CreateUserPayload, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	payload, err := c.settings.MarshalBinary(in)
	if err != nil {
		return nil, err
	}
	task := asynq.NewTask(Pattern_User_CreateUser, payload, opts...)
	return c.cc.Enqueue(task)
}

// UpdateUser 异步更新用户
func (c *UserTaskClientImpl) UpdateUser(ctx context.Context, in *UpdateUserPayload, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	payload, err := c.settings.MarshalBinary(in)
	if err != nil {
		return nil, err
	}
	task := asynq.NewTask(Pattern_User_UpdateUser, payload, opts...)
	return c.cc.Enqueue(task)
}

func RegisterUserScheduler_UpdateUser(scheduler *asynq.Scheduler, settings *asynq_auxiliary.ClientSettings, in *UpdateUserPayload, opts ...asynq.Option) (entryId string, err error) {
	var payload []byte

	if settings.MarshalBinary != nil {
		payload, err = settings.MarshalBinary(in)
	} else {
		payload, err = proto.Marshal(in)
	}
	if err != nil {
		return "", err
	}
	return scheduler.Register(CronSpec_User_UpdateUser, asynq.NewTask(Pattern_User_UpdateUser, payload, opts...))
}
