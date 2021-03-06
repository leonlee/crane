package dockerclient

import (
	"testing"
	"time"

	"github.com/Dataman-Cloud/crane/src/dockerclient/model"

	"github.com/docker/engine-api/types/swarm"
)

func TestValidateResources(t *testing.T) {
	if err := validateResources(&swarm.Resources{}); err != nil {
		t.Error("faild")
	} else {
		t.Log("pass")
	}

	if err := validateResources(&swarm.Resources{NanoCPUs: 1e6, MemoryBytes: 4 * 1024 * 1024}); err != nil {
		t.Error("faild")
	} else {
		t.Log("pass")
	}
}

func TestValidateResourceRequirements(t *testing.T) {
	if err := validateResourceRequirements(&swarm.ResourceRequirements{}); err != nil {
		t.Error("faild")
	} else {
		t.Log("pass")
	}

	if err := validateResourceRequirements(&swarm.ResourceRequirements{
		Limits:       &swarm.Resources{},
		Reservations: &swarm.Resources{},
	}); err != nil {
		t.Error("faild")
	} else {
		t.Log("pass")
	}
}

func TestValidateRestartPolicy(t *testing.T) {
	s := time.Duration(time.Second * 10)
	if err := validateRestartPolicy(&swarm.RestartPolicy{
		Delay:  &s,
		Window: &s,
	}); err != nil {
		t.Error(err)
	} else {
		t.Log("pass")
	}
}

func TestValidatePlacement(t *testing.T) {
	if err := validatePlacement(&swarm.Placement{
		Constraints: []string{"node==1"},
	}); err != nil {
		t.Error(err)
	} else {
		t.Log("pass")
	}
}

func TestValidateUpdate(t *testing.T) {
	s := time.Duration(time.Second * 10)
	if err := validateUpdate(&swarm.UpdateConfig{
		Delay: s,
	}); err != nil {
		t.Error(err)
	} else {
		t.Log("pass")
	}
}

func TestValidateTask(t *testing.T) {
	if err := validateTask(swarm.TaskSpec{}); err != nil {
		t.Error(err)
	} else {
		t.Log("pass")
	}
}

func TestValidateEndpointSpec(t *testing.T) {
	if err := validateEndpointSpec(&swarm.EndpointSpec{}); err != nil {
		t.Error(err)
	} else {
		t.Log("pass")
	}
}

func TestValidateCraneServiceSpec(t *testing.T) {
	if err := ValidateCraneServiceSpec(&model.CraneServiceSpec{
		Name: "test",
		TaskTemplate: swarm.TaskSpec{
			ContainerSpec: swarm.ContainerSpec{
				Image: "testimage:latest",
			},
		},
		UpdateConfig: &swarm.UpdateConfig{},
		EndpointSpec: &swarm.EndpointSpec{},
	}); err != nil {
		t.Error(err)
	} else {
		t.Log("pass")
	}
}

func TestValidateName(t *testing.T) {
	if err := validateName("testname"); err != nil {
		t.Error(err)
	} else {
		t.Log("pass")
	}
}

func TestValidateImageName(t *testing.T) {
	if err := validateImageName("testimage:latest"); err != nil {
		t.Error(err)
	} else {
		t.Log("pass")
	}
}
