# WebRTCGame
A template for making real time online games with web browsers as clients.

The biggest challenge with making real time online browser games (AKA .io games) is that there is no way to directly send a UDP packet to a browser client.
WebRTC solves this problem by allowing you to make a secure connection with browser clients and then send packets acting exactly like UDP to them.
These packets are unordered, unreliable SCTP packets.

#How to Use This Code
This code is structured with main.go being the web server/ webrtc client and index.html in the public folder being the web client.
You will need to add your public IP or domain name of the location you're hosting the web server (main.go) at to the `New Websocket` line in index.html.
Next you can add any game code you'd like to index.html.
Send unreliable game state updates across the network to browser clients using the `dataChannel.SendText` or `dataChannel.Send` methods in main.go.


