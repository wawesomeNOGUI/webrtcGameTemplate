# WebRTCGame
A template for making real time online games with a server-client structure, web browsers as clients and a pion backend as a server.

The biggest challenge with making real time online browser games (AKA .io games) is that there is no way to send an arbitray UDP packet to a browser client.
WebRTC solves this problem by allowing you to make a secure connection between browser clients and a [pion](https://pion.ly/) client and then send packets acting exactly like UDP between them.
Unordered, unreliable SCTP packets will behave exactly like vanilla UDP, but with encryption. In this example, we will be using an unreliable, **ordered** SCTP packets. We will include the ordered tag in the datachannel because any packets received older than the last packet received will be dropped, allowing every client the ability to assume every message received is the most recent update from the server.

For player controls we send what the client does to the server reliably (SCTP acting like TCP), so every player action can be taken into account by the server without any dropped messages.

# How to Use This Code
This code is structured with main.go being the web server/ webrtc client and index.html in the public folder being the web client.
You will need to add your public IP or domain name of the location you're hosting the web server (main.go) at to the `New Websocket` line in index.html, [link](https://github.com/wawesomeNOGUI/webrtcGameTemplate/blob/81adbc8efa806678abe4b296b417c00eae7f6ac7/public/index.html#L31). You'll also need to add the IP of your server in main.go [here](https://github.com/wawesomeNOGUI/webrtcGameTemplate/blob/b729c0f31b376b70ee5d6554f5fb4044ba09d60e/main.go#L298). e.g `settingEngine.SetNAT1To1IPs([]string{"172.16.0.0"}, webrtc.ICECandidateTypeHost)`. (But leaving the IP line in main.go as is will default to using your link-local address.)
Next you can add any game code you'd like to index.html, or a seperate .js file.

Send no-retransmit, but ordered, game state updates across the network to browser clients using the `dataChannel.SendText` or `dataChannel.Send` methods in main.go.

To receive client actions I chose to create another datachannel with ordered messages and retransmits enabled to be able to receive reliable messages from clients.
(No one wants to have to guess if pressing right on their controller will actually make their character move :P)

Also I added an example of how to do entity interpolation client side in public/interpolationExample. This example just takes each update from the server, saves a copy of the previous update, and divides how much the players have moved by a specified number (`interpolationFrames` in index.html). Then the quotient is used to move the player square only that much each render, for the specified amount of frames, `interpolationFrames`, creating smoother movement to the actual update positions.

# How To Build
- You need to first install [golang](https://golang.org/)
- Next make sure you have at least golang 1.15 by running `go version`
- Then git clone, or download this repository and place it inside your GOPATH
  (You can find your GOPATH by executing `go env`)
- Next enter:
  - `set GO111MODULE=on` for windows
  - `export GO111MODULW=on` for linux or mac
- Finally navigate inside the downloaded repository and run `go build` which will produce the binary server file for you to run!

# Librarys Used
Thank you to:
- [Pion](https://pion.ly/)
  For the wonderful WebRTC library in Go!
- [Gorilla/Websocket](https://github.com/gorilla/websocket)
  For the wonderful websockets library in Go!

