package go_tanks

import (
  "net"
  "encoding/json"
  i "./interfaces"
  "bufio"
  "errors"
)

const (
  NON_AUTHORIZED = iota
  AUTHORIZED
  EOL = "\n"
)

type Client struct {
  Connection  net.Conn
  State       int
  Reader      *bufio.Reader
  login       string
  password    string
  TankId      int
  channel     i.MessageChan
}

func (c *Client) RemoteAddr () (net.Addr) {
  return c.Connection.RemoteAddr()
}

func NewClient(conn net.Conn) (*Client) {
  return &Client{
    Connection: conn,
    State: NON_AUTHORIZED,
    Reader: bufio.NewReader(conn),
    channel: make(i.MessageChan, 5),
  }
}

func (c *Client) Close () {
  c.Connection.Close()
}

func (c *Client) SendMessage ( m *i.Message ) error {
  jsonStr, err := json.Marshal(m)
  if( err != nil ){ return err }

  c.Connection.Write(jsonStr)
  c.Connection.Write([]byte(EOL))

  return nil
}

func (c *Client) ReadMessage () ( *i.Message, error ) {
  buffer := []byte(nil)

  for {
    part, prefix, err := c.Reader.ReadLine()
    if err != nil { return nil, errors.New("Connection read error.") }

    if len(part) == 0 { continue }

    buffer = append(buffer, part...)

    for prefix && err == nil {
      part, prefix, err = c.Reader.ReadLine()
      if err != nil { return nil, errors.New("Connection read error.") }
      buffer = append(buffer, part...)
    }

    break
  }

  message := i.Message{}
  err := json.Unmarshal(buffer, &message)
  if( err != nil) { return nil, err }

  return &message, nil
}

func (c *Client) SetAuthCredentials ( login, password string) {
  c.login = login
  c.password = password
}

func (c *Client) Login () *string {
  return &c.login
}

func (c *Client) Password () *string {
  return &c.password
}

func (c *Client) SetTankId (id int) {
  c.TankId = id
}

func ( c *Client ) Channel () i.MessageChan {
  return c.channel
}