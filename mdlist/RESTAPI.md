# RESTAPI
```

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL:         "/",
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         "/blocks",
			Method:      "POST",
			Description: "Add A Block",
			Payload:     "data:string",
		},
	}
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)
}
```
- json.Marshal:encoding한 interface를 return,메모리 형식으로 저장된 객체를 ,저장/송신 할수 있도록 변환
- Unmarshal:JSON을 받아서 go로 변환
- json.NewEncoder():writer에 작성하는 encoder를 반환
- Encode:go의 interface를 받아

## struct field tag
```
type URLDescription struct {
	URL         string `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

```
- 기본적으로 struct가 json으로 변환하는 방식을 바꾸는 방법
- 소문자로 쓰면 export되지 않기 때문에 소문자를 쓸수없음 이럴떄 사용
- omitempty :Field가 비어있으면 Field를 숨겨줌



```
type URL string

func (u URL) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

```
- MarshalText :byte slice와 error를 return
- URL이 어떻게 json으로 marshal될지를 정할수 있음

## blocks/get,blocks/post

```
func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBlocks())
	case "POST":
		var addBlockBody AddBlockBody
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
		blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated)
	}
}

```
- post는 블록을 생성할떄 사용
- get은 블록체인을 확인할 떄 사용,블록체인의 모든 블록을 확인 가능
- Encode가 Marshal의 일을 해주고 결과를 ResponseWrite에 작성해줌
- Post에서 받아와서 golang의 struct로 변환
- REST Client :백엔드에 REST quest를 보내도록
