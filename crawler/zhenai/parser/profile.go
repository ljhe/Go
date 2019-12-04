package parser

import (
	"Grammar/crawler/engine"
	"Grammar/crawler/model"
	"regexp"
)

var personRe = regexp.MustCompile(`<div .* class="purple-btns">
					   <div .* class="m-btn purple">([^<]+)</div>
					   <div .* class="m-btn purple">([^<]+)</div>
					   <div .* class="m-btn purple">([^<]+)</div>
					   <div .* class="m-btn purple">([^<]+)</div>
					   <div .* class="m-btn purple">([^<]+)</div>
					   <div .* class="m-btn purple">([^<]+)</div>
					   <div .* class="m-btn purple">([^<]+)</div>
					   <div .* class="m-btn purple">([^<]+)</div>
               	       <div .* class="m-btn purple">([^<]+)</div>
				  </div>`)
var imgRe = regexp.MustCompile(`<div .* class="logo f-fl" style="background-image: url([^;]*;([^?]+)?[^)]*);"></div>`)
var nameRe = regexp.MustCompile(`<h1 .* class="nickName">([^<]+)</h1>`)

func ParseProfile(contents []byte) engine.ParseResult {
	profile := model.Profile{}
	//person := personRe.FindSubmatch(contents)
	img := imgRe.FindSubmatch(contents)
	name := nameRe.FindSubmatch(contents)
	profile.Img = string(img[1])
	profile.Name = string(name[1])
	result := engine.ParseResult{
		Items:    []interface{}{profile},
	}
	return result
}

