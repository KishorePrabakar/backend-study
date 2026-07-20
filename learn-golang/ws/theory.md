# WebSocket

## What is the difference between HTTP and WebSockets?

- HTTP:
  - Request-response protocol, stateless (forgets the client and server status after a session/ fetch request)
  - Used to fetch data for a definite number of times
- WebSockets:
  - Supports full duplex communication
  - Best used for Multiplayer games as it supports polling (listening for data indefinitely from server) after initializing the connection.
 
## Concurrency:
-  Concurrent processes play a vital role in online multiplayer games because they handle multiple asynchronous lightweight threads (GoRoutines in GoLang).

## What happens when I type a URL and press enter?
``` Client(Browser) -->  DNS (URL to IP) --> HTTP/TCP Connection Request --> Server```
``` Server --> Response --> Client```
