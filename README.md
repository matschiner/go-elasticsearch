```go
package main
import "github.com/matschiner/go-elasticsearch"

func main() {
    client := elastic.Create()
    // to connect to a different Host, Port or Protocol use 
    // elastic.Create(elastic.Client{Host: "example.com",Port:"9200",Protocol:"https"})

    //index
    insert := client.Index(elastic.Doc{Index: "test", Type: "test", Id:"42"}, `{
        "v1": "test1",
        "v2":"test2"
    }`)
    fmt.Println(insert)

    r := client.Search(elastic.Query{Index: "test", Type: "test",Query: `{
        "query": {"match_all": {}}
    }`})

    // interate through every hit of search
    for _,doc := range r.Hits.Hits {
        fmt.Println(doc.Id)
    }


    // Delete
    delete:=client.Delete(elastic.Doc{Index: "test",Type:"test",Id:"42"})
    fmt.Println(delete)
}
```

```javascript
var s = "JavaScript syntax highlighting";
alert(s);
```
