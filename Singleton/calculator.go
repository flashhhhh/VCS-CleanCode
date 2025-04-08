package main

import "sync"

type Calculator struct {
}

var lock = &sync.Mutex{}
var instance *Calculator

func GetInstance() *Calculator {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()

		if instance == nil {
			instance = &Calculator{}
		}
	}
	return instance
}

func (c *Calculator) Add(a, b int) int {
	return a + b
}

func (c *Calculator) Subtract(a, b int) int {
	return a - b
}

func (c *Calculator) Multiply(a, b int) int {
	return a * b
}

func (c *Calculator) Divide(a, b int) int {
	return a / b
}