package main

var instance *singleton = new(singleton)

type singleton struct{}

func GetInstance() *singleton {
	return instance
}
