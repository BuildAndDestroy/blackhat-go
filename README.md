# blackhat-go
From the book BlackHat Go, spin of their code but writing my way to easily remember how to program in Go.

# Build Environments

## Windows
```
$ env GOOS=windows GOARCH=amd64 go build

$ file blackhat-go.exe
blackhat-go.exe: PE32+ executable (console) x86-64, for MS Windows
```

## Linux
```
# static binary, no linked libraries
$ CGO_ENABLED=0 go build

$ file blackhat-go 
blackhat-go: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), statically linked, Go BuildID=P83-EZpwXiaTx8UEc6km/1fJtGaHFhIhYy31U2CdL/3ELf7H_vRHE-k5kN61Q3/n8xRjKrQJmLfPMAJ-JsY, with debug_info, not stripped

$ ldd blackhat-go 
        not a dynamic executable
```

## OSX
```
$ env GOOS=darwin GOARCH=amd64 go build

$ file blackhat-go 
blackhat-go: Mach-O 64-bit x86_64 executable
```

## Compression

* Golang binaries are beefy bois. We can try to trim them down a bit.

``` 
$ ls -alh blackhat-go.exe 
-rwxrwxr-x 1 someuser somegroup 2.9M Dec 20 12:25 blackhat-go.exe

$ upx-ucl -9 blackhat-go.exe 
                       Ultimate Packer for eXecutables
                          Copyright (C) 1996 - 2020
UPX 3.96        Markus Oberhumer, Laszlo Molnar & John Reiser   Jan 23rd 2020

        File size         Ratio      Format      Name
   --------------------   ------   -----------   -----------
   3040256 ->   1830400   60.21%    win64/pe     blackhat-go.exe               

Packed 1 file.

$ ls -alh blackhat-go.exe 
-rwxrwxr-x 1 someuser somegroup 1.8M Dec 20 12:25 blackhat-go.exe
```
