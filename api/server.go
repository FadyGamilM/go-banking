package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	// dbStore *transaction.PgTxStore
	router *gin.Engine
	// handler *AccountHandler
}

// func NewServer(txStore *transaction.PgTxStore) *Server {
// 	// initialize the server
// 	server := &Server{
// 		dbStore: txStore,
// 	}

// 	// initialize the router
// 	r := gin.Default()

// 	// define the routes and map them to the handlers
// 	accountsRouter := r.Group("/api/v1/accounts")
// 	// transfersRouter := r.Group("/api/v1/transfers")
// 	// entriesRouter := r.Group("/api/v1/entries")
// 	accountsRouter.POST("/")

// 	// set the router to the server.router
// 	server.router = r

// 	// return the server
// 	return server
// }

func NewServer(accountHandler *AccountHandler, entryHandler *EntryHandler, transferHandler *TransferHandler) *Server {
	// initialize the server
	server := &Server{}

	// initialize the router
	r := gin.Default()

	// define the routes and map them to the handlers
	accountsRouter := r.Group("/api/v1/accounts")
	transfersRouter := r.Group("/api/v1/transfers")
	entriesRouter := r.Group("/api/v1/entries")

	// map the methods to the routes
	accountsRouter.POST("/", accountHandler.HandleCreateAccount)
	entriesRouter.POST("/", entryHandler.HandleCreateEntry)
	transfersRouter.POST("/", transferHandler.HandleCreateTransfer)

	// set the router to the server.router
	server.router = r

	// return the server
	return server
}

func responseWithError(e error) string {
	return e.Error()
}

func (s *Server) Start(addr string) error {
	log.Println("starting the server ..")
	return s.router.Run(addr)
}
