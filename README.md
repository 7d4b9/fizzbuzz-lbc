# Synopsis

# Run the unit tests
```sh
make tests
```

# Run the integration tests
```sh
make integration-tests
```

# Run the fizzbuzz service loccally
```sh
make up
```
A running http server is listening on localhost:8080

# Example

Then

```sh
curl --header "Content-Type: application/json"   --request POST   --data '{"Int1": 2,"Int2": 3, "Limit": 10, "Str1": "fizz", "Str2": "buzz"}' localhost:8080
```

Retrieves

```
1 fizz buzz fizz 5 fizzbuzz 7 fizz buzz fizz
```