package go_kit_rest_api

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func encodePostRegionRequest(ctx context.Context, req *http.Request, request interface{}) error {
	// r.Methods("POST").Path("/regions/")
	req.URL.Path = "/regions/"
	return encodeRequest(ctx, req, request)
}

func decodePostRegionRequest(_ context.Context, resp *http.Response) (interface{}, error) {
	var response postRegionResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

func encodeRequest(_ context.Context, req *http.Request, request interface{}) error {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(request)
	if err != nil {
		return err
	}
	req.Body = ioutil.NopCloser(&buf)
	return nil
}
