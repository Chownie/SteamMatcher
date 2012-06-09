package main

import (
	"encoding/xml"
	"net/http"
	"strings"
	"io/ioutil"
	"fmt"
	"github.com/hoisie/mustache"
	"github.com/hoisie/web"
)

type Person struct {
	Name    string `xml:-`
	SteamID string `xml:"steamID"`
	Avatar  string `xml:"avatarMedium"`
	Games   []Game `xml:"games>game"`
}

type Game struct {
	Name      string `xml:"name"`
	Storelink string `xml:"storeLink"`
	Logo      string `xml:"logo"`
}

func LogError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func (pe *Person) GetXML(url []string) {
	XMLResource, _HTTPerr := http.Get(strings.Join(url, ""))
	XMLSource, _iotuilerr := ioutil.ReadAll(XMLResource.Body)
	LogError(_HTTPerr)
	LogError(_iotuilerr)
	XMLResource.Body.Close()
	xml.Unmarshal(XMLSource, pe)
}

func (pe *Person) GetData(name string) {
	pe.Name = name
	pe.GetXML([]string{"http://steamcommunity.com/id/", name, "/games?xml=1"})
	pe.GetXML([]string{"http://steamcommunity.com/id/", name, "?xml=1"})
}

func Loadmustache(filename string, args *map[string]string) string {
	file, _err := ioutil.ReadFile("Mst/" + filename)
	if _err != nil {
		fmt.Println(_err)
		return "File not found"
	}
	data := mustache.Render(string(file), args)
	return data
}

func GetMustache(filename string) string {
	file, _err := ioutil.ReadFile("Mst/" + filename)
	LogError(_err)
	return string(file)
}

func DelIndex(slice []Game, i int) []Game {
	slice = append(slice[:i], slice[i+1:]...)
	return slice
}

func Rendermustache(mst string, args *map[string]string) string {
	data := mustache.Render(mst, args)
	return data
}

func Sendstatic(ctx *web.Context, val string) {
	file, _err := ioutil.ReadFile("static/" + val)
	if _err != nil {
		return
	}
	filetype := strings.Split(val, ".")
	ctx.ContentType(filetype[len(filetype)-1])
	ctx.WriteString(string(file))
}
