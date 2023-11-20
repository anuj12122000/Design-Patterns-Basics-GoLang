package main

import "fmt"

type Publisher interface {
	Register(Subscriber)
	NotifyAll(msg string)
}

type publisher struct {
	subscriberList []Subscriber
}

func (pub *publisher) Register(sub Subscriber) {
	pub.subscriberList = append(pub.subscriberList, sub)
}

func (pub *publisher) NotifyAll(msg string) {
	for _, sub := range pub.subscriberList {
		fmt.Println("informing the Subscriber with Subscriber ID ", sub.(subscriber).subscriberId) // use of Type Assertion Here
		sub.ReactToPublisherMsg(msg)
	}
}

type Subscriber interface {
	ReactToPublisherMsg(msg string)
}

type subscriber struct {
	subscriberId string
}

func GetNewSubscriber(id string) subscriber {
	return subscriber{subscriberId: id}
}

func (s subscriber) ReactToPublisherMsg(msg string) {
	fmt.Println("Subscriber Received the Message ", msg, " For Subscriber Id -> ", s.subscriberId)
}

func main() {
	pub := &publisher{subscriberList: make([]Subscriber, 0)}

	sub1 := GetNewSubscriber("1")
	sub2 := GetNewSubscriber("2")

	pub.Register(sub1)
	pub.Register(sub2)
	pub.NotifyAll("RELEASE THIS MESSAGE")
}
