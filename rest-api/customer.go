package api

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"
)

func (c *initAPI) GetCustomerList(ctx context.Context, req *CustomerList) (*CustomerListResponse, error) {
	limit := 5
	page := 0

	if req.Limit != 0 {
		limit = int(req.Limit)
	}

	if req.Page != 0 {
		page = int(req.Page)
	}

	rows, err := c.Db.Query(`
		SELECT id, data FROM users LIMIT $1 OFFSET $2
	`, limit, page)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()
	var data []*CustomerListItem
	for rows.Next() {
		var item CustomerListItem
		var code string
		err = rows.Scan(
			&item.Id,
			&code,
		)

		if err != nil {
			log.Println(err)
			return nil, err
		}

		items := struct {
			Status   string `json:"status"`
			Username string `json:"username"`
			Email    string `json:"email"`
		}{}

		err = json.Unmarshal([]byte(code), &items)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		item.Username = items.Username
		item.Email = items.Email
		item.Status = items.Status

		data = append(data, &item)
	}

	if len(data) == 0 {
		return nil, errors.New("user-not-found")
	}

	return &CustomerListResponse{
		List: data,
	}, nil
}

func (c *initAPI) CreateUser(ctx context.Context, req *CustomerItem) (*CustomerId, error) {
	if req.UserId == "" {
		return nil, errors.New("missing-id")
	}

	roles, err := c.GetRoles(req.UserId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if roles.RolesName != "ADMIN" {
		return nil, errors.New("only-admin-can-create-users")
	}

	rolesId, err := c.CreateRoles(&RoleItem{
		RolesName:   "USER",
		CreatedTime: time.Now().Unix(),
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	userCode, err := json.Marshal(struct {
		Status   string `json:"status"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}{req.Status, req.Username, req.Email})

	if err != nil {
		log.Println(err)
		return nil, err
	}
	var id string
	err = c.Db.QueryRow(`
		INSERT INTO users (data, role_id) VALUES ($1, $2) RETURNING id
	`, userCode, rolesId).Scan(&id)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &CustomerId{
		Id: id,
	}, nil
}

func (c *initAPI) UpdateUser(ctx context.Context, req *UpdateRequest) (*CustomerId, error) {
	if req.Id == "" {
		return nil, errors.New("missing-id")
	}

	roles, err := c.GetRoles(req.Id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if roles.RolesName != "ADMIN" {
		return nil, errors.New("only-admin-can-create-users")
	}

	item := req.Item

	user, err := c.GetUserById(&CustomerId{
		Id: item.UserId,
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if item.Username == "" {
		item.Username = user.Username
	}

	if item.Email == "" {
		item.Email = user.Email
	}

	if item.Status == "" {
		item.Status = user.Status
	}

	if item.Roles != user.Roles {
		err := c.UpdateRoles(&RoleItem{
			Id:        user.RoleId,
			RolesName: item.Roles,
		})

		if err != nil {
			return nil, err
		}
	}

	userCode, err := json.Marshal(struct {
		Status   string `json:"status"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}{item.Status, item.Username, item.Email})

	_, err = c.Db.Exec(`
		UPDATE users SET data = $1 WHERE id = $2
	`, userCode, item.UserId)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &CustomerId{
		Id: user.UserId,
	}, nil
}

func (c *initAPI) DeleteUser(ctx context.Context, req *DeleteRequest) (*CustomerId, error) {
	if req.Id == "" {
		return nil, errors.New("missing-id")
	}

	roles, err := c.GetRoles(req.Id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if roles.RolesName != "ADMIN" {
		return nil, errors.New("only-admin-can-create-users")
	}

	var roleId string
	err = c.Db.QueryRow(`
		DELETE FROM users WHERE id = $1 RETURNING role_id
	`, req.UserId).Scan(&roleId)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = c.deleteRoles(roleId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &CustomerId{
		Id: req.Id,
	}, nil
}

func (c *initAPI) GetCustomerById(ctx context.Context, req *CustomerId) (*CustomerItem, error) {
	return c.GetUserById(req)
}

func (c *initAPI) GetUserById(req *CustomerId) (*CustomerItem, error) {
	if req.Id == "" {
		return nil, errors.New("missing-id")
	}
	var id, code, roleId string
	err := c.Db.QueryRow(`
		SELECT id, data, role_id FROM users WHERE id = $1
	`, req.Id).Scan(&id, &code, &roleId)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	items := struct {
		Status   string `json:"status"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}{}

	err = json.Unmarshal([]byte(code), &items)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	roles, err := c.GetRolesById(roleId)
	if err != nil {
		return nil, err
	}

	return &CustomerItem{
		UserId:   id,
		Username: items.Username,
		Email:    items.Email,
		RoleId:   roleId,
		Status:   items.Status,
		Roles:    roles.RolesName,
	}, nil
}

func (c *initAPI) GetRolesById(id string) (*RoleItem, error) {
	var rolesId, code string
	err := c.Db.QueryRow(`
		SELECT id, data FROM roles WHERE id = $1
	`, id).Scan(&rolesId, &code)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	items := struct {
		RolesName   string `json:"roles-name"`
		CreatedTime int64  `json:"created-time"`
	}{}

	err = json.Unmarshal([]byte(code), &items)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &RoleItem{
		Id:          rolesId,
		RolesName:   items.RolesName,
		CreatedTime: items.CreatedTime,
	}, nil
}

func (c *initAPI) GetRoles(id string) (*RoleItem, error) {
	resp, err := c.GetUserById(&CustomerId{
		Id: id,
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &RoleItem{
		Id:        resp.RoleId,
		RolesName: resp.Roles,
	}, nil
}

func (c *initAPI) CreateRoles(req *RoleItem) (string, error) {
	if req.RolesName == "" {
		return "", errors.New("missing-roles-name")
	}

	if req.CreatedTime == 0 {
		return "", errors.New("missing-created-time")
	}

	userRoles, err := json.Marshal(struct {
		RolesName   string `json:"roles-name"`
		CreatedTime int64  `json:"created-time"`
	}{req.RolesName, req.CreatedTime})

	var id string
	err = c.Db.QueryRow(`
		INSERT INTO roles (data) VALUES ($1) RETURNING id
	`, userRoles).Scan(&id)

	if err != nil {
		log.Println(err)
		return "", err
	}

	return id, nil
}

func (c *initAPI) deleteRoles(id string) error {
	_, err := c.Db.Exec(`
		DELETE FROM roles WHERE id = $1
	`, id)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (c *initAPI) UpdateRoles(req *RoleItem) error {

	role, err := c.GetRolesById(req.Id)
	if err != nil {
		log.Println(err)
		return err
	}

	req.CreatedTime = role.CreatedTime
	userRoles, err := json.Marshal(struct {
		RolesName   string `json:"roles-name"`
		CreatedTime int64  `json:"created-time"`
	}{req.RolesName, req.CreatedTime})

	_, err = c.Db.Exec(`
		UPDATE roles SET data = $1 WHERE id = $2
	`, userRoles, req.Id)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil

}
