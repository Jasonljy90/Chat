package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bregydoc/gtranslate"
	"github.com/nsqio/go-nsq"
	"golang.org/x/text/language"
)

var (
	mainInput int
)

type Message struct {
	Name      string
	Content   string
	Timestamp string
}

// Ask user to enter message and validate input. Returns the messages
// Added encryption to use message name and content
// func getMsg() (string, string) {
// 	var (
// 		msgName        string
// 		msgContent     string
// 		messageName    string
// 		messageContent string
// 	)
// 	fmt.Println("Please enter message name:")
// 	fmt.Scanf("%s\n", &messageName)
// 	fmt.Println("Please enter message content:")
// 	fmt.Scanf("%s\n", &messageContent)

// 	// Trim space of username Input
// 	msgName = strings.TrimSpace(messageName)
// 	msgContent = strings.TrimSpace(messageContent)

// 	if msgName == "" || msgContent == "" {
// 		log.Println("Message Name and Content cannot be empty")
// 	} else {
// 		encryptedMsgName := encryption([]byte(messageName))
// 		encryptedMsgContent := encryption([]byte(messageContent))
// 		return string(encryptedMsgName), string(encryptedMsgContent)
// 	}
// 	return "", ""
//}

// Ask user to enter message and validate input. Returns the messages
func getMsg() (string, string) {
	var (
		msgName        string
		msgContent     string
		messageName    string
		messageContent string
	)
	fmt.Println("Please enter message name:")
	fmt.Scanf("%s\n", &messageName)
	fmt.Println("Please enter message content:")
	fmt.Scanf("%s\n", &messageContent)

	// Trim space of username Input
	msgName = strings.TrimSpace(messageName)
	msgContent = strings.TrimSpace(messageContent)

	if msgName == "" || msgContent == "" {
		log.Println("Message Name and Content cannot be empty")
	} else {
		return msgName, msgContent
	}
	return "", ""
}

//go:generate mockgen -source=producer.go -destination=./mock/mock.go -package=mock
type publisher interface {
	Publish(topic string, body []byte) error
}

type Handler struct {
	producer publisher
}

func New(producer publisher) *Handler {
	return &Handler{
		producer: producer,
	}
}

func SendMsg(msgName, msgContent string) {
	//The only valid way to instantiate the Config
	config := nsq.NewConfig()
	//Creating the Producer using NSQD Address
	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Fatal(err)
	}

	//Init topic name and message
	topic := "Topic_Example"
	msg := Message{
		Name:      msgName,
		Content:   msgContent,
		Timestamp: time.Now().String(),
	}
	//Convert message as []byte
	payload, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
	}
	//Publish the Message
	err = producer.Publish(topic, payload)
	if err != nil {
		log.Println(err)
	}
}

func sendjapMsg(msgName, msgContent string) {
	//The only valid way to instantiate the Config
	config := nsq.NewConfig()
	//Creating the Producer using NSQD Address
	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Fatal(err)
	}
	//Init topic name and message
	topic := "Topic_Example"
	msg := Message{
		Name:      msgName,
		Content:   msgContent,
		Timestamp: time.Now().String(),
	}
	//Convert message as []byte
	payload, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
	}
	//Publish the Message
	err = producer.Publish(topic, payload)
	if err != nil {
		log.Println(err)
	}
	//Translate English To Japanese
	translated, err := gtranslate.TranslateWithParams(
		msgContent,
		gtranslate.TranslationParams{
			From: "en",
			To:   "ja",
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("en: %s | ja: %s \n", msgContent, translated)
}

func sendspainMsg(msgName, msgContent string) {
	//The only valid way to instantiate the Config
	config := nsq.NewConfig()
	//Creating the Producer using NSQD Address
	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Fatal(err)
	}
	//Init topic name and message
	topic := "Topic_Example"
	msg := Message{
		Name:      msgName,
		Content:   msgContent,
		Timestamp: time.Now().String(),
	}
	//Convert message as []byte
	payload, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
	}
	//Publish the Message
	err = producer.Publish(topic, payload)
	if err != nil {
		log.Println(err)
	}
	//Translate English To Spainish
	translatedText, _ := gtranslate.Translate(msgContent, language.English, language.Spanish)

	if err != nil {
		panic(err)
	}

	fmt.Printf("en: %s | spainish: %s \n", msgContent, translatedText)
}

