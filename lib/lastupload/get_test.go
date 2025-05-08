package lastupload_test

import (
	"internal/env"
	"testing"

	"github.com/ArkhamCookie/crabot/lib/lastupload"
)

func TestGetData(t *testing.T) {
	YOUTUBE_KEY, err := env.GetEnvValue("YOUTUBE_KEY", "./../../.env")
	if err != nil {
		t.Fatal(err)
	}

	_, err = lastupload.GetData("spiritvoices", YOUTUBE_KEY)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetDate(t *testing.T) {
	YOUTUBE_KEY, err := env.GetEnvValue("YOUTUBE_KEY", "./../../.env")
	if err != nil {
		t.Fatal(err)
	}

	lastUploadData, err := lastupload.GetData("spiritvoices", YOUTUBE_KEY)
	if err != nil {
		t.Fatal(err)
	}

	_, err = lastupload.GetDate(lastUploadData)
	if err != nil {
		t.Fatal(err)
	}
}
