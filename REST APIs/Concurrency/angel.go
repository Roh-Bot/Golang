package main

//
//import (
//	"fmt"
//	"github.com/angel-one/smartapigo/websocket"
//	"time"
//)
//
//var socketClient *websocket.SocketClient
//
//// Triggered when any error is raised
//func onError(err error) {
//	fmt.Println("Error: ", err)
//}
//
//// Triggered when websocket connection is closed
//func onClose(code int, reason string) {
//	fmt.Println("Close: ", code, reason)
//}
//
//// Triggered when connection is established and ready to send and accept data
//func onConnect() {
//	fmt.Println("Connected")
//	err := socketClient.Subscribe()
//	if err != nil {
//		fmt.Println("err: ", err)
//	}
//}
//
//// Triggered when a message is received
//func onMessage(message []map[string]interface{}) {
//	fmt.Printf("Message Received :- %v\n", message)
//}
//
//// Triggered when reconnection is attempted which is enabled by default
//func onReconnect(attempt int, delay time.Duration) {
//	fmt.Printf("Reconnect attempt %d in %fs\n", attempt, delay.Seconds())
//}
//
//// Triggered when maximum number of reconnect attempt is made and the program is terminated
//func onNoReconnect(attempt int) {
//	fmt.Printf("Maximum no of reconnect attempt reached: %d\n", attempt)
//}
//
//func main() {
//	feedToken := "eyJhbGciOiJIUzUxMiJ9.eyJ1c2VybmFtZSI6IkEyNzEwNDEiLCJpYXQiOjE2OTAxMDA5ODcsImV4cCI6MTY5MDE4NzM4N30.5Mp5cMvevbw0f_idcwyjlx4_mCk4hI4FUUSKMTwH47ZBBU63WLgUIWXdOfvIKiTyRzt3PSPm0JZQd2aL09mdoQ"
//	clientCode := "A271041"
//
//	// New Websocket Client
//	socketClient = websocket.New(clientCode, feedToken, "nse_cm|17963&nse_cm|3499&nse_cm|11536&nse_cm|21808&nse_cm|317")
//
//	// Assign callbacks
//	socketClient.OnError(onError)
//	socketClient.OnClose(onClose)
//	socketClient.OnMessage(onMessage)
//	socketClient.OnConnect(onConnect)
//	socketClient.OnReconnect(onReconnect)
//	socketClient.OnNoReconnect(onNoReconnect)
//
//	// Start Consuming Data
//	socketClient.Serve()
//
//}
