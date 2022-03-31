# CRUD
Clean Architecture, REST API, PostgeSQL

#Create user {"name":"some name","age":"24"}
#Make friends {"source_id":"1","target_id":"2"}
#Update user age {"id":"1","new_age":"28"}
#Delete User {"id":"1"}



    POST("/create", h.createNewUser)
		POST("/make_friends", h.makeFriends)
		GET("/", h.getAllUser)
		GET("/friends/:id", h.getUserByID)
		PUT("/age_updated/:id", h.updateUser)
		DELETE("/delete/:id", h.deleteUser)
