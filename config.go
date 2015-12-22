package main

type DFConfig struct {
	URLs []string

	Period *int64
}

type ConfigSettings struct {
	Input DFConfig
}
