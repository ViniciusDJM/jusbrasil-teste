package datasources

import (
	"errors"
	"io"
	"net/http"
)

const tjceURL = "https://esaj.tjce.jus.br"

type TJCearaDatasource struct{}

// processo.codigo=01Z081I9T0000&processo.foro=1&processo.numero=0070337-91.2008.8.06.0001
func (d TJCearaDatasource) buildCompleteURL() string {
	params := map[string]any{
		"processo.codigo": "01Z081I9T0000",
		"processo.foro":   "1",
		"processo.numero": "0070337-91.2008.8.06.0001",
	}

	return addQueryParamsToURL(tjceURL+"/cpopg/show.do", params)
}

func (d TJCearaDatasource) DoRequest() (respBody []byte, err error) {
	var req *http.Request
	if req, err = http.NewRequest(http.MethodGet, d.buildCompleteURL(), nil); err != nil {
		return
	}

	req.Header.Add("cookie", "JSESSIONID=AED4E4A2A8603ADEB26440A87551AEA8.cpopg8")

	var res *http.Response
	if res, err = http.DefaultClient.Do(req); err != nil {
		return
	}

	defer func(Body io.ReadCloser) {
		if closeErr := Body.Close(); closeErr != nil {
			err = errors.Join(err, closeErr)
		}
	}(res.Body)

	var body []byte
	if body, err = io.ReadAll(res.Body); err == nil && body != nil {
		respBody = body
	}

	if contentErr := checkHTMLContentType(res, body); contentErr != nil {
		err = errors.Join(err, contentErr)
	}

	if statusErr := checkStatus(res.StatusCode, http.StatusOK); statusErr != nil {
		err = errors.Join(err, statusErr)
	}

	return
}
