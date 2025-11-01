# AndroidLibXrayLite

## Build requirements
* JDK
* Android SDK
* Go
* gomobile

## Build instructions
1. `git clone [repo] && cd AndroidLibXrayLite`
2. `gomobile init`
3. `go mod tidy -v`
4. `gomobile bind -v -androidapi 21 -ldflags='-s -w' ./`
5. 
## Build 16kb
5. `env CGO_LDFLAGS="-s -w -Wl,-z,max-page-size=16384" gomobile bind -v -androidapi 21 ./`
