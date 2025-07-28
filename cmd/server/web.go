package main

import (
	"os"
	"todosAPI/routes"

	"azugo.io/azugo"
	"azugo.io/azugo/middleware"
	"azugo.io/azugo/server"
	"github.com/spf13/cobra"
)

// webCmd represents the web command
var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Start web server",
	Long: `Web server is the only thing you need to run,
and it takes care of all the other things for you`,
	RunE: runWeb,
}

func runWeb(cmd *cobra.Command, args []string) error {
	app, err := server.New(cmd, server.Options{
		AppName: "Example application",
		AppVer:  Version,
	})
	if err != nil {
		return err
	}

	app.Get("/", func(ctx *azugo.Context) {
		ctx.Text("Hello, World!")
	})

	app.Get("/gettodo/{taskID}", func(ctx *azugo.Context) {
		routes.GetInfobyID(ctx)
	})

	app.Get("/listall", func(ctx *azugo.Context) {
		routes.GetAllTodos(ctx)
	})

	app.Post("/posttodo", func(ctx *azugo.Context) {
		routes.AddATodo(ctx)
	})

	app.Delete("/delete/{taskID}", func(ctx *azugo.Context) {
		routes.DeleteToDo(ctx)
	})
	app.Patch("/edit/{taskID}", func(ctx *azugo.Context) {
		routes.EditAToDo(ctx)
	})

	corsOpts := app.RouterOptions().CORS
	corsOpts.
		SetOrigins(os.Getenv("CORS_ORIGINS")).
		SetMethods("GET", "POST", "OPTIONS", "DELETE", "PATCH", "PUT").
		SetHeaders("Content-Type", "Authorization", "Accept")

	app.Use(middleware.CORS(&corsOpts))

	server.Run(app)
	return nil
}

func init() { // to run do :go run ./cmd/server
	initRootCmd()
	RootCmd.AddCommand(webCmd)
}
