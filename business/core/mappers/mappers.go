package mappers

import (
	"com.cross-join.crossviewer.authservice/app/service/auth-api/graph/model"
	"com.cross-join.crossviewer.authservice/business/data/schema"
)

func MapUsers(us []schema.User) []*model.User {
	result := make([]*model.User, 0, len(us))

	for i, _ := range us {
		result = append(result, MapUser(us[i]))
	}

	return result
}

func MapUser(u schema.User) *model.User {
	return &model.User{
		ID:          u.Id,
		FullName:    u.FullName,
		Username:    u.Username,
		Email:       u.Email,
		Preferences: &u.Preferences,
	}
}

func MapRoles(rs []schema.Role) []*model.Role {
	result := make([]*model.Role, 0, len(rs))

	for i, _ := range rs {
		result = append(result, MapRole(rs[i]))
	}

	return result
}

func MapRole(r schema.Role) *model.Role {
	return &model.Role{
		ID:          r.Id,
		Name:        r.Name,
		Description: r.Description,
	}
}
