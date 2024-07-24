package e2e

import (
	"bytes"
	"os"
	"os/exec"
	"path"
	"strings"
)

type KubectlRancherx struct {
	binPath string
}

func NewKubectlRancherx() (*KubectlRancherx, error) {
	cmd := exec.Command(
		"git",
		"rev-parse",
		"--show-toplevel",
	)

	root, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	repoRoot := strings.TrimSpace(string(root))

	return &KubectlRancherx{
		binPath: path.Join(repoRoot, "dist", "kubectl-rancherx"),
	}, nil
}

func (e *KubectlRancherx) Run(args ...string) (string, string, error) {
	cmd := exec.Command(e.binPath, args...)
	cmd.Env = os.Environ()

	outBuff, errBuff := &bytes.Buffer{}, &bytes.Buffer{}
	cmd.Stdout = outBuff
	cmd.Stderr = errBuff

	err := cmd.Run()
	return outBuff.String(), errBuff.String(), err
}

func ParseOutTable(out string) [][]string {
	outTable := [][]string{}

	out = strings.TrimSpace(out)
	rows := strings.Split(out, "\n")

	for _, row := range rows {
		rowCells := strings.FieldsFunc(row, func(r rune) bool {
			return r == '\t'
		})

		outTable = append(outTable, rowCells)

	}
	return outTable
}
