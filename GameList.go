package main

import (
	"github.com/hoisie/web"
)

//The HTML+MST frame, stored in const to eliminate disk read slowdown
const FRAME = `<!DOCTYPE HTML>
<html>
	<head>
		<meta charset="UTF-8">
		<link rel="stylesheet" type="text/css" href="/static/site.css" />
		<title>Steam Game Comparison</title>
	</head>
	<body>
		<div id="header">
			{{{header}}}
		</div>
		<div id="content">
			{{{content}}}
		</div>
	</body>
</html>`

const FORM = `<form method="GET" action="/match/">
	<input type="text" name="PersonA" />
	<input type="text" name="PersonB" /><br />
	<input type="submit" value="Go!" />
</form>
`

const RESULT = `
			<div id="columnA">
				{{{colA}}}
			</div>
			<div id="columnB">
				{{{colB}}}
			</div>`

//The HTML+MST game div, stored in const to eliminate disk read slowdown
const GAME = `<div class="game {{common}}">
	<div class="gameicon"><img src="{{logo}}" /></div>
	<div class="gamename"><a href="{{gameurl}}">{{gamename}}</a></div>
</div>`

func TestLoadHome(val string) string {
	person := Person{}
	person.GetData(val)
	GameList := ""
	for _, value := range person.Games {
		gameinfo := map[string]string{"logo": value.Logo, "gameurl": value.Storelink, "gamename": value.Name}
		GameList += Rendermustache(GAME, &gameinfo)
	}
	return Rendermustache(FRAME, &map[string]string{"colA": GameList})
}

func IndexLoad(val string) string {
	content := Rendermustache(FORM, &map[string]string{})
	return Rendermustache(FRAME, &map[string]string{"content": content})
}

func CompareNames(ctx *web.Context, val string) string {
	personA := Person{}
	personB := Person{}
	personA.GetData(ctx.Params["PersonA"])
	personB.GetData(ctx.Params["PersonB"])
	dupes := FindDuplicates(&personA, &personB)
	listA := ""
	listB := ""
	dupelist := ""
	for _, keys := range *dupes {
		dupeinfo := map[string]string{"logo": keys.Logo, "gameurl": keys.Storelink, "gamename": keys.Name, "common": "common"}
		dupelist += Rendermustache(GAME, &dupeinfo)
	}
	for _, keys := range personA.Games {
		listAinfo := map[string]string{"logo": keys.Logo, "gameurl": keys.Storelink, "gamename": keys.Name}
		listA += Rendermustache(GAME, &listAinfo)
	}
	for _, keys := range personB.Games {
		listBinfo := map[string]string{"logo": keys.Logo, "gameurl": keys.Storelink, "gamename": keys.Name}
		listB += Rendermustache(GAME, &listBinfo)
	}
	content := Rendermustache(RESULT, &map[string]string{"colA": dupelist + "<hr />" + listA, "colB": dupelist + "<hr />" + listB})
	return Rendermustache(FRAME, &map[string]string{"content": content})
}

func FindDuplicates(personA, personB *Person) *[]Game {
	common := []Game{}
	i := 0
	ii := 0
	for ii < len(personA.Games) {
		j := 0
		jj := 0
		for jj < len(personB.Games) {
			if personA.Games[ii].Name != personB.Games[jj].Name {
				personB.Games[j] = personB.Games[jj]
				j += 1
			}
			jj += 1
		}
		personB.Games = personB.Games[:j]
		if j == jj {
			personA.Games[i] = personA.Games[ii]
			i += 1
		} else {
			common = append(common, personA.Games[ii])
		}
		ii += 1
	}
	personA.Games = personA.Games[:i]
	return &common
}
