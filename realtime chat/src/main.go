package main

import (
	"log"
	"net/http"
	"database/sql"
	"os"
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/bregydoc/gtranslate"
	"golang.org/x/text/language"
	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	UserName  string
	Password  string
	FirstName string
	LastName  string
	Language  string
}

var(
	clients = make(map[*websocket.Conn]bool) // connected clients
 	broadcast = make(chan Message)           // broadcast channel
	db    *sql.DB
	err1 error
)

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
	// Get environment variables
	getEnvVars()
	mysqlAccount := os.Getenv("MySql_Account")

	// Open database
	db, err1 = sql.Open("mysql", mysqlAccount)
	if err1 != nil {
		panic(err1.Error())
	} else {
		fmt.Println("Database opened")
	}
	defer db.Close()

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

	// get user email to fetch language from database
	//language := getLanguageOfUser(userName)
	
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

// Need modify it to translate language before sending to every client
func handleMessages() {
	var msgNew = ""
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast

		// get user language preference(now convert key in english to preferred language)
		clientLanguage := getLanguageOfUser(msg.Email) 
		
		// Send it out to every client that is currently connected
		for client := range clients {
			if clientLanguage == "japanese"{
				msgNew = engToJapMsg(msg.Message)
			}else if clientLanguage == "chinese"{
				msgNew = engToChineseMsg(msg.Message)
			}else if clientLanguage == "german"{
				msgNew = engToGermanMsg(msg.Message)
			}else if clientLanguage == "spanish"{
				msgNew = engToSpanishMsg(msg.Message)
			}else{
				msgNew = msg.Message	
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

// get the first name of user in string type
func getLanguageOfUser(userName string) string {
	results, err := db.Query("SELECT * FROM MYSTOREDBFOODPANDA.Users where Username=?", userName)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	for results.Next() {
		var person User
		err = results.Scan(&person.UserName, &person.Password, &person.FirstName, &person.LastName, &person.Language)
		if err != nil {
			fmt.Println(err)
			return ""
		} else {
			return person.Language
		}
	}
	return ""
}

func getEnvVars() {
	err := godotenv.Load("credentials.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}