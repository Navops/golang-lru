golang-lru
==========

This provides the `lru` package which implements a fixed-size thread
safe LRU cache. It is based on the Hashicorp LRU cache, which is based
on the cache in Groupcache. This package has two differences:
  * Cache limit is based on object-size, not object-count
  * There is an optional TTL for cached items. 
  
If there is a nonzero TTL, items are removed from the cache lazily as
they're accessed. This is useful to set a maximum time for an item to
stay in the cache to prevent stale reads if a process updates the
data.

Documentation
=============

Full docs are available on [Godoc](http://godoc.org/github.com/bserdar/golang-lru)

Example
=======

Using the LRU is very simple:

```go
// Cache with a 128-byte limit
l, _ := New(128)
for i := 0; i < 256; i++ {
    // You have to pass the byte-size of the item you're putting into the cache
    l.Add(i, i, 8)
}
// Len gives the number of items in the cache
if l.Len() != 16 {
    panic(fmt.Sprintf("bad len: %v", l.Len()))
}
```
