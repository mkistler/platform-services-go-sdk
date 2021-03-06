/**
 * (C) Copyright IBM Corp. 2020.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package catalogmanagementv1_test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/platform-services-go-sdk/catalogmanagementv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

var _ = Describe(`CatalogManagementV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(catalogManagementService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(catalogManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
				URL: "https://catalogmanagementv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(catalogManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CATALOG_MANAGEMENT_URL": "https://catalogmanagementv1/api",
				"CATALOG_MANAGEMENT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
				})
				Expect(catalogManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
					URL: "https://testService/api",
				})
				Expect(catalogManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
				})
				err := catalogManagementService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CATALOG_MANAGEMENT_URL": "https://catalogmanagementv1/api",
				"CATALOG_MANAGEMENT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(catalogManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CATALOG_MANAGEMENT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(catalogManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`GetCatalogAccount(getCatalogAccountOptions *GetCatalogAccountOptions) - Operation response error`, func() {
		getCatalogAccountPath := "/catalogaccount"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCatalogAccountPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCatalogAccount with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetCatalogAccountOptions model
				getCatalogAccountOptionsModel := new(catalogmanagementv1.GetCatalogAccountOptions)
				getCatalogAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.GetCatalogAccount(getCatalogAccountOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.GetCatalogAccount(getCatalogAccountOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetCatalogAccount(getCatalogAccountOptions *GetCatalogAccountOptions)`, func() {
		getCatalogAccountPath := "/catalogaccount"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCatalogAccountPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "account_filters": {"include_all": true, "category_filters": {"mapKey": {"include": false, "filter": {"filter_terms": ["FilterTerms"]}}}, "id_filters": {"include": {"filter_terms": ["FilterTerms"]}, "exclude": {"filter_terms": ["FilterTerms"]}}}}`)
				}))
			})
			It(`Invoke GetCatalogAccount successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.GetCatalogAccount(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCatalogAccountOptions model
				getCatalogAccountOptionsModel := new(catalogmanagementv1.GetCatalogAccountOptions)
				getCatalogAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.GetCatalogAccount(getCatalogAccountOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetCatalogAccountWithContext(ctx, getCatalogAccountOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.GetCatalogAccount(getCatalogAccountOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetCatalogAccountWithContext(ctx, getCatalogAccountOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetCatalogAccount with error: Operation request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetCatalogAccountOptions model
				getCatalogAccountOptionsModel := new(catalogmanagementv1.GetCatalogAccountOptions)
				getCatalogAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.GetCatalogAccount(getCatalogAccountOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateCatalogAccount(updateCatalogAccountOptions *UpdateCatalogAccountOptions)`, func() {
		updateCatalogAccountPath := "/catalogaccount"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCatalogAccountPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateCatalogAccount successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := catalogManagementService.UpdateCatalogAccount(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the FilterTerms model
				filterTermsModel := new(catalogmanagementv1.FilterTerms)
				filterTermsModel.FilterTerms = []string{"testString"}

				// Construct an instance of the CategoryFilter model
				categoryFilterModel := new(catalogmanagementv1.CategoryFilter)
				categoryFilterModel.Include = core.BoolPtr(true)
				categoryFilterModel.Filter = filterTermsModel

				// Construct an instance of the IDFilter model
				idFilterModel := new(catalogmanagementv1.IDFilter)
				idFilterModel.Include = filterTermsModel
				idFilterModel.Exclude = filterTermsModel

				// Construct an instance of the Filters model
				filtersModel := new(catalogmanagementv1.Filters)
				filtersModel.IncludeAll = core.BoolPtr(true)
				filtersModel.CategoryFilters = make(map[string]catalogmanagementv1.CategoryFilter)
				filtersModel.IdFilters = idFilterModel
				filtersModel.CategoryFilters["foo"] = *categoryFilterModel

				// Construct an instance of the UpdateCatalogAccountOptions model
				updateCatalogAccountOptionsModel := new(catalogmanagementv1.UpdateCatalogAccountOptions)
				updateCatalogAccountOptionsModel.ID = core.StringPtr("testString")
				updateCatalogAccountOptionsModel.AccountFilters = filtersModel
				updateCatalogAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = catalogManagementService.UpdateCatalogAccount(updateCatalogAccountOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				response, operationErr = catalogManagementService.UpdateCatalogAccount(updateCatalogAccountOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UpdateCatalogAccount with error: Operation request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the FilterTerms model
				filterTermsModel := new(catalogmanagementv1.FilterTerms)
				filterTermsModel.FilterTerms = []string{"testString"}

				// Construct an instance of the CategoryFilter model
				categoryFilterModel := new(catalogmanagementv1.CategoryFilter)
				categoryFilterModel.Include = core.BoolPtr(true)
				categoryFilterModel.Filter = filterTermsModel

				// Construct an instance of the IDFilter model
				idFilterModel := new(catalogmanagementv1.IDFilter)
				idFilterModel.Include = filterTermsModel
				idFilterModel.Exclude = filterTermsModel

				// Construct an instance of the Filters model
				filtersModel := new(catalogmanagementv1.Filters)
				filtersModel.IncludeAll = core.BoolPtr(true)
				filtersModel.CategoryFilters = make(map[string]catalogmanagementv1.CategoryFilter)
				filtersModel.IdFilters = idFilterModel
				filtersModel.CategoryFilters["foo"] = *categoryFilterModel

				// Construct an instance of the UpdateCatalogAccountOptions model
				updateCatalogAccountOptionsModel := new(catalogmanagementv1.UpdateCatalogAccountOptions)
				updateCatalogAccountOptionsModel.ID = core.StringPtr("testString")
				updateCatalogAccountOptionsModel.AccountFilters = filtersModel
				updateCatalogAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := catalogManagementService.UpdateCatalogAccount(updateCatalogAccountOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetCatalogAccountAudit(getCatalogAccountAuditOptions *GetCatalogAccountAuditOptions)`, func() {
		getCatalogAccountAuditPath := "/catalogaccount/audit"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCatalogAccountAuditPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["id"]).To(Equal([]string{"testString"}))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetCatalogAccountAudit successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := catalogManagementService.GetCatalogAccountAudit(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the GetCatalogAccountAuditOptions model
				getCatalogAccountAuditOptionsModel := new(catalogmanagementv1.GetCatalogAccountAuditOptions)
				getCatalogAccountAuditOptionsModel.ID = core.StringPtr("testString")
				getCatalogAccountAuditOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = catalogManagementService.GetCatalogAccountAudit(getCatalogAccountAuditOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				response, operationErr = catalogManagementService.GetCatalogAccountAudit(getCatalogAccountAuditOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke GetCatalogAccountAudit with error: Operation request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetCatalogAccountAuditOptions model
				getCatalogAccountAuditOptionsModel := new(catalogmanagementv1.GetCatalogAccountAuditOptions)
				getCatalogAccountAuditOptionsModel.ID = core.StringPtr("testString")
				getCatalogAccountAuditOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := catalogManagementService.GetCatalogAccountAudit(getCatalogAccountAuditOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCatalogAccountFilters(getCatalogAccountFiltersOptions *GetCatalogAccountFiltersOptions) - Operation response error`, func() {
		getCatalogAccountFiltersPath := "/catalogaccount/filters"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCatalogAccountFiltersPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["catalog"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCatalogAccountFilters with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetCatalogAccountFiltersOptions model
				getCatalogAccountFiltersOptionsModel := new(catalogmanagementv1.GetCatalogAccountFiltersOptions)
				getCatalogAccountFiltersOptionsModel.Catalog = core.StringPtr("testString")
				getCatalogAccountFiltersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.GetCatalogAccountFilters(getCatalogAccountFiltersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.GetCatalogAccountFilters(getCatalogAccountFiltersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetCatalogAccountFilters(getCatalogAccountFiltersOptions *GetCatalogAccountFiltersOptions)`, func() {
		getCatalogAccountFiltersPath := "/catalogaccount/filters"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCatalogAccountFiltersPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["catalog"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"account_filters": [{"include_all": true, "category_filters": {"mapKey": {"include": false, "filter": {"filter_terms": ["FilterTerms"]}}}, "id_filters": {"include": {"filter_terms": ["FilterTerms"]}, "exclude": {"filter_terms": ["FilterTerms"]}}}], "catalog_filters": [{"catalog": {"id": "ID", "name": "Name"}, "filters": {"include_all": true, "category_filters": {"mapKey": {"include": false, "filter": {"filter_terms": ["FilterTerms"]}}}, "id_filters": {"include": {"filter_terms": ["FilterTerms"]}, "exclude": {"filter_terms": ["FilterTerms"]}}}}]}`)
				}))
			})
			It(`Invoke GetCatalogAccountFilters successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.GetCatalogAccountFilters(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCatalogAccountFiltersOptions model
				getCatalogAccountFiltersOptionsModel := new(catalogmanagementv1.GetCatalogAccountFiltersOptions)
				getCatalogAccountFiltersOptionsModel.Catalog = core.StringPtr("testString")
				getCatalogAccountFiltersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.GetCatalogAccountFilters(getCatalogAccountFiltersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetCatalogAccountFiltersWithContext(ctx, getCatalogAccountFiltersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.GetCatalogAccountFilters(getCatalogAccountFiltersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetCatalogAccountFiltersWithContext(ctx, getCatalogAccountFiltersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetCatalogAccountFilters with error: Operation request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetCatalogAccountFiltersOptions model
				getCatalogAccountFiltersOptionsModel := new(catalogmanagementv1.GetCatalogAccountFiltersOptions)
				getCatalogAccountFiltersOptionsModel.Catalog = core.StringPtr("testString")
				getCatalogAccountFiltersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.GetCatalogAccountFilters(getCatalogAccountFiltersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(catalogManagementService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(catalogManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
				URL: "https://catalogmanagementv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(catalogManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CATALOG_MANAGEMENT_URL": "https://catalogmanagementv1/api",
				"CATALOG_MANAGEMENT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
				})
				Expect(catalogManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
					URL: "https://testService/api",
				})
				Expect(catalogManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
				})
				err := catalogManagementService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CATALOG_MANAGEMENT_URL": "https://catalogmanagementv1/api",
				"CATALOG_MANAGEMENT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(catalogManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CATALOG_MANAGEMENT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(catalogManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListCatalogs(listCatalogsOptions *ListCatalogsOptions) - Operation response error`, func() {
		listCatalogsPath := "/catalogs"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCatalogsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListCatalogs with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the ListCatalogsOptions model
				listCatalogsOptionsModel := new(catalogmanagementv1.ListCatalogsOptions)
				listCatalogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.ListCatalogs(listCatalogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.ListCatalogs(listCatalogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListCatalogs(listCatalogsOptions *ListCatalogsOptions)`, func() {
		listCatalogsPath := "/catalogs"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCatalogsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 6, "limit": 5, "total_count": 10, "resource_count": 13, "first": "First", "last": "Last", "prev": "Prev", "next": "Next", "resources": [{"id": "ID", "_rev": "Rev", "label": "Label", "short_description": "ShortDescription", "catalog_icon_url": "CatalogIconURL", "tags": ["Tags"], "url": "URL", "crn": "Crn", "offerings_url": "OfferingsURL", "features": [{"title": "Title", "description": "Description"}], "disabled": true, "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "resource_group_id": "ResourceGroupID", "owning_account": "OwningAccount", "catalog_filters": {"include_all": true, "category_filters": {"mapKey": {"include": false, "filter": {"filter_terms": ["FilterTerms"]}}}, "id_filters": {"include": {"filter_terms": ["FilterTerms"]}, "exclude": {"filter_terms": ["FilterTerms"]}}}, "syndication_settings": {"remove_related_components": false, "clusters": [{"region": "Region", "id": "ID", "name": "Name", "resource_group_name": "ResourceGroupName", "type": "Type", "namespaces": ["Namespaces"], "all_namespaces": false}], "history": {"namespaces": ["Namespaces"], "clusters": [{"region": "Region", "id": "ID", "name": "Name", "resource_group_name": "ResourceGroupName", "type": "Type", "namespaces": ["Namespaces"], "all_namespaces": false}], "last_run": "2019-01-01T12:00:00"}, "authorization": {"token": "Token", "last_run": "2019-01-01T12:00:00"}}}]}`)
				}))
			})
			It(`Invoke ListCatalogs successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.ListCatalogs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListCatalogsOptions model
				listCatalogsOptionsModel := new(catalogmanagementv1.ListCatalogsOptions)
				listCatalogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.ListCatalogs(listCatalogsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.ListCatalogsWithContext(ctx, listCatalogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.ListCatalogs(listCatalogsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.ListCatalogsWithContext(ctx, listCatalogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListCatalogs with error: Operation request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the ListCatalogsOptions model
				listCatalogsOptionsModel := new(catalogmanagementv1.ListCatalogsOptions)
				listCatalogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.ListCatalogs(listCatalogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateCatalog(createCatalogOptions *CreateCatalogOptions) - Operation response error`, func() {
		createCatalogPath := "/catalogs"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCatalogPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateCatalog with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the Feature model
				featureModel := new(catalogmanagementv1.Feature)
				featureModel.Title = core.StringPtr("testString")
				featureModel.Description = core.StringPtr("testString")

				// Construct an instance of the FilterTerms model
				filterTermsModel := new(catalogmanagementv1.FilterTerms)
				filterTermsModel.FilterTerms = []string{"testString"}

				// Construct an instance of the CategoryFilter model
				categoryFilterModel := new(catalogmanagementv1.CategoryFilter)
				categoryFilterModel.Include = core.BoolPtr(true)
				categoryFilterModel.Filter = filterTermsModel

				// Construct an instance of the IDFilter model
				idFilterModel := new(catalogmanagementv1.IDFilter)
				idFilterModel.Include = filterTermsModel
				idFilterModel.Exclude = filterTermsModel

				// Construct an instance of the Filters model
				filtersModel := new(catalogmanagementv1.Filters)
				filtersModel.IncludeAll = core.BoolPtr(true)
				filtersModel.CategoryFilters = make(map[string]catalogmanagementv1.CategoryFilter)
				filtersModel.IdFilters = idFilterModel
				filtersModel.CategoryFilters["foo"] = *categoryFilterModel

				// Construct an instance of the SyndicationCluster model
				syndicationClusterModel := new(catalogmanagementv1.SyndicationCluster)
				syndicationClusterModel.Region = core.StringPtr("testString")
				syndicationClusterModel.ID = core.StringPtr("testString")
				syndicationClusterModel.Name = core.StringPtr("testString")
				syndicationClusterModel.ResourceGroupName = core.StringPtr("testString")
				syndicationClusterModel.Type = core.StringPtr("testString")
				syndicationClusterModel.Namespaces = []string{"testString"}
				syndicationClusterModel.AllNamespaces = core.BoolPtr(true)

				// Construct an instance of the SyndicationHistory model
				syndicationHistoryModel := new(catalogmanagementv1.SyndicationHistory)
				syndicationHistoryModel.Namespaces = []string{"testString"}
				syndicationHistoryModel.Clusters = []catalogmanagementv1.SyndicationCluster{*syndicationClusterModel}
				syndicationHistoryModel.LastRun = CreateMockDateTime()

				// Construct an instance of the SyndicationAuthorization model
				syndicationAuthorizationModel := new(catalogmanagementv1.SyndicationAuthorization)
				syndicationAuthorizationModel.Token = core.StringPtr("testString")
				syndicationAuthorizationModel.LastRun = CreateMockDateTime()

				// Construct an instance of the SyndicationResource model
				syndicationResourceModel := new(catalogmanagementv1.SyndicationResource)
				syndicationResourceModel.RemoveRelatedComponents = core.BoolPtr(true)
				syndicationResourceModel.Clusters = []catalogmanagementv1.SyndicationCluster{*syndicationClusterModel}
				syndicationResourceModel.History = syndicationHistoryModel
				syndicationResourceModel.Authorization = syndicationAuthorizationModel

				// Construct an instance of the CreateCatalogOptions model
				createCatalogOptionsModel := new(catalogmanagementv1.CreateCatalogOptions)
				createCatalogOptionsModel.ID = core.StringPtr("testString")
				createCatalogOptionsModel.Rev = core.StringPtr("testString")
				createCatalogOptionsModel.Label = core.StringPtr("testString")
				createCatalogOptionsModel.ShortDescription = core.StringPtr("testString")
				createCatalogOptionsModel.CatalogIconURL = core.StringPtr("testString")
				createCatalogOptionsModel.Tags = []string{"testString"}
				createCatalogOptionsModel.URL = core.StringPtr("testString")
				createCatalogOptionsModel.Crn = core.StringPtr("testString")
				createCatalogOptionsModel.OfferingsURL = core.StringPtr("testString")
				createCatalogOptionsModel.Features = []catalogmanagementv1.Feature{*featureModel}
				createCatalogOptionsModel.Disabled = core.BoolPtr(true)
				createCatalogOptionsModel.Created = CreateMockDateTime()
				createCatalogOptionsModel.Updated = CreateMockDateTime()
				createCatalogOptionsModel.ResourceGroupID = core.StringPtr("testString")
				createCatalogOptionsModel.OwningAccount = core.StringPtr("testString")
				createCatalogOptionsModel.CatalogFilters = filtersModel
				createCatalogOptionsModel.SyndicationSettings = syndicationResourceModel
				createCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.CreateCatalog(createCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.CreateCatalog(createCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateCatalog(createCatalogOptions *CreateCatalogOptions)`, func() {
		createCatalogPath := "/catalogs"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCatalogPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "_rev": "Rev", "label": "Label", "short_description": "ShortDescription", "catalog_icon_url": "CatalogIconURL", "tags": ["Tags"], "url": "URL", "crn": "Crn", "offerings_url": "OfferingsURL", "features": [{"title": "Title", "description": "Description"}], "disabled": true, "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "resource_group_id": "ResourceGroupID", "owning_account": "OwningAccount", "catalog_filters": {"include_all": true, "category_filters": {"mapKey": {"include": false, "filter": {"filter_terms": ["FilterTerms"]}}}, "id_filters": {"include": {"filter_terms": ["FilterTerms"]}, "exclude": {"filter_terms": ["FilterTerms"]}}}, "syndication_settings": {"remove_related_components": false, "clusters": [{"region": "Region", "id": "ID", "name": "Name", "resource_group_name": "ResourceGroupName", "type": "Type", "namespaces": ["Namespaces"], "all_namespaces": false}], "history": {"namespaces": ["Namespaces"], "clusters": [{"region": "Region", "id": "ID", "name": "Name", "resource_group_name": "ResourceGroupName", "type": "Type", "namespaces": ["Namespaces"], "all_namespaces": false}], "last_run": "2019-01-01T12:00:00"}, "authorization": {"token": "Token", "last_run": "2019-01-01T12:00:00"}}}`)
				}))
			})
			It(`Invoke CreateCatalog successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.CreateCatalog(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Feature model
				featureModel := new(catalogmanagementv1.Feature)
				featureModel.Title = core.StringPtr("testString")
				featureModel.Description = core.StringPtr("testString")

				// Construct an instance of the FilterTerms model
				filterTermsModel := new(catalogmanagementv1.FilterTerms)
				filterTermsModel.FilterTerms = []string{"testString"}

				// Construct an instance of the CategoryFilter model
				categoryFilterModel := new(catalogmanagementv1.CategoryFilter)
				categoryFilterModel.Include = core.BoolPtr(true)
				categoryFilterModel.Filter = filterTermsModel

				// Construct an instance of the IDFilter model
				idFilterModel := new(catalogmanagementv1.IDFilter)
				idFilterModel.Include = filterTermsModel
				idFilterModel.Exclude = filterTermsModel

				// Construct an instance of the Filters model
				filtersModel := new(catalogmanagementv1.Filters)
				filtersModel.IncludeAll = core.BoolPtr(true)
				filtersModel.CategoryFilters = make(map[string]catalogmanagementv1.CategoryFilter)
				filtersModel.IdFilters = idFilterModel
				filtersModel.CategoryFilters["foo"] = *categoryFilterModel

				// Construct an instance of the SyndicationCluster model
				syndicationClusterModel := new(catalogmanagementv1.SyndicationCluster)
				syndicationClusterModel.Region = core.StringPtr("testString")
				syndicationClusterModel.ID = core.StringPtr("testString")
				syndicationClusterModel.Name = core.StringPtr("testString")
				syndicationClusterModel.ResourceGroupName = core.StringPtr("testString")
				syndicationClusterModel.Type = core.StringPtr("testString")
				syndicationClusterModel.Namespaces = []string{"testString"}
				syndicationClusterModel.AllNamespaces = core.BoolPtr(true)

				// Construct an instance of the SyndicationHistory model
				syndicationHistoryModel := new(catalogmanagementv1.SyndicationHistory)
				syndicationHistoryModel.Namespaces = []string{"testString"}
				syndicationHistoryModel.Clusters = []catalogmanagementv1.SyndicationCluster{*syndicationClusterModel}
				syndicationHistoryModel.LastRun = CreateMockDateTime()

				// Construct an instance of the SyndicationAuthorization model
				syndicationAuthorizationModel := new(catalogmanagementv1.SyndicationAuthorization)
				syndicationAuthorizationModel.Token = core.StringPtr("testString")
				syndicationAuthorizationModel.LastRun = CreateMockDateTime()

				// Construct an instance of the SyndicationResource model
				syndicationResourceModel := new(catalogmanagementv1.SyndicationResource)
				syndicationResourceModel.RemoveRelatedComponents = core.BoolPtr(true)
				syndicationResourceModel.Clusters = []catalogmanagementv1.SyndicationCluster{*syndicationClusterModel}
				syndicationResourceModel.History = syndicationHistoryModel
				syndicationResourceModel.Authorization = syndicationAuthorizationModel

				// Construct an instance of the CreateCatalogOptions model
				createCatalogOptionsModel := new(catalogmanagementv1.CreateCatalogOptions)
				createCatalogOptionsModel.ID = core.StringPtr("testString")
				createCatalogOptionsModel.Rev = core.StringPtr("testString")
				createCatalogOptionsModel.Label = core.StringPtr("testString")
				createCatalogOptionsModel.ShortDescription = core.StringPtr("testString")
				createCatalogOptionsModel.CatalogIconURL = core.StringPtr("testString")
				createCatalogOptionsModel.Tags = []string{"testString"}
				createCatalogOptionsModel.URL = core.StringPtr("testString")
				createCatalogOptionsModel.Crn = core.StringPtr("testString")
				createCatalogOptionsModel.OfferingsURL = core.StringPtr("testString")
				createCatalogOptionsModel.Features = []catalogmanagementv1.Feature{*featureModel}
				createCatalogOptionsModel.Disabled = core.BoolPtr(true)
				createCatalogOptionsModel.Created = CreateMockDateTime()
				createCatalogOptionsModel.Updated = CreateMockDateTime()
				createCatalogOptionsModel.ResourceGroupID = core.StringPtr("testString")
				createCatalogOptionsModel.OwningAccount = core.StringPtr("testString")
				createCatalogOptionsModel.CatalogFilters = filtersModel
				createCatalogOptionsModel.SyndicationSettings = syndicationResourceModel
				createCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.CreateCatalog(createCatalogOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.CreateCatalogWithContext(ctx, createCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.CreateCatalog(createCatalogOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.CreateCatalogWithContext(ctx, createCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke CreateCatalog with error: Operation request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the Feature model
				featureModel := new(catalogmanagementv1.Feature)
				featureModel.Title = core.StringPtr("testString")
				featureModel.Description = core.StringPtr("testString")

				// Construct an instance of the FilterTerms model
				filterTermsModel := new(catalogmanagementv1.FilterTerms)
				filterTermsModel.FilterTerms = []string{"testString"}

				// Construct an instance of the CategoryFilter model
				categoryFilterModel := new(catalogmanagementv1.CategoryFilter)
				categoryFilterModel.Include = core.BoolPtr(true)
				categoryFilterModel.Filter = filterTermsModel

				// Construct an instance of the IDFilter model
				idFilterModel := new(catalogmanagementv1.IDFilter)
				idFilterModel.Include = filterTermsModel
				idFilterModel.Exclude = filterTermsModel

				// Construct an instance of the Filters model
				filtersModel := new(catalogmanagementv1.Filters)
				filtersModel.IncludeAll = core.BoolPtr(true)
				filtersModel.CategoryFilters = make(map[string]catalogmanagementv1.CategoryFilter)
				filtersModel.IdFilters = idFilterModel
				filtersModel.CategoryFilters["foo"] = *categoryFilterModel

				// Construct an instance of the SyndicationCluster model
				syndicationClusterModel := new(catalogmanagementv1.SyndicationCluster)
				syndicationClusterModel.Region = core.StringPtr("testString")
				syndicationClusterModel.ID = core.StringPtr("testString")
				syndicationClusterModel.Name = core.StringPtr("testString")
				syndicationClusterModel.ResourceGroupName = core.StringPtr("testString")
				syndicationClusterModel.Type = core.StringPtr("testString")
				syndicationClusterModel.Namespaces = []string{"testString"}
				syndicationClusterModel.AllNamespaces = core.BoolPtr(true)

				// Construct an instance of the SyndicationHistory model
				syndicationHistoryModel := new(catalogmanagementv1.SyndicationHistory)
				syndicationHistoryModel.Namespaces = []string{"testString"}
				syndicationHistoryModel.Clusters = []catalogmanagementv1.SyndicationCluster{*syndicationClusterModel}
				syndicationHistoryModel.LastRun = CreateMockDateTime()

				// Construct an instance of the SyndicationAuthorization model
				syndicationAuthorizationModel := new(catalogmanagementv1.SyndicationAuthorization)
				syndicationAuthorizationModel.Token = core.StringPtr("testString")
				syndicationAuthorizationModel.LastRun = CreateMockDateTime()

				// Construct an instance of the SyndicationResource model
				syndicationResourceModel := new(catalogmanagementv1.SyndicationResource)
				syndicationResourceModel.RemoveRelatedComponents = core.BoolPtr(true)
				syndicationResourceModel.Clusters = []catalogmanagementv1.SyndicationCluster{*syndicationClusterModel}
				syndicationResourceModel.History = syndicationHistoryModel
				syndicationResourceModel.Authorization = syndicationAuthorizationModel

				// Construct an instance of the CreateCatalogOptions model
				createCatalogOptionsModel := new(catalogmanagementv1.CreateCatalogOptions)
				createCatalogOptionsModel.ID = core.StringPtr("testString")
				createCatalogOptionsModel.Rev = core.StringPtr("testString")
				createCatalogOptionsModel.Label = core.StringPtr("testString")
				createCatalogOptionsModel.ShortDescription = core.StringPtr("testString")
				createCatalogOptionsModel.CatalogIconURL = core.StringPtr("testString")
				createCatalogOptionsModel.Tags = []string{"testString"}
				createCatalogOptionsModel.URL = core.StringPtr("testString")
				createCatalogOptionsModel.Crn = core.StringPtr("testString")
				createCatalogOptionsModel.OfferingsURL = core.StringPtr("testString")
				createCatalogOptionsModel.Features = []catalogmanagementv1.Feature{*featureModel}
				createCatalogOptionsModel.Disabled = core.BoolPtr(true)
				createCatalogOptionsModel.Created = CreateMockDateTime()
				createCatalogOptionsModel.Updated = CreateMockDateTime()
				createCatalogOptionsModel.ResourceGroupID = core.StringPtr("testString")
				createCatalogOptionsModel.OwningAccount = core.StringPtr("testString")
				createCatalogOptionsModel.CatalogFilters = filtersModel
				createCatalogOptionsModel.SyndicationSettings = syndicationResourceModel
				createCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.CreateCatalog(createCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCatalog(getCatalogOptions *GetCatalogOptions) - Operation response error`, func() {
		getCatalogPath := "/catalogs/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCatalogPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCatalog with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetCatalogOptions model
				getCatalogOptionsModel := new(catalogmanagementv1.GetCatalogOptions)
				getCatalogOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				getCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.GetCatalog(getCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.GetCatalog(getCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetCatalog(getCatalogOptions *GetCatalogOptions)`, func() {
		getCatalogPath := "/catalogs/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCatalogPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "_rev": "Rev", "label": "Label", "short_description": "ShortDescription", "catalog_icon_url": "CatalogIconURL", "tags": ["Tags"], "url": "URL", "crn": "Crn", "offerings_url": "OfferingsURL", "features": [{"title": "Title", "description": "Description"}], "disabled": true, "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "resource_group_id": "ResourceGroupID", "owning_account": "OwningAccount", "catalog_filters": {"include_all": true, "category_filters": {"mapKey": {"include": false, "filter": {"filter_terms": ["FilterTerms"]}}}, "id_filters": {"include": {"filter_terms": ["FilterTerms"]}, "exclude": {"filter_terms": ["FilterTerms"]}}}, "syndication_settings": {"remove_related_components": false, "clusters": [{"region": "Region", "id": "ID", "name": "Name", "resource_group_name": "ResourceGroupName", "type": "Type", "namespaces": ["Namespaces"], "all_namespaces": false}], "history": {"namespaces": ["Namespaces"], "clusters": [{"region": "Region", "id": "ID", "name": "Name", "resource_group_name": "ResourceGroupName", "type": "Type", "namespaces": ["Namespaces"], "all_namespaces": false}], "last_run": "2019-01-01T12:00:00"}, "authorization": {"token": "Token", "last_run": "2019-01-01T12:00:00"}}}`)
				}))
			})
			It(`Invoke GetCatalog successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.GetCatalog(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCatalogOptions model
				getCatalogOptionsModel := new(catalogmanagementv1.GetCatalogOptions)
				getCatalogOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				getCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.GetCatalog(getCatalogOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetCatalogWithContext(ctx, getCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.GetCatalog(getCatalogOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetCatalogWithContext(ctx, getCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetCatalog with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetCatalogOptions model
				getCatalogOptionsModel := new(catalogmanagementv1.GetCatalogOptions)
				getCatalogOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				getCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.GetCatalog(getCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetCatalogOptions model with no property values
				getCatalogOptionsModelNew := new(catalogmanagementv1.GetCatalogOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = catalogManagementService.GetCatalog(getCatalogOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceCatalog(replaceCatalogOptions *ReplaceCatalogOptions) - Operation response error`, func() {
		replaceCatalogPath := "/catalogs/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceCatalogPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceCatalog with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the Feature model
				featureModel := new(catalogmanagementv1.Feature)
				featureModel.Title = core.StringPtr("testString")
				featureModel.Description = core.StringPtr("testString")

				// Construct an instance of the FilterTerms model
				filterTermsModel := new(catalogmanagementv1.FilterTerms)
				filterTermsModel.FilterTerms = []string{"testString"}

				// Construct an instance of the CategoryFilter model
				categoryFilterModel := new(catalogmanagementv1.CategoryFilter)
				categoryFilterModel.Include = core.BoolPtr(true)
				categoryFilterModel.Filter = filterTermsModel

				// Construct an instance of the IDFilter model
				idFilterModel := new(catalogmanagementv1.IDFilter)
				idFilterModel.Include = filterTermsModel
				idFilterModel.Exclude = filterTermsModel

				// Construct an instance of the Filters model
				filtersModel := new(catalogmanagementv1.Filters)
				filtersModel.IncludeAll = core.BoolPtr(true)
				filtersModel.CategoryFilters = make(map[string]catalogmanagementv1.CategoryFilter)
				filtersModel.IdFilters = idFilterModel
				filtersModel.CategoryFilters["foo"] = *categoryFilterModel

				// Construct an instance of the SyndicationCluster model
				syndicationClusterModel := new(catalogmanagementv1.SyndicationCluster)
				syndicationClusterModel.Region = core.StringPtr("testString")
				syndicationClusterModel.ID = core.StringPtr("testString")
				syndicationClusterModel.Name = core.StringPtr("testString")
				syndicationClusterModel.ResourceGroupName = core.StringPtr("testString")
				syndicationClusterModel.Type = core.StringPtr("testString")
				syndicationClusterModel.Namespaces = []string{"testString"}
				syndicationClusterModel.AllNamespaces = core.BoolPtr(true)

				// Construct an instance of the SyndicationHistory model
				syndicationHistoryModel := new(catalogmanagementv1.SyndicationHistory)
				syndicationHistoryModel.Namespaces = []string{"testString"}
				syndicationHistoryModel.Clusters = []catalogmanagementv1.SyndicationCluster{*syndicationClusterModel}
				syndicationHistoryModel.LastRun = CreateMockDateTime()

				// Construct an instance of the SyndicationAuthorization model
				syndicationAuthorizationModel := new(catalogmanagementv1.SyndicationAuthorization)
				syndicationAuthorizationModel.Token = core.StringPtr("testString")
				syndicationAuthorizationModel.LastRun = CreateMockDateTime()

				// Construct an instance of the SyndicationResource model
				syndicationResourceModel := new(catalogmanagementv1.SyndicationResource)
				syndicationResourceModel.RemoveRelatedComponents = core.BoolPtr(true)
				syndicationResourceModel.Clusters = []catalogmanagementv1.SyndicationCluster{*syndicationClusterModel}
				syndicationResourceModel.History = syndicationHistoryModel
				syndicationResourceModel.Authorization = syndicationAuthorizationModel

				// Construct an instance of the ReplaceCatalogOptions model
				replaceCatalogOptionsModel := new(catalogmanagementv1.ReplaceCatalogOptions)
				replaceCatalogOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				replaceCatalogOptionsModel.ID = core.StringPtr("testString")
				replaceCatalogOptionsModel.Rev = core.StringPtr("testString")
				replaceCatalogOptionsModel.Label = core.StringPtr("testString")
				replaceCatalogOptionsModel.ShortDescription = core.StringPtr("testString")
				replaceCatalogOptionsModel.CatalogIconURL = core.StringPtr("testString")
				replaceCatalogOptionsModel.Tags = []string{"testString"}
				replaceCatalogOptionsModel.URL = core.StringPtr("testString")
				replaceCatalogOptionsModel.Crn = core.StringPtr("testString")
				replaceCatalogOptionsModel.OfferingsURL = core.StringPtr("testString")
				replaceCatalogOptionsModel.Features = []catalogmanagementv1.Feature{*featureModel}
				replaceCatalogOptionsModel.Disabled = core.BoolPtr(true)
				replaceCatalogOptionsModel.Created = CreateMockDateTime()
				replaceCatalogOptionsModel.Updated = CreateMockDateTime()
				replaceCatalogOptionsModel.ResourceGroupID = core.StringPtr("testString")
				replaceCatalogOptionsModel.OwningAccount = core.StringPtr("testString")
				replaceCatalogOptionsModel.CatalogFilters = filtersModel
				replaceCatalogOptionsModel.SyndicationSettings = syndicationResourceModel
				replaceCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.ReplaceCatalog(replaceCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.ReplaceCatalog(replaceCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ReplaceCatalog(replaceCatalogOptions *ReplaceCatalogOptions)`, func() {
		replaceCatalogPath := "/catalogs/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceCatalogPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "_rev": "Rev", "label": "Label", "short_description": "ShortDescription", "catalog_icon_url": "CatalogIconURL", "tags": ["Tags"], "url": "URL", "crn": "Crn", "offerings_url": "OfferingsURL", "features": [{"title": "Title", "description": "Description"}], "disabled": true, "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "resource_group_id": "ResourceGroupID", "owning_account": "OwningAccount", "catalog_filters": {"include_all": true, "category_filters": {"mapKey": {"include": false, "filter": {"filter_terms": ["FilterTerms"]}}}, "id_filters": {"include": {"filter_terms": ["FilterTerms"]}, "exclude": {"filter_terms": ["FilterTerms"]}}}, "syndication_settings": {"remove_related_components": false, "clusters": [{"region": "Region", "id": "ID", "name": "Name", "resource_group_name": "ResourceGroupName", "type": "Type", "namespaces": ["Namespaces"], "all_namespaces": false}], "history": {"namespaces": ["Namespaces"], "clusters": [{"region": "Region", "id": "ID", "name": "Name", "resource_group_name": "ResourceGroupName", "type": "Type", "namespaces": ["Namespaces"], "all_namespaces": false}], "last_run": "2019-01-01T12:00:00"}, "authorization": {"token": "Token", "last_run": "2019-01-01T12:00:00"}}}`)
				}))
			})
			It(`Invoke ReplaceCatalog successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.ReplaceCatalog(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Feature model
				featureModel := new(catalogmanagementv1.Feature)
				featureModel.Title = core.StringPtr("testString")
				featureModel.Description = core.StringPtr("testString")

				// Construct an instance of the FilterTerms model
				filterTermsModel := new(catalogmanagementv1.FilterTerms)
				filterTermsModel.FilterTerms = []string{"testString"}

				// Construct an instance of the CategoryFilter model
				categoryFilterModel := new(catalogmanagementv1.CategoryFilter)
				categoryFilterModel.Include = core.BoolPtr(true)
				categoryFilterModel.Filter = filterTermsModel

				// Construct an instance of the IDFilter model
				idFilterModel := new(catalogmanagementv1.IDFilter)
				idFilterModel.Include = filterTermsModel
				idFilterModel.Exclude = filterTermsModel

				// Construct an instance of the Filters model
				filtersModel := new(catalogmanagementv1.Filters)
				filtersModel.IncludeAll = core.BoolPtr(true)
				filtersModel.CategoryFilters = make(map[string]catalogmanagementv1.CategoryFilter)
				filtersModel.IdFilters = idFilterModel
				filtersModel.CategoryFilters["foo"] = *categoryFilterModel

				// Construct an instance of the SyndicationCluster model
				syndicationClusterModel := new(catalogmanagementv1.SyndicationCluster)
				syndicationClusterModel.Region = core.StringPtr("testString")
				syndicationClusterModel.ID = core.StringPtr("testString")
				syndicationClusterModel.Name = core.StringPtr("testString")
				syndicationClusterModel.ResourceGroupName = core.StringPtr("testString")
				syndicationClusterModel.Type = core.StringPtr("testString")
				syndicationClusterModel.Namespaces = []string{"testString"}
				syndicationClusterModel.AllNamespaces = core.BoolPtr(true)

				// Construct an instance of the SyndicationHistory model
				syndicationHistoryModel := new(catalogmanagementv1.SyndicationHistory)
				syndicationHistoryModel.Namespaces = []string{"testString"}
				syndicationHistoryModel.Clusters = []catalogmanagementv1.SyndicationCluster{*syndicationClusterModel}
				syndicationHistoryModel.LastRun = CreateMockDateTime()

				// Construct an instance of the SyndicationAuthorization model
				syndicationAuthorizationModel := new(catalogmanagementv1.SyndicationAuthorization)
				syndicationAuthorizationModel.Token = core.StringPtr("testString")
				syndicationAuthorizationModel.LastRun = CreateMockDateTime()

				// Construct an instance of the SyndicationResource model
				syndicationResourceModel := new(catalogmanagementv1.SyndicationResource)
				syndicationResourceModel.RemoveRelatedComponents = core.BoolPtr(true)
				syndicationResourceModel.Clusters = []catalogmanagementv1.SyndicationCluster{*syndicationClusterModel}
				syndicationResourceModel.History = syndicationHistoryModel
				syndicationResourceModel.Authorization = syndicationAuthorizationModel

				// Construct an instance of the ReplaceCatalogOptions model
				replaceCatalogOptionsModel := new(catalogmanagementv1.ReplaceCatalogOptions)
				replaceCatalogOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				replaceCatalogOptionsModel.ID = core.StringPtr("testString")
				replaceCatalogOptionsModel.Rev = core.StringPtr("testString")
				replaceCatalogOptionsModel.Label = core.StringPtr("testString")
				replaceCatalogOptionsModel.ShortDescription = core.StringPtr("testString")
				replaceCatalogOptionsModel.CatalogIconURL = core.StringPtr("testString")
				replaceCatalogOptionsModel.Tags = []string{"testString"}
				replaceCatalogOptionsModel.URL = core.StringPtr("testString")
				replaceCatalogOptionsModel.Crn = core.StringPtr("testString")
				replaceCatalogOptionsModel.OfferingsURL = core.StringPtr("testString")
				replaceCatalogOptionsModel.Features = []catalogmanagementv1.Feature{*featureModel}
				replaceCatalogOptionsModel.Disabled = core.BoolPtr(true)
				replaceCatalogOptionsModel.Created = CreateMockDateTime()
				replaceCatalogOptionsModel.Updated = CreateMockDateTime()
				replaceCatalogOptionsModel.ResourceGroupID = core.StringPtr("testString")
				replaceCatalogOptionsModel.OwningAccount = core.StringPtr("testString")
				replaceCatalogOptionsModel.CatalogFilters = filtersModel
				replaceCatalogOptionsModel.SyndicationSettings = syndicationResourceModel
				replaceCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.ReplaceCatalog(replaceCatalogOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.ReplaceCatalogWithContext(ctx, replaceCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.ReplaceCatalog(replaceCatalogOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.ReplaceCatalogWithContext(ctx, replaceCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ReplaceCatalog with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the Feature model
				featureModel := new(catalogmanagementv1.Feature)
				featureModel.Title = core.StringPtr("testString")
				featureModel.Description = core.StringPtr("testString")

				// Construct an instance of the FilterTerms model
				filterTermsModel := new(catalogmanagementv1.FilterTerms)
				filterTermsModel.FilterTerms = []string{"testString"}

				// Construct an instance of the CategoryFilter model
				categoryFilterModel := new(catalogmanagementv1.CategoryFilter)
				categoryFilterModel.Include = core.BoolPtr(true)
				categoryFilterModel.Filter = filterTermsModel

				// Construct an instance of the IDFilter model
				idFilterModel := new(catalogmanagementv1.IDFilter)
				idFilterModel.Include = filterTermsModel
				idFilterModel.Exclude = filterTermsModel

				// Construct an instance of the Filters model
				filtersModel := new(catalogmanagementv1.Filters)
				filtersModel.IncludeAll = core.BoolPtr(true)
				filtersModel.CategoryFilters = make(map[string]catalogmanagementv1.CategoryFilter)
				filtersModel.IdFilters = idFilterModel
				filtersModel.CategoryFilters["foo"] = *categoryFilterModel

				// Construct an instance of the SyndicationCluster model
				syndicationClusterModel := new(catalogmanagementv1.SyndicationCluster)
				syndicationClusterModel.Region = core.StringPtr("testString")
				syndicationClusterModel.ID = core.StringPtr("testString")
				syndicationClusterModel.Name = core.StringPtr("testString")
				syndicationClusterModel.ResourceGroupName = core.StringPtr("testString")
				syndicationClusterModel.Type = core.StringPtr("testString")
				syndicationClusterModel.Namespaces = []string{"testString"}
				syndicationClusterModel.AllNamespaces = core.BoolPtr(true)

				// Construct an instance of the SyndicationHistory model
				syndicationHistoryModel := new(catalogmanagementv1.SyndicationHistory)
				syndicationHistoryModel.Namespaces = []string{"testString"}
				syndicationHistoryModel.Clusters = []catalogmanagementv1.SyndicationCluster{*syndicationClusterModel}
				syndicationHistoryModel.LastRun = CreateMockDateTime()

				// Construct an instance of the SyndicationAuthorization model
				syndicationAuthorizationModel := new(catalogmanagementv1.SyndicationAuthorization)
				syndicationAuthorizationModel.Token = core.StringPtr("testString")
				syndicationAuthorizationModel.LastRun = CreateMockDateTime()

				// Construct an instance of the SyndicationResource model
				syndicationResourceModel := new(catalogmanagementv1.SyndicationResource)
				syndicationResourceModel.RemoveRelatedComponents = core.BoolPtr(true)
				syndicationResourceModel.Clusters = []catalogmanagementv1.SyndicationCluster{*syndicationClusterModel}
				syndicationResourceModel.History = syndicationHistoryModel
				syndicationResourceModel.Authorization = syndicationAuthorizationModel

				// Construct an instance of the ReplaceCatalogOptions model
				replaceCatalogOptionsModel := new(catalogmanagementv1.ReplaceCatalogOptions)
				replaceCatalogOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				replaceCatalogOptionsModel.ID = core.StringPtr("testString")
				replaceCatalogOptionsModel.Rev = core.StringPtr("testString")
				replaceCatalogOptionsModel.Label = core.StringPtr("testString")
				replaceCatalogOptionsModel.ShortDescription = core.StringPtr("testString")
				replaceCatalogOptionsModel.CatalogIconURL = core.StringPtr("testString")
				replaceCatalogOptionsModel.Tags = []string{"testString"}
				replaceCatalogOptionsModel.URL = core.StringPtr("testString")
				replaceCatalogOptionsModel.Crn = core.StringPtr("testString")
				replaceCatalogOptionsModel.OfferingsURL = core.StringPtr("testString")
				replaceCatalogOptionsModel.Features = []catalogmanagementv1.Feature{*featureModel}
				replaceCatalogOptionsModel.Disabled = core.BoolPtr(true)
				replaceCatalogOptionsModel.Created = CreateMockDateTime()
				replaceCatalogOptionsModel.Updated = CreateMockDateTime()
				replaceCatalogOptionsModel.ResourceGroupID = core.StringPtr("testString")
				replaceCatalogOptionsModel.OwningAccount = core.StringPtr("testString")
				replaceCatalogOptionsModel.CatalogFilters = filtersModel
				replaceCatalogOptionsModel.SyndicationSettings = syndicationResourceModel
				replaceCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.ReplaceCatalog(replaceCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceCatalogOptions model with no property values
				replaceCatalogOptionsModelNew := new(catalogmanagementv1.ReplaceCatalogOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = catalogManagementService.ReplaceCatalog(replaceCatalogOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteCatalog(deleteCatalogOptions *DeleteCatalogOptions)`, func() {
		deleteCatalogPath := "/catalogs/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteCatalogPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteCatalog successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := catalogManagementService.DeleteCatalog(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteCatalogOptions model
				deleteCatalogOptionsModel := new(catalogmanagementv1.DeleteCatalogOptions)
				deleteCatalogOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				deleteCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = catalogManagementService.DeleteCatalog(deleteCatalogOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				response, operationErr = catalogManagementService.DeleteCatalog(deleteCatalogOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteCatalog with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the DeleteCatalogOptions model
				deleteCatalogOptionsModel := new(catalogmanagementv1.DeleteCatalogOptions)
				deleteCatalogOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				deleteCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := catalogManagementService.DeleteCatalog(deleteCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteCatalogOptions model with no property values
				deleteCatalogOptionsModelNew := new(catalogmanagementv1.DeleteCatalogOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = catalogManagementService.DeleteCatalog(deleteCatalogOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetCatalogAudit(getCatalogAuditOptions *GetCatalogAuditOptions)`, func() {
		getCatalogAuditPath := "/catalogs/testString/audit"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCatalogAuditPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["id"]).To(Equal([]string{"testString"}))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetCatalogAudit successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := catalogManagementService.GetCatalogAudit(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the GetCatalogAuditOptions model
				getCatalogAuditOptionsModel := new(catalogmanagementv1.GetCatalogAuditOptions)
				getCatalogAuditOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				getCatalogAuditOptionsModel.ID = core.StringPtr("testString")
				getCatalogAuditOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = catalogManagementService.GetCatalogAudit(getCatalogAuditOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				response, operationErr = catalogManagementService.GetCatalogAudit(getCatalogAuditOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke GetCatalogAudit with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetCatalogAuditOptions model
				getCatalogAuditOptionsModel := new(catalogmanagementv1.GetCatalogAuditOptions)
				getCatalogAuditOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				getCatalogAuditOptionsModel.ID = core.StringPtr("testString")
				getCatalogAuditOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := catalogManagementService.GetCatalogAudit(getCatalogAuditOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the GetCatalogAuditOptions model with no property values
				getCatalogAuditOptionsModelNew := new(catalogmanagementv1.GetCatalogAuditOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = catalogManagementService.GetCatalogAudit(getCatalogAuditOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(catalogManagementService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(catalogManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
				URL: "https://catalogmanagementv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(catalogManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CATALOG_MANAGEMENT_URL": "https://catalogmanagementv1/api",
				"CATALOG_MANAGEMENT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
				})
				Expect(catalogManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
					URL: "https://testService/api",
				})
				Expect(catalogManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
				})
				err := catalogManagementService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CATALOG_MANAGEMENT_URL": "https://catalogmanagementv1/api",
				"CATALOG_MANAGEMENT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(catalogManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CATALOG_MANAGEMENT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(catalogManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`GetEnterprise(getEnterpriseOptions *GetEnterpriseOptions) - Operation response error`, func() {
		getEnterprisePath := "/enterprises/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getEnterprisePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetEnterprise with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetEnterpriseOptions model
				getEnterpriseOptionsModel := new(catalogmanagementv1.GetEnterpriseOptions)
				getEnterpriseOptionsModel.EnterpriseID = core.StringPtr("testString")
				getEnterpriseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.GetEnterprise(getEnterpriseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.GetEnterprise(getEnterpriseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetEnterprise(getEnterpriseOptions *GetEnterpriseOptions)`, func() {
		getEnterprisePath := "/enterprises/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getEnterprisePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "_rev": "Rev", "account_filters": {"include_all": true, "category_filters": {"mapKey": {"include": false, "filter": {"filter_terms": ["FilterTerms"]}}}, "id_filters": {"include": {"filter_terms": ["FilterTerms"]}, "exclude": {"filter_terms": ["FilterTerms"]}}}, "account_groups": {"keys": {"id": "ID", "account_filters": {"include_all": true, "category_filters": {"mapKey": {"include": false, "filter": {"filter_terms": ["FilterTerms"]}}}, "id_filters": {"include": {"filter_terms": ["FilterTerms"]}, "exclude": {"filter_terms": ["FilterTerms"]}}}}}}`)
				}))
			})
			It(`Invoke GetEnterprise successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.GetEnterprise(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetEnterpriseOptions model
				getEnterpriseOptionsModel := new(catalogmanagementv1.GetEnterpriseOptions)
				getEnterpriseOptionsModel.EnterpriseID = core.StringPtr("testString")
				getEnterpriseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.GetEnterprise(getEnterpriseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetEnterpriseWithContext(ctx, getEnterpriseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.GetEnterprise(getEnterpriseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetEnterpriseWithContext(ctx, getEnterpriseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetEnterprise with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetEnterpriseOptions model
				getEnterpriseOptionsModel := new(catalogmanagementv1.GetEnterpriseOptions)
				getEnterpriseOptionsModel.EnterpriseID = core.StringPtr("testString")
				getEnterpriseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.GetEnterprise(getEnterpriseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetEnterpriseOptions model with no property values
				getEnterpriseOptionsModelNew := new(catalogmanagementv1.GetEnterpriseOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = catalogManagementService.GetEnterprise(getEnterpriseOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ReplaceEnterprise(replaceEnterpriseOptions *ReplaceEnterpriseOptions)`, func() {
		replaceEnterprisePath := "/enterprises/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceEnterprisePath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					res.WriteHeader(200)
				}))
			})
			It(`Invoke ReplaceEnterprise successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := catalogManagementService.ReplaceEnterprise(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the FilterTerms model
				filterTermsModel := new(catalogmanagementv1.FilterTerms)
				filterTermsModel.FilterTerms = []string{"testString"}

				// Construct an instance of the CategoryFilter model
				categoryFilterModel := new(catalogmanagementv1.CategoryFilter)
				categoryFilterModel.Include = core.BoolPtr(true)
				categoryFilterModel.Filter = filterTermsModel

				// Construct an instance of the IDFilter model
				idFilterModel := new(catalogmanagementv1.IDFilter)
				idFilterModel.Include = filterTermsModel
				idFilterModel.Exclude = filterTermsModel

				// Construct an instance of the Filters model
				filtersModel := new(catalogmanagementv1.Filters)
				filtersModel.IncludeAll = core.BoolPtr(true)
				filtersModel.CategoryFilters = make(map[string]catalogmanagementv1.CategoryFilter)
				filtersModel.IdFilters = idFilterModel
				filtersModel.CategoryFilters["foo"] = *categoryFilterModel

				// Construct an instance of the AccountGroup model
				accountGroupModel := new(catalogmanagementv1.AccountGroup)
				accountGroupModel.ID = core.StringPtr("testString")
				accountGroupModel.AccountFilters = filtersModel

				// Construct an instance of the EnterpriseAccountGroups model
				enterpriseAccountGroupsModel := new(catalogmanagementv1.EnterpriseAccountGroups)
				enterpriseAccountGroupsModel.Keys = accountGroupModel

				// Construct an instance of the ReplaceEnterpriseOptions model
				replaceEnterpriseOptionsModel := new(catalogmanagementv1.ReplaceEnterpriseOptions)
				replaceEnterpriseOptionsModel.EnterpriseID = core.StringPtr("testString")
				replaceEnterpriseOptionsModel.ID = core.StringPtr("testString")
				replaceEnterpriseOptionsModel.Rev = core.StringPtr("testString")
				replaceEnterpriseOptionsModel.AccountFilters = filtersModel
				replaceEnterpriseOptionsModel.AccountGroups = enterpriseAccountGroupsModel
				replaceEnterpriseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = catalogManagementService.ReplaceEnterprise(replaceEnterpriseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				response, operationErr = catalogManagementService.ReplaceEnterprise(replaceEnterpriseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke ReplaceEnterprise with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the FilterTerms model
				filterTermsModel := new(catalogmanagementv1.FilterTerms)
				filterTermsModel.FilterTerms = []string{"testString"}

				// Construct an instance of the CategoryFilter model
				categoryFilterModel := new(catalogmanagementv1.CategoryFilter)
				categoryFilterModel.Include = core.BoolPtr(true)
				categoryFilterModel.Filter = filterTermsModel

				// Construct an instance of the IDFilter model
				idFilterModel := new(catalogmanagementv1.IDFilter)
				idFilterModel.Include = filterTermsModel
				idFilterModel.Exclude = filterTermsModel

				// Construct an instance of the Filters model
				filtersModel := new(catalogmanagementv1.Filters)
				filtersModel.IncludeAll = core.BoolPtr(true)
				filtersModel.CategoryFilters = make(map[string]catalogmanagementv1.CategoryFilter)
				filtersModel.IdFilters = idFilterModel
				filtersModel.CategoryFilters["foo"] = *categoryFilterModel

				// Construct an instance of the AccountGroup model
				accountGroupModel := new(catalogmanagementv1.AccountGroup)
				accountGroupModel.ID = core.StringPtr("testString")
				accountGroupModel.AccountFilters = filtersModel

				// Construct an instance of the EnterpriseAccountGroups model
				enterpriseAccountGroupsModel := new(catalogmanagementv1.EnterpriseAccountGroups)
				enterpriseAccountGroupsModel.Keys = accountGroupModel

				// Construct an instance of the ReplaceEnterpriseOptions model
				replaceEnterpriseOptionsModel := new(catalogmanagementv1.ReplaceEnterpriseOptions)
				replaceEnterpriseOptionsModel.EnterpriseID = core.StringPtr("testString")
				replaceEnterpriseOptionsModel.ID = core.StringPtr("testString")
				replaceEnterpriseOptionsModel.Rev = core.StringPtr("testString")
				replaceEnterpriseOptionsModel.AccountFilters = filtersModel
				replaceEnterpriseOptionsModel.AccountGroups = enterpriseAccountGroupsModel
				replaceEnterpriseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := catalogManagementService.ReplaceEnterprise(replaceEnterpriseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the ReplaceEnterpriseOptions model with no property values
				replaceEnterpriseOptionsModelNew := new(catalogmanagementv1.ReplaceEnterpriseOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = catalogManagementService.ReplaceEnterprise(replaceEnterpriseOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetEnterprisesAudit(getEnterprisesAuditOptions *GetEnterprisesAuditOptions)`, func() {
		getEnterprisesAuditPath := "/enterprises/testString/audit"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getEnterprisesAuditPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["id"]).To(Equal([]string{"testString"}))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetEnterprisesAudit successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := catalogManagementService.GetEnterprisesAudit(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the GetEnterprisesAuditOptions model
				getEnterprisesAuditOptionsModel := new(catalogmanagementv1.GetEnterprisesAuditOptions)
				getEnterprisesAuditOptionsModel.EnterpriseID = core.StringPtr("testString")
				getEnterprisesAuditOptionsModel.ID = core.StringPtr("testString")
				getEnterprisesAuditOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = catalogManagementService.GetEnterprisesAudit(getEnterprisesAuditOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				response, operationErr = catalogManagementService.GetEnterprisesAudit(getEnterprisesAuditOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke GetEnterprisesAudit with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetEnterprisesAuditOptions model
				getEnterprisesAuditOptionsModel := new(catalogmanagementv1.GetEnterprisesAuditOptions)
				getEnterprisesAuditOptionsModel.EnterpriseID = core.StringPtr("testString")
				getEnterprisesAuditOptionsModel.ID = core.StringPtr("testString")
				getEnterprisesAuditOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := catalogManagementService.GetEnterprisesAudit(getEnterprisesAuditOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the GetEnterprisesAuditOptions model with no property values
				getEnterprisesAuditOptionsModelNew := new(catalogmanagementv1.GetEnterprisesAuditOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = catalogManagementService.GetEnterprisesAudit(getEnterprisesAuditOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(catalogManagementService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(catalogManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
				URL: "https://catalogmanagementv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(catalogManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CATALOG_MANAGEMENT_URL": "https://catalogmanagementv1/api",
				"CATALOG_MANAGEMENT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
				})
				Expect(catalogManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
					URL: "https://testService/api",
				})
				Expect(catalogManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
				})
				err := catalogManagementService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CATALOG_MANAGEMENT_URL": "https://catalogmanagementv1/api",
				"CATALOG_MANAGEMENT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(catalogManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CATALOG_MANAGEMENT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(catalogManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`GetConsumptionOfferings(getConsumptionOfferingsOptions *GetConsumptionOfferingsOptions) - Operation response error`, func() {
		getConsumptionOfferingsPath := "/offerings"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getConsumptionOfferingsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for digest query parameter

					Expect(req.URL.Query()["catalog"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["select"]).To(Equal([]string{"all"}))


					// TODO: Add check for includeHidden query parameter

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetConsumptionOfferings with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetConsumptionOfferingsOptions model
				getConsumptionOfferingsOptionsModel := new(catalogmanagementv1.GetConsumptionOfferingsOptions)
				getConsumptionOfferingsOptionsModel.Digest = core.BoolPtr(true)
				getConsumptionOfferingsOptionsModel.Catalog = core.StringPtr("testString")
				getConsumptionOfferingsOptionsModel.Select = core.StringPtr("all")
				getConsumptionOfferingsOptionsModel.IncludeHidden = core.BoolPtr(true)
				getConsumptionOfferingsOptionsModel.Limit = core.Int64Ptr(int64(38))
				getConsumptionOfferingsOptionsModel.Offset = core.Int64Ptr(int64(38))
				getConsumptionOfferingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.GetConsumptionOfferings(getConsumptionOfferingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.GetConsumptionOfferings(getConsumptionOfferingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetConsumptionOfferings(getConsumptionOfferingsOptions *GetConsumptionOfferingsOptions)`, func() {
		getConsumptionOfferingsPath := "/offerings"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getConsumptionOfferingsPath))
					Expect(req.Method).To(Equal("GET"))


					// TODO: Add check for digest query parameter

					Expect(req.URL.Query()["catalog"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["select"]).To(Equal([]string{"all"}))


					// TODO: Add check for includeHidden query parameter

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 6, "limit": 5, "total_count": 10, "resource_count": 13, "first": "First", "last": "Last", "prev": "Prev", "next": "Next", "resources": [{"id": "ID", "_rev": "Rev", "url": "URL", "crn": "Crn", "label": "Label", "name": "Name", "offering_icon_url": "OfferingIconURL", "offering_docs_url": "OfferingDocsURL", "offering_support_url": "OfferingSupportURL", "tags": ["Tags"], "rating": {"one_star_count": 12, "two_star_count": 12, "three_star_count": 14, "four_star_count": 13}, "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "short_description": "ShortDescription", "long_description": "LongDescription", "features": [{"title": "Title", "description": "Description"}], "kinds": [{"id": "ID", "format_kind": "FormatKind", "target_kind": "TargetKind", "metadata": {"anyKey": "anyValue"}, "install_description": "InstallDescription", "tags": ["Tags"], "additional_features": [{"title": "Title", "description": "Description"}], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "versions": [{"id": "ID", "_rev": "Rev", "crn": "Crn", "version": "Version", "sha": "Sha", "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "offering_id": "OfferingID", "catalog_id": "CatalogID", "kind_id": "KindID", "tags": ["Tags"], "repo_url": "RepoURL", "source_url": "SourceURL", "tgz_url": "TgzURL", "configuration": [{"key": "Key", "type": "Type", "default_value": "anyValue", "value_constraint": "ValueConstraint", "description": "Description", "required": true, "options": [{"anyKey": "anyValue"}], "hidden": true}], "metadata": {"anyKey": "anyValue"}, "validation": {"validated": "2019-01-01T12:00:00", "requested": "2019-01-01T12:00:00", "state": "State", "last_operation": "LastOperation", "target": {"anyKey": "anyValue"}}, "required_resources": [{"type": "mem", "value": "anyValue"}], "single_instance": true, "install": {"instructions": "Instructions", "script": "Script", "script_permission": "ScriptPermission", "delete_script": "DeleteScript", "scope": "Scope"}, "pre_install": [{"instructions": "Instructions", "script": "Script", "script_permission": "ScriptPermission", "delete_script": "DeleteScript", "scope": "Scope"}], "entitlement": {"provider_name": "ProviderName", "provider_id": "ProviderID", "product_id": "ProductID", "part_numbers": ["PartNumbers"], "image_repo_name": "ImageRepoName"}, "licenses": [{"id": "ID", "name": "Name", "type": "Type", "url": "URL", "description": "Description"}], "image_manifest_url": "ImageManifestURL", "deprecated": true, "package_version": "PackageVersion", "state": {"current": "Current", "current_entered": "2019-01-01T12:00:00", "pending": "Pending", "pending_requested": "2019-01-01T12:00:00", "previous": "Previous"}, "version_locator": "VersionLocator", "console_url": "ConsoleURL", "long_description": "LongDescription", "whitelisted_accounts": ["WhitelistedAccounts"]}], "plans": [{"id": "ID", "label": "Label", "name": "Name", "short_description": "ShortDescription", "long_description": "LongDescription", "metadata": {"anyKey": "anyValue"}, "tags": ["Tags"], "additional_features": [{"title": "Title", "description": "Description"}], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "deployments": [{"id": "ID", "label": "Label", "name": "Name", "short_description": "ShortDescription", "long_description": "LongDescription", "metadata": {"anyKey": "anyValue"}, "tags": ["Tags"], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00"}]}]}], "permit_request_ibm_public_publish": false, "ibm_publish_approved": true, "public_publish_approved": false, "public_original_crn": "PublicOriginalCrn", "publish_public_crn": "PublishPublicCrn", "portal_approval_record": "PortalApprovalRecord", "portal_ui_url": "PortalUiURL", "catalog_id": "CatalogID", "catalog_name": "CatalogName", "metadata": {"anyKey": "anyValue"}, "disclaimer": "Disclaimer", "hidden": true, "provider": "Provider", "repo_info": {"token": "Token", "type": "Type"}}]}`)
				}))
			})
			It(`Invoke GetConsumptionOfferings successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.GetConsumptionOfferings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetConsumptionOfferingsOptions model
				getConsumptionOfferingsOptionsModel := new(catalogmanagementv1.GetConsumptionOfferingsOptions)
				getConsumptionOfferingsOptionsModel.Digest = core.BoolPtr(true)
				getConsumptionOfferingsOptionsModel.Catalog = core.StringPtr("testString")
				getConsumptionOfferingsOptionsModel.Select = core.StringPtr("all")
				getConsumptionOfferingsOptionsModel.IncludeHidden = core.BoolPtr(true)
				getConsumptionOfferingsOptionsModel.Limit = core.Int64Ptr(int64(38))
				getConsumptionOfferingsOptionsModel.Offset = core.Int64Ptr(int64(38))
				getConsumptionOfferingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.GetConsumptionOfferings(getConsumptionOfferingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetConsumptionOfferingsWithContext(ctx, getConsumptionOfferingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.GetConsumptionOfferings(getConsumptionOfferingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetConsumptionOfferingsWithContext(ctx, getConsumptionOfferingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetConsumptionOfferings with error: Operation request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetConsumptionOfferingsOptions model
				getConsumptionOfferingsOptionsModel := new(catalogmanagementv1.GetConsumptionOfferingsOptions)
				getConsumptionOfferingsOptionsModel.Digest = core.BoolPtr(true)
				getConsumptionOfferingsOptionsModel.Catalog = core.StringPtr("testString")
				getConsumptionOfferingsOptionsModel.Select = core.StringPtr("all")
				getConsumptionOfferingsOptionsModel.IncludeHidden = core.BoolPtr(true)
				getConsumptionOfferingsOptionsModel.Limit = core.Int64Ptr(int64(38))
				getConsumptionOfferingsOptionsModel.Offset = core.Int64Ptr(int64(38))
				getConsumptionOfferingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.GetConsumptionOfferings(getConsumptionOfferingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListOfferings(listOfferingsOptions *ListOfferingsOptions) - Operation response error`, func() {
		listOfferingsPath := "/catalogs/testString/offerings"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listOfferingsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for digest query parameter

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListOfferings with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the ListOfferingsOptions model
				listOfferingsOptionsModel := new(catalogmanagementv1.ListOfferingsOptions)
				listOfferingsOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				listOfferingsOptionsModel.Digest = core.BoolPtr(true)
				listOfferingsOptionsModel.Limit = core.Int64Ptr(int64(38))
				listOfferingsOptionsModel.Offset = core.Int64Ptr(int64(38))
				listOfferingsOptionsModel.Name = core.StringPtr("testString")
				listOfferingsOptionsModel.Sort = core.StringPtr("testString")
				listOfferingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.ListOfferings(listOfferingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.ListOfferings(listOfferingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListOfferings(listOfferingsOptions *ListOfferingsOptions)`, func() {
		listOfferingsPath := "/catalogs/testString/offerings"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listOfferingsPath))
					Expect(req.Method).To(Equal("GET"))


					// TODO: Add check for digest query parameter

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 6, "limit": 5, "total_count": 10, "resource_count": 13, "first": "First", "last": "Last", "prev": "Prev", "next": "Next", "resources": [{"id": "ID", "_rev": "Rev", "url": "URL", "crn": "Crn", "label": "Label", "name": "Name", "offering_icon_url": "OfferingIconURL", "offering_docs_url": "OfferingDocsURL", "offering_support_url": "OfferingSupportURL", "tags": ["Tags"], "rating": {"one_star_count": 12, "two_star_count": 12, "three_star_count": 14, "four_star_count": 13}, "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "short_description": "ShortDescription", "long_description": "LongDescription", "features": [{"title": "Title", "description": "Description"}], "kinds": [{"id": "ID", "format_kind": "FormatKind", "target_kind": "TargetKind", "metadata": {"anyKey": "anyValue"}, "install_description": "InstallDescription", "tags": ["Tags"], "additional_features": [{"title": "Title", "description": "Description"}], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "versions": [{"id": "ID", "_rev": "Rev", "crn": "Crn", "version": "Version", "sha": "Sha", "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "offering_id": "OfferingID", "catalog_id": "CatalogID", "kind_id": "KindID", "tags": ["Tags"], "repo_url": "RepoURL", "source_url": "SourceURL", "tgz_url": "TgzURL", "configuration": [{"key": "Key", "type": "Type", "default_value": "anyValue", "value_constraint": "ValueConstraint", "description": "Description", "required": true, "options": [{"anyKey": "anyValue"}], "hidden": true}], "metadata": {"anyKey": "anyValue"}, "validation": {"validated": "2019-01-01T12:00:00", "requested": "2019-01-01T12:00:00", "state": "State", "last_operation": "LastOperation", "target": {"anyKey": "anyValue"}}, "required_resources": [{"type": "mem", "value": "anyValue"}], "single_instance": true, "install": {"instructions": "Instructions", "script": "Script", "script_permission": "ScriptPermission", "delete_script": "DeleteScript", "scope": "Scope"}, "pre_install": [{"instructions": "Instructions", "script": "Script", "script_permission": "ScriptPermission", "delete_script": "DeleteScript", "scope": "Scope"}], "entitlement": {"provider_name": "ProviderName", "provider_id": "ProviderID", "product_id": "ProductID", "part_numbers": ["PartNumbers"], "image_repo_name": "ImageRepoName"}, "licenses": [{"id": "ID", "name": "Name", "type": "Type", "url": "URL", "description": "Description"}], "image_manifest_url": "ImageManifestURL", "deprecated": true, "package_version": "PackageVersion", "state": {"current": "Current", "current_entered": "2019-01-01T12:00:00", "pending": "Pending", "pending_requested": "2019-01-01T12:00:00", "previous": "Previous"}, "version_locator": "VersionLocator", "console_url": "ConsoleURL", "long_description": "LongDescription", "whitelisted_accounts": ["WhitelistedAccounts"]}], "plans": [{"id": "ID", "label": "Label", "name": "Name", "short_description": "ShortDescription", "long_description": "LongDescription", "metadata": {"anyKey": "anyValue"}, "tags": ["Tags"], "additional_features": [{"title": "Title", "description": "Description"}], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "deployments": [{"id": "ID", "label": "Label", "name": "Name", "short_description": "ShortDescription", "long_description": "LongDescription", "metadata": {"anyKey": "anyValue"}, "tags": ["Tags"], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00"}]}]}], "permit_request_ibm_public_publish": false, "ibm_publish_approved": true, "public_publish_approved": false, "public_original_crn": "PublicOriginalCrn", "publish_public_crn": "PublishPublicCrn", "portal_approval_record": "PortalApprovalRecord", "portal_ui_url": "PortalUiURL", "catalog_id": "CatalogID", "catalog_name": "CatalogName", "metadata": {"anyKey": "anyValue"}, "disclaimer": "Disclaimer", "hidden": true, "provider": "Provider", "repo_info": {"token": "Token", "type": "Type"}}]}`)
				}))
			})
			It(`Invoke ListOfferings successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.ListOfferings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListOfferingsOptions model
				listOfferingsOptionsModel := new(catalogmanagementv1.ListOfferingsOptions)
				listOfferingsOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				listOfferingsOptionsModel.Digest = core.BoolPtr(true)
				listOfferingsOptionsModel.Limit = core.Int64Ptr(int64(38))
				listOfferingsOptionsModel.Offset = core.Int64Ptr(int64(38))
				listOfferingsOptionsModel.Name = core.StringPtr("testString")
				listOfferingsOptionsModel.Sort = core.StringPtr("testString")
				listOfferingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.ListOfferings(listOfferingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.ListOfferingsWithContext(ctx, listOfferingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.ListOfferings(listOfferingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.ListOfferingsWithContext(ctx, listOfferingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListOfferings with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the ListOfferingsOptions model
				listOfferingsOptionsModel := new(catalogmanagementv1.ListOfferingsOptions)
				listOfferingsOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				listOfferingsOptionsModel.Digest = core.BoolPtr(true)
				listOfferingsOptionsModel.Limit = core.Int64Ptr(int64(38))
				listOfferingsOptionsModel.Offset = core.Int64Ptr(int64(38))
				listOfferingsOptionsModel.Name = core.StringPtr("testString")
				listOfferingsOptionsModel.Sort = core.StringPtr("testString")
				listOfferingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.ListOfferings(listOfferingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListOfferingsOptions model with no property values
				listOfferingsOptionsModelNew := new(catalogmanagementv1.ListOfferingsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = catalogManagementService.ListOfferings(listOfferingsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateOffering(createOfferingOptions *CreateOfferingOptions) - Operation response error`, func() {
		createOfferingPath := "/catalogs/testString/offerings"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createOfferingPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateOffering with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the Rating model
				ratingModel := new(catalogmanagementv1.Rating)
				ratingModel.OneStarCount = core.Int64Ptr(int64(38))
				ratingModel.TwoStarCount = core.Int64Ptr(int64(38))
				ratingModel.ThreeStarCount = core.Int64Ptr(int64(38))
				ratingModel.FourStarCount = core.Int64Ptr(int64(38))

				// Construct an instance of the Feature model
				featureModel := new(catalogmanagementv1.Feature)
				featureModel.Title = core.StringPtr("testString")
				featureModel.Description = core.StringPtr("testString")

				// Construct an instance of the Configuration model
				configurationModel := new(catalogmanagementv1.Configuration)
				configurationModel.Key = core.StringPtr("testString")
				configurationModel.Type = core.StringPtr("testString")
				configurationModel.DefaultValue = core.StringPtr("testString")
				configurationModel.ValueConstraint = core.StringPtr("testString")
				configurationModel.Description = core.StringPtr("testString")
				configurationModel.Required = core.BoolPtr(true)
				configurationModel.Options = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				configurationModel.Hidden = core.BoolPtr(true)

				// Construct an instance of the Validation model
				validationModel := new(catalogmanagementv1.Validation)
				validationModel.Validated = CreateMockDateTime()
				validationModel.Requested = CreateMockDateTime()
				validationModel.State = core.StringPtr("testString")
				validationModel.LastOperation = core.StringPtr("testString")
				validationModel.Target = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the Resource model
				resourceModel := new(catalogmanagementv1.Resource)
				resourceModel.Type = core.StringPtr("mem")
				resourceModel.Value = core.StringPtr("testString")

				// Construct an instance of the Script model
				scriptModel := new(catalogmanagementv1.Script)
				scriptModel.Instructions = core.StringPtr("testString")
				scriptModel.Script = core.StringPtr("testString")
				scriptModel.ScriptPermission = core.StringPtr("testString")
				scriptModel.DeleteScript = core.StringPtr("testString")
				scriptModel.Scope = core.StringPtr("testString")

				// Construct an instance of the VersionEntitlement model
				versionEntitlementModel := new(catalogmanagementv1.VersionEntitlement)
				versionEntitlementModel.ProviderName = core.StringPtr("testString")
				versionEntitlementModel.ProviderID = core.StringPtr("testString")
				versionEntitlementModel.ProductID = core.StringPtr("testString")
				versionEntitlementModel.PartNumbers = []string{"testString"}
				versionEntitlementModel.ImageRepoName = core.StringPtr("testString")

				// Construct an instance of the License model
				licenseModel := new(catalogmanagementv1.License)
				licenseModel.ID = core.StringPtr("testString")
				licenseModel.Name = core.StringPtr("testString")
				licenseModel.Type = core.StringPtr("testString")
				licenseModel.URL = core.StringPtr("testString")
				licenseModel.Description = core.StringPtr("testString")

				// Construct an instance of the State model
				stateModel := new(catalogmanagementv1.State)
				stateModel.Current = core.StringPtr("testString")
				stateModel.CurrentEntered = CreateMockDateTime()
				stateModel.Pending = core.StringPtr("testString")
				stateModel.PendingRequested = CreateMockDateTime()
				stateModel.Previous = core.StringPtr("testString")

				// Construct an instance of the Version model
				versionModel := new(catalogmanagementv1.Version)
				versionModel.ID = core.StringPtr("testString")
				versionModel.Rev = core.StringPtr("testString")
				versionModel.Crn = core.StringPtr("testString")
				versionModel.Version = core.StringPtr("testString")
				versionModel.Sha = core.StringPtr("testString")
				versionModel.Created = CreateMockDateTime()
				versionModel.Updated = CreateMockDateTime()
				versionModel.OfferingID = core.StringPtr("testString")
				versionModel.CatalogID = core.StringPtr("testString")
				versionModel.KindID = core.StringPtr("testString")
				versionModel.Tags = []string{"testString"}
				versionModel.RepoURL = core.StringPtr("testString")
				versionModel.SourceURL = core.StringPtr("testString")
				versionModel.TgzURL = core.StringPtr("testString")
				versionModel.Configuration = []catalogmanagementv1.Configuration{*configurationModel}
				versionModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}
				versionModel.Validation = validationModel
				versionModel.RequiredResources = []catalogmanagementv1.Resource{*resourceModel}
				versionModel.SingleInstance = core.BoolPtr(true)
				versionModel.Install = scriptModel
				versionModel.PreInstall = []catalogmanagementv1.Script{*scriptModel}
				versionModel.Entitlement = versionEntitlementModel
				versionModel.Licenses = []catalogmanagementv1.License{*licenseModel}
				versionModel.ImageManifestURL = core.StringPtr("testString")
				versionModel.Deprecated = core.BoolPtr(true)
				versionModel.PackageVersion = core.StringPtr("testString")
				versionModel.State = stateModel
				versionModel.VersionLocator = core.StringPtr("testString")
				versionModel.ConsoleURL = core.StringPtr("testString")
				versionModel.LongDescription = core.StringPtr("testString")
				versionModel.WhitelistedAccounts = []string{"testString"}

				// Construct an instance of the Deployment model
				deploymentModel := new(catalogmanagementv1.Deployment)
				deploymentModel.ID = core.StringPtr("testString")
				deploymentModel.Label = core.StringPtr("testString")
				deploymentModel.Name = core.StringPtr("testString")
				deploymentModel.ShortDescription = core.StringPtr("testString")
				deploymentModel.LongDescription = core.StringPtr("testString")
				deploymentModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}
				deploymentModel.Tags = []string{"testString"}
				deploymentModel.Created = CreateMockDateTime()
				deploymentModel.Updated = CreateMockDateTime()

				// Construct an instance of the Plan model
				planModel := new(catalogmanagementv1.Plan)
				planModel.ID = core.StringPtr("testString")
				planModel.Label = core.StringPtr("testString")
				planModel.Name = core.StringPtr("testString")
				planModel.ShortDescription = core.StringPtr("testString")
				planModel.LongDescription = core.StringPtr("testString")
				planModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}
				planModel.Tags = []string{"testString"}
				planModel.AdditionalFeatures = []catalogmanagementv1.Feature{*featureModel}
				planModel.Created = CreateMockDateTime()
				planModel.Updated = CreateMockDateTime()
				planModel.Deployments = []catalogmanagementv1.Deployment{*deploymentModel}

				// Construct an instance of the Kind model
				kindModel := new(catalogmanagementv1.Kind)
				kindModel.ID = core.StringPtr("testString")
				kindModel.FormatKind = core.StringPtr("testString")
				kindModel.TargetKind = core.StringPtr("testString")
				kindModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}
				kindModel.InstallDescription = core.StringPtr("testString")
				kindModel.Tags = []string{"testString"}
				kindModel.AdditionalFeatures = []catalogmanagementv1.Feature{*featureModel}
				kindModel.Created = CreateMockDateTime()
				kindModel.Updated = CreateMockDateTime()
				kindModel.Versions = []catalogmanagementv1.Version{*versionModel}
				kindModel.Plans = []catalogmanagementv1.Plan{*planModel}

				// Construct an instance of the RepoInfo model
				repoInfoModel := new(catalogmanagementv1.RepoInfo)
				repoInfoModel.Token = core.StringPtr("testString")
				repoInfoModel.Type = core.StringPtr("testString")

				// Construct an instance of the CreateOfferingOptions model
				createOfferingOptionsModel := new(catalogmanagementv1.CreateOfferingOptions)
				createOfferingOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				createOfferingOptionsModel.ID = core.StringPtr("testString")
				createOfferingOptionsModel.Rev = core.StringPtr("testString")
				createOfferingOptionsModel.URL = core.StringPtr("testString")
				createOfferingOptionsModel.Crn = core.StringPtr("testString")
				createOfferingOptionsModel.Label = core.StringPtr("testString")
				createOfferingOptionsModel.Name = core.StringPtr("testString")
				createOfferingOptionsModel.OfferingIconURL = core.StringPtr("testString")
				createOfferingOptionsModel.OfferingDocsURL = core.StringPtr("testString")
				createOfferingOptionsModel.OfferingSupportURL = core.StringPtr("testString")
				createOfferingOptionsModel.Tags = []string{"testString"}
				createOfferingOptionsModel.Rating = ratingModel
				createOfferingOptionsModel.Created = CreateMockDateTime()
				createOfferingOptionsModel.Updated = CreateMockDateTime()
				createOfferingOptionsModel.ShortDescription = core.StringPtr("testString")
				createOfferingOptionsModel.LongDescription = core.StringPtr("testString")
				createOfferingOptionsModel.Features = []catalogmanagementv1.Feature{*featureModel}
				createOfferingOptionsModel.Kinds = []catalogmanagementv1.Kind{*kindModel}
				createOfferingOptionsModel.PermitRequestIbmPublicPublish = core.BoolPtr(true)
				createOfferingOptionsModel.IbmPublishApproved = core.BoolPtr(true)
				createOfferingOptionsModel.PublicPublishApproved = core.BoolPtr(true)
				createOfferingOptionsModel.PublicOriginalCrn = core.StringPtr("testString")
				createOfferingOptionsModel.PublishPublicCrn = core.StringPtr("testString")
				createOfferingOptionsModel.PortalApprovalRecord = core.StringPtr("testString")
				createOfferingOptionsModel.PortalUiURL = core.StringPtr("testString")
				createOfferingOptionsModel.CatalogID = core.StringPtr("testString")
				createOfferingOptionsModel.CatalogName = core.StringPtr("testString")
				createOfferingOptionsModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}
				createOfferingOptionsModel.Disclaimer = core.StringPtr("testString")
				createOfferingOptionsModel.Hidden = core.BoolPtr(true)
				createOfferingOptionsModel.Provider = core.StringPtr("testString")
				createOfferingOptionsModel.RepoInfo = repoInfoModel
				createOfferingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.CreateOffering(createOfferingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.CreateOffering(createOfferingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateOffering(createOfferingOptions *CreateOfferingOptions)`, func() {
		createOfferingPath := "/catalogs/testString/offerings"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createOfferingPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "_rev": "Rev", "url": "URL", "crn": "Crn", "label": "Label", "name": "Name", "offering_icon_url": "OfferingIconURL", "offering_docs_url": "OfferingDocsURL", "offering_support_url": "OfferingSupportURL", "tags": ["Tags"], "rating": {"one_star_count": 12, "two_star_count": 12, "three_star_count": 14, "four_star_count": 13}, "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "short_description": "ShortDescription", "long_description": "LongDescription", "features": [{"title": "Title", "description": "Description"}], "kinds": [{"id": "ID", "format_kind": "FormatKind", "target_kind": "TargetKind", "metadata": {"anyKey": "anyValue"}, "install_description": "InstallDescription", "tags": ["Tags"], "additional_features": [{"title": "Title", "description": "Description"}], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "versions": [{"id": "ID", "_rev": "Rev", "crn": "Crn", "version": "Version", "sha": "Sha", "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "offering_id": "OfferingID", "catalog_id": "CatalogID", "kind_id": "KindID", "tags": ["Tags"], "repo_url": "RepoURL", "source_url": "SourceURL", "tgz_url": "TgzURL", "configuration": [{"key": "Key", "type": "Type", "default_value": "anyValue", "value_constraint": "ValueConstraint", "description": "Description", "required": true, "options": [{"anyKey": "anyValue"}], "hidden": true}], "metadata": {"anyKey": "anyValue"}, "validation": {"validated": "2019-01-01T12:00:00", "requested": "2019-01-01T12:00:00", "state": "State", "last_operation": "LastOperation", "target": {"anyKey": "anyValue"}}, "required_resources": [{"type": "mem", "value": "anyValue"}], "single_instance": true, "install": {"instructions": "Instructions", "script": "Script", "script_permission": "ScriptPermission", "delete_script": "DeleteScript", "scope": "Scope"}, "pre_install": [{"instructions": "Instructions", "script": "Script", "script_permission": "ScriptPermission", "delete_script": "DeleteScript", "scope": "Scope"}], "entitlement": {"provider_name": "ProviderName", "provider_id": "ProviderID", "product_id": "ProductID", "part_numbers": ["PartNumbers"], "image_repo_name": "ImageRepoName"}, "licenses": [{"id": "ID", "name": "Name", "type": "Type", "url": "URL", "description": "Description"}], "image_manifest_url": "ImageManifestURL", "deprecated": true, "package_version": "PackageVersion", "state": {"current": "Current", "current_entered": "2019-01-01T12:00:00", "pending": "Pending", "pending_requested": "2019-01-01T12:00:00", "previous": "Previous"}, "version_locator": "VersionLocator", "console_url": "ConsoleURL", "long_description": "LongDescription", "whitelisted_accounts": ["WhitelistedAccounts"]}], "plans": [{"id": "ID", "label": "Label", "name": "Name", "short_description": "ShortDescription", "long_description": "LongDescription", "metadata": {"anyKey": "anyValue"}, "tags": ["Tags"], "additional_features": [{"title": "Title", "description": "Description"}], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "deployments": [{"id": "ID", "label": "Label", "name": "Name", "short_description": "ShortDescription", "long_description": "LongDescription", "metadata": {"anyKey": "anyValue"}, "tags": ["Tags"], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00"}]}]}], "permit_request_ibm_public_publish": false, "ibm_publish_approved": true, "public_publish_approved": false, "public_original_crn": "PublicOriginalCrn", "publish_public_crn": "PublishPublicCrn", "portal_approval_record": "PortalApprovalRecord", "portal_ui_url": "PortalUiURL", "catalog_id": "CatalogID", "catalog_name": "CatalogName", "metadata": {"anyKey": "anyValue"}, "disclaimer": "Disclaimer", "hidden": true, "provider": "Provider", "repo_info": {"token": "Token", "type": "Type"}}`)
				}))
			})
			It(`Invoke CreateOffering successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.CreateOffering(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Rating model
				ratingModel := new(catalogmanagementv1.Rating)
				ratingModel.OneStarCount = core.Int64Ptr(int64(38))
				ratingModel.TwoStarCount = core.Int64Ptr(int64(38))
				ratingModel.ThreeStarCount = core.Int64Ptr(int64(38))
				ratingModel.FourStarCount = core.Int64Ptr(int64(38))

				// Construct an instance of the Feature model
				featureModel := new(catalogmanagementv1.Feature)
				featureModel.Title = core.StringPtr("testString")
				featureModel.Description = core.StringPtr("testString")

				// Construct an instance of the Configuration model
				configurationModel := new(catalogmanagementv1.Configuration)
				configurationModel.Key = core.StringPtr("testString")
				configurationModel.Type = core.StringPtr("testString")
				configurationModel.DefaultValue = core.StringPtr("testString")
				configurationModel.ValueConstraint = core.StringPtr("testString")
				configurationModel.Description = core.StringPtr("testString")
				configurationModel.Required = core.BoolPtr(true)
				configurationModel.Options = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				configurationModel.Hidden = core.BoolPtr(true)

				// Construct an instance of the Validation model
				validationModel := new(catalogmanagementv1.Validation)
				validationModel.Validated = CreateMockDateTime()
				validationModel.Requested = CreateMockDateTime()
				validationModel.State = core.StringPtr("testString")
				validationModel.LastOperation = core.StringPtr("testString")
				validationModel.Target = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the Resource model
				resourceModel := new(catalogmanagementv1.Resource)
				resourceModel.Type = core.StringPtr("mem")
				resourceModel.Value = core.StringPtr("testString")

				// Construct an instance of the Script model
				scriptModel := new(catalogmanagementv1.Script)
				scriptModel.Instructions = core.StringPtr("testString")
				scriptModel.Script = core.StringPtr("testString")
				scriptModel.ScriptPermission = core.StringPtr("testString")
				scriptModel.DeleteScript = core.StringPtr("testString")
				scriptModel.Scope = core.StringPtr("testString")

				// Construct an instance of the VersionEntitlement model
				versionEntitlementModel := new(catalogmanagementv1.VersionEntitlement)
				versionEntitlementModel.ProviderName = core.StringPtr("testString")
				versionEntitlementModel.ProviderID = core.StringPtr("testString")
				versionEntitlementModel.ProductID = core.StringPtr("testString")
				versionEntitlementModel.PartNumbers = []string{"testString"}
				versionEntitlementModel.ImageRepoName = core.StringPtr("testString")

				// Construct an instance of the License model
				licenseModel := new(catalogmanagementv1.License)
				licenseModel.ID = core.StringPtr("testString")
				licenseModel.Name = core.StringPtr("testString")
				licenseModel.Type = core.StringPtr("testString")
				licenseModel.URL = core.StringPtr("testString")
				licenseModel.Description = core.StringPtr("testString")

				// Construct an instance of the State model
				stateModel := new(catalogmanagementv1.State)
				stateModel.Current = core.StringPtr("testString")
				stateModel.CurrentEntered = CreateMockDateTime()
				stateModel.Pending = core.StringPtr("testString")
				stateModel.PendingRequested = CreateMockDateTime()
				stateModel.Previous = core.StringPtr("testString")

				// Construct an instance of the Version model
				versionModel := new(catalogmanagementv1.Version)
				versionModel.ID = core.StringPtr("testString")
				versionModel.Rev = core.StringPtr("testString")
				versionModel.Crn = core.StringPtr("testString")
				versionModel.Version = core.StringPtr("testString")
				versionModel.Sha = core.StringPtr("testString")
				versionModel.Created = CreateMockDateTime()
				versionModel.Updated = CreateMockDateTime()
				versionModel.OfferingID = core.StringPtr("testString")
				versionModel.CatalogID = core.StringPtr("testString")
				versionModel.KindID = core.StringPtr("testString")
				versionModel.Tags = []string{"testString"}
				versionModel.RepoURL = core.StringPtr("testString")
				versionModel.SourceURL = core.StringPtr("testString")
				versionModel.TgzURL = core.StringPtr("testString")
				versionModel.Configuration = []catalogmanagementv1.Configuration{*configurationModel}
				versionModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}
				versionModel.Validation = validationModel
				versionModel.RequiredResources = []catalogmanagementv1.Resource{*resourceModel}
				versionModel.SingleInstance = core.BoolPtr(true)
				versionModel.Install = scriptModel
				versionModel.PreInstall = []catalogmanagementv1.Script{*scriptModel}
				versionModel.Entitlement = versionEntitlementModel
				versionModel.Licenses = []catalogmanagementv1.License{*licenseModel}
				versionModel.ImageManifestURL = core.StringPtr("testString")
				versionModel.Deprecated = core.BoolPtr(true)
				versionModel.PackageVersion = core.StringPtr("testString")
				versionModel.State = stateModel
				versionModel.VersionLocator = core.StringPtr("testString")
				versionModel.ConsoleURL = core.StringPtr("testString")
				versionModel.LongDescription = core.StringPtr("testString")
				versionModel.WhitelistedAccounts = []string{"testString"}

				// Construct an instance of the Deployment model
				deploymentModel := new(catalogmanagementv1.Deployment)
				deploymentModel.ID = core.StringPtr("testString")
				deploymentModel.Label = core.StringPtr("testString")
				deploymentModel.Name = core.StringPtr("testString")
				deploymentModel.ShortDescription = core.StringPtr("testString")
				deploymentModel.LongDescription = core.StringPtr("testString")
				deploymentModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}
				deploymentModel.Tags = []string{"testString"}
				deploymentModel.Created = CreateMockDateTime()
				deploymentModel.Updated = CreateMockDateTime()

				// Construct an instance of the Plan model
				planModel := new(catalogmanagementv1.Plan)
				planModel.ID = core.StringPtr("testString")
				planModel.Label = core.StringPtr("testString")
				planModel.Name = core.StringPtr("testString")
				planModel.ShortDescription = core.StringPtr("testString")
				planModel.LongDescription = core.StringPtr("testString")
				planModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}
				planModel.Tags = []string{"testString"}
				planModel.AdditionalFeatures = []catalogmanagementv1.Feature{*featureModel}
				planModel.Created = CreateMockDateTime()
				planModel.Updated = CreateMockDateTime()
				planModel.Deployments = []catalogmanagementv1.Deployment{*deploymentModel}

				// Construct an instance of the Kind model
				kindModel := new(catalogmanagementv1.Kind)
				kindModel.ID = core.StringPtr("testString")
				kindModel.FormatKind = core.StringPtr("testString")
				kindModel.TargetKind = core.StringPtr("testString")
				kindModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}
				kindModel.InstallDescription = core.StringPtr("testString")
				kindModel.Tags = []string{"testString"}
				kindModel.AdditionalFeatures = []catalogmanagementv1.Feature{*featureModel}
				kindModel.Created = CreateMockDateTime()
				kindModel.Updated = CreateMockDateTime()
				kindModel.Versions = []catalogmanagementv1.Version{*versionModel}
				kindModel.Plans = []catalogmanagementv1.Plan{*planModel}

				// Construct an instance of the RepoInfo model
				repoInfoModel := new(catalogmanagementv1.RepoInfo)
				repoInfoModel.Token = core.StringPtr("testString")
				repoInfoModel.Type = core.StringPtr("testString")

				// Construct an instance of the CreateOfferingOptions model
				createOfferingOptionsModel := new(catalogmanagementv1.CreateOfferingOptions)
				createOfferingOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				createOfferingOptionsModel.ID = core.StringPtr("testString")
				createOfferingOptionsModel.Rev = core.StringPtr("testString")
				createOfferingOptionsModel.URL = core.StringPtr("testString")
				createOfferingOptionsModel.Crn = core.StringPtr("testString")
				createOfferingOptionsModel.Label = core.StringPtr("testString")
				createOfferingOptionsModel.Name = core.StringPtr("testString")
				createOfferingOptionsModel.OfferingIconURL = core.StringPtr("testString")
				createOfferingOptionsModel.OfferingDocsURL = core.StringPtr("testString")
				createOfferingOptionsModel.OfferingSupportURL = core.StringPtr("testString")
				createOfferingOptionsModel.Tags = []string{"testString"}
				createOfferingOptionsModel.Rating = ratingModel
				createOfferingOptionsModel.Created = CreateMockDateTime()
				createOfferingOptionsModel.Updated = CreateMockDateTime()
				createOfferingOptionsModel.ShortDescription = core.StringPtr("testString")
				createOfferingOptionsModel.LongDescription = core.StringPtr("testString")
				createOfferingOptionsModel.Features = []catalogmanagementv1.Feature{*featureModel}
				createOfferingOptionsModel.Kinds = []catalogmanagementv1.Kind{*kindModel}
				createOfferingOptionsModel.PermitRequestIbmPublicPublish = core.BoolPtr(true)
				createOfferingOptionsModel.IbmPublishApproved = core.BoolPtr(true)
				createOfferingOptionsModel.PublicPublishApproved = core.BoolPtr(true)
				createOfferingOptionsModel.PublicOriginalCrn = core.StringPtr("testString")
				createOfferingOptionsModel.PublishPublicCrn = core.StringPtr("testString")
				createOfferingOptionsModel.PortalApprovalRecord = core.StringPtr("testString")
				createOfferingOptionsModel.PortalUiURL = core.StringPtr("testString")
				createOfferingOptionsModel.CatalogID = core.StringPtr("testString")
				createOfferingOptionsModel.CatalogName = core.StringPtr("testString")
				createOfferingOptionsModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}
				createOfferingOptionsModel.Disclaimer = core.StringPtr("testString")
				createOfferingOptionsModel.Hidden = core.BoolPtr(true)
				createOfferingOptionsModel.Provider = core.StringPtr("testString")
				createOfferingOptionsModel.RepoInfo = repoInfoModel
				createOfferingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.CreateOffering(createOfferingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.CreateOfferingWithContext(ctx, createOfferingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.CreateOffering(createOfferingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.CreateOfferingWithContext(ctx, createOfferingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke CreateOffering with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the Rating model
				ratingModel := new(catalogmanagementv1.Rating)
				ratingModel.OneStarCount = core.Int64Ptr(int64(38))
				ratingModel.TwoStarCount = core.Int64Ptr(int64(38))
				ratingModel.ThreeStarCount = core.Int64Ptr(int64(38))
				ratingModel.FourStarCount = core.Int64Ptr(int64(38))

				// Construct an instance of the Feature model
				featureModel := new(catalogmanagementv1.Feature)
				featureModel.Title = core.StringPtr("testString")
				featureModel.Description = core.StringPtr("testString")

				// Construct an instance of the Configuration model
				configurationModel := new(catalogmanagementv1.Configuration)
				configurationModel.Key = core.StringPtr("testString")
				configurationModel.Type = core.StringPtr("testString")
				configurationModel.DefaultValue = core.StringPtr("testString")
				configurationModel.ValueConstraint = core.StringPtr("testString")
				configurationModel.Description = core.StringPtr("testString")
				configurationModel.Required = core.BoolPtr(true)
				configurationModel.Options = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				configurationModel.Hidden = core.BoolPtr(true)

				// Construct an instance of the Validation model
				validationModel := new(catalogmanagementv1.Validation)
				validationModel.Validated = CreateMockDateTime()
				validationModel.Requested = CreateMockDateTime()
				validationModel.State = core.StringPtr("testString")
				validationModel.LastOperation = core.StringPtr("testString")
				validationModel.Target = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the Resource model
				resourceModel := new(catalogmanagementv1.Resource)
				resourceModel.Type = core.StringPtr("mem")
				resourceModel.Value = core.StringPtr("testString")

				// Construct an instance of the Script model
				scriptModel := new(catalogmanagementv1.Script)
				scriptModel.Instructions = core.StringPtr("testString")
				scriptModel.Script = core.StringPtr("testString")
				scriptModel.ScriptPermission = core.StringPtr("testString")
				scriptModel.DeleteScript = core.StringPtr("testString")
				scriptModel.Scope = core.StringPtr("testString")

				// Construct an instance of the VersionEntitlement model
				versionEntitlementModel := new(catalogmanagementv1.VersionEntitlement)
				versionEntitlementModel.ProviderName = core.StringPtr("testString")
				versionEntitlementModel.ProviderID = core.StringPtr("testString")
				versionEntitlementModel.ProductID = core.StringPtr("testString")
				versionEntitlementModel.PartNumbers = []string{"testString"}
				versionEntitlementModel.ImageRepoName = core.StringPtr("testString")

				// Construct an instance of the License model
				licenseModel := new(catalogmanagementv1.License)
				licenseModel.ID = core.StringPtr("testString")
				licenseModel.Name = core.StringPtr("testString")
				licenseModel.Type = core.StringPtr("testString")
				licenseModel.URL = core.StringPtr("testString")
				licenseModel.Description = core.StringPtr("testString")

				// Construct an instance of the State model
				stateModel := new(catalogmanagementv1.State)
				stateModel.Current = core.StringPtr("testString")
				stateModel.CurrentEntered = CreateMockDateTime()
				stateModel.Pending = core.StringPtr("testString")
				stateModel.PendingRequested = CreateMockDateTime()
				stateModel.Previous = core.StringPtr("testString")

				// Construct an instance of the Version model
				versionModel := new(catalogmanagementv1.Version)
				versionModel.ID = core.StringPtr("testString")
				versionModel.Rev = core.StringPtr("testString")
				versionModel.Crn = core.StringPtr("testString")
				versionModel.Version = core.StringPtr("testString")
				versionModel.Sha = core.StringPtr("testString")
				versionModel.Created = CreateMockDateTime()
				versionModel.Updated = CreateMockDateTime()
				versionModel.OfferingID = core.StringPtr("testString")
				versionModel.CatalogID = core.StringPtr("testString")
				versionModel.KindID = core.StringPtr("testString")
				versionModel.Tags = []string{"testString"}
				versionModel.RepoURL = core.StringPtr("testString")
				versionModel.SourceURL = core.StringPtr("testString")
				versionModel.TgzURL = core.StringPtr("testString")
				versionModel.Configuration = []catalogmanagementv1.Configuration{*configurationModel}
				versionModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}
				versionModel.Validation = validationModel
				versionModel.RequiredResources = []catalogmanagementv1.Resource{*resourceModel}
				versionModel.SingleInstance = core.BoolPtr(true)
				versionModel.Install = scriptModel
				versionModel.PreInstall = []catalogmanagementv1.Script{*scriptModel}
				versionModel.Entitlement = versionEntitlementModel
				versionModel.Licenses = []catalogmanagementv1.License{*licenseModel}
				versionModel.ImageManifestURL = core.StringPtr("testString")
				versionModel.Deprecated = core.BoolPtr(true)
				versionModel.PackageVersion = core.StringPtr("testString")
				versionModel.State = stateModel
				versionModel.VersionLocator = core.StringPtr("testString")
				versionModel.ConsoleURL = core.StringPtr("testString")
				versionModel.LongDescription = core.StringPtr("testString")
				versionModel.WhitelistedAccounts = []string{"testString"}

				// Construct an instance of the Deployment model
				deploymentModel := new(catalogmanagementv1.Deployment)
				deploymentModel.ID = core.StringPtr("testString")
				deploymentModel.Label = core.StringPtr("testString")
				deploymentModel.Name = core.StringPtr("testString")
				deploymentModel.ShortDescription = core.StringPtr("testString")
				deploymentModel.LongDescription = core.StringPtr("testString")
				deploymentModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}
				deploymentModel.Tags = []string{"testString"}
				deploymentModel.Created = CreateMockDateTime()
				deploymentModel.Updated = CreateMockDateTime()

				// Construct an instance of the Plan model
				planModel := new(catalogmanagementv1.Plan)
				planModel.ID = core.StringPtr("testString")
				planModel.Label = core.StringPtr("testString")
				planModel.Name = core.StringPtr("testString")
				planModel.ShortDescription = core.StringPtr("testString")
				planModel.LongDescription = core.StringPtr("testString")
				planModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}
				planModel.Tags = []string{"testString"}
				planModel.AdditionalFeatures = []catalogmanagementv1.Feature{*featureModel}
				planModel.Created = CreateMockDateTime()
				planModel.Updated = CreateMockDateTime()
				planModel.Deployments = []catalogmanagementv1.Deployment{*deploymentModel}

				// Construct an instance of the Kind model
				kindModel := new(catalogmanagementv1.Kind)
				kindModel.ID = core.StringPtr("testString")
				kindModel.FormatKind = core.StringPtr("testString")
				kindModel.TargetKind = core.StringPtr("testString")
				kindModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}
				kindModel.InstallDescription = core.StringPtr("testString")
				kindModel.Tags = []string{"testString"}
				kindModel.AdditionalFeatures = []catalogmanagementv1.Feature{*featureModel}
				kindModel.Created = CreateMockDateTime()
				kindModel.Updated = CreateMockDateTime()
				kindModel.Versions = []catalogmanagementv1.Version{*versionModel}
				kindModel.Plans = []catalogmanagementv1.Plan{*planModel}

				// Construct an instance of the RepoInfo model
				repoInfoModel := new(catalogmanagementv1.RepoInfo)
				repoInfoModel.Token = core.StringPtr("testString")
				repoInfoModel.Type = core.StringPtr("testString")

				// Construct an instance of the CreateOfferingOptions model
				createOfferingOptionsModel := new(catalogmanagementv1.CreateOfferingOptions)
				createOfferingOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				createOfferingOptionsModel.ID = core.StringPtr("testString")
				createOfferingOptionsModel.Rev = core.StringPtr("testString")
				createOfferingOptionsModel.URL = core.StringPtr("testString")
				createOfferingOptionsModel.Crn = core.StringPtr("testString")
				createOfferingOptionsModel.Label = core.StringPtr("testString")
				createOfferingOptionsModel.Name = core.StringPtr("testString")
				createOfferingOptionsModel.OfferingIconURL = core.StringPtr("testString")
				createOfferingOptionsModel.OfferingDocsURL = core.StringPtr("testString")
				createOfferingOptionsModel.OfferingSupportURL = core.StringPtr("testString")
				createOfferingOptionsModel.Tags = []string{"testString"}
				createOfferingOptionsModel.Rating = ratingModel
				createOfferingOptionsModel.Created = CreateMockDateTime()
				createOfferingOptionsModel.Updated = CreateMockDateTime()
				createOfferingOptionsModel.ShortDescription = core.StringPtr("testString")
				createOfferingOptionsModel.LongDescription = core.StringPtr("testString")
				createOfferingOptionsModel.Features = []catalogmanagementv1.Feature{*featureModel}
				createOfferingOptionsModel.Kinds = []catalogmanagementv1.Kind{*kindModel}
				createOfferingOptionsModel.PermitRequestIbmPublicPublish = core.BoolPtr(true)
				createOfferingOptionsModel.IbmPublishApproved = core.BoolPtr(true)
				createOfferingOptionsModel.PublicPublishApproved = core.BoolPtr(true)
				createOfferingOptionsModel.PublicOriginalCrn = core.StringPtr("testString")
				createOfferingOptionsModel.PublishPublicCrn = core.StringPtr("testString")
				createOfferingOptionsModel.PortalApprovalRecord = core.StringPtr("testString")
				createOfferingOptionsModel.PortalUiURL = core.StringPtr("testString")
				createOfferingOptionsModel.CatalogID = core.StringPtr("testString")
				createOfferingOptionsModel.CatalogName = core.StringPtr("testString")
				createOfferingOptionsModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}
				createOfferingOptionsModel.Disclaimer = core.StringPtr("testString")
				createOfferingOptionsModel.Hidden = core.BoolPtr(true)
				createOfferingOptionsModel.Provider = core.StringPtr("testString")
				createOfferingOptionsModel.RepoInfo = repoInfoModel
				createOfferingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.CreateOffering(createOfferingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateOfferingOptions model with no property values
				createOfferingOptionsModelNew := new(catalogmanagementv1.CreateOfferingOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = catalogManagementService.CreateOffering(createOfferingOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ImportOfferingVersion(importOfferingVersionOptions *ImportOfferingVersionOptions) - Operation response error`, func() {
		importOfferingVersionPath := "/catalogs/testString/offerings/testString/version"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(importOfferingVersionPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["zipurl"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["targetVersion"]).To(Equal([]string{"testString"}))


					// TODO: Add check for includeConfig query parameter

					Expect(req.URL.Query()["repoType"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ImportOfferingVersion with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the ImportOfferingVersionOptions model
				importOfferingVersionOptionsModel := new(catalogmanagementv1.ImportOfferingVersionOptions)
				importOfferingVersionOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				importOfferingVersionOptionsModel.OfferingID = core.StringPtr("testString")
				importOfferingVersionOptionsModel.Tags = []string{"testString"}
				importOfferingVersionOptionsModel.TargetKinds = []string{"testString"}
				importOfferingVersionOptionsModel.Content = []int64{int64(38)}
				importOfferingVersionOptionsModel.Zipurl = core.StringPtr("testString")
				importOfferingVersionOptionsModel.TargetVersion = core.StringPtr("testString")
				importOfferingVersionOptionsModel.IncludeConfig = core.BoolPtr(true)
				importOfferingVersionOptionsModel.RepoType = core.StringPtr("testString")
				importOfferingVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.ImportOfferingVersion(importOfferingVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.ImportOfferingVersion(importOfferingVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ImportOfferingVersion(importOfferingVersionOptions *ImportOfferingVersionOptions)`, func() {
		importOfferingVersionPath := "/catalogs/testString/offerings/testString/version"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(importOfferingVersionPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["zipurl"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["targetVersion"]).To(Equal([]string{"testString"}))


					// TODO: Add check for includeConfig query parameter

					Expect(req.URL.Query()["repoType"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "_rev": "Rev", "url": "URL", "crn": "Crn", "label": "Label", "name": "Name", "offering_icon_url": "OfferingIconURL", "offering_docs_url": "OfferingDocsURL", "offering_support_url": "OfferingSupportURL", "tags": ["Tags"], "rating": {"one_star_count": 12, "two_star_count": 12, "three_star_count": 14, "four_star_count": 13}, "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "short_description": "ShortDescription", "long_description": "LongDescription", "features": [{"title": "Title", "description": "Description"}], "kinds": [{"id": "ID", "format_kind": "FormatKind", "target_kind": "TargetKind", "metadata": {"anyKey": "anyValue"}, "install_description": "InstallDescription", "tags": ["Tags"], "additional_features": [{"title": "Title", "description": "Description"}], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "versions": [{"id": "ID", "_rev": "Rev", "crn": "Crn", "version": "Version", "sha": "Sha", "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "offering_id": "OfferingID", "catalog_id": "CatalogID", "kind_id": "KindID", "tags": ["Tags"], "repo_url": "RepoURL", "source_url": "SourceURL", "tgz_url": "TgzURL", "configuration": [{"key": "Key", "type": "Type", "default_value": "anyValue", "value_constraint": "ValueConstraint", "description": "Description", "required": true, "options": [{"anyKey": "anyValue"}], "hidden": true}], "metadata": {"anyKey": "anyValue"}, "validation": {"validated": "2019-01-01T12:00:00", "requested": "2019-01-01T12:00:00", "state": "State", "last_operation": "LastOperation", "target": {"anyKey": "anyValue"}}, "required_resources": [{"type": "mem", "value": "anyValue"}], "single_instance": true, "install": {"instructions": "Instructions", "script": "Script", "script_permission": "ScriptPermission", "delete_script": "DeleteScript", "scope": "Scope"}, "pre_install": [{"instructions": "Instructions", "script": "Script", "script_permission": "ScriptPermission", "delete_script": "DeleteScript", "scope": "Scope"}], "entitlement": {"provider_name": "ProviderName", "provider_id": "ProviderID", "product_id": "ProductID", "part_numbers": ["PartNumbers"], "image_repo_name": "ImageRepoName"}, "licenses": [{"id": "ID", "name": "Name", "type": "Type", "url": "URL", "description": "Description"}], "image_manifest_url": "ImageManifestURL", "deprecated": true, "package_version": "PackageVersion", "state": {"current": "Current", "current_entered": "2019-01-01T12:00:00", "pending": "Pending", "pending_requested": "2019-01-01T12:00:00", "previous": "Previous"}, "version_locator": "VersionLocator", "console_url": "ConsoleURL", "long_description": "LongDescription", "whitelisted_accounts": ["WhitelistedAccounts"]}], "plans": [{"id": "ID", "label": "Label", "name": "Name", "short_description": "ShortDescription", "long_description": "LongDescription", "metadata": {"anyKey": "anyValue"}, "tags": ["Tags"], "additional_features": [{"title": "Title", "description": "Description"}], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "deployments": [{"id": "ID", "label": "Label", "name": "Name", "short_description": "ShortDescription", "long_description": "LongDescription", "metadata": {"anyKey": "anyValue"}, "tags": ["Tags"], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00"}]}]}], "permit_request_ibm_public_publish": false, "ibm_publish_approved": true, "public_publish_approved": false, "public_original_crn": "PublicOriginalCrn", "publish_public_crn": "PublishPublicCrn", "portal_approval_record": "PortalApprovalRecord", "portal_ui_url": "PortalUiURL", "catalog_id": "CatalogID", "catalog_name": "CatalogName", "metadata": {"anyKey": "anyValue"}, "disclaimer": "Disclaimer", "hidden": true, "provider": "Provider", "repo_info": {"token": "Token", "type": "Type"}}`)
				}))
			})
			It(`Invoke ImportOfferingVersion successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.ImportOfferingVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ImportOfferingVersionOptions model
				importOfferingVersionOptionsModel := new(catalogmanagementv1.ImportOfferingVersionOptions)
				importOfferingVersionOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				importOfferingVersionOptionsModel.OfferingID = core.StringPtr("testString")
				importOfferingVersionOptionsModel.Tags = []string{"testString"}
				importOfferingVersionOptionsModel.TargetKinds = []string{"testString"}
				importOfferingVersionOptionsModel.Content = []int64{int64(38)}
				importOfferingVersionOptionsModel.Zipurl = core.StringPtr("testString")
				importOfferingVersionOptionsModel.TargetVersion = core.StringPtr("testString")
				importOfferingVersionOptionsModel.IncludeConfig = core.BoolPtr(true)
				importOfferingVersionOptionsModel.RepoType = core.StringPtr("testString")
				importOfferingVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.ImportOfferingVersion(importOfferingVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.ImportOfferingVersionWithContext(ctx, importOfferingVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.ImportOfferingVersion(importOfferingVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.ImportOfferingVersionWithContext(ctx, importOfferingVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ImportOfferingVersion with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the ImportOfferingVersionOptions model
				importOfferingVersionOptionsModel := new(catalogmanagementv1.ImportOfferingVersionOptions)
				importOfferingVersionOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				importOfferingVersionOptionsModel.OfferingID = core.StringPtr("testString")
				importOfferingVersionOptionsModel.Tags = []string{"testString"}
				importOfferingVersionOptionsModel.TargetKinds = []string{"testString"}
				importOfferingVersionOptionsModel.Content = []int64{int64(38)}
				importOfferingVersionOptionsModel.Zipurl = core.StringPtr("testString")
				importOfferingVersionOptionsModel.TargetVersion = core.StringPtr("testString")
				importOfferingVersionOptionsModel.IncludeConfig = core.BoolPtr(true)
				importOfferingVersionOptionsModel.RepoType = core.StringPtr("testString")
				importOfferingVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.ImportOfferingVersion(importOfferingVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ImportOfferingVersionOptions model with no property values
				importOfferingVersionOptionsModelNew := new(catalogmanagementv1.ImportOfferingVersionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = catalogManagementService.ImportOfferingVersion(importOfferingVersionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ImportOffering(importOfferingOptions *ImportOfferingOptions) - Operation response error`, func() {
		importOfferingPath := "/catalogs/testString/import/offerings"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(importOfferingPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Auth-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["zipurl"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["offeringID"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["targetVersion"]).To(Equal([]string{"testString"}))


					// TODO: Add check for includeConfig query parameter

					Expect(req.URL.Query()["repoType"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ImportOffering with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the ImportOfferingOptions model
				importOfferingOptionsModel := new(catalogmanagementv1.ImportOfferingOptions)
				importOfferingOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				importOfferingOptionsModel.Tags = []string{"testString"}
				importOfferingOptionsModel.TargetKinds = []string{"testString"}
				importOfferingOptionsModel.Content = []int64{int64(38)}
				importOfferingOptionsModel.Zipurl = core.StringPtr("testString")
				importOfferingOptionsModel.OfferingID = core.StringPtr("testString")
				importOfferingOptionsModel.TargetVersion = core.StringPtr("testString")
				importOfferingOptionsModel.IncludeConfig = core.BoolPtr(true)
				importOfferingOptionsModel.RepoType = core.StringPtr("testString")
				importOfferingOptionsModel.XAuthToken = core.StringPtr("testString")
				importOfferingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.ImportOffering(importOfferingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.ImportOffering(importOfferingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ImportOffering(importOfferingOptions *ImportOfferingOptions)`, func() {
		importOfferingPath := "/catalogs/testString/import/offerings"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(importOfferingPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["X-Auth-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["zipurl"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["offeringID"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["targetVersion"]).To(Equal([]string{"testString"}))


					// TODO: Add check for includeConfig query parameter

					Expect(req.URL.Query()["repoType"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "_rev": "Rev", "url": "URL", "crn": "Crn", "label": "Label", "name": "Name", "offering_icon_url": "OfferingIconURL", "offering_docs_url": "OfferingDocsURL", "offering_support_url": "OfferingSupportURL", "tags": ["Tags"], "rating": {"one_star_count": 12, "two_star_count": 12, "three_star_count": 14, "four_star_count": 13}, "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "short_description": "ShortDescription", "long_description": "LongDescription", "features": [{"title": "Title", "description": "Description"}], "kinds": [{"id": "ID", "format_kind": "FormatKind", "target_kind": "TargetKind", "metadata": {"anyKey": "anyValue"}, "install_description": "InstallDescription", "tags": ["Tags"], "additional_features": [{"title": "Title", "description": "Description"}], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "versions": [{"id": "ID", "_rev": "Rev", "crn": "Crn", "version": "Version", "sha": "Sha", "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "offering_id": "OfferingID", "catalog_id": "CatalogID", "kind_id": "KindID", "tags": ["Tags"], "repo_url": "RepoURL", "source_url": "SourceURL", "tgz_url": "TgzURL", "configuration": [{"key": "Key", "type": "Type", "default_value": "anyValue", "value_constraint": "ValueConstraint", "description": "Description", "required": true, "options": [{"anyKey": "anyValue"}], "hidden": true}], "metadata": {"anyKey": "anyValue"}, "validation": {"validated": "2019-01-01T12:00:00", "requested": "2019-01-01T12:00:00", "state": "State", "last_operation": "LastOperation", "target": {"anyKey": "anyValue"}}, "required_resources": [{"type": "mem", "value": "anyValue"}], "single_instance": true, "install": {"instructions": "Instructions", "script": "Script", "script_permission": "ScriptPermission", "delete_script": "DeleteScript", "scope": "Scope"}, "pre_install": [{"instructions": "Instructions", "script": "Script", "script_permission": "ScriptPermission", "delete_script": "DeleteScript", "scope": "Scope"}], "entitlement": {"provider_name": "ProviderName", "provider_id": "ProviderID", "product_id": "ProductID", "part_numbers": ["PartNumbers"], "image_repo_name": "ImageRepoName"}, "licenses": [{"id": "ID", "name": "Name", "type": "Type", "url": "URL", "description": "Description"}], "image_manifest_url": "ImageManifestURL", "deprecated": true, "package_version": "PackageVersion", "state": {"current": "Current", "current_entered": "2019-01-01T12:00:00", "pending": "Pending", "pending_requested": "2019-01-01T12:00:00", "previous": "Previous"}, "version_locator": "VersionLocator", "console_url": "ConsoleURL", "long_description": "LongDescription", "whitelisted_accounts": ["WhitelistedAccounts"]}], "plans": [{"id": "ID", "label": "Label", "name": "Name", "short_description": "ShortDescription", "long_description": "LongDescription", "metadata": {"anyKey": "anyValue"}, "tags": ["Tags"], "additional_features": [{"title": "Title", "description": "Description"}], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "deployments": [{"id": "ID", "label": "Label", "name": "Name", "short_description": "ShortDescription", "long_description": "LongDescription", "metadata": {"anyKey": "anyValue"}, "tags": ["Tags"], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00"}]}]}], "permit_request_ibm_public_publish": false, "ibm_publish_approved": true, "public_publish_approved": false, "public_original_crn": "PublicOriginalCrn", "publish_public_crn": "PublishPublicCrn", "portal_approval_record": "PortalApprovalRecord", "portal_ui_url": "PortalUiURL", "catalog_id": "CatalogID", "catalog_name": "CatalogName", "metadata": {"anyKey": "anyValue"}, "disclaimer": "Disclaimer", "hidden": true, "provider": "Provider", "repo_info": {"token": "Token", "type": "Type"}}`)
				}))
			})
			It(`Invoke ImportOffering successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.ImportOffering(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ImportOfferingOptions model
				importOfferingOptionsModel := new(catalogmanagementv1.ImportOfferingOptions)
				importOfferingOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				importOfferingOptionsModel.Tags = []string{"testString"}
				importOfferingOptionsModel.TargetKinds = []string{"testString"}
				importOfferingOptionsModel.Content = []int64{int64(38)}
				importOfferingOptionsModel.Zipurl = core.StringPtr("testString")
				importOfferingOptionsModel.OfferingID = core.StringPtr("testString")
				importOfferingOptionsModel.TargetVersion = core.StringPtr("testString")
				importOfferingOptionsModel.IncludeConfig = core.BoolPtr(true)
				importOfferingOptionsModel.RepoType = core.StringPtr("testString")
				importOfferingOptionsModel.XAuthToken = core.StringPtr("testString")
				importOfferingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.ImportOffering(importOfferingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.ImportOfferingWithContext(ctx, importOfferingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.ImportOffering(importOfferingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.ImportOfferingWithContext(ctx, importOfferingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ImportOffering with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the ImportOfferingOptions model
				importOfferingOptionsModel := new(catalogmanagementv1.ImportOfferingOptions)
				importOfferingOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				importOfferingOptionsModel.Tags = []string{"testString"}
				importOfferingOptionsModel.TargetKinds = []string{"testString"}
				importOfferingOptionsModel.Content = []int64{int64(38)}
				importOfferingOptionsModel.Zipurl = core.StringPtr("testString")
				importOfferingOptionsModel.OfferingID = core.StringPtr("testString")
				importOfferingOptionsModel.TargetVersion = core.StringPtr("testString")
				importOfferingOptionsModel.IncludeConfig = core.BoolPtr(true)
				importOfferingOptionsModel.RepoType = core.StringPtr("testString")
				importOfferingOptionsModel.XAuthToken = core.StringPtr("testString")
				importOfferingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.ImportOffering(importOfferingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ImportOfferingOptions model with no property values
				importOfferingOptionsModelNew := new(catalogmanagementv1.ImportOfferingOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = catalogManagementService.ImportOffering(importOfferingOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReloadOffering(reloadOfferingOptions *ReloadOfferingOptions) - Operation response error`, func() {
		reloadOfferingPath := "/catalogs/testString/offerings/testString/reload"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(reloadOfferingPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.URL.Query()["targetVersion"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["zipurl"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["repoType"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReloadOffering with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the ReloadOfferingOptions model
				reloadOfferingOptionsModel := new(catalogmanagementv1.ReloadOfferingOptions)
				reloadOfferingOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				reloadOfferingOptionsModel.OfferingID = core.StringPtr("testString")
				reloadOfferingOptionsModel.TargetVersion = core.StringPtr("testString")
				reloadOfferingOptionsModel.Tags = []string{"testString"}
				reloadOfferingOptionsModel.TargetKinds = []string{"testString"}
				reloadOfferingOptionsModel.Content = []int64{int64(38)}
				reloadOfferingOptionsModel.Zipurl = core.StringPtr("testString")
				reloadOfferingOptionsModel.RepoType = core.StringPtr("testString")
				reloadOfferingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.ReloadOffering(reloadOfferingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.ReloadOffering(reloadOfferingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ReloadOffering(reloadOfferingOptions *ReloadOfferingOptions)`, func() {
		reloadOfferingPath := "/catalogs/testString/offerings/testString/reload"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(reloadOfferingPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["targetVersion"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["zipurl"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["repoType"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "_rev": "Rev", "url": "URL", "crn": "Crn", "label": "Label", "name": "Name", "offering_icon_url": "OfferingIconURL", "offering_docs_url": "OfferingDocsURL", "offering_support_url": "OfferingSupportURL", "tags": ["Tags"], "rating": {"one_star_count": 12, "two_star_count": 12, "three_star_count": 14, "four_star_count": 13}, "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "short_description": "ShortDescription", "long_description": "LongDescription", "features": [{"title": "Title", "description": "Description"}], "kinds": [{"id": "ID", "format_kind": "FormatKind", "target_kind": "TargetKind", "metadata": {"anyKey": "anyValue"}, "install_description": "InstallDescription", "tags": ["Tags"], "additional_features": [{"title": "Title", "description": "Description"}], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "versions": [{"id": "ID", "_rev": "Rev", "crn": "Crn", "version": "Version", "sha": "Sha", "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "offering_id": "OfferingID", "catalog_id": "CatalogID", "kind_id": "KindID", "tags": ["Tags"], "repo_url": "RepoURL", "source_url": "SourceURL", "tgz_url": "TgzURL", "configuration": [{"key": "Key", "type": "Type", "default_value": "anyValue", "value_constraint": "ValueConstraint", "description": "Description", "required": true, "options": [{"anyKey": "anyValue"}], "hidden": true}], "metadata": {"anyKey": "anyValue"}, "validation": {"validated": "2019-01-01T12:00:00", "requested": "2019-01-01T12:00:00", "state": "State", "last_operation": "LastOperation", "target": {"anyKey": "anyValue"}}, "required_resources": [{"type": "mem", "value": "anyValue"}], "single_instance": true, "install": {"instructions": "Instructions", "script": "Script", "script_permission": "ScriptPermission", "delete_script": "DeleteScript", "scope": "Scope"}, "pre_install": [{"instructions": "Instructions", "script": "Script", "script_permission": "ScriptPermission", "delete_script": "DeleteScript", "scope": "Scope"}], "entitlement": {"provider_name": "ProviderName", "provider_id": "ProviderID", "product_id": "ProductID", "part_numbers": ["PartNumbers"], "image_repo_name": "ImageRepoName"}, "licenses": [{"id": "ID", "name": "Name", "type": "Type", "url": "URL", "description": "Description"}], "image_manifest_url": "ImageManifestURL", "deprecated": true, "package_version": "PackageVersion", "state": {"current": "Current", "current_entered": "2019-01-01T12:00:00", "pending": "Pending", "pending_requested": "2019-01-01T12:00:00", "previous": "Previous"}, "version_locator": "VersionLocator", "console_url": "ConsoleURL", "long_description": "LongDescription", "whitelisted_accounts": ["WhitelistedAccounts"]}], "plans": [{"id": "ID", "label": "Label", "name": "Name", "short_description": "ShortDescription", "long_description": "LongDescription", "metadata": {"anyKey": "anyValue"}, "tags": ["Tags"], "additional_features": [{"title": "Title", "description": "Description"}], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "deployments": [{"id": "ID", "label": "Label", "name": "Name", "short_description": "ShortDescription", "long_description": "LongDescription", "metadata": {"anyKey": "anyValue"}, "tags": ["Tags"], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00"}]}]}], "permit_request_ibm_public_publish": false, "ibm_publish_approved": true, "public_publish_approved": false, "public_original_crn": "PublicOriginalCrn", "publish_public_crn": "PublishPublicCrn", "portal_approval_record": "PortalApprovalRecord", "portal_ui_url": "PortalUiURL", "catalog_id": "CatalogID", "catalog_name": "CatalogName", "metadata": {"anyKey": "anyValue"}, "disclaimer": "Disclaimer", "hidden": true, "provider": "Provider", "repo_info": {"token": "Token", "type": "Type"}}`)
				}))
			})
			It(`Invoke ReloadOffering successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.ReloadOffering(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ReloadOfferingOptions model
				reloadOfferingOptionsModel := new(catalogmanagementv1.ReloadOfferingOptions)
				reloadOfferingOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				reloadOfferingOptionsModel.OfferingID = core.StringPtr("testString")
				reloadOfferingOptionsModel.TargetVersion = core.StringPtr("testString")
				reloadOfferingOptionsModel.Tags = []string{"testString"}
				reloadOfferingOptionsModel.TargetKinds = []string{"testString"}
				reloadOfferingOptionsModel.Content = []int64{int64(38)}
				reloadOfferingOptionsModel.Zipurl = core.StringPtr("testString")
				reloadOfferingOptionsModel.RepoType = core.StringPtr("testString")
				reloadOfferingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.ReloadOffering(reloadOfferingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.ReloadOfferingWithContext(ctx, reloadOfferingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.ReloadOffering(reloadOfferingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.ReloadOfferingWithContext(ctx, reloadOfferingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ReloadOffering with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the ReloadOfferingOptions model
				reloadOfferingOptionsModel := new(catalogmanagementv1.ReloadOfferingOptions)
				reloadOfferingOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				reloadOfferingOptionsModel.OfferingID = core.StringPtr("testString")
				reloadOfferingOptionsModel.TargetVersion = core.StringPtr("testString")
				reloadOfferingOptionsModel.Tags = []string{"testString"}
				reloadOfferingOptionsModel.TargetKinds = []string{"testString"}
				reloadOfferingOptionsModel.Content = []int64{int64(38)}
				reloadOfferingOptionsModel.Zipurl = core.StringPtr("testString")
				reloadOfferingOptionsModel.RepoType = core.StringPtr("testString")
				reloadOfferingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.ReloadOffering(reloadOfferingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReloadOfferingOptions model with no property values
				reloadOfferingOptionsModelNew := new(catalogmanagementv1.ReloadOfferingOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = catalogManagementService.ReloadOffering(reloadOfferingOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetOffering(getOfferingOptions *GetOfferingOptions) - Operation response error`, func() {
		getOfferingPath := "/catalogs/testString/offerings/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getOfferingPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetOffering with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetOfferingOptions model
				getOfferingOptionsModel := new(catalogmanagementv1.GetOfferingOptions)
				getOfferingOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				getOfferingOptionsModel.OfferingID = core.StringPtr("testString")
				getOfferingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.GetOffering(getOfferingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.GetOffering(getOfferingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetOffering(getOfferingOptions *GetOfferingOptions)`, func() {
		getOfferingPath := "/catalogs/testString/offerings/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getOfferingPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "_rev": "Rev", "url": "URL", "crn": "Crn", "label": "Label", "name": "Name", "offering_icon_url": "OfferingIconURL", "offering_docs_url": "OfferingDocsURL", "offering_support_url": "OfferingSupportURL", "tags": ["Tags"], "rating": {"one_star_count": 12, "two_star_count": 12, "three_star_count": 14, "four_star_count": 13}, "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "short_description": "ShortDescription", "long_description": "LongDescription", "features": [{"title": "Title", "description": "Description"}], "kinds": [{"id": "ID", "format_kind": "FormatKind", "target_kind": "TargetKind", "metadata": {"anyKey": "anyValue"}, "install_description": "InstallDescription", "tags": ["Tags"], "additional_features": [{"title": "Title", "description": "Description"}], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "versions": [{"id": "ID", "_rev": "Rev", "crn": "Crn", "version": "Version", "sha": "Sha", "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "offering_id": "OfferingID", "catalog_id": "CatalogID", "kind_id": "KindID", "tags": ["Tags"], "repo_url": "RepoURL", "source_url": "SourceURL", "tgz_url": "TgzURL", "configuration": [{"key": "Key", "type": "Type", "default_value": "anyValue", "value_constraint": "ValueConstraint", "description": "Description", "required": true, "options": [{"anyKey": "anyValue"}], "hidden": true}], "metadata": {"anyKey": "anyValue"}, "validation": {"validated": "2019-01-01T12:00:00", "requested": "2019-01-01T12:00:00", "state": "State", "last_operation": "LastOperation", "target": {"anyKey": "anyValue"}}, "required_resources": [{"type": "mem", "value": "anyValue"}], "single_instance": true, "install": {"instructions": "Instructions", "script": "Script", "script_permission": "ScriptPermission", "delete_script": "DeleteScript", "scope": "Scope"}, "pre_install": [{"instructions": "Instructions", "script": "Script", "script_permission": "ScriptPermission", "delete_script": "DeleteScript", "scope": "Scope"}], "entitlement": {"provider_name": "ProviderName", "provider_id": "ProviderID", "product_id": "ProductID", "part_numbers": ["PartNumbers"], "image_repo_name": "ImageRepoName"}, "licenses": [{"id": "ID", "name": "Name", "type": "Type", "url": "URL", "description": "Description"}], "image_manifest_url": "ImageManifestURL", "deprecated": true, "package_version": "PackageVersion", "state": {"current": "Current", "current_entered": "2019-01-01T12:00:00", "pending": "Pending", "pending_requested": "2019-01-01T12:00:00", "previous": "Previous"}, "version_locator": "VersionLocator", "console_url": "ConsoleURL", "long_description": "LongDescription", "whitelisted_accounts": ["WhitelistedAccounts"]}], "plans": [{"id": "ID", "label": "Label", "name": "Name", "short_description": "ShortDescription", "long_description": "LongDescription", "metadata": {"anyKey": "anyValue"}, "tags": ["Tags"], "additional_features": [{"title": "Title", "description": "Description"}], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "deployments": [{"id": "ID", "label": "Label", "name": "Name", "short_description": "ShortDescription", "long_description": "LongDescription", "metadata": {"anyKey": "anyValue"}, "tags": ["Tags"], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00"}]}]}], "permit_request_ibm_public_publish": false, "ibm_publish_approved": true, "public_publish_approved": false, "public_original_crn": "PublicOriginalCrn", "publish_public_crn": "PublishPublicCrn", "portal_approval_record": "PortalApprovalRecord", "portal_ui_url": "PortalUiURL", "catalog_id": "CatalogID", "catalog_name": "CatalogName", "metadata": {"anyKey": "anyValue"}, "disclaimer": "Disclaimer", "hidden": true, "provider": "Provider", "repo_info": {"token": "Token", "type": "Type"}}`)
				}))
			})
			It(`Invoke GetOffering successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.GetOffering(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetOfferingOptions model
				getOfferingOptionsModel := new(catalogmanagementv1.GetOfferingOptions)
				getOfferingOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				getOfferingOptionsModel.OfferingID = core.StringPtr("testString")
				getOfferingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.GetOffering(getOfferingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetOfferingWithContext(ctx, getOfferingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.GetOffering(getOfferingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetOfferingWithContext(ctx, getOfferingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetOffering with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetOfferingOptions model
				getOfferingOptionsModel := new(catalogmanagementv1.GetOfferingOptions)
				getOfferingOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				getOfferingOptionsModel.OfferingID = core.StringPtr("testString")
				getOfferingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.GetOffering(getOfferingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetOfferingOptions model with no property values
				getOfferingOptionsModelNew := new(catalogmanagementv1.GetOfferingOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = catalogManagementService.GetOffering(getOfferingOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceOffering(replaceOfferingOptions *ReplaceOfferingOptions) - Operation response error`, func() {
		replaceOfferingPath := "/catalogs/testString/offerings/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceOfferingPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceOffering with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the Rating model
				ratingModel := new(catalogmanagementv1.Rating)
				ratingModel.OneStarCount = core.Int64Ptr(int64(38))
				ratingModel.TwoStarCount = core.Int64Ptr(int64(38))
				ratingModel.ThreeStarCount = core.Int64Ptr(int64(38))
				ratingModel.FourStarCount = core.Int64Ptr(int64(38))

				// Construct an instance of the Feature model
				featureModel := new(catalogmanagementv1.Feature)
				featureModel.Title = core.StringPtr("testString")
				featureModel.Description = core.StringPtr("testString")

				// Construct an instance of the Configuration model
				configurationModel := new(catalogmanagementv1.Configuration)
				configurationModel.Key = core.StringPtr("testString")
				configurationModel.Type = core.StringPtr("testString")
				configurationModel.DefaultValue = core.StringPtr("testString")
				configurationModel.ValueConstraint = core.StringPtr("testString")
				configurationModel.Description = core.StringPtr("testString")
				configurationModel.Required = core.BoolPtr(true)
				configurationModel.Options = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				configurationModel.Hidden = core.BoolPtr(true)

				// Construct an instance of the Validation model
				validationModel := new(catalogmanagementv1.Validation)
				validationModel.Validated = CreateMockDateTime()
				validationModel.Requested = CreateMockDateTime()
				validationModel.State = core.StringPtr("testString")
				validationModel.LastOperation = core.StringPtr("testString")
				validationModel.Target = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the Resource model
				resourceModel := new(catalogmanagementv1.Resource)
				resourceModel.Type = core.StringPtr("mem")
				resourceModel.Value = core.StringPtr("testString")

				// Construct an instance of the Script model
				scriptModel := new(catalogmanagementv1.Script)
				scriptModel.Instructions = core.StringPtr("testString")
				scriptModel.Script = core.StringPtr("testString")
				scriptModel.ScriptPermission = core.StringPtr("testString")
				scriptModel.DeleteScript = core.StringPtr("testString")
				scriptModel.Scope = core.StringPtr("testString")

				// Construct an instance of the VersionEntitlement model
				versionEntitlementModel := new(catalogmanagementv1.VersionEntitlement)
				versionEntitlementModel.ProviderName = core.StringPtr("testString")
				versionEntitlementModel.ProviderID = core.StringPtr("testString")
				versionEntitlementModel.ProductID = core.StringPtr("testString")
				versionEntitlementModel.PartNumbers = []string{"testString"}
				versionEntitlementModel.ImageRepoName = core.StringPtr("testString")

				// Construct an instance of the License model
				licenseModel := new(catalogmanagementv1.License)
				licenseModel.ID = core.StringPtr("testString")
				licenseModel.Name = core.StringPtr("testString")
				licenseModel.Type = core.StringPtr("testString")
				licenseModel.URL = core.StringPtr("testString")
				licenseModel.Description = core.StringPtr("testString")

				// Construct an instance of the State model
				stateModel := new(catalogmanagementv1.State)
				stateModel.Current = core.StringPtr("testString")
				stateModel.CurrentEntered = CreateMockDateTime()
				stateModel.Pending = core.StringPtr("testString")
				stateModel.PendingRequested = CreateMockDateTime()
				stateModel.Previous = core.StringPtr("testString")

				// Construct an instance of the Version model
				versionModel := new(catalogmanagementv1.Version)
				versionModel.ID = core.StringPtr("testString")
				versionModel.Rev = core.StringPtr("testString")
				versionModel.Crn = core.StringPtr("testString")
				versionModel.Version = core.StringPtr("testString")
				versionModel.Sha = core.StringPtr("testString")
				versionModel.Created = CreateMockDateTime()
				versionModel.Updated = CreateMockDateTime()
				versionModel.OfferingID = core.StringPtr("testString")
				versionModel.CatalogID = core.StringPtr("testString")
				versionModel.KindID = core.StringPtr("testString")
				versionModel.Tags = []string{"testString"}
				versionModel.RepoURL = core.StringPtr("testString")
				versionModel.SourceURL = core.StringPtr("testString")
				versionModel.TgzURL = core.StringPtr("testString")
				versionModel.Configuration = []catalogmanagementv1.Configuration{*configurationModel}
				versionModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}
				versionModel.Validation = validationModel
				versionModel.RequiredResources = []catalogmanagementv1.Resource{*resourceModel}
				versionModel.SingleInstance = core.BoolPtr(true)
				versionModel.Install = scriptModel
				versionModel.PreInstall = []catalogmanagementv1.Script{*scriptModel}
				versionModel.Entitlement = versionEntitlementModel
				versionModel.Licenses = []catalogmanagementv1.License{*licenseModel}
				versionModel.ImageManifestURL = core.StringPtr("testString")
				versionModel.Deprecated = core.BoolPtr(true)
				versionModel.PackageVersion = core.StringPtr("testString")
				versionModel.State = stateModel
				versionModel.VersionLocator = core.StringPtr("testString")
				versionModel.ConsoleURL = core.StringPtr("testString")
				versionModel.LongDescription = core.StringPtr("testString")
				versionModel.WhitelistedAccounts = []string{"testString"}

				// Construct an instance of the Deployment model
				deploymentModel := new(catalogmanagementv1.Deployment)
				deploymentModel.ID = core.StringPtr("testString")
				deploymentModel.Label = core.StringPtr("testString")
				deploymentModel.Name = core.StringPtr("testString")
				deploymentModel.ShortDescription = core.StringPtr("testString")
				deploymentModel.LongDescription = core.StringPtr("testString")
				deploymentModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}
				deploymentModel.Tags = []string{"testString"}
				deploymentModel.Created = CreateMockDateTime()
				deploymentModel.Updated = CreateMockDateTime()

				// Construct an instance of the Plan model
				planModel := new(catalogmanagementv1.Plan)
				planModel.ID = core.StringPtr("testString")
				planModel.Label = core.StringPtr("testString")
				planModel.Name = core.StringPtr("testString")
				planModel.ShortDescription = core.StringPtr("testString")
				planModel.LongDescription = core.StringPtr("testString")
				planModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}
				planModel.Tags = []string{"testString"}
				planModel.AdditionalFeatures = []catalogmanagementv1.Feature{*featureModel}
				planModel.Created = CreateMockDateTime()
				planModel.Updated = CreateMockDateTime()
				planModel.Deployments = []catalogmanagementv1.Deployment{*deploymentModel}

				// Construct an instance of the Kind model
				kindModel := new(catalogmanagementv1.Kind)
				kindModel.ID = core.StringPtr("testString")
				kindModel.FormatKind = core.StringPtr("testString")
				kindModel.TargetKind = core.StringPtr("testString")
				kindModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}
				kindModel.InstallDescription = core.StringPtr("testString")
				kindModel.Tags = []string{"testString"}
				kindModel.AdditionalFeatures = []catalogmanagementv1.Feature{*featureModel}
				kindModel.Created = CreateMockDateTime()
				kindModel.Updated = CreateMockDateTime()
				kindModel.Versions = []catalogmanagementv1.Version{*versionModel}
				kindModel.Plans = []catalogmanagementv1.Plan{*planModel}

				// Construct an instance of the RepoInfo model
				repoInfoModel := new(catalogmanagementv1.RepoInfo)
				repoInfoModel.Token = core.StringPtr("testString")
				repoInfoModel.Type = core.StringPtr("testString")

				// Construct an instance of the ReplaceOfferingOptions model
				replaceOfferingOptionsModel := new(catalogmanagementv1.ReplaceOfferingOptions)
				replaceOfferingOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				replaceOfferingOptionsModel.OfferingID = core.StringPtr("testString")
				replaceOfferingOptionsModel.ID = core.StringPtr("testString")
				replaceOfferingOptionsModel.Rev = core.StringPtr("testString")
				replaceOfferingOptionsModel.URL = core.StringPtr("testString")
				replaceOfferingOptionsModel.Crn = core.StringPtr("testString")
				replaceOfferingOptionsModel.Label = core.StringPtr("testString")
				replaceOfferingOptionsModel.Name = core.StringPtr("testString")
				replaceOfferingOptionsModel.OfferingIconURL = core.StringPtr("testString")
				replaceOfferingOptionsModel.OfferingDocsURL = core.StringPtr("testString")
				replaceOfferingOptionsModel.OfferingSupportURL = core.StringPtr("testString")
				replaceOfferingOptionsModel.Tags = []string{"testString"}
				replaceOfferingOptionsModel.Rating = ratingModel
				replaceOfferingOptionsModel.Created = CreateMockDateTime()
				replaceOfferingOptionsModel.Updated = CreateMockDateTime()
				replaceOfferingOptionsModel.ShortDescription = core.StringPtr("testString")
				replaceOfferingOptionsModel.LongDescription = core.StringPtr("testString")
				replaceOfferingOptionsModel.Features = []catalogmanagementv1.Feature{*featureModel}
				replaceOfferingOptionsModel.Kinds = []catalogmanagementv1.Kind{*kindModel}
				replaceOfferingOptionsModel.PermitRequestIbmPublicPublish = core.BoolPtr(true)
				replaceOfferingOptionsModel.IbmPublishApproved = core.BoolPtr(true)
				replaceOfferingOptionsModel.PublicPublishApproved = core.BoolPtr(true)
				replaceOfferingOptionsModel.PublicOriginalCrn = core.StringPtr("testString")
				replaceOfferingOptionsModel.PublishPublicCrn = core.StringPtr("testString")
				replaceOfferingOptionsModel.PortalApprovalRecord = core.StringPtr("testString")
				replaceOfferingOptionsModel.PortalUiURL = core.StringPtr("testString")
				replaceOfferingOptionsModel.CatalogID = core.StringPtr("testString")
				replaceOfferingOptionsModel.CatalogName = core.StringPtr("testString")
				replaceOfferingOptionsModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}
				replaceOfferingOptionsModel.Disclaimer = core.StringPtr("testString")
				replaceOfferingOptionsModel.Hidden = core.BoolPtr(true)
				replaceOfferingOptionsModel.Provider = core.StringPtr("testString")
				replaceOfferingOptionsModel.RepoInfo = repoInfoModel
				replaceOfferingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.ReplaceOffering(replaceOfferingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.ReplaceOffering(replaceOfferingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ReplaceOffering(replaceOfferingOptions *ReplaceOfferingOptions)`, func() {
		replaceOfferingPath := "/catalogs/testString/offerings/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceOfferingPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "_rev": "Rev", "label": "Label", "short_description": "ShortDescription", "catalog_icon_url": "CatalogIconURL", "tags": ["Tags"], "url": "URL", "crn": "Crn", "offerings_url": "OfferingsURL", "features": [{"title": "Title", "description": "Description"}], "disabled": true, "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "resource_group_id": "ResourceGroupID", "owning_account": "OwningAccount", "catalog_filters": {"include_all": true, "category_filters": {"mapKey": {"include": false, "filter": {"filter_terms": ["FilterTerms"]}}}, "id_filters": {"include": {"filter_terms": ["FilterTerms"]}, "exclude": {"filter_terms": ["FilterTerms"]}}}, "syndication_settings": {"remove_related_components": false, "clusters": [{"region": "Region", "id": "ID", "name": "Name", "resource_group_name": "ResourceGroupName", "type": "Type", "namespaces": ["Namespaces"], "all_namespaces": false}], "history": {"namespaces": ["Namespaces"], "clusters": [{"region": "Region", "id": "ID", "name": "Name", "resource_group_name": "ResourceGroupName", "type": "Type", "namespaces": ["Namespaces"], "all_namespaces": false}], "last_run": "2019-01-01T12:00:00"}, "authorization": {"token": "Token", "last_run": "2019-01-01T12:00:00"}}}`)
				}))
			})
			It(`Invoke ReplaceOffering successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.ReplaceOffering(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Rating model
				ratingModel := new(catalogmanagementv1.Rating)
				ratingModel.OneStarCount = core.Int64Ptr(int64(38))
				ratingModel.TwoStarCount = core.Int64Ptr(int64(38))
				ratingModel.ThreeStarCount = core.Int64Ptr(int64(38))
				ratingModel.FourStarCount = core.Int64Ptr(int64(38))

				// Construct an instance of the Feature model
				featureModel := new(catalogmanagementv1.Feature)
				featureModel.Title = core.StringPtr("testString")
				featureModel.Description = core.StringPtr("testString")

				// Construct an instance of the Configuration model
				configurationModel := new(catalogmanagementv1.Configuration)
				configurationModel.Key = core.StringPtr("testString")
				configurationModel.Type = core.StringPtr("testString")
				configurationModel.DefaultValue = core.StringPtr("testString")
				configurationModel.ValueConstraint = core.StringPtr("testString")
				configurationModel.Description = core.StringPtr("testString")
				configurationModel.Required = core.BoolPtr(true)
				configurationModel.Options = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				configurationModel.Hidden = core.BoolPtr(true)

				// Construct an instance of the Validation model
				validationModel := new(catalogmanagementv1.Validation)
				validationModel.Validated = CreateMockDateTime()
				validationModel.Requested = CreateMockDateTime()
				validationModel.State = core.StringPtr("testString")
				validationModel.LastOperation = core.StringPtr("testString")
				validationModel.Target = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the Resource model
				resourceModel := new(catalogmanagementv1.Resource)
				resourceModel.Type = core.StringPtr("mem")
				resourceModel.Value = core.StringPtr("testString")

				// Construct an instance of the Script model
				scriptModel := new(catalogmanagementv1.Script)
				scriptModel.Instructions = core.StringPtr("testString")
				scriptModel.Script = core.StringPtr("testString")
				scriptModel.ScriptPermission = core.StringPtr("testString")
				scriptModel.DeleteScript = core.StringPtr("testString")
				scriptModel.Scope = core.StringPtr("testString")

				// Construct an instance of the VersionEntitlement model
				versionEntitlementModel := new(catalogmanagementv1.VersionEntitlement)
				versionEntitlementModel.ProviderName = core.StringPtr("testString")
				versionEntitlementModel.ProviderID = core.StringPtr("testString")
				versionEntitlementModel.ProductID = core.StringPtr("testString")
				versionEntitlementModel.PartNumbers = []string{"testString"}
				versionEntitlementModel.ImageRepoName = core.StringPtr("testString")

				// Construct an instance of the License model
				licenseModel := new(catalogmanagementv1.License)
				licenseModel.ID = core.StringPtr("testString")
				licenseModel.Name = core.StringPtr("testString")
				licenseModel.Type = core.StringPtr("testString")
				licenseModel.URL = core.StringPtr("testString")
				licenseModel.Description = core.StringPtr("testString")

				// Construct an instance of the State model
				stateModel := new(catalogmanagementv1.State)
				stateModel.Current = core.StringPtr("testString")
				stateModel.CurrentEntered = CreateMockDateTime()
				stateModel.Pending = core.StringPtr("testString")
				stateModel.PendingRequested = CreateMockDateTime()
				stateModel.Previous = core.StringPtr("testString")

				// Construct an instance of the Version model
				versionModel := new(catalogmanagementv1.Version)
				versionModel.ID = core.StringPtr("testString")
				versionModel.Rev = core.StringPtr("testString")
				versionModel.Crn = core.StringPtr("testString")
				versionModel.Version = core.StringPtr("testString")
				versionModel.Sha = core.StringPtr("testString")
				versionModel.Created = CreateMockDateTime()
				versionModel.Updated = CreateMockDateTime()
				versionModel.OfferingID = core.StringPtr("testString")
				versionModel.CatalogID = core.StringPtr("testString")
				versionModel.KindID = core.StringPtr("testString")
				versionModel.Tags = []string{"testString"}
				versionModel.RepoURL = core.StringPtr("testString")
				versionModel.SourceURL = core.StringPtr("testString")
				versionModel.TgzURL = core.StringPtr("testString")
				versionModel.Configuration = []catalogmanagementv1.Configuration{*configurationModel}
				versionModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}
				versionModel.Validation = validationModel
				versionModel.RequiredResources = []catalogmanagementv1.Resource{*resourceModel}
				versionModel.SingleInstance = core.BoolPtr(true)
				versionModel.Install = scriptModel
				versionModel.PreInstall = []catalogmanagementv1.Script{*scriptModel}
				versionModel.Entitlement = versionEntitlementModel
				versionModel.Licenses = []catalogmanagementv1.License{*licenseModel}
				versionModel.ImageManifestURL = core.StringPtr("testString")
				versionModel.Deprecated = core.BoolPtr(true)
				versionModel.PackageVersion = core.StringPtr("testString")
				versionModel.State = stateModel
				versionModel.VersionLocator = core.StringPtr("testString")
				versionModel.ConsoleURL = core.StringPtr("testString")
				versionModel.LongDescription = core.StringPtr("testString")
				versionModel.WhitelistedAccounts = []string{"testString"}

				// Construct an instance of the Deployment model
				deploymentModel := new(catalogmanagementv1.Deployment)
				deploymentModel.ID = core.StringPtr("testString")
				deploymentModel.Label = core.StringPtr("testString")
				deploymentModel.Name = core.StringPtr("testString")
				deploymentModel.ShortDescription = core.StringPtr("testString")
				deploymentModel.LongDescription = core.StringPtr("testString")
				deploymentModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}
				deploymentModel.Tags = []string{"testString"}
				deploymentModel.Created = CreateMockDateTime()
				deploymentModel.Updated = CreateMockDateTime()

				// Construct an instance of the Plan model
				planModel := new(catalogmanagementv1.Plan)
				planModel.ID = core.StringPtr("testString")
				planModel.Label = core.StringPtr("testString")
				planModel.Name = core.StringPtr("testString")
				planModel.ShortDescription = core.StringPtr("testString")
				planModel.LongDescription = core.StringPtr("testString")
				planModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}
				planModel.Tags = []string{"testString"}
				planModel.AdditionalFeatures = []catalogmanagementv1.Feature{*featureModel}
				planModel.Created = CreateMockDateTime()
				planModel.Updated = CreateMockDateTime()
				planModel.Deployments = []catalogmanagementv1.Deployment{*deploymentModel}

				// Construct an instance of the Kind model
				kindModel := new(catalogmanagementv1.Kind)
				kindModel.ID = core.StringPtr("testString")
				kindModel.FormatKind = core.StringPtr("testString")
				kindModel.TargetKind = core.StringPtr("testString")
				kindModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}
				kindModel.InstallDescription = core.StringPtr("testString")
				kindModel.Tags = []string{"testString"}
				kindModel.AdditionalFeatures = []catalogmanagementv1.Feature{*featureModel}
				kindModel.Created = CreateMockDateTime()
				kindModel.Updated = CreateMockDateTime()
				kindModel.Versions = []catalogmanagementv1.Version{*versionModel}
				kindModel.Plans = []catalogmanagementv1.Plan{*planModel}

				// Construct an instance of the RepoInfo model
				repoInfoModel := new(catalogmanagementv1.RepoInfo)
				repoInfoModel.Token = core.StringPtr("testString")
				repoInfoModel.Type = core.StringPtr("testString")

				// Construct an instance of the ReplaceOfferingOptions model
				replaceOfferingOptionsModel := new(catalogmanagementv1.ReplaceOfferingOptions)
				replaceOfferingOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				replaceOfferingOptionsModel.OfferingID = core.StringPtr("testString")
				replaceOfferingOptionsModel.ID = core.StringPtr("testString")
				replaceOfferingOptionsModel.Rev = core.StringPtr("testString")
				replaceOfferingOptionsModel.URL = core.StringPtr("testString")
				replaceOfferingOptionsModel.Crn = core.StringPtr("testString")
				replaceOfferingOptionsModel.Label = core.StringPtr("testString")
				replaceOfferingOptionsModel.Name = core.StringPtr("testString")
				replaceOfferingOptionsModel.OfferingIconURL = core.StringPtr("testString")
				replaceOfferingOptionsModel.OfferingDocsURL = core.StringPtr("testString")
				replaceOfferingOptionsModel.OfferingSupportURL = core.StringPtr("testString")
				replaceOfferingOptionsModel.Tags = []string{"testString"}
				replaceOfferingOptionsModel.Rating = ratingModel
				replaceOfferingOptionsModel.Created = CreateMockDateTime()
				replaceOfferingOptionsModel.Updated = CreateMockDateTime()
				replaceOfferingOptionsModel.ShortDescription = core.StringPtr("testString")
				replaceOfferingOptionsModel.LongDescription = core.StringPtr("testString")
				replaceOfferingOptionsModel.Features = []catalogmanagementv1.Feature{*featureModel}
				replaceOfferingOptionsModel.Kinds = []catalogmanagementv1.Kind{*kindModel}
				replaceOfferingOptionsModel.PermitRequestIbmPublicPublish = core.BoolPtr(true)
				replaceOfferingOptionsModel.IbmPublishApproved = core.BoolPtr(true)
				replaceOfferingOptionsModel.PublicPublishApproved = core.BoolPtr(true)
				replaceOfferingOptionsModel.PublicOriginalCrn = core.StringPtr("testString")
				replaceOfferingOptionsModel.PublishPublicCrn = core.StringPtr("testString")
				replaceOfferingOptionsModel.PortalApprovalRecord = core.StringPtr("testString")
				replaceOfferingOptionsModel.PortalUiURL = core.StringPtr("testString")
				replaceOfferingOptionsModel.CatalogID = core.StringPtr("testString")
				replaceOfferingOptionsModel.CatalogName = core.StringPtr("testString")
				replaceOfferingOptionsModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}
				replaceOfferingOptionsModel.Disclaimer = core.StringPtr("testString")
				replaceOfferingOptionsModel.Hidden = core.BoolPtr(true)
				replaceOfferingOptionsModel.Provider = core.StringPtr("testString")
				replaceOfferingOptionsModel.RepoInfo = repoInfoModel
				replaceOfferingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.ReplaceOffering(replaceOfferingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.ReplaceOfferingWithContext(ctx, replaceOfferingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.ReplaceOffering(replaceOfferingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.ReplaceOfferingWithContext(ctx, replaceOfferingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ReplaceOffering with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the Rating model
				ratingModel := new(catalogmanagementv1.Rating)
				ratingModel.OneStarCount = core.Int64Ptr(int64(38))
				ratingModel.TwoStarCount = core.Int64Ptr(int64(38))
				ratingModel.ThreeStarCount = core.Int64Ptr(int64(38))
				ratingModel.FourStarCount = core.Int64Ptr(int64(38))

				// Construct an instance of the Feature model
				featureModel := new(catalogmanagementv1.Feature)
				featureModel.Title = core.StringPtr("testString")
				featureModel.Description = core.StringPtr("testString")

				// Construct an instance of the Configuration model
				configurationModel := new(catalogmanagementv1.Configuration)
				configurationModel.Key = core.StringPtr("testString")
				configurationModel.Type = core.StringPtr("testString")
				configurationModel.DefaultValue = core.StringPtr("testString")
				configurationModel.ValueConstraint = core.StringPtr("testString")
				configurationModel.Description = core.StringPtr("testString")
				configurationModel.Required = core.BoolPtr(true)
				configurationModel.Options = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				configurationModel.Hidden = core.BoolPtr(true)

				// Construct an instance of the Validation model
				validationModel := new(catalogmanagementv1.Validation)
				validationModel.Validated = CreateMockDateTime()
				validationModel.Requested = CreateMockDateTime()
				validationModel.State = core.StringPtr("testString")
				validationModel.LastOperation = core.StringPtr("testString")
				validationModel.Target = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the Resource model
				resourceModel := new(catalogmanagementv1.Resource)
				resourceModel.Type = core.StringPtr("mem")
				resourceModel.Value = core.StringPtr("testString")

				// Construct an instance of the Script model
				scriptModel := new(catalogmanagementv1.Script)
				scriptModel.Instructions = core.StringPtr("testString")
				scriptModel.Script = core.StringPtr("testString")
				scriptModel.ScriptPermission = core.StringPtr("testString")
				scriptModel.DeleteScript = core.StringPtr("testString")
				scriptModel.Scope = core.StringPtr("testString")

				// Construct an instance of the VersionEntitlement model
				versionEntitlementModel := new(catalogmanagementv1.VersionEntitlement)
				versionEntitlementModel.ProviderName = core.StringPtr("testString")
				versionEntitlementModel.ProviderID = core.StringPtr("testString")
				versionEntitlementModel.ProductID = core.StringPtr("testString")
				versionEntitlementModel.PartNumbers = []string{"testString"}
				versionEntitlementModel.ImageRepoName = core.StringPtr("testString")

				// Construct an instance of the License model
				licenseModel := new(catalogmanagementv1.License)
				licenseModel.ID = core.StringPtr("testString")
				licenseModel.Name = core.StringPtr("testString")
				licenseModel.Type = core.StringPtr("testString")
				licenseModel.URL = core.StringPtr("testString")
				licenseModel.Description = core.StringPtr("testString")

				// Construct an instance of the State model
				stateModel := new(catalogmanagementv1.State)
				stateModel.Current = core.StringPtr("testString")
				stateModel.CurrentEntered = CreateMockDateTime()
				stateModel.Pending = core.StringPtr("testString")
				stateModel.PendingRequested = CreateMockDateTime()
				stateModel.Previous = core.StringPtr("testString")

				// Construct an instance of the Version model
				versionModel := new(catalogmanagementv1.Version)
				versionModel.ID = core.StringPtr("testString")
				versionModel.Rev = core.StringPtr("testString")
				versionModel.Crn = core.StringPtr("testString")
				versionModel.Version = core.StringPtr("testString")
				versionModel.Sha = core.StringPtr("testString")
				versionModel.Created = CreateMockDateTime()
				versionModel.Updated = CreateMockDateTime()
				versionModel.OfferingID = core.StringPtr("testString")
				versionModel.CatalogID = core.StringPtr("testString")
				versionModel.KindID = core.StringPtr("testString")
				versionModel.Tags = []string{"testString"}
				versionModel.RepoURL = core.StringPtr("testString")
				versionModel.SourceURL = core.StringPtr("testString")
				versionModel.TgzURL = core.StringPtr("testString")
				versionModel.Configuration = []catalogmanagementv1.Configuration{*configurationModel}
				versionModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}
				versionModel.Validation = validationModel
				versionModel.RequiredResources = []catalogmanagementv1.Resource{*resourceModel}
				versionModel.SingleInstance = core.BoolPtr(true)
				versionModel.Install = scriptModel
				versionModel.PreInstall = []catalogmanagementv1.Script{*scriptModel}
				versionModel.Entitlement = versionEntitlementModel
				versionModel.Licenses = []catalogmanagementv1.License{*licenseModel}
				versionModel.ImageManifestURL = core.StringPtr("testString")
				versionModel.Deprecated = core.BoolPtr(true)
				versionModel.PackageVersion = core.StringPtr("testString")
				versionModel.State = stateModel
				versionModel.VersionLocator = core.StringPtr("testString")
				versionModel.ConsoleURL = core.StringPtr("testString")
				versionModel.LongDescription = core.StringPtr("testString")
				versionModel.WhitelistedAccounts = []string{"testString"}

				// Construct an instance of the Deployment model
				deploymentModel := new(catalogmanagementv1.Deployment)
				deploymentModel.ID = core.StringPtr("testString")
				deploymentModel.Label = core.StringPtr("testString")
				deploymentModel.Name = core.StringPtr("testString")
				deploymentModel.ShortDescription = core.StringPtr("testString")
				deploymentModel.LongDescription = core.StringPtr("testString")
				deploymentModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}
				deploymentModel.Tags = []string{"testString"}
				deploymentModel.Created = CreateMockDateTime()
				deploymentModel.Updated = CreateMockDateTime()

				// Construct an instance of the Plan model
				planModel := new(catalogmanagementv1.Plan)
				planModel.ID = core.StringPtr("testString")
				planModel.Label = core.StringPtr("testString")
				planModel.Name = core.StringPtr("testString")
				planModel.ShortDescription = core.StringPtr("testString")
				planModel.LongDescription = core.StringPtr("testString")
				planModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}
				planModel.Tags = []string{"testString"}
				planModel.AdditionalFeatures = []catalogmanagementv1.Feature{*featureModel}
				planModel.Created = CreateMockDateTime()
				planModel.Updated = CreateMockDateTime()
				planModel.Deployments = []catalogmanagementv1.Deployment{*deploymentModel}

				// Construct an instance of the Kind model
				kindModel := new(catalogmanagementv1.Kind)
				kindModel.ID = core.StringPtr("testString")
				kindModel.FormatKind = core.StringPtr("testString")
				kindModel.TargetKind = core.StringPtr("testString")
				kindModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}
				kindModel.InstallDescription = core.StringPtr("testString")
				kindModel.Tags = []string{"testString"}
				kindModel.AdditionalFeatures = []catalogmanagementv1.Feature{*featureModel}
				kindModel.Created = CreateMockDateTime()
				kindModel.Updated = CreateMockDateTime()
				kindModel.Versions = []catalogmanagementv1.Version{*versionModel}
				kindModel.Plans = []catalogmanagementv1.Plan{*planModel}

				// Construct an instance of the RepoInfo model
				repoInfoModel := new(catalogmanagementv1.RepoInfo)
				repoInfoModel.Token = core.StringPtr("testString")
				repoInfoModel.Type = core.StringPtr("testString")

				// Construct an instance of the ReplaceOfferingOptions model
				replaceOfferingOptionsModel := new(catalogmanagementv1.ReplaceOfferingOptions)
				replaceOfferingOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				replaceOfferingOptionsModel.OfferingID = core.StringPtr("testString")
				replaceOfferingOptionsModel.ID = core.StringPtr("testString")
				replaceOfferingOptionsModel.Rev = core.StringPtr("testString")
				replaceOfferingOptionsModel.URL = core.StringPtr("testString")
				replaceOfferingOptionsModel.Crn = core.StringPtr("testString")
				replaceOfferingOptionsModel.Label = core.StringPtr("testString")
				replaceOfferingOptionsModel.Name = core.StringPtr("testString")
				replaceOfferingOptionsModel.OfferingIconURL = core.StringPtr("testString")
				replaceOfferingOptionsModel.OfferingDocsURL = core.StringPtr("testString")
				replaceOfferingOptionsModel.OfferingSupportURL = core.StringPtr("testString")
				replaceOfferingOptionsModel.Tags = []string{"testString"}
				replaceOfferingOptionsModel.Rating = ratingModel
				replaceOfferingOptionsModel.Created = CreateMockDateTime()
				replaceOfferingOptionsModel.Updated = CreateMockDateTime()
				replaceOfferingOptionsModel.ShortDescription = core.StringPtr("testString")
				replaceOfferingOptionsModel.LongDescription = core.StringPtr("testString")
				replaceOfferingOptionsModel.Features = []catalogmanagementv1.Feature{*featureModel}
				replaceOfferingOptionsModel.Kinds = []catalogmanagementv1.Kind{*kindModel}
				replaceOfferingOptionsModel.PermitRequestIbmPublicPublish = core.BoolPtr(true)
				replaceOfferingOptionsModel.IbmPublishApproved = core.BoolPtr(true)
				replaceOfferingOptionsModel.PublicPublishApproved = core.BoolPtr(true)
				replaceOfferingOptionsModel.PublicOriginalCrn = core.StringPtr("testString")
				replaceOfferingOptionsModel.PublishPublicCrn = core.StringPtr("testString")
				replaceOfferingOptionsModel.PortalApprovalRecord = core.StringPtr("testString")
				replaceOfferingOptionsModel.PortalUiURL = core.StringPtr("testString")
				replaceOfferingOptionsModel.CatalogID = core.StringPtr("testString")
				replaceOfferingOptionsModel.CatalogName = core.StringPtr("testString")
				replaceOfferingOptionsModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}
				replaceOfferingOptionsModel.Disclaimer = core.StringPtr("testString")
				replaceOfferingOptionsModel.Hidden = core.BoolPtr(true)
				replaceOfferingOptionsModel.Provider = core.StringPtr("testString")
				replaceOfferingOptionsModel.RepoInfo = repoInfoModel
				replaceOfferingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.ReplaceOffering(replaceOfferingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceOfferingOptions model with no property values
				replaceOfferingOptionsModelNew := new(catalogmanagementv1.ReplaceOfferingOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = catalogManagementService.ReplaceOffering(replaceOfferingOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteOffering(deleteOfferingOptions *DeleteOfferingOptions)`, func() {
		deleteOfferingPath := "/catalogs/testString/offerings/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteOfferingPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteOffering successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := catalogManagementService.DeleteOffering(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteOfferingOptions model
				deleteOfferingOptionsModel := new(catalogmanagementv1.DeleteOfferingOptions)
				deleteOfferingOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				deleteOfferingOptionsModel.OfferingID = core.StringPtr("testString")
				deleteOfferingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = catalogManagementService.DeleteOffering(deleteOfferingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				response, operationErr = catalogManagementService.DeleteOffering(deleteOfferingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteOffering with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the DeleteOfferingOptions model
				deleteOfferingOptionsModel := new(catalogmanagementv1.DeleteOfferingOptions)
				deleteOfferingOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				deleteOfferingOptionsModel.OfferingID = core.StringPtr("testString")
				deleteOfferingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := catalogManagementService.DeleteOffering(deleteOfferingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteOfferingOptions model with no property values
				deleteOfferingOptionsModelNew := new(catalogmanagementv1.DeleteOfferingOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = catalogManagementService.DeleteOffering(deleteOfferingOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetOfferingAudit(getOfferingAuditOptions *GetOfferingAuditOptions)`, func() {
		getOfferingAuditPath := "/catalogs/testString/offerings/testString/audit"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getOfferingAuditPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["id"]).To(Equal([]string{"testString"}))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetOfferingAudit successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := catalogManagementService.GetOfferingAudit(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the GetOfferingAuditOptions model
				getOfferingAuditOptionsModel := new(catalogmanagementv1.GetOfferingAuditOptions)
				getOfferingAuditOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				getOfferingAuditOptionsModel.OfferingID = core.StringPtr("testString")
				getOfferingAuditOptionsModel.ID = core.StringPtr("testString")
				getOfferingAuditOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = catalogManagementService.GetOfferingAudit(getOfferingAuditOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				response, operationErr = catalogManagementService.GetOfferingAudit(getOfferingAuditOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke GetOfferingAudit with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetOfferingAuditOptions model
				getOfferingAuditOptionsModel := new(catalogmanagementv1.GetOfferingAuditOptions)
				getOfferingAuditOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				getOfferingAuditOptionsModel.OfferingID = core.StringPtr("testString")
				getOfferingAuditOptionsModel.ID = core.StringPtr("testString")
				getOfferingAuditOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := catalogManagementService.GetOfferingAudit(getOfferingAuditOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the GetOfferingAuditOptions model with no property values
				getOfferingAuditOptionsModelNew := new(catalogmanagementv1.GetOfferingAuditOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = catalogManagementService.GetOfferingAudit(getOfferingAuditOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceOfferingIcon(replaceOfferingIconOptions *ReplaceOfferingIconOptions) - Operation response error`, func() {
		replaceOfferingIconPath := "/catalogs/testString/offerings/testString/icon/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceOfferingIconPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceOfferingIcon with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the ReplaceOfferingIconOptions model
				replaceOfferingIconOptionsModel := new(catalogmanagementv1.ReplaceOfferingIconOptions)
				replaceOfferingIconOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				replaceOfferingIconOptionsModel.OfferingID = core.StringPtr("testString")
				replaceOfferingIconOptionsModel.FileName = core.StringPtr("testString")
				replaceOfferingIconOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.ReplaceOfferingIcon(replaceOfferingIconOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.ReplaceOfferingIcon(replaceOfferingIconOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ReplaceOfferingIcon(replaceOfferingIconOptions *ReplaceOfferingIconOptions)`, func() {
		replaceOfferingIconPath := "/catalogs/testString/offerings/testString/icon/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceOfferingIconPath))
					Expect(req.Method).To(Equal("PUT"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "_rev": "Rev", "url": "URL", "crn": "Crn", "label": "Label", "name": "Name", "offering_icon_url": "OfferingIconURL", "offering_docs_url": "OfferingDocsURL", "offering_support_url": "OfferingSupportURL", "tags": ["Tags"], "rating": {"one_star_count": 12, "two_star_count": 12, "three_star_count": 14, "four_star_count": 13}, "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "short_description": "ShortDescription", "long_description": "LongDescription", "features": [{"title": "Title", "description": "Description"}], "kinds": [{"id": "ID", "format_kind": "FormatKind", "target_kind": "TargetKind", "metadata": {"anyKey": "anyValue"}, "install_description": "InstallDescription", "tags": ["Tags"], "additional_features": [{"title": "Title", "description": "Description"}], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "versions": [{"id": "ID", "_rev": "Rev", "crn": "Crn", "version": "Version", "sha": "Sha", "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "offering_id": "OfferingID", "catalog_id": "CatalogID", "kind_id": "KindID", "tags": ["Tags"], "repo_url": "RepoURL", "source_url": "SourceURL", "tgz_url": "TgzURL", "configuration": [{"key": "Key", "type": "Type", "default_value": "anyValue", "value_constraint": "ValueConstraint", "description": "Description", "required": true, "options": [{"anyKey": "anyValue"}], "hidden": true}], "metadata": {"anyKey": "anyValue"}, "validation": {"validated": "2019-01-01T12:00:00", "requested": "2019-01-01T12:00:00", "state": "State", "last_operation": "LastOperation", "target": {"anyKey": "anyValue"}}, "required_resources": [{"type": "mem", "value": "anyValue"}], "single_instance": true, "install": {"instructions": "Instructions", "script": "Script", "script_permission": "ScriptPermission", "delete_script": "DeleteScript", "scope": "Scope"}, "pre_install": [{"instructions": "Instructions", "script": "Script", "script_permission": "ScriptPermission", "delete_script": "DeleteScript", "scope": "Scope"}], "entitlement": {"provider_name": "ProviderName", "provider_id": "ProviderID", "product_id": "ProductID", "part_numbers": ["PartNumbers"], "image_repo_name": "ImageRepoName"}, "licenses": [{"id": "ID", "name": "Name", "type": "Type", "url": "URL", "description": "Description"}], "image_manifest_url": "ImageManifestURL", "deprecated": true, "package_version": "PackageVersion", "state": {"current": "Current", "current_entered": "2019-01-01T12:00:00", "pending": "Pending", "pending_requested": "2019-01-01T12:00:00", "previous": "Previous"}, "version_locator": "VersionLocator", "console_url": "ConsoleURL", "long_description": "LongDescription", "whitelisted_accounts": ["WhitelistedAccounts"]}], "plans": [{"id": "ID", "label": "Label", "name": "Name", "short_description": "ShortDescription", "long_description": "LongDescription", "metadata": {"anyKey": "anyValue"}, "tags": ["Tags"], "additional_features": [{"title": "Title", "description": "Description"}], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "deployments": [{"id": "ID", "label": "Label", "name": "Name", "short_description": "ShortDescription", "long_description": "LongDescription", "metadata": {"anyKey": "anyValue"}, "tags": ["Tags"], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00"}]}]}], "permit_request_ibm_public_publish": false, "ibm_publish_approved": true, "public_publish_approved": false, "public_original_crn": "PublicOriginalCrn", "publish_public_crn": "PublishPublicCrn", "portal_approval_record": "PortalApprovalRecord", "portal_ui_url": "PortalUiURL", "catalog_id": "CatalogID", "catalog_name": "CatalogName", "metadata": {"anyKey": "anyValue"}, "disclaimer": "Disclaimer", "hidden": true, "provider": "Provider", "repo_info": {"token": "Token", "type": "Type"}}`)
				}))
			})
			It(`Invoke ReplaceOfferingIcon successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.ReplaceOfferingIcon(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ReplaceOfferingIconOptions model
				replaceOfferingIconOptionsModel := new(catalogmanagementv1.ReplaceOfferingIconOptions)
				replaceOfferingIconOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				replaceOfferingIconOptionsModel.OfferingID = core.StringPtr("testString")
				replaceOfferingIconOptionsModel.FileName = core.StringPtr("testString")
				replaceOfferingIconOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.ReplaceOfferingIcon(replaceOfferingIconOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.ReplaceOfferingIconWithContext(ctx, replaceOfferingIconOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.ReplaceOfferingIcon(replaceOfferingIconOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.ReplaceOfferingIconWithContext(ctx, replaceOfferingIconOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ReplaceOfferingIcon with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the ReplaceOfferingIconOptions model
				replaceOfferingIconOptionsModel := new(catalogmanagementv1.ReplaceOfferingIconOptions)
				replaceOfferingIconOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				replaceOfferingIconOptionsModel.OfferingID = core.StringPtr("testString")
				replaceOfferingIconOptionsModel.FileName = core.StringPtr("testString")
				replaceOfferingIconOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.ReplaceOfferingIcon(replaceOfferingIconOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceOfferingIconOptions model with no property values
				replaceOfferingIconOptionsModelNew := new(catalogmanagementv1.ReplaceOfferingIconOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = catalogManagementService.ReplaceOfferingIcon(replaceOfferingIconOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateOfferingIbm(updateOfferingIbmOptions *UpdateOfferingIbmOptions) - Operation response error`, func() {
		updateOfferingIbmPath := "/catalogs/testString/offerings/testString/publish/allow_request/true"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateOfferingIbmPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateOfferingIbm with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateOfferingIbmOptions model
				updateOfferingIbmOptionsModel := new(catalogmanagementv1.UpdateOfferingIbmOptions)
				updateOfferingIbmOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				updateOfferingIbmOptionsModel.OfferingID = core.StringPtr("testString")
				updateOfferingIbmOptionsModel.ApprovalType = core.StringPtr("allow_request")
				updateOfferingIbmOptionsModel.Approved = core.StringPtr("true")
				updateOfferingIbmOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.UpdateOfferingIbm(updateOfferingIbmOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.UpdateOfferingIbm(updateOfferingIbmOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateOfferingIbm(updateOfferingIbmOptions *UpdateOfferingIbmOptions)`, func() {
		updateOfferingIbmPath := "/catalogs/testString/offerings/testString/publish/allow_request/true"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateOfferingIbmPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"allow_request": true, "ibm": false, "public": true, "changed": false}`)
				}))
			})
			It(`Invoke UpdateOfferingIbm successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.UpdateOfferingIbm(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateOfferingIbmOptions model
				updateOfferingIbmOptionsModel := new(catalogmanagementv1.UpdateOfferingIbmOptions)
				updateOfferingIbmOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				updateOfferingIbmOptionsModel.OfferingID = core.StringPtr("testString")
				updateOfferingIbmOptionsModel.ApprovalType = core.StringPtr("allow_request")
				updateOfferingIbmOptionsModel.Approved = core.StringPtr("true")
				updateOfferingIbmOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.UpdateOfferingIbm(updateOfferingIbmOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.UpdateOfferingIbmWithContext(ctx, updateOfferingIbmOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.UpdateOfferingIbm(updateOfferingIbmOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.UpdateOfferingIbmWithContext(ctx, updateOfferingIbmOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateOfferingIbm with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateOfferingIbmOptions model
				updateOfferingIbmOptionsModel := new(catalogmanagementv1.UpdateOfferingIbmOptions)
				updateOfferingIbmOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				updateOfferingIbmOptionsModel.OfferingID = core.StringPtr("testString")
				updateOfferingIbmOptionsModel.ApprovalType = core.StringPtr("allow_request")
				updateOfferingIbmOptionsModel.Approved = core.StringPtr("true")
				updateOfferingIbmOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.UpdateOfferingIbm(updateOfferingIbmOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateOfferingIbmOptions model with no property values
				updateOfferingIbmOptionsModelNew := new(catalogmanagementv1.UpdateOfferingIbmOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = catalogManagementService.UpdateOfferingIbm(updateOfferingIbmOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(catalogManagementService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(catalogManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
				URL: "https://catalogmanagementv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(catalogManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CATALOG_MANAGEMENT_URL": "https://catalogmanagementv1/api",
				"CATALOG_MANAGEMENT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
				})
				Expect(catalogManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
					URL: "https://testService/api",
				})
				Expect(catalogManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
				})
				err := catalogManagementService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CATALOG_MANAGEMENT_URL": "https://catalogmanagementv1/api",
				"CATALOG_MANAGEMENT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(catalogManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CATALOG_MANAGEMENT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(catalogManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})

	Describe(`GetVersionAbout(getVersionAboutOptions *GetVersionAboutOptions)`, func() {
		getVersionAboutPath := "/versions/testString/about"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getVersionAboutPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "text/markdown")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `"OperationResponse"`)
				}))
			})
			It(`Invoke GetVersionAbout successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.GetVersionAbout(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetVersionAboutOptions model
				getVersionAboutOptionsModel := new(catalogmanagementv1.GetVersionAboutOptions)
				getVersionAboutOptionsModel.VersionLocID = core.StringPtr("testString")
				getVersionAboutOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.GetVersionAbout(getVersionAboutOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetVersionAboutWithContext(ctx, getVersionAboutOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.GetVersionAbout(getVersionAboutOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetVersionAboutWithContext(ctx, getVersionAboutOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetVersionAbout with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetVersionAboutOptions model
				getVersionAboutOptionsModel := new(catalogmanagementv1.GetVersionAboutOptions)
				getVersionAboutOptionsModel.VersionLocID = core.StringPtr("testString")
				getVersionAboutOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.GetVersionAbout(getVersionAboutOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetVersionAboutOptions model with no property values
				getVersionAboutOptionsModelNew := new(catalogmanagementv1.GetVersionAboutOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = catalogManagementService.GetVersionAbout(getVersionAboutOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetVersionLicense(getVersionLicenseOptions *GetVersionLicenseOptions)`, func() {
		getVersionLicensePath := "/versions/testString/licenses/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getVersionLicensePath))
					Expect(req.Method).To(Equal("GET"))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetVersionLicense successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := catalogManagementService.GetVersionLicense(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the GetVersionLicenseOptions model
				getVersionLicenseOptionsModel := new(catalogmanagementv1.GetVersionLicenseOptions)
				getVersionLicenseOptionsModel.VersionLocID = core.StringPtr("testString")
				getVersionLicenseOptionsModel.LicenseID = core.StringPtr("testString")
				getVersionLicenseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = catalogManagementService.GetVersionLicense(getVersionLicenseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				response, operationErr = catalogManagementService.GetVersionLicense(getVersionLicenseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke GetVersionLicense with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetVersionLicenseOptions model
				getVersionLicenseOptionsModel := new(catalogmanagementv1.GetVersionLicenseOptions)
				getVersionLicenseOptionsModel.VersionLocID = core.StringPtr("testString")
				getVersionLicenseOptionsModel.LicenseID = core.StringPtr("testString")
				getVersionLicenseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := catalogManagementService.GetVersionLicense(getVersionLicenseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the GetVersionLicenseOptions model with no property values
				getVersionLicenseOptionsModelNew := new(catalogmanagementv1.GetVersionLicenseOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = catalogManagementService.GetVersionLicense(getVersionLicenseOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetVersionContainerImages(getVersionContainerImagesOptions *GetVersionContainerImagesOptions) - Operation response error`, func() {
		getVersionContainerImagesPath := "/versions/testString/containerImages"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getVersionContainerImagesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetVersionContainerImages with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetVersionContainerImagesOptions model
				getVersionContainerImagesOptionsModel := new(catalogmanagementv1.GetVersionContainerImagesOptions)
				getVersionContainerImagesOptionsModel.VersionLocID = core.StringPtr("testString")
				getVersionContainerImagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.GetVersionContainerImages(getVersionContainerImagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.GetVersionContainerImages(getVersionContainerImagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetVersionContainerImages(getVersionContainerImagesOptions *GetVersionContainerImagesOptions)`, func() {
		getVersionContainerImagesPath := "/versions/testString/containerImages"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getVersionContainerImagesPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"description": "Description", "images": [{"image": "Image"}]}`)
				}))
			})
			It(`Invoke GetVersionContainerImages successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.GetVersionContainerImages(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetVersionContainerImagesOptions model
				getVersionContainerImagesOptionsModel := new(catalogmanagementv1.GetVersionContainerImagesOptions)
				getVersionContainerImagesOptionsModel.VersionLocID = core.StringPtr("testString")
				getVersionContainerImagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.GetVersionContainerImages(getVersionContainerImagesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetVersionContainerImagesWithContext(ctx, getVersionContainerImagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.GetVersionContainerImages(getVersionContainerImagesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetVersionContainerImagesWithContext(ctx, getVersionContainerImagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetVersionContainerImages with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetVersionContainerImagesOptions model
				getVersionContainerImagesOptionsModel := new(catalogmanagementv1.GetVersionContainerImagesOptions)
				getVersionContainerImagesOptionsModel.VersionLocID = core.StringPtr("testString")
				getVersionContainerImagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.GetVersionContainerImages(getVersionContainerImagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetVersionContainerImagesOptions model with no property values
				getVersionContainerImagesOptionsModelNew := new(catalogmanagementv1.GetVersionContainerImagesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = catalogManagementService.GetVersionContainerImages(getVersionContainerImagesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeprecateVersion(deprecateVersionOptions *DeprecateVersionOptions)`, func() {
		deprecateVersionPath := "/versions/testString/deprecate"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deprecateVersionPath))
					Expect(req.Method).To(Equal("POST"))

					res.WriteHeader(202)
				}))
			})
			It(`Invoke DeprecateVersion successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := catalogManagementService.DeprecateVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeprecateVersionOptions model
				deprecateVersionOptionsModel := new(catalogmanagementv1.DeprecateVersionOptions)
				deprecateVersionOptionsModel.VersionLocID = core.StringPtr("testString")
				deprecateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = catalogManagementService.DeprecateVersion(deprecateVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				response, operationErr = catalogManagementService.DeprecateVersion(deprecateVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeprecateVersion with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the DeprecateVersionOptions model
				deprecateVersionOptionsModel := new(catalogmanagementv1.DeprecateVersionOptions)
				deprecateVersionOptionsModel.VersionLocID = core.StringPtr("testString")
				deprecateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := catalogManagementService.DeprecateVersion(deprecateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeprecateVersionOptions model with no property values
				deprecateVersionOptionsModelNew := new(catalogmanagementv1.DeprecateVersionOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = catalogManagementService.DeprecateVersion(deprecateVersionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`AccountPublishVersion(accountPublishVersionOptions *AccountPublishVersionOptions)`, func() {
		accountPublishVersionPath := "/versions/testString/account-publish"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(accountPublishVersionPath))
					Expect(req.Method).To(Equal("POST"))

					res.WriteHeader(202)
				}))
			})
			It(`Invoke AccountPublishVersion successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := catalogManagementService.AccountPublishVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the AccountPublishVersionOptions model
				accountPublishVersionOptionsModel := new(catalogmanagementv1.AccountPublishVersionOptions)
				accountPublishVersionOptionsModel.VersionLocID = core.StringPtr("testString")
				accountPublishVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = catalogManagementService.AccountPublishVersion(accountPublishVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				response, operationErr = catalogManagementService.AccountPublishVersion(accountPublishVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke AccountPublishVersion with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the AccountPublishVersionOptions model
				accountPublishVersionOptionsModel := new(catalogmanagementv1.AccountPublishVersionOptions)
				accountPublishVersionOptionsModel.VersionLocID = core.StringPtr("testString")
				accountPublishVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := catalogManagementService.AccountPublishVersion(accountPublishVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the AccountPublishVersionOptions model with no property values
				accountPublishVersionOptionsModelNew := new(catalogmanagementv1.AccountPublishVersionOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = catalogManagementService.AccountPublishVersion(accountPublishVersionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`IbmPublishVersion(ibmPublishVersionOptions *IbmPublishVersionOptions)`, func() {
		ibmPublishVersionPath := "/versions/testString/ibm-publish"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(ibmPublishVersionPath))
					Expect(req.Method).To(Equal("POST"))

					res.WriteHeader(202)
				}))
			})
			It(`Invoke IbmPublishVersion successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := catalogManagementService.IbmPublishVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the IbmPublishVersionOptions model
				ibmPublishVersionOptionsModel := new(catalogmanagementv1.IbmPublishVersionOptions)
				ibmPublishVersionOptionsModel.VersionLocID = core.StringPtr("testString")
				ibmPublishVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = catalogManagementService.IbmPublishVersion(ibmPublishVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				response, operationErr = catalogManagementService.IbmPublishVersion(ibmPublishVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke IbmPublishVersion with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the IbmPublishVersionOptions model
				ibmPublishVersionOptionsModel := new(catalogmanagementv1.IbmPublishVersionOptions)
				ibmPublishVersionOptionsModel.VersionLocID = core.StringPtr("testString")
				ibmPublishVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := catalogManagementService.IbmPublishVersion(ibmPublishVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the IbmPublishVersionOptions model with no property values
				ibmPublishVersionOptionsModelNew := new(catalogmanagementv1.IbmPublishVersionOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = catalogManagementService.IbmPublishVersion(ibmPublishVersionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PublicPublishVersion(publicPublishVersionOptions *PublicPublishVersionOptions)`, func() {
		publicPublishVersionPath := "/versions/testString/public-publish"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(publicPublishVersionPath))
					Expect(req.Method).To(Equal("POST"))

					res.WriteHeader(202)
				}))
			})
			It(`Invoke PublicPublishVersion successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := catalogManagementService.PublicPublishVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the PublicPublishVersionOptions model
				publicPublishVersionOptionsModel := new(catalogmanagementv1.PublicPublishVersionOptions)
				publicPublishVersionOptionsModel.VersionLocID = core.StringPtr("testString")
				publicPublishVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = catalogManagementService.PublicPublishVersion(publicPublishVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				response, operationErr = catalogManagementService.PublicPublishVersion(publicPublishVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke PublicPublishVersion with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the PublicPublishVersionOptions model
				publicPublishVersionOptionsModel := new(catalogmanagementv1.PublicPublishVersionOptions)
				publicPublishVersionOptionsModel.VersionLocID = core.StringPtr("testString")
				publicPublishVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := catalogManagementService.PublicPublishVersion(publicPublishVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the PublicPublishVersionOptions model with no property values
				publicPublishVersionOptionsModelNew := new(catalogmanagementv1.PublicPublishVersionOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = catalogManagementService.PublicPublishVersion(publicPublishVersionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CommitVersion(commitVersionOptions *CommitVersionOptions)`, func() {
		commitVersionPath := "/versions/testString/commit"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(commitVersionPath))
					Expect(req.Method).To(Equal("POST"))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke CommitVersion successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := catalogManagementService.CommitVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the CommitVersionOptions model
				commitVersionOptionsModel := new(catalogmanagementv1.CommitVersionOptions)
				commitVersionOptionsModel.VersionLocID = core.StringPtr("testString")
				commitVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = catalogManagementService.CommitVersion(commitVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				response, operationErr = catalogManagementService.CommitVersion(commitVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke CommitVersion with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the CommitVersionOptions model
				commitVersionOptionsModel := new(catalogmanagementv1.CommitVersionOptions)
				commitVersionOptionsModel.VersionLocID = core.StringPtr("testString")
				commitVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := catalogManagementService.CommitVersion(commitVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the CommitVersionOptions model with no property values
				commitVersionOptionsModelNew := new(catalogmanagementv1.CommitVersionOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = catalogManagementService.CommitVersion(commitVersionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CopyVersion(copyVersionOptions *CopyVersionOptions)`, func() {
		copyVersionPath := "/versions/testString/copy"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(copyVersionPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					res.WriteHeader(200)
				}))
			})
			It(`Invoke CopyVersion successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := catalogManagementService.CopyVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the CopyVersionOptions model
				copyVersionOptionsModel := new(catalogmanagementv1.CopyVersionOptions)
				copyVersionOptionsModel.VersionLocID = core.StringPtr("testString")
				copyVersionOptionsModel.Tags = []string{"testString"}
				copyVersionOptionsModel.TargetKinds = []string{"testString"}
				copyVersionOptionsModel.Content = []int64{int64(38)}
				copyVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = catalogManagementService.CopyVersion(copyVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				response, operationErr = catalogManagementService.CopyVersion(copyVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke CopyVersion with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the CopyVersionOptions model
				copyVersionOptionsModel := new(catalogmanagementv1.CopyVersionOptions)
				copyVersionOptionsModel.VersionLocID = core.StringPtr("testString")
				copyVersionOptionsModel.Tags = []string{"testString"}
				copyVersionOptionsModel.TargetKinds = []string{"testString"}
				copyVersionOptionsModel.Content = []int64{int64(38)}
				copyVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := catalogManagementService.CopyVersion(copyVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the CopyVersionOptions model with no property values
				copyVersionOptionsModelNew := new(catalogmanagementv1.CopyVersionOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = catalogManagementService.CopyVersion(copyVersionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetVersionWorkingCopy(getVersionWorkingCopyOptions *GetVersionWorkingCopyOptions) - Operation response error`, func() {
		getVersionWorkingCopyPath := "/versions/testString/workingcopy"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getVersionWorkingCopyPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetVersionWorkingCopy with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetVersionWorkingCopyOptions model
				getVersionWorkingCopyOptionsModel := new(catalogmanagementv1.GetVersionWorkingCopyOptions)
				getVersionWorkingCopyOptionsModel.VersionLocID = core.StringPtr("testString")
				getVersionWorkingCopyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.GetVersionWorkingCopy(getVersionWorkingCopyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.GetVersionWorkingCopy(getVersionWorkingCopyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetVersionWorkingCopy(getVersionWorkingCopyOptions *GetVersionWorkingCopyOptions)`, func() {
		getVersionWorkingCopyPath := "/versions/testString/workingcopy"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getVersionWorkingCopyPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "_rev": "Rev", "crn": "Crn", "version": "Version", "sha": "Sha", "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "offering_id": "OfferingID", "catalog_id": "CatalogID", "kind_id": "KindID", "tags": ["Tags"], "repo_url": "RepoURL", "source_url": "SourceURL", "tgz_url": "TgzURL", "configuration": [{"key": "Key", "type": "Type", "default_value": "anyValue", "value_constraint": "ValueConstraint", "description": "Description", "required": true, "options": [{"anyKey": "anyValue"}], "hidden": true}], "metadata": {"anyKey": "anyValue"}, "validation": {"validated": "2019-01-01T12:00:00", "requested": "2019-01-01T12:00:00", "state": "State", "last_operation": "LastOperation", "target": {"anyKey": "anyValue"}}, "required_resources": [{"type": "mem", "value": "anyValue"}], "single_instance": true, "install": {"instructions": "Instructions", "script": "Script", "script_permission": "ScriptPermission", "delete_script": "DeleteScript", "scope": "Scope"}, "pre_install": [{"instructions": "Instructions", "script": "Script", "script_permission": "ScriptPermission", "delete_script": "DeleteScript", "scope": "Scope"}], "entitlement": {"provider_name": "ProviderName", "provider_id": "ProviderID", "product_id": "ProductID", "part_numbers": ["PartNumbers"], "image_repo_name": "ImageRepoName"}, "licenses": [{"id": "ID", "name": "Name", "type": "Type", "url": "URL", "description": "Description"}], "image_manifest_url": "ImageManifestURL", "deprecated": true, "package_version": "PackageVersion", "state": {"current": "Current", "current_entered": "2019-01-01T12:00:00", "pending": "Pending", "pending_requested": "2019-01-01T12:00:00", "previous": "Previous"}, "version_locator": "VersionLocator", "console_url": "ConsoleURL", "long_description": "LongDescription", "whitelisted_accounts": ["WhitelistedAccounts"]}`)
				}))
			})
			It(`Invoke GetVersionWorkingCopy successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.GetVersionWorkingCopy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetVersionWorkingCopyOptions model
				getVersionWorkingCopyOptionsModel := new(catalogmanagementv1.GetVersionWorkingCopyOptions)
				getVersionWorkingCopyOptionsModel.VersionLocID = core.StringPtr("testString")
				getVersionWorkingCopyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.GetVersionWorkingCopy(getVersionWorkingCopyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetVersionWorkingCopyWithContext(ctx, getVersionWorkingCopyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.GetVersionWorkingCopy(getVersionWorkingCopyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetVersionWorkingCopyWithContext(ctx, getVersionWorkingCopyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetVersionWorkingCopy with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetVersionWorkingCopyOptions model
				getVersionWorkingCopyOptionsModel := new(catalogmanagementv1.GetVersionWorkingCopyOptions)
				getVersionWorkingCopyOptionsModel.VersionLocID = core.StringPtr("testString")
				getVersionWorkingCopyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.GetVersionWorkingCopy(getVersionWorkingCopyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetVersionWorkingCopyOptions model with no property values
				getVersionWorkingCopyOptionsModelNew := new(catalogmanagementv1.GetVersionWorkingCopyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = catalogManagementService.GetVersionWorkingCopy(getVersionWorkingCopyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetVersionUpdates(getVersionUpdatesOptions *GetVersionUpdatesOptions) - Operation response error`, func() {
		getVersionUpdatesPath := "/versions/testString/updates"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getVersionUpdatesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["cluster_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["region"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["namespace"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetVersionUpdates with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetVersionUpdatesOptions model
				getVersionUpdatesOptionsModel := new(catalogmanagementv1.GetVersionUpdatesOptions)
				getVersionUpdatesOptionsModel.VersionLocID = core.StringPtr("testString")
				getVersionUpdatesOptionsModel.ClusterID = core.StringPtr("testString")
				getVersionUpdatesOptionsModel.Region = core.StringPtr("testString")
				getVersionUpdatesOptionsModel.ResourceGroupID = core.StringPtr("testString")
				getVersionUpdatesOptionsModel.Namespace = core.StringPtr("testString")
				getVersionUpdatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.GetVersionUpdates(getVersionUpdatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.GetVersionUpdates(getVersionUpdatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetVersionUpdates(getVersionUpdatesOptions *GetVersionUpdatesOptions)`, func() {
		getVersionUpdatesPath := "/versions/testString/updates"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getVersionUpdatesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["cluster_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["region"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["namespace"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `[{"version_locator": "VersionLocator", "version": "Version", "state": {"current": "Current", "current_entered": "2019-01-01T12:00:00", "pending": "Pending", "pending_requested": "2019-01-01T12:00:00", "previous": "Previous"}, "required_resources": [{"type": "mem", "value": "anyValue"}], "package_version": "PackageVersion", "can_update": false, "messages": {"anyKey": "anyValue"}}]`)
				}))
			})
			It(`Invoke GetVersionUpdates successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.GetVersionUpdates(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetVersionUpdatesOptions model
				getVersionUpdatesOptionsModel := new(catalogmanagementv1.GetVersionUpdatesOptions)
				getVersionUpdatesOptionsModel.VersionLocID = core.StringPtr("testString")
				getVersionUpdatesOptionsModel.ClusterID = core.StringPtr("testString")
				getVersionUpdatesOptionsModel.Region = core.StringPtr("testString")
				getVersionUpdatesOptionsModel.ResourceGroupID = core.StringPtr("testString")
				getVersionUpdatesOptionsModel.Namespace = core.StringPtr("testString")
				getVersionUpdatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.GetVersionUpdates(getVersionUpdatesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetVersionUpdatesWithContext(ctx, getVersionUpdatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.GetVersionUpdates(getVersionUpdatesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetVersionUpdatesWithContext(ctx, getVersionUpdatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetVersionUpdates with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetVersionUpdatesOptions model
				getVersionUpdatesOptionsModel := new(catalogmanagementv1.GetVersionUpdatesOptions)
				getVersionUpdatesOptionsModel.VersionLocID = core.StringPtr("testString")
				getVersionUpdatesOptionsModel.ClusterID = core.StringPtr("testString")
				getVersionUpdatesOptionsModel.Region = core.StringPtr("testString")
				getVersionUpdatesOptionsModel.ResourceGroupID = core.StringPtr("testString")
				getVersionUpdatesOptionsModel.Namespace = core.StringPtr("testString")
				getVersionUpdatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.GetVersionUpdates(getVersionUpdatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetVersionUpdatesOptions model with no property values
				getVersionUpdatesOptionsModelNew := new(catalogmanagementv1.GetVersionUpdatesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = catalogManagementService.GetVersionUpdates(getVersionUpdatesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetVersion(getVersionOptions *GetVersionOptions) - Operation response error`, func() {
		getVersionPath := "/versions/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getVersionPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetVersion with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetVersionOptions model
				getVersionOptionsModel := new(catalogmanagementv1.GetVersionOptions)
				getVersionOptionsModel.VersionLocID = core.StringPtr("testString")
				getVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.GetVersion(getVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.GetVersion(getVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetVersion(getVersionOptions *GetVersionOptions)`, func() {
		getVersionPath := "/versions/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getVersionPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "_rev": "Rev", "url": "URL", "crn": "Crn", "label": "Label", "name": "Name", "offering_icon_url": "OfferingIconURL", "offering_docs_url": "OfferingDocsURL", "offering_support_url": "OfferingSupportURL", "tags": ["Tags"], "rating": {"one_star_count": 12, "two_star_count": 12, "three_star_count": 14, "four_star_count": 13}, "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "short_description": "ShortDescription", "long_description": "LongDescription", "features": [{"title": "Title", "description": "Description"}], "kinds": [{"id": "ID", "format_kind": "FormatKind", "target_kind": "TargetKind", "metadata": {"anyKey": "anyValue"}, "install_description": "InstallDescription", "tags": ["Tags"], "additional_features": [{"title": "Title", "description": "Description"}], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "versions": [{"id": "ID", "_rev": "Rev", "crn": "Crn", "version": "Version", "sha": "Sha", "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "offering_id": "OfferingID", "catalog_id": "CatalogID", "kind_id": "KindID", "tags": ["Tags"], "repo_url": "RepoURL", "source_url": "SourceURL", "tgz_url": "TgzURL", "configuration": [{"key": "Key", "type": "Type", "default_value": "anyValue", "value_constraint": "ValueConstraint", "description": "Description", "required": true, "options": [{"anyKey": "anyValue"}], "hidden": true}], "metadata": {"anyKey": "anyValue"}, "validation": {"validated": "2019-01-01T12:00:00", "requested": "2019-01-01T12:00:00", "state": "State", "last_operation": "LastOperation", "target": {"anyKey": "anyValue"}}, "required_resources": [{"type": "mem", "value": "anyValue"}], "single_instance": true, "install": {"instructions": "Instructions", "script": "Script", "script_permission": "ScriptPermission", "delete_script": "DeleteScript", "scope": "Scope"}, "pre_install": [{"instructions": "Instructions", "script": "Script", "script_permission": "ScriptPermission", "delete_script": "DeleteScript", "scope": "Scope"}], "entitlement": {"provider_name": "ProviderName", "provider_id": "ProviderID", "product_id": "ProductID", "part_numbers": ["PartNumbers"], "image_repo_name": "ImageRepoName"}, "licenses": [{"id": "ID", "name": "Name", "type": "Type", "url": "URL", "description": "Description"}], "image_manifest_url": "ImageManifestURL", "deprecated": true, "package_version": "PackageVersion", "state": {"current": "Current", "current_entered": "2019-01-01T12:00:00", "pending": "Pending", "pending_requested": "2019-01-01T12:00:00", "previous": "Previous"}, "version_locator": "VersionLocator", "console_url": "ConsoleURL", "long_description": "LongDescription", "whitelisted_accounts": ["WhitelistedAccounts"]}], "plans": [{"id": "ID", "label": "Label", "name": "Name", "short_description": "ShortDescription", "long_description": "LongDescription", "metadata": {"anyKey": "anyValue"}, "tags": ["Tags"], "additional_features": [{"title": "Title", "description": "Description"}], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "deployments": [{"id": "ID", "label": "Label", "name": "Name", "short_description": "ShortDescription", "long_description": "LongDescription", "metadata": {"anyKey": "anyValue"}, "tags": ["Tags"], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00"}]}]}], "permit_request_ibm_public_publish": false, "ibm_publish_approved": true, "public_publish_approved": false, "public_original_crn": "PublicOriginalCrn", "publish_public_crn": "PublishPublicCrn", "portal_approval_record": "PortalApprovalRecord", "portal_ui_url": "PortalUiURL", "catalog_id": "CatalogID", "catalog_name": "CatalogName", "metadata": {"anyKey": "anyValue"}, "disclaimer": "Disclaimer", "hidden": true, "provider": "Provider", "repo_info": {"token": "Token", "type": "Type"}}`)
				}))
			})
			It(`Invoke GetVersion successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.GetVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetVersionOptions model
				getVersionOptionsModel := new(catalogmanagementv1.GetVersionOptions)
				getVersionOptionsModel.VersionLocID = core.StringPtr("testString")
				getVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.GetVersion(getVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetVersionWithContext(ctx, getVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.GetVersion(getVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetVersionWithContext(ctx, getVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetVersion with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetVersionOptions model
				getVersionOptionsModel := new(catalogmanagementv1.GetVersionOptions)
				getVersionOptionsModel.VersionLocID = core.StringPtr("testString")
				getVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.GetVersion(getVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetVersionOptions model with no property values
				getVersionOptionsModelNew := new(catalogmanagementv1.GetVersionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = catalogManagementService.GetVersion(getVersionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteVersion(deleteVersionOptions *DeleteVersionOptions)`, func() {
		deleteVersionPath := "/versions/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteVersionPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteVersion successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := catalogManagementService.DeleteVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteVersionOptions model
				deleteVersionOptionsModel := new(catalogmanagementv1.DeleteVersionOptions)
				deleteVersionOptionsModel.VersionLocID = core.StringPtr("testString")
				deleteVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = catalogManagementService.DeleteVersion(deleteVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				response, operationErr = catalogManagementService.DeleteVersion(deleteVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteVersion with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the DeleteVersionOptions model
				deleteVersionOptionsModel := new(catalogmanagementv1.DeleteVersionOptions)
				deleteVersionOptionsModel.VersionLocID = core.StringPtr("testString")
				deleteVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := catalogManagementService.DeleteVersion(deleteVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteVersionOptions model with no property values
				deleteVersionOptionsModelNew := new(catalogmanagementv1.DeleteVersionOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = catalogManagementService.DeleteVersion(deleteVersionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListVersions(listVersionsOptions *ListVersionsOptions)`, func() {
		listVersionsPath := "/versions"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listVersionsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["q"]).To(Equal([]string{"testString"}))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListVersions successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := catalogManagementService.ListVersions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the ListVersionsOptions model
				listVersionsOptionsModel := new(catalogmanagementv1.ListVersionsOptions)
				listVersionsOptionsModel.Q = core.StringPtr("testString")
				listVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = catalogManagementService.ListVersions(listVersionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				response, operationErr = catalogManagementService.ListVersions(listVersionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke ListVersions with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the ListVersionsOptions model
				listVersionsOptionsModel := new(catalogmanagementv1.ListVersionsOptions)
				listVersionsOptionsModel.Q = core.StringPtr("testString")
				listVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := catalogManagementService.ListVersions(listVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the ListVersionsOptions model with no property values
				listVersionsOptionsModelNew := new(catalogmanagementv1.ListVersionsOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = catalogManagementService.ListVersions(listVersionsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(catalogManagementService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(catalogManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
				URL: "https://catalogmanagementv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(catalogManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CATALOG_MANAGEMENT_URL": "https://catalogmanagementv1/api",
				"CATALOG_MANAGEMENT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
				})
				Expect(catalogManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
					URL: "https://testService/api",
				})
				Expect(catalogManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
				})
				err := catalogManagementService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CATALOG_MANAGEMENT_URL": "https://catalogmanagementv1/api",
				"CATALOG_MANAGEMENT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(catalogManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CATALOG_MANAGEMENT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(catalogManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`GetRepos(getReposOptions *GetReposOptions) - Operation response error`, func() {
		getReposPath := "/repo/testString/entries"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReposPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["repourl"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetRepos with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetReposOptions model
				getReposOptionsModel := new(catalogmanagementv1.GetReposOptions)
				getReposOptionsModel.Type = core.StringPtr("testString")
				getReposOptionsModel.Repourl = core.StringPtr("testString")
				getReposOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.GetRepos(getReposOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.GetRepos(getReposOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetRepos(getReposOptions *GetReposOptions)`, func() {
		getReposPath := "/repo/testString/entries"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReposPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["repourl"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"chart": {"api_version": "ApiVersion", "created": "2019-01-01T12:00:00", "description": "Description", "deprecated": true, "digest": "Digest", "home": "Home", "icon": "Icon", "keywords": ["Keywords"], "maintainers": [{"email": "Email", "name": "Name"}], "name": "Name", "tiller_version": "TillerVersion", "urls": ["Urls"], "sources": ["Sources"], "version": "Version", "appVersion": "AppVersion"}}`)
				}))
			})
			It(`Invoke GetRepos successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.GetRepos(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetReposOptions model
				getReposOptionsModel := new(catalogmanagementv1.GetReposOptions)
				getReposOptionsModel.Type = core.StringPtr("testString")
				getReposOptionsModel.Repourl = core.StringPtr("testString")
				getReposOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.GetRepos(getReposOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetReposWithContext(ctx, getReposOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.GetRepos(getReposOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetReposWithContext(ctx, getReposOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetRepos with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetReposOptions model
				getReposOptionsModel := new(catalogmanagementv1.GetReposOptions)
				getReposOptionsModel.Type = core.StringPtr("testString")
				getReposOptionsModel.Repourl = core.StringPtr("testString")
				getReposOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.GetRepos(getReposOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetReposOptions model with no property values
				getReposOptionsModelNew := new(catalogmanagementv1.GetReposOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = catalogManagementService.GetRepos(getReposOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetRepo(getRepoOptions *GetRepoOptions) - Operation response error`, func() {
		getRepoPath := "/repo/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRepoPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["charturl"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetRepo with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetRepoOptions model
				getRepoOptionsModel := new(catalogmanagementv1.GetRepoOptions)
				getRepoOptionsModel.Type = core.StringPtr("testString")
				getRepoOptionsModel.Charturl = core.StringPtr("testString")
				getRepoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.GetRepo(getRepoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.GetRepo(getRepoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetRepo(getRepoOptions *GetRepoOptions)`, func() {
		getRepoPath := "/repo/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRepoPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["charturl"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"chart": {"Chart.yaml": {"name": "Name", "description": "Description", "icon": "Icon", "version": "Version", "appVersion": "AppVersion"}, "sha": {"anyKey": "anyValue"}, "README.md": "READMEMd", "values-metadata": {"anyKey": "anyValue"}, "license-metadata": {"anyKey": "anyValue"}}}`)
				}))
			})
			It(`Invoke GetRepo successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.GetRepo(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetRepoOptions model
				getRepoOptionsModel := new(catalogmanagementv1.GetRepoOptions)
				getRepoOptionsModel.Type = core.StringPtr("testString")
				getRepoOptionsModel.Charturl = core.StringPtr("testString")
				getRepoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.GetRepo(getRepoOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetRepoWithContext(ctx, getRepoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.GetRepo(getRepoOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetRepoWithContext(ctx, getRepoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetRepo with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetRepoOptions model
				getRepoOptionsModel := new(catalogmanagementv1.GetRepoOptions)
				getRepoOptionsModel.Type = core.StringPtr("testString")
				getRepoOptionsModel.Charturl = core.StringPtr("testString")
				getRepoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.GetRepo(getRepoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetRepoOptions model with no property values
				getRepoOptionsModelNew := new(catalogmanagementv1.GetRepoOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = catalogManagementService.GetRepo(getRepoOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(catalogManagementService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(catalogManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
				URL: "https://catalogmanagementv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(catalogManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CATALOG_MANAGEMENT_URL": "https://catalogmanagementv1/api",
				"CATALOG_MANAGEMENT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
				})
				Expect(catalogManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
					URL: "https://testService/api",
				})
				Expect(catalogManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
				})
				err := catalogManagementService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CATALOG_MANAGEMENT_URL": "https://catalogmanagementv1/api",
				"CATALOG_MANAGEMENT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(catalogManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CATALOG_MANAGEMENT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(catalogManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListClusters(listClustersOptions *ListClustersOptions) - Operation response error`, func() {
		listClustersPath := "/deploy/kubernetes/clusters"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listClustersPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["type"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListClusters with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the ListClustersOptions model
				listClustersOptionsModel := new(catalogmanagementv1.ListClustersOptions)
				listClustersOptionsModel.Limit = core.Int64Ptr(int64(38))
				listClustersOptionsModel.Offset = core.Int64Ptr(int64(38))
				listClustersOptionsModel.Type = core.StringPtr("testString")
				listClustersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.ListClusters(listClustersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.ListClusters(listClustersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListClusters(listClustersOptions *ListClustersOptions)`, func() {
		listClustersPath := "/deploy/kubernetes/clusters"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listClustersPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["type"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 6, "limit": 5, "total_count": 10, "resource_count": 13, "first": "First", "last": "Last", "prev": "Prev", "next": "Next", "resources": [{"resource_group_id": "ResourceGroupID", "resource_group_name": "ResourceGroupName", "id": "ID", "name": "Name", "region": "Region"}]}`)
				}))
			})
			It(`Invoke ListClusters successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.ListClusters(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListClustersOptions model
				listClustersOptionsModel := new(catalogmanagementv1.ListClustersOptions)
				listClustersOptionsModel.Limit = core.Int64Ptr(int64(38))
				listClustersOptionsModel.Offset = core.Int64Ptr(int64(38))
				listClustersOptionsModel.Type = core.StringPtr("testString")
				listClustersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.ListClusters(listClustersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.ListClustersWithContext(ctx, listClustersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.ListClusters(listClustersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.ListClustersWithContext(ctx, listClustersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListClusters with error: Operation request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the ListClustersOptions model
				listClustersOptionsModel := new(catalogmanagementv1.ListClustersOptions)
				listClustersOptionsModel.Limit = core.Int64Ptr(int64(38))
				listClustersOptionsModel.Offset = core.Int64Ptr(int64(38))
				listClustersOptionsModel.Type = core.StringPtr("testString")
				listClustersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.ListClusters(listClustersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCluster(getClusterOptions *GetClusterOptions) - Operation response error`, func() {
		getClusterPath := "/deploy/kubernetes/clusters/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getClusterPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Auth-Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["region"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCluster with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetClusterOptions model
				getClusterOptionsModel := new(catalogmanagementv1.GetClusterOptions)
				getClusterOptionsModel.ClusterID = core.StringPtr("testString")
				getClusterOptionsModel.Region = core.StringPtr("testString")
				getClusterOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				getClusterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.GetCluster(getClusterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.GetCluster(getClusterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetCluster(getClusterOptions *GetClusterOptions)`, func() {
		getClusterPath := "/deploy/kubernetes/clusters/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getClusterPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Auth-Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["region"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"resource_group_id": "ResourceGroupID", "resource_group_name": "ResourceGroupName", "id": "ID", "name": "Name", "region": "Region"}`)
				}))
			})
			It(`Invoke GetCluster successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.GetCluster(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetClusterOptions model
				getClusterOptionsModel := new(catalogmanagementv1.GetClusterOptions)
				getClusterOptionsModel.ClusterID = core.StringPtr("testString")
				getClusterOptionsModel.Region = core.StringPtr("testString")
				getClusterOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				getClusterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.GetCluster(getClusterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetClusterWithContext(ctx, getClusterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.GetCluster(getClusterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetClusterWithContext(ctx, getClusterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetCluster with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetClusterOptions model
				getClusterOptionsModel := new(catalogmanagementv1.GetClusterOptions)
				getClusterOptionsModel.ClusterID = core.StringPtr("testString")
				getClusterOptionsModel.Region = core.StringPtr("testString")
				getClusterOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				getClusterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.GetCluster(getClusterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetClusterOptions model with no property values
				getClusterOptionsModelNew := new(catalogmanagementv1.GetClusterOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = catalogManagementService.GetCluster(getClusterOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetNamespaces(getNamespacesOptions *GetNamespacesOptions) - Operation response error`, func() {
		getNamespacesPath := "/deploy/kubernetes/clusters/testString/namespaces"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getNamespacesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Auth-Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["region"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetNamespaces with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetNamespacesOptions model
				getNamespacesOptionsModel := new(catalogmanagementv1.GetNamespacesOptions)
				getNamespacesOptionsModel.ClusterID = core.StringPtr("testString")
				getNamespacesOptionsModel.Region = core.StringPtr("testString")
				getNamespacesOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				getNamespacesOptionsModel.Limit = core.Int64Ptr(int64(38))
				getNamespacesOptionsModel.Offset = core.Int64Ptr(int64(38))
				getNamespacesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.GetNamespaces(getNamespacesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.GetNamespaces(getNamespacesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetNamespaces(getNamespacesOptions *GetNamespacesOptions)`, func() {
		getNamespacesPath := "/deploy/kubernetes/clusters/testString/namespaces"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getNamespacesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Auth-Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["region"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 6, "limit": 5, "total_count": 10, "resource_count": 13, "first": "First", "last": "Last", "prev": "Prev", "next": "Next", "resources": ["Resources"]}`)
				}))
			})
			It(`Invoke GetNamespaces successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.GetNamespaces(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetNamespacesOptions model
				getNamespacesOptionsModel := new(catalogmanagementv1.GetNamespacesOptions)
				getNamespacesOptionsModel.ClusterID = core.StringPtr("testString")
				getNamespacesOptionsModel.Region = core.StringPtr("testString")
				getNamespacesOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				getNamespacesOptionsModel.Limit = core.Int64Ptr(int64(38))
				getNamespacesOptionsModel.Offset = core.Int64Ptr(int64(38))
				getNamespacesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.GetNamespaces(getNamespacesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetNamespacesWithContext(ctx, getNamespacesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.GetNamespaces(getNamespacesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetNamespacesWithContext(ctx, getNamespacesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetNamespaces with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetNamespacesOptions model
				getNamespacesOptionsModel := new(catalogmanagementv1.GetNamespacesOptions)
				getNamespacesOptionsModel.ClusterID = core.StringPtr("testString")
				getNamespacesOptionsModel.Region = core.StringPtr("testString")
				getNamespacesOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				getNamespacesOptionsModel.Limit = core.Int64Ptr(int64(38))
				getNamespacesOptionsModel.Offset = core.Int64Ptr(int64(38))
				getNamespacesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.GetNamespaces(getNamespacesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetNamespacesOptions model with no property values
				getNamespacesOptionsModelNew := new(catalogmanagementv1.GetNamespacesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = catalogManagementService.GetNamespaces(getNamespacesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateOperator(createOperatorOptions *CreateOperatorOptions) - Operation response error`, func() {
		createOperatorPath := "/deploy/kubernetes/olm/operator"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createOperatorPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Auth-Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateOperator with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the CreateOperatorOptions model
				createOperatorOptionsModel := new(catalogmanagementv1.CreateOperatorOptions)
				createOperatorOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				createOperatorOptionsModel.ClusterID = core.StringPtr("testString")
				createOperatorOptionsModel.Region = core.StringPtr("testString")
				createOperatorOptionsModel.Namespaces = []string{"testString"}
				createOperatorOptionsModel.AllNamespaces = core.BoolPtr(true)
				createOperatorOptionsModel.VersionLocatorID = core.StringPtr("testString")
				createOperatorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.CreateOperator(createOperatorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.CreateOperator(createOperatorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateOperator(createOperatorOptions *CreateOperatorOptions)`, func() {
		createOperatorPath := "/deploy/kubernetes/olm/operator"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createOperatorPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["X-Auth-Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `[{"phase": "Phase", "message": "Message", "link": "Link", "name": "Name", "version": "Version", "namespace": "Namespace", "package_name": "PackageName", "catalog_id": "CatalogID"}]`)
				}))
			})
			It(`Invoke CreateOperator successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.CreateOperator(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateOperatorOptions model
				createOperatorOptionsModel := new(catalogmanagementv1.CreateOperatorOptions)
				createOperatorOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				createOperatorOptionsModel.ClusterID = core.StringPtr("testString")
				createOperatorOptionsModel.Region = core.StringPtr("testString")
				createOperatorOptionsModel.Namespaces = []string{"testString"}
				createOperatorOptionsModel.AllNamespaces = core.BoolPtr(true)
				createOperatorOptionsModel.VersionLocatorID = core.StringPtr("testString")
				createOperatorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.CreateOperator(createOperatorOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.CreateOperatorWithContext(ctx, createOperatorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.CreateOperator(createOperatorOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.CreateOperatorWithContext(ctx, createOperatorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke CreateOperator with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the CreateOperatorOptions model
				createOperatorOptionsModel := new(catalogmanagementv1.CreateOperatorOptions)
				createOperatorOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				createOperatorOptionsModel.ClusterID = core.StringPtr("testString")
				createOperatorOptionsModel.Region = core.StringPtr("testString")
				createOperatorOptionsModel.Namespaces = []string{"testString"}
				createOperatorOptionsModel.AllNamespaces = core.BoolPtr(true)
				createOperatorOptionsModel.VersionLocatorID = core.StringPtr("testString")
				createOperatorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.CreateOperator(createOperatorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateOperatorOptions model with no property values
				createOperatorOptionsModelNew := new(catalogmanagementv1.CreateOperatorOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = catalogManagementService.CreateOperator(createOperatorOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListOperators(listOperatorsOptions *ListOperatorsOptions) - Operation response error`, func() {
		listOperatorsPath := "/deploy/kubernetes/olm/operator"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listOperatorsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Auth-Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["cluster_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["region"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["version_locator_id"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListOperators with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the ListOperatorsOptions model
				listOperatorsOptionsModel := new(catalogmanagementv1.ListOperatorsOptions)
				listOperatorsOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				listOperatorsOptionsModel.ClusterID = core.StringPtr("testString")
				listOperatorsOptionsModel.Region = core.StringPtr("testString")
				listOperatorsOptionsModel.VersionLocatorID = core.StringPtr("testString")
				listOperatorsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.ListOperators(listOperatorsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.ListOperators(listOperatorsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListOperators(listOperatorsOptions *ListOperatorsOptions)`, func() {
		listOperatorsPath := "/deploy/kubernetes/olm/operator"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listOperatorsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Auth-Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["cluster_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["region"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["version_locator_id"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `[{"phase": "Phase", "message": "Message", "link": "Link", "name": "Name", "version": "Version", "namespace": "Namespace", "package_name": "PackageName", "catalog_id": "CatalogID"}]`)
				}))
			})
			It(`Invoke ListOperators successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.ListOperators(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListOperatorsOptions model
				listOperatorsOptionsModel := new(catalogmanagementv1.ListOperatorsOptions)
				listOperatorsOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				listOperatorsOptionsModel.ClusterID = core.StringPtr("testString")
				listOperatorsOptionsModel.Region = core.StringPtr("testString")
				listOperatorsOptionsModel.VersionLocatorID = core.StringPtr("testString")
				listOperatorsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.ListOperators(listOperatorsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.ListOperatorsWithContext(ctx, listOperatorsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.ListOperators(listOperatorsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.ListOperatorsWithContext(ctx, listOperatorsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListOperators with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the ListOperatorsOptions model
				listOperatorsOptionsModel := new(catalogmanagementv1.ListOperatorsOptions)
				listOperatorsOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				listOperatorsOptionsModel.ClusterID = core.StringPtr("testString")
				listOperatorsOptionsModel.Region = core.StringPtr("testString")
				listOperatorsOptionsModel.VersionLocatorID = core.StringPtr("testString")
				listOperatorsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.ListOperators(listOperatorsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListOperatorsOptions model with no property values
				listOperatorsOptionsModelNew := new(catalogmanagementv1.ListOperatorsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = catalogManagementService.ListOperators(listOperatorsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceOperator(replaceOperatorOptions *ReplaceOperatorOptions) - Operation response error`, func() {
		replaceOperatorPath := "/deploy/kubernetes/olm/operator"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceOperatorPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["X-Auth-Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceOperator with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the ReplaceOperatorOptions model
				replaceOperatorOptionsModel := new(catalogmanagementv1.ReplaceOperatorOptions)
				replaceOperatorOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				replaceOperatorOptionsModel.ClusterID = core.StringPtr("testString")
				replaceOperatorOptionsModel.Region = core.StringPtr("testString")
				replaceOperatorOptionsModel.Namespaces = []string{"testString"}
				replaceOperatorOptionsModel.AllNamespaces = core.BoolPtr(true)
				replaceOperatorOptionsModel.VersionLocatorID = core.StringPtr("testString")
				replaceOperatorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.ReplaceOperator(replaceOperatorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.ReplaceOperator(replaceOperatorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ReplaceOperator(replaceOperatorOptions *ReplaceOperatorOptions)`, func() {
		replaceOperatorPath := "/deploy/kubernetes/olm/operator"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceOperatorPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["X-Auth-Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `[{"phase": "Phase", "message": "Message", "link": "Link", "name": "Name", "version": "Version", "namespace": "Namespace", "package_name": "PackageName", "catalog_id": "CatalogID"}]`)
				}))
			})
			It(`Invoke ReplaceOperator successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.ReplaceOperator(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ReplaceOperatorOptions model
				replaceOperatorOptionsModel := new(catalogmanagementv1.ReplaceOperatorOptions)
				replaceOperatorOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				replaceOperatorOptionsModel.ClusterID = core.StringPtr("testString")
				replaceOperatorOptionsModel.Region = core.StringPtr("testString")
				replaceOperatorOptionsModel.Namespaces = []string{"testString"}
				replaceOperatorOptionsModel.AllNamespaces = core.BoolPtr(true)
				replaceOperatorOptionsModel.VersionLocatorID = core.StringPtr("testString")
				replaceOperatorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.ReplaceOperator(replaceOperatorOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.ReplaceOperatorWithContext(ctx, replaceOperatorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.ReplaceOperator(replaceOperatorOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.ReplaceOperatorWithContext(ctx, replaceOperatorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ReplaceOperator with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the ReplaceOperatorOptions model
				replaceOperatorOptionsModel := new(catalogmanagementv1.ReplaceOperatorOptions)
				replaceOperatorOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				replaceOperatorOptionsModel.ClusterID = core.StringPtr("testString")
				replaceOperatorOptionsModel.Region = core.StringPtr("testString")
				replaceOperatorOptionsModel.Namespaces = []string{"testString"}
				replaceOperatorOptionsModel.AllNamespaces = core.BoolPtr(true)
				replaceOperatorOptionsModel.VersionLocatorID = core.StringPtr("testString")
				replaceOperatorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.ReplaceOperator(replaceOperatorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceOperatorOptions model with no property values
				replaceOperatorOptionsModelNew := new(catalogmanagementv1.ReplaceOperatorOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = catalogManagementService.ReplaceOperator(replaceOperatorOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteOperator(deleteOperatorOptions *DeleteOperatorOptions)`, func() {
		deleteOperatorPath := "/deploy/kubernetes/olm/operator"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteOperatorPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["X-Auth-Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["cluster_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["region"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["version_locator_id"]).To(Equal([]string{"testString"}))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteOperator successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := catalogManagementService.DeleteOperator(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteOperatorOptions model
				deleteOperatorOptionsModel := new(catalogmanagementv1.DeleteOperatorOptions)
				deleteOperatorOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				deleteOperatorOptionsModel.ClusterID = core.StringPtr("testString")
				deleteOperatorOptionsModel.Region = core.StringPtr("testString")
				deleteOperatorOptionsModel.VersionLocatorID = core.StringPtr("testString")
				deleteOperatorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = catalogManagementService.DeleteOperator(deleteOperatorOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				response, operationErr = catalogManagementService.DeleteOperator(deleteOperatorOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteOperator with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the DeleteOperatorOptions model
				deleteOperatorOptionsModel := new(catalogmanagementv1.DeleteOperatorOptions)
				deleteOperatorOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				deleteOperatorOptionsModel.ClusterID = core.StringPtr("testString")
				deleteOperatorOptionsModel.Region = core.StringPtr("testString")
				deleteOperatorOptionsModel.VersionLocatorID = core.StringPtr("testString")
				deleteOperatorOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := catalogManagementService.DeleteOperator(deleteOperatorOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteOperatorOptions model with no property values
				deleteOperatorOptionsModelNew := new(catalogmanagementv1.DeleteOperatorOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = catalogManagementService.DeleteOperator(deleteOperatorOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`InstallVersion(installVersionOptions *InstallVersionOptions)`, func() {
		installVersionPath := "/versions/testString/install"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(installVersionPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["X-Auth-Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(202)
				}))
			})
			It(`Invoke InstallVersion successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := catalogManagementService.InstallVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeployRequestBodySchematics model
				deployRequestBodySchematicsModel := new(catalogmanagementv1.DeployRequestBodySchematics)
				deployRequestBodySchematicsModel.Name = core.StringPtr("testString")
				deployRequestBodySchematicsModel.Description = core.StringPtr("testString")
				deployRequestBodySchematicsModel.Tags = []string{"testString"}
				deployRequestBodySchematicsModel.ResourceGroupID = core.StringPtr("testString")

				// Construct an instance of the InstallVersionOptions model
				installVersionOptionsModel := new(catalogmanagementv1.InstallVersionOptions)
				installVersionOptionsModel.VersionLocID = core.StringPtr("testString")
				installVersionOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				installVersionOptionsModel.ClusterID = core.StringPtr("testString")
				installVersionOptionsModel.Region = core.StringPtr("testString")
				installVersionOptionsModel.Namespace = core.StringPtr("testString")
				installVersionOptionsModel.OverrideValues = map[string]interface{}{"anyKey": "anyValue"}
				installVersionOptionsModel.EntitlementApikey = core.StringPtr("testString")
				installVersionOptionsModel.Schematics = deployRequestBodySchematicsModel
				installVersionOptionsModel.Script = core.StringPtr("testString")
				installVersionOptionsModel.ScriptID = core.StringPtr("testString")
				installVersionOptionsModel.VersionLocatorID = core.StringPtr("testString")
				installVersionOptionsModel.VcenterID = core.StringPtr("testString")
				installVersionOptionsModel.VcenterUser = core.StringPtr("testString")
				installVersionOptionsModel.VcenterPassword = core.StringPtr("testString")
				installVersionOptionsModel.VcenterLocation = core.StringPtr("testString")
				installVersionOptionsModel.VcenterDatastore = core.StringPtr("testString")
				installVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = catalogManagementService.InstallVersion(installVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				response, operationErr = catalogManagementService.InstallVersion(installVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke InstallVersion with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the DeployRequestBodySchematics model
				deployRequestBodySchematicsModel := new(catalogmanagementv1.DeployRequestBodySchematics)
				deployRequestBodySchematicsModel.Name = core.StringPtr("testString")
				deployRequestBodySchematicsModel.Description = core.StringPtr("testString")
				deployRequestBodySchematicsModel.Tags = []string{"testString"}
				deployRequestBodySchematicsModel.ResourceGroupID = core.StringPtr("testString")

				// Construct an instance of the InstallVersionOptions model
				installVersionOptionsModel := new(catalogmanagementv1.InstallVersionOptions)
				installVersionOptionsModel.VersionLocID = core.StringPtr("testString")
				installVersionOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				installVersionOptionsModel.ClusterID = core.StringPtr("testString")
				installVersionOptionsModel.Region = core.StringPtr("testString")
				installVersionOptionsModel.Namespace = core.StringPtr("testString")
				installVersionOptionsModel.OverrideValues = map[string]interface{}{"anyKey": "anyValue"}
				installVersionOptionsModel.EntitlementApikey = core.StringPtr("testString")
				installVersionOptionsModel.Schematics = deployRequestBodySchematicsModel
				installVersionOptionsModel.Script = core.StringPtr("testString")
				installVersionOptionsModel.ScriptID = core.StringPtr("testString")
				installVersionOptionsModel.VersionLocatorID = core.StringPtr("testString")
				installVersionOptionsModel.VcenterID = core.StringPtr("testString")
				installVersionOptionsModel.VcenterUser = core.StringPtr("testString")
				installVersionOptionsModel.VcenterPassword = core.StringPtr("testString")
				installVersionOptionsModel.VcenterLocation = core.StringPtr("testString")
				installVersionOptionsModel.VcenterDatastore = core.StringPtr("testString")
				installVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := catalogManagementService.InstallVersion(installVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the InstallVersionOptions model with no property values
				installVersionOptionsModelNew := new(catalogmanagementv1.InstallVersionOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = catalogManagementService.InstallVersion(installVersionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PreinstallVersion(preinstallVersionOptions *PreinstallVersionOptions)`, func() {
		preinstallVersionPath := "/versions/testString/preinstall"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(preinstallVersionPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["X-Auth-Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(202)
				}))
			})
			It(`Invoke PreinstallVersion successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := catalogManagementService.PreinstallVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeployRequestBodySchematics model
				deployRequestBodySchematicsModel := new(catalogmanagementv1.DeployRequestBodySchematics)
				deployRequestBodySchematicsModel.Name = core.StringPtr("testString")
				deployRequestBodySchematicsModel.Description = core.StringPtr("testString")
				deployRequestBodySchematicsModel.Tags = []string{"testString"}
				deployRequestBodySchematicsModel.ResourceGroupID = core.StringPtr("testString")

				// Construct an instance of the PreinstallVersionOptions model
				preinstallVersionOptionsModel := new(catalogmanagementv1.PreinstallVersionOptions)
				preinstallVersionOptionsModel.VersionLocID = core.StringPtr("testString")
				preinstallVersionOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				preinstallVersionOptionsModel.ClusterID = core.StringPtr("testString")
				preinstallVersionOptionsModel.Region = core.StringPtr("testString")
				preinstallVersionOptionsModel.Namespace = core.StringPtr("testString")
				preinstallVersionOptionsModel.OverrideValues = map[string]interface{}{"anyKey": "anyValue"}
				preinstallVersionOptionsModel.EntitlementApikey = core.StringPtr("testString")
				preinstallVersionOptionsModel.Schematics = deployRequestBodySchematicsModel
				preinstallVersionOptionsModel.Script = core.StringPtr("testString")
				preinstallVersionOptionsModel.ScriptID = core.StringPtr("testString")
				preinstallVersionOptionsModel.VersionLocatorID = core.StringPtr("testString")
				preinstallVersionOptionsModel.VcenterID = core.StringPtr("testString")
				preinstallVersionOptionsModel.VcenterUser = core.StringPtr("testString")
				preinstallVersionOptionsModel.VcenterPassword = core.StringPtr("testString")
				preinstallVersionOptionsModel.VcenterLocation = core.StringPtr("testString")
				preinstallVersionOptionsModel.VcenterDatastore = core.StringPtr("testString")
				preinstallVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = catalogManagementService.PreinstallVersion(preinstallVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				response, operationErr = catalogManagementService.PreinstallVersion(preinstallVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke PreinstallVersion with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the DeployRequestBodySchematics model
				deployRequestBodySchematicsModel := new(catalogmanagementv1.DeployRequestBodySchematics)
				deployRequestBodySchematicsModel.Name = core.StringPtr("testString")
				deployRequestBodySchematicsModel.Description = core.StringPtr("testString")
				deployRequestBodySchematicsModel.Tags = []string{"testString"}
				deployRequestBodySchematicsModel.ResourceGroupID = core.StringPtr("testString")

				// Construct an instance of the PreinstallVersionOptions model
				preinstallVersionOptionsModel := new(catalogmanagementv1.PreinstallVersionOptions)
				preinstallVersionOptionsModel.VersionLocID = core.StringPtr("testString")
				preinstallVersionOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				preinstallVersionOptionsModel.ClusterID = core.StringPtr("testString")
				preinstallVersionOptionsModel.Region = core.StringPtr("testString")
				preinstallVersionOptionsModel.Namespace = core.StringPtr("testString")
				preinstallVersionOptionsModel.OverrideValues = map[string]interface{}{"anyKey": "anyValue"}
				preinstallVersionOptionsModel.EntitlementApikey = core.StringPtr("testString")
				preinstallVersionOptionsModel.Schematics = deployRequestBodySchematicsModel
				preinstallVersionOptionsModel.Script = core.StringPtr("testString")
				preinstallVersionOptionsModel.ScriptID = core.StringPtr("testString")
				preinstallVersionOptionsModel.VersionLocatorID = core.StringPtr("testString")
				preinstallVersionOptionsModel.VcenterID = core.StringPtr("testString")
				preinstallVersionOptionsModel.VcenterUser = core.StringPtr("testString")
				preinstallVersionOptionsModel.VcenterPassword = core.StringPtr("testString")
				preinstallVersionOptionsModel.VcenterLocation = core.StringPtr("testString")
				preinstallVersionOptionsModel.VcenterDatastore = core.StringPtr("testString")
				preinstallVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := catalogManagementService.PreinstallVersion(preinstallVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the PreinstallVersionOptions model with no property values
				preinstallVersionOptionsModelNew := new(catalogmanagementv1.PreinstallVersionOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = catalogManagementService.PreinstallVersion(preinstallVersionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetPreinstall(getPreinstallOptions *GetPreinstallOptions) - Operation response error`, func() {
		getPreinstallPath := "/versions/testString/preinstall"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPreinstallPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Auth-Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["cluster_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["region"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["namespace"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetPreinstall with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetPreinstallOptions model
				getPreinstallOptionsModel := new(catalogmanagementv1.GetPreinstallOptions)
				getPreinstallOptionsModel.VersionLocID = core.StringPtr("testString")
				getPreinstallOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				getPreinstallOptionsModel.ClusterID = core.StringPtr("testString")
				getPreinstallOptionsModel.Region = core.StringPtr("testString")
				getPreinstallOptionsModel.Namespace = core.StringPtr("testString")
				getPreinstallOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.GetPreinstall(getPreinstallOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.GetPreinstall(getPreinstallOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetPreinstall(getPreinstallOptions *GetPreinstallOptions)`, func() {
		getPreinstallPath := "/versions/testString/preinstall"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPreinstallPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Auth-Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["cluster_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["region"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["namespace"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"metadata": {"cluster_id": "ClusterID", "region": "Region", "namespace": "Namespace", "workspace_id": "WorkspaceID", "workspace_name": "WorkspaceName"}, "release": {"deployments": [{"anyKey": "anyValue"}], "replicasets": [{"anyKey": "anyValue"}], "statefulsets": [{"anyKey": "anyValue"}], "pods": [{"anyKey": "anyValue"}], "errors": [{"anyKey": "anyValue"}]}, "content_mgmt": {"pods": [{"anyKey": "anyValue"}], "errors": [{"anyKey": "anyValue"}]}}`)
				}))
			})
			It(`Invoke GetPreinstall successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.GetPreinstall(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetPreinstallOptions model
				getPreinstallOptionsModel := new(catalogmanagementv1.GetPreinstallOptions)
				getPreinstallOptionsModel.VersionLocID = core.StringPtr("testString")
				getPreinstallOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				getPreinstallOptionsModel.ClusterID = core.StringPtr("testString")
				getPreinstallOptionsModel.Region = core.StringPtr("testString")
				getPreinstallOptionsModel.Namespace = core.StringPtr("testString")
				getPreinstallOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.GetPreinstall(getPreinstallOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetPreinstallWithContext(ctx, getPreinstallOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.GetPreinstall(getPreinstallOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetPreinstallWithContext(ctx, getPreinstallOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetPreinstall with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetPreinstallOptions model
				getPreinstallOptionsModel := new(catalogmanagementv1.GetPreinstallOptions)
				getPreinstallOptionsModel.VersionLocID = core.StringPtr("testString")
				getPreinstallOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				getPreinstallOptionsModel.ClusterID = core.StringPtr("testString")
				getPreinstallOptionsModel.Region = core.StringPtr("testString")
				getPreinstallOptionsModel.Namespace = core.StringPtr("testString")
				getPreinstallOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.GetPreinstall(getPreinstallOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetPreinstallOptions model with no property values
				getPreinstallOptionsModelNew := new(catalogmanagementv1.GetPreinstallOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = catalogManagementService.GetPreinstall(getPreinstallOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ValidationInstall(validationInstallOptions *ValidationInstallOptions)`, func() {
		validationInstallPath := "/versions/testString/validation/install"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(validationInstallPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["X-Auth-Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(202)
				}))
			})
			It(`Invoke ValidationInstall successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := catalogManagementService.ValidationInstall(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeployRequestBodySchematics model
				deployRequestBodySchematicsModel := new(catalogmanagementv1.DeployRequestBodySchematics)
				deployRequestBodySchematicsModel.Name = core.StringPtr("testString")
				deployRequestBodySchematicsModel.Description = core.StringPtr("testString")
				deployRequestBodySchematicsModel.Tags = []string{"testString"}
				deployRequestBodySchematicsModel.ResourceGroupID = core.StringPtr("testString")

				// Construct an instance of the ValidationInstallOptions model
				validationInstallOptionsModel := new(catalogmanagementv1.ValidationInstallOptions)
				validationInstallOptionsModel.VersionLocID = core.StringPtr("testString")
				validationInstallOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				validationInstallOptionsModel.ClusterID = core.StringPtr("testString")
				validationInstallOptionsModel.Region = core.StringPtr("testString")
				validationInstallOptionsModel.Namespace = core.StringPtr("testString")
				validationInstallOptionsModel.OverrideValues = map[string]interface{}{"anyKey": "anyValue"}
				validationInstallOptionsModel.EntitlementApikey = core.StringPtr("testString")
				validationInstallOptionsModel.Schematics = deployRequestBodySchematicsModel
				validationInstallOptionsModel.Script = core.StringPtr("testString")
				validationInstallOptionsModel.ScriptID = core.StringPtr("testString")
				validationInstallOptionsModel.VersionLocatorID = core.StringPtr("testString")
				validationInstallOptionsModel.VcenterID = core.StringPtr("testString")
				validationInstallOptionsModel.VcenterUser = core.StringPtr("testString")
				validationInstallOptionsModel.VcenterPassword = core.StringPtr("testString")
				validationInstallOptionsModel.VcenterLocation = core.StringPtr("testString")
				validationInstallOptionsModel.VcenterDatastore = core.StringPtr("testString")
				validationInstallOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = catalogManagementService.ValidationInstall(validationInstallOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				response, operationErr = catalogManagementService.ValidationInstall(validationInstallOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke ValidationInstall with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the DeployRequestBodySchematics model
				deployRequestBodySchematicsModel := new(catalogmanagementv1.DeployRequestBodySchematics)
				deployRequestBodySchematicsModel.Name = core.StringPtr("testString")
				deployRequestBodySchematicsModel.Description = core.StringPtr("testString")
				deployRequestBodySchematicsModel.Tags = []string{"testString"}
				deployRequestBodySchematicsModel.ResourceGroupID = core.StringPtr("testString")

				// Construct an instance of the ValidationInstallOptions model
				validationInstallOptionsModel := new(catalogmanagementv1.ValidationInstallOptions)
				validationInstallOptionsModel.VersionLocID = core.StringPtr("testString")
				validationInstallOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				validationInstallOptionsModel.ClusterID = core.StringPtr("testString")
				validationInstallOptionsModel.Region = core.StringPtr("testString")
				validationInstallOptionsModel.Namespace = core.StringPtr("testString")
				validationInstallOptionsModel.OverrideValues = map[string]interface{}{"anyKey": "anyValue"}
				validationInstallOptionsModel.EntitlementApikey = core.StringPtr("testString")
				validationInstallOptionsModel.Schematics = deployRequestBodySchematicsModel
				validationInstallOptionsModel.Script = core.StringPtr("testString")
				validationInstallOptionsModel.ScriptID = core.StringPtr("testString")
				validationInstallOptionsModel.VersionLocatorID = core.StringPtr("testString")
				validationInstallOptionsModel.VcenterID = core.StringPtr("testString")
				validationInstallOptionsModel.VcenterUser = core.StringPtr("testString")
				validationInstallOptionsModel.VcenterPassword = core.StringPtr("testString")
				validationInstallOptionsModel.VcenterLocation = core.StringPtr("testString")
				validationInstallOptionsModel.VcenterDatastore = core.StringPtr("testString")
				validationInstallOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := catalogManagementService.ValidationInstall(validationInstallOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the ValidationInstallOptions model with no property values
				validationInstallOptionsModelNew := new(catalogmanagementv1.ValidationInstallOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = catalogManagementService.ValidationInstall(validationInstallOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetValidationStatus(getValidationStatusOptions *GetValidationStatusOptions) - Operation response error`, func() {
		getValidationStatusPath := "/versions/testString/validation/install"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getValidationStatusPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Auth-Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetValidationStatus with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetValidationStatusOptions model
				getValidationStatusOptionsModel := new(catalogmanagementv1.GetValidationStatusOptions)
				getValidationStatusOptionsModel.VersionLocID = core.StringPtr("testString")
				getValidationStatusOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				getValidationStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.GetValidationStatus(getValidationStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.GetValidationStatus(getValidationStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetValidationStatus(getValidationStatusOptions *GetValidationStatusOptions)`, func() {
		getValidationStatusPath := "/versions/testString/validation/install"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getValidationStatusPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Auth-Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"validated": "2019-01-01T12:00:00", "requested": "2019-01-01T12:00:00", "state": "State", "last_operation": "LastOperation", "target": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke GetValidationStatus successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.GetValidationStatus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetValidationStatusOptions model
				getValidationStatusOptionsModel := new(catalogmanagementv1.GetValidationStatusOptions)
				getValidationStatusOptionsModel.VersionLocID = core.StringPtr("testString")
				getValidationStatusOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				getValidationStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.GetValidationStatus(getValidationStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetValidationStatusWithContext(ctx, getValidationStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.GetValidationStatus(getValidationStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetValidationStatusWithContext(ctx, getValidationStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetValidationStatus with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetValidationStatusOptions model
				getValidationStatusOptionsModel := new(catalogmanagementv1.GetValidationStatusOptions)
				getValidationStatusOptionsModel.VersionLocID = core.StringPtr("testString")
				getValidationStatusOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				getValidationStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.GetValidationStatus(getValidationStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetValidationStatusOptions model with no property values
				getValidationStatusOptionsModelNew := new(catalogmanagementv1.GetValidationStatusOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = catalogManagementService.GetValidationStatus(getValidationStatusOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetOverrideValues(getOverrideValuesOptions *GetOverrideValuesOptions)`, func() {
		getOverrideValuesPath := "/versions/testString/validation/overridevalues"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getOverrideValuesPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"mapKey": "anyValue"}`)
				}))
			})
			It(`Invoke GetOverrideValues successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.GetOverrideValues(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetOverrideValuesOptions model
				getOverrideValuesOptionsModel := new(catalogmanagementv1.GetOverrideValuesOptions)
				getOverrideValuesOptionsModel.VersionLocID = core.StringPtr("testString")
				getOverrideValuesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.GetOverrideValues(getOverrideValuesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetOverrideValuesWithContext(ctx, getOverrideValuesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.GetOverrideValues(getOverrideValuesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetOverrideValuesWithContext(ctx, getOverrideValuesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetOverrideValues with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetOverrideValuesOptions model
				getOverrideValuesOptionsModel := new(catalogmanagementv1.GetOverrideValuesOptions)
				getOverrideValuesOptionsModel.VersionLocID = core.StringPtr("testString")
				getOverrideValuesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.GetOverrideValues(getOverrideValuesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetOverrideValuesOptions model with no property values
				getOverrideValuesOptionsModelNew := new(catalogmanagementv1.GetOverrideValuesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = catalogManagementService.GetOverrideValues(getOverrideValuesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSchematicsWorkspaces(getSchematicsWorkspacesOptions *GetSchematicsWorkspacesOptions) - Operation response error`, func() {
		getSchematicsWorkspacesPath := "/versions/testString/workspaces"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSchematicsWorkspacesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Auth-Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSchematicsWorkspaces with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetSchematicsWorkspacesOptions model
				getSchematicsWorkspacesOptionsModel := new(catalogmanagementv1.GetSchematicsWorkspacesOptions)
				getSchematicsWorkspacesOptionsModel.VersionLocID = core.StringPtr("testString")
				getSchematicsWorkspacesOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				getSchematicsWorkspacesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.GetSchematicsWorkspaces(getSchematicsWorkspacesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.GetSchematicsWorkspaces(getSchematicsWorkspacesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetSchematicsWorkspaces(getSchematicsWorkspacesOptions *GetSchematicsWorkspacesOptions)`, func() {
		getSchematicsWorkspacesPath := "/versions/testString/workspaces"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSchematicsWorkspacesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Auth-Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 6, "limit": 5, "total_count": 10, "resource_count": 13, "first": "First", "last": "Last", "prev": "Prev", "next": "Next", "resources": [{"id": "ID", "name": "Name", "type": ["Type"], "description": "Description", "tags": ["Tags"], "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "status": "Status", "workspace_status": {"frozen": true, "locked": true}, "template_ref": "TemplateRef", "template_repo": {"repo_url": "RepoURL", "chart_name": "ChartName", "script_name": "ScriptName", "uninstall_script_name": "UninstallScriptName", "folder_name": "FolderName", "repo_sha_value": "RepoShaValue"}, "template_data": [{"anyKey": "anyValue"}], "runtime_data": {"id": "ID", "engine_name": "EngineName", "engine_version": "EngineVersion", "state_store_url": "StateStoreURL", "log_store_url": "LogStoreURL"}, "shared_data": {"anyKey": "anyValue"}, "catalog_ref": {"item_id": "ItemID", "item_name": "ItemName", "item_url": "ItemURL"}}]}`)
				}))
			})
			It(`Invoke GetSchematicsWorkspaces successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.GetSchematicsWorkspaces(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSchematicsWorkspacesOptions model
				getSchematicsWorkspacesOptionsModel := new(catalogmanagementv1.GetSchematicsWorkspacesOptions)
				getSchematicsWorkspacesOptionsModel.VersionLocID = core.StringPtr("testString")
				getSchematicsWorkspacesOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				getSchematicsWorkspacesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.GetSchematicsWorkspaces(getSchematicsWorkspacesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetSchematicsWorkspacesWithContext(ctx, getSchematicsWorkspacesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.GetSchematicsWorkspaces(getSchematicsWorkspacesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetSchematicsWorkspacesWithContext(ctx, getSchematicsWorkspacesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetSchematicsWorkspaces with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetSchematicsWorkspacesOptions model
				getSchematicsWorkspacesOptionsModel := new(catalogmanagementv1.GetSchematicsWorkspacesOptions)
				getSchematicsWorkspacesOptionsModel.VersionLocID = core.StringPtr("testString")
				getSchematicsWorkspacesOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				getSchematicsWorkspacesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.GetSchematicsWorkspaces(getSchematicsWorkspacesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetSchematicsWorkspacesOptions model with no property values
				getSchematicsWorkspacesOptionsModelNew := new(catalogmanagementv1.GetSchematicsWorkspacesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = catalogManagementService.GetSchematicsWorkspaces(getSchematicsWorkspacesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CanDeploySchematics(canDeploySchematicsOptions *CanDeploySchematicsOptions) - Operation response error`, func() {
		canDeploySchematicsPath := "/versions/testString/candeploy"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(canDeploySchematicsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["cluster_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["region"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["namespace"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CanDeploySchematics with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the CanDeploySchematicsOptions model
				canDeploySchematicsOptionsModel := new(catalogmanagementv1.CanDeploySchematicsOptions)
				canDeploySchematicsOptionsModel.VersionLocID = core.StringPtr("testString")
				canDeploySchematicsOptionsModel.ClusterID = core.StringPtr("testString")
				canDeploySchematicsOptionsModel.Region = core.StringPtr("testString")
				canDeploySchematicsOptionsModel.Namespace = core.StringPtr("testString")
				canDeploySchematicsOptionsModel.ResourceGroupID = core.StringPtr("testString")
				canDeploySchematicsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.CanDeploySchematics(canDeploySchematicsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.CanDeploySchematics(canDeploySchematicsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CanDeploySchematics(canDeploySchematicsOptions *CanDeploySchematicsOptions)`, func() {
		canDeploySchematicsPath := "/versions/testString/candeploy"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(canDeploySchematicsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["cluster_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["region"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["namespace"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"pre_install": {"anyKey": "anyValue"}, "install": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke CanDeploySchematics successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.CanDeploySchematics(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CanDeploySchematicsOptions model
				canDeploySchematicsOptionsModel := new(catalogmanagementv1.CanDeploySchematicsOptions)
				canDeploySchematicsOptionsModel.VersionLocID = core.StringPtr("testString")
				canDeploySchematicsOptionsModel.ClusterID = core.StringPtr("testString")
				canDeploySchematicsOptionsModel.Region = core.StringPtr("testString")
				canDeploySchematicsOptionsModel.Namespace = core.StringPtr("testString")
				canDeploySchematicsOptionsModel.ResourceGroupID = core.StringPtr("testString")
				canDeploySchematicsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.CanDeploySchematics(canDeploySchematicsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.CanDeploySchematicsWithContext(ctx, canDeploySchematicsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.CanDeploySchematics(canDeploySchematicsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.CanDeploySchematicsWithContext(ctx, canDeploySchematicsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke CanDeploySchematics with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the CanDeploySchematicsOptions model
				canDeploySchematicsOptionsModel := new(catalogmanagementv1.CanDeploySchematicsOptions)
				canDeploySchematicsOptionsModel.VersionLocID = core.StringPtr("testString")
				canDeploySchematicsOptionsModel.ClusterID = core.StringPtr("testString")
				canDeploySchematicsOptionsModel.Region = core.StringPtr("testString")
				canDeploySchematicsOptionsModel.Namespace = core.StringPtr("testString")
				canDeploySchematicsOptionsModel.ResourceGroupID = core.StringPtr("testString")
				canDeploySchematicsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.CanDeploySchematics(canDeploySchematicsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CanDeploySchematicsOptions model with no property values
				canDeploySchematicsOptionsModelNew := new(catalogmanagementv1.CanDeploySchematicsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = catalogManagementService.CanDeploySchematics(canDeploySchematicsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetResourceGroups(getResourceGroupsOptions *GetResourceGroupsOptions) - Operation response error`, func() {
		getResourceGroupsPath := "/deploy/schematics/resourcegroups"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getResourceGroupsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetResourceGroups with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetResourceGroupsOptions model
				getResourceGroupsOptionsModel := new(catalogmanagementv1.GetResourceGroupsOptions)
				getResourceGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.GetResourceGroups(getResourceGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.GetResourceGroups(getResourceGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetResourceGroups(getResourceGroupsOptions *GetResourceGroupsOptions)`, func() {
		getResourceGroupsPath := "/deploy/schematics/resourcegroups"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getResourceGroupsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 6, "limit": 5, "total_count": 10, "resource_count": 13, "first": "First", "last": "Last", "prev": "Prev", "next": "Next", "resources": [{"id": "ID", "name": "Name", "crn": "Crn", "account_id": "AccountID", "state": "State", "default": false}]}`)
				}))
			})
			It(`Invoke GetResourceGroups successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.GetResourceGroups(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetResourceGroupsOptions model
				getResourceGroupsOptionsModel := new(catalogmanagementv1.GetResourceGroupsOptions)
				getResourceGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.GetResourceGroups(getResourceGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetResourceGroupsWithContext(ctx, getResourceGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.GetResourceGroups(getResourceGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetResourceGroupsWithContext(ctx, getResourceGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetResourceGroups with error: Operation request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetResourceGroupsOptions model
				getResourceGroupsOptionsModel := new(catalogmanagementv1.GetResourceGroupsOptions)
				getResourceGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.GetResourceGroups(getResourceGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(catalogManagementService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(catalogManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
				URL: "https://catalogmanagementv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(catalogManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CATALOG_MANAGEMENT_URL": "https://catalogmanagementv1/api",
				"CATALOG_MANAGEMENT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
				})
				Expect(catalogManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
					URL: "https://testService/api",
				})
				Expect(catalogManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
				})
				err := catalogManagementService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CATALOG_MANAGEMENT_URL": "https://catalogmanagementv1/api",
				"CATALOG_MANAGEMENT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(catalogManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CATALOG_MANAGEMENT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(catalogManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`GetLicenseProviders(getLicenseProvidersOptions *GetLicenseProvidersOptions) - Operation response error`, func() {
		getLicenseProvidersPath := "/license/license_providers"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLicenseProvidersPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetLicenseProviders with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetLicenseProvidersOptions model
				getLicenseProvidersOptionsModel := new(catalogmanagementv1.GetLicenseProvidersOptions)
				getLicenseProvidersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.GetLicenseProviders(getLicenseProvidersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.GetLicenseProviders(getLicenseProvidersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetLicenseProviders(getLicenseProvidersOptions *GetLicenseProvidersOptions)`, func() {
		getLicenseProvidersPath := "/license/license_providers"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLicenseProvidersPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_results": 12, "total_pages": 10, "prev_url": "PrevURL", "next_url": "NextURL", "resources": [{"name": "Name", "short_description": "ShortDescription", "id": "ID", "licence_type": "LicenceType", "offering_type": "OfferingType", "create_url": "CreateURL", "info_url": "InfoURL", "url": "URL", "crn": "Crn", "state": "State"}]}`)
				}))
			})
			It(`Invoke GetLicenseProviders successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.GetLicenseProviders(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLicenseProvidersOptions model
				getLicenseProvidersOptionsModel := new(catalogmanagementv1.GetLicenseProvidersOptions)
				getLicenseProvidersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.GetLicenseProviders(getLicenseProvidersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetLicenseProvidersWithContext(ctx, getLicenseProvidersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.GetLicenseProviders(getLicenseProvidersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetLicenseProvidersWithContext(ctx, getLicenseProvidersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetLicenseProviders with error: Operation request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetLicenseProvidersOptions model
				getLicenseProvidersOptionsModel := new(catalogmanagementv1.GetLicenseProvidersOptions)
				getLicenseProvidersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.GetLicenseProviders(getLicenseProvidersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListLicenseEntitlements(listLicenseEntitlementsOptions *ListLicenseEntitlementsOptions) - Operation response error`, func() {
		listLicenseEntitlementsPath := "/license/entitlements"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listLicenseEntitlementsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["license_product_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["version_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["state"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListLicenseEntitlements with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the ListLicenseEntitlementsOptions model
				listLicenseEntitlementsOptionsModel := new(catalogmanagementv1.ListLicenseEntitlementsOptions)
				listLicenseEntitlementsOptionsModel.AccountID = core.StringPtr("testString")
				listLicenseEntitlementsOptionsModel.LicenseProductID = core.StringPtr("testString")
				listLicenseEntitlementsOptionsModel.VersionID = core.StringPtr("testString")
				listLicenseEntitlementsOptionsModel.State = core.StringPtr("testString")
				listLicenseEntitlementsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.ListLicenseEntitlements(listLicenseEntitlementsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.ListLicenseEntitlements(listLicenseEntitlementsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListLicenseEntitlements(listLicenseEntitlementsOptions *ListLicenseEntitlementsOptions)`, func() {
		listLicenseEntitlementsPath := "/license/entitlements"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listLicenseEntitlementsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["license_product_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["version_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["state"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_results": 12, "total_pages": 10, "prev_url": "PrevURL", "next_url": "NextURL", "resources": [{"name": "Name", "id": "ID", "crn": "Crn", "url": "URL", "offering_type": "OfferingType", "state": "State", "effective_from": "EffectiveFrom", "effective_until": "EffectiveUntil", "account_id": "AccountID", "owner_id": "OwnerID", "version_id": "VersionID", "license_offering_id": "LicenseOfferingID", "license_id": "LicenseID", "license_owner_id": "LicenseOwnerID", "license_type": "LicenseType", "license_provider_id": "LicenseProviderID", "license_provider_url": "LicenseProviderURL", "license_product_id": "LicenseProductID", "namespace_repository": "NamespaceRepository", "apikey": "Apikey", "create_by": "CreateBy", "update_by": "UpdateBy", "create_at": "CreateAt", "updated_at": "UpdatedAt", "history": [{"action": "Action", "user": "User", "date": "Date"}], "offering_list": [{"id": "ID", "name": "Name", "label": "Label", "offering_icon_url": "OfferingIconURL", "account_id": "AccountID", "catalog_id": "CatalogID"}]}]}`)
				}))
			})
			It(`Invoke ListLicenseEntitlements successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.ListLicenseEntitlements(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListLicenseEntitlementsOptions model
				listLicenseEntitlementsOptionsModel := new(catalogmanagementv1.ListLicenseEntitlementsOptions)
				listLicenseEntitlementsOptionsModel.AccountID = core.StringPtr("testString")
				listLicenseEntitlementsOptionsModel.LicenseProductID = core.StringPtr("testString")
				listLicenseEntitlementsOptionsModel.VersionID = core.StringPtr("testString")
				listLicenseEntitlementsOptionsModel.State = core.StringPtr("testString")
				listLicenseEntitlementsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.ListLicenseEntitlements(listLicenseEntitlementsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.ListLicenseEntitlementsWithContext(ctx, listLicenseEntitlementsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.ListLicenseEntitlements(listLicenseEntitlementsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.ListLicenseEntitlementsWithContext(ctx, listLicenseEntitlementsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListLicenseEntitlements with error: Operation request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the ListLicenseEntitlementsOptions model
				listLicenseEntitlementsOptionsModel := new(catalogmanagementv1.ListLicenseEntitlementsOptions)
				listLicenseEntitlementsOptionsModel.AccountID = core.StringPtr("testString")
				listLicenseEntitlementsOptionsModel.LicenseProductID = core.StringPtr("testString")
				listLicenseEntitlementsOptionsModel.VersionID = core.StringPtr("testString")
				listLicenseEntitlementsOptionsModel.State = core.StringPtr("testString")
				listLicenseEntitlementsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.ListLicenseEntitlements(listLicenseEntitlementsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateLicenseEntitlement(createLicenseEntitlementOptions *CreateLicenseEntitlementOptions) - Operation response error`, func() {
		createLicenseEntitlementPath := "/license/entitlements"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createLicenseEntitlementPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateLicenseEntitlement with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the CreateLicenseEntitlementOptions model
				createLicenseEntitlementOptionsModel := new(catalogmanagementv1.CreateLicenseEntitlementOptions)
				createLicenseEntitlementOptionsModel.Name = core.StringPtr("testString")
				createLicenseEntitlementOptionsModel.EffectiveFrom = core.StringPtr("testString")
				createLicenseEntitlementOptionsModel.EffectiveUntil = core.StringPtr("testString")
				createLicenseEntitlementOptionsModel.VersionID = core.StringPtr("testString")
				createLicenseEntitlementOptionsModel.LicenseID = core.StringPtr("testString")
				createLicenseEntitlementOptionsModel.LicenseOwnerID = core.StringPtr("testString")
				createLicenseEntitlementOptionsModel.LicenseProviderID = core.StringPtr("testString")
				createLicenseEntitlementOptionsModel.LicenseProductID = core.StringPtr("testString")
				createLicenseEntitlementOptionsModel.AccountID = core.StringPtr("testString")
				createLicenseEntitlementOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.CreateLicenseEntitlement(createLicenseEntitlementOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.CreateLicenseEntitlement(createLicenseEntitlementOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateLicenseEntitlement(createLicenseEntitlementOptions *CreateLicenseEntitlementOptions)`, func() {
		createLicenseEntitlementPath := "/license/entitlements"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createLicenseEntitlementPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "id": "ID", "crn": "Crn", "url": "URL", "offering_type": "OfferingType", "state": "State", "effective_from": "EffectiveFrom", "effective_until": "EffectiveUntil", "account_id": "AccountID", "owner_id": "OwnerID", "version_id": "VersionID", "license_offering_id": "LicenseOfferingID", "license_id": "LicenseID", "license_owner_id": "LicenseOwnerID", "license_type": "LicenseType", "license_provider_id": "LicenseProviderID", "license_provider_url": "LicenseProviderURL", "license_product_id": "LicenseProductID", "namespace_repository": "NamespaceRepository", "apikey": "Apikey", "create_by": "CreateBy", "update_by": "UpdateBy", "create_at": "CreateAt", "updated_at": "UpdatedAt", "history": [{"action": "Action", "user": "User", "date": "Date"}], "offering_list": [{"id": "ID", "name": "Name", "label": "Label", "offering_icon_url": "OfferingIconURL", "account_id": "AccountID", "catalog_id": "CatalogID"}]}`)
				}))
			})
			It(`Invoke CreateLicenseEntitlement successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.CreateLicenseEntitlement(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateLicenseEntitlementOptions model
				createLicenseEntitlementOptionsModel := new(catalogmanagementv1.CreateLicenseEntitlementOptions)
				createLicenseEntitlementOptionsModel.Name = core.StringPtr("testString")
				createLicenseEntitlementOptionsModel.EffectiveFrom = core.StringPtr("testString")
				createLicenseEntitlementOptionsModel.EffectiveUntil = core.StringPtr("testString")
				createLicenseEntitlementOptionsModel.VersionID = core.StringPtr("testString")
				createLicenseEntitlementOptionsModel.LicenseID = core.StringPtr("testString")
				createLicenseEntitlementOptionsModel.LicenseOwnerID = core.StringPtr("testString")
				createLicenseEntitlementOptionsModel.LicenseProviderID = core.StringPtr("testString")
				createLicenseEntitlementOptionsModel.LicenseProductID = core.StringPtr("testString")
				createLicenseEntitlementOptionsModel.AccountID = core.StringPtr("testString")
				createLicenseEntitlementOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.CreateLicenseEntitlement(createLicenseEntitlementOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.CreateLicenseEntitlementWithContext(ctx, createLicenseEntitlementOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.CreateLicenseEntitlement(createLicenseEntitlementOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.CreateLicenseEntitlementWithContext(ctx, createLicenseEntitlementOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke CreateLicenseEntitlement with error: Operation request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the CreateLicenseEntitlementOptions model
				createLicenseEntitlementOptionsModel := new(catalogmanagementv1.CreateLicenseEntitlementOptions)
				createLicenseEntitlementOptionsModel.Name = core.StringPtr("testString")
				createLicenseEntitlementOptionsModel.EffectiveFrom = core.StringPtr("testString")
				createLicenseEntitlementOptionsModel.EffectiveUntil = core.StringPtr("testString")
				createLicenseEntitlementOptionsModel.VersionID = core.StringPtr("testString")
				createLicenseEntitlementOptionsModel.LicenseID = core.StringPtr("testString")
				createLicenseEntitlementOptionsModel.LicenseOwnerID = core.StringPtr("testString")
				createLicenseEntitlementOptionsModel.LicenseProviderID = core.StringPtr("testString")
				createLicenseEntitlementOptionsModel.LicenseProductID = core.StringPtr("testString")
				createLicenseEntitlementOptionsModel.AccountID = core.StringPtr("testString")
				createLicenseEntitlementOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.CreateLicenseEntitlement(createLicenseEntitlementOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetLicenseEntitlements(getLicenseEntitlementsOptions *GetLicenseEntitlementsOptions) - Operation response error`, func() {
		getLicenseEntitlementsPath := "/license/entitlements/productID/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLicenseEntitlementsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["version_id"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetLicenseEntitlements with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetLicenseEntitlementsOptions model
				getLicenseEntitlementsOptionsModel := new(catalogmanagementv1.GetLicenseEntitlementsOptions)
				getLicenseEntitlementsOptionsModel.LicenseProductID = core.StringPtr("testString")
				getLicenseEntitlementsOptionsModel.AccountID = core.StringPtr("testString")
				getLicenseEntitlementsOptionsModel.VersionID = core.StringPtr("testString")
				getLicenseEntitlementsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.GetLicenseEntitlements(getLicenseEntitlementsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.GetLicenseEntitlements(getLicenseEntitlementsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetLicenseEntitlements(getLicenseEntitlementsOptions *GetLicenseEntitlementsOptions)`, func() {
		getLicenseEntitlementsPath := "/license/entitlements/productID/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLicenseEntitlementsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["version_id"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_results": 12, "total_pages": 10, "prev_url": "PrevURL", "next_url": "NextURL", "resources": [{"name": "Name", "id": "ID", "crn": "Crn", "url": "URL", "offering_type": "OfferingType", "state": "State", "effective_from": "EffectiveFrom", "effective_until": "EffectiveUntil", "account_id": "AccountID", "owner_id": "OwnerID", "version_id": "VersionID", "license_offering_id": "LicenseOfferingID", "license_id": "LicenseID", "license_owner_id": "LicenseOwnerID", "license_type": "LicenseType", "license_provider_id": "LicenseProviderID", "license_provider_url": "LicenseProviderURL", "license_product_id": "LicenseProductID", "namespace_repository": "NamespaceRepository", "apikey": "Apikey", "create_by": "CreateBy", "update_by": "UpdateBy", "create_at": "CreateAt", "updated_at": "UpdatedAt", "history": [{"action": "Action", "user": "User", "date": "Date"}], "offering_list": [{"id": "ID", "name": "Name", "label": "Label", "offering_icon_url": "OfferingIconURL", "account_id": "AccountID", "catalog_id": "CatalogID"}]}]}`)
				}))
			})
			It(`Invoke GetLicenseEntitlements successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.GetLicenseEntitlements(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLicenseEntitlementsOptions model
				getLicenseEntitlementsOptionsModel := new(catalogmanagementv1.GetLicenseEntitlementsOptions)
				getLicenseEntitlementsOptionsModel.LicenseProductID = core.StringPtr("testString")
				getLicenseEntitlementsOptionsModel.AccountID = core.StringPtr("testString")
				getLicenseEntitlementsOptionsModel.VersionID = core.StringPtr("testString")
				getLicenseEntitlementsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.GetLicenseEntitlements(getLicenseEntitlementsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetLicenseEntitlementsWithContext(ctx, getLicenseEntitlementsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.GetLicenseEntitlements(getLicenseEntitlementsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetLicenseEntitlementsWithContext(ctx, getLicenseEntitlementsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetLicenseEntitlements with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetLicenseEntitlementsOptions model
				getLicenseEntitlementsOptionsModel := new(catalogmanagementv1.GetLicenseEntitlementsOptions)
				getLicenseEntitlementsOptionsModel.LicenseProductID = core.StringPtr("testString")
				getLicenseEntitlementsOptionsModel.AccountID = core.StringPtr("testString")
				getLicenseEntitlementsOptionsModel.VersionID = core.StringPtr("testString")
				getLicenseEntitlementsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.GetLicenseEntitlements(getLicenseEntitlementsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetLicenseEntitlementsOptions model with no property values
				getLicenseEntitlementsOptionsModelNew := new(catalogmanagementv1.GetLicenseEntitlementsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = catalogManagementService.GetLicenseEntitlements(getLicenseEntitlementsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteLicenseEntitlement(deleteLicenseEntitlementOptions *DeleteLicenseEntitlementOptions)`, func() {
		deleteLicenseEntitlementPath := "/license/entitlements/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteLicenseEntitlementPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteLicenseEntitlement successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := catalogManagementService.DeleteLicenseEntitlement(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteLicenseEntitlementOptions model
				deleteLicenseEntitlementOptionsModel := new(catalogmanagementv1.DeleteLicenseEntitlementOptions)
				deleteLicenseEntitlementOptionsModel.EntitlementID = core.StringPtr("testString")
				deleteLicenseEntitlementOptionsModel.AccountID = core.StringPtr("testString")
				deleteLicenseEntitlementOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = catalogManagementService.DeleteLicenseEntitlement(deleteLicenseEntitlementOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				response, operationErr = catalogManagementService.DeleteLicenseEntitlement(deleteLicenseEntitlementOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteLicenseEntitlement with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the DeleteLicenseEntitlementOptions model
				deleteLicenseEntitlementOptionsModel := new(catalogmanagementv1.DeleteLicenseEntitlementOptions)
				deleteLicenseEntitlementOptionsModel.EntitlementID = core.StringPtr("testString")
				deleteLicenseEntitlementOptionsModel.AccountID = core.StringPtr("testString")
				deleteLicenseEntitlementOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := catalogManagementService.DeleteLicenseEntitlement(deleteLicenseEntitlementOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteLicenseEntitlementOptions model with no property values
				deleteLicenseEntitlementOptionsModelNew := new(catalogmanagementv1.DeleteLicenseEntitlementOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = catalogManagementService.DeleteLicenseEntitlement(deleteLicenseEntitlementOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetLicenses(getLicensesOptions *GetLicensesOptions) - Operation response error`, func() {
		getLicensesPath := "/license/licenses"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLicensesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["license_provider_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["license_type"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["license_product_id"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetLicenses with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetLicensesOptions model
				getLicensesOptionsModel := new(catalogmanagementv1.GetLicensesOptions)
				getLicensesOptionsModel.LicenseProviderID = core.StringPtr("testString")
				getLicensesOptionsModel.AccountID = core.StringPtr("testString")
				getLicensesOptionsModel.Name = core.StringPtr("testString")
				getLicensesOptionsModel.LicenseType = core.StringPtr("testString")
				getLicensesOptionsModel.LicenseProductID = core.StringPtr("testString")
				getLicensesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.GetLicenses(getLicensesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.GetLicenses(getLicensesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetLicenses(getLicensesOptions *GetLicensesOptions)`, func() {
		getLicensesPath := "/license/licenses"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLicensesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["license_provider_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["license_type"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["license_product_id"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_results": 12, "total_pages": 10, "prev_url": "PrevURL", "next_url": "NextURL", "resources": [{"name": "Name", "offering_type": "OfferingType", "seats_allowed": "SeatsAllowed", "seats_used": "SeatsUsed", "owner_id": "OwnerID", "license_offering_id": "LicenseOfferingID", "license_id": "LicenseID", "license_owner_id": "LicenseOwnerID", "license_type": "LicenseType", "license_provider_id": "LicenseProviderID", "license_product_id": "LicenseProductID", "license_provider_url": "LicenseProviderURL", "effective_from": "EffectiveFrom", "effective_until": "EffectiveUntil", "internal": true, "offering_list": [{"id": "ID", "name": "Name", "label": "Label", "offering_icon_url": "OfferingIconURL", "account_id": "AccountID", "catalog_id": "CatalogID"}]}]}`)
				}))
			})
			It(`Invoke GetLicenses successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.GetLicenses(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLicensesOptions model
				getLicensesOptionsModel := new(catalogmanagementv1.GetLicensesOptions)
				getLicensesOptionsModel.LicenseProviderID = core.StringPtr("testString")
				getLicensesOptionsModel.AccountID = core.StringPtr("testString")
				getLicensesOptionsModel.Name = core.StringPtr("testString")
				getLicensesOptionsModel.LicenseType = core.StringPtr("testString")
				getLicensesOptionsModel.LicenseProductID = core.StringPtr("testString")
				getLicensesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.GetLicenses(getLicensesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetLicensesWithContext(ctx, getLicensesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.GetLicenses(getLicensesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetLicensesWithContext(ctx, getLicensesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetLicenses with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetLicensesOptions model
				getLicensesOptionsModel := new(catalogmanagementv1.GetLicensesOptions)
				getLicensesOptionsModel.LicenseProviderID = core.StringPtr("testString")
				getLicensesOptionsModel.AccountID = core.StringPtr("testString")
				getLicensesOptionsModel.Name = core.StringPtr("testString")
				getLicensesOptionsModel.LicenseType = core.StringPtr("testString")
				getLicensesOptionsModel.LicenseProductID = core.StringPtr("testString")
				getLicensesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.GetLicenses(getLicensesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetLicensesOptions model with no property values
				getLicensesOptionsModelNew := new(catalogmanagementv1.GetLicensesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = catalogManagementService.GetLicenses(getLicensesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(catalogManagementService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(catalogManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
				URL: "https://catalogmanagementv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(catalogManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CATALOG_MANAGEMENT_URL": "https://catalogmanagementv1/api",
				"CATALOG_MANAGEMENT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
				})
				Expect(catalogManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
					URL: "https://testService/api",
				})
				Expect(catalogManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
				})
				err := catalogManagementService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CATALOG_MANAGEMENT_URL": "https://catalogmanagementv1/api",
				"CATALOG_MANAGEMENT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(catalogManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CATALOG_MANAGEMENT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(catalogManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})

	Describe(`SearchLicenseVersions(searchLicenseVersionsOptions *SearchLicenseVersionsOptions)`, func() {
		searchLicenseVersionsPath := "/search/license/versions"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(searchLicenseVersionsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["q"]).To(Equal([]string{"testString"}))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke SearchLicenseVersions successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := catalogManagementService.SearchLicenseVersions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the SearchLicenseVersionsOptions model
				searchLicenseVersionsOptionsModel := new(catalogmanagementv1.SearchLicenseVersionsOptions)
				searchLicenseVersionsOptionsModel.Q = core.StringPtr("testString")
				searchLicenseVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = catalogManagementService.SearchLicenseVersions(searchLicenseVersionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				response, operationErr = catalogManagementService.SearchLicenseVersions(searchLicenseVersionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke SearchLicenseVersions with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the SearchLicenseVersionsOptions model
				searchLicenseVersionsOptionsModel := new(catalogmanagementv1.SearchLicenseVersionsOptions)
				searchLicenseVersionsOptionsModel.Q = core.StringPtr("testString")
				searchLicenseVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := catalogManagementService.SearchLicenseVersions(searchLicenseVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the SearchLicenseVersionsOptions model with no property values
				searchLicenseVersionsOptionsModelNew := new(catalogmanagementv1.SearchLicenseVersionsOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = catalogManagementService.SearchLicenseVersions(searchLicenseVersionsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`SearchLicenseOfferings(searchLicenseOfferingsOptions *SearchLicenseOfferingsOptions)`, func() {
		searchLicenseOfferingsPath := "/search/license/offerings"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(searchLicenseOfferingsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["q"]).To(Equal([]string{"testString"}))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke SearchLicenseOfferings successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := catalogManagementService.SearchLicenseOfferings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the SearchLicenseOfferingsOptions model
				searchLicenseOfferingsOptionsModel := new(catalogmanagementv1.SearchLicenseOfferingsOptions)
				searchLicenseOfferingsOptionsModel.Q = core.StringPtr("testString")
				searchLicenseOfferingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = catalogManagementService.SearchLicenseOfferings(searchLicenseOfferingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				response, operationErr = catalogManagementService.SearchLicenseOfferings(searchLicenseOfferingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke SearchLicenseOfferings with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the SearchLicenseOfferingsOptions model
				searchLicenseOfferingsOptionsModel := new(catalogmanagementv1.SearchLicenseOfferingsOptions)
				searchLicenseOfferingsOptionsModel.Q = core.StringPtr("testString")
				searchLicenseOfferingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := catalogManagementService.SearchLicenseOfferings(searchLicenseOfferingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the SearchLicenseOfferingsOptions model with no property values
				searchLicenseOfferingsOptionsModelNew := new(catalogmanagementv1.SearchLicenseOfferingsOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = catalogManagementService.SearchLicenseOfferings(searchLicenseOfferingsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(catalogManagementService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(catalogManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
				URL: "https://catalogmanagementv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(catalogManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CATALOG_MANAGEMENT_URL": "https://catalogmanagementv1/api",
				"CATALOG_MANAGEMENT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
				})
				Expect(catalogManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
					URL: "https://testService/api",
				})
				Expect(catalogManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
				})
				err := catalogManagementService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CATALOG_MANAGEMENT_URL": "https://catalogmanagementv1/api",
				"CATALOG_MANAGEMENT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(catalogManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CATALOG_MANAGEMENT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1UsingExternalConfig(&catalogmanagementv1.CatalogManagementV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(catalogManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`SearchObjects(searchObjectsOptions *SearchObjectsOptions) - Operation response error`, func() {
		searchObjectsPath := "/objects"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(searchObjectsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["query"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))


					// TODO: Add check for collapse query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SearchObjects with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the SearchObjectsOptions model
				searchObjectsOptionsModel := new(catalogmanagementv1.SearchObjectsOptions)
				searchObjectsOptionsModel.Query = core.StringPtr("testString")
				searchObjectsOptionsModel.Limit = core.Int64Ptr(int64(38))
				searchObjectsOptionsModel.Offset = core.Int64Ptr(int64(38))
				searchObjectsOptionsModel.Collapse = core.BoolPtr(true)
				searchObjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.SearchObjects(searchObjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.SearchObjects(searchObjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`SearchObjects(searchObjectsOptions *SearchObjectsOptions)`, func() {
		searchObjectsPath := "/objects"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(searchObjectsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["query"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))


					// TODO: Add check for collapse query parameter

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 6, "limit": 5, "total_count": 10, "resource_count": 13, "first": "First", "last": "Last", "prev": "Prev", "next": "Next", "resources": [{"id": "ID", "order": [5], "fields": {"catalog_id": "CatalogID", "name": "Name", "parent_id": "ParentID", "label": "Label", "updated": "2019-01-01T12:00:00", "kind": "Kind", "parent_name": "ParentName"}}]}`)
				}))
			})
			It(`Invoke SearchObjects successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.SearchObjects(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SearchObjectsOptions model
				searchObjectsOptionsModel := new(catalogmanagementv1.SearchObjectsOptions)
				searchObjectsOptionsModel.Query = core.StringPtr("testString")
				searchObjectsOptionsModel.Limit = core.Int64Ptr(int64(38))
				searchObjectsOptionsModel.Offset = core.Int64Ptr(int64(38))
				searchObjectsOptionsModel.Collapse = core.BoolPtr(true)
				searchObjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.SearchObjects(searchObjectsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.SearchObjectsWithContext(ctx, searchObjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.SearchObjects(searchObjectsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.SearchObjectsWithContext(ctx, searchObjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke SearchObjects with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the SearchObjectsOptions model
				searchObjectsOptionsModel := new(catalogmanagementv1.SearchObjectsOptions)
				searchObjectsOptionsModel.Query = core.StringPtr("testString")
				searchObjectsOptionsModel.Limit = core.Int64Ptr(int64(38))
				searchObjectsOptionsModel.Offset = core.Int64Ptr(int64(38))
				searchObjectsOptionsModel.Collapse = core.BoolPtr(true)
				searchObjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.SearchObjects(searchObjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SearchObjectsOptions model with no property values
				searchObjectsOptionsModelNew := new(catalogmanagementv1.SearchObjectsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = catalogManagementService.SearchObjects(searchObjectsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListObjects(listObjectsOptions *ListObjectsOptions) - Operation response error`, func() {
		listObjectsPath := "/catalogs/testString/objects"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listObjectsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListObjects with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the ListObjectsOptions model
				listObjectsOptionsModel := new(catalogmanagementv1.ListObjectsOptions)
				listObjectsOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				listObjectsOptionsModel.Limit = core.Int64Ptr(int64(38))
				listObjectsOptionsModel.Offset = core.Int64Ptr(int64(38))
				listObjectsOptionsModel.Name = core.StringPtr("testString")
				listObjectsOptionsModel.Sort = core.StringPtr("testString")
				listObjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.ListObjects(listObjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.ListObjects(listObjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListObjects(listObjectsOptions *ListObjectsOptions)`, func() {
		listObjectsPath := "/catalogs/testString/objects"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listObjectsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"offset": 6, "limit": 5, "total_count": 10, "resource_count": 13, "first": "First", "last": "Last", "prev": "Prev", "next": "Next", "resources": [{"id": "ID", "name": "Name", "_rev": "Rev", "crn": "Crn", "url": "URL", "parent_id": "ParentID", "allow_list": ["AllowList"], "label_i18n": "LabelI18n", "label": "Label", "tags": ["Tags"], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "short_description": "ShortDescription", "short_description_i18n": "ShortDescriptionI18n", "kind": "Kind", "publish": {"permit_ibm_public_publish": true, "ibm_approved": false, "public_approved": true, "portal_approval_record": "PortalApprovalRecord", "portal_url": "PortalURL"}, "state": {"current": "Current", "current_entered": "2019-01-01T12:00:00", "pending": "Pending", "pending_requested": "2019-01-01T12:00:00", "previous": "Previous"}, "catalog_id": "CatalogID", "catalog_name": "CatalogName", "data": {"anyKey": "anyValue"}}]}`)
				}))
			})
			It(`Invoke ListObjects successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.ListObjects(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListObjectsOptions model
				listObjectsOptionsModel := new(catalogmanagementv1.ListObjectsOptions)
				listObjectsOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				listObjectsOptionsModel.Limit = core.Int64Ptr(int64(38))
				listObjectsOptionsModel.Offset = core.Int64Ptr(int64(38))
				listObjectsOptionsModel.Name = core.StringPtr("testString")
				listObjectsOptionsModel.Sort = core.StringPtr("testString")
				listObjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.ListObjects(listObjectsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.ListObjectsWithContext(ctx, listObjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.ListObjects(listObjectsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.ListObjectsWithContext(ctx, listObjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListObjects with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the ListObjectsOptions model
				listObjectsOptionsModel := new(catalogmanagementv1.ListObjectsOptions)
				listObjectsOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				listObjectsOptionsModel.Limit = core.Int64Ptr(int64(38))
				listObjectsOptionsModel.Offset = core.Int64Ptr(int64(38))
				listObjectsOptionsModel.Name = core.StringPtr("testString")
				listObjectsOptionsModel.Sort = core.StringPtr("testString")
				listObjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.ListObjects(listObjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListObjectsOptions model with no property values
				listObjectsOptionsModelNew := new(catalogmanagementv1.ListObjectsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = catalogManagementService.ListObjects(listObjectsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateObject(createObjectOptions *CreateObjectOptions) - Operation response error`, func() {
		createObjectPath := "/catalogs/testString/objects"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createObjectPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateObject with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the PublishObject model
				publishObjectModel := new(catalogmanagementv1.PublishObject)
				publishObjectModel.PermitIbmPublicPublish = core.BoolPtr(true)
				publishObjectModel.IbmApproved = core.BoolPtr(true)
				publishObjectModel.PublicApproved = core.BoolPtr(true)
				publishObjectModel.PortalApprovalRecord = core.StringPtr("testString")
				publishObjectModel.PortalURL = core.StringPtr("testString")

				// Construct an instance of the State model
				stateModel := new(catalogmanagementv1.State)
				stateModel.Current = core.StringPtr("testString")
				stateModel.CurrentEntered = CreateMockDateTime()
				stateModel.Pending = core.StringPtr("testString")
				stateModel.PendingRequested = CreateMockDateTime()
				stateModel.Previous = core.StringPtr("testString")

				// Construct an instance of the CreateObjectOptions model
				createObjectOptionsModel := new(catalogmanagementv1.CreateObjectOptions)
				createObjectOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				createObjectOptionsModel.ID = core.StringPtr("testString")
				createObjectOptionsModel.Name = core.StringPtr("testString")
				createObjectOptionsModel.Rev = core.StringPtr("testString")
				createObjectOptionsModel.Crn = core.StringPtr("testString")
				createObjectOptionsModel.URL = core.StringPtr("testString")
				createObjectOptionsModel.ParentID = core.StringPtr("testString")
				createObjectOptionsModel.AllowList = []string{"testString"}
				createObjectOptionsModel.LabelI18n = core.StringPtr("testString")
				createObjectOptionsModel.Label = core.StringPtr("testString")
				createObjectOptionsModel.Tags = []string{"testString"}
				createObjectOptionsModel.Created = CreateMockDateTime()
				createObjectOptionsModel.Updated = CreateMockDateTime()
				createObjectOptionsModel.ShortDescription = core.StringPtr("testString")
				createObjectOptionsModel.ShortDescriptionI18n = core.StringPtr("testString")
				createObjectOptionsModel.Kind = core.StringPtr("testString")
				createObjectOptionsModel.Publish = publishObjectModel
				createObjectOptionsModel.State = stateModel
				createObjectOptionsModel.CatalogID = core.StringPtr("testString")
				createObjectOptionsModel.CatalogName = core.StringPtr("testString")
				createObjectOptionsModel.Data = map[string]interface{}{"anyKey": "anyValue"}
				createObjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.CreateObject(createObjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.CreateObject(createObjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateObject(createObjectOptions *CreateObjectOptions)`, func() {
		createObjectPath := "/catalogs/testString/objects"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createObjectPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "_rev": "Rev", "crn": "Crn", "url": "URL", "parent_id": "ParentID", "allow_list": ["AllowList"], "label_i18n": "LabelI18n", "label": "Label", "tags": ["Tags"], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "short_description": "ShortDescription", "short_description_i18n": "ShortDescriptionI18n", "kind": "Kind", "publish": {"permit_ibm_public_publish": true, "ibm_approved": false, "public_approved": true, "portal_approval_record": "PortalApprovalRecord", "portal_url": "PortalURL"}, "state": {"current": "Current", "current_entered": "2019-01-01T12:00:00", "pending": "Pending", "pending_requested": "2019-01-01T12:00:00", "previous": "Previous"}, "catalog_id": "CatalogID", "catalog_name": "CatalogName", "data": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke CreateObject successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.CreateObject(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PublishObject model
				publishObjectModel := new(catalogmanagementv1.PublishObject)
				publishObjectModel.PermitIbmPublicPublish = core.BoolPtr(true)
				publishObjectModel.IbmApproved = core.BoolPtr(true)
				publishObjectModel.PublicApproved = core.BoolPtr(true)
				publishObjectModel.PortalApprovalRecord = core.StringPtr("testString")
				publishObjectModel.PortalURL = core.StringPtr("testString")

				// Construct an instance of the State model
				stateModel := new(catalogmanagementv1.State)
				stateModel.Current = core.StringPtr("testString")
				stateModel.CurrentEntered = CreateMockDateTime()
				stateModel.Pending = core.StringPtr("testString")
				stateModel.PendingRequested = CreateMockDateTime()
				stateModel.Previous = core.StringPtr("testString")

				// Construct an instance of the CreateObjectOptions model
				createObjectOptionsModel := new(catalogmanagementv1.CreateObjectOptions)
				createObjectOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				createObjectOptionsModel.ID = core.StringPtr("testString")
				createObjectOptionsModel.Name = core.StringPtr("testString")
				createObjectOptionsModel.Rev = core.StringPtr("testString")
				createObjectOptionsModel.Crn = core.StringPtr("testString")
				createObjectOptionsModel.URL = core.StringPtr("testString")
				createObjectOptionsModel.ParentID = core.StringPtr("testString")
				createObjectOptionsModel.AllowList = []string{"testString"}
				createObjectOptionsModel.LabelI18n = core.StringPtr("testString")
				createObjectOptionsModel.Label = core.StringPtr("testString")
				createObjectOptionsModel.Tags = []string{"testString"}
				createObjectOptionsModel.Created = CreateMockDateTime()
				createObjectOptionsModel.Updated = CreateMockDateTime()
				createObjectOptionsModel.ShortDescription = core.StringPtr("testString")
				createObjectOptionsModel.ShortDescriptionI18n = core.StringPtr("testString")
				createObjectOptionsModel.Kind = core.StringPtr("testString")
				createObjectOptionsModel.Publish = publishObjectModel
				createObjectOptionsModel.State = stateModel
				createObjectOptionsModel.CatalogID = core.StringPtr("testString")
				createObjectOptionsModel.CatalogName = core.StringPtr("testString")
				createObjectOptionsModel.Data = map[string]interface{}{"anyKey": "anyValue"}
				createObjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.CreateObject(createObjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.CreateObjectWithContext(ctx, createObjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.CreateObject(createObjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.CreateObjectWithContext(ctx, createObjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke CreateObject with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the PublishObject model
				publishObjectModel := new(catalogmanagementv1.PublishObject)
				publishObjectModel.PermitIbmPublicPublish = core.BoolPtr(true)
				publishObjectModel.IbmApproved = core.BoolPtr(true)
				publishObjectModel.PublicApproved = core.BoolPtr(true)
				publishObjectModel.PortalApprovalRecord = core.StringPtr("testString")
				publishObjectModel.PortalURL = core.StringPtr("testString")

				// Construct an instance of the State model
				stateModel := new(catalogmanagementv1.State)
				stateModel.Current = core.StringPtr("testString")
				stateModel.CurrentEntered = CreateMockDateTime()
				stateModel.Pending = core.StringPtr("testString")
				stateModel.PendingRequested = CreateMockDateTime()
				stateModel.Previous = core.StringPtr("testString")

				// Construct an instance of the CreateObjectOptions model
				createObjectOptionsModel := new(catalogmanagementv1.CreateObjectOptions)
				createObjectOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				createObjectOptionsModel.ID = core.StringPtr("testString")
				createObjectOptionsModel.Name = core.StringPtr("testString")
				createObjectOptionsModel.Rev = core.StringPtr("testString")
				createObjectOptionsModel.Crn = core.StringPtr("testString")
				createObjectOptionsModel.URL = core.StringPtr("testString")
				createObjectOptionsModel.ParentID = core.StringPtr("testString")
				createObjectOptionsModel.AllowList = []string{"testString"}
				createObjectOptionsModel.LabelI18n = core.StringPtr("testString")
				createObjectOptionsModel.Label = core.StringPtr("testString")
				createObjectOptionsModel.Tags = []string{"testString"}
				createObjectOptionsModel.Created = CreateMockDateTime()
				createObjectOptionsModel.Updated = CreateMockDateTime()
				createObjectOptionsModel.ShortDescription = core.StringPtr("testString")
				createObjectOptionsModel.ShortDescriptionI18n = core.StringPtr("testString")
				createObjectOptionsModel.Kind = core.StringPtr("testString")
				createObjectOptionsModel.Publish = publishObjectModel
				createObjectOptionsModel.State = stateModel
				createObjectOptionsModel.CatalogID = core.StringPtr("testString")
				createObjectOptionsModel.CatalogName = core.StringPtr("testString")
				createObjectOptionsModel.Data = map[string]interface{}{"anyKey": "anyValue"}
				createObjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.CreateObject(createObjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateObjectOptions model with no property values
				createObjectOptionsModelNew := new(catalogmanagementv1.CreateObjectOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = catalogManagementService.CreateObject(createObjectOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetObject(getObjectOptions *GetObjectOptions) - Operation response error`, func() {
		getObjectPath := "/catalogs/testString/objects/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getObjectPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetObject with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetObjectOptions model
				getObjectOptionsModel := new(catalogmanagementv1.GetObjectOptions)
				getObjectOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				getObjectOptionsModel.ObjectIdentifier = core.StringPtr("testString")
				getObjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.GetObject(getObjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.GetObject(getObjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetObject(getObjectOptions *GetObjectOptions)`, func() {
		getObjectPath := "/catalogs/testString/objects/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getObjectPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "_rev": "Rev", "crn": "Crn", "url": "URL", "parent_id": "ParentID", "allow_list": ["AllowList"], "label_i18n": "LabelI18n", "label": "Label", "tags": ["Tags"], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "short_description": "ShortDescription", "short_description_i18n": "ShortDescriptionI18n", "kind": "Kind", "publish": {"permit_ibm_public_publish": true, "ibm_approved": false, "public_approved": true, "portal_approval_record": "PortalApprovalRecord", "portal_url": "PortalURL"}, "state": {"current": "Current", "current_entered": "2019-01-01T12:00:00", "pending": "Pending", "pending_requested": "2019-01-01T12:00:00", "previous": "Previous"}, "catalog_id": "CatalogID", "catalog_name": "CatalogName", "data": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke GetObject successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.GetObject(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetObjectOptions model
				getObjectOptionsModel := new(catalogmanagementv1.GetObjectOptions)
				getObjectOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				getObjectOptionsModel.ObjectIdentifier = core.StringPtr("testString")
				getObjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.GetObject(getObjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetObjectWithContext(ctx, getObjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.GetObject(getObjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.GetObjectWithContext(ctx, getObjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetObject with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetObjectOptions model
				getObjectOptionsModel := new(catalogmanagementv1.GetObjectOptions)
				getObjectOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				getObjectOptionsModel.ObjectIdentifier = core.StringPtr("testString")
				getObjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.GetObject(getObjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetObjectOptions model with no property values
				getObjectOptionsModelNew := new(catalogmanagementv1.GetObjectOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = catalogManagementService.GetObject(getObjectOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceObject(replaceObjectOptions *ReplaceObjectOptions) - Operation response error`, func() {
		replaceObjectPath := "/catalogs/testString/objects/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceObjectPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceObject with error: Operation response processing error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the PublishObject model
				publishObjectModel := new(catalogmanagementv1.PublishObject)
				publishObjectModel.PermitIbmPublicPublish = core.BoolPtr(true)
				publishObjectModel.IbmApproved = core.BoolPtr(true)
				publishObjectModel.PublicApproved = core.BoolPtr(true)
				publishObjectModel.PortalApprovalRecord = core.StringPtr("testString")
				publishObjectModel.PortalURL = core.StringPtr("testString")

				// Construct an instance of the State model
				stateModel := new(catalogmanagementv1.State)
				stateModel.Current = core.StringPtr("testString")
				stateModel.CurrentEntered = CreateMockDateTime()
				stateModel.Pending = core.StringPtr("testString")
				stateModel.PendingRequested = CreateMockDateTime()
				stateModel.Previous = core.StringPtr("testString")

				// Construct an instance of the ReplaceObjectOptions model
				replaceObjectOptionsModel := new(catalogmanagementv1.ReplaceObjectOptions)
				replaceObjectOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				replaceObjectOptionsModel.ObjectIdentifier = core.StringPtr("testString")
				replaceObjectOptionsModel.ID = core.StringPtr("testString")
				replaceObjectOptionsModel.Name = core.StringPtr("testString")
				replaceObjectOptionsModel.Rev = core.StringPtr("testString")
				replaceObjectOptionsModel.Crn = core.StringPtr("testString")
				replaceObjectOptionsModel.URL = core.StringPtr("testString")
				replaceObjectOptionsModel.ParentID = core.StringPtr("testString")
				replaceObjectOptionsModel.AllowList = []string{"testString"}
				replaceObjectOptionsModel.LabelI18n = core.StringPtr("testString")
				replaceObjectOptionsModel.Label = core.StringPtr("testString")
				replaceObjectOptionsModel.Tags = []string{"testString"}
				replaceObjectOptionsModel.Created = CreateMockDateTime()
				replaceObjectOptionsModel.Updated = CreateMockDateTime()
				replaceObjectOptionsModel.ShortDescription = core.StringPtr("testString")
				replaceObjectOptionsModel.ShortDescriptionI18n = core.StringPtr("testString")
				replaceObjectOptionsModel.Kind = core.StringPtr("testString")
				replaceObjectOptionsModel.Publish = publishObjectModel
				replaceObjectOptionsModel.State = stateModel
				replaceObjectOptionsModel.CatalogID = core.StringPtr("testString")
				replaceObjectOptionsModel.CatalogName = core.StringPtr("testString")
				replaceObjectOptionsModel.Data = map[string]interface{}{"anyKey": "anyValue"}
				replaceObjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := catalogManagementService.ReplaceObject(replaceObjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				catalogManagementService.EnableRetries(0, 0)
				result, response, operationErr = catalogManagementService.ReplaceObject(replaceObjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ReplaceObject(replaceObjectOptions *ReplaceObjectOptions)`, func() {
		replaceObjectPath := "/catalogs/testString/objects/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceObjectPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "_rev": "Rev", "crn": "Crn", "url": "URL", "parent_id": "ParentID", "allow_list": ["AllowList"], "label_i18n": "LabelI18n", "label": "Label", "tags": ["Tags"], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "short_description": "ShortDescription", "short_description_i18n": "ShortDescriptionI18n", "kind": "Kind", "publish": {"permit_ibm_public_publish": true, "ibm_approved": false, "public_approved": true, "portal_approval_record": "PortalApprovalRecord", "portal_url": "PortalURL"}, "state": {"current": "Current", "current_entered": "2019-01-01T12:00:00", "pending": "Pending", "pending_requested": "2019-01-01T12:00:00", "previous": "Previous"}, "catalog_id": "CatalogID", "catalog_name": "CatalogName", "data": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke ReplaceObject successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := catalogManagementService.ReplaceObject(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PublishObject model
				publishObjectModel := new(catalogmanagementv1.PublishObject)
				publishObjectModel.PermitIbmPublicPublish = core.BoolPtr(true)
				publishObjectModel.IbmApproved = core.BoolPtr(true)
				publishObjectModel.PublicApproved = core.BoolPtr(true)
				publishObjectModel.PortalApprovalRecord = core.StringPtr("testString")
				publishObjectModel.PortalURL = core.StringPtr("testString")

				// Construct an instance of the State model
				stateModel := new(catalogmanagementv1.State)
				stateModel.Current = core.StringPtr("testString")
				stateModel.CurrentEntered = CreateMockDateTime()
				stateModel.Pending = core.StringPtr("testString")
				stateModel.PendingRequested = CreateMockDateTime()
				stateModel.Previous = core.StringPtr("testString")

				// Construct an instance of the ReplaceObjectOptions model
				replaceObjectOptionsModel := new(catalogmanagementv1.ReplaceObjectOptions)
				replaceObjectOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				replaceObjectOptionsModel.ObjectIdentifier = core.StringPtr("testString")
				replaceObjectOptionsModel.ID = core.StringPtr("testString")
				replaceObjectOptionsModel.Name = core.StringPtr("testString")
				replaceObjectOptionsModel.Rev = core.StringPtr("testString")
				replaceObjectOptionsModel.Crn = core.StringPtr("testString")
				replaceObjectOptionsModel.URL = core.StringPtr("testString")
				replaceObjectOptionsModel.ParentID = core.StringPtr("testString")
				replaceObjectOptionsModel.AllowList = []string{"testString"}
				replaceObjectOptionsModel.LabelI18n = core.StringPtr("testString")
				replaceObjectOptionsModel.Label = core.StringPtr("testString")
				replaceObjectOptionsModel.Tags = []string{"testString"}
				replaceObjectOptionsModel.Created = CreateMockDateTime()
				replaceObjectOptionsModel.Updated = CreateMockDateTime()
				replaceObjectOptionsModel.ShortDescription = core.StringPtr("testString")
				replaceObjectOptionsModel.ShortDescriptionI18n = core.StringPtr("testString")
				replaceObjectOptionsModel.Kind = core.StringPtr("testString")
				replaceObjectOptionsModel.Publish = publishObjectModel
				replaceObjectOptionsModel.State = stateModel
				replaceObjectOptionsModel.CatalogID = core.StringPtr("testString")
				replaceObjectOptionsModel.CatalogName = core.StringPtr("testString")
				replaceObjectOptionsModel.Data = map[string]interface{}{"anyKey": "anyValue"}
				replaceObjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = catalogManagementService.ReplaceObject(replaceObjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.ReplaceObjectWithContext(ctx, replaceObjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				result, response, operationErr = catalogManagementService.ReplaceObject(replaceObjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = catalogManagementService.ReplaceObjectWithContext(ctx, replaceObjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ReplaceObject with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the PublishObject model
				publishObjectModel := new(catalogmanagementv1.PublishObject)
				publishObjectModel.PermitIbmPublicPublish = core.BoolPtr(true)
				publishObjectModel.IbmApproved = core.BoolPtr(true)
				publishObjectModel.PublicApproved = core.BoolPtr(true)
				publishObjectModel.PortalApprovalRecord = core.StringPtr("testString")
				publishObjectModel.PortalURL = core.StringPtr("testString")

				// Construct an instance of the State model
				stateModel := new(catalogmanagementv1.State)
				stateModel.Current = core.StringPtr("testString")
				stateModel.CurrentEntered = CreateMockDateTime()
				stateModel.Pending = core.StringPtr("testString")
				stateModel.PendingRequested = CreateMockDateTime()
				stateModel.Previous = core.StringPtr("testString")

				// Construct an instance of the ReplaceObjectOptions model
				replaceObjectOptionsModel := new(catalogmanagementv1.ReplaceObjectOptions)
				replaceObjectOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				replaceObjectOptionsModel.ObjectIdentifier = core.StringPtr("testString")
				replaceObjectOptionsModel.ID = core.StringPtr("testString")
				replaceObjectOptionsModel.Name = core.StringPtr("testString")
				replaceObjectOptionsModel.Rev = core.StringPtr("testString")
				replaceObjectOptionsModel.Crn = core.StringPtr("testString")
				replaceObjectOptionsModel.URL = core.StringPtr("testString")
				replaceObjectOptionsModel.ParentID = core.StringPtr("testString")
				replaceObjectOptionsModel.AllowList = []string{"testString"}
				replaceObjectOptionsModel.LabelI18n = core.StringPtr("testString")
				replaceObjectOptionsModel.Label = core.StringPtr("testString")
				replaceObjectOptionsModel.Tags = []string{"testString"}
				replaceObjectOptionsModel.Created = CreateMockDateTime()
				replaceObjectOptionsModel.Updated = CreateMockDateTime()
				replaceObjectOptionsModel.ShortDescription = core.StringPtr("testString")
				replaceObjectOptionsModel.ShortDescriptionI18n = core.StringPtr("testString")
				replaceObjectOptionsModel.Kind = core.StringPtr("testString")
				replaceObjectOptionsModel.Publish = publishObjectModel
				replaceObjectOptionsModel.State = stateModel
				replaceObjectOptionsModel.CatalogID = core.StringPtr("testString")
				replaceObjectOptionsModel.CatalogName = core.StringPtr("testString")
				replaceObjectOptionsModel.Data = map[string]interface{}{"anyKey": "anyValue"}
				replaceObjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := catalogManagementService.ReplaceObject(replaceObjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceObjectOptions model with no property values
				replaceObjectOptionsModelNew := new(catalogmanagementv1.ReplaceObjectOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = catalogManagementService.ReplaceObject(replaceObjectOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteObject(deleteObjectOptions *DeleteObjectOptions)`, func() {
		deleteObjectPath := "/catalogs/testString/objects/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteObjectPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteObject successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := catalogManagementService.DeleteObject(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteObjectOptions model
				deleteObjectOptionsModel := new(catalogmanagementv1.DeleteObjectOptions)
				deleteObjectOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				deleteObjectOptionsModel.ObjectIdentifier = core.StringPtr("testString")
				deleteObjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = catalogManagementService.DeleteObject(deleteObjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				response, operationErr = catalogManagementService.DeleteObject(deleteObjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteObject with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the DeleteObjectOptions model
				deleteObjectOptionsModel := new(catalogmanagementv1.DeleteObjectOptions)
				deleteObjectOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				deleteObjectOptionsModel.ObjectIdentifier = core.StringPtr("testString")
				deleteObjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := catalogManagementService.DeleteObject(deleteObjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteObjectOptions model with no property values
				deleteObjectOptionsModelNew := new(catalogmanagementv1.DeleteObjectOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = catalogManagementService.DeleteObject(deleteObjectOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetObjectAudit(getObjectAuditOptions *GetObjectAuditOptions)`, func() {
		getObjectAuditPath := "/catalogs/testString/offerings/testString/audit"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getObjectAuditPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["id"]).To(Equal([]string{"testString"}))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetObjectAudit successfully`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())
				catalogManagementService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := catalogManagementService.GetObjectAudit(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the GetObjectAuditOptions model
				getObjectAuditOptionsModel := new(catalogmanagementv1.GetObjectAuditOptions)
				getObjectAuditOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				getObjectAuditOptionsModel.ObjectIdentifier = core.StringPtr("testString")
				getObjectAuditOptionsModel.ID = core.StringPtr("testString")
				getObjectAuditOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = catalogManagementService.GetObjectAudit(getObjectAuditOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				catalogManagementService.DisableRetries()
				response, operationErr = catalogManagementService.GetObjectAudit(getObjectAuditOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke GetObjectAudit with error: Operation validation and request error`, func() {
				catalogManagementService, serviceErr := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(catalogManagementService).ToNot(BeNil())

				// Construct an instance of the GetObjectAuditOptions model
				getObjectAuditOptionsModel := new(catalogmanagementv1.GetObjectAuditOptions)
				getObjectAuditOptionsModel.CatalogIdentifier = core.StringPtr("testString")
				getObjectAuditOptionsModel.ObjectIdentifier = core.StringPtr("testString")
				getObjectAuditOptionsModel.ID = core.StringPtr("testString")
				getObjectAuditOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := catalogManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := catalogManagementService.GetObjectAudit(getObjectAuditOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the GetObjectAuditOptions model with no property values
				getObjectAuditOptionsModelNew := new(catalogmanagementv1.GetObjectAuditOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = catalogManagementService.GetObjectAudit(getObjectAuditOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			catalogManagementService, _ := catalogmanagementv1.NewCatalogManagementV1(&catalogmanagementv1.CatalogManagementV1Options{
				URL:           "http://catalogmanagementv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewAccountPublishVersionOptions successfully`, func() {
				// Construct an instance of the AccountPublishVersionOptions model
				versionLocID := "testString"
				accountPublishVersionOptionsModel := catalogManagementService.NewAccountPublishVersionOptions(versionLocID)
				accountPublishVersionOptionsModel.SetVersionLocID("testString")
				accountPublishVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(accountPublishVersionOptionsModel).ToNot(BeNil())
				Expect(accountPublishVersionOptionsModel.VersionLocID).To(Equal(core.StringPtr("testString")))
				Expect(accountPublishVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCanDeploySchematicsOptions successfully`, func() {
				// Construct an instance of the CanDeploySchematicsOptions model
				versionLocID := "testString"
				clusterID := "testString"
				region := "testString"
				canDeploySchematicsOptionsModel := catalogManagementService.NewCanDeploySchematicsOptions(versionLocID, clusterID, region)
				canDeploySchematicsOptionsModel.SetVersionLocID("testString")
				canDeploySchematicsOptionsModel.SetClusterID("testString")
				canDeploySchematicsOptionsModel.SetRegion("testString")
				canDeploySchematicsOptionsModel.SetNamespace("testString")
				canDeploySchematicsOptionsModel.SetResourceGroupID("testString")
				canDeploySchematicsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(canDeploySchematicsOptionsModel).ToNot(BeNil())
				Expect(canDeploySchematicsOptionsModel.VersionLocID).To(Equal(core.StringPtr("testString")))
				Expect(canDeploySchematicsOptionsModel.ClusterID).To(Equal(core.StringPtr("testString")))
				Expect(canDeploySchematicsOptionsModel.Region).To(Equal(core.StringPtr("testString")))
				Expect(canDeploySchematicsOptionsModel.Namespace).To(Equal(core.StringPtr("testString")))
				Expect(canDeploySchematicsOptionsModel.ResourceGroupID).To(Equal(core.StringPtr("testString")))
				Expect(canDeploySchematicsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCommitVersionOptions successfully`, func() {
				// Construct an instance of the CommitVersionOptions model
				versionLocID := "testString"
				commitVersionOptionsModel := catalogManagementService.NewCommitVersionOptions(versionLocID)
				commitVersionOptionsModel.SetVersionLocID("testString")
				commitVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(commitVersionOptionsModel).ToNot(BeNil())
				Expect(commitVersionOptionsModel.VersionLocID).To(Equal(core.StringPtr("testString")))
				Expect(commitVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCopyVersionOptions successfully`, func() {
				// Construct an instance of the CopyVersionOptions model
				versionLocID := "testString"
				copyVersionOptionsModel := catalogManagementService.NewCopyVersionOptions(versionLocID)
				copyVersionOptionsModel.SetVersionLocID("testString")
				copyVersionOptionsModel.SetTags([]string{"testString"})
				copyVersionOptionsModel.SetTargetKinds([]string{"testString"})
				copyVersionOptionsModel.SetContent([]int64{int64(38)})
				copyVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(copyVersionOptionsModel).ToNot(BeNil())
				Expect(copyVersionOptionsModel.VersionLocID).To(Equal(core.StringPtr("testString")))
				Expect(copyVersionOptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(copyVersionOptionsModel.TargetKinds).To(Equal([]string{"testString"}))
				Expect(copyVersionOptionsModel.Content).To(Equal([]int64{int64(38)}))
				Expect(copyVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateCatalogOptions successfully`, func() {
				// Construct an instance of the Feature model
				featureModel := new(catalogmanagementv1.Feature)
				Expect(featureModel).ToNot(BeNil())
				featureModel.Title = core.StringPtr("testString")
				featureModel.Description = core.StringPtr("testString")
				Expect(featureModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(featureModel.Description).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the FilterTerms model
				filterTermsModel := new(catalogmanagementv1.FilterTerms)
				Expect(filterTermsModel).ToNot(BeNil())
				filterTermsModel.FilterTerms = []string{"testString"}
				Expect(filterTermsModel.FilterTerms).To(Equal([]string{"testString"}))

				// Construct an instance of the CategoryFilter model
				categoryFilterModel := new(catalogmanagementv1.CategoryFilter)
				Expect(categoryFilterModel).ToNot(BeNil())
				categoryFilterModel.Include = core.BoolPtr(true)
				categoryFilterModel.Filter = filterTermsModel
				Expect(categoryFilterModel.Include).To(Equal(core.BoolPtr(true)))
				Expect(categoryFilterModel.Filter).To(Equal(filterTermsModel))

				// Construct an instance of the IDFilter model
				idFilterModel := new(catalogmanagementv1.IDFilter)
				Expect(idFilterModel).ToNot(BeNil())
				idFilterModel.Include = filterTermsModel
				idFilterModel.Exclude = filterTermsModel
				Expect(idFilterModel.Include).To(Equal(filterTermsModel))
				Expect(idFilterModel.Exclude).To(Equal(filterTermsModel))

				// Construct an instance of the Filters model
				filtersModel := new(catalogmanagementv1.Filters)
				Expect(filtersModel).ToNot(BeNil())
				filtersModel.IncludeAll = core.BoolPtr(true)
				filtersModel.CategoryFilters = make(map[string]catalogmanagementv1.CategoryFilter)
				filtersModel.IdFilters = idFilterModel
				filtersModel.CategoryFilters["foo"] = *categoryFilterModel
				Expect(filtersModel.IncludeAll).To(Equal(core.BoolPtr(true)))
				Expect(filtersModel.IdFilters).To(Equal(idFilterModel))
				Expect(filtersModel.CategoryFilters["foo"]).To(Equal(*categoryFilterModel))

				// Construct an instance of the SyndicationCluster model
				syndicationClusterModel := new(catalogmanagementv1.SyndicationCluster)
				Expect(syndicationClusterModel).ToNot(BeNil())
				syndicationClusterModel.Region = core.StringPtr("testString")
				syndicationClusterModel.ID = core.StringPtr("testString")
				syndicationClusterModel.Name = core.StringPtr("testString")
				syndicationClusterModel.ResourceGroupName = core.StringPtr("testString")
				syndicationClusterModel.Type = core.StringPtr("testString")
				syndicationClusterModel.Namespaces = []string{"testString"}
				syndicationClusterModel.AllNamespaces = core.BoolPtr(true)
				Expect(syndicationClusterModel.Region).To(Equal(core.StringPtr("testString")))
				Expect(syndicationClusterModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(syndicationClusterModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(syndicationClusterModel.ResourceGroupName).To(Equal(core.StringPtr("testString")))
				Expect(syndicationClusterModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(syndicationClusterModel.Namespaces).To(Equal([]string{"testString"}))
				Expect(syndicationClusterModel.AllNamespaces).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the SyndicationHistory model
				syndicationHistoryModel := new(catalogmanagementv1.SyndicationHistory)
				Expect(syndicationHistoryModel).ToNot(BeNil())
				syndicationHistoryModel.Namespaces = []string{"testString"}
				syndicationHistoryModel.Clusters = []catalogmanagementv1.SyndicationCluster{*syndicationClusterModel}
				syndicationHistoryModel.LastRun = CreateMockDateTime()
				Expect(syndicationHistoryModel.Namespaces).To(Equal([]string{"testString"}))
				Expect(syndicationHistoryModel.Clusters).To(Equal([]catalogmanagementv1.SyndicationCluster{*syndicationClusterModel}))
				Expect(syndicationHistoryModel.LastRun).To(Equal(CreateMockDateTime()))

				// Construct an instance of the SyndicationAuthorization model
				syndicationAuthorizationModel := new(catalogmanagementv1.SyndicationAuthorization)
				Expect(syndicationAuthorizationModel).ToNot(BeNil())
				syndicationAuthorizationModel.Token = core.StringPtr("testString")
				syndicationAuthorizationModel.LastRun = CreateMockDateTime()
				Expect(syndicationAuthorizationModel.Token).To(Equal(core.StringPtr("testString")))
				Expect(syndicationAuthorizationModel.LastRun).To(Equal(CreateMockDateTime()))

				// Construct an instance of the SyndicationResource model
				syndicationResourceModel := new(catalogmanagementv1.SyndicationResource)
				Expect(syndicationResourceModel).ToNot(BeNil())
				syndicationResourceModel.RemoveRelatedComponents = core.BoolPtr(true)
				syndicationResourceModel.Clusters = []catalogmanagementv1.SyndicationCluster{*syndicationClusterModel}
				syndicationResourceModel.History = syndicationHistoryModel
				syndicationResourceModel.Authorization = syndicationAuthorizationModel
				Expect(syndicationResourceModel.RemoveRelatedComponents).To(Equal(core.BoolPtr(true)))
				Expect(syndicationResourceModel.Clusters).To(Equal([]catalogmanagementv1.SyndicationCluster{*syndicationClusterModel}))
				Expect(syndicationResourceModel.History).To(Equal(syndicationHistoryModel))
				Expect(syndicationResourceModel.Authorization).To(Equal(syndicationAuthorizationModel))

				// Construct an instance of the CreateCatalogOptions model
				createCatalogOptionsModel := catalogManagementService.NewCreateCatalogOptions()
				createCatalogOptionsModel.SetID("testString")
				createCatalogOptionsModel.SetRev("testString")
				createCatalogOptionsModel.SetLabel("testString")
				createCatalogOptionsModel.SetShortDescription("testString")
				createCatalogOptionsModel.SetCatalogIconURL("testString")
				createCatalogOptionsModel.SetTags([]string{"testString"})
				createCatalogOptionsModel.SetURL("testString")
				createCatalogOptionsModel.SetCrn("testString")
				createCatalogOptionsModel.SetOfferingsURL("testString")
				createCatalogOptionsModel.SetFeatures([]catalogmanagementv1.Feature{*featureModel})
				createCatalogOptionsModel.SetDisabled(true)
				createCatalogOptionsModel.SetCreated(CreateMockDateTime())
				createCatalogOptionsModel.SetUpdated(CreateMockDateTime())
				createCatalogOptionsModel.SetResourceGroupID("testString")
				createCatalogOptionsModel.SetOwningAccount("testString")
				createCatalogOptionsModel.SetCatalogFilters(filtersModel)
				createCatalogOptionsModel.SetSyndicationSettings(syndicationResourceModel)
				createCatalogOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createCatalogOptionsModel).ToNot(BeNil())
				Expect(createCatalogOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(createCatalogOptionsModel.Rev).To(Equal(core.StringPtr("testString")))
				Expect(createCatalogOptionsModel.Label).To(Equal(core.StringPtr("testString")))
				Expect(createCatalogOptionsModel.ShortDescription).To(Equal(core.StringPtr("testString")))
				Expect(createCatalogOptionsModel.CatalogIconURL).To(Equal(core.StringPtr("testString")))
				Expect(createCatalogOptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(createCatalogOptionsModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(createCatalogOptionsModel.Crn).To(Equal(core.StringPtr("testString")))
				Expect(createCatalogOptionsModel.OfferingsURL).To(Equal(core.StringPtr("testString")))
				Expect(createCatalogOptionsModel.Features).To(Equal([]catalogmanagementv1.Feature{*featureModel}))
				Expect(createCatalogOptionsModel.Disabled).To(Equal(core.BoolPtr(true)))
				Expect(createCatalogOptionsModel.Created).To(Equal(CreateMockDateTime()))
				Expect(createCatalogOptionsModel.Updated).To(Equal(CreateMockDateTime()))
				Expect(createCatalogOptionsModel.ResourceGroupID).To(Equal(core.StringPtr("testString")))
				Expect(createCatalogOptionsModel.OwningAccount).To(Equal(core.StringPtr("testString")))
				Expect(createCatalogOptionsModel.CatalogFilters).To(Equal(filtersModel))
				Expect(createCatalogOptionsModel.SyndicationSettings).To(Equal(syndicationResourceModel))
				Expect(createCatalogOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateLicenseEntitlementOptions successfully`, func() {
				// Construct an instance of the CreateLicenseEntitlementOptions model
				createLicenseEntitlementOptionsModel := catalogManagementService.NewCreateLicenseEntitlementOptions()
				createLicenseEntitlementOptionsModel.SetName("testString")
				createLicenseEntitlementOptionsModel.SetEffectiveFrom("testString")
				createLicenseEntitlementOptionsModel.SetEffectiveUntil("testString")
				createLicenseEntitlementOptionsModel.SetVersionID("testString")
				createLicenseEntitlementOptionsModel.SetLicenseID("testString")
				createLicenseEntitlementOptionsModel.SetLicenseOwnerID("testString")
				createLicenseEntitlementOptionsModel.SetLicenseProviderID("testString")
				createLicenseEntitlementOptionsModel.SetLicenseProductID("testString")
				createLicenseEntitlementOptionsModel.SetAccountID("testString")
				createLicenseEntitlementOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createLicenseEntitlementOptionsModel).ToNot(BeNil())
				Expect(createLicenseEntitlementOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createLicenseEntitlementOptionsModel.EffectiveFrom).To(Equal(core.StringPtr("testString")))
				Expect(createLicenseEntitlementOptionsModel.EffectiveUntil).To(Equal(core.StringPtr("testString")))
				Expect(createLicenseEntitlementOptionsModel.VersionID).To(Equal(core.StringPtr("testString")))
				Expect(createLicenseEntitlementOptionsModel.LicenseID).To(Equal(core.StringPtr("testString")))
				Expect(createLicenseEntitlementOptionsModel.LicenseOwnerID).To(Equal(core.StringPtr("testString")))
				Expect(createLicenseEntitlementOptionsModel.LicenseProviderID).To(Equal(core.StringPtr("testString")))
				Expect(createLicenseEntitlementOptionsModel.LicenseProductID).To(Equal(core.StringPtr("testString")))
				Expect(createLicenseEntitlementOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(createLicenseEntitlementOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateObjectOptions successfully`, func() {
				// Construct an instance of the PublishObject model
				publishObjectModel := new(catalogmanagementv1.PublishObject)
				Expect(publishObjectModel).ToNot(BeNil())
				publishObjectModel.PermitIbmPublicPublish = core.BoolPtr(true)
				publishObjectModel.IbmApproved = core.BoolPtr(true)
				publishObjectModel.PublicApproved = core.BoolPtr(true)
				publishObjectModel.PortalApprovalRecord = core.StringPtr("testString")
				publishObjectModel.PortalURL = core.StringPtr("testString")
				Expect(publishObjectModel.PermitIbmPublicPublish).To(Equal(core.BoolPtr(true)))
				Expect(publishObjectModel.IbmApproved).To(Equal(core.BoolPtr(true)))
				Expect(publishObjectModel.PublicApproved).To(Equal(core.BoolPtr(true)))
				Expect(publishObjectModel.PortalApprovalRecord).To(Equal(core.StringPtr("testString")))
				Expect(publishObjectModel.PortalURL).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the State model
				stateModel := new(catalogmanagementv1.State)
				Expect(stateModel).ToNot(BeNil())
				stateModel.Current = core.StringPtr("testString")
				stateModel.CurrentEntered = CreateMockDateTime()
				stateModel.Pending = core.StringPtr("testString")
				stateModel.PendingRequested = CreateMockDateTime()
				stateModel.Previous = core.StringPtr("testString")
				Expect(stateModel.Current).To(Equal(core.StringPtr("testString")))
				Expect(stateModel.CurrentEntered).To(Equal(CreateMockDateTime()))
				Expect(stateModel.Pending).To(Equal(core.StringPtr("testString")))
				Expect(stateModel.PendingRequested).To(Equal(CreateMockDateTime()))
				Expect(stateModel.Previous).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CreateObjectOptions model
				catalogIdentifier := "testString"
				createObjectOptionsModel := catalogManagementService.NewCreateObjectOptions(catalogIdentifier)
				createObjectOptionsModel.SetCatalogIdentifier("testString")
				createObjectOptionsModel.SetID("testString")
				createObjectOptionsModel.SetName("testString")
				createObjectOptionsModel.SetRev("testString")
				createObjectOptionsModel.SetCrn("testString")
				createObjectOptionsModel.SetURL("testString")
				createObjectOptionsModel.SetParentID("testString")
				createObjectOptionsModel.SetAllowList([]string{"testString"})
				createObjectOptionsModel.SetLabelI18n("testString")
				createObjectOptionsModel.SetLabel("testString")
				createObjectOptionsModel.SetTags([]string{"testString"})
				createObjectOptionsModel.SetCreated(CreateMockDateTime())
				createObjectOptionsModel.SetUpdated(CreateMockDateTime())
				createObjectOptionsModel.SetShortDescription("testString")
				createObjectOptionsModel.SetShortDescriptionI18n("testString")
				createObjectOptionsModel.SetKind("testString")
				createObjectOptionsModel.SetPublish(publishObjectModel)
				createObjectOptionsModel.SetState(stateModel)
				createObjectOptionsModel.SetCatalogID("testString")
				createObjectOptionsModel.SetCatalogName("testString")
				createObjectOptionsModel.SetData(map[string]interface{}{"anyKey": "anyValue"})
				createObjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createObjectOptionsModel).ToNot(BeNil())
				Expect(createObjectOptionsModel.CatalogIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(createObjectOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(createObjectOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createObjectOptionsModel.Rev).To(Equal(core.StringPtr("testString")))
				Expect(createObjectOptionsModel.Crn).To(Equal(core.StringPtr("testString")))
				Expect(createObjectOptionsModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(createObjectOptionsModel.ParentID).To(Equal(core.StringPtr("testString")))
				Expect(createObjectOptionsModel.AllowList).To(Equal([]string{"testString"}))
				Expect(createObjectOptionsModel.LabelI18n).To(Equal(core.StringPtr("testString")))
				Expect(createObjectOptionsModel.Label).To(Equal(core.StringPtr("testString")))
				Expect(createObjectOptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(createObjectOptionsModel.Created).To(Equal(CreateMockDateTime()))
				Expect(createObjectOptionsModel.Updated).To(Equal(CreateMockDateTime()))
				Expect(createObjectOptionsModel.ShortDescription).To(Equal(core.StringPtr("testString")))
				Expect(createObjectOptionsModel.ShortDescriptionI18n).To(Equal(core.StringPtr("testString")))
				Expect(createObjectOptionsModel.Kind).To(Equal(core.StringPtr("testString")))
				Expect(createObjectOptionsModel.Publish).To(Equal(publishObjectModel))
				Expect(createObjectOptionsModel.State).To(Equal(stateModel))
				Expect(createObjectOptionsModel.CatalogID).To(Equal(core.StringPtr("testString")))
				Expect(createObjectOptionsModel.CatalogName).To(Equal(core.StringPtr("testString")))
				Expect(createObjectOptionsModel.Data).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(createObjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateOfferingOptions successfully`, func() {
				// Construct an instance of the Rating model
				ratingModel := new(catalogmanagementv1.Rating)
				Expect(ratingModel).ToNot(BeNil())
				ratingModel.OneStarCount = core.Int64Ptr(int64(38))
				ratingModel.TwoStarCount = core.Int64Ptr(int64(38))
				ratingModel.ThreeStarCount = core.Int64Ptr(int64(38))
				ratingModel.FourStarCount = core.Int64Ptr(int64(38))
				Expect(ratingModel.OneStarCount).To(Equal(core.Int64Ptr(int64(38))))
				Expect(ratingModel.TwoStarCount).To(Equal(core.Int64Ptr(int64(38))))
				Expect(ratingModel.ThreeStarCount).To(Equal(core.Int64Ptr(int64(38))))
				Expect(ratingModel.FourStarCount).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the Feature model
				featureModel := new(catalogmanagementv1.Feature)
				Expect(featureModel).ToNot(BeNil())
				featureModel.Title = core.StringPtr("testString")
				featureModel.Description = core.StringPtr("testString")
				Expect(featureModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(featureModel.Description).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Configuration model
				configurationModel := new(catalogmanagementv1.Configuration)
				Expect(configurationModel).ToNot(BeNil())
				configurationModel.Key = core.StringPtr("testString")
				configurationModel.Type = core.StringPtr("testString")
				configurationModel.DefaultValue = core.StringPtr("testString")
				configurationModel.ValueConstraint = core.StringPtr("testString")
				configurationModel.Description = core.StringPtr("testString")
				configurationModel.Required = core.BoolPtr(true)
				configurationModel.Options = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				configurationModel.Hidden = core.BoolPtr(true)
				Expect(configurationModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(configurationModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(configurationModel.DefaultValue).To(Equal(core.StringPtr("testString")))
				Expect(configurationModel.ValueConstraint).To(Equal(core.StringPtr("testString")))
				Expect(configurationModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(configurationModel.Required).To(Equal(core.BoolPtr(true)))
				Expect(configurationModel.Options).To(Equal([]interface{}{map[string]interface{}{"anyKey": "anyValue"}}))
				Expect(configurationModel.Hidden).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the Validation model
				validationModel := new(catalogmanagementv1.Validation)
				Expect(validationModel).ToNot(BeNil())
				validationModel.Validated = CreateMockDateTime()
				validationModel.Requested = CreateMockDateTime()
				validationModel.State = core.StringPtr("testString")
				validationModel.LastOperation = core.StringPtr("testString")
				validationModel.Target = map[string]interface{}{"anyKey": "anyValue"}
				Expect(validationModel.Validated).To(Equal(CreateMockDateTime()))
				Expect(validationModel.Requested).To(Equal(CreateMockDateTime()))
				Expect(validationModel.State).To(Equal(core.StringPtr("testString")))
				Expect(validationModel.LastOperation).To(Equal(core.StringPtr("testString")))
				Expect(validationModel.Target).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))

				// Construct an instance of the Resource model
				resourceModel := new(catalogmanagementv1.Resource)
				Expect(resourceModel).ToNot(BeNil())
				resourceModel.Type = core.StringPtr("mem")
				resourceModel.Value = core.StringPtr("testString")
				Expect(resourceModel.Type).To(Equal(core.StringPtr("mem")))
				Expect(resourceModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Script model
				scriptModel := new(catalogmanagementv1.Script)
				Expect(scriptModel).ToNot(BeNil())
				scriptModel.Instructions = core.StringPtr("testString")
				scriptModel.Script = core.StringPtr("testString")
				scriptModel.ScriptPermission = core.StringPtr("testString")
				scriptModel.DeleteScript = core.StringPtr("testString")
				scriptModel.Scope = core.StringPtr("testString")
				Expect(scriptModel.Instructions).To(Equal(core.StringPtr("testString")))
				Expect(scriptModel.Script).To(Equal(core.StringPtr("testString")))
				Expect(scriptModel.ScriptPermission).To(Equal(core.StringPtr("testString")))
				Expect(scriptModel.DeleteScript).To(Equal(core.StringPtr("testString")))
				Expect(scriptModel.Scope).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the VersionEntitlement model
				versionEntitlementModel := new(catalogmanagementv1.VersionEntitlement)
				Expect(versionEntitlementModel).ToNot(BeNil())
				versionEntitlementModel.ProviderName = core.StringPtr("testString")
				versionEntitlementModel.ProviderID = core.StringPtr("testString")
				versionEntitlementModel.ProductID = core.StringPtr("testString")
				versionEntitlementModel.PartNumbers = []string{"testString"}
				versionEntitlementModel.ImageRepoName = core.StringPtr("testString")
				Expect(versionEntitlementModel.ProviderName).To(Equal(core.StringPtr("testString")))
				Expect(versionEntitlementModel.ProviderID).To(Equal(core.StringPtr("testString")))
				Expect(versionEntitlementModel.ProductID).To(Equal(core.StringPtr("testString")))
				Expect(versionEntitlementModel.PartNumbers).To(Equal([]string{"testString"}))
				Expect(versionEntitlementModel.ImageRepoName).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the License model
				licenseModel := new(catalogmanagementv1.License)
				Expect(licenseModel).ToNot(BeNil())
				licenseModel.ID = core.StringPtr("testString")
				licenseModel.Name = core.StringPtr("testString")
				licenseModel.Type = core.StringPtr("testString")
				licenseModel.URL = core.StringPtr("testString")
				licenseModel.Description = core.StringPtr("testString")
				Expect(licenseModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(licenseModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(licenseModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(licenseModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(licenseModel.Description).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the State model
				stateModel := new(catalogmanagementv1.State)
				Expect(stateModel).ToNot(BeNil())
				stateModel.Current = core.StringPtr("testString")
				stateModel.CurrentEntered = CreateMockDateTime()
				stateModel.Pending = core.StringPtr("testString")
				stateModel.PendingRequested = CreateMockDateTime()
				stateModel.Previous = core.StringPtr("testString")
				Expect(stateModel.Current).To(Equal(core.StringPtr("testString")))
				Expect(stateModel.CurrentEntered).To(Equal(CreateMockDateTime()))
				Expect(stateModel.Pending).To(Equal(core.StringPtr("testString")))
				Expect(stateModel.PendingRequested).To(Equal(CreateMockDateTime()))
				Expect(stateModel.Previous).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Version model
				versionModel := new(catalogmanagementv1.Version)
				Expect(versionModel).ToNot(BeNil())
				versionModel.ID = core.StringPtr("testString")
				versionModel.Rev = core.StringPtr("testString")
				versionModel.Crn = core.StringPtr("testString")
				versionModel.Version = core.StringPtr("testString")
				versionModel.Sha = core.StringPtr("testString")
				versionModel.Created = CreateMockDateTime()
				versionModel.Updated = CreateMockDateTime()
				versionModel.OfferingID = core.StringPtr("testString")
				versionModel.CatalogID = core.StringPtr("testString")
				versionModel.KindID = core.StringPtr("testString")
				versionModel.Tags = []string{"testString"}
				versionModel.RepoURL = core.StringPtr("testString")
				versionModel.SourceURL = core.StringPtr("testString")
				versionModel.TgzURL = core.StringPtr("testString")
				versionModel.Configuration = []catalogmanagementv1.Configuration{*configurationModel}
				versionModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}
				versionModel.Validation = validationModel
				versionModel.RequiredResources = []catalogmanagementv1.Resource{*resourceModel}
				versionModel.SingleInstance = core.BoolPtr(true)
				versionModel.Install = scriptModel
				versionModel.PreInstall = []catalogmanagementv1.Script{*scriptModel}
				versionModel.Entitlement = versionEntitlementModel
				versionModel.Licenses = []catalogmanagementv1.License{*licenseModel}
				versionModel.ImageManifestURL = core.StringPtr("testString")
				versionModel.Deprecated = core.BoolPtr(true)
				versionModel.PackageVersion = core.StringPtr("testString")
				versionModel.State = stateModel
				versionModel.VersionLocator = core.StringPtr("testString")
				versionModel.ConsoleURL = core.StringPtr("testString")
				versionModel.LongDescription = core.StringPtr("testString")
				versionModel.WhitelistedAccounts = []string{"testString"}
				Expect(versionModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(versionModel.Rev).To(Equal(core.StringPtr("testString")))
				Expect(versionModel.Crn).To(Equal(core.StringPtr("testString")))
				Expect(versionModel.Version).To(Equal(core.StringPtr("testString")))
				Expect(versionModel.Sha).To(Equal(core.StringPtr("testString")))
				Expect(versionModel.Created).To(Equal(CreateMockDateTime()))
				Expect(versionModel.Updated).To(Equal(CreateMockDateTime()))
				Expect(versionModel.OfferingID).To(Equal(core.StringPtr("testString")))
				Expect(versionModel.CatalogID).To(Equal(core.StringPtr("testString")))
				Expect(versionModel.KindID).To(Equal(core.StringPtr("testString")))
				Expect(versionModel.Tags).To(Equal([]string{"testString"}))
				Expect(versionModel.RepoURL).To(Equal(core.StringPtr("testString")))
				Expect(versionModel.SourceURL).To(Equal(core.StringPtr("testString")))
				Expect(versionModel.TgzURL).To(Equal(core.StringPtr("testString")))
				Expect(versionModel.Configuration).To(Equal([]catalogmanagementv1.Configuration{*configurationModel}))
				Expect(versionModel.Metadata).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(versionModel.Validation).To(Equal(validationModel))
				Expect(versionModel.RequiredResources).To(Equal([]catalogmanagementv1.Resource{*resourceModel}))
				Expect(versionModel.SingleInstance).To(Equal(core.BoolPtr(true)))
				Expect(versionModel.Install).To(Equal(scriptModel))
				Expect(versionModel.PreInstall).To(Equal([]catalogmanagementv1.Script{*scriptModel}))
				Expect(versionModel.Entitlement).To(Equal(versionEntitlementModel))
				Expect(versionModel.Licenses).To(Equal([]catalogmanagementv1.License{*licenseModel}))
				Expect(versionModel.ImageManifestURL).To(Equal(core.StringPtr("testString")))
				Expect(versionModel.Deprecated).To(Equal(core.BoolPtr(true)))
				Expect(versionModel.PackageVersion).To(Equal(core.StringPtr("testString")))
				Expect(versionModel.State).To(Equal(stateModel))
				Expect(versionModel.VersionLocator).To(Equal(core.StringPtr("testString")))
				Expect(versionModel.ConsoleURL).To(Equal(core.StringPtr("testString")))
				Expect(versionModel.LongDescription).To(Equal(core.StringPtr("testString")))
				Expect(versionModel.WhitelistedAccounts).To(Equal([]string{"testString"}))

				// Construct an instance of the Deployment model
				deploymentModel := new(catalogmanagementv1.Deployment)
				Expect(deploymentModel).ToNot(BeNil())
				deploymentModel.ID = core.StringPtr("testString")
				deploymentModel.Label = core.StringPtr("testString")
				deploymentModel.Name = core.StringPtr("testString")
				deploymentModel.ShortDescription = core.StringPtr("testString")
				deploymentModel.LongDescription = core.StringPtr("testString")
				deploymentModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}
				deploymentModel.Tags = []string{"testString"}
				deploymentModel.Created = CreateMockDateTime()
				deploymentModel.Updated = CreateMockDateTime()
				Expect(deploymentModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deploymentModel.Label).To(Equal(core.StringPtr("testString")))
				Expect(deploymentModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(deploymentModel.ShortDescription).To(Equal(core.StringPtr("testString")))
				Expect(deploymentModel.LongDescription).To(Equal(core.StringPtr("testString")))
				Expect(deploymentModel.Metadata).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(deploymentModel.Tags).To(Equal([]string{"testString"}))
				Expect(deploymentModel.Created).To(Equal(CreateMockDateTime()))
				Expect(deploymentModel.Updated).To(Equal(CreateMockDateTime()))

				// Construct an instance of the Plan model
				planModel := new(catalogmanagementv1.Plan)
				Expect(planModel).ToNot(BeNil())
				planModel.ID = core.StringPtr("testString")
				planModel.Label = core.StringPtr("testString")
				planModel.Name = core.StringPtr("testString")
				planModel.ShortDescription = core.StringPtr("testString")
				planModel.LongDescription = core.StringPtr("testString")
				planModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}
				planModel.Tags = []string{"testString"}
				planModel.AdditionalFeatures = []catalogmanagementv1.Feature{*featureModel}
				planModel.Created = CreateMockDateTime()
				planModel.Updated = CreateMockDateTime()
				planModel.Deployments = []catalogmanagementv1.Deployment{*deploymentModel}
				Expect(planModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(planModel.Label).To(Equal(core.StringPtr("testString")))
				Expect(planModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(planModel.ShortDescription).To(Equal(core.StringPtr("testString")))
				Expect(planModel.LongDescription).To(Equal(core.StringPtr("testString")))
				Expect(planModel.Metadata).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(planModel.Tags).To(Equal([]string{"testString"}))
				Expect(planModel.AdditionalFeatures).To(Equal([]catalogmanagementv1.Feature{*featureModel}))
				Expect(planModel.Created).To(Equal(CreateMockDateTime()))
				Expect(planModel.Updated).To(Equal(CreateMockDateTime()))
				Expect(planModel.Deployments).To(Equal([]catalogmanagementv1.Deployment{*deploymentModel}))

				// Construct an instance of the Kind model
				kindModel := new(catalogmanagementv1.Kind)
				Expect(kindModel).ToNot(BeNil())
				kindModel.ID = core.StringPtr("testString")
				kindModel.FormatKind = core.StringPtr("testString")
				kindModel.TargetKind = core.StringPtr("testString")
				kindModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}
				kindModel.InstallDescription = core.StringPtr("testString")
				kindModel.Tags = []string{"testString"}
				kindModel.AdditionalFeatures = []catalogmanagementv1.Feature{*featureModel}
				kindModel.Created = CreateMockDateTime()
				kindModel.Updated = CreateMockDateTime()
				kindModel.Versions = []catalogmanagementv1.Version{*versionModel}
				kindModel.Plans = []catalogmanagementv1.Plan{*planModel}
				Expect(kindModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(kindModel.FormatKind).To(Equal(core.StringPtr("testString")))
				Expect(kindModel.TargetKind).To(Equal(core.StringPtr("testString")))
				Expect(kindModel.Metadata).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(kindModel.InstallDescription).To(Equal(core.StringPtr("testString")))
				Expect(kindModel.Tags).To(Equal([]string{"testString"}))
				Expect(kindModel.AdditionalFeatures).To(Equal([]catalogmanagementv1.Feature{*featureModel}))
				Expect(kindModel.Created).To(Equal(CreateMockDateTime()))
				Expect(kindModel.Updated).To(Equal(CreateMockDateTime()))
				Expect(kindModel.Versions).To(Equal([]catalogmanagementv1.Version{*versionModel}))
				Expect(kindModel.Plans).To(Equal([]catalogmanagementv1.Plan{*planModel}))

				// Construct an instance of the RepoInfo model
				repoInfoModel := new(catalogmanagementv1.RepoInfo)
				Expect(repoInfoModel).ToNot(BeNil())
				repoInfoModel.Token = core.StringPtr("testString")
				repoInfoModel.Type = core.StringPtr("testString")
				Expect(repoInfoModel.Token).To(Equal(core.StringPtr("testString")))
				Expect(repoInfoModel.Type).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CreateOfferingOptions model
				catalogIdentifier := "testString"
				createOfferingOptionsModel := catalogManagementService.NewCreateOfferingOptions(catalogIdentifier)
				createOfferingOptionsModel.SetCatalogIdentifier("testString")
				createOfferingOptionsModel.SetID("testString")
				createOfferingOptionsModel.SetRev("testString")
				createOfferingOptionsModel.SetURL("testString")
				createOfferingOptionsModel.SetCrn("testString")
				createOfferingOptionsModel.SetLabel("testString")
				createOfferingOptionsModel.SetName("testString")
				createOfferingOptionsModel.SetOfferingIconURL("testString")
				createOfferingOptionsModel.SetOfferingDocsURL("testString")
				createOfferingOptionsModel.SetOfferingSupportURL("testString")
				createOfferingOptionsModel.SetTags([]string{"testString"})
				createOfferingOptionsModel.SetRating(ratingModel)
				createOfferingOptionsModel.SetCreated(CreateMockDateTime())
				createOfferingOptionsModel.SetUpdated(CreateMockDateTime())
				createOfferingOptionsModel.SetShortDescription("testString")
				createOfferingOptionsModel.SetLongDescription("testString")
				createOfferingOptionsModel.SetFeatures([]catalogmanagementv1.Feature{*featureModel})
				createOfferingOptionsModel.SetKinds([]catalogmanagementv1.Kind{*kindModel})
				createOfferingOptionsModel.SetPermitRequestIbmPublicPublish(true)
				createOfferingOptionsModel.SetIbmPublishApproved(true)
				createOfferingOptionsModel.SetPublicPublishApproved(true)
				createOfferingOptionsModel.SetPublicOriginalCrn("testString")
				createOfferingOptionsModel.SetPublishPublicCrn("testString")
				createOfferingOptionsModel.SetPortalApprovalRecord("testString")
				createOfferingOptionsModel.SetPortalUiURL("testString")
				createOfferingOptionsModel.SetCatalogID("testString")
				createOfferingOptionsModel.SetCatalogName("testString")
				createOfferingOptionsModel.SetMetadata(map[string]interface{}{"anyKey": "anyValue"})
				createOfferingOptionsModel.SetDisclaimer("testString")
				createOfferingOptionsModel.SetHidden(true)
				createOfferingOptionsModel.SetProvider("testString")
				createOfferingOptionsModel.SetRepoInfo(repoInfoModel)
				createOfferingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createOfferingOptionsModel).ToNot(BeNil())
				Expect(createOfferingOptionsModel.CatalogIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(createOfferingOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(createOfferingOptionsModel.Rev).To(Equal(core.StringPtr("testString")))
				Expect(createOfferingOptionsModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(createOfferingOptionsModel.Crn).To(Equal(core.StringPtr("testString")))
				Expect(createOfferingOptionsModel.Label).To(Equal(core.StringPtr("testString")))
				Expect(createOfferingOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createOfferingOptionsModel.OfferingIconURL).To(Equal(core.StringPtr("testString")))
				Expect(createOfferingOptionsModel.OfferingDocsURL).To(Equal(core.StringPtr("testString")))
				Expect(createOfferingOptionsModel.OfferingSupportURL).To(Equal(core.StringPtr("testString")))
				Expect(createOfferingOptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(createOfferingOptionsModel.Rating).To(Equal(ratingModel))
				Expect(createOfferingOptionsModel.Created).To(Equal(CreateMockDateTime()))
				Expect(createOfferingOptionsModel.Updated).To(Equal(CreateMockDateTime()))
				Expect(createOfferingOptionsModel.ShortDescription).To(Equal(core.StringPtr("testString")))
				Expect(createOfferingOptionsModel.LongDescription).To(Equal(core.StringPtr("testString")))
				Expect(createOfferingOptionsModel.Features).To(Equal([]catalogmanagementv1.Feature{*featureModel}))
				Expect(createOfferingOptionsModel.Kinds).To(Equal([]catalogmanagementv1.Kind{*kindModel}))
				Expect(createOfferingOptionsModel.PermitRequestIbmPublicPublish).To(Equal(core.BoolPtr(true)))
				Expect(createOfferingOptionsModel.IbmPublishApproved).To(Equal(core.BoolPtr(true)))
				Expect(createOfferingOptionsModel.PublicPublishApproved).To(Equal(core.BoolPtr(true)))
				Expect(createOfferingOptionsModel.PublicOriginalCrn).To(Equal(core.StringPtr("testString")))
				Expect(createOfferingOptionsModel.PublishPublicCrn).To(Equal(core.StringPtr("testString")))
				Expect(createOfferingOptionsModel.PortalApprovalRecord).To(Equal(core.StringPtr("testString")))
				Expect(createOfferingOptionsModel.PortalUiURL).To(Equal(core.StringPtr("testString")))
				Expect(createOfferingOptionsModel.CatalogID).To(Equal(core.StringPtr("testString")))
				Expect(createOfferingOptionsModel.CatalogName).To(Equal(core.StringPtr("testString")))
				Expect(createOfferingOptionsModel.Metadata).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(createOfferingOptionsModel.Disclaimer).To(Equal(core.StringPtr("testString")))
				Expect(createOfferingOptionsModel.Hidden).To(Equal(core.BoolPtr(true)))
				Expect(createOfferingOptionsModel.Provider).To(Equal(core.StringPtr("testString")))
				Expect(createOfferingOptionsModel.RepoInfo).To(Equal(repoInfoModel))
				Expect(createOfferingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateOperatorOptions successfully`, func() {
				// Construct an instance of the CreateOperatorOptions model
				xAuthRefreshToken := "testString"
				createOperatorOptionsModel := catalogManagementService.NewCreateOperatorOptions(xAuthRefreshToken)
				createOperatorOptionsModel.SetXAuthRefreshToken("testString")
				createOperatorOptionsModel.SetClusterID("testString")
				createOperatorOptionsModel.SetRegion("testString")
				createOperatorOptionsModel.SetNamespaces([]string{"testString"})
				createOperatorOptionsModel.SetAllNamespaces(true)
				createOperatorOptionsModel.SetVersionLocatorID("testString")
				createOperatorOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createOperatorOptionsModel).ToNot(BeNil())
				Expect(createOperatorOptionsModel.XAuthRefreshToken).To(Equal(core.StringPtr("testString")))
				Expect(createOperatorOptionsModel.ClusterID).To(Equal(core.StringPtr("testString")))
				Expect(createOperatorOptionsModel.Region).To(Equal(core.StringPtr("testString")))
				Expect(createOperatorOptionsModel.Namespaces).To(Equal([]string{"testString"}))
				Expect(createOperatorOptionsModel.AllNamespaces).To(Equal(core.BoolPtr(true)))
				Expect(createOperatorOptionsModel.VersionLocatorID).To(Equal(core.StringPtr("testString")))
				Expect(createOperatorOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteCatalogOptions successfully`, func() {
				// Construct an instance of the DeleteCatalogOptions model
				catalogIdentifier := "testString"
				deleteCatalogOptionsModel := catalogManagementService.NewDeleteCatalogOptions(catalogIdentifier)
				deleteCatalogOptionsModel.SetCatalogIdentifier("testString")
				deleteCatalogOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteCatalogOptionsModel).ToNot(BeNil())
				Expect(deleteCatalogOptionsModel.CatalogIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(deleteCatalogOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteLicenseEntitlementOptions successfully`, func() {
				// Construct an instance of the DeleteLicenseEntitlementOptions model
				entitlementID := "testString"
				deleteLicenseEntitlementOptionsModel := catalogManagementService.NewDeleteLicenseEntitlementOptions(entitlementID)
				deleteLicenseEntitlementOptionsModel.SetEntitlementID("testString")
				deleteLicenseEntitlementOptionsModel.SetAccountID("testString")
				deleteLicenseEntitlementOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteLicenseEntitlementOptionsModel).ToNot(BeNil())
				Expect(deleteLicenseEntitlementOptionsModel.EntitlementID).To(Equal(core.StringPtr("testString")))
				Expect(deleteLicenseEntitlementOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(deleteLicenseEntitlementOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteObjectOptions successfully`, func() {
				// Construct an instance of the DeleteObjectOptions model
				catalogIdentifier := "testString"
				objectIdentifier := "testString"
				deleteObjectOptionsModel := catalogManagementService.NewDeleteObjectOptions(catalogIdentifier, objectIdentifier)
				deleteObjectOptionsModel.SetCatalogIdentifier("testString")
				deleteObjectOptionsModel.SetObjectIdentifier("testString")
				deleteObjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteObjectOptionsModel).ToNot(BeNil())
				Expect(deleteObjectOptionsModel.CatalogIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(deleteObjectOptionsModel.ObjectIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(deleteObjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteOfferingOptions successfully`, func() {
				// Construct an instance of the DeleteOfferingOptions model
				catalogIdentifier := "testString"
				offeringID := "testString"
				deleteOfferingOptionsModel := catalogManagementService.NewDeleteOfferingOptions(catalogIdentifier, offeringID)
				deleteOfferingOptionsModel.SetCatalogIdentifier("testString")
				deleteOfferingOptionsModel.SetOfferingID("testString")
				deleteOfferingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteOfferingOptionsModel).ToNot(BeNil())
				Expect(deleteOfferingOptionsModel.CatalogIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(deleteOfferingOptionsModel.OfferingID).To(Equal(core.StringPtr("testString")))
				Expect(deleteOfferingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteOperatorOptions successfully`, func() {
				// Construct an instance of the DeleteOperatorOptions model
				xAuthRefreshToken := "testString"
				clusterID := "testString"
				region := "testString"
				versionLocatorID := "testString"
				deleteOperatorOptionsModel := catalogManagementService.NewDeleteOperatorOptions(xAuthRefreshToken, clusterID, region, versionLocatorID)
				deleteOperatorOptionsModel.SetXAuthRefreshToken("testString")
				deleteOperatorOptionsModel.SetClusterID("testString")
				deleteOperatorOptionsModel.SetRegion("testString")
				deleteOperatorOptionsModel.SetVersionLocatorID("testString")
				deleteOperatorOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteOperatorOptionsModel).ToNot(BeNil())
				Expect(deleteOperatorOptionsModel.XAuthRefreshToken).To(Equal(core.StringPtr("testString")))
				Expect(deleteOperatorOptionsModel.ClusterID).To(Equal(core.StringPtr("testString")))
				Expect(deleteOperatorOptionsModel.Region).To(Equal(core.StringPtr("testString")))
				Expect(deleteOperatorOptionsModel.VersionLocatorID).To(Equal(core.StringPtr("testString")))
				Expect(deleteOperatorOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteVersionOptions successfully`, func() {
				// Construct an instance of the DeleteVersionOptions model
				versionLocID := "testString"
				deleteVersionOptionsModel := catalogManagementService.NewDeleteVersionOptions(versionLocID)
				deleteVersionOptionsModel.SetVersionLocID("testString")
				deleteVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteVersionOptionsModel).ToNot(BeNil())
				Expect(deleteVersionOptionsModel.VersionLocID).To(Equal(core.StringPtr("testString")))
				Expect(deleteVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeprecateVersionOptions successfully`, func() {
				// Construct an instance of the DeprecateVersionOptions model
				versionLocID := "testString"
				deprecateVersionOptionsModel := catalogManagementService.NewDeprecateVersionOptions(versionLocID)
				deprecateVersionOptionsModel.SetVersionLocID("testString")
				deprecateVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deprecateVersionOptionsModel).ToNot(BeNil())
				Expect(deprecateVersionOptionsModel.VersionLocID).To(Equal(core.StringPtr("testString")))
				Expect(deprecateVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCatalogAccountAuditOptions successfully`, func() {
				// Construct an instance of the GetCatalogAccountAuditOptions model
				getCatalogAccountAuditOptionsModel := catalogManagementService.NewGetCatalogAccountAuditOptions()
				getCatalogAccountAuditOptionsModel.SetID("testString")
				getCatalogAccountAuditOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCatalogAccountAuditOptionsModel).ToNot(BeNil())
				Expect(getCatalogAccountAuditOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getCatalogAccountAuditOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCatalogAccountFiltersOptions successfully`, func() {
				// Construct an instance of the GetCatalogAccountFiltersOptions model
				getCatalogAccountFiltersOptionsModel := catalogManagementService.NewGetCatalogAccountFiltersOptions()
				getCatalogAccountFiltersOptionsModel.SetCatalog("testString")
				getCatalogAccountFiltersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCatalogAccountFiltersOptionsModel).ToNot(BeNil())
				Expect(getCatalogAccountFiltersOptionsModel.Catalog).To(Equal(core.StringPtr("testString")))
				Expect(getCatalogAccountFiltersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCatalogAccountOptions successfully`, func() {
				// Construct an instance of the GetCatalogAccountOptions model
				getCatalogAccountOptionsModel := catalogManagementService.NewGetCatalogAccountOptions()
				getCatalogAccountOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCatalogAccountOptionsModel).ToNot(BeNil())
				Expect(getCatalogAccountOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCatalogAuditOptions successfully`, func() {
				// Construct an instance of the GetCatalogAuditOptions model
				catalogIdentifier := "testString"
				getCatalogAuditOptionsModel := catalogManagementService.NewGetCatalogAuditOptions(catalogIdentifier)
				getCatalogAuditOptionsModel.SetCatalogIdentifier("testString")
				getCatalogAuditOptionsModel.SetID("testString")
				getCatalogAuditOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCatalogAuditOptionsModel).ToNot(BeNil())
				Expect(getCatalogAuditOptionsModel.CatalogIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(getCatalogAuditOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getCatalogAuditOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCatalogOptions successfully`, func() {
				// Construct an instance of the GetCatalogOptions model
				catalogIdentifier := "testString"
				getCatalogOptionsModel := catalogManagementService.NewGetCatalogOptions(catalogIdentifier)
				getCatalogOptionsModel.SetCatalogIdentifier("testString")
				getCatalogOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCatalogOptionsModel).ToNot(BeNil())
				Expect(getCatalogOptionsModel.CatalogIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(getCatalogOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetClusterOptions successfully`, func() {
				// Construct an instance of the GetClusterOptions model
				clusterID := "testString"
				region := "testString"
				xAuthRefreshToken := "testString"
				getClusterOptionsModel := catalogManagementService.NewGetClusterOptions(clusterID, region, xAuthRefreshToken)
				getClusterOptionsModel.SetClusterID("testString")
				getClusterOptionsModel.SetRegion("testString")
				getClusterOptionsModel.SetXAuthRefreshToken("testString")
				getClusterOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getClusterOptionsModel).ToNot(BeNil())
				Expect(getClusterOptionsModel.ClusterID).To(Equal(core.StringPtr("testString")))
				Expect(getClusterOptionsModel.Region).To(Equal(core.StringPtr("testString")))
				Expect(getClusterOptionsModel.XAuthRefreshToken).To(Equal(core.StringPtr("testString")))
				Expect(getClusterOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetConsumptionOfferingsOptions successfully`, func() {
				// Construct an instance of the GetConsumptionOfferingsOptions model
				getConsumptionOfferingsOptionsModel := catalogManagementService.NewGetConsumptionOfferingsOptions()
				getConsumptionOfferingsOptionsModel.SetDigest(true)
				getConsumptionOfferingsOptionsModel.SetCatalog("testString")
				getConsumptionOfferingsOptionsModel.SetSelect("all")
				getConsumptionOfferingsOptionsModel.SetIncludeHidden(true)
				getConsumptionOfferingsOptionsModel.SetLimit(int64(38))
				getConsumptionOfferingsOptionsModel.SetOffset(int64(38))
				getConsumptionOfferingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getConsumptionOfferingsOptionsModel).ToNot(BeNil())
				Expect(getConsumptionOfferingsOptionsModel.Digest).To(Equal(core.BoolPtr(true)))
				Expect(getConsumptionOfferingsOptionsModel.Catalog).To(Equal(core.StringPtr("testString")))
				Expect(getConsumptionOfferingsOptionsModel.Select).To(Equal(core.StringPtr("all")))
				Expect(getConsumptionOfferingsOptionsModel.IncludeHidden).To(Equal(core.BoolPtr(true)))
				Expect(getConsumptionOfferingsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(38))))
				Expect(getConsumptionOfferingsOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(38))))
				Expect(getConsumptionOfferingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetEnterpriseOptions successfully`, func() {
				// Construct an instance of the GetEnterpriseOptions model
				enterpriseID := "testString"
				getEnterpriseOptionsModel := catalogManagementService.NewGetEnterpriseOptions(enterpriseID)
				getEnterpriseOptionsModel.SetEnterpriseID("testString")
				getEnterpriseOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getEnterpriseOptionsModel).ToNot(BeNil())
				Expect(getEnterpriseOptionsModel.EnterpriseID).To(Equal(core.StringPtr("testString")))
				Expect(getEnterpriseOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetEnterprisesAuditOptions successfully`, func() {
				// Construct an instance of the GetEnterprisesAuditOptions model
				enterpriseID := "testString"
				getEnterprisesAuditOptionsModel := catalogManagementService.NewGetEnterprisesAuditOptions(enterpriseID)
				getEnterprisesAuditOptionsModel.SetEnterpriseID("testString")
				getEnterprisesAuditOptionsModel.SetID("testString")
				getEnterprisesAuditOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getEnterprisesAuditOptionsModel).ToNot(BeNil())
				Expect(getEnterprisesAuditOptionsModel.EnterpriseID).To(Equal(core.StringPtr("testString")))
				Expect(getEnterprisesAuditOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getEnterprisesAuditOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetLicenseEntitlementsOptions successfully`, func() {
				// Construct an instance of the GetLicenseEntitlementsOptions model
				licenseProductID := "testString"
				getLicenseEntitlementsOptionsModel := catalogManagementService.NewGetLicenseEntitlementsOptions(licenseProductID)
				getLicenseEntitlementsOptionsModel.SetLicenseProductID("testString")
				getLicenseEntitlementsOptionsModel.SetAccountID("testString")
				getLicenseEntitlementsOptionsModel.SetVersionID("testString")
				getLicenseEntitlementsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getLicenseEntitlementsOptionsModel).ToNot(BeNil())
				Expect(getLicenseEntitlementsOptionsModel.LicenseProductID).To(Equal(core.StringPtr("testString")))
				Expect(getLicenseEntitlementsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(getLicenseEntitlementsOptionsModel.VersionID).To(Equal(core.StringPtr("testString")))
				Expect(getLicenseEntitlementsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetLicenseProvidersOptions successfully`, func() {
				// Construct an instance of the GetLicenseProvidersOptions model
				getLicenseProvidersOptionsModel := catalogManagementService.NewGetLicenseProvidersOptions()
				getLicenseProvidersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getLicenseProvidersOptionsModel).ToNot(BeNil())
				Expect(getLicenseProvidersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetLicensesOptions successfully`, func() {
				// Construct an instance of the GetLicensesOptions model
				licenseProviderID := "testString"
				getLicensesOptionsModel := catalogManagementService.NewGetLicensesOptions(licenseProviderID)
				getLicensesOptionsModel.SetLicenseProviderID("testString")
				getLicensesOptionsModel.SetAccountID("testString")
				getLicensesOptionsModel.SetName("testString")
				getLicensesOptionsModel.SetLicenseType("testString")
				getLicensesOptionsModel.SetLicenseProductID("testString")
				getLicensesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getLicensesOptionsModel).ToNot(BeNil())
				Expect(getLicensesOptionsModel.LicenseProviderID).To(Equal(core.StringPtr("testString")))
				Expect(getLicensesOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(getLicensesOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(getLicensesOptionsModel.LicenseType).To(Equal(core.StringPtr("testString")))
				Expect(getLicensesOptionsModel.LicenseProductID).To(Equal(core.StringPtr("testString")))
				Expect(getLicensesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetNamespacesOptions successfully`, func() {
				// Construct an instance of the GetNamespacesOptions model
				clusterID := "testString"
				region := "testString"
				xAuthRefreshToken := "testString"
				getNamespacesOptionsModel := catalogManagementService.NewGetNamespacesOptions(clusterID, region, xAuthRefreshToken)
				getNamespacesOptionsModel.SetClusterID("testString")
				getNamespacesOptionsModel.SetRegion("testString")
				getNamespacesOptionsModel.SetXAuthRefreshToken("testString")
				getNamespacesOptionsModel.SetLimit(int64(38))
				getNamespacesOptionsModel.SetOffset(int64(38))
				getNamespacesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getNamespacesOptionsModel).ToNot(BeNil())
				Expect(getNamespacesOptionsModel.ClusterID).To(Equal(core.StringPtr("testString")))
				Expect(getNamespacesOptionsModel.Region).To(Equal(core.StringPtr("testString")))
				Expect(getNamespacesOptionsModel.XAuthRefreshToken).To(Equal(core.StringPtr("testString")))
				Expect(getNamespacesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(38))))
				Expect(getNamespacesOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(38))))
				Expect(getNamespacesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetObjectAuditOptions successfully`, func() {
				// Construct an instance of the GetObjectAuditOptions model
				catalogIdentifier := "testString"
				objectIdentifier := "testString"
				getObjectAuditOptionsModel := catalogManagementService.NewGetObjectAuditOptions(catalogIdentifier, objectIdentifier)
				getObjectAuditOptionsModel.SetCatalogIdentifier("testString")
				getObjectAuditOptionsModel.SetObjectIdentifier("testString")
				getObjectAuditOptionsModel.SetID("testString")
				getObjectAuditOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getObjectAuditOptionsModel).ToNot(BeNil())
				Expect(getObjectAuditOptionsModel.CatalogIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(getObjectAuditOptionsModel.ObjectIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(getObjectAuditOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getObjectAuditOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetObjectOptions successfully`, func() {
				// Construct an instance of the GetObjectOptions model
				catalogIdentifier := "testString"
				objectIdentifier := "testString"
				getObjectOptionsModel := catalogManagementService.NewGetObjectOptions(catalogIdentifier, objectIdentifier)
				getObjectOptionsModel.SetCatalogIdentifier("testString")
				getObjectOptionsModel.SetObjectIdentifier("testString")
				getObjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getObjectOptionsModel).ToNot(BeNil())
				Expect(getObjectOptionsModel.CatalogIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(getObjectOptionsModel.ObjectIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(getObjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetOfferingAuditOptions successfully`, func() {
				// Construct an instance of the GetOfferingAuditOptions model
				catalogIdentifier := "testString"
				offeringID := "testString"
				getOfferingAuditOptionsModel := catalogManagementService.NewGetOfferingAuditOptions(catalogIdentifier, offeringID)
				getOfferingAuditOptionsModel.SetCatalogIdentifier("testString")
				getOfferingAuditOptionsModel.SetOfferingID("testString")
				getOfferingAuditOptionsModel.SetID("testString")
				getOfferingAuditOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getOfferingAuditOptionsModel).ToNot(BeNil())
				Expect(getOfferingAuditOptionsModel.CatalogIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(getOfferingAuditOptionsModel.OfferingID).To(Equal(core.StringPtr("testString")))
				Expect(getOfferingAuditOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getOfferingAuditOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetOfferingOptions successfully`, func() {
				// Construct an instance of the GetOfferingOptions model
				catalogIdentifier := "testString"
				offeringID := "testString"
				getOfferingOptionsModel := catalogManagementService.NewGetOfferingOptions(catalogIdentifier, offeringID)
				getOfferingOptionsModel.SetCatalogIdentifier("testString")
				getOfferingOptionsModel.SetOfferingID("testString")
				getOfferingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getOfferingOptionsModel).ToNot(BeNil())
				Expect(getOfferingOptionsModel.CatalogIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(getOfferingOptionsModel.OfferingID).To(Equal(core.StringPtr("testString")))
				Expect(getOfferingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetOverrideValuesOptions successfully`, func() {
				// Construct an instance of the GetOverrideValuesOptions model
				versionLocID := "testString"
				getOverrideValuesOptionsModel := catalogManagementService.NewGetOverrideValuesOptions(versionLocID)
				getOverrideValuesOptionsModel.SetVersionLocID("testString")
				getOverrideValuesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getOverrideValuesOptionsModel).ToNot(BeNil())
				Expect(getOverrideValuesOptionsModel.VersionLocID).To(Equal(core.StringPtr("testString")))
				Expect(getOverrideValuesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetPreinstallOptions successfully`, func() {
				// Construct an instance of the GetPreinstallOptions model
				versionLocID := "testString"
				xAuthRefreshToken := "testString"
				getPreinstallOptionsModel := catalogManagementService.NewGetPreinstallOptions(versionLocID, xAuthRefreshToken)
				getPreinstallOptionsModel.SetVersionLocID("testString")
				getPreinstallOptionsModel.SetXAuthRefreshToken("testString")
				getPreinstallOptionsModel.SetClusterID("testString")
				getPreinstallOptionsModel.SetRegion("testString")
				getPreinstallOptionsModel.SetNamespace("testString")
				getPreinstallOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getPreinstallOptionsModel).ToNot(BeNil())
				Expect(getPreinstallOptionsModel.VersionLocID).To(Equal(core.StringPtr("testString")))
				Expect(getPreinstallOptionsModel.XAuthRefreshToken).To(Equal(core.StringPtr("testString")))
				Expect(getPreinstallOptionsModel.ClusterID).To(Equal(core.StringPtr("testString")))
				Expect(getPreinstallOptionsModel.Region).To(Equal(core.StringPtr("testString")))
				Expect(getPreinstallOptionsModel.Namespace).To(Equal(core.StringPtr("testString")))
				Expect(getPreinstallOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetRepoOptions successfully`, func() {
				// Construct an instance of the GetRepoOptions model
				typeVar := "testString"
				charturl := "testString"
				getRepoOptionsModel := catalogManagementService.NewGetRepoOptions(typeVar, charturl)
				getRepoOptionsModel.SetType("testString")
				getRepoOptionsModel.SetCharturl("testString")
				getRepoOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getRepoOptionsModel).ToNot(BeNil())
				Expect(getRepoOptionsModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(getRepoOptionsModel.Charturl).To(Equal(core.StringPtr("testString")))
				Expect(getRepoOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetReposOptions successfully`, func() {
				// Construct an instance of the GetReposOptions model
				typeVar := "testString"
				repourl := "testString"
				getReposOptionsModel := catalogManagementService.NewGetReposOptions(typeVar, repourl)
				getReposOptionsModel.SetType("testString")
				getReposOptionsModel.SetRepourl("testString")
				getReposOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getReposOptionsModel).ToNot(BeNil())
				Expect(getReposOptionsModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(getReposOptionsModel.Repourl).To(Equal(core.StringPtr("testString")))
				Expect(getReposOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetResourceGroupsOptions successfully`, func() {
				// Construct an instance of the GetResourceGroupsOptions model
				getResourceGroupsOptionsModel := catalogManagementService.NewGetResourceGroupsOptions()
				getResourceGroupsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getResourceGroupsOptionsModel).ToNot(BeNil())
				Expect(getResourceGroupsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSchematicsWorkspacesOptions successfully`, func() {
				// Construct an instance of the GetSchematicsWorkspacesOptions model
				versionLocID := "testString"
				xAuthRefreshToken := "testString"
				getSchematicsWorkspacesOptionsModel := catalogManagementService.NewGetSchematicsWorkspacesOptions(versionLocID, xAuthRefreshToken)
				getSchematicsWorkspacesOptionsModel.SetVersionLocID("testString")
				getSchematicsWorkspacesOptionsModel.SetXAuthRefreshToken("testString")
				getSchematicsWorkspacesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSchematicsWorkspacesOptionsModel).ToNot(BeNil())
				Expect(getSchematicsWorkspacesOptionsModel.VersionLocID).To(Equal(core.StringPtr("testString")))
				Expect(getSchematicsWorkspacesOptionsModel.XAuthRefreshToken).To(Equal(core.StringPtr("testString")))
				Expect(getSchematicsWorkspacesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetValidationStatusOptions successfully`, func() {
				// Construct an instance of the GetValidationStatusOptions model
				versionLocID := "testString"
				xAuthRefreshToken := "testString"
				getValidationStatusOptionsModel := catalogManagementService.NewGetValidationStatusOptions(versionLocID, xAuthRefreshToken)
				getValidationStatusOptionsModel.SetVersionLocID("testString")
				getValidationStatusOptionsModel.SetXAuthRefreshToken("testString")
				getValidationStatusOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getValidationStatusOptionsModel).ToNot(BeNil())
				Expect(getValidationStatusOptionsModel.VersionLocID).To(Equal(core.StringPtr("testString")))
				Expect(getValidationStatusOptionsModel.XAuthRefreshToken).To(Equal(core.StringPtr("testString")))
				Expect(getValidationStatusOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetVersionAboutOptions successfully`, func() {
				// Construct an instance of the GetVersionAboutOptions model
				versionLocID := "testString"
				getVersionAboutOptionsModel := catalogManagementService.NewGetVersionAboutOptions(versionLocID)
				getVersionAboutOptionsModel.SetVersionLocID("testString")
				getVersionAboutOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getVersionAboutOptionsModel).ToNot(BeNil())
				Expect(getVersionAboutOptionsModel.VersionLocID).To(Equal(core.StringPtr("testString")))
				Expect(getVersionAboutOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetVersionContainerImagesOptions successfully`, func() {
				// Construct an instance of the GetVersionContainerImagesOptions model
				versionLocID := "testString"
				getVersionContainerImagesOptionsModel := catalogManagementService.NewGetVersionContainerImagesOptions(versionLocID)
				getVersionContainerImagesOptionsModel.SetVersionLocID("testString")
				getVersionContainerImagesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getVersionContainerImagesOptionsModel).ToNot(BeNil())
				Expect(getVersionContainerImagesOptionsModel.VersionLocID).To(Equal(core.StringPtr("testString")))
				Expect(getVersionContainerImagesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetVersionLicenseOptions successfully`, func() {
				// Construct an instance of the GetVersionLicenseOptions model
				versionLocID := "testString"
				licenseID := "testString"
				getVersionLicenseOptionsModel := catalogManagementService.NewGetVersionLicenseOptions(versionLocID, licenseID)
				getVersionLicenseOptionsModel.SetVersionLocID("testString")
				getVersionLicenseOptionsModel.SetLicenseID("testString")
				getVersionLicenseOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getVersionLicenseOptionsModel).ToNot(BeNil())
				Expect(getVersionLicenseOptionsModel.VersionLocID).To(Equal(core.StringPtr("testString")))
				Expect(getVersionLicenseOptionsModel.LicenseID).To(Equal(core.StringPtr("testString")))
				Expect(getVersionLicenseOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetVersionOptions successfully`, func() {
				// Construct an instance of the GetVersionOptions model
				versionLocID := "testString"
				getVersionOptionsModel := catalogManagementService.NewGetVersionOptions(versionLocID)
				getVersionOptionsModel.SetVersionLocID("testString")
				getVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getVersionOptionsModel).ToNot(BeNil())
				Expect(getVersionOptionsModel.VersionLocID).To(Equal(core.StringPtr("testString")))
				Expect(getVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetVersionUpdatesOptions successfully`, func() {
				// Construct an instance of the GetVersionUpdatesOptions model
				versionLocID := "testString"
				getVersionUpdatesOptionsModel := catalogManagementService.NewGetVersionUpdatesOptions(versionLocID)
				getVersionUpdatesOptionsModel.SetVersionLocID("testString")
				getVersionUpdatesOptionsModel.SetClusterID("testString")
				getVersionUpdatesOptionsModel.SetRegion("testString")
				getVersionUpdatesOptionsModel.SetResourceGroupID("testString")
				getVersionUpdatesOptionsModel.SetNamespace("testString")
				getVersionUpdatesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getVersionUpdatesOptionsModel).ToNot(BeNil())
				Expect(getVersionUpdatesOptionsModel.VersionLocID).To(Equal(core.StringPtr("testString")))
				Expect(getVersionUpdatesOptionsModel.ClusterID).To(Equal(core.StringPtr("testString")))
				Expect(getVersionUpdatesOptionsModel.Region).To(Equal(core.StringPtr("testString")))
				Expect(getVersionUpdatesOptionsModel.ResourceGroupID).To(Equal(core.StringPtr("testString")))
				Expect(getVersionUpdatesOptionsModel.Namespace).To(Equal(core.StringPtr("testString")))
				Expect(getVersionUpdatesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetVersionWorkingCopyOptions successfully`, func() {
				// Construct an instance of the GetVersionWorkingCopyOptions model
				versionLocID := "testString"
				getVersionWorkingCopyOptionsModel := catalogManagementService.NewGetVersionWorkingCopyOptions(versionLocID)
				getVersionWorkingCopyOptionsModel.SetVersionLocID("testString")
				getVersionWorkingCopyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getVersionWorkingCopyOptionsModel).ToNot(BeNil())
				Expect(getVersionWorkingCopyOptionsModel.VersionLocID).To(Equal(core.StringPtr("testString")))
				Expect(getVersionWorkingCopyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewIbmPublishVersionOptions successfully`, func() {
				// Construct an instance of the IbmPublishVersionOptions model
				versionLocID := "testString"
				ibmPublishVersionOptionsModel := catalogManagementService.NewIbmPublishVersionOptions(versionLocID)
				ibmPublishVersionOptionsModel.SetVersionLocID("testString")
				ibmPublishVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(ibmPublishVersionOptionsModel).ToNot(BeNil())
				Expect(ibmPublishVersionOptionsModel.VersionLocID).To(Equal(core.StringPtr("testString")))
				Expect(ibmPublishVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewImportOfferingOptions successfully`, func() {
				// Construct an instance of the ImportOfferingOptions model
				catalogIdentifier := "testString"
				importOfferingOptionsModel := catalogManagementService.NewImportOfferingOptions(catalogIdentifier)
				importOfferingOptionsModel.SetCatalogIdentifier("testString")
				importOfferingOptionsModel.SetTags([]string{"testString"})
				importOfferingOptionsModel.SetTargetKinds([]string{"testString"})
				importOfferingOptionsModel.SetContent([]int64{int64(38)})
				importOfferingOptionsModel.SetZipurl("testString")
				importOfferingOptionsModel.SetOfferingID("testString")
				importOfferingOptionsModel.SetTargetVersion("testString")
				importOfferingOptionsModel.SetIncludeConfig(true)
				importOfferingOptionsModel.SetRepoType("testString")
				importOfferingOptionsModel.SetXAuthToken("testString")
				importOfferingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(importOfferingOptionsModel).ToNot(BeNil())
				Expect(importOfferingOptionsModel.CatalogIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(importOfferingOptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(importOfferingOptionsModel.TargetKinds).To(Equal([]string{"testString"}))
				Expect(importOfferingOptionsModel.Content).To(Equal([]int64{int64(38)}))
				Expect(importOfferingOptionsModel.Zipurl).To(Equal(core.StringPtr("testString")))
				Expect(importOfferingOptionsModel.OfferingID).To(Equal(core.StringPtr("testString")))
				Expect(importOfferingOptionsModel.TargetVersion).To(Equal(core.StringPtr("testString")))
				Expect(importOfferingOptionsModel.IncludeConfig).To(Equal(core.BoolPtr(true)))
				Expect(importOfferingOptionsModel.RepoType).To(Equal(core.StringPtr("testString")))
				Expect(importOfferingOptionsModel.XAuthToken).To(Equal(core.StringPtr("testString")))
				Expect(importOfferingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewImportOfferingVersionOptions successfully`, func() {
				// Construct an instance of the ImportOfferingVersionOptions model
				catalogIdentifier := "testString"
				offeringID := "testString"
				importOfferingVersionOptionsModel := catalogManagementService.NewImportOfferingVersionOptions(catalogIdentifier, offeringID)
				importOfferingVersionOptionsModel.SetCatalogIdentifier("testString")
				importOfferingVersionOptionsModel.SetOfferingID("testString")
				importOfferingVersionOptionsModel.SetTags([]string{"testString"})
				importOfferingVersionOptionsModel.SetTargetKinds([]string{"testString"})
				importOfferingVersionOptionsModel.SetContent([]int64{int64(38)})
				importOfferingVersionOptionsModel.SetZipurl("testString")
				importOfferingVersionOptionsModel.SetTargetVersion("testString")
				importOfferingVersionOptionsModel.SetIncludeConfig(true)
				importOfferingVersionOptionsModel.SetRepoType("testString")
				importOfferingVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(importOfferingVersionOptionsModel).ToNot(BeNil())
				Expect(importOfferingVersionOptionsModel.CatalogIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(importOfferingVersionOptionsModel.OfferingID).To(Equal(core.StringPtr("testString")))
				Expect(importOfferingVersionOptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(importOfferingVersionOptionsModel.TargetKinds).To(Equal([]string{"testString"}))
				Expect(importOfferingVersionOptionsModel.Content).To(Equal([]int64{int64(38)}))
				Expect(importOfferingVersionOptionsModel.Zipurl).To(Equal(core.StringPtr("testString")))
				Expect(importOfferingVersionOptionsModel.TargetVersion).To(Equal(core.StringPtr("testString")))
				Expect(importOfferingVersionOptionsModel.IncludeConfig).To(Equal(core.BoolPtr(true)))
				Expect(importOfferingVersionOptionsModel.RepoType).To(Equal(core.StringPtr("testString")))
				Expect(importOfferingVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewInstallVersionOptions successfully`, func() {
				// Construct an instance of the DeployRequestBodySchematics model
				deployRequestBodySchematicsModel := new(catalogmanagementv1.DeployRequestBodySchematics)
				Expect(deployRequestBodySchematicsModel).ToNot(BeNil())
				deployRequestBodySchematicsModel.Name = core.StringPtr("testString")
				deployRequestBodySchematicsModel.Description = core.StringPtr("testString")
				deployRequestBodySchematicsModel.Tags = []string{"testString"}
				deployRequestBodySchematicsModel.ResourceGroupID = core.StringPtr("testString")
				Expect(deployRequestBodySchematicsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(deployRequestBodySchematicsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(deployRequestBodySchematicsModel.Tags).To(Equal([]string{"testString"}))
				Expect(deployRequestBodySchematicsModel.ResourceGroupID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the InstallVersionOptions model
				versionLocID := "testString"
				xAuthRefreshToken := "testString"
				installVersionOptionsModel := catalogManagementService.NewInstallVersionOptions(versionLocID, xAuthRefreshToken)
				installVersionOptionsModel.SetVersionLocID("testString")
				installVersionOptionsModel.SetXAuthRefreshToken("testString")
				installVersionOptionsModel.SetClusterID("testString")
				installVersionOptionsModel.SetRegion("testString")
				installVersionOptionsModel.SetNamespace("testString")
				installVersionOptionsModel.SetOverrideValues(map[string]interface{}{"anyKey": "anyValue"})
				installVersionOptionsModel.SetEntitlementApikey("testString")
				installVersionOptionsModel.SetSchematics(deployRequestBodySchematicsModel)
				installVersionOptionsModel.SetScript("testString")
				installVersionOptionsModel.SetScriptID("testString")
				installVersionOptionsModel.SetVersionLocatorID("testString")
				installVersionOptionsModel.SetVcenterID("testString")
				installVersionOptionsModel.SetVcenterUser("testString")
				installVersionOptionsModel.SetVcenterPassword("testString")
				installVersionOptionsModel.SetVcenterLocation("testString")
				installVersionOptionsModel.SetVcenterDatastore("testString")
				installVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(installVersionOptionsModel).ToNot(BeNil())
				Expect(installVersionOptionsModel.VersionLocID).To(Equal(core.StringPtr("testString")))
				Expect(installVersionOptionsModel.XAuthRefreshToken).To(Equal(core.StringPtr("testString")))
				Expect(installVersionOptionsModel.ClusterID).To(Equal(core.StringPtr("testString")))
				Expect(installVersionOptionsModel.Region).To(Equal(core.StringPtr("testString")))
				Expect(installVersionOptionsModel.Namespace).To(Equal(core.StringPtr("testString")))
				Expect(installVersionOptionsModel.OverrideValues).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(installVersionOptionsModel.EntitlementApikey).To(Equal(core.StringPtr("testString")))
				Expect(installVersionOptionsModel.Schematics).To(Equal(deployRequestBodySchematicsModel))
				Expect(installVersionOptionsModel.Script).To(Equal(core.StringPtr("testString")))
				Expect(installVersionOptionsModel.ScriptID).To(Equal(core.StringPtr("testString")))
				Expect(installVersionOptionsModel.VersionLocatorID).To(Equal(core.StringPtr("testString")))
				Expect(installVersionOptionsModel.VcenterID).To(Equal(core.StringPtr("testString")))
				Expect(installVersionOptionsModel.VcenterUser).To(Equal(core.StringPtr("testString")))
				Expect(installVersionOptionsModel.VcenterPassword).To(Equal(core.StringPtr("testString")))
				Expect(installVersionOptionsModel.VcenterLocation).To(Equal(core.StringPtr("testString")))
				Expect(installVersionOptionsModel.VcenterDatastore).To(Equal(core.StringPtr("testString")))
				Expect(installVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListCatalogsOptions successfully`, func() {
				// Construct an instance of the ListCatalogsOptions model
				listCatalogsOptionsModel := catalogManagementService.NewListCatalogsOptions()
				listCatalogsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listCatalogsOptionsModel).ToNot(BeNil())
				Expect(listCatalogsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListClustersOptions successfully`, func() {
				// Construct an instance of the ListClustersOptions model
				listClustersOptionsModel := catalogManagementService.NewListClustersOptions()
				listClustersOptionsModel.SetLimit(int64(38))
				listClustersOptionsModel.SetOffset(int64(38))
				listClustersOptionsModel.SetType("testString")
				listClustersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listClustersOptionsModel).ToNot(BeNil())
				Expect(listClustersOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listClustersOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listClustersOptionsModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(listClustersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListLicenseEntitlementsOptions successfully`, func() {
				// Construct an instance of the ListLicenseEntitlementsOptions model
				listLicenseEntitlementsOptionsModel := catalogManagementService.NewListLicenseEntitlementsOptions()
				listLicenseEntitlementsOptionsModel.SetAccountID("testString")
				listLicenseEntitlementsOptionsModel.SetLicenseProductID("testString")
				listLicenseEntitlementsOptionsModel.SetVersionID("testString")
				listLicenseEntitlementsOptionsModel.SetState("testString")
				listLicenseEntitlementsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listLicenseEntitlementsOptionsModel).ToNot(BeNil())
				Expect(listLicenseEntitlementsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listLicenseEntitlementsOptionsModel.LicenseProductID).To(Equal(core.StringPtr("testString")))
				Expect(listLicenseEntitlementsOptionsModel.VersionID).To(Equal(core.StringPtr("testString")))
				Expect(listLicenseEntitlementsOptionsModel.State).To(Equal(core.StringPtr("testString")))
				Expect(listLicenseEntitlementsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListObjectsOptions successfully`, func() {
				// Construct an instance of the ListObjectsOptions model
				catalogIdentifier := "testString"
				listObjectsOptionsModel := catalogManagementService.NewListObjectsOptions(catalogIdentifier)
				listObjectsOptionsModel.SetCatalogIdentifier("testString")
				listObjectsOptionsModel.SetLimit(int64(38))
				listObjectsOptionsModel.SetOffset(int64(38))
				listObjectsOptionsModel.SetName("testString")
				listObjectsOptionsModel.SetSort("testString")
				listObjectsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listObjectsOptionsModel).ToNot(BeNil())
				Expect(listObjectsOptionsModel.CatalogIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(listObjectsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listObjectsOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listObjectsOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(listObjectsOptionsModel.Sort).To(Equal(core.StringPtr("testString")))
				Expect(listObjectsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListOfferingsOptions successfully`, func() {
				// Construct an instance of the ListOfferingsOptions model
				catalogIdentifier := "testString"
				listOfferingsOptionsModel := catalogManagementService.NewListOfferingsOptions(catalogIdentifier)
				listOfferingsOptionsModel.SetCatalogIdentifier("testString")
				listOfferingsOptionsModel.SetDigest(true)
				listOfferingsOptionsModel.SetLimit(int64(38))
				listOfferingsOptionsModel.SetOffset(int64(38))
				listOfferingsOptionsModel.SetName("testString")
				listOfferingsOptionsModel.SetSort("testString")
				listOfferingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listOfferingsOptionsModel).ToNot(BeNil())
				Expect(listOfferingsOptionsModel.CatalogIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(listOfferingsOptionsModel.Digest).To(Equal(core.BoolPtr(true)))
				Expect(listOfferingsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listOfferingsOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listOfferingsOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(listOfferingsOptionsModel.Sort).To(Equal(core.StringPtr("testString")))
				Expect(listOfferingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListOperatorsOptions successfully`, func() {
				// Construct an instance of the ListOperatorsOptions model
				xAuthRefreshToken := "testString"
				clusterID := "testString"
				region := "testString"
				versionLocatorID := "testString"
				listOperatorsOptionsModel := catalogManagementService.NewListOperatorsOptions(xAuthRefreshToken, clusterID, region, versionLocatorID)
				listOperatorsOptionsModel.SetXAuthRefreshToken("testString")
				listOperatorsOptionsModel.SetClusterID("testString")
				listOperatorsOptionsModel.SetRegion("testString")
				listOperatorsOptionsModel.SetVersionLocatorID("testString")
				listOperatorsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listOperatorsOptionsModel).ToNot(BeNil())
				Expect(listOperatorsOptionsModel.XAuthRefreshToken).To(Equal(core.StringPtr("testString")))
				Expect(listOperatorsOptionsModel.ClusterID).To(Equal(core.StringPtr("testString")))
				Expect(listOperatorsOptionsModel.Region).To(Equal(core.StringPtr("testString")))
				Expect(listOperatorsOptionsModel.VersionLocatorID).To(Equal(core.StringPtr("testString")))
				Expect(listOperatorsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListVersionsOptions successfully`, func() {
				// Construct an instance of the ListVersionsOptions model
				q := "testString"
				listVersionsOptionsModel := catalogManagementService.NewListVersionsOptions(q)
				listVersionsOptionsModel.SetQ("testString")
				listVersionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listVersionsOptionsModel).ToNot(BeNil())
				Expect(listVersionsOptionsModel.Q).To(Equal(core.StringPtr("testString")))
				Expect(listVersionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPreinstallVersionOptions successfully`, func() {
				// Construct an instance of the DeployRequestBodySchematics model
				deployRequestBodySchematicsModel := new(catalogmanagementv1.DeployRequestBodySchematics)
				Expect(deployRequestBodySchematicsModel).ToNot(BeNil())
				deployRequestBodySchematicsModel.Name = core.StringPtr("testString")
				deployRequestBodySchematicsModel.Description = core.StringPtr("testString")
				deployRequestBodySchematicsModel.Tags = []string{"testString"}
				deployRequestBodySchematicsModel.ResourceGroupID = core.StringPtr("testString")
				Expect(deployRequestBodySchematicsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(deployRequestBodySchematicsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(deployRequestBodySchematicsModel.Tags).To(Equal([]string{"testString"}))
				Expect(deployRequestBodySchematicsModel.ResourceGroupID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the PreinstallVersionOptions model
				versionLocID := "testString"
				xAuthRefreshToken := "testString"
				preinstallVersionOptionsModel := catalogManagementService.NewPreinstallVersionOptions(versionLocID, xAuthRefreshToken)
				preinstallVersionOptionsModel.SetVersionLocID("testString")
				preinstallVersionOptionsModel.SetXAuthRefreshToken("testString")
				preinstallVersionOptionsModel.SetClusterID("testString")
				preinstallVersionOptionsModel.SetRegion("testString")
				preinstallVersionOptionsModel.SetNamespace("testString")
				preinstallVersionOptionsModel.SetOverrideValues(map[string]interface{}{"anyKey": "anyValue"})
				preinstallVersionOptionsModel.SetEntitlementApikey("testString")
				preinstallVersionOptionsModel.SetSchematics(deployRequestBodySchematicsModel)
				preinstallVersionOptionsModel.SetScript("testString")
				preinstallVersionOptionsModel.SetScriptID("testString")
				preinstallVersionOptionsModel.SetVersionLocatorID("testString")
				preinstallVersionOptionsModel.SetVcenterID("testString")
				preinstallVersionOptionsModel.SetVcenterUser("testString")
				preinstallVersionOptionsModel.SetVcenterPassword("testString")
				preinstallVersionOptionsModel.SetVcenterLocation("testString")
				preinstallVersionOptionsModel.SetVcenterDatastore("testString")
				preinstallVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(preinstallVersionOptionsModel).ToNot(BeNil())
				Expect(preinstallVersionOptionsModel.VersionLocID).To(Equal(core.StringPtr("testString")))
				Expect(preinstallVersionOptionsModel.XAuthRefreshToken).To(Equal(core.StringPtr("testString")))
				Expect(preinstallVersionOptionsModel.ClusterID).To(Equal(core.StringPtr("testString")))
				Expect(preinstallVersionOptionsModel.Region).To(Equal(core.StringPtr("testString")))
				Expect(preinstallVersionOptionsModel.Namespace).To(Equal(core.StringPtr("testString")))
				Expect(preinstallVersionOptionsModel.OverrideValues).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(preinstallVersionOptionsModel.EntitlementApikey).To(Equal(core.StringPtr("testString")))
				Expect(preinstallVersionOptionsModel.Schematics).To(Equal(deployRequestBodySchematicsModel))
				Expect(preinstallVersionOptionsModel.Script).To(Equal(core.StringPtr("testString")))
				Expect(preinstallVersionOptionsModel.ScriptID).To(Equal(core.StringPtr("testString")))
				Expect(preinstallVersionOptionsModel.VersionLocatorID).To(Equal(core.StringPtr("testString")))
				Expect(preinstallVersionOptionsModel.VcenterID).To(Equal(core.StringPtr("testString")))
				Expect(preinstallVersionOptionsModel.VcenterUser).To(Equal(core.StringPtr("testString")))
				Expect(preinstallVersionOptionsModel.VcenterPassword).To(Equal(core.StringPtr("testString")))
				Expect(preinstallVersionOptionsModel.VcenterLocation).To(Equal(core.StringPtr("testString")))
				Expect(preinstallVersionOptionsModel.VcenterDatastore).To(Equal(core.StringPtr("testString")))
				Expect(preinstallVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPublicPublishVersionOptions successfully`, func() {
				// Construct an instance of the PublicPublishVersionOptions model
				versionLocID := "testString"
				publicPublishVersionOptionsModel := catalogManagementService.NewPublicPublishVersionOptions(versionLocID)
				publicPublishVersionOptionsModel.SetVersionLocID("testString")
				publicPublishVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(publicPublishVersionOptionsModel).ToNot(BeNil())
				Expect(publicPublishVersionOptionsModel.VersionLocID).To(Equal(core.StringPtr("testString")))
				Expect(publicPublishVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReloadOfferingOptions successfully`, func() {
				// Construct an instance of the ReloadOfferingOptions model
				catalogIdentifier := "testString"
				offeringID := "testString"
				targetVersion := "testString"
				reloadOfferingOptionsModel := catalogManagementService.NewReloadOfferingOptions(catalogIdentifier, offeringID, targetVersion)
				reloadOfferingOptionsModel.SetCatalogIdentifier("testString")
				reloadOfferingOptionsModel.SetOfferingID("testString")
				reloadOfferingOptionsModel.SetTargetVersion("testString")
				reloadOfferingOptionsModel.SetTags([]string{"testString"})
				reloadOfferingOptionsModel.SetTargetKinds([]string{"testString"})
				reloadOfferingOptionsModel.SetContent([]int64{int64(38)})
				reloadOfferingOptionsModel.SetZipurl("testString")
				reloadOfferingOptionsModel.SetRepoType("testString")
				reloadOfferingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(reloadOfferingOptionsModel).ToNot(BeNil())
				Expect(reloadOfferingOptionsModel.CatalogIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(reloadOfferingOptionsModel.OfferingID).To(Equal(core.StringPtr("testString")))
				Expect(reloadOfferingOptionsModel.TargetVersion).To(Equal(core.StringPtr("testString")))
				Expect(reloadOfferingOptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(reloadOfferingOptionsModel.TargetKinds).To(Equal([]string{"testString"}))
				Expect(reloadOfferingOptionsModel.Content).To(Equal([]int64{int64(38)}))
				Expect(reloadOfferingOptionsModel.Zipurl).To(Equal(core.StringPtr("testString")))
				Expect(reloadOfferingOptionsModel.RepoType).To(Equal(core.StringPtr("testString")))
				Expect(reloadOfferingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceCatalogOptions successfully`, func() {
				// Construct an instance of the Feature model
				featureModel := new(catalogmanagementv1.Feature)
				Expect(featureModel).ToNot(BeNil())
				featureModel.Title = core.StringPtr("testString")
				featureModel.Description = core.StringPtr("testString")
				Expect(featureModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(featureModel.Description).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the FilterTerms model
				filterTermsModel := new(catalogmanagementv1.FilterTerms)
				Expect(filterTermsModel).ToNot(BeNil())
				filterTermsModel.FilterTerms = []string{"testString"}
				Expect(filterTermsModel.FilterTerms).To(Equal([]string{"testString"}))

				// Construct an instance of the CategoryFilter model
				categoryFilterModel := new(catalogmanagementv1.CategoryFilter)
				Expect(categoryFilterModel).ToNot(BeNil())
				categoryFilterModel.Include = core.BoolPtr(true)
				categoryFilterModel.Filter = filterTermsModel
				Expect(categoryFilterModel.Include).To(Equal(core.BoolPtr(true)))
				Expect(categoryFilterModel.Filter).To(Equal(filterTermsModel))

				// Construct an instance of the IDFilter model
				idFilterModel := new(catalogmanagementv1.IDFilter)
				Expect(idFilterModel).ToNot(BeNil())
				idFilterModel.Include = filterTermsModel
				idFilterModel.Exclude = filterTermsModel
				Expect(idFilterModel.Include).To(Equal(filterTermsModel))
				Expect(idFilterModel.Exclude).To(Equal(filterTermsModel))

				// Construct an instance of the Filters model
				filtersModel := new(catalogmanagementv1.Filters)
				Expect(filtersModel).ToNot(BeNil())
				filtersModel.IncludeAll = core.BoolPtr(true)
				filtersModel.CategoryFilters = make(map[string]catalogmanagementv1.CategoryFilter)
				filtersModel.IdFilters = idFilterModel
				filtersModel.CategoryFilters["foo"] = *categoryFilterModel
				Expect(filtersModel.IncludeAll).To(Equal(core.BoolPtr(true)))
				Expect(filtersModel.IdFilters).To(Equal(idFilterModel))
				Expect(filtersModel.CategoryFilters["foo"]).To(Equal(*categoryFilterModel))

				// Construct an instance of the SyndicationCluster model
				syndicationClusterModel := new(catalogmanagementv1.SyndicationCluster)
				Expect(syndicationClusterModel).ToNot(BeNil())
				syndicationClusterModel.Region = core.StringPtr("testString")
				syndicationClusterModel.ID = core.StringPtr("testString")
				syndicationClusterModel.Name = core.StringPtr("testString")
				syndicationClusterModel.ResourceGroupName = core.StringPtr("testString")
				syndicationClusterModel.Type = core.StringPtr("testString")
				syndicationClusterModel.Namespaces = []string{"testString"}
				syndicationClusterModel.AllNamespaces = core.BoolPtr(true)
				Expect(syndicationClusterModel.Region).To(Equal(core.StringPtr("testString")))
				Expect(syndicationClusterModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(syndicationClusterModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(syndicationClusterModel.ResourceGroupName).To(Equal(core.StringPtr("testString")))
				Expect(syndicationClusterModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(syndicationClusterModel.Namespaces).To(Equal([]string{"testString"}))
				Expect(syndicationClusterModel.AllNamespaces).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the SyndicationHistory model
				syndicationHistoryModel := new(catalogmanagementv1.SyndicationHistory)
				Expect(syndicationHistoryModel).ToNot(BeNil())
				syndicationHistoryModel.Namespaces = []string{"testString"}
				syndicationHistoryModel.Clusters = []catalogmanagementv1.SyndicationCluster{*syndicationClusterModel}
				syndicationHistoryModel.LastRun = CreateMockDateTime()
				Expect(syndicationHistoryModel.Namespaces).To(Equal([]string{"testString"}))
				Expect(syndicationHistoryModel.Clusters).To(Equal([]catalogmanagementv1.SyndicationCluster{*syndicationClusterModel}))
				Expect(syndicationHistoryModel.LastRun).To(Equal(CreateMockDateTime()))

				// Construct an instance of the SyndicationAuthorization model
				syndicationAuthorizationModel := new(catalogmanagementv1.SyndicationAuthorization)
				Expect(syndicationAuthorizationModel).ToNot(BeNil())
				syndicationAuthorizationModel.Token = core.StringPtr("testString")
				syndicationAuthorizationModel.LastRun = CreateMockDateTime()
				Expect(syndicationAuthorizationModel.Token).To(Equal(core.StringPtr("testString")))
				Expect(syndicationAuthorizationModel.LastRun).To(Equal(CreateMockDateTime()))

				// Construct an instance of the SyndicationResource model
				syndicationResourceModel := new(catalogmanagementv1.SyndicationResource)
				Expect(syndicationResourceModel).ToNot(BeNil())
				syndicationResourceModel.RemoveRelatedComponents = core.BoolPtr(true)
				syndicationResourceModel.Clusters = []catalogmanagementv1.SyndicationCluster{*syndicationClusterModel}
				syndicationResourceModel.History = syndicationHistoryModel
				syndicationResourceModel.Authorization = syndicationAuthorizationModel
				Expect(syndicationResourceModel.RemoveRelatedComponents).To(Equal(core.BoolPtr(true)))
				Expect(syndicationResourceModel.Clusters).To(Equal([]catalogmanagementv1.SyndicationCluster{*syndicationClusterModel}))
				Expect(syndicationResourceModel.History).To(Equal(syndicationHistoryModel))
				Expect(syndicationResourceModel.Authorization).To(Equal(syndicationAuthorizationModel))

				// Construct an instance of the ReplaceCatalogOptions model
				catalogIdentifier := "testString"
				replaceCatalogOptionsModel := catalogManagementService.NewReplaceCatalogOptions(catalogIdentifier)
				replaceCatalogOptionsModel.SetCatalogIdentifier("testString")
				replaceCatalogOptionsModel.SetID("testString")
				replaceCatalogOptionsModel.SetRev("testString")
				replaceCatalogOptionsModel.SetLabel("testString")
				replaceCatalogOptionsModel.SetShortDescription("testString")
				replaceCatalogOptionsModel.SetCatalogIconURL("testString")
				replaceCatalogOptionsModel.SetTags([]string{"testString"})
				replaceCatalogOptionsModel.SetURL("testString")
				replaceCatalogOptionsModel.SetCrn("testString")
				replaceCatalogOptionsModel.SetOfferingsURL("testString")
				replaceCatalogOptionsModel.SetFeatures([]catalogmanagementv1.Feature{*featureModel})
				replaceCatalogOptionsModel.SetDisabled(true)
				replaceCatalogOptionsModel.SetCreated(CreateMockDateTime())
				replaceCatalogOptionsModel.SetUpdated(CreateMockDateTime())
				replaceCatalogOptionsModel.SetResourceGroupID("testString")
				replaceCatalogOptionsModel.SetOwningAccount("testString")
				replaceCatalogOptionsModel.SetCatalogFilters(filtersModel)
				replaceCatalogOptionsModel.SetSyndicationSettings(syndicationResourceModel)
				replaceCatalogOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceCatalogOptionsModel).ToNot(BeNil())
				Expect(replaceCatalogOptionsModel.CatalogIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(replaceCatalogOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(replaceCatalogOptionsModel.Rev).To(Equal(core.StringPtr("testString")))
				Expect(replaceCatalogOptionsModel.Label).To(Equal(core.StringPtr("testString")))
				Expect(replaceCatalogOptionsModel.ShortDescription).To(Equal(core.StringPtr("testString")))
				Expect(replaceCatalogOptionsModel.CatalogIconURL).To(Equal(core.StringPtr("testString")))
				Expect(replaceCatalogOptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(replaceCatalogOptionsModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(replaceCatalogOptionsModel.Crn).To(Equal(core.StringPtr("testString")))
				Expect(replaceCatalogOptionsModel.OfferingsURL).To(Equal(core.StringPtr("testString")))
				Expect(replaceCatalogOptionsModel.Features).To(Equal([]catalogmanagementv1.Feature{*featureModel}))
				Expect(replaceCatalogOptionsModel.Disabled).To(Equal(core.BoolPtr(true)))
				Expect(replaceCatalogOptionsModel.Created).To(Equal(CreateMockDateTime()))
				Expect(replaceCatalogOptionsModel.Updated).To(Equal(CreateMockDateTime()))
				Expect(replaceCatalogOptionsModel.ResourceGroupID).To(Equal(core.StringPtr("testString")))
				Expect(replaceCatalogOptionsModel.OwningAccount).To(Equal(core.StringPtr("testString")))
				Expect(replaceCatalogOptionsModel.CatalogFilters).To(Equal(filtersModel))
				Expect(replaceCatalogOptionsModel.SyndicationSettings).To(Equal(syndicationResourceModel))
				Expect(replaceCatalogOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceEnterpriseOptions successfully`, func() {
				// Construct an instance of the FilterTerms model
				filterTermsModel := new(catalogmanagementv1.FilterTerms)
				Expect(filterTermsModel).ToNot(BeNil())
				filterTermsModel.FilterTerms = []string{"testString"}
				Expect(filterTermsModel.FilterTerms).To(Equal([]string{"testString"}))

				// Construct an instance of the CategoryFilter model
				categoryFilterModel := new(catalogmanagementv1.CategoryFilter)
				Expect(categoryFilterModel).ToNot(BeNil())
				categoryFilterModel.Include = core.BoolPtr(true)
				categoryFilterModel.Filter = filterTermsModel
				Expect(categoryFilterModel.Include).To(Equal(core.BoolPtr(true)))
				Expect(categoryFilterModel.Filter).To(Equal(filterTermsModel))

				// Construct an instance of the IDFilter model
				idFilterModel := new(catalogmanagementv1.IDFilter)
				Expect(idFilterModel).ToNot(BeNil())
				idFilterModel.Include = filterTermsModel
				idFilterModel.Exclude = filterTermsModel
				Expect(idFilterModel.Include).To(Equal(filterTermsModel))
				Expect(idFilterModel.Exclude).To(Equal(filterTermsModel))

				// Construct an instance of the Filters model
				filtersModel := new(catalogmanagementv1.Filters)
				Expect(filtersModel).ToNot(BeNil())
				filtersModel.IncludeAll = core.BoolPtr(true)
				filtersModel.CategoryFilters = make(map[string]catalogmanagementv1.CategoryFilter)
				filtersModel.IdFilters = idFilterModel
				filtersModel.CategoryFilters["foo"] = *categoryFilterModel
				Expect(filtersModel.IncludeAll).To(Equal(core.BoolPtr(true)))
				Expect(filtersModel.IdFilters).To(Equal(idFilterModel))
				Expect(filtersModel.CategoryFilters["foo"]).To(Equal(*categoryFilterModel))

				// Construct an instance of the AccountGroup model
				accountGroupModel := new(catalogmanagementv1.AccountGroup)
				Expect(accountGroupModel).ToNot(BeNil())
				accountGroupModel.ID = core.StringPtr("testString")
				accountGroupModel.AccountFilters = filtersModel
				Expect(accountGroupModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(accountGroupModel.AccountFilters).To(Equal(filtersModel))

				// Construct an instance of the EnterpriseAccountGroups model
				enterpriseAccountGroupsModel := new(catalogmanagementv1.EnterpriseAccountGroups)
				Expect(enterpriseAccountGroupsModel).ToNot(BeNil())
				enterpriseAccountGroupsModel.Keys = accountGroupModel
				Expect(enterpriseAccountGroupsModel.Keys).To(Equal(accountGroupModel))

				// Construct an instance of the ReplaceEnterpriseOptions model
				enterpriseID := "testString"
				replaceEnterpriseOptionsModel := catalogManagementService.NewReplaceEnterpriseOptions(enterpriseID)
				replaceEnterpriseOptionsModel.SetEnterpriseID("testString")
				replaceEnterpriseOptionsModel.SetID("testString")
				replaceEnterpriseOptionsModel.SetRev("testString")
				replaceEnterpriseOptionsModel.SetAccountFilters(filtersModel)
				replaceEnterpriseOptionsModel.SetAccountGroups(enterpriseAccountGroupsModel)
				replaceEnterpriseOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceEnterpriseOptionsModel).ToNot(BeNil())
				Expect(replaceEnterpriseOptionsModel.EnterpriseID).To(Equal(core.StringPtr("testString")))
				Expect(replaceEnterpriseOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(replaceEnterpriseOptionsModel.Rev).To(Equal(core.StringPtr("testString")))
				Expect(replaceEnterpriseOptionsModel.AccountFilters).To(Equal(filtersModel))
				Expect(replaceEnterpriseOptionsModel.AccountGroups).To(Equal(enterpriseAccountGroupsModel))
				Expect(replaceEnterpriseOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceObjectOptions successfully`, func() {
				// Construct an instance of the PublishObject model
				publishObjectModel := new(catalogmanagementv1.PublishObject)
				Expect(publishObjectModel).ToNot(BeNil())
				publishObjectModel.PermitIbmPublicPublish = core.BoolPtr(true)
				publishObjectModel.IbmApproved = core.BoolPtr(true)
				publishObjectModel.PublicApproved = core.BoolPtr(true)
				publishObjectModel.PortalApprovalRecord = core.StringPtr("testString")
				publishObjectModel.PortalURL = core.StringPtr("testString")
				Expect(publishObjectModel.PermitIbmPublicPublish).To(Equal(core.BoolPtr(true)))
				Expect(publishObjectModel.IbmApproved).To(Equal(core.BoolPtr(true)))
				Expect(publishObjectModel.PublicApproved).To(Equal(core.BoolPtr(true)))
				Expect(publishObjectModel.PortalApprovalRecord).To(Equal(core.StringPtr("testString")))
				Expect(publishObjectModel.PortalURL).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the State model
				stateModel := new(catalogmanagementv1.State)
				Expect(stateModel).ToNot(BeNil())
				stateModel.Current = core.StringPtr("testString")
				stateModel.CurrentEntered = CreateMockDateTime()
				stateModel.Pending = core.StringPtr("testString")
				stateModel.PendingRequested = CreateMockDateTime()
				stateModel.Previous = core.StringPtr("testString")
				Expect(stateModel.Current).To(Equal(core.StringPtr("testString")))
				Expect(stateModel.CurrentEntered).To(Equal(CreateMockDateTime()))
				Expect(stateModel.Pending).To(Equal(core.StringPtr("testString")))
				Expect(stateModel.PendingRequested).To(Equal(CreateMockDateTime()))
				Expect(stateModel.Previous).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ReplaceObjectOptions model
				catalogIdentifier := "testString"
				objectIdentifier := "testString"
				replaceObjectOptionsModel := catalogManagementService.NewReplaceObjectOptions(catalogIdentifier, objectIdentifier)
				replaceObjectOptionsModel.SetCatalogIdentifier("testString")
				replaceObjectOptionsModel.SetObjectIdentifier("testString")
				replaceObjectOptionsModel.SetID("testString")
				replaceObjectOptionsModel.SetName("testString")
				replaceObjectOptionsModel.SetRev("testString")
				replaceObjectOptionsModel.SetCrn("testString")
				replaceObjectOptionsModel.SetURL("testString")
				replaceObjectOptionsModel.SetParentID("testString")
				replaceObjectOptionsModel.SetAllowList([]string{"testString"})
				replaceObjectOptionsModel.SetLabelI18n("testString")
				replaceObjectOptionsModel.SetLabel("testString")
				replaceObjectOptionsModel.SetTags([]string{"testString"})
				replaceObjectOptionsModel.SetCreated(CreateMockDateTime())
				replaceObjectOptionsModel.SetUpdated(CreateMockDateTime())
				replaceObjectOptionsModel.SetShortDescription("testString")
				replaceObjectOptionsModel.SetShortDescriptionI18n("testString")
				replaceObjectOptionsModel.SetKind("testString")
				replaceObjectOptionsModel.SetPublish(publishObjectModel)
				replaceObjectOptionsModel.SetState(stateModel)
				replaceObjectOptionsModel.SetCatalogID("testString")
				replaceObjectOptionsModel.SetCatalogName("testString")
				replaceObjectOptionsModel.SetData(map[string]interface{}{"anyKey": "anyValue"})
				replaceObjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceObjectOptionsModel).ToNot(BeNil())
				Expect(replaceObjectOptionsModel.CatalogIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(replaceObjectOptionsModel.ObjectIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(replaceObjectOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(replaceObjectOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(replaceObjectOptionsModel.Rev).To(Equal(core.StringPtr("testString")))
				Expect(replaceObjectOptionsModel.Crn).To(Equal(core.StringPtr("testString")))
				Expect(replaceObjectOptionsModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(replaceObjectOptionsModel.ParentID).To(Equal(core.StringPtr("testString")))
				Expect(replaceObjectOptionsModel.AllowList).To(Equal([]string{"testString"}))
				Expect(replaceObjectOptionsModel.LabelI18n).To(Equal(core.StringPtr("testString")))
				Expect(replaceObjectOptionsModel.Label).To(Equal(core.StringPtr("testString")))
				Expect(replaceObjectOptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(replaceObjectOptionsModel.Created).To(Equal(CreateMockDateTime()))
				Expect(replaceObjectOptionsModel.Updated).To(Equal(CreateMockDateTime()))
				Expect(replaceObjectOptionsModel.ShortDescription).To(Equal(core.StringPtr("testString")))
				Expect(replaceObjectOptionsModel.ShortDescriptionI18n).To(Equal(core.StringPtr("testString")))
				Expect(replaceObjectOptionsModel.Kind).To(Equal(core.StringPtr("testString")))
				Expect(replaceObjectOptionsModel.Publish).To(Equal(publishObjectModel))
				Expect(replaceObjectOptionsModel.State).To(Equal(stateModel))
				Expect(replaceObjectOptionsModel.CatalogID).To(Equal(core.StringPtr("testString")))
				Expect(replaceObjectOptionsModel.CatalogName).To(Equal(core.StringPtr("testString")))
				Expect(replaceObjectOptionsModel.Data).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(replaceObjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceOfferingIconOptions successfully`, func() {
				// Construct an instance of the ReplaceOfferingIconOptions model
				catalogIdentifier := "testString"
				offeringID := "testString"
				fileName := "testString"
				replaceOfferingIconOptionsModel := catalogManagementService.NewReplaceOfferingIconOptions(catalogIdentifier, offeringID, fileName)
				replaceOfferingIconOptionsModel.SetCatalogIdentifier("testString")
				replaceOfferingIconOptionsModel.SetOfferingID("testString")
				replaceOfferingIconOptionsModel.SetFileName("testString")
				replaceOfferingIconOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceOfferingIconOptionsModel).ToNot(BeNil())
				Expect(replaceOfferingIconOptionsModel.CatalogIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(replaceOfferingIconOptionsModel.OfferingID).To(Equal(core.StringPtr("testString")))
				Expect(replaceOfferingIconOptionsModel.FileName).To(Equal(core.StringPtr("testString")))
				Expect(replaceOfferingIconOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceOfferingOptions successfully`, func() {
				// Construct an instance of the Rating model
				ratingModel := new(catalogmanagementv1.Rating)
				Expect(ratingModel).ToNot(BeNil())
				ratingModel.OneStarCount = core.Int64Ptr(int64(38))
				ratingModel.TwoStarCount = core.Int64Ptr(int64(38))
				ratingModel.ThreeStarCount = core.Int64Ptr(int64(38))
				ratingModel.FourStarCount = core.Int64Ptr(int64(38))
				Expect(ratingModel.OneStarCount).To(Equal(core.Int64Ptr(int64(38))))
				Expect(ratingModel.TwoStarCount).To(Equal(core.Int64Ptr(int64(38))))
				Expect(ratingModel.ThreeStarCount).To(Equal(core.Int64Ptr(int64(38))))
				Expect(ratingModel.FourStarCount).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the Feature model
				featureModel := new(catalogmanagementv1.Feature)
				Expect(featureModel).ToNot(BeNil())
				featureModel.Title = core.StringPtr("testString")
				featureModel.Description = core.StringPtr("testString")
				Expect(featureModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(featureModel.Description).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Configuration model
				configurationModel := new(catalogmanagementv1.Configuration)
				Expect(configurationModel).ToNot(BeNil())
				configurationModel.Key = core.StringPtr("testString")
				configurationModel.Type = core.StringPtr("testString")
				configurationModel.DefaultValue = core.StringPtr("testString")
				configurationModel.ValueConstraint = core.StringPtr("testString")
				configurationModel.Description = core.StringPtr("testString")
				configurationModel.Required = core.BoolPtr(true)
				configurationModel.Options = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				configurationModel.Hidden = core.BoolPtr(true)
				Expect(configurationModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(configurationModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(configurationModel.DefaultValue).To(Equal(core.StringPtr("testString")))
				Expect(configurationModel.ValueConstraint).To(Equal(core.StringPtr("testString")))
				Expect(configurationModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(configurationModel.Required).To(Equal(core.BoolPtr(true)))
				Expect(configurationModel.Options).To(Equal([]interface{}{map[string]interface{}{"anyKey": "anyValue"}}))
				Expect(configurationModel.Hidden).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the Validation model
				validationModel := new(catalogmanagementv1.Validation)
				Expect(validationModel).ToNot(BeNil())
				validationModel.Validated = CreateMockDateTime()
				validationModel.Requested = CreateMockDateTime()
				validationModel.State = core.StringPtr("testString")
				validationModel.LastOperation = core.StringPtr("testString")
				validationModel.Target = map[string]interface{}{"anyKey": "anyValue"}
				Expect(validationModel.Validated).To(Equal(CreateMockDateTime()))
				Expect(validationModel.Requested).To(Equal(CreateMockDateTime()))
				Expect(validationModel.State).To(Equal(core.StringPtr("testString")))
				Expect(validationModel.LastOperation).To(Equal(core.StringPtr("testString")))
				Expect(validationModel.Target).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))

				// Construct an instance of the Resource model
				resourceModel := new(catalogmanagementv1.Resource)
				Expect(resourceModel).ToNot(BeNil())
				resourceModel.Type = core.StringPtr("mem")
				resourceModel.Value = core.StringPtr("testString")
				Expect(resourceModel.Type).To(Equal(core.StringPtr("mem")))
				Expect(resourceModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Script model
				scriptModel := new(catalogmanagementv1.Script)
				Expect(scriptModel).ToNot(BeNil())
				scriptModel.Instructions = core.StringPtr("testString")
				scriptModel.Script = core.StringPtr("testString")
				scriptModel.ScriptPermission = core.StringPtr("testString")
				scriptModel.DeleteScript = core.StringPtr("testString")
				scriptModel.Scope = core.StringPtr("testString")
				Expect(scriptModel.Instructions).To(Equal(core.StringPtr("testString")))
				Expect(scriptModel.Script).To(Equal(core.StringPtr("testString")))
				Expect(scriptModel.ScriptPermission).To(Equal(core.StringPtr("testString")))
				Expect(scriptModel.DeleteScript).To(Equal(core.StringPtr("testString")))
				Expect(scriptModel.Scope).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the VersionEntitlement model
				versionEntitlementModel := new(catalogmanagementv1.VersionEntitlement)
				Expect(versionEntitlementModel).ToNot(BeNil())
				versionEntitlementModel.ProviderName = core.StringPtr("testString")
				versionEntitlementModel.ProviderID = core.StringPtr("testString")
				versionEntitlementModel.ProductID = core.StringPtr("testString")
				versionEntitlementModel.PartNumbers = []string{"testString"}
				versionEntitlementModel.ImageRepoName = core.StringPtr("testString")
				Expect(versionEntitlementModel.ProviderName).To(Equal(core.StringPtr("testString")))
				Expect(versionEntitlementModel.ProviderID).To(Equal(core.StringPtr("testString")))
				Expect(versionEntitlementModel.ProductID).To(Equal(core.StringPtr("testString")))
				Expect(versionEntitlementModel.PartNumbers).To(Equal([]string{"testString"}))
				Expect(versionEntitlementModel.ImageRepoName).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the License model
				licenseModel := new(catalogmanagementv1.License)
				Expect(licenseModel).ToNot(BeNil())
				licenseModel.ID = core.StringPtr("testString")
				licenseModel.Name = core.StringPtr("testString")
				licenseModel.Type = core.StringPtr("testString")
				licenseModel.URL = core.StringPtr("testString")
				licenseModel.Description = core.StringPtr("testString")
				Expect(licenseModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(licenseModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(licenseModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(licenseModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(licenseModel.Description).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the State model
				stateModel := new(catalogmanagementv1.State)
				Expect(stateModel).ToNot(BeNil())
				stateModel.Current = core.StringPtr("testString")
				stateModel.CurrentEntered = CreateMockDateTime()
				stateModel.Pending = core.StringPtr("testString")
				stateModel.PendingRequested = CreateMockDateTime()
				stateModel.Previous = core.StringPtr("testString")
				Expect(stateModel.Current).To(Equal(core.StringPtr("testString")))
				Expect(stateModel.CurrentEntered).To(Equal(CreateMockDateTime()))
				Expect(stateModel.Pending).To(Equal(core.StringPtr("testString")))
				Expect(stateModel.PendingRequested).To(Equal(CreateMockDateTime()))
				Expect(stateModel.Previous).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Version model
				versionModel := new(catalogmanagementv1.Version)
				Expect(versionModel).ToNot(BeNil())
				versionModel.ID = core.StringPtr("testString")
				versionModel.Rev = core.StringPtr("testString")
				versionModel.Crn = core.StringPtr("testString")
				versionModel.Version = core.StringPtr("testString")
				versionModel.Sha = core.StringPtr("testString")
				versionModel.Created = CreateMockDateTime()
				versionModel.Updated = CreateMockDateTime()
				versionModel.OfferingID = core.StringPtr("testString")
				versionModel.CatalogID = core.StringPtr("testString")
				versionModel.KindID = core.StringPtr("testString")
				versionModel.Tags = []string{"testString"}
				versionModel.RepoURL = core.StringPtr("testString")
				versionModel.SourceURL = core.StringPtr("testString")
				versionModel.TgzURL = core.StringPtr("testString")
				versionModel.Configuration = []catalogmanagementv1.Configuration{*configurationModel}
				versionModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}
				versionModel.Validation = validationModel
				versionModel.RequiredResources = []catalogmanagementv1.Resource{*resourceModel}
				versionModel.SingleInstance = core.BoolPtr(true)
				versionModel.Install = scriptModel
				versionModel.PreInstall = []catalogmanagementv1.Script{*scriptModel}
				versionModel.Entitlement = versionEntitlementModel
				versionModel.Licenses = []catalogmanagementv1.License{*licenseModel}
				versionModel.ImageManifestURL = core.StringPtr("testString")
				versionModel.Deprecated = core.BoolPtr(true)
				versionModel.PackageVersion = core.StringPtr("testString")
				versionModel.State = stateModel
				versionModel.VersionLocator = core.StringPtr("testString")
				versionModel.ConsoleURL = core.StringPtr("testString")
				versionModel.LongDescription = core.StringPtr("testString")
				versionModel.WhitelistedAccounts = []string{"testString"}
				Expect(versionModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(versionModel.Rev).To(Equal(core.StringPtr("testString")))
				Expect(versionModel.Crn).To(Equal(core.StringPtr("testString")))
				Expect(versionModel.Version).To(Equal(core.StringPtr("testString")))
				Expect(versionModel.Sha).To(Equal(core.StringPtr("testString")))
				Expect(versionModel.Created).To(Equal(CreateMockDateTime()))
				Expect(versionModel.Updated).To(Equal(CreateMockDateTime()))
				Expect(versionModel.OfferingID).To(Equal(core.StringPtr("testString")))
				Expect(versionModel.CatalogID).To(Equal(core.StringPtr("testString")))
				Expect(versionModel.KindID).To(Equal(core.StringPtr("testString")))
				Expect(versionModel.Tags).To(Equal([]string{"testString"}))
				Expect(versionModel.RepoURL).To(Equal(core.StringPtr("testString")))
				Expect(versionModel.SourceURL).To(Equal(core.StringPtr("testString")))
				Expect(versionModel.TgzURL).To(Equal(core.StringPtr("testString")))
				Expect(versionModel.Configuration).To(Equal([]catalogmanagementv1.Configuration{*configurationModel}))
				Expect(versionModel.Metadata).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(versionModel.Validation).To(Equal(validationModel))
				Expect(versionModel.RequiredResources).To(Equal([]catalogmanagementv1.Resource{*resourceModel}))
				Expect(versionModel.SingleInstance).To(Equal(core.BoolPtr(true)))
				Expect(versionModel.Install).To(Equal(scriptModel))
				Expect(versionModel.PreInstall).To(Equal([]catalogmanagementv1.Script{*scriptModel}))
				Expect(versionModel.Entitlement).To(Equal(versionEntitlementModel))
				Expect(versionModel.Licenses).To(Equal([]catalogmanagementv1.License{*licenseModel}))
				Expect(versionModel.ImageManifestURL).To(Equal(core.StringPtr("testString")))
				Expect(versionModel.Deprecated).To(Equal(core.BoolPtr(true)))
				Expect(versionModel.PackageVersion).To(Equal(core.StringPtr("testString")))
				Expect(versionModel.State).To(Equal(stateModel))
				Expect(versionModel.VersionLocator).To(Equal(core.StringPtr("testString")))
				Expect(versionModel.ConsoleURL).To(Equal(core.StringPtr("testString")))
				Expect(versionModel.LongDescription).To(Equal(core.StringPtr("testString")))
				Expect(versionModel.WhitelistedAccounts).To(Equal([]string{"testString"}))

				// Construct an instance of the Deployment model
				deploymentModel := new(catalogmanagementv1.Deployment)
				Expect(deploymentModel).ToNot(BeNil())
				deploymentModel.ID = core.StringPtr("testString")
				deploymentModel.Label = core.StringPtr("testString")
				deploymentModel.Name = core.StringPtr("testString")
				deploymentModel.ShortDescription = core.StringPtr("testString")
				deploymentModel.LongDescription = core.StringPtr("testString")
				deploymentModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}
				deploymentModel.Tags = []string{"testString"}
				deploymentModel.Created = CreateMockDateTime()
				deploymentModel.Updated = CreateMockDateTime()
				Expect(deploymentModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deploymentModel.Label).To(Equal(core.StringPtr("testString")))
				Expect(deploymentModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(deploymentModel.ShortDescription).To(Equal(core.StringPtr("testString")))
				Expect(deploymentModel.LongDescription).To(Equal(core.StringPtr("testString")))
				Expect(deploymentModel.Metadata).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(deploymentModel.Tags).To(Equal([]string{"testString"}))
				Expect(deploymentModel.Created).To(Equal(CreateMockDateTime()))
				Expect(deploymentModel.Updated).To(Equal(CreateMockDateTime()))

				// Construct an instance of the Plan model
				planModel := new(catalogmanagementv1.Plan)
				Expect(planModel).ToNot(BeNil())
				planModel.ID = core.StringPtr("testString")
				planModel.Label = core.StringPtr("testString")
				planModel.Name = core.StringPtr("testString")
				planModel.ShortDescription = core.StringPtr("testString")
				planModel.LongDescription = core.StringPtr("testString")
				planModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}
				planModel.Tags = []string{"testString"}
				planModel.AdditionalFeatures = []catalogmanagementv1.Feature{*featureModel}
				planModel.Created = CreateMockDateTime()
				planModel.Updated = CreateMockDateTime()
				planModel.Deployments = []catalogmanagementv1.Deployment{*deploymentModel}
				Expect(planModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(planModel.Label).To(Equal(core.StringPtr("testString")))
				Expect(planModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(planModel.ShortDescription).To(Equal(core.StringPtr("testString")))
				Expect(planModel.LongDescription).To(Equal(core.StringPtr("testString")))
				Expect(planModel.Metadata).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(planModel.Tags).To(Equal([]string{"testString"}))
				Expect(planModel.AdditionalFeatures).To(Equal([]catalogmanagementv1.Feature{*featureModel}))
				Expect(planModel.Created).To(Equal(CreateMockDateTime()))
				Expect(planModel.Updated).To(Equal(CreateMockDateTime()))
				Expect(planModel.Deployments).To(Equal([]catalogmanagementv1.Deployment{*deploymentModel}))

				// Construct an instance of the Kind model
				kindModel := new(catalogmanagementv1.Kind)
				Expect(kindModel).ToNot(BeNil())
				kindModel.ID = core.StringPtr("testString")
				kindModel.FormatKind = core.StringPtr("testString")
				kindModel.TargetKind = core.StringPtr("testString")
				kindModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}
				kindModel.InstallDescription = core.StringPtr("testString")
				kindModel.Tags = []string{"testString"}
				kindModel.AdditionalFeatures = []catalogmanagementv1.Feature{*featureModel}
				kindModel.Created = CreateMockDateTime()
				kindModel.Updated = CreateMockDateTime()
				kindModel.Versions = []catalogmanagementv1.Version{*versionModel}
				kindModel.Plans = []catalogmanagementv1.Plan{*planModel}
				Expect(kindModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(kindModel.FormatKind).To(Equal(core.StringPtr("testString")))
				Expect(kindModel.TargetKind).To(Equal(core.StringPtr("testString")))
				Expect(kindModel.Metadata).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(kindModel.InstallDescription).To(Equal(core.StringPtr("testString")))
				Expect(kindModel.Tags).To(Equal([]string{"testString"}))
				Expect(kindModel.AdditionalFeatures).To(Equal([]catalogmanagementv1.Feature{*featureModel}))
				Expect(kindModel.Created).To(Equal(CreateMockDateTime()))
				Expect(kindModel.Updated).To(Equal(CreateMockDateTime()))
				Expect(kindModel.Versions).To(Equal([]catalogmanagementv1.Version{*versionModel}))
				Expect(kindModel.Plans).To(Equal([]catalogmanagementv1.Plan{*planModel}))

				// Construct an instance of the RepoInfo model
				repoInfoModel := new(catalogmanagementv1.RepoInfo)
				Expect(repoInfoModel).ToNot(BeNil())
				repoInfoModel.Token = core.StringPtr("testString")
				repoInfoModel.Type = core.StringPtr("testString")
				Expect(repoInfoModel.Token).To(Equal(core.StringPtr("testString")))
				Expect(repoInfoModel.Type).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ReplaceOfferingOptions model
				catalogIdentifier := "testString"
				offeringID := "testString"
				replaceOfferingOptionsModel := catalogManagementService.NewReplaceOfferingOptions(catalogIdentifier, offeringID)
				replaceOfferingOptionsModel.SetCatalogIdentifier("testString")
				replaceOfferingOptionsModel.SetOfferingID("testString")
				replaceOfferingOptionsModel.SetID("testString")
				replaceOfferingOptionsModel.SetRev("testString")
				replaceOfferingOptionsModel.SetURL("testString")
				replaceOfferingOptionsModel.SetCrn("testString")
				replaceOfferingOptionsModel.SetLabel("testString")
				replaceOfferingOptionsModel.SetName("testString")
				replaceOfferingOptionsModel.SetOfferingIconURL("testString")
				replaceOfferingOptionsModel.SetOfferingDocsURL("testString")
				replaceOfferingOptionsModel.SetOfferingSupportURL("testString")
				replaceOfferingOptionsModel.SetTags([]string{"testString"})
				replaceOfferingOptionsModel.SetRating(ratingModel)
				replaceOfferingOptionsModel.SetCreated(CreateMockDateTime())
				replaceOfferingOptionsModel.SetUpdated(CreateMockDateTime())
				replaceOfferingOptionsModel.SetShortDescription("testString")
				replaceOfferingOptionsModel.SetLongDescription("testString")
				replaceOfferingOptionsModel.SetFeatures([]catalogmanagementv1.Feature{*featureModel})
				replaceOfferingOptionsModel.SetKinds([]catalogmanagementv1.Kind{*kindModel})
				replaceOfferingOptionsModel.SetPermitRequestIbmPublicPublish(true)
				replaceOfferingOptionsModel.SetIbmPublishApproved(true)
				replaceOfferingOptionsModel.SetPublicPublishApproved(true)
				replaceOfferingOptionsModel.SetPublicOriginalCrn("testString")
				replaceOfferingOptionsModel.SetPublishPublicCrn("testString")
				replaceOfferingOptionsModel.SetPortalApprovalRecord("testString")
				replaceOfferingOptionsModel.SetPortalUiURL("testString")
				replaceOfferingOptionsModel.SetCatalogID("testString")
				replaceOfferingOptionsModel.SetCatalogName("testString")
				replaceOfferingOptionsModel.SetMetadata(map[string]interface{}{"anyKey": "anyValue"})
				replaceOfferingOptionsModel.SetDisclaimer("testString")
				replaceOfferingOptionsModel.SetHidden(true)
				replaceOfferingOptionsModel.SetProvider("testString")
				replaceOfferingOptionsModel.SetRepoInfo(repoInfoModel)
				replaceOfferingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceOfferingOptionsModel).ToNot(BeNil())
				Expect(replaceOfferingOptionsModel.CatalogIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(replaceOfferingOptionsModel.OfferingID).To(Equal(core.StringPtr("testString")))
				Expect(replaceOfferingOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(replaceOfferingOptionsModel.Rev).To(Equal(core.StringPtr("testString")))
				Expect(replaceOfferingOptionsModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(replaceOfferingOptionsModel.Crn).To(Equal(core.StringPtr("testString")))
				Expect(replaceOfferingOptionsModel.Label).To(Equal(core.StringPtr("testString")))
				Expect(replaceOfferingOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(replaceOfferingOptionsModel.OfferingIconURL).To(Equal(core.StringPtr("testString")))
				Expect(replaceOfferingOptionsModel.OfferingDocsURL).To(Equal(core.StringPtr("testString")))
				Expect(replaceOfferingOptionsModel.OfferingSupportURL).To(Equal(core.StringPtr("testString")))
				Expect(replaceOfferingOptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(replaceOfferingOptionsModel.Rating).To(Equal(ratingModel))
				Expect(replaceOfferingOptionsModel.Created).To(Equal(CreateMockDateTime()))
				Expect(replaceOfferingOptionsModel.Updated).To(Equal(CreateMockDateTime()))
				Expect(replaceOfferingOptionsModel.ShortDescription).To(Equal(core.StringPtr("testString")))
				Expect(replaceOfferingOptionsModel.LongDescription).To(Equal(core.StringPtr("testString")))
				Expect(replaceOfferingOptionsModel.Features).To(Equal([]catalogmanagementv1.Feature{*featureModel}))
				Expect(replaceOfferingOptionsModel.Kinds).To(Equal([]catalogmanagementv1.Kind{*kindModel}))
				Expect(replaceOfferingOptionsModel.PermitRequestIbmPublicPublish).To(Equal(core.BoolPtr(true)))
				Expect(replaceOfferingOptionsModel.IbmPublishApproved).To(Equal(core.BoolPtr(true)))
				Expect(replaceOfferingOptionsModel.PublicPublishApproved).To(Equal(core.BoolPtr(true)))
				Expect(replaceOfferingOptionsModel.PublicOriginalCrn).To(Equal(core.StringPtr("testString")))
				Expect(replaceOfferingOptionsModel.PublishPublicCrn).To(Equal(core.StringPtr("testString")))
				Expect(replaceOfferingOptionsModel.PortalApprovalRecord).To(Equal(core.StringPtr("testString")))
				Expect(replaceOfferingOptionsModel.PortalUiURL).To(Equal(core.StringPtr("testString")))
				Expect(replaceOfferingOptionsModel.CatalogID).To(Equal(core.StringPtr("testString")))
				Expect(replaceOfferingOptionsModel.CatalogName).To(Equal(core.StringPtr("testString")))
				Expect(replaceOfferingOptionsModel.Metadata).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(replaceOfferingOptionsModel.Disclaimer).To(Equal(core.StringPtr("testString")))
				Expect(replaceOfferingOptionsModel.Hidden).To(Equal(core.BoolPtr(true)))
				Expect(replaceOfferingOptionsModel.Provider).To(Equal(core.StringPtr("testString")))
				Expect(replaceOfferingOptionsModel.RepoInfo).To(Equal(repoInfoModel))
				Expect(replaceOfferingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceOperatorOptions successfully`, func() {
				// Construct an instance of the ReplaceOperatorOptions model
				xAuthRefreshToken := "testString"
				replaceOperatorOptionsModel := catalogManagementService.NewReplaceOperatorOptions(xAuthRefreshToken)
				replaceOperatorOptionsModel.SetXAuthRefreshToken("testString")
				replaceOperatorOptionsModel.SetClusterID("testString")
				replaceOperatorOptionsModel.SetRegion("testString")
				replaceOperatorOptionsModel.SetNamespaces([]string{"testString"})
				replaceOperatorOptionsModel.SetAllNamespaces(true)
				replaceOperatorOptionsModel.SetVersionLocatorID("testString")
				replaceOperatorOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceOperatorOptionsModel).ToNot(BeNil())
				Expect(replaceOperatorOptionsModel.XAuthRefreshToken).To(Equal(core.StringPtr("testString")))
				Expect(replaceOperatorOptionsModel.ClusterID).To(Equal(core.StringPtr("testString")))
				Expect(replaceOperatorOptionsModel.Region).To(Equal(core.StringPtr("testString")))
				Expect(replaceOperatorOptionsModel.Namespaces).To(Equal([]string{"testString"}))
				Expect(replaceOperatorOptionsModel.AllNamespaces).To(Equal(core.BoolPtr(true)))
				Expect(replaceOperatorOptionsModel.VersionLocatorID).To(Equal(core.StringPtr("testString")))
				Expect(replaceOperatorOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSearchLicenseOfferingsOptions successfully`, func() {
				// Construct an instance of the SearchLicenseOfferingsOptions model
				q := "testString"
				searchLicenseOfferingsOptionsModel := catalogManagementService.NewSearchLicenseOfferingsOptions(q)
				searchLicenseOfferingsOptionsModel.SetQ("testString")
				searchLicenseOfferingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(searchLicenseOfferingsOptionsModel).ToNot(BeNil())
				Expect(searchLicenseOfferingsOptionsModel.Q).To(Equal(core.StringPtr("testString")))
				Expect(searchLicenseOfferingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSearchLicenseVersionsOptions successfully`, func() {
				// Construct an instance of the SearchLicenseVersionsOptions model
				q := "testString"
				searchLicenseVersionsOptionsModel := catalogManagementService.NewSearchLicenseVersionsOptions(q)
				searchLicenseVersionsOptionsModel.SetQ("testString")
				searchLicenseVersionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(searchLicenseVersionsOptionsModel).ToNot(BeNil())
				Expect(searchLicenseVersionsOptionsModel.Q).To(Equal(core.StringPtr("testString")))
				Expect(searchLicenseVersionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSearchObjectsOptions successfully`, func() {
				// Construct an instance of the SearchObjectsOptions model
				query := "testString"
				searchObjectsOptionsModel := catalogManagementService.NewSearchObjectsOptions(query)
				searchObjectsOptionsModel.SetQuery("testString")
				searchObjectsOptionsModel.SetLimit(int64(38))
				searchObjectsOptionsModel.SetOffset(int64(38))
				searchObjectsOptionsModel.SetCollapse(true)
				searchObjectsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(searchObjectsOptionsModel).ToNot(BeNil())
				Expect(searchObjectsOptionsModel.Query).To(Equal(core.StringPtr("testString")))
				Expect(searchObjectsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(38))))
				Expect(searchObjectsOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(38))))
				Expect(searchObjectsOptionsModel.Collapse).To(Equal(core.BoolPtr(true)))
				Expect(searchObjectsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateCatalogAccountOptions successfully`, func() {
				// Construct an instance of the FilterTerms model
				filterTermsModel := new(catalogmanagementv1.FilterTerms)
				Expect(filterTermsModel).ToNot(BeNil())
				filterTermsModel.FilterTerms = []string{"testString"}
				Expect(filterTermsModel.FilterTerms).To(Equal([]string{"testString"}))

				// Construct an instance of the CategoryFilter model
				categoryFilterModel := new(catalogmanagementv1.CategoryFilter)
				Expect(categoryFilterModel).ToNot(BeNil())
				categoryFilterModel.Include = core.BoolPtr(true)
				categoryFilterModel.Filter = filterTermsModel
				Expect(categoryFilterModel.Include).To(Equal(core.BoolPtr(true)))
				Expect(categoryFilterModel.Filter).To(Equal(filterTermsModel))

				// Construct an instance of the IDFilter model
				idFilterModel := new(catalogmanagementv1.IDFilter)
				Expect(idFilterModel).ToNot(BeNil())
				idFilterModel.Include = filterTermsModel
				idFilterModel.Exclude = filterTermsModel
				Expect(idFilterModel.Include).To(Equal(filterTermsModel))
				Expect(idFilterModel.Exclude).To(Equal(filterTermsModel))

				// Construct an instance of the Filters model
				filtersModel := new(catalogmanagementv1.Filters)
				Expect(filtersModel).ToNot(BeNil())
				filtersModel.IncludeAll = core.BoolPtr(true)
				filtersModel.CategoryFilters = make(map[string]catalogmanagementv1.CategoryFilter)
				filtersModel.IdFilters = idFilterModel
				filtersModel.CategoryFilters["foo"] = *categoryFilterModel
				Expect(filtersModel.IncludeAll).To(Equal(core.BoolPtr(true)))
				Expect(filtersModel.IdFilters).To(Equal(idFilterModel))
				Expect(filtersModel.CategoryFilters["foo"]).To(Equal(*categoryFilterModel))

				// Construct an instance of the UpdateCatalogAccountOptions model
				updateCatalogAccountOptionsModel := catalogManagementService.NewUpdateCatalogAccountOptions()
				updateCatalogAccountOptionsModel.SetID("testString")
				updateCatalogAccountOptionsModel.SetAccountFilters(filtersModel)
				updateCatalogAccountOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateCatalogAccountOptionsModel).ToNot(BeNil())
				Expect(updateCatalogAccountOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateCatalogAccountOptionsModel.AccountFilters).To(Equal(filtersModel))
				Expect(updateCatalogAccountOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateOfferingIbmOptions successfully`, func() {
				// Construct an instance of the UpdateOfferingIbmOptions model
				catalogIdentifier := "testString"
				offeringID := "testString"
				approvalType := "allow_request"
				approved := "true"
				updateOfferingIbmOptionsModel := catalogManagementService.NewUpdateOfferingIbmOptions(catalogIdentifier, offeringID, approvalType, approved)
				updateOfferingIbmOptionsModel.SetCatalogIdentifier("testString")
				updateOfferingIbmOptionsModel.SetOfferingID("testString")
				updateOfferingIbmOptionsModel.SetApprovalType("allow_request")
				updateOfferingIbmOptionsModel.SetApproved("true")
				updateOfferingIbmOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateOfferingIbmOptionsModel).ToNot(BeNil())
				Expect(updateOfferingIbmOptionsModel.CatalogIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(updateOfferingIbmOptionsModel.OfferingID).To(Equal(core.StringPtr("testString")))
				Expect(updateOfferingIbmOptionsModel.ApprovalType).To(Equal(core.StringPtr("allow_request")))
				Expect(updateOfferingIbmOptionsModel.Approved).To(Equal(core.StringPtr("true")))
				Expect(updateOfferingIbmOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewValidationInstallOptions successfully`, func() {
				// Construct an instance of the DeployRequestBodySchematics model
				deployRequestBodySchematicsModel := new(catalogmanagementv1.DeployRequestBodySchematics)
				Expect(deployRequestBodySchematicsModel).ToNot(BeNil())
				deployRequestBodySchematicsModel.Name = core.StringPtr("testString")
				deployRequestBodySchematicsModel.Description = core.StringPtr("testString")
				deployRequestBodySchematicsModel.Tags = []string{"testString"}
				deployRequestBodySchematicsModel.ResourceGroupID = core.StringPtr("testString")
				Expect(deployRequestBodySchematicsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(deployRequestBodySchematicsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(deployRequestBodySchematicsModel.Tags).To(Equal([]string{"testString"}))
				Expect(deployRequestBodySchematicsModel.ResourceGroupID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ValidationInstallOptions model
				versionLocID := "testString"
				xAuthRefreshToken := "testString"
				validationInstallOptionsModel := catalogManagementService.NewValidationInstallOptions(versionLocID, xAuthRefreshToken)
				validationInstallOptionsModel.SetVersionLocID("testString")
				validationInstallOptionsModel.SetXAuthRefreshToken("testString")
				validationInstallOptionsModel.SetClusterID("testString")
				validationInstallOptionsModel.SetRegion("testString")
				validationInstallOptionsModel.SetNamespace("testString")
				validationInstallOptionsModel.SetOverrideValues(map[string]interface{}{"anyKey": "anyValue"})
				validationInstallOptionsModel.SetEntitlementApikey("testString")
				validationInstallOptionsModel.SetSchematics(deployRequestBodySchematicsModel)
				validationInstallOptionsModel.SetScript("testString")
				validationInstallOptionsModel.SetScriptID("testString")
				validationInstallOptionsModel.SetVersionLocatorID("testString")
				validationInstallOptionsModel.SetVcenterID("testString")
				validationInstallOptionsModel.SetVcenterUser("testString")
				validationInstallOptionsModel.SetVcenterPassword("testString")
				validationInstallOptionsModel.SetVcenterLocation("testString")
				validationInstallOptionsModel.SetVcenterDatastore("testString")
				validationInstallOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(validationInstallOptionsModel).ToNot(BeNil())
				Expect(validationInstallOptionsModel.VersionLocID).To(Equal(core.StringPtr("testString")))
				Expect(validationInstallOptionsModel.XAuthRefreshToken).To(Equal(core.StringPtr("testString")))
				Expect(validationInstallOptionsModel.ClusterID).To(Equal(core.StringPtr("testString")))
				Expect(validationInstallOptionsModel.Region).To(Equal(core.StringPtr("testString")))
				Expect(validationInstallOptionsModel.Namespace).To(Equal(core.StringPtr("testString")))
				Expect(validationInstallOptionsModel.OverrideValues).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(validationInstallOptionsModel.EntitlementApikey).To(Equal(core.StringPtr("testString")))
				Expect(validationInstallOptionsModel.Schematics).To(Equal(deployRequestBodySchematicsModel))
				Expect(validationInstallOptionsModel.Script).To(Equal(core.StringPtr("testString")))
				Expect(validationInstallOptionsModel.ScriptID).To(Equal(core.StringPtr("testString")))
				Expect(validationInstallOptionsModel.VersionLocatorID).To(Equal(core.StringPtr("testString")))
				Expect(validationInstallOptionsModel.VcenterID).To(Equal(core.StringPtr("testString")))
				Expect(validationInstallOptionsModel.VcenterUser).To(Equal(core.StringPtr("testString")))
				Expect(validationInstallOptionsModel.VcenterPassword).To(Equal(core.StringPtr("testString")))
				Expect(validationInstallOptionsModel.VcenterLocation).To(Equal(core.StringPtr("testString")))
				Expect(validationInstallOptionsModel.VcenterDatastore).To(Equal(core.StringPtr("testString")))
				Expect(validationInstallOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate()
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime()
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, 0)
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate() *strfmt.Date {
	d := strfmt.Date(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	return &d
}

func CreateMockDateTime() *strfmt.DateTime {
	d := strfmt.DateTime(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
