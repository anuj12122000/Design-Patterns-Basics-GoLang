package main

import (
	"fmt"
	"strconv"
)

type IDBConnection interface {
	getId() string
}

type connection struct {
	id string
}

func (c *connection) getId() string {
	return c.id
}

type poolConnections struct {
	IdleConnection   []IDBConnection
	ActiveConnection []IDBConnection
	capacity         int
}

func (pc *poolConnections) getConnection() IDBConnection {

	if len(pc.IdleConnection) == 0 {
		panic("NO CONNECTION PRESENT IN THE POOL")
	}

	obj := pc.IdleConnection[0]
	pc.IdleConnection = pc.IdleConnection[1:] // yaha pe 1 st ko hataa diya hain
	pc.ActiveConnection = append(pc.ActiveConnection, obj)
	return obj

}

func (pc *poolConnections) returnConnection(target IDBConnection) error {

	err := pc.remove(target)

	if err != nil {
		return err
	}
	pc.IdleConnection = append(pc.IdleConnection, target)
	return nil

}

func (pc *poolConnections) remove(target IDBConnection) error {

	currentActiveLength := len(pc.ActiveConnection)

	for i, obj := range pc.ActiveConnection {
		if obj.getId() == target.getId() {
			pc.ActiveConnection[currentActiveLength-1], pc.ActiveConnection[i] = pc.ActiveConnection[i], pc.ActiveConnection[currentActiveLength-1]
			pc.ActiveConnection = pc.ActiveConnection[:currentActiveLength-1]
			return nil
		}
	}
	return fmt.Errorf("This Target Pool Doesn't Belong To The Pool")
}

func initPool(connections []IDBConnection) *poolConnections {
	if len(connections) == 0 {
		return nil
	}

	active := make([]IDBConnection, 0)

	var pool *poolConnections
	pool = &poolConnections{
		IdleConnection:   connections,
		ActiveConnection: active,
		capacity:         len(connections),
	}
	return pool
}

func main() {

	connections := make([]IDBConnection, 0)

	for i := 0; i < 3; i++ {
		connections = append(connections, &connection{id: strconv.Itoa(i)})
	}

	Pools := initPool(connections)

	conn1 := Pools.getConnection()
	fmt.Println("CONNECT COUNT ", len(Pools.ActiveConnection), " ", conn1.getId())

	err := Pools.returnConnection(conn1)

	if err != nil {
		fmt.Println(err)
	}

}
