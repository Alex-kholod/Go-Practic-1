package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"pz10/internal/core"
	"pz10/internal/http/middleware"
	"pz10/internal/platform/config"
	"pz10/internal/platform/jwt"
	"pz10/internal/repo"
)

func Build(cfg config.Config) http.Handler {
	r := chi.NewRouter()

	// DI
	userRepository := repo.NewUserMem() // храним заранее захэшированных юзеров (email, bcrypt)
	jwtv := jwt.NewHS256(cfg.JWTSecret, cfg.JWTTTL)
	svc := core.NewService(userRepository, jwtv)

	// Публичные маршруты
	r.Post("/api/v1/login", svc.LoginHandler) // выдаёт JWT по email+password

	// Защищённые маршруты
	r.Group(func(priv chi.Router) {
		priv.Use(middleware.AuthN(jwtv))                 // аутентификация JWT
		priv.Use(middleware.AuthZRoles("admin", "user")) // базовая RBAC
		priv.Get("/api/v1/me", svc.MeHandler)            // вернёт профиль из токена
	})

	// Пример только для админов
	r.Group(func(admin chi.Router) {
		admin.Use(middleware.AuthN(jwtv))
		admin.Use(middleware.AuthZRoles("admin"))
		admin.Get("/api/v1/admin/stats", svc.AdminStats)
	})

	return r
}
