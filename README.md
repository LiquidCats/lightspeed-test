# Unique Address Counter

Helps to count number of unique IP addresses in the large file

## Usage

**Example:**
```shell
go run cmd/counter/main.go test/assets/ips.txt
```

**Output:**
```shell
Number of unique IP addresses: 7000
```

### OR

**Example**
```shell
go build -o unique_ips main.go
./unique_ips test/assets/ips.txt
```

**Output:**
```shell
Number of unique IP addresses: 7000
```

## Generating Test Data

```shell
go run cmd/generator/main.go
```

**OR**

```shell
go build -o generate_test_data generate_test_data.go
./generate_test_data
```

### Set the number of rows and duplication ration

Open `cmd/generator/main.go` and change this line 

```go
filename := "test/assets/ips.txt" // output file name and directory
totalLines := 10000 // total number of rows
duplicateRatio := 0.3 // duplication percentage
```