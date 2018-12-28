package main

import (
	"github.com/Southclaws/sampctl/print"
	"github.com/Southclaws/sampctl/types"
)

func main() {
	pkg, err := types.PackageFromDir(".")
	if err != nil {
		print.Erro(err)
	}

}
