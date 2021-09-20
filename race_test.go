package http2datarace

import (
	"context"
	"testing"

	"cloud.google.com/go/compute/metadata"

	"google.golang.org/api/compute/v1"
)

func TestDataRace(t *testing.T) {
	ctx := context.Background()

	computeService, err := compute.NewService(ctx)
	if err != nil {
		t.Fatalf("error creating compute service: %s", err)
	}

	projectID, err := metadata.ProjectID()
	if err != nil {
		t.Fatalf("error retrieving project ID: %s", err)
	}

	req := computeService.Instances.AggregatedList(projectID)

	err = req.Pages(
		ctx,
		func(l *compute.InstanceAggregatedList) error {
			return nil
		},
	)
	if err != nil {
		t.Fatalf("error performing aggregated list call: %s", err)
	}
}
