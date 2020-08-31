package Models

import (
	"net"
	"sync/atomic"
)

const userLimit = 16

var userSlice []user
var userMap map[string]*user
var userCounter uint32 = 0

type userID uint8

type user struct {
	name string
	ID userID
	connection net.Conn
	privateMessages chan []string
	currentRoom *room
	admin bool

}

func InitUser(){
	userMap = make(map[string]*user)
	userSlice = make([]user, 5, userLimit)
}

func NewUser(name string, con net.Conn) *user{
	userSlice = append(userSlice, user{name: name, ID: userID(userCounter), connection: con,
		privateMessages: make(chan []string), currentRoom: getRoom("main")})
	userMap[name] = &userSlice[userCounter]
	atomic.AddUint32(&userCounter, 1)
	return userMap[name]
}