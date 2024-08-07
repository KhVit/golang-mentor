package main

import (
	"fmt"
	"strings"
)

type Room struct {
	Name     string
	Comment  string
	NextRoom []string
	Loot     map[string][]string
}

func (room Room) LootInRoom() string {
	var res string
	if len(room.Loot) > 0 {
		for place, items := range room.Loot {
			res += place + ": "
			for _, item := range items {
				res += item + ", "
			}
		}
	}

	return res
}

func (room Room) itemInRoom(item string) bool {
	for _, vals := range room.Loot {
		for _, val := range vals {
			if val == item {
				return true
			}
		}
	}
	return false
}

func (room Room) ListNextRooms() string {
	var res string
	if len(room.NextRoom) > 0 {
		res += "можно пройти - "
		for index, nextRoom := range room.NextRoom {
			if index > 0 {
				res += ", "
			}
			res += nextRoom
		}
	}

	return res
}

func (room *Room) deleteItemRoom(item string) {
	for key, vals := range room.Loot {
		for ind, val := range vals {
			if val == item {
				room.Loot[key] = append(vals[:ind], vals[ind+1:]...)
				if len(room.Loot[key]) == 0 {
					delete(room.Loot, key)
				}
			}
		}
	}
}

type Gamer struct {
	CurRoom   string
	Inventory []string
	BackPack  bool
	DoorOpen  bool
	Rooms     map[string]Room
}

func (g *Gamer) LookAround() string {
	room := g.Rooms[g.CurRoom]
	resMsg := room.Comment + " " + room.LootInRoom() + room.ListNextRooms()

	return resMsg
}

func containsItem(items []string, item string) bool {
	for _, i := range items {
		if i == item {
			return true
		}
	}
	return false
}

func (g *Gamer) Move(roomName string) string {
	var res string
	room := g.Rooms[g.CurRoom]
	if _, ok := g.Rooms[roomName]; ok {
		if containsItem(room.NextRoom, roomName) {
			if roomName == "улица" && g.DoorOpen == false {
				return "дверь закрыта"
			}
			g.CurRoom = roomName
			room := g.Rooms[g.CurRoom]
			res = room.Comment + ". " + room.ListNextRooms()
		} else {
			res = "нет пути в " + roomName
		}
	} else {
		res = "нет такой комнаты"
	}

	return res
}

func (g *Gamer) Wear(item string) string {
	var res string
	if item == "рюкзак" && g.BackPack == false {
		room := g.Rooms[g.CurRoom]
		if room.itemInRoom(item) {
			g.BackPack = true
			room.deleteItemRoom(item)
			res = "вы надели: " + item
		}
	} else {
		res = "нет такого"
	}

	return res
}

func (g *Gamer) Take(item string) string {
	var res string
	room := g.Rooms[g.CurRoom]
	if room.itemInRoom(item) {
		if g.BackPack {
			g.Inventory = append(g.Inventory, item)
			room.deleteItemRoom(item)
			res = "предмет добавлен в инвентарь: " + item
		} else {
			res = "некуда класть"
		}
	} else {
		res = "нет такого"
	}

	return res
}

func (g *Gamer) Apply(data []string) string {
	var res string
	key := data[0]
	applyTo := (data[1])
	if containsItem(g.Inventory, key) {
		if key == "ключи" && applyTo == "дверь" {
			g.DoorOpen = true
			res = "дверь открыта"
		} else {
			res = "не к чему применить"
		}
	} else {
		res = "нет предмета в инвентаре - " + key
	}

	return res
}

var GameWorld Gamer

