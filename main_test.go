package main_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	. "github.com/athornton2012/10x"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var server http.Server
var csvPath string

var _ = Describe("Main", func() {
	JustBeforeEach(func() {
		server, err := SetupServer(csvPath)
		Expect(err).NotTo(HaveOccurred())

		go func() {
			err := server.ListenAndServe()
			Expect(err).NotTo(HaveOccurred())

		}()
	})

	AfterEach(func() {
		server.Shutdown(context.Background())
	})

	Context("happy path", func() {
		BeforeEach(func() {
			csvPath = "./fixtures/weather-short.csv"
		})

		It("Only returns the first n records it finds when limit is set", func() {
			request, _ := http.NewRequest("GET", "http://localhost:8080/query?limit=1&weather=rain", nil)
			request.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			response, err := client.Do(request)
			Expect(err).NotTo(HaveOccurred())

			responseBody, err := io.ReadAll(response.Body)
			Expect(err).NotTo(HaveOccurred())

			var resp []map[string]string
			err = json.Unmarshal(responseBody, &resp)
			Expect(err).NotTo(HaveOccurred())

			Expect(resp).To(Equal([]map[string]string{{"date": "2012-01-02", "precipitation": "10.9", "temp_max": "10.6", "temp_min": "2.8", "wind": "4.5", "weather": "rain"}}))
		})
	})
})
