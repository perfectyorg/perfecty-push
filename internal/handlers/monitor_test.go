package handlers_test

import (
	"github.com/julienschmidt/httprouter"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rwngallego/perfecty-push/internal/handlers"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Monitor", func() {
	Describe("Calling the monitor", func() {

		It("Should return OK", func() {
			req := httptest.NewRequest("GET", "/monitor", nil)
			rr := httptest.NewRecorder()

			router := httprouter.New()
			router.GET("/monitor", handlers.Monitor)
			router.ServeHTTP(rr, req)

			Expect(rr.Code).To(Equal(http.StatusOK))
		})
	})

})
