package server

import (
	"RMS/handlers"
	"RMS/middlewares"
	"RMS/models"
	"RMS/utils"
	"context"
	"github.com/go-chi/chi/v5"
	"net/http"
	"time"
)

type Server struct {
	chi.Router
	server *http.Server
}

const (
	readTimeout       = 5 * time.Minute
	readHeaderTimeout = 30 * time.Second
	writeTimeout      = 5 * time.Minute
)

func SetupRoutes() *Server {
	router := chi.NewRouter()

	router.Use(middlewares.CommonMiddlewares()...)

	router.Route("/v1", func(r chi.Router) {

		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			utils.RespondJSON(w, http.StatusOK, struct {
				Status string `json:"status"`
			}{Status: "server is running"})
		})

		r.Post("/login", handlers.LoginUser)
		r.Post("/logout", handlers.LogoutUser)

		r.Route("/admin", func(admin chi.Router) {
			admin.Use(middlewares.Authenticate)
			admin.Use(middlewares.ShouldHaveRole(models.RoleAdmin))
			admin.Post("/create-sub-admin", handlers.CreateSubAdmin)
			admin.Get("/get-all-sub-admin", handlers.GetAllSubAdmins)
			admin.Post("/create-user", handlers.CreateUser)
			admin.Get("/get-all-user", handlers.GetAllUsersByAdmin)
			admin.Post("/create-restaurant", handlers.CreateRestaurant)
			admin.Get("/get-all-restaurant", handlers.GetAllRestaurantsByAdmin)
			admin.Post("/create-dish", handlers.CreateDish)
			admin.Get("/get-all-dish", handlers.GetAllDishesByAdmin)
			admin.Get("/dishes-by-restaurant", handlers.DishesByRestaurant)
		})

		r.Route("/subadmin", func(subadmin chi.Router) {
			subadmin.Use(middlewares.Authenticate)
			subadmin.Use(middlewares.ShouldHaveRole(models.RoleSubAdmin))
			subadmin.Post("/create-user", handlers.CreateUser)
			subadmin.Get("/get-all-user", handlers.GetAllUsersBySubAdmin)
			subadmin.Post("/create-restaurant", handlers.CreateRestaurant)
			subadmin.Get("/get-all-restaurant", handlers.GetAllRestaurantsBySubAdmin)
			subadmin.Post("/create-dish", handlers.CreateDish)
			subadmin.Get("/get-all-dish", handlers.GetAllDishesBySubAdmin)
			subadmin.Get("/dishes-by-restaurant", handlers.DishesByRestaurant)
		})

		r.Route("/user", func(user chi.Router) {
			user.Use(middlewares.Authenticate)
			user.Use(middlewares.ShouldHaveRole(models.RoleUser))
			user.Get("/get-all-restaurant", handlers.GetAllRestaurantsByAdmin)
			user.Get("/dishes-by-restaurant", handlers.DishesByRestaurant)
		})
	})

	return &Server{
		Router: router,
	}
}

func (svc *Server) Run(port string) error {
	svc.server = &http.Server{
		Addr:              port,
		Handler:           svc.Router,
		ReadTimeout:       readTimeout,
		ReadHeaderTimeout: readHeaderTimeout,
		WriteTimeout:      writeTimeout,
	}
	return svc.server.ListenAndServe()
}

func (svc *Server) Shutdown(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return svc.server.Shutdown(ctx)
}
