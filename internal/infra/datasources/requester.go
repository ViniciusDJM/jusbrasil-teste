package datasources

import (
	"errors"
	"io"
	"net/http"
	"strings"
)

func checkHTMLContentType(res *http.Response, body []byte) (err error) {
	contextType := res.Header.Get("Context-Type")
	contentA, _, _ := strings.Cut(contextType, ";")
	if strings.TrimSpace(contentA) != "text/html" {
		contextType = http.DetectContentType(body)
		contentA, _, _ = strings.Cut(contextType, ";")
		if contentA != "text/html" {
			err = ErrInvalidContentType
		}
	}
	return
}

func checkStatus(status int, oneOf ...int) (err error) {
	for _, toCheck := range oneOf {
		if status == toCheck {
			return
		}
	}

	err = ErrInvalidStatus
	return
}

func mountRequest(url, method, cookie string) (req *http.Request, err error) {
	req, err = http.NewRequest(method, url, nil)
	if cookie != "" {
		req.Header.Add("cookie", cookie)
	}

	return
}

func doRequest(req *http.Request) (respBody []byte, err error) {
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
