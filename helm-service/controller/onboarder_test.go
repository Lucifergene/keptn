package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/cloudevents/sdk-go/pkg/cloudevents"
	"github.com/google/uuid"
	"github.com/keptn/keptn/helm-service/controller/helm"
	"github.com/keptn/keptn/helm-service/controller/mesh"

	keptnevents "github.com/keptn/go-utils/pkg/events"
	"github.com/keptn/go-utils/pkg/models"
	keptnmodels "github.com/keptn/go-utils/pkg/models"
	keptnutils "github.com/keptn/go-utils/pkg/utils"
	"github.com/stretchr/testify/assert"
)

const configBaseURL = "localhost:8080"
const projectName = "sockshop"
const serviceName = "carts"
const stage1 = "dev"
const stage2 = "staging"
const stage3 = "production"

const shipyard = `project: sockshop
stages:
    - {deployment_strategy: direct, name: dev, test_strategy: functional}
    - {deployment_strategy: blue_green_service, name: staging, test_strategy: performance}
    - {deployment_strategy: blue_green_service, name: production}`

func createTestProjet(t *testing.T) {

	prjHandler := keptnutils.NewProjectHandler(configBaseURL)
	prj := keptnmodels.Project{ProjectName: projectName}
	respErr, err := prjHandler.CreateProject(prj)
	check(err, t)
	assert.Nil(t, respErr, "Creating a project failed")

	// Send shipyard
	rHandler := keptnutils.NewResourceHandler(configBaseURL)
	shipyardURI := "shipyard.yaml"
	shipyardResource := models.Resource{ResourceURI: &shipyardURI, ResourceContent: shipyard}
	resources := []*models.Resource{&shipyardResource}
	_, err = rHandler.CreateProjectResources(projectName, resources)
	check(err, t)

	// Create stages
	stageHandler := keptnutils.NewStageHandler(configBaseURL)
	for _, stage := range []string{stage1, stage2, stage3} {

		respErr, err := stageHandler.CreateStage(projectName, stage)
		check(err, t)
		assert.Nil(t, respErr, "Creating a stage failed")
	}
}

// TestDoOnboard tests the onboarding of a new chart. Therefore, this test requires the configuration-service
// on localhost:8080
func TestDoOnboard(t *testing.T) {

	_, err := http.Get("http://" + configBaseURL)
	if err != nil {
		fmt.Println(err.Error())
		t.Skip("Skipping test; no local configuration-service available")
	}

	createTestProjet(t)

	data := helm.CreateHelmChartData(t)
	ce := cloudevents.New("0.2")
	dataBytes, err := json.Marshal(keptnevents.ServiceCreateEventData{Project: projectName, Service: serviceName, HelmChart: data})
	check(err, t)
	ce.Data = dataBytes

	id := uuid.New().String()
	err = DoOnboard(ce, mesh.NewIstioMesh(), keptnutils.NewLogger(id, "service.create", "helm-service"), id, configBaseURL, "test.keptn.sh")

	check(err, t)

	// TODO: Add assertions
}

func check(e error, t *testing.T) {
	if e != nil {
		t.Error(e)
	}
}
