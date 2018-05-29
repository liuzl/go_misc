package main

import (
	"context"
	"fmt"
	"github.com/rs/dnscache"
	"time"
)

func main() {
	fmt.Println("vim-go")
	resolver := &dnscache.Resolver{}

	// First call will cache the result
	addrs, err := resolver.LookupHost(context.Background(), "baidu.com")

	// Subsequent calls will use the cached result
	addrs, err = resolver.LookupHost(context.Background(), "baidu.com")
	fmt.Println(addrs, err)

	// Call to refresh will refresh names in cache. If you pass true, it will also
	// remove cached names not looked up since the last call to Refresh. It is a good idea
	// to call this method on a regular interval.
	go func() {
		clearUnused := true
		t := time.NewTicker(5 * time.Minute)
		defer t.Stop()
		for range t.C {
			resolver.Refresh(clearUnused)
		}
	}()
}
