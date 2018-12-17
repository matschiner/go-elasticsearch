package elastic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	Host     string
	Port     int
	Protocol string
}
type Query struct {
	Index, Type, Query string
	Size               int
}
type Insert struct {
	Index, Type, Values, Id string
}
type ResponseSearch struct {
	Hits struct {
		Total    int     `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index  string                 `json:"_index"`
			Type   string                 `json:"_type"`
			Id     string                 `json:"_id"`
			Score  float64                `json:"_score"`
			Source map[string]interface{} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}
type Doc struct {
	Index, Type, Id string
}

func Create(in ...Client) Client {
	var c Client
	if len(in) == 0 {
		c = Client{}
	} else {
		c = in[0]
	}
	if c.Host == "" {
		c.Host = "localhost"
	}
	if c.Protocol == "" {
		c.Protocol = "http"
	}
	if c.Port == 0 {
		c.Port = 9200
	}
	return c
}
func (c Client) Search(q Query) ResponseSearch {
	var jsonStr = []byte(q.Query)
	if q.Size == 0 {
		q.Size = 10
	}
	url := fmt.Sprintf("%s://%s:%d/%s/%s/_search?pretty&size=%d", c.Protocol, c.Host, c.Port, q.Index, q.Type, q.Size)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	r := ResponseSearch{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		print(err.Error())
	}
	return r

}

func (c Client) Index(d Doc, values string) map[string]interface{} {
	var jsonStr = []byte(values)
	var url string
	var req *http.Request
	var err error

	if d.Id == "" {
		url = fmt.Sprintf("%s://%s:%d/%s/%s", c.Protocol, c.Host, c.Port, d.Index, d.Type)
		req, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	} else {
		url = fmt.Sprintf("%s://%s:%d/%s/%s/%s", c.Protocol, c.Host, c.Port, d.Index, d.Type, d.Id)
		req, err = http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
	}
	// print("1",url)

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	r := new(map[string]interface{})
	// r.String = string(body)
	// m := map[string]string{}
	err = json.Unmarshal(body, &r)
	// fmt.Println(r)
	if err != nil {
		print(err.Error())
	}
	return *r

}

func (c Client) Update(d Doc, values string) map[string]interface{} {
	var jsonStr = []byte(`{"doc": ` + values + `}`)
	var url string
	var req *http.Request
	var err error

	url = fmt.Sprintf("%s://%s:%d/%s/%s/%s/_update", c.Protocol, c.Host, c.Port, d.Index, d.Type, d.Id)
	req, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	// print("1",url)

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	r := new(map[string]interface{})
	// r.String = string(body)
	// m := map[string]string{}
	err = json.Unmarshal(body, &r)
	// fmt.Println(r)
	if err != nil {
		print(err.Error())
	}
	return *r

}
func (c Client) Delete(d Doc) map[string]interface{} {
	var jsonStr = []byte(``)
	var url string
	var req *http.Request
	var err error
	if d.Id != "" {
		url = fmt.Sprintf("%s://%s:%d/%s/%s/%s", c.Protocol, c.Host, c.Port, d.Index, d.Type, d.Id)
		req, err = http.NewRequest("DELETE", url, bytes.NewBuffer(jsonStr))
	} else {
		print("No Id specified")
	}
	// print("1",url)

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	r := new(map[string]interface{})
	// r.String = string(body)
	// m := map[string]string{}
	err = json.Unmarshal(body, &r)
	// fmt.Println(r)
	if err != nil {
		print(err.Error())
	}
	return *r

}
