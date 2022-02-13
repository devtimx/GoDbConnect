# GoDbConnect

A basic structure with a connection to the PostgreSQL database, and structure of directories, is ideal for initializing project in GO.

Replace for your data in to next files:

-connectDB.go
    - Replace for your data \<user\>\<password\>\<server address\>
	    
        - connStr := "postgres://<user>:<password>@<server address>:5432/gotest?sslmode=disable"

-handlers.go

    - Add routers for endpoints
	- Example:
	    -router.HandleFunc("/signup", middleware.CheckDB(routers.SignUp)).Methods("POST")