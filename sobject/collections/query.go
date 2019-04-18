package collections

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/g8rswimmer/goforce"
	"github.com/g8rswimmer/goforce/session"
	"github.com/g8rswimmer/goforce/sobject"
)

type collectionQueryPayload struct {
	IDs    []string `json:"ids"`
	Fields []string `json:"fields"`
}

type query struct {
	session session.ServiceFormatter
}

func (q *query) callout(sobject string, records []sobject.Querier) ([]*goforce.Record, error) {
	if q == nil {
		panic("collections: Collection Query can not be nil")
	}
	payload, err := q.payload(sobject, records)
	if err != nil {
		return nil, err
	}
	c := &collection{
		method:   http.MethodPost,
		body:     payload,
		endpoint: endpoint + "/" + sobject,
	}
	var values []*goforce.Record
	err = c.send(q.session, &values)
	if err != nil {
		return nil, err
	}
	return values, nil
}
func (q *query) payload(sobject string, records []sobject.Querier) (io.Reader, error) {
	fields := make(map[string]interface{})
	ids := make(map[string]interface{})
	for _, querier := range records {
		if sobject != querier.SObject() {
			return nil, fmt.Errorf("sobject collections: sobjects do not match got %s want %s", querier.SObject(), sobject)
		}
		ids[querier.ID()] = nil
		for _, field := range querier.Fields() {
			fields[field] = nil
		}
	}
	queryPayload := collectionQueryPayload{
		IDs:    q.keyArray(ids),
		Fields: q.keyArray(fields),
	}
	payload, err := json.Marshal(queryPayload)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(payload), nil
}
func (q *query) keyArray(m map[string]interface{}) []string {
	array := make([]string, len(m))
	idx := 0
	for k := range m {
		array[idx] = k
		idx++
	}
	return array
}
