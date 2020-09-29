# Customer Invitation
This project implements the system to generate a list of customers who should be invited for an office party. 
We want to invite only those customers who are within the given distance limit.

Given the input file with customers, location of the office, and the distance limit, this program processes the
input file line-by-line and outputs all eligible customers
who are within the given distance limit, sorted by the User ID (in ascending order).



## Input format
The input file follows the JSON lines format. 
```
{"latitude": "52.986375", "user_id": 12, "name": "Christina McArdle", "longitude": "-6.043701"}
{"latitude": "51.92893", "user_id": 1, "name": "Alice Cahill", "longitude": "-10.27699"}
...
```
## Output format
The output file follows the JSON format. 
```
[
    {
        "user_id": 4,
        "name": "Ian Kehoe"
    },
    {
        "user_id": 5,
        "name": "Nora Dempsey"
    },
    {
        ...
    }
]
```

## Implementation Details
**Programming Language:** GoLang 1.14 (refer to go.mod file)

**Installation:** Golang (Version 1.14) - can be installed from [here](https://golang.org/doc/install?download=go1.14.6.darwin-amd64.pkgl)

**Requirements:**
- github.com/stretchr/testify v1.6.1 (refer to go.mod file)

## Instructions to run the code

1) Go to `CustomerInvitation` folder on the command prompt

2) Build the main.go file
    ```
    $ go build main.go
    ```
5) Run the main.go file
    ```
    $ ./main
    ```

## Instructions for System Tests

**Option 1:** Go to `invitation` package and run the tests

1) Go to `CustomerInvitation/internal/invitation` folder

2) Run `go test`
    ```
    $ go test
    ```
**Option 2:** Run tests from `CustomerInvitation` folder
```
$ go test ./...
```