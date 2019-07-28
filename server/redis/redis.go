package redis

import (
	"encoding/json"
	redigo "github.com/garyburd/redigo/redis"
)

var pool *redigo.Pool

func init() {
	pool = &redigo.Pool{
		// Other pool configuration not shown in this example.
		Dial: func () (redigo.Conn, error) {
			c, err := redigo.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", "123456"); err != nil {
				c.Close()
				return nil, err
			}
			if _, err := c.Do("SELECT", 0); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
	}
}

func Set(key string,value interface{},timeout int64) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	conn := pool.Get()
	err = conn.Send("SET",key,data)
	if err != nil {
		return err
	}
	err = conn.Send("EXPIRE",key,timeout)
	if err != nil {
		return err
	}
	err = conn.Flush()

	return err
}

func Get(key string,value interface{}) error  {
	data,err := redigo.Bytes(pool.Get().Do("GET",key))
	if err != nil {
		return err
	}
	err = json.Unmarshal(data,value)
	return err
}

func Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	return pool.Get().Do(commandName,args)
}