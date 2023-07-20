package datasources

import (
	"net/http"
	"strconv"
)

// Constants
const tjalURL = "https://www2.tjal.jus.br"

// TJAlagoasDatasource represents a data source for the Alagoas Court of Justice (TJAL).
type TJAlagoasDatasource struct{}

// buildFirstInstanceURL constructs the URL for searching the first instance of a judicial process.
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

// buildSecondInstanceSearchURL constructs the URL for searching the second instance of a judicial process.
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

// buildSecondInstanceShowURL constructs the URL for displaying the details of a judicial process in the second instance.
func (d TJAlagoasDatasource) buildSecondInstanceShowURL(filter SearchFilter) string {
	return tjalURL + "/cposg5/show.do?processo.codigo=" + filter.ProcessCode
}

// SearchFirstInstance performs a search for the first instance of a judicial process using the provided filter.
func (d TJAlagoasDatasource) SearchFirstInstance(filter SearchFilter) (respBody []byte, err error) {
	var req *http.Request
	req, err = mountRequest(
		d.buildFirstInstanceURL(filter), http.MethodGet,
		"JSESSIONID=AED4E4A2A8603ADEB26440A87551AEA8.cpopg8",
	)
	return doRequest(req)
}

// SearchSecondInstance performs a search for the second instance of a judicial process using the provided filter.
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
