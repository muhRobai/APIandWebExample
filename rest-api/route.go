package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
)

//done
func (c *initAPI) GetCustomerByIdHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	vars := mux.Vars(r)
	resp, err := c.GetCustomerById(ctx, &CustomerId{
		Id: vars["id"],
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "faild-convert-json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

//done
func (c *initAPI) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var p DeleteRequest
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Println(err)
		http.Error(w, "faild-convert-json", http.StatusInternalServerError)
		return
	}

	resp, err := c.DeleteUser(ctx, &p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "faild-convert-json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

//done
func (c *initAPI) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var p UpdateRequest
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Println(err)
		http.Error(w, "faild-convert-json", http.StatusInternalServerError)
		return
	}

	resp, err := c.UpdateUser(ctx, &p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "faild-convert-json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

//done
func (c *initAPI) CreateUserhandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var p CustomerItem
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Println(err)
		http.Error(w, "faild-convert-json", http.StatusInternalServerError)
		return
	}

	resp, err := c.CreateUser(ctx, &p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "faild-convert-json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

//done
func (c *initAPI) GetCustomerListHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var p CustomerList
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Println(err)
		http.Error(w, "faild-convert-json", http.StatusInternalServerError)
		return
	}

	resp, err := c.GetCustomerList(ctx, &p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "faild-convert-json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func (c *initAPI) initDB() {

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	port, err := strconv.Atoi(dbPort)
	if err != nil {
		log.Println(err.Error())
		return
	}

	dbConfig := &pgx.ConnConfig{
		Port:     uint16(port),
		Host:     dbHost,
		User:     dbUser,
		Password: dbPass,
		Database: dbName,
	}

	connection := pgx.ConnPoolConfig{
		ConnConfig:     *dbConfig,
		MaxConnections: 5,
	}

	c.Db, err = pgx.NewConnPool(connection)
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func StartHTTP() http.Handler {
	api, err := createAPI()
	if err != nil {
		log.Println(err)
		return nil
	}

	api.initDB()

	r := mux.NewRouter()
	// get customer list
	r.HandleFunc("/api/customer/{id}", api.GetCustomerByIdHandler).Methods("GET")
	r.HandleFunc("/api/customer/delete", api.DeleteUserHandler).Methods("DELETE")
	r.HandleFunc("/api/customer/update", api.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/api/customer/create", api.CreateUserhandler).Methods("POST")
	r.HandleFunc("/api/customer/list", api.GetCustomerListHandler).Methods("GET")
	return r
}
