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