func sendsimplifiedchineseMsg(msgName, msgContent string) {
	//The only valid way to instantiate the Config
	config := nsq.NewConfig()
	//Creating the Producer using NSQD Address
	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Fatal(err)
	}
	//Init topic name and message
	topic := "Topic_Example"
	msg := Message{
		Name:      msgName,
		Content:   msgContent,
		Timestamp: time.Now().String(),
	}
	//Convert message as []byte
	payload, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
	}
	//Publish the Message
	err = producer.Publish(topic, payload)
	if err != nil {
		log.Println(err)
	}
	//Translate English To Simplified Chinese
	translatedText, _ := gtranslate.Translate(msgContent, language.English, language.SimplifiedChinese)

	if err != nil {
		panic(err)
	}

	fmt.Printf("en: %s | simplified chinese: %s \n", msgContent, translatedText)
}

func sendgermanMsg(msgName, msgContent string) {
	//The only valid way to instantiate the Config
	config := nsq.NewConfig()
	//Creating the Producer using NSQD Address
	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Fatal(err)
	}
	//Init topic name and message
	topic := "Topic_Example"
	msg := Message{
		Name:      msgName,
		Content:   msgContent,
		Timestamp: time.Now().String(),
	}
	//Convert message as []byte
	payload, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
	}
	//Publish the Message
	err = producer.Publish(topic, payload)
	if err != nil {
		log.Println(err)
	}
	//Translate English To German
	translatedText, _ := gtranslate.Translate(msgContent, language.English, language.German)

	if err != nil {
		panic(err)
	}

	fmt.Printf("en: %s | german: %s \n", msgContent, translatedText)
}

// func main() {
// 	//The only valid way to instantiate the Config
// 	config := nsq.NewConfig()
// 	//Creating the Producer using NSQD Address
// 	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	h := New(producer)
// 	h.SendMsg("name example", "content example")
// for {
// 	fmt.Println("Chat Application")
// 	fmt.Println(strings.Repeat("=", 16))
// 	fmt.Println("1. Send Message")
// 	fmt.Println("2. Translate English To Japanese")
// 	fmt.Println("3. Translate English To Spanish")
// 	fmt.Println("4. Translate English To Simplified Chinese")
// 	fmt.Println("5. Translate English To German")
// 	// fmt.Println("6. Print Current Data.")
// 	// fmt.Println("7. Add New Category Name")
// 	// fmt.Println("8. Modify Category")
// 	// fmt.Println("9. Delete Category")
// 	// fmt.Println("10. Save Shopping List")
// 	// fmt.Println("11. Previous Shopping List")
// 	fmt.Println("Select your choice:")
// 	fmt.Scanf("%d\n", &mainInput)

// 	switch mainInput {
// 	case 1:
// 		//msgName, msgContent := getMsg()
// 		//sendMsg(msgName, msgContent)
// 	case 2:
// 		msgName, msgContent := getMsg()
// 		sendjapMsg(msgName, msgContent)
// 	case 3:
// 		msgName, msgContent := getMsg()
// 		sendspainMsg(msgName, msgContent)
// 	case 4:
// 		msgName, msgContent := getMsg()
// 		sendsimplifiedchineseMsg(msgName, msgContent)
// 	case 5:
// 		msgName, msgContent := getMsg()
// 		sendgermanMsg(msgName, msgContent)
// 	//case 6:

// 	//case 7:

// 	//case 8:

// 	//case 9:

// 	//case 10:

// 	//case 11:

// 	default:
// 		fmt.Println("Invalid Input")
// 	}
// }
//}

/*
func main() {
	//The only valid way to instantiate the Config
	config := nsq.NewConfig()
	//Creating the Producer using NSQD Address
	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Fatal(err)
	}
	//Init topic name and message
	topic := "Topic_Example"
	msg := Message{
		Name:      "Message Name Example",
		Content:   "Message Content Example",
		Timestamp: time.Now().String(),
	}
	//Convert message as []byte
	payload, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
	}
	//Publish the Message
	err = producer.Publish(topic, payload)
	if err != nil {
		log.Println(err)
	}
}
*/

/*
func (r *RoomMap) CreateRoom() string {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, 8)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	roomID := string(b)
	r.Map[roomID] = []Participant{}

	return roomID
}
Instead of formulating your own way to generate unique id, use uuid instead, it's easier and more proven to be unique across the industry
*/
