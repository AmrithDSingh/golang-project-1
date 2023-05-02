package main

import (
    "fmt"
    "math/rand"
    "net/http"
    "time"
)

//Sets up the server, unsure how this works
func main() {
    http.HandleFunc("/flipcoin", flipCoinHandler)
    http.ListenAndServe(":0808", nil)
}

//Handles request for flip coin path
func flipCoinHandler(w http.ResponseWriter, r *http.Request) {
    //Initialising the random number, unix nano ensures the seed to be called
    //The nano timer compares the time from the first request, if time has passed it provides a new result
    rand.Seed(time.Now().UnixNano())
    //assigns result to heads
    result := "heads"
    //if the new number is not equal to heads value then assign tails
    if rand.Intn(2) == 0 {
        result = "tails"
    }
    fmt.Println("The coin flip result is =", result)
}