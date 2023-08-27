package main

import (
	"go.uber.org/fx"
	tcppow "tcp-pow"
)

func run(lc fx.Lifecycle) error {
	closer, err := tcppow.TcpServerStart(tcppow.HasherProvider, "localhost:8888")
	if err != nil {
		return err
	}

	lc.Append(fx.StopHook(closer))

	return nil
}

func main() {
	fx.New(fx.Invoke(run)).Run()
}
