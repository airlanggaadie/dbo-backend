package handler

func (h *Handler) routes() {
	// health check segment
	h.router.GET("/healthz", h.healthCheck)

	// auth segment
	h.authRoutes()

	// user segment
	h.userRoutes()

	// order segment
	h.orderRoutes()
}

func (h *Handler) authRoutes() {
	authGroup := h.router.Group("/auth")
	authGroup.POST("/login", h.login)
	authGroup.GET("/login/report", h.loginReport)
}

func (h *Handler) userRoutes() {
	h.router.GET("/users", h.listUser)

	userGroup := h.router.Group("/user")
	userGroup.GET("/:id", h.userDetail)
	userGroup.GET("/search", h.searchUser)
	userGroup.POST("", h.addNewUser)
	userGroup.PUT("/:id", h.updateUser)
	userGroup.DELETE("/:id", h.deleteUser)
}

func (h *Handler) orderRoutes() {
	h.router.GET("/orders", h.listOrder)

	orderGroup := h.router.Group("/order")
	orderGroup.GET("/:id", h.orderDetail)
	orderGroup.GET("/search", h.orderSearch)
	orderGroup.POST("", h.addNewOrder)
	orderGroup.PUT("/:id", h.updateOrder)
	orderGroup.DELETE("/:id", h.deleteOrder)
}
