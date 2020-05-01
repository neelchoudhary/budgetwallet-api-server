package main

import (
	"flag"

	"github.com/neelchoudhary/budgetmanagergrpc/services/plaidfinances"
	"github.com/neelchoudhary/budgetmanagergrpc/services/userfinances"

	"github.com/neelchoudhary/budgetmanagergrpc/config"

	"github.com/neelchoudhary/budgetmanagergrpc/postgresql"
	"github.com/neelchoudhary/budgetmanagergrpc/services/auth"

	"github.com/neelchoudhary/budgetmanagergrpc/utils"
)

func main() {
	// Flags
	plaidClientID := flag.String("plaidClientID", "", "Plaid Client ID")
	plaidSecret := flag.String("plaidSecret", "", "Plaid Secret")
	plaidPublicKey := flag.String("plaidPublicKey", "", "Plaid Public Key")

	dbHost := flag.String("dbHost", "localhost", "Database Host")
	dbPort := flag.Int("dbPort", 5432, "Database Port")
	dbUser := flag.String("dbUser", "postgres", "Database User")
	dbPassword := flag.String("dbPassword", "password", "Database User Password")
	dbName := flag.String("dbName", "budget_manager", "Database Name")

	serverEnv := flag.String("serverEnv", "local", "Server Environment (local, prd)")
	serverHost := flag.String("serverHost", "localhost", "Server Host")
	serverAPIPort := flag.String("serverAPIPort", "50051", "API Server Port")
	serverWebhookPort := flag.String("serverWebhookPort", "8080", "Webhook Server Port")
	serverTLSKeyPath := flag.String("serverTLSKeyPath", "", "Server TLS Key Path")
	serverTLSCertPath := flag.String("serverTLSCertPath", "", "Server TLS Cert Path")

	jwtSecret := flag.String("jwtSecret", "", "JWT Secret")
	jwtExpiryMin := flag.Int("jwtExpiryMin", 0, "JWT Expiry Time in Minutes")

	flag.Parse()

	// Create configs from flags
	dbConfig := config.NewDBConfig(*dbHost, *dbPort, *dbUser, *dbPassword, *dbName)
	serverConfig := config.NewServerConfig(*serverEnv, *serverHost, *serverAPIPort, *serverWebhookPort, *serverTLSKeyPath, *serverTLSCertPath)
	plaidConfig := config.NewPlaidConfig(*plaidClientID, *plaidSecret, *plaidPublicKey)
	jwtManager := utils.NewJWTManager(*jwtExpiryMin, *jwtSecret)

	// Establish connection to data sources
	db := postgresql.ConnectDB(dbConfig)
	plaidClient := config.ConnectToPlaid(plaidConfig)

	// Create data repositories
	authRepo := postgresql.NewUserRepository(db)
	itemRepo := postgresql.NewFinancialItemRepository(db)
	accountRepo := postgresql.NewFinancialAccountRepository(db)
	transactionRepo := postgresql.NewFinancialTransactionRepository(db)

	// Create Services
	authService := auth.NewAuthServiceServer(&authRepo, jwtManager)
	userFinancesService := userfinances.NewUserFinancesServer(&itemRepo, &accountRepo)
	plaidFinancesService := plaidfinances.NewPlaidFinancesServer(&itemRepo, &accountRepo, &transactionRepo, plaidClient)

	// Create Server
	srv := NewServer(serverConfig, jwtManager, &authService, &userFinancesService, &plaidFinancesService)

	// Run Server
	go func() {
		err := srv.runHTTPServer()
		utils.LogIfFatalAndExit(err, "Failed to run http server: ")
	}()

	err := srv.runGRPCServer()
	utils.LogIfFatalAndExit(err, "Failed to run gRPC server: ")
}