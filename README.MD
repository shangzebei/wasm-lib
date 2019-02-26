```$shell
export GO111MODULE=on
go mod vendor
go bulid
```
wasm 
```$shell
EMCC_WASM_BACKEND=1 emcc  -std=c++11 -s WASM=1 main.cpp
```