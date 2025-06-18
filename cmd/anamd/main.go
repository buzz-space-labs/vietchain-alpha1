package main

import (
	"github.com/cosmos/cosmos-sdk/server"
	"os"

	"github.com/anam-nw/anam/app"
	"github.com/anam-nw/anam/cmd/anamd/cmd"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
