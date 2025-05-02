package lastupload_test

import (
	"internal/env"
	"testing"

	"crabot/lastupload"
)

func TestGet(t *testing.T) {
	YOUTUBE_KEY, err := env.GetEnvValue("YOUTUBE_KEY", "./../../.env")
	if err != nil {
		t.Fatal(err)
	}

	_, err = lastupload.GetLastUploadData("UUi_0J4Zm1qwiYVSGRsA0_Bg", YOUTUBE_KEY)
	if err != nil {
		t.Fatal(err)
	}
}
