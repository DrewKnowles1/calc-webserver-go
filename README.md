# calc-webserver-go

basic go webserver calculator

run with:
``` go run . ```

query with:
```  curl -X POST -d '{"firstnumber": 2, "secondnumber": 0, "operator": "/"}' localhost:8080/calc ```

feel free to substitute operators/values