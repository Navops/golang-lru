// Package lru provides LRU cache based on hashicorp's LRU cache that
// uses size instead of count
//
// The LRU cache is a simple LRU cache that has a memory size
// limit. You should know the sizes of objects you're adding to the
// cache, and you should not be adding objects that are larger than
// the limit. There are no limits on the number of items in the cache
// as long as the memory size limit is not exceeded.
//
// This is forked from Hashicorp's LRU cache: https://github.com/hashicorp/golang-lru
//
package lru
