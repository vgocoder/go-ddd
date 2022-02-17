/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"flag"
	"fmt"
	"github.com/go-chi/chi/v5"
	"kingsmith.com.cn/ks-horizon/config"
	"kingsmith.com.cn/ks-horizon/infrastructure/database"
	"kingsmith.com.cn/ks-horizon/interfaces"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/docgen"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server called")
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// init db
	setupDb()

	// init router
	r := setupRouter()

	// start server
	setupServer(r)
}

var routes = flag.Bool("routes", false, "Generate router documentation")

func setupServer(r *chi.Mux) {
	appConfig := config.Conf.App
	port := "3000"
	if appConfig.Port != "" {
		port = appConfig.Port
	}
	http.ListenAndServe(fmt.Sprintf(":%v", port), r)
}

func setupRouter() *chi.Mux {
	flag.Parse()
	r := chi.NewRouter()
	h := interfaces.Handler()
	r.Mount("/", h)
	r.Mount("/debug", middleware.Profiler())

	// Print Router
	printAllRegisteredRoutes(r)

	// Passing -routes to the program will generate docs for the above
	// router definition. See the `routes.json` file in this folder for
	// the output.
	if *routes {
		// fmt.Println(docgen.JSONRoutesDoc(r))
		fmt.Println(docgen.MarkdownRoutesDoc(r, docgen.MarkdownOpts{
			ProjectPath: "github.com/go-chi/chi/v5",
			Intro:       "Welcome to the chi/_examples/rest generated docs.",
		}))
	}
	return r
}

func setupDb() {
	err := database.SetupDb()
	if err != nil {
		panic(err)
	}
}

func printAllRegisteredRoutes(r *chi.Mux) {
	walkFunc := func(method string, path string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Printf("%-7s %s\n", method, path)
		return nil
	}
	if err := chi.Walk(r, walkFunc); err != nil {
		fmt.Print(err)
	}
}
