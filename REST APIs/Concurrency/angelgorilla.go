package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
)

func main() {
	req := http.Header{}
	req.Add("Authorization", "eyJhbGciOiJIUzUxMiJ9.eyJ1c2VybmFtZSI6IkEyNzEwNDEiLCJyb2xlcyI6MCwidXNlcnR5cGUiOiJVU0VSIiwiaWF0IjoxNjkwMTAwOTg3LCJleHAiOjE2OTAxODczODd9.13_vruwtAtXCAEG9HJWuDTzNUjDtRni7JUkMmQws_YiEI_L1_uMyuHEYWIcczyKgrgrl7aNPneXMirMmiSJhgA")
	req.Add("x-api-key", "6YzATTiE")
	req.Add("x-clientcode", "A271041")
	req.Add("x-feed-token", "eyJhbGciOiJIUzUxMiJ9.eyJ1c2VybmFtZSI6IkEyNzEwNDEiLCJpYXQiOjE2OTAxMDA5ODcsImV4cCI6MTY5MDE4NzM4N30.5Mp5cMvevbw0f_idcwyjlx4_mCk4hI4FUUSKMTwH47ZBBU63WLgUIWXdOfvIKiTyRzt3PSPm0JZQd2aL09mdoQ")
	conn, _, err := websocket.DefaultDialer.Dial("ws://smartapisocket.angelone.in/smart-stream", req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	message := map[string]any{
		"correlationID": "abcde12345",
		"action":        1,
		"params": map[string]any{
			"mode": 1,
			"tokenList": []map[string]any{
				{
					"exchangeType": 1,
					"tokens": []string{
						"10626",
						"5290",
					},
				},
				{
					"exchangeType": 1,
					"tokens": []string{
						"234230",
						"234235",
						"234219",
					},
				},
			},
		},
	}
	messageJson, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err)
		return
	}
	err2 := conn.WriteMessage(1, messageJson)
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	for {
		_, message2, err := conn.ReadMessage()
		if err != nil || err == io.EOF {
			log.Printf("Error reading: %s", err)
			break
		}
		messageString := binary.LittleEndian.
			fmt.Printf("recv: %s", messageString)
	}

}
