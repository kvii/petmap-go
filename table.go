package main

// 数据源采用内存数据模拟

const tableName = "tables.json"

var tables = Tables{
	Users: []User{
		{UserName: "kvii", Password: "123"},
		{UserName: "张三", Password: "321"},
	},
	Pets: []Pet{
		{PetName: "狗", Owner: "kvii", Longitude: 35.9518869, Latitude: 120.1850354},
	},
	Homes: []Home{
		{Owner: "kvii", Longitude: 35.9518869, Latitude: 120.1850354},
	},
	Messages: []Message{},
}
