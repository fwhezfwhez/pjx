package main

import (
	"log"
	"os"
)

var logger = log.New(os.Stderr, "[pjx] ", log.LstdFlags|log.Llongfile)
