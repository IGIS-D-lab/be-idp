package v2

import (
	"IGISBackEnd/orm"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

func GetMessage(database *redis.Client, w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	msgNum, err := strconv.Atoi(values.Get("msgNum"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-type", "application/json")
	}

	msg := getMsgBlock(database, msgNum)
	packet, _ := json.Marshal(msg)
	if len(msg) == 0 {
		e := MessageOk{"no message under requested msgNum"}
		packet, _ = json.Marshal(e)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-type", "application/json")
		w.Write(packet)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	w.Write(packet)
}

func PostMessage(database *redis.Client, w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	// message board number (msg index)
	msgNum, err := strconv.Atoi(v.Get("msgNum"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-type", "application/json")
	}
	msgBlock := getMsgBlock(database, msgNum)
	if len(msgBlock) == 0 {
		log.Println("This is the first message for the block")
	}
	msgThreadNo := len(msgBlock)

	sendPacket := Message{
		MessageNum:  msgNum,
		Username:    v.Get("username"),
		Password:    v.Get("password"),
		MessageTime: time.Now().Format(LayOut),
		Subthread:   msgThreadNo,
		Content:     v.Get("content"),
	}
	msg, err := checkPacket(sendPacket)

	e := MessageOk{msg}
	packet, _ := json.Marshal(e)
	if err != nil {
		// condition not satisfied
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-type", "application/json")
		w.Write(packet)
		return
	}
	// condition satisfied
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	w.Write(packet)
	msgIndex := fmt.Sprintf("%v:%v", MessageKey, msgNum)
	orm.JSONArrAppend[Message](database, msgIndex, "$", &sendPacket)
}

func checkPacket(m Message) (string, error) {
	cndUsername := m.Username == ""
	cndPassword := m.Password == ""
	cndContent := m.Content == ""

	switch {
	case cndUsername:
		return "enter username", errors.New("no username")
	case cndPassword:
		return "enter password", errors.New("no password")
	case cndContent:
		return "enter content. need minimum one word", errors.New("no content")
	default:
		return "Ok", nil
	}
}

func getMsgBlock(rdb *redis.Client, msgNum int) []Message {
	var m [][]Message
	msgIndex := fmt.Sprintf("%v:%v", MessageKey, msgNum)
	m, err := orm.JSONGet[[][]Message](rdb, msgIndex, "$", &m)
	if err != nil {
		fmt.Println(err)
	}
	return m[0]
}
