# Golang Workshop 19.03.2016
The workshop aims to focus on the language's key features. During the event we are going to write a simple chat program consisting of a number of views representing a list of currently active users, the main chat history area and the text input box. The application will use the master-slave architecture where a master node receiving a portion of data through the HTTP from one of the nodes shares it with the rest of the connected instances via the TCP socket connection.

## Learning objectives
By the end of this workshop you will be able to:
- Configure and run the HTTP server
- Transmit binary data over the socket connection
- Encode custom structures into the byte arrays and back
- Perform the inter-routine communication using channels

## Dependencies
- github.com/jroimartin/gocui

Itoa - Error why not pass handlers, how to ?
