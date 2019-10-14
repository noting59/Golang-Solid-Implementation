package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/noting59/Golang-Solid-Implementation/controllers"
	"github.com/noting59/Golang-Solid-Implementation/infrastructures"
	"github.com/noting59/Golang-Solid-Implementation/repositories"
	"github.com/noting59/Golang-Solid-Implementation/services"
	"github.com/thedevsaddam/renderer"
	"log"
	"sync"
)

type IServiceContainer interface {
	InjectCartController() controllers.CartController
	InjectProductController() controllers.ProductController
	InjectOrderController() controllers.OrderController
	InjectPayController() controllers.PayController
}

var (
	rnd               *renderer.Render
	postgreSQLHandler *infrastructures.PostgreSQLHandler
	orderRepository   *repositories.OrderRepository
	productRepository *repositories.ProductRepository
	userRepository    *repositories.UserRepository
	cartRepository    *repositories.CartRepository
	userService       *services.UserService
	productService    *services.ProductService
	orderService      *services.OrderService
	cartService       *services.CartService
	paymentService    *infrastructures.SolidPayments
)

func init() {
	rnd = renderer.New(renderer.Options{
		TemplateDir: "templates",
	})

	config := infrastructures.Config{}

	db, err := sql.Open("postgres", config.GetConf().PostgreSQLConn)
	if err != nil {
		log.Fatal(err)
	}
	postgreSQLHandler = &infrastructures.PostgreSQLHandler{}
	postgreSQLHandler.Conn = db

	orderRepository = &repositories.OrderRepository{IDbHandler: postgreSQLHandler}
	userRepository = &repositories.UserRepository{IDbHandler: postgreSQLHandler}
	productRepository = &repositories.ProductRepository{IDbHandler: postgreSQLHandler}
	cartRepository = &repositories.CartRepository{IDbHandler: postgreSQLHandler}

	userService = &services.UserService{IUserRepository: userRepository}
	productService = &services.ProductService{IProductRepository: productRepository, ICartRepository: cartRepository}
	cartService = &services.CartService{ICartRepository: cartRepository, IProductService: productService, IUserService: userService}

	orderService = &services.OrderService{IOrderRepository: orderRepository, IProductService: productService, IUserService: userService}
	paymentService = &infrastructures.SolidPayments{
		IOrderService: orderService,
		ICartService:  cartService,
		IUserService:  userService,
	}
}

type kernel struct{}

func (k *kernel) InjectPayController() controllers.PayController {
	payController := controllers.PayController{
		Render:      rnd,
		IPayHandler: paymentService,
	}
	return payController
}

func (k *kernel) InjectOrderController() controllers.OrderController {
	orderController := controllers.OrderController{
		Render:        rnd,
		ICartService:  cartService,
		IOrderService: orderService,
	}
	return orderController
}

func (k *kernel) InjectCartController() controllers.CartController {
	cartController := controllers.CartController{
		Render:       rnd,
		ICartService: cartService,
	}
	return cartController
}

func (k *kernel) InjectProductController() controllers.ProductController {
	productController := controllers.ProductController{
		Render:          rnd,
		IProductService: productService,
	}
	return productController
}

var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}
