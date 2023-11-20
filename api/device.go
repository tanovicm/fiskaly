package api

import (
	"encoding/json"
	"net/http"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/persistence"
)

type ListDevicesResponse struct {
	Devices []persistence.Device `json:"devices"`
}

type GetDeviceResponse struct {
	Device persistence.Device `json:"device"`
}

type CreateSignatureDeviceRequest struct {
	Algorithm string `json:"algorithm"`
	Label     string `json:"label,omitempty"`
}

type CreateSignatureDeviceResponse struct {
	ID string `json:"id"`
}

func (s *Server) ListDevices(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		WriteErrorResponse(response, http.StatusMethodNotAllowed, []string{
			http.StatusText(http.StatusMethodNotAllowed),
		})
		return
	}
	devices, err := s.storage.ListDevices()
	if err != nil {
		WriteErrorResponse(response, http.StatusBadRequest, []string{err.Error()})
		return
	}

	resp := ListDevicesResponse{
		Devices: devices,
	}

	WriteAPIResponse(response, http.StatusOK, resp)
}

func (s *Server) GetDevice(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		WriteErrorResponse(response, http.StatusMethodNotAllowed, []string{
			http.StatusText(http.StatusMethodNotAllowed),
		})
		return
	}
	deviceID := request.URL.Query().Get("id")

	// Check if the parameter exists
	if deviceID == "" {
		WriteErrorResponse(response, http.StatusBadRequest, []string{"DeviceID is missing"})
		return
	}

	device, err := s.storage.GetDevice(deviceID)
	if err != nil {
		WriteErrorResponse(response, http.StatusBadRequest, []string{err.Error()})
		return
	}

	resp := GetDeviceResponse{
		Device: *device,
	}

	WriteAPIResponse(response, http.StatusOK, resp)
}

func (s *Server) CreateSignatureDevice(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		WriteErrorResponse(response, http.StatusMethodNotAllowed, []string{
			http.StatusText(http.StatusMethodNotAllowed),
		})
		return
	}

	var requestBody CreateSignatureDeviceRequest
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		WriteErrorResponse(response, http.StatusBadRequest, []string{"Error parsing request body"})
		return
	}
	defer request.Body.Close()

	device, err := s.storage.CreateDevice(requestBody.Algorithm, requestBody.Label)
	if err != nil {
		WriteErrorResponse(response, http.StatusInternalServerError, []string{"Error storing device"})
		return
	}

	WriteAPIResponse(response, http.StatusOK, CreateSignatureDeviceResponse{ID: device.ID})
}
