package integration_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"time"

	logcache "code.cloudfoundry.org/go-log-cache/v2/rpc/logcache_v1"
	"code.cloudfoundry.org/log-cache/integration/integrationfakes"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Gateway", func() {
	var (
		fakeLogCache *integrationfakes.FakeLogCache

		gatewayPort int
		gateway     *gexec.Session
	)

	BeforeEach(func() {
		port := 8000 + GinkgoParallelProcess()
		fakeLogCache = integrationfakes.NewFakeLogCache(port, nil)

		gatewayPort = 8081 + GinkgoParallelProcess()
	})

	JustBeforeEach(func() {
		fakeLogCache.Start()

		envVars := map[string]string{
			"ADDR":           fmt.Sprintf(":%d", gatewayPort),
			"LOG_CACHE_ADDR": fakeLogCache.Address(),
			"METRICS_PORT":   "0",
		}
		command := exec.Command(componentPaths.Gateway)
		for k, v := range envVars {
			command.Env = append(command.Env, fmt.Sprintf("%s=%s", k, v))
		}
		var err error
		gateway, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).ShouldNot(HaveOccurred())
	})

	JustAfterEach(func() {
		gateway.Interrupt().Wait(2 * time.Second)
		fakeLogCache.Stop()
	})

	Context("/api/v1/info endpoint", func() {
		var resp *http.Response

		JustBeforeEach(func() {
			u := fmt.Sprintf("http://localhost:%d/api/v1/info", gatewayPort)
			Eventually(func() error {
				var err error
				resp, err = http.Get(u)
				return err
			}, "5s").ShouldNot(HaveOccurred())
		})

		AfterEach(func() {
			resp.Body.Close()
		})

		It("returns 200", func() {
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
		})

		PIt("returns Content-Type as application/json", func() {
			Expect(resp.Header.Get("Content-Type")).To(Equal("application/json"))
		})

		PIt("returns Content-Length", func() {
			Expect(resp.Header.Get("Content-Length")).To(MatchRegexp("\\d+"))
		})

		Context("response body", func() {
			var body []byte

			JustBeforeEach(func() {
				var err error
				body, err = io.ReadAll(resp.Body)
				Expect(err).ToNot(HaveOccurred())
			})

			It("is a JSON with version and uptime information", func() {
				result := struct {
					Version  string `json:"version"`
					VMUptime string `json:"vm_uptime"`
				}{}
				err := json.Unmarshal(body, &result)
				Expect(err).ToNot(HaveOccurred())
				Expect(result.Version).To(Equal("1.2.3"))
				Expect(result.VMUptime).To(MatchRegexp("\\d+"))
			})

			It("has a newline at the end", func() {
				Expect(string(body)).To(MatchRegexp(".*\\n$"))
			})
		})
	})

	Context("api/v1/read endpoint", func() {
		var resp *http.Response

		JustBeforeEach(func() {
			u := fmt.Sprintf("http://localhost:%d/api/v1/read", gatewayPort)
			Eventually(func() error {
				var err error
				resp, err = http.Get(u)
				return err
			}, "5s").ShouldNot(HaveOccurred())
		})

		AfterEach(func() {
			resp.Body.Close()
		})

		It("returns 200", func() {
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
		})

		PIt("returns Content-Type as application/json", func() {
			Expect(resp.Header.Get("Content-Type")).To(Equal("application/json"))
		})

		PIt("returns Content-Length", func() {
			Expect(resp.Header.Get("Content-Length")).To(MatchRegexp("\\d+"))
		})

		PIt("forwards the request to Log Cache", func() {
			// Expect(flc.requestCount).To(Equal(1))
		})

		Context("response body", func() {
			var body []byte

			JustBeforeEach(func() {
				var err error
				body, err = io.ReadAll(resp.Body)
				Expect(err).ToNot(HaveOccurred())
			})

			PIt("is a JSON with envelopes", func() {
				var rr logcache.ReadResponse
				err := json.Unmarshal(body, &rr)
				Expect(err).ToNot(HaveOccurred())
				Expect(rr.Envelopes).To(HaveLen(0))
			})

			It("has a newline at the end", func() {
				Expect(string(body)).To(MatchRegexp(".*\\n$"))
			})
		})
	})
})
