package user

import (
	"ecommerce/domain"
	userHandler "ecommerce/rest/handlers/user"
)

type Service interface {
	userHandler.Service //embedding
}

// user jader upor dependent sob signature ekhane thakbe
type UserRepo interface { // I need these two methods // I won't implement this, repo layer will do this // this interface will be embedded in repo layer
	Create(user domain.User) (*domain.User, error)
	Find(email, pass string) (*domain.User, error)
}

/*
UserRepo interface
------------------
- Purpose: Define a contract for how you can interact with the persistence layer (database).
- It says: “Any repository that stores users must implement these methods: Create and Find.
- It does not care how the data is stored — MySQL, MongoDB, in-memory, fake for tests.

Service interface
-----------------
- Purpose: Define what the service layer (application logic) must do.
- Handlers (create_user.go) only know about Service interface.
- Go lets us embed interfaces inside other interfaces.
- This means: “Service must implement all methods from userHandler.Service

	Go interface embedding is not classical inheritance
	---------------------------------------------------
		- Service includes all methods of userHandler.Service.
		- userHandler.Service doesn’t know anything about Service.
		- Service promises to have all methods that userHandler.Service has, plus any extra methods I add.
		- The embedded interface (userHandler.Service) cannot force the embedding interface (Service) to implement anything beyond what it already defines.


**** SEE DETAILS ON service.go for how Service is implementing Create and Find
*/
