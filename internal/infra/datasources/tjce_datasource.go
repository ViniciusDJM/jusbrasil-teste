package datasources

import (
	"net/http"
	"strconv"
)

// Constants
const tjceURL = "https://esaj.tjce.jus.br"

// TJCearaDatasource represents a data source for the Ceará Court of Justice (TJCE).
type TJCearaDatasource struct{}

// buildFirstInstanceURL constructs the URL for searching the first instance of a judicial process.
func (d TJCearaDatasource) buildFirstInstanceURL(filter SearchFilter) string {
	params := map[string]any{
		"processo.codigo": "",
		"processo.foro":   filter.UnifiedCourtNumber,
		"processo.numero": filter.ProcessNumber,
	}

	return addQueryParamsToURL(tjceURL+"/cpopg/show.do", params)
}

// buildSecondInstanceSearchURL constructs the URL for searching the second instance of a judicial process.
func (d TJCearaDatasource) buildSecondInstanceSearchURL(filter SearchFilter) string {
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

	return addQueryParamsToURL(tjceURL+"/cposg5/search.do", params)
}

// buildSecondInstanceShowURL constructs the URL for displaying the details of a judicial process in the second instance.
func (d TJCearaDatasource) buildSecondInstanceShowURL(filter SearchFilter) string {
	return tjceURL + "/cposg5/show.do?processo.codigo=" + filter.ProcessCode
}

// SearchFirstInstance performs a search for the first instance of a judicial process using the provided filter.
func (d TJCearaDatasource) SearchFirstInstance(filter SearchFilter) (respBody []byte, err error) {
	var req *http.Request
	req, _ = mountRequest(
		d.buildFirstInstanceURL(filter), http.MethodGet,
		"JSESSIONID=AED4E4A2A8603ADEB26440A87551AEA8.cpopg8",
	)
	return doRequest(req)
}

// SearchSecondInstance performs a search for the second instance of a judicial process using the provided filter.
func (d TJCearaDatasource) SearchSecondInstance(
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
	req, _ = mountRequest(
		url,
		http.MethodGet,
		"JSESSIONID=FF3394BFA7C0949B815725A0D03A1ED9.cposg5",
	)
	return doRequest(req)
}
