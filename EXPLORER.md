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
- template폴더 생성
- 



```

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}
func home(rw http.ResponseWriter, r *http.Request) {
tmpl := template.Must(template.ParseFiles("templates/home.html"))
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
	tmpl.Execute(rw, data)
  }
```
```
...gohtml
 <h1>{{.PageTitle}}</h1>
```
- go에서 받아온 데이터 출력(대문자 사용)
