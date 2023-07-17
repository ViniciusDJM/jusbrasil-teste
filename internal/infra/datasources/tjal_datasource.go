package datasources

import (
	"net/http"
	"strconv"
)

const tjalURL = "https://www2.tjal.jus.br"

type TJAlagoasDatasource struct{}

func (d TJAlagoasDatasource) buildFirstInstanceURL(filter SearchFilter) string {
	params := map[string]any{
		"conversationId":           "",
		"cbPesquisa":               "NUMPROC",
		"numeroDigitoAnoUnificado": filter.UnifiedYearNumber,
		"foroNumeroUnificado":      filter.UnifiedCourtNumber,
		"dadosConsulta.valorConsultaNuUnificado": []string{
			filter.ProcessNumber,
			"UNIFICADO",
		},
		"dadosConsulta.valorConsulta":  "",
		"dadosConsulta.tipoNuProcesso": "UNIFICADO",
	}

	return addQueryParamsToURL(tjalURL+"/cpopg/search.do", params)
}

func (d TJAlagoasDatasource) buildSecondInstanceSearchURL(filter SearchFilter) string {
	params := map[string]any{
		"conversationId":           "",
		"paginaConsulta":           strconv.Itoa(filter.PageNumber),
		"cbPesquisa":               "NUMPROC",
		"numeroDigitoAnoUnificado": filter.UnifiedYearNumber,
		"foroNumeroUnificado":      filter.UnifiedCourtNumber,
		"dePesquisaNuUnificado": []string{
			filter.ProcessNumber,
			"UNIFICADO",
		},
		"dePesquisa":     "",
		"tipoNuProcesso": "UNIFICADO",
	}

	return addQueryParamsToURL(tjalURL+"/cposg5/search.do", params)
}

func (d TJAlagoasDatasource) buildSecondInstanceShowURL(filter SearchFilter) string {
	return tjalURL + "/cposg5/show.do?processo.codigo=" + filter.ProcessCode
}

func (d TJAlagoasDatasource) SearchFirstInstance(filter SearchFilter) (respBody []byte, err error) {
	var req *http.Request
	req, err = mountRequest(
		d.buildFirstInstanceURL(filter), http.MethodGet,
		"JSESSIONID=AED4E4A2A8603ADEB26440A87551AEA8.cpopg8",
	)
	return doRequest(req)
}

func (d TJAlagoasDatasource) SearchSecondInstance(
	filter SearchFilter,
) (respBody []byte, err error) {
	var (
		req *http.Request
		url string
	)
	if filter.ProcessCode == "" {
		url = d.buildSecondInstanceSearchURL(filter)
	} else {
		url = d.buildSecondInstanceShowURL(filter)
	}
	req, err = mountRequest(
		url,
		http.MethodGet,
		"JSESSIONID=200E02602EA1C65D673C2C33F013AC19.cposg3",
	)
	return doRequest(req)
}
