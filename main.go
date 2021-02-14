package main

import (
	"fmt"
	"time"
	"strconv"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/pion/webrtc/v3"
	"github.com/wawesomeNOGUI/webrtcGamerServer/internal/signal"
)

var upgrader = websocket.Upgrader{} // use default options for upgrader

func echo(w http.ResponseWriter, r *http.Request) {
        c, err := upgrader.Upgrade(w, r, nil)
        if err != nil {
                fmt.Println("Upgrade: ", err)
                return
        }
        defer c.Close()
				fmt.Println("User connected from: ", c.RemoteAddr())

				//===========WEBRTC====================================
				// Prepare the configuration
				config := webrtc.Configuration{
					ICEServers: []webrtc.ICEServer{
						{
							URLs: []string{"stun:stun.l.google.com:19302"},
						},
					},
				}

				// Create a new RTCPeerConnection
				peerConnection, err := webrtc.NewPeerConnection(config)
				if err != nil {
					panic(err)
				}

				//Setup dataChannel to act like UDP (unordered, no retransmits) with the
				//DataChannelInit struct
				var udpPls webrtc.DataChannelInit
				var retransmits uint16 = 0
				var ordered = false

				udpPls.Ordered = &ordered
				udpPls.MaxRetransmits = &retransmits


				// Create a datachannel with label 'data' and options udpPls
				dataChannel, err := peerConnection.CreateDataChannel("data", &udpPls)
				if err != nil {
					panic(err)
				}



				// Set the handler for ICE connection state
				// This will notify you when the peer has connected/disconnected
				peerConnection.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
					fmt.Printf("ICE Connection State has changed: %s\n", connectionState.String())
				})

				// Register channel opening handling
				dataChannel.OnOpen(func() {
					fmt.Printf("Data channel '%s'-'%d' open. Random messages will now be sent to any connected DataChannels\n", dataChannel.Label(), dataChannel.ID())

					var message int

					INFINITE:
						time.Sleep(50) //50 nanoseconds
						message++ //add 1 to message
						//fmt.Printf("Sending '%s'\n", message)


						// Send the message as text
						sendErr := dataChannel.SendText(strconv.Itoa(message))    //make new byte slice with message as the only field
						if sendErr != nil {
							panic(sendErr)
						}

						goto INFINITE
				})

				// Register text message handling
				dataChannel.OnMessage(func(msg webrtc.DataChannelMessage) {
					fmt.Printf("Message from DataChannel '%s': '%s'\n", dataChannel.Label(), string(msg.Data))
				})

				// Create an offer to send to the browser
				offer, err := peerConnection.CreateOffer(nil)
				if err != nil {
					panic(err)
				}

				// Create channel that is blocked until ICE Gathering is complete
				gatherComplete := webrtc.GatheringCompletePromise(peerConnection)

				// Sets the LocalDescription, and starts our UDP listeners
				err = peerConnection.SetLocalDescription(offer)
				if err != nil {
					panic(err)
				}

				// Block until ICE Gathering is complete, disabling trickle ICE
				// we do this because we only can exchange one signaling message
				// in a production application you should exchange ICE Candidates via OnICECandidate
				<-gatherComplete

				fmt.Println( *peerConnection.LocalDescription() )

				//Send the SDP with the final ICE candidate to the browser as our offer
        err = c.WriteMessage(1, []byte(signal.Encode(*peerConnection.LocalDescription()))) //write message back to browser, 1 means message in byte format?
        if err != nil {
                fmt.Println("write:", err)
        }

				//Wait for the browser to return an answer (its SDP)
        msgType, message, err2 := c.ReadMessage() //ReadMessage blocks until message received
        if err2 != nil {
                fmt.Println("read:", err)
        }

				answer := webrtc.SessionDescription{}

        signal.Decode(string(message), &answer) //set answer to the decoded SDP
        fmt.Println(answer, msgType)

				// Set the remote SessionDescription
        err = peerConnection.SetRemoteDescription(answer)
        if err != nil {
                panic(err)
        }
}
				//WEBRTC connection made!

func main() {


	fileServer := http.FileServer(http.Dir("./public"))
	http.HandleFunc("/echo", echo) //this request comes from webrtc.html
	http.Handle("/", fileServer)

	err := http.ListenAndServe(":80", nil) //Http server blocks
	if err != nil {
					panic(err)
	}
}
