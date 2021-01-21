package handlers_test

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rwngallego/perfecty-push/internal/handlers"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("PutPublicUsers", func() {
	Describe("Registering a user", func() {
		It("Should return OK", func() {
			req := httptest.NewRequest("PUT", "/v1/public/users", nil)
			rr := httptest.NewRecorder()
			router := httprouter.New()
			router.PUT("/v1/public/users", handlers.PutPublicUsers)
			router.ServeHTTP(rr, req)

			body := handlers.PutPublicUsersResponse{}
			err := json.NewDecoder(rr.Body).Decode(&body)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(rr.Code).To(Equal(http.StatusOK))
			Expect(body.Success).To(Equal(nil))
			Expect(body.UUID).NotTo(Equal(nil))
		})
	})
})
