package view

import (
	"github.com/zitadel/zitadel/internal/errors"
	es_models "github.com/zitadel/zitadel/internal/eventstore/v1/models"
	"github.com/zitadel/zitadel/internal/repository/project"
)

func ProjectByIDQuery(id, instanceID string, latestSequence uint64) (*es_models.SearchQuery, error) {
	if id == "" {
		return nil, errors.ThrowPreconditionFailed(nil, "EVENT-dke74", "Errors.Project.ProjectIDMissing")
	}
	return es_models.NewSearchQuery().
		AddQuery().
		AggregateIDFilter(id).
		AggregateTypeFilter(project.AggregateType).
		LatestSequenceFilter(latestSequence).
		InstanceIDFilter(instanceID).
		SearchQuery(), nil
}
