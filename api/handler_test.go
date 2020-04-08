package api

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Alma-media/taxi/model"
	"github.com/Alma-media/taxi/repository/mock"
)

func Test_NewHandler(t *testing.T) {
	type testCase struct {
		title      string
		repository OrderRepository
		url        string
		code       int
		output     string
	}

	testCases := []testCase{
		{
			title: "test if error is reurned when call to repository.Order() fails",
			url:   "/request/",
			repository: mock.OrderRepository{
				Err: errors.New("order failure"),
			},
			code:   http.StatusInternalServerError,
			output: "order failure\n",
		},
		{
			title: "test getting order from repository with success",
			url:   "/request/",
			repository: mock.OrderRepository{
				OrderResponse: &model.Order{
					ID: "XX",
				},
			},
			code:   http.StatusOK,
			output: "XX\n",
		},
		{
			title: "test if error is reurned when call to repository.List() fails",
			url:   "/admin/requests",
			repository: mock.OrderRepository{
				Err: errors.New("list failure"),
			},
			code:   http.StatusInternalServerError,
			output: "list failure\n",
		},
		{
			title: "test getting the list of orders with success",
			url:   "/admin/requests",
			repository: mock.OrderRepository{
				ListResponse: []*model.Order{
					{
						ID: "XX",
					},
					{
						ID: "YY",
					},
					{
						ID: "ZZ",
					},
				},
			},
			code:   http.StatusOK,
			output: "XX - 0\nYY - 0\nZZ - 0\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			r, err := http.NewRequest(http.MethodGet, tc.url, nil)
			if err != nil {
				t.Fatal(err)
			}

			w := httptest.NewRecorder()

			NewHandler(tc.repository).ServeHTTP(w, r)

			if w.Code != tc.code {
				t.Errorf("response code %d was expected to be %d", w.Code, tc.code)
			}

			actual := w.Body.String()
			if actual != tc.output {
				t.Errorf("response body %q was expected to be %q", actual, tc.output)
			}
		})
	}
}
