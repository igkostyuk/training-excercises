package main

import (
	"log"
	"os"

	srv "github.com/igkostyuk/training-excercises/day_2/service"
)

// Data to decode
var jsonStr = []byte(`
{
    "things": [
        {
            "name": "Alice",
            "age": 37
        },
        {
            "city": "Ipoh",
            "country": "Malaysia"
        },
        {
            "name": "Bob",
            "age": 36
        },
        {
            "city": "Northampton",
            "country": "England"
        },
 		{
            "name": "Albert",
            "age": 3
        },
		{
            "city": "Dnipro",
            "country": "Ukraine"
        },
		{
            "name": "Roman",
            "age": 32
        },
		{
            "city": "New York City",
            "country": "US"
        }
    ]
}`)

func main() {
	// logger to Inject
	logger := log.New(os.Stdout, "INFO: ", 0)
	service := srv.New(logger)
	service.Decode(jsonStr)
}
