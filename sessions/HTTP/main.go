package main
import(
	"fmt"
	"net/http"
)


func main(){
	http.HandleFunc("/hello",helloHandler)
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080",nil)

}

func helloHandler(w http.ResponseWriter,r *http.Request){

	/*fmt.Println("Method:",r.Method)
	fmt.Println("URL:",r.URL.Path)
	fmt.Fprintln(w,"Hello From Go!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type","text/plain")
*/
		if r.Method!=http.MethodGet{
			http.Error(w,"Method not allowed",http.StatusMethodNotAllowed)
			return
		}
		fmt.Fprintln(w,"Hello From Go!")
	}