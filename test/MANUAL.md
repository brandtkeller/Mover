# Manual Testing
```
curl http://localhost:8080/status \
    --include \
    --header "Content-Type: application/json" \
    --request "GET" 
```
```
curl http://localhost:8080/status \
    --include \
    --header "Content-Type: application/json" \
    --request "PUT" \
    --data '{"Status": "up"}'
```
```
curl http://localhost:8080/status \
    --include \
    --header "Content-Type: application/json" \
    --request "GET" 
```
```
curl http://localhost:8080/state \
    --include \
    --header "Content-Type: application/json" \
    --request "GET" 
```
```
curl http://localhost:8080/lifecycle \
    --include \
    --header "Content-Type: application/json" \
    --request "GET" 
```
```
curl http://localhost:8080/lifecycle \
    --include \
    --header "Content-Type: application/json" \
    --request "PUT" \
    --data '{"Operating": true}'
```