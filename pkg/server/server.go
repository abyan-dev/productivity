package server

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/abyan-dev/productivity/pkg/handler"
	"github.com/abyan-dev/productivity/pkg/middleware"
	"github.com/abyan-dev/productivity/pkg/model"
	"github.com/abyan-dev/productivity/pkg/utils"
	"github.com/goccy/go-json"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Server struct {
	DB *gorm.DB
}

func (s *Server) New() *fiber.App {
	slog.Info("Loading environment variables...")

	config, err := utils.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	slog.Info("Connecting to database...")
	db, err := utils.InitDB(config)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	s.DB = db

	slog.Info("Applying database migrations...")
	if err := db.AutoMigrate(&model.Task{}, &model.PomodoroSession{}); err != nil {
		log.Fatalf("Error auto-migrating database: %v", err)
	}

	slog.Info("Setting up the app...")

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	slog.Info("Loading routes...")

	s.initRouter(app)

	return app
}

func (s *Server) initRouter(app fiber.Router) {
	api := app.Group("/api")

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",                  // Allow specific origin
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS", // Allow all methods
		AllowHeaders:     "Content-Type, Authorization",            // Allow specific headers
		AllowCredentials: true,
	}))

	// Health check
	api.Get("/health", handler.Health)
	api.Get("/health/protected", middleware.RequireAuthenticated(), handler.HealthProtected)

	api.Use(middleware.RequireAuthenticated())

	// Task management
	api.Post("/productivity/tasks", handler.CreateTask)
	api.Get("/productivity/tasks", handler.GetAllTasks)
	api.Get("/productivity/tasks/:id", handler.GetTask)
	api.Put("/productivity/tasks/:id", handler.UpdateTask)
	api.Delete("/productivity/tasks/:id", handler.DeleteTask)

	// Pomodoro timer
	api.Post("/productivity/pomodoro/start", handler.StartPomodoro)
	api.Put("/productivity/pomodoro/stop", handler.StopPomodoro)

	// Study performance metrics
	api.Get("/productivity/metrics/study", handler.GenerateStudyMetrics)
}

func (s *Server) Run(app *fiber.App) {
	slog.Info("Server is now listening on port 8081...")
	if err := app.Listen(":8081"); err != nil {
		slog.Error(fmt.Sprintf("Failed to start server: %v", err))
	}
}
