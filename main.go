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

*/
