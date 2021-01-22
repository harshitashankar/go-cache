package cache

import (
	"testing"
)

func TestSet(t *testing.T) {
	cache := New()
	cache.Set("lemon.com", "172.28.5.64")
	_, exists := cache.Get("lemon.com")
	if !exists {
		t.Errorf("TestSet: no entry for %v in cache", "lemon.com")
	}
}

func TestGet(t *testing.T) {
	cache := New()
	cache.Set("lemon.com", "172.28.5.64")
	ip, exists := cache.Get("lemon.com")
	if !exists {
		t.Errorf("TestGet: no entry for %v in cache", "lemon.com")
	}
	if ip!="172.28.5.64" {
		t.Errorf("TestGet: cache entry does not match")
	}
	//test when key does not exist
	_, exists = cache.Get("ginger.com")
	if exists {
		t.Errorf("TestGet: Unexpected entry for %v in cache", "lemon.com")
	}

}

func TestRemove(t *testing.T) {
	cache := New()
	cache.Set("lemon.com", "172.28.5.64")
	err := cache.Remove("lemon.com")
	if err != nil{
		t.Error(err)
	}
	_, exists := cache.Get("ginger.com")
	if exists {
		t.Errorf("TestRemove: Unexpected entry for %v in cache", "lemon.com")
	}
	//test to check nothing bad happens when try to remove key that doesn't exist

	err = cache.Remove("IDon'tExist")
	if err != nil{
		t.Error(err)
	}
	
}