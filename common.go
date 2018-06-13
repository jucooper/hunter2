package hunter2

import (
	"bufio"
	"log"
	"os"
)

const defaultFilePath = "./passwords.txt"

// Common ...
type Common interface {
	IsCommon(password string) bool
	IsNotCommon(password string) bool
}

// Passwords ...
type Passwords struct {
	List *os.File
}

// IsCommon ...
func (p *Passwords) IsCommon(password string) bool {
	scanner := bufio.NewScanner(p.List)
	for scanner.Scan() {
		if scanner.Text() == password {
			return true
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	defer p.List.Close()
	return false
}

// IsNotCommon ...
func (p *Passwords) IsNotCommon(password string) bool {
	return !p.IsCommon(password)
}

// New ...
func New() *Passwords {
	file, err := os.Open(defaultFilePath)
	if err != nil {
		log.Fatal(err)
	}
	return &Passwords{file}
}
