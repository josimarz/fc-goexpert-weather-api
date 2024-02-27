package weather

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/josimarz/fc-goexpert-weather-api/internal/temp"
)

type MockClient struct{}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	body := io.NopCloser(bytes.NewReader([]byte(`{"current": {"temp_c": 30.5, "temp_f": 581}}`)))
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       body,
	}, nil
}

func TestWeatherApi(t *testing.T) {
	key := "a9b53f9910744477a77220725242602"
	api := NewWeatherApi(key)
	api.client = &MockClient{}
	t.Run("Search", func(t *testing.T) {
		location := "Blumenau, Santa Catarina"
		want := &temp.Temp{
			C: 30.5,
			F: 581,
			K: 303.65,
		}
		got, err := api.Search(location)
		if err != nil || !want.Equals(got) {
			t.Errorf("Search(%v) = (%v, %v), want (%v, %v)", location, got, err, want, nil)
		}
	})
}
