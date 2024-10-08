package flag

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/pkg/errors"
	"github.com/torchiaf/kubectl-rancherx/pkg/log"
	"k8s.io/helm/pkg/strvals"
)

func MergeValues(ctx context.Context, obj any, set []string) error {

	log.Debug(
		ctx,
		"merge values",
		slog.Group("args",
			"configs", set,
		),
	)

	dest := make(map[string]interface{})

	jsonData, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonData, &dest)
	if err != nil {
		return err
	}

	// User specified a value via --set
	for _, value := range set {
		if err := strvals.ParseInto(value, dest); err != nil {
			return errors.Wrap(err, "failed parsing --set data")
		}
	}

	// Add here common flags values

	jsonData, err = json.Marshal(dest)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonData, &obj)
	if err != nil {
		return err
	}

	return nil
}
