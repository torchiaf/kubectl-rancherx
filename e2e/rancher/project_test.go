package e2e

import (
	"encoding/json"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/torchiaf/kubectl-rancherx/e2e"

	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
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

		It("should find project 'pippo'", FlakeAttempts(5), func() {
			out, _, err := rancherx.Run("get", "project", "--cluster-name", "local")
			Expect(err).To(BeNil())

			outTable := ParseOutTable(out)

			Expect(outTable[2][1]).To(Equal("pippo"))
		})

		It("should create project 'pippo2' and --set spec.description = bar1", func() {
			out, _, err := rancherx.Run("create", "project", "pippo2", "--display-name", "pippo2", "--cluster-name", "local", "--set", "foo=bar", "--set", "spec.description=bar1")
			Expect(err).To(BeNil())

			Expect(out).To(ContainSubstring("Project: \"pippo2\" created"))
		})

		It("should find project 'pippo2' with spec.description = bar1", FlakeAttempts(5), func() {
			out, _, err := rancherx.Run("get", "project", "pippo2", "--cluster-name", "local", "-o", "json")
			Expect(err).To(BeNil())

			project := v3.Project{}

			err = json.Unmarshal([]byte(out), &project)
			Expect(err).To(BeNil())

			Expect(project.Spec.DisplayName).To(Equal("pippo2"))
			Expect(project.Spec.Description).To(Equal("bar1"))
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

	Context("DeleteProject", Ordered, func() {
		It("should delete project 'pippo'", func() {
			out, _, err := rancherx.Run("delete", "project", "pippo", "--cluster-name", "local")
			Expect(err).To(BeNil())

			Expect(out).To(Equal("Project: \"pippo\" deleted \n"))
		})

		It("should not find project 'pippo'", FlakeAttempts(5), func() {
			out, _, err := rancherx.Run("get", "project", "pippo", "--cluster-name", "local")
			Expect(err).To(BeNil())

			Expect(out).To(Equal(""))
		})
	})

	AfterAll(func() {
		_, _, err := rancherx.Run("delete", "project", "pippo2", "--cluster-name", "local")
		Expect(err).To(BeNil())
	})
})
