package utils

import (
	"container/list"
	"errors"

	"github.com/Knetic/govaluate"
)

type Queue struct {
	queue *list.List
}

func NewQueue() *Queue {
	newQueue := &Queue{
		queue: list.New(),
	}
	return newQueue
}

func (this *Queue) PushBack(item interface{}) {
	this.queue.PushBack(item)
}

func (this *Queue) PushFront(item interface{}) {
	this.queue.PushFront(item)
}

func (this *Queue) GetFront() *list.Element {
	return this.queue.Front()
}

func (this *Queue) GetBack() *list.Element {
	return this.queue.Back()
}

func (this *Queue) PopFront() (interface{}, error) {
	item := this.queue.Front()
	var ret interface{}
	if item != nil {
		ret = item.Value
		this.queue.Remove(item)
	} else {
		return nil, errors.New("queue empty")
	}
	return ret, nil
}

func (this *Queue) Clear() {
	this.queue.Init()
}

func ParseExpr(expr string, parameters map[string]interface{}) bool {
	expression, err := govaluate.NewEvaluableExpression(expr)

	// parameters := make(map[string]interface{}, 8)

	result, err := expression.Evaluate(parameters)
	if err != nil {
		panic(err)
	}

	if ret, ok := result.(bool); ok {
		return ret
	} else {
		panic("result is not bool")
	}
}
