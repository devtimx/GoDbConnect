# GoDbConnect

A basic structure with a connection to the PostgreSQL database, and structure of directories, is ideal for initializing project in GO.

Replace for your data into next files:

- connectDB.go
    
        Replace for your data \<user\>\<password\>\<server address\>
	    
        connStr := "postgres://<user>:<password>@<server address>:5432/gotest?sslmode=disable"

- handlers.go

        Add routers for endpoints
	    Example:

	    router.HandleFunc("/signup", middleware.CheckDB(routers.SignUp)).Methods("POST")

- Packages used

    - github.com/dgrijalva/jwt-go
    - github.com/gorilla/mux
    - github.com/lib/pq
    - github.com/rs/cors