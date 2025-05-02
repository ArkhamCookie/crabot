package lastupload_test

import (
	"crabot/lastupload"
	"internal/env"
	"testing"
)

func TestGetLastUpload(t *testing.T) {
	YOUTUBE_KEY, err := env.GetEnvValue("YOUTUBE_KEY", "./../../.env")
	if err != nil {
		t.Fatal(err)
	}

	_, err = lastupload.GetLastUpload("spiritvoices", YOUTUBE_KEY)
	if err != nil {
		t.Fatal(err)
	}
}
