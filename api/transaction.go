package api

import (
	"encoding/json"
	"net/http"
)

type SignTransactionRequest struct {
	DeviceID string `json:"device_id"`
	Data     string `json:"data"`
}

type SignTransactionResponse struct {
	Signature  string `json:"signature"`
	SignedData string `json:"signed_data"`
}

func (s *Server) SignTransaction(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		WriteErrorResponse(response, http.StatusMethodNotAllowed, []string{
			http.StatusText(http.StatusMethodNotAllowed),
		})
		return
	}

	var requestBody SignTransactionRequest
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		WriteErrorResponse(response, http.StatusBadRequest, []string{"Error parsing request body"})
		return
	}
	defer request.Body.Close()

	resp := SignTransactionResponse{}

	WriteAPIResponse(response, http.StatusOK, resp)
}
