package e2e

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/torchiaf/kubectl-rancherx/e2e"
)

var _ = Describe("Project", Ordered, func() {

	rancherx, cliErr := NewKubectlRancherx()

	BeforeAll(func() {
		Expect(cliErr).To(BeNil())
	})

	Context("ListProjects", func() {
		It("should get default projects list", func() {
			out, _, err := rancherx.Run("get", "project", "--cluster-name", "local")
			Expect(err).To(BeNil())

			outTable := ParseOutTable(out)

			Expect(len(outTable)).To(Equal(2))
			Expect(outTable[0][1]).To(Equal("Default"))
			Expect(outTable[1][1]).To(Equal("System"))
		})
	})

	Context("CreateProject", Ordered, func() {
		It("should create project 'pippo'", func() {
			out, _, err := rancherx.Run("create", "project", "pippo", "--display-name", "pippo", "--cluster-name", "local")
			Expect(err).To(BeNil())

			Expect(out).To(ContainSubstring("Project: \"pippo\" created"))
		})

		It("should get project 'pippo'", FlakeAttempts(5), func() {
			out, _, err := rancherx.Run("get", "project", "--cluster-name", "local")
			Expect(err).To(BeNil())

			outTable := ParseOutTable(out)

			Expect(outTable[2][1]).To(Equal("pippo"))
		})
	})

	Context("GetProject", Ordered, func() {
		It("should get project 'pippo'", func() {
			out, _, err := rancherx.Run("get", "project", "pippo", "--cluster-name", "local")
			Expect(err).To(BeNil())

			outTable := ParseOutTable(out)

			Expect(outTable[0][1]).To(Equal("pippo"))
		})
	})

})
