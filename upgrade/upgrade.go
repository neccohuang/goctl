package upgrade

import (
	"fmt"
	"runtime"

	"github.com/urfave/cli"
	"github.com/neccohuang/goctl/rpc/execx"
)

// Upgrade gets the latest goctl by
// go install github.com/neccohuang/goctl@latest
func Upgrade(_ *cli.Context) error {
	cmd := `GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go install github.com/neccohuang/goctl@latest`
	if runtime.GOOS == "windows" {
		cmd = `set GOPROXY=https://goproxy.cn,direct && go install github.com/neccohuang/goctl@latest`
	}
	info, err := execx.Run(cmd, "")
	if err != nil {
		return err
	}

	fmt.Print(info)
	return nil
}
