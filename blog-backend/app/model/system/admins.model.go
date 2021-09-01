package system

import "blog-backend/app/model"

type AdminHasOneRole struct {
	Admin *model.Admins `json:"admin"`
	Role  *model.Role   `json:"role"`
}
