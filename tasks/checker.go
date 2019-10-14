package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"test-task/infrastructures"
	"test-task/repositories"
	"test-task/services"
)

func main() {
	//nextTime := time.Now().Truncate(time.Minute)
	//nextTime = nextTime.Add(time.Minute)
	//time.Sleep(time.Until(nextTime))
	Task()
}

// Spider scans website's market.
func Task() {
	config := infrastructures.Config{}

	db, err := sql.Open("postgres", config.GetConf().PostgreSQLConn)
	if err != nil {
		log.Fatal(err)
	}
	postgreSQLHandler := &infrastructures.PostgreSQLHandler{}
	postgreSQLHandler.Conn = db

	orderRepository := &repositories.OrderRepository{IDbHandler: postgreSQLHandler}
	userRepository := &repositories.UserRepository{IDbHandler: postgreSQLHandler}
	productRepository := &repositories.ProductRepository{IDbHandler: postgreSQLHandler}
	cartRepository := &repositories.CartRepository{IDbHandler: postgreSQLHandler}

	userService := &services.UserService{IUserRepository: userRepository}
	productService := &services.ProductService{IProductRepository: productRepository, ICartRepository: cartRepository}
	cartService := &services.CartService{ICartRepository: cartRepository, IProductService: productService, IUserService: userService}

	orderService := &services.OrderService{IOrderRepository: orderRepository, IProductService: productService, IUserService: userService}

	paymentService := &infrastructures.SolidPayments{
		IOrderService: orderService,
		ICartService: cartService,
		IUserService: userService,
	}

	list := orderService.GetUnProcessedOrders("processing")

	log.Print(list)

	for _, elem := range list {
		_, err := paymentService.StatusCheck(elem.Id)

		if err != nil {
			log.Print(err)
		}

		log.Print(fmt.Sprintf("Order %d processed", elem.Id))
	}
}
