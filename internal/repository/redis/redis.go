package repository

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"time"

	"github.com/gomodule/redigo/redis"
)

type Union struct {
	TimeLimit     time.Time
	QuantityQuery int
}

type RedisRepository struct {
	data redis.Conn
}

func (r *RedisRepository) Check(ctx context.Context, userID int64,
	TimeLimit time.Duration, MaxQuantityQuery int) (bool, error) {
	if values, err := redis.Bytes(r.data.Do("get", userID)); values != nil && err == nil {

		out := Union{}
		dec := gob.NewDecoder(bytes.NewReader(values))
		err = dec.Decode(&out)
		if err != nil {
			return false, err
		}
		if out.TimeLimit.After(time.Now()) && out.QuantityQuery < MaxQuantityQuery {

			var buff1 bytes.Buffer
			out.QuantityQuery = out.QuantityQuery + 1
			enc := gob.NewEncoder(&buff1)
			enc.Encode(out)
			r.data.Do("set", userID, buff1.Bytes())
			return true, nil
		} else if out.TimeLimit.Before(time.Now()) {
			var buff1 bytes.Buffer
			out.TimeLimit = out.TimeLimit.Add(TimeLimit)
			out.QuantityQuery = 1
			enc := gob.NewEncoder(&buff1)
			enc.Encode(out)
			r.data.Do("set", userID, buff1.Bytes())
			return true, nil
		}
		return false, errors.New("Spam")
	} else if err != nil && err != redis.ErrNil {
		return false, err
	}

	var buff bytes.Buffer
	in := Union{
		TimeLimit:     time.Now().Add(TimeLimit),
		QuantityQuery: 1,
	}

	enc := gob.NewEncoder(&buff)
	enc.Encode(in)
	_, err := redis.Bytes(r.data.Do("set", userID, buff.Bytes()))
	if err != nil {
		return false, err
	}
	return true, nil
}
