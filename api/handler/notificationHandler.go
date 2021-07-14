package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"gitlab.luizalabs.com/luizalabs/tio-patinhas-notificacao-api/config"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"gitlab.luizalabs.com/luizalabs/tio-patinhas-notificacao-api/interfaces"
	"gitlab.luizalabs.com/luizalabs/tio-patinhas-notificacao-api/models"
	"gitlab.luizalabs.com/luizalabs/tio-patinhas-notificacao-api/utils"
)

var (
	maxID   int
	clients []models.Client
)

//HTTPNotificationHandler ...
type HTTPNotificationHandler struct {
	NotificationUseCase interfaces.NotificationUseCase
}

//NewNotificationHTTPHandler ...
func NewNotificationHTTPHandler(e *mux.Router, us interfaces.NotificationUseCase) {
	handler := &HTTPNotificationHandler{
		NotificationUseCase: us,
	}

	s := e.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/ws", handler.Connect).Methods(http.MethodGet)
	s.HandleFunc("/notifications", handler.GetAll).Methods(http.MethodGet)
	s.HandleFunc("/notification/{id:[0-9]+}/read", handler.Read).Methods(http.MethodPut)

	go handler.SearchNotifications()
}

//Connect é a função de gerenciamento de conexões do socket
func (a *HTTPNotificationHandler) Connect(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(_ *http.Request) bool { return true },
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	reader(ws)
}

//SearchNotifications é a função que procura por novas notificações
func (a *HTTPNotificationHandler) SearchNotifications() {
	for {
		not := &models.Notification{}
		list, err := a.NotificationUseCase.GetAll(not)
		if err != nil {
			log.Fatalln("[SearchNotifications] An error occurred when trying to get notifications: ", err)
		}

		for _, item := range list {
			connections := utils.FindConn(clients, strings.ToLower(item.Usuario))

			itemJSON, err := json.Marshal(item)
			if err != nil {
				log.Fatalln("[SearchNotifications] An error occurred when trying to convert Notification to JSON: ", err)
			}

			for _, conn := range connections {
				err := conn.WriteMessage(1, itemJSON)
				if err != nil {
					log.Fatalln("[SearchNotifications] An error occurred when trying to send a notification: ", err)
				}

				item.DataEnvio = time.Now().Format("2006-01-02 15:04:05")
				a.NotificationUseCase.Update(item)
			}
		}

		time.Sleep(time.Second * 2)
	}
}

//GetAll é a função que retorna todas as notificações de acordo com o filtro
func (a *HTTPNotificationHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	var (
		c config.Config
	)

	usuario := r.FormValue("usuario")

	not := &models.Notification{Usuario: usuario}

	list, err := a.NotificationUseCase.GetAll(not)
	if err != nil {
		switch err {
		case models.ErrNotFound:
			c.RespondWithErrorNew(w, c.GetStatusCode(err), 20002)
			return
		default:
			c.RespondWithError(w, c.GetStatusCode(err), err.Error(), "")
			return
		}
	}

	c.ResponseWithJSON(w, http.StatusOK, list)
}

//Read é a função que altera a notificação para o status de visualizada
func (a *HTTPNotificationHandler) Read(w http.ResponseWriter, r *http.Request) {
	var (
		c config.Config
	)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		c.RespondWithError(w, http.StatusBadRequest, "Invalid id", "")
		return
	}

	not := &models.Notification{ID: int64(id)}

	err = a.NotificationUseCase.Read(not)
	if err != nil {
		switch err {
		case models.ErrNotFound:
			c.RespondWithErrorNew(w, c.GetStatusCode(err), 20002)
			return
		default:
			c.RespondWithError(w, c.GetStatusCode(err), err.Error(), "")
			return
		}
	}

	c.ResponseWithJSON(w, http.StatusOK, nil)
}

func reader(conn *websocket.Conn) {
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			//Disconnect error codes
			codes := []int{1000, 1001, 1002, 1003, 1005, 1006, 1010, 1009, 1008, 1007, 1011}
			if !websocket.IsCloseError(err, codes...) {
				err = fmt.Errorf("[notificationHandler] Error receiving data: %v", err)
				fmt.Println(err)
			} else {
				index := utils.Find(clients, conn)
				if index != len(clients) {
					clients = utils.Remove(clients, index)
				}
			}

			return
		}

		var message models.Message

		err = json.Unmarshal(p, &message)
		if err != nil {
			log.Println(err)
			return
		}
		if message.MessageType == 1 {
			client := models.Client{
				Connection: conn,
				User:       strings.ToLower(message.Message),
				ID:         maxID,
			}

			clients = append(clients, client)
			maxID++
		}
	}
}
