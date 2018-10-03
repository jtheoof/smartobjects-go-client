package mnubo

import (
	"encoding/json"
	"fmt"
)

const (
	searchPath = "/api/v3/search"
)

func (m *Mnubo) createBasicQuery(mql interface{}, results interface{}) (error) {
	payload, err := json.Marshal(mql)

	if err != nil {
		return fmt.Errorf("Unable to marshal the mql: %s (%s)", mql, err)
	}

	return m.createBasicQueryWithBytes(payload, results)
}

func (m *Mnubo) createBasicQueryWithString(mql string, results interface{}) (error) {
	return m.createBasicQueryWithBytes([]byte(mql), results)
}

func (m *Mnubo) createBasicQueryWithBytes(mql []byte, results interface{}) (error) {
	cr := ClientRequest{
		method: "POST",
		contentType: "application/json",
		path: fmt.Sprintf("%s/basic", searchPath),
		payload: mql,
	}

	return m.doRequestWithAuthentication(cr, results)
}
