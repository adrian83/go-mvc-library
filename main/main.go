package main

import (
	"fmt"
	"github.com/adrian83/go-redis-session"
	"github.com/gorilla/mux"
	"net/http"
	"reflect"
	"time"
	"web/cookie"
	"web/handler"
	myhttp "web/html"
	"web/session"
)

const (
	defSessionDuration = time.Duration(30) * time.Minute
)

var (
	id = 0
)

func Index(w http.ResponseWriter, r *http.Request, s redissession.Session) (map[string]interface{}, error) {

	var val int = 0

	valI, ok := s.Get("id")
	fmt.Printf("get val with key 'id': %v %v\n", valI, ok)

	if ok {

		switch i := valI.(type) {
		case int:
			val = i
		case string:
			val = 666
		case float64:
			val = int(i)
		default:
			fmt.Println(reflect.TypeOf(valI))
			val = 777
		}

	}

	model := map[string]interface{}{"idk": val}

	s.Add("id", val+1)

	return model, nil
}

func withSession(sessionStore redissession.SessionStore, handler session.SessionHandler) handler.ModelHandler {

	return func(w http.ResponseWriter, r *http.Request) (map[string]interface{}, error) {

		sessionID, ok := cookie.SessionID(r)

		var session redissession.Session

		if !ok {

			// create new session
			sess, err := sessionStore.NewSession(defSessionDuration)
			if err != nil {
				return nil, err
			}

			// create and save cookie with session id
			cookie.SaveSessionID(sess.ID(), w)

			session = sess

		} else {

			// get existing session
			sess, err := sessionStore.FindSession(sessionID)
			if err != nil {
				return nil, err
			}
			session = sess
		}

		// execute controller function
		model, err := handler(w, r, session)
		if err != nil {
			return nil, err
		}

		// save all session changes
		if err := sessionStore.SaveSession(session); err != nil {
			return nil, err
		}

		fmt.Printf("Model: %v\n", model)
		return model, err

	}
}

func main() {

	sessionStoreConfig := redissession.Config{
		DB:       0,
		Password: "",
		Host:     "localhost",
		Port:     6379,
		IDLength: 50,
	}

	sessionStore, err := redissession.NewSessionStore(sessionStoreConfig)
	if err != nil {
		panic(fmt.Sprintf("SessionStore cannot be created because of: %v", err))
	}

	fmt.Println(sessionStore)

	//mux := http.NewServeMux()
	mux := mux.NewRouter()
	files := http.FileServer(http.Dir("/static"))

	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.Handle("/", &myhttp.HttpHandler{
		View:        "index",
		HandlerFunc: withSession(sessionStore, Index)})

	server := &http.Server{Addr: "0.0.0.0:9090", Handler: mux}
	server.ListenAndServe()
}
