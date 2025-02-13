package cmd

import (
	"easy-cooking/config"
	"easy-cooking/database"
	"easy-cooking/internal/handler"
	"easy-cooking/internal/router"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"time"
)

var rootCmd = &cobra.Command{
	Use:   "omni",
	Short: "Omni Channel Manager",
	Run: func(cmd *cobra.Command, args []string) {

		config.LoadConfig()

		dbConn := connectDatabase()

		handler := createHandler(dbConn)

		router.InitRouter(handler)

		port := config.Config.ServerPort
		router.StartRouter(port)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func connectDatabase() *database.Database {
	dbConn, err := database.NewDatabase(config.Config.DatabaseDSN)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return dbConn
}

func createHandler(dbConn *database.Database) *handler.Handler {
	return handler.NewHandler(dbConn.GetDB(), 5*time.Second)
}
