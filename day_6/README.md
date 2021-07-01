# Http server

Your main tasks:
0. Replace Existing in-memory Repository with PostgreSQL repo.
1. Extend this application to have: book in-memory concurrent-safe cache with 4 methods (Set,Get,Delete, Flush all), which will be used in service layer and will update\remove books when they are modified.
Use interfaces and not call it directly. Put it into separate package.
1. Implement a Job that will Flush a cache in the backgroun every 15 minutes. This value should be configurable, not hardcoded. This is a business logic so put it in the corresponding place!
2. Add new auth middleware that will check for an `Authorization` header. Header should contain `Basic <encoded base64 username:password>` . Just Retrieve this values from request using `request.BasicAuth()` and parse using
```b
    result, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}
```
For auth to be successful user should be `test` with `test_pwd` password. In other cases throw 403.
3. (If there is time) Add new `Business layer` that will work only with repository and will be called by transport layer, from http handlers code.
You should have `type BookService interface` for that layer. This service layer will know nothing about jsons/http etc.
###Steps to do:

1. Run `go mod download` or `go mod tidy` to add packages from module to your local cache so they can be used by your application
2. Add new repository package `postgres` that will work with PostgreSQL and replace it with existing in-memory repo. 
3. Add `NewRepository` function that will return new PostgreRepo instance with connection inside.  
4. `	conn, err := sql.Open("postgres", dsn)` How to open conn.
5. Call `MigrateUp` during db initialization. Provide correct path to the migrate path. No hardcoding!   
5. Import github.com/lib/pq
4. Use blank import to import driver using `_ "github.com/lib/pq"` inside your new repo package file where connection with DB is initialized
5. Use provided docker-compose file to run this db when you're ready to run the server. 
6. Set `SQL_DSN` connection string to connect to postgre and read it on application startup using `os.Getenv()` function. DSN string:
 `SQL_DSN=postgres://postgres:secret@localhost:5432/test?sslmode=disable`
7. Implement methods to work with `Books`
8. In the main file just replace the Repo with new Repo.
9. Implement Auth middlware structure in analogy with existing middlware and add it to `r.Use(logMD, authMD)`
10. Implement Cache.
11. Implement background Flush job.
12. 