```$shell
export GO111MODULE=on
go mod vendor
go bulid
```
wasm 
```$shell
EMCC_WASM_BACKEND=1 emcc  -std=c++11 -s WASM=1 main.cpp
```
### progress
- [X] File IO
- [x] Sigal
- [x] Time
- [x] Math
- [x] http
- [x] regex
- [ ] Thread
- [ ] c++11 std

### how to use
```$shell
➜  wasmgo git:(master) ✗ ./wasmgo                                                                   
  -f string
        wasm file name 
  -n string
        wasm method name (default "_main")
  -v    open debug mode (-v open)

```
### eg.
```$shell
   ./wasmgo -f /Users/shang/Desktop/code/CoinIDLib/build/core.wasm -n _main -v
```
## windows build cgo
```bash
1. install gcc-5.1.0-tdm64-1-core.zip #set to path
2. install Cygwin # install make 
```