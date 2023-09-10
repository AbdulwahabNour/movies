package permission

type Permission struct {
	ID   int64  `json:"id"  `
	Code string `json:"code" validate:"required,max=200"`
}
type UserPermission struct {
	UserId       int64 `json:"user_id" validate:"required"`
	PermissionId int64 `json:"permission_id" validate:"required"`
}

func HasCode(p []*Permission, code string) bool {

	for _, v := range p {
		if v.Code == code {
			return true
		}
	}

	return false
}
