package config_test

import (
	"testing"

	"github.com/ilho-tiger/slack-notifier/pkg/config"
)

func Test_Add(t *testing.T) {
	c := config.InitConfig("testing")

	c.Add("CONCURRENCY", "concurrency", "3", "number of concurrent process(es)")
	c.Parse()

	v, ok := c.Get("CONCURRENCY")
	if !ok {
		t.Fail()
	}
	t.Log(v)
}
