package datasources

import "net/http"

const tjceURL = "https://esaj.tjce.jus.br"

type TJCearaDatasource struct{}

func (d TJCearaDatasource) buildCompleteURL() string {
	params := map[string]any{
		"processo.codigo": "01Z081I9T0000",
		"processo.foro":   "1",
		"processo.numero": "0070337-91.2008.8.06.0001",
	}

	return addQueryParamsToURL(tjceURL+"/cpopg/show.do", params)
}

func (d TJCearaDatasource) SearchFirstInstance(filter SearchFilter) (respBody []byte, err error) {
	var req *http.Request
	req, err = mountRequest(d.buildCompleteURL(), http.MethodGet, "")
	return doRequest(req)
}

func (d TJCearaDatasource) SearchSecondInstance(
	filter SearchFilter,
) (respBody []byte, err error) {
	return
}
