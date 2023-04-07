package ctype

import "encoding/json"

type Role int

const (
	PermissionAdmin       Role = 1 // 管理员
	PermissionUser        Role = 2 // 普通用户
	PermissionVisitor     Role = 3 // 访客
	PermissionDisableUser Role = 4 // 被禁言的用户
)

func (r Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r Role) String() string {
	var str string
	switch r {
	case PermissionAdmin:
		str = "管理员"
	case PermissionUser:
		str = "普通用户"
	case PermissionVisitor:
		str = "访客"
	case PermissionDisableUser:
		str = "禁言用户"
	default:
		str = "其他"
	}
	return str
}
