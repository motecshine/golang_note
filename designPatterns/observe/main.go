package main

// RabbitMQ 可以使用这样的方式去处理断线重连 更新channel

import (
	"fmt"
	"time"
)

// Event 定一个事件类型，发生事件驱动时将事件传递给观察者们
type Event struct {
	Data string
}

// Observer 定义观察者和观察对象的接口。
// Observer 定义了一个更新发生事件的标准接口，
// Subject是具体被观察的接口，他有注册观察者、注销观察者和发布通知的三个主要函数接口
type Observer interface {
	//更新事件
	Update(*Event)
}

// Subject 被观察的对象接口
type Subject interface {
	//注册观察者
	Regist(Observer)
	//注销观察者
	Deregist(Observer)

	//通知观察者事件
	Notify(*Event)
}

// ConcreteObserver 定义一个观察者
type ConcreteObserver struct {
	Id int
}

// Update 实现Update的方法
func (co *ConcreteObserver) Update(e *Event) {
	fmt.Printf("observer [%d] recieved msg: %s.\n", co.Id, e.Data)
}

// ConcreteSubject 被观察者对象
type ConcreteSubject struct {
	Observers map[Observer]struct{}
}

// Regist 注册要观察的对象
func (cs *ConcreteSubject) Regist(ob Observer) {
	cs.Observers[ob] = struct{}{}
}

// Deregist 删除要观察的对象
func (cs *ConcreteSubject) Deregist(ob Observer) {
	delete(cs.Observers, ob)
}

// 通知每个观察者事件
func (cs *ConcreteSubject) Notify(e *Event) {
	for ob, _ := range cs.Observers {
		ob.Update(e)
	}
}

func main() {
	cs := &ConcreteSubject{
		Observers: make(map[Observer]struct{}),
	}
	//实例化两个观察者
	cobserver1 := &ConcreteObserver{1}
	cobserver2 := &ConcreteObserver{2}
	//注册观察者
	cs.Regist(cobserver1)
	cs.Regist(cobserver2)

	for i := 0; i < 5; i++ {
		e := &Event{fmt.Sprintf("msg [%d]", i)}
		cs.Notify(e)

		time.Sleep(time.Duration(1) * time.Second)
	}
}
