// bins/api_test.go
package bins_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"hw/3/bins"
)

func setupMockServer() (*httptest.Server, *bins.JsonBinApi) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("X-Master-Key")
		if key != "test-key" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		switch r.Method {
		case "POST":
			body := map[string]any{}
			json.NewDecoder(r.Body).Decode(&body)
			resp := bins.BinResponse{
				Record:   body,
				Metadata: bins.Bin{ID: "new-bin-id", Private: false, Name: "test"},
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)

		case "GET":
			resp := bins.BinResponse{
				Record:   map[string]any{"message": "hello"},
				Metadata: bins.Bin{ID: "existing-id", Private: false, Name: "test"},
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)

		case "PUT":
			body := map[string]any{}
			json.NewDecoder(r.Body).Decode(&body)
			resp := bins.BinResponse{
				Record:   body,
				Metadata: bins.Bin{ID: "updated-id", Private: false, Name: "test"},
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)

		case "DELETE":
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, `{"message":"Bin deleted successfully"}`)
		}
	}))

	api := bins.NewJsonBinApi(server.URL, "test-key")

	return server, api
}

func TestAPICreate(t *testing.T) {
	server, api := setupMockServer()
	defer server.Close()

	data := map[string]any{"field": "value"}
	body, err := api.Create(data)
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}

	var resp bins.BinResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if resp.Metadata.ID != "new-bin-id" {
		t.Errorf("Expected ID 'new-bin-id', got '%s'", resp.Metadata.ID)
	}
}

func TestAPIGet(t *testing.T) {
	server, api := setupMockServer()
	defer server.Close()

	body, err := api.Get("any-id")
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}

	var resp bins.BinResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if resp.Metadata.ID != "existing-id" {
		t.Errorf("Expected ID 'existing-id', got '%s'", resp.Metadata.ID)
	}
}

func TestAPIUpdate(t *testing.T) {
	server, api := setupMockServer()
	defer server.Close()

	data := map[string]any{"updated": "data"}
	body, err := api.Update("some-id", data)
	if err != nil {
		t.Fatalf("Update failed: %v", err)
	}

	var resp bins.BinResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if resp.Metadata.ID != "updated-id" {
		t.Errorf("Expected ID 'updated-id', got '%s'", resp.Metadata.ID)
	}
}

func TestAPIDelete(t *testing.T) {
	server, api := setupMockServer()
	defer server.Close()

	_, err := api.Delete("some-id")
	if err != nil {
		t.Fatalf("Delete failed: %v", err)
	}
}

func TestAPIDelete_MissingID(t *testing.T) {
	server, api := setupMockServer()
	defer server.Close()

	_, err := api.Delete("")
	if err == nil {
		t.Fatal("Expected error for empty ID, got nil")
	}
}
