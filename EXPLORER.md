# Server-side 렌더링

## server
```
const port string = ":4000

func home(rw http.ResponseWriter, r *http.Request){
  fmt.Fprint(rw,"hello from hone!")
  }
 func main(){
http.HandleFunc("/",home)
fmt.println("Listening on http://localhost%s\n",port)
http.ListenAndServer(port,nil)
}
```
