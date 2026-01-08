package main

import (
	"ecommerce/cmd"
	"sync"
)

var cnt int64
var mu sync.Mutex

func main() {
	cmd.Serve()
}

/*
Locking
=======
- multiple goroutine tried to access
	- one will lock then other cannot access that variable (blocked)
	- blocked goroutines will go to sleep (no extra cpu cost)
	- will wake up when other goroutine will unlock that variable

- process p1 and process p2 both needs shared data1 and data2
- p1 comes then lock data1 and p2 comes then lock data2
- now p1 is waiting for data1 and p2 is waiting for data1 as they both need that data (cannot unlock without getting data1/data2)
- DEADLOCK!
- THAT'S WHY GO CHANNEL IS MORE RELIABLE THAT LOCKING!

Go Channel
==========
- pipe (input and output)
- data input and data output
- go routine uses go channel/pipe
- G1 sends data, G3 will receive it (G1 ---->go channel---> G3)
- G1 and G2 sends data (G3 will receive whatever sent data first)

- what if receiver goroutines executes first?
	- receiver goroutine will wait for data and go runtime will sent it to sleep
	- go runtime will awake it when data will be available on the channel
	- G1 will go to sleep after publishing data to the channel (NOT DONE YET!)
	- As long as noone is receiving data it will be stuck/deadlock
	- when someone will receive that publisied data, G1 will be awaken

- 2 types of go channel
	a: unbuffered
		- one slot, at a time cannot store more than 1 data
		- what if go routines send data to the channel and nothing receives from that channel (ERROR: BLOCKED)
		- sender go routine will not awake (deadlock)
	b: buffered
		- multiple slot
		- sender never go to sleep (publish and done)
		- in bufferred channel if there are slots available after publishing data and sender goroutine will be done (will not go to sleep)
		- but if we publish more than slot then it will be stuck
*/

/*
Domain Driven Design
====================
** Business Rules = Decisions That Affect Business Value

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


The “cleanest” DDD practice is:
-------------------------------
Always return interfaces from constructors that represent cross-layer boundaries
(like repository → service, service → handler, etc.)

DDD Rules
---------
- High-level code (like services) should depend on interfaces, not concrete structs.
- The layer that uses an interface defines it — not the one that implements it.
	- handler defines the interface that it uses
	- service implementes those for handler
	- and what service needs it just defines in interface
	- and repo implements those interface for service
- Keep Domain “Pure” - domain contain entities, value objects, and domain logic
- Keep Infrastructure Replaceable
- Composition Root Wires Everything
- Dependencies point inward — toward the domain, not outward.


PAGINATION
----------
- when we send request like GET, it can return huge data and that data is saved to the RAM
- If we load huge data to RAM that would be inefficient
*/
