package config_test

import (
	"os"
	"testing"

	"github.com/ilho-tiger/slack-notifier/pkg/config"
)

func Test_Add(t *testing.T) {
	c := config.InitConfig()

	c.Add("CONCURRENCY", "concurrency", "3", "number of concurrent process(es)")

	// set env var
	if err := os.Setenv("CONCURRENCY", "5"); err != nil {
		t.Fatal("failed to set env var")
	}

	c.Parse()

	v, ok := c.Get("CONCURRENCY")
	if !ok {
		t.Fatal("expected CONCURRENCY to be set")
	}
	if v != "5" {
		t.Fatalf("expected CONCURRENCY to be 5, got %s", v)
	}
}
