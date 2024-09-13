/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/cham/elk/internal/gen/calc"
	"github.com/cham/elk/internal/gen/http/calc/server"
	"github.com/cham/elk/internal/service"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	goahttp "goa.design/goa/v3/http"
)

var (
	cfgFile     string
	userLicense string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "elk",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		// 初始化服务

		svc := service.NewCalcService()                       // Create Service
		endpoints := calc.NewEndpoints(svc)                   // Create endpoints
		mux := goahttp.NewMuxer()                             // Create HTTP muxer
		dec := goahttp.RequestDecoder                         // Set HTTP request decoder
		enc := goahttp.ResponseEncoder                        // Set HTTP response encoder
		svr := server.New(endpoints, mux, dec, enc, nil, nil) // Create Goa HTTP server
		server.Mount(mux, svr)                                // Mount Goa server on mux
		httpsvr := &http.Server{                              // Create Go HTTP server
			Addr:    "localhost:8081", // Configure server address
			Handler: mux,              // Set request handler
		}
		if err := httpsvr.ListenAndServe(); err != nil { // Start HTTP server
			panic(err)
		}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.coderx.toml)")
	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	viper.SetDefault("license", "apache")

	// rootCmd.AddCommand(addCmd)
	// rootCmd.AddCommand(initCmd)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".coderx" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("toml")
		viper.SetConfigName(".coderx")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
