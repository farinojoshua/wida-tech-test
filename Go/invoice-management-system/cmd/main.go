package main

import (
	"invoice-management-system/internal/configs"
	"invoice-management-system/internal/handlers/invoices"
	invoiceRepo "invoice-management-system/internal/repository/invoices"
	invoiceService "invoice-management-system/internal/service/invoices"
	"invoice-management-system/pkg/internalsql"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder(
			[]string{"./internal/configs"},
		),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil {
		log.Fatal("failed to init configs", err)
	}

	cfg = configs.GetConfig()
	log.Println("config: ", cfg)

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("failed to connect to database", err)
	}

	// initialize repositories
	invoiceRepo := invoiceRepo.NewRepository(db)

	// initialize services
	invoiceService := invoiceService.NewService(cfg, invoiceRepo)

	// initialize handlers
	invoiceHandler := invoices.NewHandler(r, invoiceService)
	invoiceHandler.RegisterRoutes()

	r.Run(cfg.Service.Port)
}
