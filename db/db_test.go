package db

import (
	"os"
	"testing"
)

func init() {
	os.Setenv("MONGODB_HOSTNAME", "mongo")
}

type mockClient struct {}

func mockDBConnector(dbURI string) interface{} {
	return &mockClient{}
}

func TestSingleton(t *testing.T) {
	connector := newDBConnector(mockDBConnector)
	client1 := connector.connect("dummy-url").(*mockClient)

	connector = newDBConnector(mockDBConnector)
	client2 := connector.connect("dummy-url").(*mockClient)

	// object should be same
	if client1 != client2 {
		t.Error("Objects have different references.")
	}
}