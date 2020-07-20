package main
import (
 "fmt"
 "github.com/julienschmidt/httprouter"
 "net/http"
 "log"
 "strconv"
 "os"
)
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
 fmt.Fprint(w, "Welcome!\n")
}
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
 fmt.Fprintf(w, "aloha, %s!\n", ps.ByName("name"))
}

//checks if input is a number
func IsInt(s string) bool {
	_, err := strconv.ParseInt(s,10,64)
	return err == nil
 }

 //recursive fibonacci function 
func FibonacciRecursion(n int) int {
    if n <= 1 {
        return n
    }
    return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
}

// fibonacci sequence generator 
func Fib(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var number string = ps.ByName("number")
	if IsInt(number){
		//if input is an integer, casts variable number from type string to int
		number, err := strconv.Atoi(number)
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}else{
			//prints fibonacci sequence up until specified number 
			for i := 0; i <= number; i++ { 
				fmt.Fprintf(w, strconv.Itoa(FibonacciRecursion(i)) + " ")
			}
		}
	} else{
		fmt.Fprintf(w, "Error: %s is not an integer", number)
	}
}

func main() {
 router := httprouter.New()
 router.GET("/api", Index)
 router.GET("/api/hello/:name", Hello)
 router.GET("/api/fibonacci/:number", Fib)
 log.Fatal(http.ListenAndServe(":8080", router))
}