func main() {
	// в этой функции можно ничего не писать но тогда у вас не будет работать через go run main.go очень круто будет сделать построчный ввод команд тут, хотя это и не требуется по заданию
	//fmt.Println("Application start...\n")

	initGame()
	fmt.Println(handleCommand("осмотреться"))
	fmt.Println(handleCommand("идти коридор"))
	fmt.Println(handleCommand("идти комната"))
	fmt.Println(handleCommand("осмотреться"))
	fmt.Println(handleCommand("надеть рюкзак"))
	fmt.Println(handleCommand("взять ключи"))
	fmt.Println(handleCommand("взять конспекты"))
	fmt.Println(handleCommand("идти коридор"))
	fmt.Println(handleCommand("применить ключи дверь"))
	fmt.Println(handleCommand("идти улица"))

	initGame()
	fmt.Println(handleCommand("осмотреться"))
	fmt.Println(handleCommand("завтракать"))
	fmt.Println(handleCommand("идти комната"))
	fmt.Println(handleCommand("идти коридор"))
	fmt.Println(handleCommand("применить ключи дверь"))
	fmt.Println(handleCommand("идти комната"))
	fmt.Println(handleCommand("осмотреться"))
	fmt.Println(handleCommand("взять ключи"))
	fmt.Println(handleCommand("надеть рюкзак"))
	fmt.Println(handleCommand("осмотреться"))
	fmt.Println(handleCommand("взять ключи"))
	fmt.Println(handleCommand("взять телефон"))
	fmt.Println(handleCommand("взять ключи"))
	fmt.Println(handleCommand("осмотреться"))
	fmt.Println(handleCommand("взять конспекты"))
	fmt.Println(handleCommand("осмотреться"))
	fmt.Println(handleCommand("идти коридор"))
	fmt.Println(handleCommand("идти кухня"))
	fmt.Println(handleCommand("осмотреться"))
	fmt.Println(handleCommand("идти коридор"))
	fmt.Println(handleCommand("идти улица"))
	fmt.Println(handleCommand("применить ключи дверь"))
	fmt.Println(handleCommand("применить телефон шкаф"))
	fmt.Println(handleCommand("применить ключи шкаф"))
	fmt.Println(handleCommand("идти улица"))
}

// эта функция инициализирует игровой мир - все команты, если что-то было - оно корректно перезатирается
func initGame() {
	fmt.Println(`Welcome to the game!
You can use 5 actions:
-> 1 - LookAround (осмотреться в комнате)
-> 2 - Move (переместиться в другую комнату)
-> 3 - Wear (надеть предмет)
-> 4 - Take (положить предмет в рюкзак)
-> 5 - Apply (применить предмет)`)

	GameWorld = Gamer{
		CurRoom:   "кухня",
		Inventory: make([]string, 4),
		BackPack:  false,
		DoorOpen:  false,

		Rooms: map[string]Room{
			"кухня": Room{
				Name:     "кухня",
				Comment:  "ты находишься на кухне",
				NextRoom: []string{"коридор"},
				Loot: map[string][]string{
					"на столе": []string{"чай"},
				},
			},
			"коридор": Room{
				Name:     "коридор",
				Comment:  "ничего интересного",
				NextRoom: []string{"кухня", "комната", "улица"},
				Loot:     map[string][]string{}, //make(map[string][]string),
			},
			"комната": Room{
				Name:     "комната",
				Comment:  "ты в своей комнате",
				NextRoom: []string{"коридор"},
				Loot: map[string][]string{
					"на столе": {"ключи", "конспекты"},
					"на стуле": {"рюкзак"},
				},
			},
			"улица": Room{
				Name:     "улица",
				Comment:  "на улице весна",
				NextRoom: []string{"домой"},
				Loot:     map[string][]string{},
			},
		},
	}
}

// данная функция принимает команду от "пользователя" и наверняка вызывает какой-то другой метод или функцию у "мира" - списка комнат
func handleCommand(command string) string {
	cmd := strings.Split(command, " ")
	l := len(cmd)

	switch {
	case cmd[0] == "осмотреться" && l == 1:
		return GameWorld.LookAround()
	case cmd[0] == "идти" && l == 2:
		return GameWorld.Move(cmd[1])
	case cmd[0] == "надеть" && l == 2:
		return GameWorld.Wear(cmd[1])
	case cmd[0] == "взять" && l == 2:
		return GameWorld.Take(cmd[1])
	case cmd[0] == "применить" && l == 3:
		return GameWorld.Apply([]string{cmd[1], cmd[2]})
	}
	return "неизвестная команда"
}
