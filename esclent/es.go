package esclent

import (
	"context"

	"github.com/olivere/elastic/v7"
)

var es *elastic.Client

func EsConn(esUrl string) {
	if esUrl == "" {
		esUrl = "http://localhost:9200"
	}
	client, err := elastic.NewClient(elastic.SetURL(esUrl), elastic.SetSniff(false))
	if err != nil {
		return
	}
	es = client
}

func InIsExists(indexname string) (bool, error) {
	isexist, err := es.IndexExists(indexname).Do(context.Background())
	return isexist, err

}

func InitCreateIndex(indexname string, doc string) (*elastic.IndicesCreateResult, error) {
	createIndex, err := es.CreateIndex(indexname).BodyString(doc).Do(context.Background())
	return createIndex, err
}
