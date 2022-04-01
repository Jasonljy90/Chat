package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/bregydoc/gtranslate"
	"golang.org/x/text/language"
)

var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Message)           // broadcast channel

// Configure the upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Define our message object
type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

func main() {
	// Create a simple file server
	fs := http.FileServer(http.Dir("../public"))
	http.Handle("/", fs)

	// Configure websocket route
	http.HandleFunc("/ws", handleConnections)

	// Start listening for incoming chat messages
	go handleMessages()

	// Start the server on localhost port 8000 and log any errors
	log.Println("http server started on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Register our new client
	clients[ws] = true

	for {
		var msg Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		// Send the newly received message to the broadcast channel
		broadcast <- msg
	}
}

// Need modify it to translate language before seding to every client
func handleMessages() {
	var msgNew = ""
	clientLanguage := "Chinese"
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		// Send it out to every client that is currently connected
		for client := range clients {
			if clientLanguage == "Japanese"{
				msgNew = engToJapMsg(msg.Message)
			}else if clientLanguage == "Chinese"{
				msgNew = engToChineseMsg(msg.Message)
			}else if clientLanguage == "German"{
				msgNew = engToGermanMsg(msg.Message)
			}else if clientLanguage == "Spanish"{
				msgNew = engToSpanishMsg(msg.Message)
			}
			msg.Message = msgNew
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

/*
func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
*/

// Translate English To Japanese
func engToJapMsg(msgContent string) string{
	translatedText, err := gtranslate.TranslateWithParams(
		msgContent,
		gtranslate.TranslationParams{
			From: "en",
			To:   "ja",
		},
	)
	if err != nil {
		panic(err)
	}
	
	return translatedText
	//fmt.Printf("en: %s | ja: %s \n", msgContent, translated)
}

// Translate English To Spainish 
func engToSpanishMsg(msgContent string) string{
	translatedText, err := gtranslate.Translate(msgContent, language.English, language.Spanish)

	if err != nil {
		panic(err)
	}

	return translatedText
	//fmt.Printf("en: %s | spainish: %s \n", msgContent, translatedText)
}

// Translate English To Chinese
func engToChineseMsg(msgContent string) string{
	translatedText, err := gtranslate.Translate(msgContent, language.English, language.SimplifiedChinese)

	if err != nil {
		panic(err)
	}

	return translatedText
	//fmt.Printf("en: %s | simplified chinese: %s \n", msgContent, translatedText)
}

// Translate English To German
func engToGermanMsg(msgContent string) string{
	translatedText, err := gtranslate.Translate(msgContent, language.English, language.German)

	if err != nil {
		panic(err)
	}
	
	return translatedText
	//fmt.Printf("en: %s | german: %s \n", msgContent, translatedText)
}