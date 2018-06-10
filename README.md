# hunter2

A super simple go package that validates if a desired password is in the top 100k most commonly used.

## Installation
1. `go get github.com/justinwcooper/hunter2`

## Usage
```go
validator := hunter2.New()

validator.isCommon("hunter2") // => true
validator.isNotCommon("3smVBU348%P8qp") // => true
```

## Acknowledgements
* [SecLists](https://github.com/danielmiessler/SecLists)
* [Name Inspiration](http://knowyourmeme.com/memes/hunter2)
