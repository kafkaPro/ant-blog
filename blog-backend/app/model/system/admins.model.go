package system

import "blog-backend/app/model"

type AdminHasOneAuthority struct {
	Admin     *model.Admins      `json:"admin"`
	Authority *model.Authorities `json:"authority"`
}
