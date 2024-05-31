package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rodrigoachilles/simple-weather/internal/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockFinder struct {
	mock.Mock
}

func (m *MockFinder) Execute(str string) (interface{}, error) {
	args := m.Called(str)
	if args.Get(0) != nil {
		return args.Get(0).(interface{}), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestLocaleHandler_Handle(t *testing.T) {
	tests := []struct {
		name                string
		cep                 string
		mockLocaleResponse  dto.OutputLocale
		mockLocaleError     error
		mockWeatherResponse dto.OutputWeather
		mockWeatherError    error
		expectedStatusCode  int
		expectedResponse    interface{}
	}{
		{
			name:               "invalid cep length - number < 8",
			cep:                "123",
			expectedStatusCode: http.StatusUnprocessableEntity,
			expectedResponse:   dto.OutputError{StatusCode: http.StatusUnprocessableEntity, Message: "invalid zipcode"},
		},
		{
			name:               "invalid cep length - number > 8",
			cep:                "123456789",
			expectedStatusCode: http.StatusUnprocessableEntity,
			expectedResponse:   dto.OutputError{StatusCode: http.StatusUnprocessableEntity, Message: "invalid zipcode"},
		},
		{
			name:               "locale finder error",
			cep:                "12345678",
			mockLocaleError:    errors.New("locale finder error"),
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   dto.OutputError{StatusCode: http.StatusInternalServerError, Message: "locale finder error"},
		},
		{
			name:               "locale not found",
			cep:                "12345678",
			mockLocaleResponse: dto.OutputLocale{Localidade: ""},
			expectedStatusCode: http.StatusNotFound,
			expectedResponse:   dto.OutputError{StatusCode: http.StatusNotFound, Message: "can not find zipcode"},
		},
		{
			name:               "weather finder error - unauthorized",
			cep:                "12345678",
			mockLocaleResponse: dto.OutputLocale{Localidade: "Localidade"},
			mockWeatherError:   errors.New("API key is invalid or not provided"),
			expectedStatusCode: http.StatusUnauthorized,
			expectedResponse:   dto.OutputError{StatusCode: http.StatusUnauthorized, Message: "API key is invalid or not provided"},
		},
		{
			name:               "weather finder error - internal server error",
			cep:                "12345678",
			mockLocaleResponse: dto.OutputLocale{Localidade: "Localidade"},
			mockWeatherError:   errors.New("weather finder error"),
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   dto.OutputError{StatusCode: http.StatusInternalServerError, Message: "weather finder error"},
		},
		{
			name:                "successful response",
			cep:                 "12345678",
			mockLocaleResponse:  dto.OutputLocale{Localidade: "Localidade"},
			mockWeatherResponse: dto.OutputWeather{Current: dto.CurrentWeather{TempC: 25.0, TempF: 77.0}},
			expectedStatusCode:  http.StatusOK,
			expectedResponse:    dto.OutputResult{Locale: "Localidade", TempC: 25.0, TempF: 77.0, TempK: 298.15},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockLocaleFinder := new(MockFinder)
			mockWeatherFinder := new(MockFinder)
			handler := NewLocaleHandler(mockLocaleFinder, mockWeatherFinder)

			req := httptest.NewRequest("GET", fmt.Sprintf("/locale/%s", tt.cep), nil)
			w := httptest.NewRecorder()

			req = mux.SetURLVars(req, map[string]string{"cep": tt.cep})

			mockLocaleFinder.On("Execute", tt.cep).Return(tt.mockLocaleResponse, tt.mockLocaleError)
			if tt.mockLocaleResponse.Localidade != "" {
				mockWeatherFinder.On("Execute", tt.mockLocaleResponse.Localidade).Return(tt.mockWeatherResponse, tt.mockWeatherError)
			}

			handler.Handle(w, req)

			res := w.Result()
			defer func(Body io.ReadCloser) {
				_ = Body.Close()
			}(res.Body)

			assert.Equal(t, tt.expectedStatusCode, res.StatusCode)

			var actualResponse interface{}
			if res.StatusCode == http.StatusOK {
				var result dto.OutputResult
				err := json.NewDecoder(res.Body).Decode(&result)
				require.NoError(t, err)
				actualResponse = result
			} else {
				var outputError dto.OutputError
				err := json.NewDecoder(res.Body).Decode(&outputError)
				require.NoError(t, err)
				actualResponse = outputError
			}

			assert.Equal(t, tt.expectedResponse, actualResponse)
		})
	}
}
