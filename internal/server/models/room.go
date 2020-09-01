package models

import "sync/atomic"

const roomLimit = 16

var roomCounter uint32

// Array contains every room by value
var roomSlice []room

// Map is used to quickly look up any room through it's name. Returns a pointer to the room in roomSlice
var roomMap map[string]*room

type roomID uint8

type room struct {
	name        string
	users       []userID
	owner       userID
	ID          roomID
	chatHistory messageManager
}

func InitRoom() {
	roomMap = make(map[string]*room)
	roomSlice = make([]room, 5, roomLimit)
}

func NewRoom(name string, owner userID) *room {
	roomSlice = append(roomSlice, room{name: name, users: make([]userID, 5), owner: owner,
		ID: roomID(roomCounter), chatHistory: *newMessageManager()})

	roomMap[name] = &roomSlice[roomCounter]
	atomic.AddUint32(&roomCounter, 1)

	return roomMap[name]
}

func getRoom(name string) *room {
	return roomMap[name]
}
