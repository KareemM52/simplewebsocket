// Copyright 2023 Kareem Mansour. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//
// Overview
//
// The Conn type is a WebSocket connection. The Server type is
// a WebSocket server with configs that the user can config. This library
// implements as callbacks would to handle incoming connections
// and messages. Here is and example on how you can use the library:
//
//  //  logger if you want to see debug messages (optional)
//	logger := log.New(os.Stdout, "websocket: ", log.Ltime)
//	ws := websocket.CreateServer("localhost", "8000", logger)
//  // you can put user defined function to handle events
//	ws.HandleDisconnected(onDisconnected)
//  ws.HandleConnected(onConnected)
//	ws.HandlePong(onPong)

//	ws.HandleTextMsg(onMsg)
//  // start a new go routine in a listen loop
//	go ws.ListenAndServe()
//
// Call the connection SendTextMsg and SendBinMsg methods to send  messages
// as a string/bytes. Example below:
//
//  binMsg := []byte{1,1,3,4,5}
//  conn.SendBinMsg(binMsg)
//  conn.SendTextMsg("Hello!")
package websocket
