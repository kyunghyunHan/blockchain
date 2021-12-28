# Server-side 렌더링

## server
```
const port string = ":4000

func home(rw http.ResponseWriter, r *http.Request){ 
  fmt.Fprint(rw,"hello from hone!")
  }
 func main(){
http.HandleFunc("/",home) 
fmt.println("Listening on http://localhost%s\n",port)  //프린트
http.ListenAndServer(port,nil)    //서버여는 용도
}
```
- template폴더 생성
- Fatal은 os.Exit(1)다음에 따라나오는 error를 Print()하는것과 동일
- Must:Template이나 error를 반환하는 함수의 호출을 감싸서 error가 있다면 error를 출력



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
- home 은 2개의 인자를 받는 함수 
- ResponseWrite : 유저에게 보내고 싶은 데이터
- Request의 pointer -> request가 file이 될수도 있고 빅데이터 가될수도 있음 그래서 복사하기보다는 실제 request를 사용하는게 알맞음
- homedata 구조체에 정의 후 data변수에 데이타 저장후 출력
```
...gohtml
 <h1>{{.PageTitle}}</h1>
 {{range .Blocks}}
 <div>
   <ul>
     <li>{{.Data}}</li>
     <li>{{.Hash}}</li>
     <li>{{.PrevHash}}</li>
 {{end}}
 
```

- go에서 받아온 데이터 출력(대문자 사용)
- range .구조체
- {{define""}}{{end}} ->재사용 저장
- {{template""}} ->불러오기

```
const templateDir string = "templates/"
templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))	  
```
- pages 폴더안에서 .gohtml로 끝나는 모든 파일을 가져오기

```
func add(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(rw, "add", nil)
	case "POST":
		r.ParseForm()
		data := r.Form.Get("blockData")
		blockchain.GetBlockchain().AddBlock(data)
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
	}
}

```
- templates.ExecuteTemplate(rw, "add", nil) = add페이지 렌더링
- Get = 
- Post = 
