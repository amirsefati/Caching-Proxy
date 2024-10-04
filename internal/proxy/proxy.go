package proxy

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type ProxyObject struct {
	Origin string
	Cache  map[string]*cache.CacheObject
	Mutex  sync.RWMutex
}

func NewProxy(origin string) *ProxyObject {
	return &ProxyObject{
		Origin: origin,
		Cache:  make(map[string]*cache.CacheObject),
	}
}

func (p *ProxyObject) ClearCache() {
	p.Mutex.Lock()
	p.Cache = make(map[string]*cache.CacheObject)
	p.Mutex.Unlock()
	fmt.Println("Cache Cleared Successfully")
}

// Flow of the Request
// Create cache key -> Check if the Key is present
// 	 Using Simple key for now for example GET method on dummyjson.com
// 	 CACHE_KEY = GET:https://dummyjson.com
//
// If Yes -> Respond with the cached http.Response and Body
//    Set the HEaders
//    X-Cache ; HIT
//
// If No -> Forward the request to Origin
//    Cache the Origin Response
//    Set the Headers
//    X-Cache : MISS
//    Write the reponse


func RespondWithHeaders(w http.ResponseWriter, response http.Response, body []byte, cacheHeader, KEY string) {
	fmt.Printf("Cache : %s %s \n", cacheHeader, KEY)
	w.Header().Set("X-Cache", cacheHeader)
	w.WriteHeader(response.StatusCode)
	for k, v := range response.Header {
		w.Header()[k] = v
	}
	w.Write(body)
}
