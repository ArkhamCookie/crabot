package lastupload_test

import (
	"crabot/lastupload"
	"internal/env"
	"testing"
)

func TestGetLastUploadData(t *testing.T) {
	YOUTUBE_KEY, err := env.GetEnvValue("YOUTUBE_KEY", "./../../.env")
	if err != nil {
		t.Fatal(err)
	}

	_, err = lastupload.GetLastUploadData("spiritvoices", YOUTUBE_KEY)
	if err != nil {
		t.Fatal(err)
	}
}
