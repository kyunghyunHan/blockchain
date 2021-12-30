# boltdb

## 권한 변경
```
chmod ??? 파일명
```

## 환경변수 설정
```
vi ~/.zshrc

export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin


source ~/.zshrc
```
## db 다운
```
Boltbrowser:

- URL : https://github.com/br0xen/boltbrowser

- command : go get github.com/br0xen/boltbrowser

Boltdbweb:

- URL : https://github.com/evnix/boltdbweb

- command : go get github.com/evnix/boltdbweb
```
## db실행
```

boltdbweb --db-name=blockchain.db
```

