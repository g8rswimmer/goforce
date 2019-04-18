package collections

import (
	"io"
	"net/http"

	"github.com/g8rswimmer/goforce/session"
	"github.com/g8rswimmer/goforce/sobject"
)

type UpdateValue struct {
	sobject.InsertValue
}

type update struct {
	session session.ServiceFormatter
}

func (u *update) callout(allOrNone bool, records []sobject.Updater) ([]UpdateValue, error) {
	payload, err := u.payload(allOrNone, records)
	if err != nil {
		return nil, err
	}
	c := &collection{
		method:   http.MethodPatch,
		body:     payload,
		endpoint: endpoint,
	}
	var values []UpdateValue
	err = c.send(u.session, &values)
	if err != nil {
		return nil, err
	}
	return values, nil
}
func (u *update) payload(allOrNone bool, recs []sobject.Updater) (io.Reader, error) {
	records := make([]interface{}, len(recs))
	for idx, updater := range recs {
		rec := map[string]interface{}{
			"attributes": map[string]string{
				"type": updater.SObject(),
			},
		}
		for field, value := range updater.Fields() {
			rec[field] = value
		}
		rec["id"] = updater.ID()
		records[idx] = rec
	}
	return dmlpayload(allOrNone, records)
}
