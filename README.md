# gopher-translator-service

REST service which transforms English words to gopher language using predefined rules

### Requirements:
    1. go 
    2. docker && docker-compose
    3. make (for ease of use)
    4. cURL (if you want to kick off with some requests)

### Start service:
######(will default to 8080)
   ```
   make run
   ```
###### to run on specific port
```
go run ./cmd/main.go -port <PORT>
```
### Start in container:
```
make run-container
```

### Build image:
```
make build-image
```

### How to run tests:
1. ##### Run all tests
    ```
    make test
    ```
2. ##### Run only unit tests
   ```
   make test-unit
   ```
3. ##### Run only integration tests
   ```
   make test-integration
   ```

### Generate mocks:
```
make generate
```

### Requests to get started (API is versioned!) 
```
curl -X POST localhost:<PORT>/v1/word \
   -H 'Content-Type: application/json' \
   -d '{"english_word":"apple"}'
```
```
curl -X POST localhost:<PORT>/v1/sentence \
   -H 'Content-Type: application/json' \
   -d '{"english_sentence":"Apples grow on trees."}'
```
```
curl GET localhost:<PORT>/v1/history 
```