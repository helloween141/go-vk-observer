package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jmoiron/sqlx"
	"go-vk-observer/config"
	"go-vk-observer/internal/database"
	"go-vk-observer/internal/services/telegram"
	"go-vk-observer/internal/services/vk"
	"log"
)

type Application struct {
	cfg             *config.Config
	db              *sqlx.DB
	telegramClient  telegram.Client
	vkClient        vk.Client
	telegramHandler telegram.Handler
	vkHandler       vk.Handler
}

func New(cfg *config.Config) (*Application, error) {
	db, err := database.Init(cfg.Database.DSN)
	if err != nil {
		return nil, err
	}

	telegramRepository := telegram.NewRepository(db)
	vkRepository := vk.NewRepository(db)

	telegramClient, err := telegram.NewClient(cfg.Telegram.Token)
	if err != nil {
		return nil, err
	}

	vkClient := vk.NewClient(cfg.Vk.BaseUrl, cfg.Vk.AccessToken, cfg.Vk.ApiVersion)
	vkService := vk.NewService()
	vkHandler := vk.NewHandler(*vkClient, telegramClient, telegramRepository, vkService)

	telegramService := telegram.NewService(telegramClient, vkClient, telegramRepository, vkRepository)
	telegramHandler := telegram.NewHandler(telegramClient, telegramService)

	return &Application{
		cfg:             cfg,
		db:              db,
		telegramClient:  *telegramClient,
		vkClient:        *vkClient,
		telegramHandler: *telegramHandler,
		vkHandler:       *vkHandler,
	}, nil
}

func (app *Application) Run() {
	go func() {
		app.telegramHandler.HandleCommands()
	}()

	http := fiber.New()
	http.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))
	http.Get("/send-notifications", func(c *fiber.Ctx) error {
		err := app.vkHandler.HandleNotifications()
		if err != nil {
			log.Println(err)
			return c.SendStatus(500)
		}

		return c.SendStatus(200)
	})
	http.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	err := http.Listen(":" + app.cfg.Application.Port)
	if err != nil {
		log.Fatal("Error starting server", err)
	}
}
