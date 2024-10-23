package flag

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/pkg/errors"
	"github.com/torchiaf/kubectl-rancherx/pkg/log"
	"k8s.io/helm/pkg/strvals"
)

func MergeValues[T any](ctx context.Context, orig T, set []string) (T, error) {

	log.Debug(
		ctx,
		"merge values",
		slog.Group("args",
			"configs", set,
		),
	)

	res := orig

	dest := make(map[string]interface{})

	jsonData, err := json.Marshal(orig)
	if err != nil {
		return orig, err
	}

	err = json.Unmarshal(jsonData, &dest)
	if err != nil {
		return orig, err
	}

	// User specified a value via --set
	for _, value := range set {
		if err := strvals.ParseInto(value, dest); err != nil {
			return orig, errors.Wrap(err, "failed parsing --set data")
		}
	}

	// Add here common flags values

	jsonData, err = json.Marshal(dest)
	if err != nil {
		return orig, err
	}

	err = json.Unmarshal(jsonData, &res)
	if err != nil {
		return orig, err
	}

	return res, nil
}
