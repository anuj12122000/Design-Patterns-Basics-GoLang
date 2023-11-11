package main

import "fmt"

// target
type Mobile interface {
	chargeAppleMobile()
}

type appleMobile struct{}

func (a *appleMobile) chargeAppleMobile() {
	fmt.Printf("Charging My Apple Mobile")
}

//adaptee
type androidMobile struct{}

func (android *androidMobile) chargeAndroidMobile() {
	fmt.Printf("Charging My Android Mobile")
}

//making Adapter here
type androidAdapter struct {
	androidMobile *androidMobile
}

func (androidAdapter *androidAdapter) chargeAppleMobile() {
	androidAdapter.androidMobile.chargeAndroidMobile()
}

//client
type client struct{}

func (c *client) chargeMobile(mob Mobile) { // yaha pe mob Mobile main jo bhi struct aapka interface ka method ka implementation hoga usko aap waha pe paas kar sakte hoon
	// in short jo mob pass ho raha hain necessary hain ke wo interface[target] ke methods ko satisfy kare

	mob.chargeAppleMobile()
}

func main() {
	apples := &appleMobile{} //made apple struct here
	client := &client{}
	client.chargeMobile(apples)

	androids := &androidMobile{}
	androidAdapter := &androidAdapter{
		androidMobile: androids,
	}
	client.chargeMobile(androidAdapter)
	main.goC:\Users\anujj\OneDrive\Desktop\design_patterns\adapter\main.go
}
