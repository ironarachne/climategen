package main

import (
	"math/rand"

	"github.com/ironarachne/climategen"
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.Writef("climategend")
	})

	app.Get("/{id:int64}", func(ctx iris.Context) {
		id, err := ctx.Params().GetInt64("id")
		if err != nil {
			ctx.Writef("error while trying to parse id parameter")
			ctx.StatusCode(iris.StatusBadRequest)
			return
		}

		rand.Seed(id)
		climate := climategen.Generate()

		ctx.JSON(climate)
	})

	app.Run(iris.Addr(":7515"))
}
