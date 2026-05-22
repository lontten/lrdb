package main

import (
	"encoding/json"
	"example/dbinit"
	"fmt"

	"github.com/lontten/lcore/v2/types"
	"github.com/lontten/lrdb"
)

func main2() {
	var u = dbinit.User{
		Id:   types.NewInt(1),
		Name: nil,
		Age:  45,
	}
	err := lrdb.Set(dbinit.DB, "a", u, 0)
	if err != nil {
		panic(err)
	}

	user, err := lrdb.Get[dbinit.User](dbinit.DB, "a")
	if err != nil {
		panic(err)
	}
	b, _ := json.MarshalIndent(user, "", "  ")
	fmt.Println(string(b))
}

func main() {
	var u = dbinit.User{
		Id:   types.NewInt(1),
		Name: nil,
		Age:  45,
	}
	var u2 = dbinit.User{
		Id:   types.NewInt(2),
		Name: nil,
		Age:  222,
	}
	var list = []dbinit.User{u, u2}

	err := lrdb.Set(dbinit.DB, "a", list, 0)
	if err != nil {
		panic(err)
	}

	list2, err := lrdb.Get[[]dbinit.User](dbinit.DB, "a")
	if err != nil {
		panic(err)
	}
	b, _ := json.MarshalIndent(list2, "", "  ")
	fmt.Println(string(b))
}
