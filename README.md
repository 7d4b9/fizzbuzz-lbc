# Fizzbuzz

## Unit tests

```sh
make unit-tests
```

## Integration tests

```sh
make integration-tests
```

## Example usage

```sh
go run main.go
```

```sh
curl --header "Content-Type: application/json"   --request POST   --data '{"Int1": 2,"Int2": 3, "Limit": 10, "Str1": "fizz", "Str2": "buzz"}' localhost:8080
```

```sh
1 fizz buzz fizz 5 fizzbuzz 7 fizz buzz fizz
```

## Statistics endpoint

Return the parameters corresponding to the most used request, as well as the number of hits for this requests.

```sh
curl  localhost:8080/statistics
```

```sh
{"Pattern":"/fizzbuzz/{\"Int1\": 2,\"Int2\": 3, \"Limit\": 10, \"Str1\": \"fizz\", \"Str2\": \"buzz\"}","Count":4}
```
