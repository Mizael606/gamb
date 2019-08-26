package controllers

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Soap(url string, xml string) revel.Result {
	req, _ := http.NewRequest("POST", url, strings.NewReader(xml))
	req.Header.Add("content-type", "application/xml")
	req.Header.Add("cache-control", "no-cache")

	type response struct {
		body    string
		success bool
		err     bool
	}

	res, err := http.DefaultClient.Do(req)
	Response := make(map[string]interface{})
	if res.StatusCode == 200 {
		ResponseBody, _ := ioutil.ReadAll(res.Body)
		Response["status"] = true
		Response["body"] = string(ResponseBody)
		Response["success"] = true
		Response["err"] = false
		defer res.Body.Close()
		c.Response.ContentType = "text/json"
		return c.RenderJSON(Response)
	} else {
		ResponseBody, _ := ioutil.ReadAll(res.Body)
		Response["status"] = false
		Response["body"] = err
		Response["success"] = false
		Response["err"] = true
		defer res.Body.Close()
		c.Response.ContentType = "text/json"
		return c.RenderJSON(ResponseBody)
	}

}
