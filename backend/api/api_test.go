package api

import (
	"log"
	"testing"
)

func TestSetup(t *testing.T) {
	engine := Setup()
	if engine == nil {
		t.Errorf("engine is nil which should be impossible")
	}
}

func TestRoutes(t *testing.T) {
	engine := Setup()
	routes := engine.Routes()
	if routes == nil {
		t.Errorf("routes is nil which should be impossible")
	}
	log.Printf("routes: %v\n", routes)

}
