package config

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestLoadNotFoundFile(t *testing.T) {
	os.Clearenv()
	_, err := Load(".env.this_file_does_not_exist")
	if err == nil {
		t.Errorf("Unexpected Load() success: expected:failure, got:success")
	}
}

func TestLoadErrors(t *testing.T) {
	expectedErrors := [4]string{
		`"NOTION_API_SECRET" should not be empty`,
		`"NOTION_PAGE_ID" should not be empty`,
		`"IGDB_CLIENT_ID" should not be empty`,
		`"IGDB_SECRET" should not be empty`,
	}

	os.Clearenv()
	_, err := Load(".env.test_invalid")

	for _, expected := range expectedErrors {
		if !strings.Contains(err.Error(), expected) {
			t.Errorf("Expected to find error '%s'", expected)
		}
	}
}

func TestLoadDefaultValues(t *testing.T) {
	os.Clearenv()
	config, err := Load(".env.test_default")
	if err != nil {
		t.Errorf("Unexpected Load() error:'%s'", err)
	}

	if assertConfig("UpdaterHost", config.UpdaterHost, "127.0.0.1") != nil {
		t.Error(err)
	}

	if assertConfig("UpdaterPort", config.UpdaterPort, 8080) != nil {
		t.Error(err)
	}

	if assertConfig("WatcherTickDelay", config.WatcherTickDelay, 5) != nil {
		t.Error(err)
	}
}

func TestLoadSuccess(t *testing.T) {
	os.Clearenv()
	config, err := Load(".env.test_success")
	if err != nil {
		t.Errorf("Unexpected Load() error:'%s'", err)
	}

	if assertConfig("NotionAPISecret", config.NotionAPISecret, "u3M3sAdjqkLm8JSJihaGu3GfscAGE2") != nil {
		t.Error(err)
	}

	if assertConfig("NotionPageID", config.NotionPageID, "sEQCpMkvVCRfd4JAzVjmsd4EGYCFNC") != nil {
		t.Error(err)
	}

	if assertConfig("IGDBClientID", config.IGDBClientID, "pUnDQo4LzgjrYifQR5W7cAeaxACx92") != nil {
		t.Error(err)
	}

	if assertConfig("IGDBSecret", config.IGDBSecret, "WjMpHfNk6Mt4b4FV8TFT2pDiwiUeKu") != nil {
		t.Error(err)
	}

	if assertConfig("UpdaterHost", config.UpdaterHost, "76.388.29.99") != nil {
		t.Error(err)
	}

	if assertConfig("UpdaterPort", config.UpdaterPort, 7654) != nil {
		t.Error(err)
	}

	if assertConfig("WatcherTickDelay", config.WatcherTickDelay, 99) != nil {
		t.Error(err)
	}
}

func TestUpdaterURL(t *testing.T) {
	config := Config{
		UpdaterHost: "172.16.4.97",
		UpdaterPort: 9999,
	}

	result := config.UpdaterURL()
	expected := "http://172.16.4.97:9999/"

	if result != expected {
		t.Errorf("Unexpected UpdaterURL: expected:%s, got:%s", expected, result)
	}
}

func assertConfig(key string, value any, expected any) (err error) {
	if value != expected {
		err = fmt.Errorf("Unexpected %s value: expected:%v, got:%v", key, expected, value)
	}

	return
}
