# blockchain

## Hash 사용

```
genesisBlock.hash = sha256.Sum256(genesisBlock.data+genesisBlock.provHash)
```


## 첫번쨰 블록생성
```
type block struct {
	data     string
	hash     string
	prevHash string
}

func main() {
	genesisBlock := block{"Genesis Block", "", ""}
	hash := sha256.Sum256([]byte(genesisBlock.data + genesisBlock.prevHash))
	hexHash := fmt.Sprintf("%x", hash)
	genesisBlock.hash = hexHash
	fmt.Println(genesisBlock)
}

```
첫번쨰 블록 생성

```
{Genesis Block 89eb0ac031a63d2421cd05a2fbe41f3ea35f5c3712ca839cbf6b85c4ee07b7a3 }
```
## 2번쨰 3번쨰 블록생성 후 연결

- blockchain struct -> block의 slice를 정의

- 이전 hash값 연결
- 블록체인의 길이가 0인지 확인
- 이전 해시값은 블록체인의 마지막 블록의 값과 동일
- -1의 경우 배열의 길이에서 -1값이 마지막 값
- 해시값으로 바꾼 후 블록 배열에서 추가

