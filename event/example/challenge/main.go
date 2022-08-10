package main

import (
	"net/http"

	"github.com/go-zoox/core-utils/fmt"
	"github.com/go-zoox/feishu/event"
	"github.com/go-zoox/zoox"
	zd "github.com/go-zoox/zoox/default"
)

func main() {
	app := zd.Default()

	app.Any("/v1/feishu/challenge", func(ctx *zoox.Context) {
		var req event.ChallengeRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}

		fmt.PrintJSON("req", req)

		ctx.JSON(http.StatusOK, map[string]string{
			"challenge": req.Challenge,
		})
	})

	app.Run(":9999")
}
