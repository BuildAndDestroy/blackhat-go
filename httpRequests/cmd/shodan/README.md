# shodan

## Help Menu

```
$ ./shodan 
2024/02/06 15:41:18 Missing search term and/or API key

$ ./shodan nginx -h
Usage of nginx:
  -shodan-api-key string
        Your shodan API key. This flag as plaintext or define as a global variable.
            Example: --shodan-api-key $SHODAN_API_KEY (default "NOTDEFINED")
```

## Example shodan
```
$ ./shodan nginx --shodan-api-key $SHODAN_API_KEY 
Query Credits: 100
Scan Credits: 100
Plan: dev

      100.24.89.80     443
         52.44.0.0     443
     217.160.0.119     443
      64.225.91.73     443
       20.71.80.38      80
..
..
REDACTED
```


# Build Environments

## Windows
```
~/blackhat-go/httpRequests/cmd/shodan$ env GOOS=windows GOARCH=amd64 go build

~/blackhat-go/httpRequests/cmd/shodan$ file shodan.exe
shodan.exe: PE32+ executable (console) x86-64, for MS Windows
```

## Linux
```
# static binary, no linked libraries
~/blackhat-go/httpRequests/cmd/shodan$ CGO_ENABLED=0 go build

~/blackhat-go/httpRequests/cmd/shodan$ file shodan
shodan: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), statically linked, Go BuildID=P83-EZpwXiaTx8UEc6km/1fJtGaHFhIhYy31U2CdL/3ELf7H_vRHE-k5kN61Q3/n8xRjKrQJmLfPMAJ-JsY, with debug_info, not stripped

~/blackhat-go/httpRequests/cmd/shodan$ ldd shodan
        not a dynamic executable
```

## OSX
```
~/blackhat-go/httpRequests/cmd/shodan$ env GOOS=darwin GOARCH=amd64 go build

~/blackhat-go/httpRequests/cmd/shodan$ file shodan
shodan: Mach-O 64-bit x86_64 executable
```

## Compression

* Golang binaries are beefy bois. We can try to trim them down a bit.

``` 
~/blackhat-go/httpRequests/cmd/shodan$ ls -alh shodan.exe 
-rwxrwxr-x 1 someuser somegroup 2.9M Dec 20 12:25 shodan.exe

~/blackhat-go/httpRequests/cmd/shodan$ upx-ucl -9 shodan.exe 
                       Ultimate Packer for eXecutables
                          Copyright (C) 1996 - 2020
UPX 3.96        Markus Oberhumer, Laszlo Molnar & John Reiser   Jan 23rd 2020

        File size         Ratio      Format      Name
   --------------------   ------   -----------   -----------
   3040256 ->   1830400   60.21%    win64/pe     shodan.exe               

Packed 1 file.

~/blackhat-go/httpRequests/cmd/shodan$ ls -alh shodan.exe 
-rwxrwxr-x 1 someuser somegroup 1.8M Dec 20 12:25 shodan.exe
```
