package lastupload_test

import (
	"testing"

	"internal/env"

	"github.com/ArkhamCookie/crabot/lib/lastupload"
)

func TestGetLastUploadData(t *testing.T) {
	YOUTUBE_KEY, err := env.GetEnvValue("YOUTUBE_KEY", "./../../.env")
	if err != nil {
		t.Fatal(err)
	}

	_, err = lastupload.GetLastUploadData("UUi_0J4Zm1qwiYVSGRsA0_Bg", YOUTUBE_KEY)
	if err != nil {
		t.Fatal(err)
	}
}
