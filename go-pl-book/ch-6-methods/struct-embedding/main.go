package main

import "sync"

/*
 When we embed structs into another structs, all of the methods are accessble by the instances of this new structs
 But if the embedded struct belongs to another package, then only the exported properties and methods are accessible

 Methods can only be declared on named types Like type Point struct{X,Y int}, but due to embedding, it's also possible for "unnamed types" to have their own methods

 Consider this simple cache implementation with a mutex to guard it's access
*/

var (
	mu      sync.Mutex // mutexes that guard resources, are often placed on top of those resources
	mapping = make(map[string]string)
)

func Lookup(key string) string {
	mu.Lock()
	v := mapping[key]
	mu.Unlock()
	return v
}

// Instead of creating 2 seperate global variables, we create a single variable and compose it using struct embedding
// cache's type is an unnamed struct type, but it got all the methods of embedded structs in it
var cache = struct {
	sync.Mutex // This is struct embedding, now we can directly access lock() unlock() through cache
	mapping    map[string]string
}{mapping: make(map[string]string)}

func Lookup2(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}
