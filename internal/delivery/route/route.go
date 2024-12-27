package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"riz.it/nurul-faizah/internal/delivery/controller"
	"riz.it/nurul-faizah/internal/utils"
)

type RouterConfig struct {
	App *fiber.App
}

func NewRouter(r *fiber.App, auth fiber.Handler, authController *controller.AuthController) *RouterConfig {
	// Konfigurasi logger menggunakan utils
	logConfig := utils.ConfigureLogger("./logs", "access_log.json")
	r.Use(logger.New(logConfig))

	// Middleware CORS
	r.Use(cors.New())

	// Route
	r.Post("/api/auth/signup", authController.SignUp)
	r.Post("/api/auth/signin", authController.SignIn)
	r.Post("/api/auth/refresh", authController.RefreshToken)
	// r.Delete("/api/auth/signout", auth, authController.SignOut)

	// Mengembalikan RouterConfig
	return &RouterConfig{
		App: r,
	}
}
