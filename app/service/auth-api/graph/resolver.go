package graph

import (
	"com.cross-join.crossviewer.authservice/business/core/users"
	"go.uber.org/zap"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UsersCore users.Core
	Log       *zap.SugaredLogger
}
