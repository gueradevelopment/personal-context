package db

// Model interface to wrap data types
type Model interface{}

// Result struct to wrap channel tuple
type Result struct {
	Result Model
	Err    error
}

// ResultArray struct to wrap channel tuple
type ResultArray struct {
	Result []Model
	Err    error
}

// Database interface to wrap data accessors
type Database interface {
	Get(string, chan Result)
	GetAll(chan ResultArray, map[string][]string)
	Add(Model, chan Result)
	Edit(Model, chan Result)
	Delete(string, chan Result)
}
