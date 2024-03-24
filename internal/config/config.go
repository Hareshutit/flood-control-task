package config

import "time"

type ConfigFloodControl struct {
	TimeLimit        time.Duration
	MaxQuantityQuery int
}
