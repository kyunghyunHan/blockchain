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
```
func (b *blockchain) getLastHash() string {
	if len(b.blocks) > 0 {
		return b.blocks[len(b.blocks)-1].hash
	}
	return ""
}

```
- 이전 hash값 연결
- 블록체인의 길이가 0인지 확인
- 이전 해시값은 블록체인의 마지막 블록의 값과 동일
- -1의 경우 배열의 길이에서 -1값이 마지막 값
```

func (b *blockchain) addBlock(data string) {
	newBlock := block{data, "", b.getLastHash()}
	hash := sha256.Sum256([]byte(newBlock.data + newBlock.prevHash))
	newBlock.hash = fmt.Sprintf("%x", hash)
	b.blocks = append(b.blocks, newBlock)
}

```
- 해시값으로 바꾼 후 블록 배열에서 추가


## Singleton Patton

```
func GetBlockchain() *blockchain {
	if b == nil {
		b = &blockchain{}
	}
	return b
}
```
- Singleton 패턴이란 우리의 application내에서 어디든지 블록체인의 단 하나의 instance만을 공유하는 방법
- getBlockchatin이라는 function이라는 생성
- b변수와동일한 타입인 blockchain 포인터를 반환. 블록체인의 instance를 생성후 리턴 b를 반환 = 블록체인이 어떻게 최기화 되고 공유될지를 제어
