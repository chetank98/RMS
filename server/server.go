package server

import (
	"RMS/handlers"
	"RMS/middlewares"
	"RMS/models"
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

	router.Route("/v1", func(v1 chi.Router) {

		v1.Post("/login", handlers.LoginUser)

		v1.Group(func(r chi.Router) {

			r.Use(middlewares.Authenticate)

			r.Get("/dishes-by-restaurant", handlers.DishesByRestaurant)
			r.Post("/logout", handlers.LogoutUser)

			r.Route("/admin", func(admin chi.Router) {
				admin.Use(middlewares.ShouldHaveRole(models.RoleAdmin))
				admin.Post("/create-sub-admin", handlers.CreateSubAdmin)
				admin.Get("/all-sub-admin", handlers.GetAllSubAdmins)
				admin.Post("/create-user", handlers.CreateUser)
				admin.Get("/all-users", handlers.GetAllUsersByAdmin)
				admin.Post("/create-restaurant", handlers.CreateRestaurant)
				admin.Get("/all-restaurants", handlers.GetAllRestaurants)

				admin.Route("/{restaurantId}", func(restaurantIDRoute chi.Router) {
					restaurantIDRoute.Post("/", handlers.CreateDish)
				})

				admin.Get("/all-dishes", handlers.GetAllDishes)
			})

			r.Route("/sub-admin", func(subAdmin chi.Router) {
				subAdmin.Use(middlewares.ShouldHaveRole(models.RoleSubAdmin))
				subAdmin.Post("/create-user", handlers.CreateUser)
				subAdmin.Get("/all-users", handlers.GetAllUsersBySubAdmin)
				subAdmin.Post("/create-restaurant", handlers.CreateRestaurant)
				subAdmin.Get("/all-restaurants", handlers.GetAllRestaurantsBySubAdmin)

				subAdmin.Route("/{restaurantId}", func(restaurantIDRoute chi.Router) {
					restaurantIDRoute.Post("/", handlers.CreateDish)
				})

				subAdmin.Get("/all-dishes", handlers.GetAllDishesBySubAdmin)
			})

			r.Route("/user", func(user chi.Router) {
				user.Use(middlewares.ShouldHaveRole(models.RoleUser))
				user.Get("/all-restaurants", handlers.GetAllRestaurants)
				user.Get("/all-dishes", handlers.GetAllDishes)
				user.Get("/calculate-distance", handlers.CalculateDistance)
			})
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
