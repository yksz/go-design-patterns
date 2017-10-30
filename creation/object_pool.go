package main

import (
	"container/list"
	"fmt"
	"strconv"
	"sync"
)

var (
	uniqueID int
)

type PooledObject struct {
	id int
}

func (o *PooledObject) String() string {
	return "PooledObject [id=" + strconv.Itoa(o.id) + "]"
}

type ObjectPool struct {
	mu     sync.Mutex
	idle   *list.List
	active *list.List
}

func NewObjectPool() *ObjectPool {
	idle := list.New()
	active := list.New()
	return &ObjectPool{idle: idle, active: active}
}

func (p *ObjectPool) BorrowObject() *PooledObject {
	p.mu.Lock()
	defer p.mu.Unlock()

	var object *PooledObject
	if p.idle.Len() <= 0 {
		object = &PooledObject{uniqueID}
		uniqueID++
	} else {
		object = p.removeAt(p.idle, 0)
	}
	fmt.Printf("Borrow: %s\n", object)
	p.active.PushBack(object)
	return object
}

func (p *ObjectPool) ReturnObject(object *PooledObject) {
	p.mu.Lock()
	defer p.mu.Unlock()

	fmt.Printf("Return: %s\n", object)
	p.idle.PushBack(object)
	p.remove(p.active, object)
}

func (p *ObjectPool) remove(list *list.List, object *PooledObject) {
	for e := list.Front(); e != nil; e = e.Next() {
		if object == e.Value.(*PooledObject) {
			list.Remove(e)
			return
		}
	}
}

func (p *ObjectPool) removeAt(list *list.List, index int) *PooledObject {
	for e, i := list.Front(), 0; e != nil; e, i = e.Next(), i+1 {
		if index == i {
			return list.Remove(e).(*PooledObject)
		}
	}
	return nil
}

func main() {
	objectPool := NewObjectPool()
	object1 := objectPool.BorrowObject()
	objectPool.ReturnObject(object1)
	object2 := objectPool.BorrowObject()
	object3 := objectPool.BorrowObject()
	objectPool.ReturnObject(object2)
	objectPool.ReturnObject(object3)
}
