package storage

import (
	"github.com/nebulaclouds/nebulastdlib/contextutils"
	"github.com/nebulaclouds/nebulastdlib/promutils"
	"github.com/nebulaclouds/nebulastdlib/promutils/labeled"
)

var metrics *dataStoreMetrics

func init() {
	labeled.SetMetricKeys(contextutils.ProjectKey, contextutils.DomainKey, contextutils.WorkflowIDKey, contextutils.TaskIDKey)
	scope := promutils.NewTestScope()
	metrics = newDataStoreMetrics(scope)
}
