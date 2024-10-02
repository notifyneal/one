/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var webroot string
var port string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "A golang web application starter project",
	Long:  `A golang web application starter porject `,
	Run: func(cmd *cobra.Command, args []string) {
		http.Handle("/status", Get())
		http.Handle("/echo", Post())
		// Check if the webroot is valid directory if not panic
		if _, err := os.Stat(webroot); os.IsNotExist(err) {
			panic("webroot is not a valid directory")
		}
		index := http.FileServer(http.Dir(webroot))
		http.Handle("/", index)
		fmt.Println("Server Listening on port: " + port + " ...")
		srv := &http.Server{
			Handler: nil,
			Addr:    ":" + port,
		}

		if err := srv.ListenAndServe(); err != nil {
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

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.server.yaml)")
	rootCmd.PersistentFlags().StringVarP(&webroot, "webroot", "w", "./web/dist", " define the web root folder")
	rootCmd.PersistentFlags().StringVarP(&port, "port", "p", "8080", "specify port to listen on (default: 8080)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".server" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".server")
	}

	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
