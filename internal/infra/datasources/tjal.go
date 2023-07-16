package datasources

import (
	"errors"
	"io"
	"net/http"
)

const tjalURL = "https://www2.tjal.jus.br"

type TJAlagoasDatasource struct{}

func (d TJAlagoasDatasource) buildCompleteURL() string {
	params := map[string]any{
		"conversationId":           "",
		"cbPesquisa":               "NUMPROC",
		"numeroDigitoAnoUnificado": "0710802-55.2018",
		"foroNumeroUnificado":      "0001",
		"dadosConsulta.valorConsultaNuUnificado": []string{
			"0710802-55.2018.8.02.0001",
			"UNIFICADO",
		},
		"dadosConsulta.valorConsulta":  "",
		"dadosConsulta.tipoNuProcesso": "UNIFICADO",
	}

	return addQueryParamsToURL(tjalURL+"/cpopg/search.do", params)
}

func (d TJAlagoasDatasource) DoRequest() (respBody []byte, err error) {
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
