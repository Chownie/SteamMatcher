package main

import (
	"github.com/hoisie/web"
)

func main() {
	web.Get("/()", IndexLoad)
	web.Post("/()", CompareNames)
	//web.Get("/test/(.*)", TestLoadHome)
	web.Get("/static/(.*)", Sendstatic)
	//STARTING PROCEDURE
	web.Run("0.0.0.0:8830")
}

/*func TestStruct(testPerson *Person) {
	testPerson := Person{}
	testPerson.GetData("plattypuss")
	fmt.Println(testPerson.SteamID)
	for _, value := range testPerson.Games {
		fmt.Println(value.Name)
		fmt.Println(value.Storelink)
		fmt.Println(value.Logo)
		fmt.Println("")
	}
	fmt.Println(testPerson.Avatar)
}*/
