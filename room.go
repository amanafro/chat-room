package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type Room struct {
	name    string
	clients map[*Client]bool
}

func NewRoom (name string) *Room {
  return &Room{
    name: name,
    clients: make(map[*Client type]bool),
  }
}

func RunRoom (room *Room) {

} 
