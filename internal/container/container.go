package container

import (
	"log"
	"os"

	"github.com/emuthianimbithi/pos-service/internal/config"
	"github.com/emuthianimbithi/pos-service/internal/database"
	"github.com/emuthianimbithi/pos-service/internal/repository"
	"github.com/emuthianimbithi/pos-service/internal/services"
	"gorm.io/gorm"
)

// Repositories holds all repository instances
type Repositories struct {
	Mpesa        *repository.MpesaRepository
	Notification *repository.NotificationRepository
	User         *repository.UserRepository
	Product      *repository.ProductRepository
	Inventory    *repository.InventoryRepository
	Shift        *repository.ShiftRepository
	Sale         *repository.SaleRepository
	Customer     *repository.CustomerRepository
	Business     *repository.BusinessRepository
	Branch       *repository.BranchRepository
}

// Services holds all service instances
type Services struct {
	Mpesa        *services.MpesaService
	Storage      *services.StorageService
	Notification *services.NotificationService
	Alert        *services.AlertService
	Auth         *services.AuthService
	Product      *services.ProductService
	Inventory    *services.InventoryService
	Shift        *services.ShiftService
	Sale         *services.SaleService
	Customer     *services.CustomerService
	Business     *services.BusinessService
	Branch       *services.BranchService
	User         *services.UserService
	Report       *services.ReportService
}

// Container holds all application dependencies
type Container struct {
	Config *config.Config
	DB     *gorm.DB

	Repositories *Repositories
	Services     *Services
}

// NewContainer creates a new dependency container
func NewContainer(cfg *config.Config) *Container {
	container := &Container{
		Config: cfg,
	}

	container.initInfrastructure()
	container.initRepositories()
	container.initServices()

	return container
}

func (c *Container) initInfrastructure() {
	db, err := database.InitDB(c.Config)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	c.DB = db
}

func (c *Container) initRepositories() {
	c.Repositories = &Repositories{
		Mpesa:        repository.NewMpesaRepository(c.DB),
		Notification: repository.NewNotificationRepository(c.DB),
		User:         repository.NewUserRepository(c.DB),
		Product:      repository.NewProductRepository(c.DB),
		Inventory:    repository.NewInventoryRepository(c.DB),
		Shift:        repository.NewShiftRepository(c.DB),
		Sale:         repository.NewSaleRepository(c.DB),
		Customer:     repository.NewCustomerRepository(c.DB),
		Business:     repository.NewBusinessRepository(c.DB),
		Branch:       repository.NewBranchRepository(c.DB),
	}
}

func (c *Container) initServices() {
	// Initialize Storage Service
	storageType := os.Getenv("STORAGE_TYPE")
	storageService, err := services.NewStorageService(storageType)
	if err != nil {
		log.Printf("Warning: Storage service initialization failed: %v", err)
	}

	// Initialize Notification Service
	notifType := os.Getenv("NOTIFICATION_TYPE")
	notifService := services.NewNotificationService(notifType)

	c.Services = &Services{
		Mpesa:        services.NewMpesaService(c.Repositories.Mpesa),
		Storage:      storageService,
		Notification: notifService,
		Alert:        services.NewAlertService(c.Repositories.Notification),
		Auth:         services.NewAuthService(c.Repositories.User, c.Config),
		Product:      services.NewProductService(c.Repositories.Product),
		Inventory:    services.NewInventoryService(c.Repositories.Inventory, c.Repositories.Product),
		Shift:        services.NewShiftService(c.Repositories.Shift, c.Repositories.Sale),
		Sale:         services.NewSaleService(c.Repositories.Sale, c.Repositories.Product, c.Repositories.Inventory, c.Repositories.Shift),
		Customer:     services.NewCustomerService(c.Repositories.Customer),
		Business:     services.NewBusinessService(c.Repositories.Business),
		Branch:       services.NewBranchService(c.Repositories.Branch),
		User:         services.NewUserService(c.Repositories.User),
		Report:       services.NewReportService(c.DB),
	}
}
