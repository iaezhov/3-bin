package bins

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type JsonBinApi struct {
	apiKey string
	apiUrl string
	client *http.Client
}

type BinResponse struct {
	Record   any `json:"record"`
	Metadata Bin `json:"metadata"`
}

type requestOptions struct {
	method string
	id     string
	body   any
}

func (jba *JsonBinApi) doRequest(opts requestOptions) ([]byte, error) {
	var bodyReader io.Reader
	if opts.body != nil {
		jsonData, err := json.Marshal(opts.body)
		if err != nil {
			return nil, fmt.Errorf("ошибка сериализации: %w", err)
		}
		bodyReader = bytes.NewBuffer(jsonData)
	}

	url := jba.apiUrl
	if opts.id != "" {
		url = fmt.Sprintf("%s/%s", jba.apiUrl, opts.id)
	}

	req, err := http.NewRequest(opts.method, url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("ошибка формирования запроса: %w", err)
	}

	req.Header.Set("X-Master-Key", jba.apiKey)
	if opts.body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := jba.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения ответа: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("статус %d: %s", resp.StatusCode, string(body))
	}

	return body, nil
}

func (jba *JsonBinApi) Create(body any) ([]byte, error) {
	return jba.doRequest(requestOptions{
		method: "POST",
		body:   body,
	})
}

func (jba *JsonBinApi) Get(id string) ([]byte, error) {
	return jba.doRequest(requestOptions{
		method: "GET",
		id:     id,
	})
}

func (jba *JsonBinApi) Delete(id string) ([]byte, error) {
	if id == "" {
		return nil, fmt.Errorf("Отсутствует ID")
	}
	return jba.doRequest(requestOptions{
		method: "DELETE",
		id:     id,
	})
}

func (jba *JsonBinApi) Update(id string, body any) ([]byte, error) {
	if id == "" {
		return nil, fmt.Errorf("Отсутствует ID")
	}
	return jba.doRequest(requestOptions{
		method: "PUT",
		id:     id,
		body:   body,
	})
}

func NewJsonBinApi(apiUrl, apiKey string) *JsonBinApi {
	return &JsonBinApi{
		apiKey: apiKey,
		apiUrl: apiUrl,
		client: &http.Client{},
	}
}
