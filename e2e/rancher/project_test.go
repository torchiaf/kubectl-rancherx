package rancher_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/torchiaf/kubectl-rancherx/e2e"
)

var _ = Describe("Project", Ordered, func() {

	rancherx, err := e2e.NewKubectlRancherx()

	BeforeAll(func() {
		Expect(err).To(BeNil())
	})

	Context("ListProjects", func() {
		It("should get default projects list", func() {
			out, _, err := rancherx.Run("get", "project", "--cluster-name", "local")
			Expect(err).To(BeNil())

			Expect(out).To(ContainSubstring("System"))
		})
	})
})
