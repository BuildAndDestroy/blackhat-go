package rpc

import (
	"bytes"
	"fmt"
	"net/http"

	"gopkg.in/vmihailenco/msgpack.v2"
)

type Metasploit struct {
	host  string
	user  string
	pass  string
	token string
}

func NewMetasploit(host, user, pass string) *Metasploit {
	msf := &Metasploit{
		host: host,
		user: user,
		pass: pass,
	}
	return msf
}

func (msf *Metasploit) Send(req interface{}, res interface{}) error {
	buf := new(bytes.Buffer)
	msgpack.NewEncoder(buf).Encode(req)
	dest := fmt.Sprintf("http://%s/api", msf.host)
	r, err := http.Post(dest, "binary/message-pack", buf)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if err := msgpack.NewDecoder(r.Body).Decode(&res); err != nil {
		return err
	}

	return nil
}

func (msf *Metasploit) Login() error {
	ctx := &LoginReq{
		Method:   "auth.login",
		Username: msf.user,
		Password: msf.pass,
	}
	var res LoginRes
	if err := msf.Send(ctx, &res); err != nil {
		return err
	}
	msf.token = res.Token
	return nil
}

func (msf *Metasploit) Logout() error {
	ctx := &LogoutReq{
		Method:      "auth.logout",
		Token:       msf.token,
		LogoutToken: msf.token,
	}
	var res LogoutRes
	if err := msf.Send(ctx, &res); err != nil {
		return err
	}
	msf.token = ""
	return nil
}

func (msf *Metasploit) SessionList() (map[uint32]SessionListRes, error) {
	req := &SessionListReq{Method: "session.list", Token: msf.token}
	res := make(map[uint32]SessionListRes)
	if err := msf.Send(req, &res); err != nil {
		return nil, err
	}

	for id, session := range res {
		session.ID = id
		res[id] = session
	}
	return res, nil
}

type SessionListReq struct {
	// Session request to Metasploit RPC
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type SessionListRes struct {
	// Session response from Metasploit RPC
	ID          uint32 `msgpack:",omitempty"`
	Type        string `msgpack:"type"`
	TunnelLocal string `msgpack:"tunnel_local"`
	TunnelPeer  string `msgpack:"tunnel_peer"`
	ViaExploit  string `msgpack:"via_exploit"`
	ViaPayload  string `msgpack:"via_payload"`
	Description string `msgpack:"desc"`
	Info        string `msgpack:"info"`
	Workspace   string `msgpack:"workspace"`
	SessionHost string `msgpack:"session_host"`
	SessionPort int    `msgpack:"session_port"`
	Username    string `msgpack:"username"`
	UUID        string `msgpack:"uuid"`
	ExploitUUID string `msgpack:"exploit_uuid"`
}

type LoginReq struct {
	// Login request to Metasploit RPC
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Username string
	Password string
}

type LoginRes struct {
	// Login Response from Metasploit RPC
	Result       string `msgpack:"result"`
	Token        string `msgpack:"token"`
	Error        bool   `msgpack:"error"`
	ErrorClass   string `msgpack:"error_class"`
	ErrorMessage string `msgpack:"error_message"`
}

type LogoutReq struct {
	// Request to Metasploit RPC
	_msgpack    struct{} `msgpack:",asArray"`
	Method      string
	Token       string
	LogoutToken string
}

type LogoutRes struct {
	// Response to Metasploit RPC when logging out
	Result string `msgpack:"result"`
}
