package api

import (
	"github.com/jackc/pgx"
)

type initAPI struct {
	Db *pgx.ConnPool
}

type CustomerItem struct {
	UserId   string `json:"user-id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	RoleId   string `json:"role-id"`
	Status   string `json:"status"`
	Roles    string `json:"roles"`
}

type UpdateRequest struct {
	Id   string        `json:"id"`
	Item *CustomerItem `json:"item"`
}

type CustomerList struct {
	Limit   int32           `json:"limit"`
	Page    int32           `json:"page"`
	List    []*CustomerItem `json:"list"`
	OrderBy string          `json:"order-by"`
	Order   string          `json:"order"`
}

type CustomerListResponse struct {
	List []*CustomerListItem `json:"list"`
}

type CustomerListItem struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Status   string `json:"status"`
}

type RoleItem struct {
	Id          string `json:"id"`
	RolesName   string `json:"roles-name"`
	CreatedTime int64  `json:"created-time"`
}

var StatusType = map[string]int32{
	"UNKNOW_STATUS": 0,
	"ACTIVE":        1,
	"INACTIVE":      2,
}

type DeleteRequest struct {
	Id     string `json:"id"`
	UserId string `json:"user-id"`
}

type CustomerId struct {
	Id string `json:"id"`
}
