package controlplane

import (
	"context"
	"github.com/benbjohnson/clock"
	"github.com/keptn/go-utils/pkg/api/models"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

type UniformInterfaceMock struct {
	RegisterIntegrationFn func(models.Integration) (string, error)
	PingFn                func(string) (*models.Integration, error)
}

func (m *UniformInterfaceMock) Ping(integrationID string) (*models.Integration, error) {
	if m.PingFn != nil {
		return m.PingFn(integrationID)
	}
	panic("Ping() not implemented")
}
func (m *UniformInterfaceMock) RegisterIntegration(integration models.Integration) (string, error) {
	if m.RegisterIntegrationFn != nil {
		return m.RegisterIntegrationFn(integration)
	}
	panic("RegisterIntegraiton not imiplemented")
}

func (m *UniformInterfaceMock) CreateSubscription(integrationID string, subscription models.EventSubscription) (string, error) {
	panic("implement me")
}

func (m *UniformInterfaceMock) UnregisterIntegration(integrationID string) error {
	panic("implement me")
}

func (m *UniformInterfaceMock) GetRegistrations() ([]*models.Integration, error) {
	panic("implement me")
}

func TestSubscriptionSource(t *testing.T) {
	integrationID := "iID"
	integrationName := "integrationName"
	subscriptionID := "sID"

	initialRegistrationData := RegistrationData{
		ID:       "",
		Name:     integrationName,
		MetaData: models.MetaData{},
		Subscriptions: []models.EventSubscription{
			{
				ID:     "",
				Event:  "keptn.event",
				Filter: models.EventSubscriptionFilter{},
			},
		},
	}

	uniformInterface := &UniformInterfaceMock{
		RegisterIntegrationFn: func(integration models.Integration) (string, error) {
			require.Equal(t, initialRegistrationData, initialRegistrationData)
			return integrationID, nil
		},
		PingFn: func(id string) (*models.Integration, error) {
			require.Equal(t, id, integrationID)
			return &models.Integration{
				ID:       integrationID,
				Name:     integrationName,
				MetaData: models.MetaData{},
				Subscriptions: []models.EventSubscription{
					{
						ID:     subscriptionID,
						Event:  "keptn.event",
						Filter: models.EventSubscriptionFilter{},
					},
				},
			}, nil
		},
	}

	subscriptionSource := NewSubscriptionSource(uniformInterface)
	clock := clock.NewMock()
	subscriptionSource.clock = clock

	subscriptionUpdates := make(chan []models.EventSubscription)

	err := subscriptionSource.Start(context.TODO(), initialRegistrationData, subscriptionUpdates)
	require.NoError(t, err)
	clock.Add(5 * time.Second)
	subs := <-subscriptionUpdates
	require.Equal(t, 1, len(subs))
	clock.Add(5 * time.Second)
	subs = <-subscriptionUpdates
	require.Equal(t, 1, len(subs))
}