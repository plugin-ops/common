package action

import (
	"context"
)

type ActionSet interface {
	Name() string

	Do(ctx context.Context, actionName string, params ...Parameter) ([]Parameter, error)
	AddAction(action Action)
	GetAction(actionName string) Action
}

type Action interface {
	Name() string
	Do(ctx context.Context, params ...Parameter) ([]Parameter, error)
}
