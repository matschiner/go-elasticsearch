# go-elasticsearch




import "github.com/matschiner/go-elasticsearch"

func main() {
	client := go-elasticsearch.Create()
  // to connect to a different Host, Port or Protocol use go-elasticsearch.Create(go-elasticsearch.Client{Host: "example.com",Port:"9200",Protocol:"https"})
  
  //index
	i := client.Index(go-elasticsearch.Insert{Index: "test", Type: "test",Id:"42" , Values: `{
		"v1": "test1",
		"v2":"test2"
	}`})
	print(i)
	
	
	r:=client.Delete(go-elasticsearch.Doc{Index: "test",Type:"test",Id:"42"})
	print(r)
	
	r := client.Search(go-elasticsearch.Query{Index: "test", Type: "test",Query: `{
  	"query": {"match_all": {}}
	}`})
	for _,doc := range r.Hits.Hits {
		print(doc.Id)
	}
}
