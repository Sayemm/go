package main

import (
	"ecommerce/cmd"
)

func main() {
	cmd.Serve()
}

/*
Domain Driven Design
====================
- Each Domain will contain it's OWN BUSINESS LOGIN
- Domain will be independent, smaller/larger but independent
- If anything happens with a domain, nothing will happen to another domain

Rules
=====
- Parent will never access the child directly
- child can access parent/grandparent....

Facebook Post (Feature)
-----------------------
Post Domain
	- Profile Domain
	- Reaction Domain
	- Comment Domain
	- Share Domain
	- created_at
	- content

	Profile Domain
		- name
		- dp
	Reaction Domain
		- like
		- haha
		- wow
		- angry
		- care
	Comment Domain
		- who
		- what
	Share Domain
		- by who
		- total Shares

Total Reaction
Total Comments

- One domain might Contains another domain
- Each Domain will contain it's OWN BUSINESS LOGIN

DOMAIN
======
- domain folder
- each domain folder will have 2 files
	- service.go
	- port.go
		- user domain can communicate with db, redis, another domain
		- user jar sathe communicate kore / dependent -> port.go maintains signature of those


- DDD is a way to organize and structure your code around the business logic
┌──────────────────────────────┐
│        REST Handlers         │  ← API Layer (HTTP)
├──────────────────────────────┤
│         Service Layer        │  ← Business Logic
├──────────────────────────────┤
│         Repository Layer     │  ← Talks to Database
├──────────────────────────────┤
│          Domain Layer        │  ← Pure Business Entities
└──────────────────────────────┘

CreateUser (handler)
   ↓
svc.Create(user)       ← calls service method
   ↓
repo.Create(user)      ← service calls repository
   ↓
INSERT INTO users...   ← repository writes to DB

=>  The handler depends on a Service interface.
	The service depends on a UserRepo interface.
	The repo is the actual implementation that talks to the database.
Each layer depends on an interface, not on a concrete struct.
That’s what makes the code loosely coupled and testable.


Why do I need a Service layer?
Why can’t my HTTP handler just call repo.CreateUser() directly?
===============================================================
- Contain the business logic, service layer is the business brain
- Keep the handler (HTTP) focused only on input/output
- Decouple the app from the database
- Let's say Before creating a user, we must hash the password, validate email, and check if it already exists.”
	- Where do we write that logic?
	- If we don’t have a service layer, we’ll likely add that inside the handler or repo. (WRONG APPROACH)
*/
