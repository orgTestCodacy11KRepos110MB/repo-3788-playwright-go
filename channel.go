package playwright

import (
	"fmt"
	"reflect"
)

type Channel struct {
	EventEmitter
	guid       string
	connection *Connection
	object     interface{}
}

func (c *Channel) Send(method string, options ...interface{}) (interface{}, error) {
	params := transformOptions(options...)
	result, err := c.connection.SendMessageToServer(c.guid, method, params)
	if err != nil {
		return nil, fmt.Errorf("could not send message to server: %v", err)
	}
	if result == nil {
		return nil, nil
	}
	if reflect.TypeOf(result).Kind() == reflect.Map {
		mapV := result.(map[string]interface{})
		for key := range mapV {
			return mapV[key], nil
		}
	}
	return result, nil
}

func newChannel(connection *Connection, guid string) *Channel {
	channel := &Channel{
		connection: connection,
		guid:       guid,
	}
	channel.initEventEmitter()
	return channel
}