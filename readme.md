# WebAssembly 예제 - winodw 기준

## 일반 Go 버전

### 실행 방법

1. WebAssembly 모듈 빌드:
    ```sh
    $env:GOOS = "js"
    $env:GOARCH = "wasm"
    go build -o main.wasm main.go
    ```

    ```sh
    $env:GOOS = "js"
    $env:GOARCH = "wasm"
    go build -o image_process.wasm image_process.go
    ```

2. `wasm_exec.js` 파일 위치:
    - `GOROOT/misc/wasm` 폴더 내부에 위치해 있습니다.

3. 웹 서버 실행:
    ```sh
    python -m http.server 3000
    ```

    - 웹 브라우저에서 [http://localhost:3000](http://localhost:3000)을 열어 확인합니다.

---

## TinyGo 버전

### 실행 방법

1. WebAssembly 모듈 빌드:
    ```sh
    tinygo build -o ./tiny/tiny-main.wasm -target wasm -opt=2 ./main.go
    ```

    ```sh
    tinygo build -o ./tiny/tiny-image_process.wasm -target wasm -opt=2 ./image_process.go
    ```

2. `wasm_exec.js` 파일 위치:
    - `tinygoRoot/misc/wasm` 폴더 내부에 위치해 있습니다.

3. 웹 서버 실행:
    ```sh
    python -m http.server 3000
    ```

    - 웹 브라우저에서 [http://localhost:3000](http://localhost:3000)을 열어 확인합니다.

---

## 필요사항

- Go 또는 TinyGo 설치
- Python 설치 (웹 서버 실행용)

## 파일 구조

```plaintext
.
├── main.go                # 이미지 처리 Go 코드
├── main.wasm              # 빌드된 WebAssembly 파일 (Go 버전)
├── tiny
│   └── main.wasm          # 빌드된 WebAssembly 파일 (TinyGo 버전)
├── wasm_exec.js           # WebAssembly 실행 환경 스크립트
└── index.html             # 웹 인터페이스
```
---
### 실행 방법 요약
```sh
$env:GOOS = "js"
$env:GOARCH = "wasm"
go build -o main.wasm main.go

python -m http.server 3000
```

```sh
tinygo build -o ./tiny/main.wasm -target wasm -opt=2 ./main.go

python -m http.server 3000
```