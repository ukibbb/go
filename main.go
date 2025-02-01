package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/mail"
	"time"

	"github.com/google/uuid"
	"golang.org/x/net/websocket"
)

const minUserNameLength int = 3

type Server struct {
	listenAddr string
}

func (s *Server) HandleWebSocket(ws *websocket.Conn) {
	log.Printf("New incoming websocket connection %s", ws.RemoteAddr())

	// remember to add mutex
	// s.ws[conn] = ws.RemoteAddr()
	s.readFromWs(ws)
}
func (s *Server) readFromWs(ws *websocket.Conn) {
	buff := make([]byte, 1024)
	for {
		n, err := ws.Read(buff)
		if err != nil {
			if err == io.EOF {
				// means conn on other side close itself
				break
			}
			log.Println("read error: ", err)
			continue
		}
		msg := buff[:n]
		log.Println("msg received: ", string(msg))
		ws.Write([]byte("Thank you for the message!"))

	}
}

func NewServer() *Server {
	return &Server{}
}

type User struct {
	Id        uuid.UUID
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	IsActive  bool
}

type Entity = User

type Database[T Entity] interface {
	Get(id uuid.UUID) (*T, error)
	Create(*T) (*T, error)
	Update(*T) (*T, error)
	Delete(id uuid.UUID) (*T, error)
}

type InMemoryDatabase[T Entity] struct {
	storage map[uuid.UUID]*T
}

func NewMemoryDatabase[T Entity]() *InMemoryDatabase[T] {
	return &InMemoryDatabase[T]{
		storage: make(map[uuid.UUID]*T),
	}
}

func (db *InMemoryDatabase[Entity]) Get(id uuid.UUID) (e *Entity, err error) {
	e, ok := db.storage[id]
	if !ok {
		return e, fmt.Errorf("error: entity %s not found in database", id)
	}
	return e, err

}
func (db *InMemoryDatabase[Entity]) Create(e *Entity) (*Entity, error) {
	uuid := uuid.New()
	db.storage[uuid] = e
	return e, nil
}

func (db *InMemoryDatabase[Entity]) Update(e *Entity) (*Entity, error) {
	return e, nil
}
func (db *InMemoryDatabase[Entity]) Delete(id uuid.UUID) (e *Entity, err error) {
	return e, err
}

type Repository[T Entity] struct {
	db Database[T]
}

func (r *Repository[Entity]) Get() {}
func (r *Repository[Entity]) Create(e *Entity) (*Entity, error) {
	ce, err := r.db.Create(e)
	if err != nil {
		return nil, err
	}
	return ce, nil
}
func (r *Repository[Entity]) Update() {}
func (r *Repository[Entity]) Delete() {}

func NewRepository[T Entity](db Database[T]) *Repository[T] {
	return &Repository[T]{
		db: db,
	}
}

type UserRegisterRequest struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

func (r *UserRegisterRequest) validate() map[string]string {
	var errors = map[string]string{}
	// enforce some kind of password
	if r.Password != r.PasswordConfirm {
		errors["password"] = "password and password confirm doesn't match"
	}
	if len(r.Username) < minUserNameLength {
		errors["username"] = "username minimum length equals 3"
	}
	_, err := mail.ParseAddress(r.Email)
	if err != nil {
		errors["email"] = "invalid email"
	}

	// // if username is not unique
	// errors["username"] = fmt.Sprintf("%s and username is not unique", errors["username"])
	// errors["email"] = fmt.Sprintf("%s and email is not unique", errors["email"])
	// http.StatusConflict

	return errors

}

type UserLoginRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// return erros or user indetifier email or username
func (r *UserLoginRequest) validate() (rrors map[string]string) {
	var errors map[string]string
	if len(r.Username) == 0 && len(r.Email) == 0 {
		errors["user"] = "username or passowrd must be provided"
	}
	if len(r.Username) < minUserNameLength {
		errors["username"] = "username minimum length equals 3"
	}
	_, err := mail.ParseAddress(r.Email)
	if err != nil {
		errors["email"] = "invalid email"
	}
	return errors
}

type ApiError struct {
	Status int         `json:"status"`
	Msg    interface{} `json:"msg"`
}

type ApiResponse struct {
	Status int         `json:"status"`
	Msg    interface{} `json:"msg"`
}

func (e ApiError) Error() string {
	return fmt.Sprintf("error: %s status code: %d\n", e.Msg, e.Status)
}

func main() {
	s := NewServer()
	const listenAddr string = ":3000"

	router := http.NewServeMux()

	router.Handle("/ws", websocket.Handler(s.HandleWebSocket))

	user := NewUserHandlers()

	router.HandleFunc("POST /register", handler(user.handleRegister))
	router.HandleFunc("POST /login", handler(user.handleLogin))

	log.Printf("Listening on %s\n", listenAddr)

	http.ListenAndServe(listenAddr, router)
}

type UserHandlers struct {
	r *Repository[User]
}

func NewUserHandlers() *UserHandlers {
	db := NewMemoryDatabase[User]()
	return &UserHandlers{
		r: NewRepository[User](db),
	}
}

func (h *UserHandlers) handleLogin(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h *UserHandlers) handleRegister(w http.ResponseWriter, r *http.Request) error {
	ur := new(UserRegisterRequest)
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(ur); err != nil {
		return ApiError{
			Status: http.StatusBadRequest,
			Msg:    "Wrong register payload",
		}
	}

	if errors := ur.validate(); len(errors) > 0 {
		log.Printf("errors occured: %+v\n", errors)
		return ApiError{
			Status: http.StatusUnprocessableEntity,
			Msg:    errors,
		}
	}

	user := User{
		Id:        uuid.New(),
		Username:  ur.Username,
		Email:     ur.Email,
		Password:  ur.Password,
		CreatedAt: time.Now(),
		IsActive:  false,
	}

	e, err := h.r.Create(&user)
	if err != nil {
		return ApiError{
			Status: http.StatusInternalServerError,
			Msg:    err.Error(),
		}
	}
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(ApiResponse{
		Status: http.StatusCreated,
		Msg:    fmt.Sprintf("User %s has been successfully created activation email has been sent to %s", e.Username, e.Email),
	})
	return nil

}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func handler(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			if e, ok := err.(ApiError); ok {
				w.WriteHeader(e.Status)
				json.NewEncoder(w).Encode(e)
			}

		}
	}
}
