package handlers

import (
	"github.com/rancher/event-subscriber/events"
	_ "github.com/rancher/go-machine-service/logging"
	"github.com/rancher/go-rancher/v2"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var (
	eventDOHostCreate = &events.Event{
		Name:         "physicalhost.create;handler=goMachineService",
		ID:           "7c933690-cfb8-4317-a84b-ee8b1be62f1",
		ReplyTo:      "reply.5495258476566540213",
		ResourceID:   "1ph1",
		ResourceType: "physicalHost",

		Data: map[string]interface{}{
			"name":       "ivan-do-h1",
			"kind":       "machine",
			"externalId": "a2ac1475-c64c-4331-a27f-f5e9b00be7b7",
			"accountId":  5,
			"hostname":   "ivan-do-h1",

			"digitaloceanConfig": map[string]interface{}{
				"image":  "ubuntu-16-04-x64",
				"region": "sfo2",
				"size":   "1gb",
			},
		},
	}
)

var apiClient, err = client.NewRancherClient(&client.ClientOpts{
	Timeout:   time.Second * 30,
	Url:       "http://localhost:8080/v2-beta",
	AccessKey: "service",
	SecretKey: "servicepass",
})

func TestCreateMachine(t *testing.T) {
	assert := require.New(t)

	err := CreateMachine(eventDOHostCreate, apiClient)
	assert.Nil(err)
}

func TestGetConnectionConfig(t *testing.T) {
	assert := require.New(t)

	machineDir := "/Users/ivan/.cattle/machine/machines/921d6c24-d22b-4032-a1bd-b3a7318b402e"
	machineName := "ivan-do3"

	conf, err := getConnectionConfig(machineDir, machineName)
	assert.Nil(err)
	assert.NotNil(conf)
}

func TestMachineExists(t *testing.T) {
	assert := require.New(t)

	exists, _ := machineExists("/Users/ivan/.cattle/machine/machines/fdccb274-92b2-4e77-ad79-f0a772622e80", "ivan-do1")
	assert.False(exists)
}
