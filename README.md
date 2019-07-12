# MD5 Hash tool

This is a tool written in golang that hashes the response of urls asynchronously using the MD5 hashing algorithm and writes the string encoded value of the result to the standard output.

## Installation

Install Golang and configure it properly. See docs here for more info : 
https://golang.org/doc/install

## Usage

```python
go build main.go

./main http://www.adjust.com http://www.facebook.com http://www.google.com

./main -parallel 3 http://www.adjust.com http://www.facebook.com http://www.google.com
```

To limit the number of parallel requests, use flag 'parallel'. By default, the value is 10.

```python
./main -parallel 3 http://www.adjust.com http://www.facebook.com http://www.google.com
```

To run all the tests:  

```python
make test
```
