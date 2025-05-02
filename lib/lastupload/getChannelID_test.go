package lastupload_test

import (
	"testing"

	"crabot/lastupload"
	"internal/env"
)

func TestGetChannelID(t *testing.T) {
	want := "UCi_0J4Zm1qwiYVSGRsA0_Bg"

	YOUTUBE_KEY, err := env.GetEnvValue("YOUTUBE_KEY", "./../../.env")
	if err != nil {
		t.Fatal(err)
	}

	output, err := lastupload.GetChannelID("spiritvoices", YOUTUBE_KEY)
	if err != nil {
		t.Fatal(err)
	}

	if want != output {
		t.Fatalf("`lastupload.GetChannelID` = %q, wanted %q", output, want)
	}
}
