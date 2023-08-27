package main

import (
	"go.uber.org/fx"
	"os"
	tcppow "tcp-pow"
)

func run(lc fx.Lifecycle) error {
	address, ok := os.LookupEnv("ADDRESS")
	if !ok {
		println("ADDRESS env variable is not set, defaulting to :8888")
		address = ":8888"
	}

	closer, err := tcppow.TcpServerStart(tcppow.HasherProvider, address)
	if err != nil {
		return err
	}

	lc.Append(fx.StopHook(closer))

	return nil
}

func main() {
	fx.New(fx.Invoke(run)).Run()
}
