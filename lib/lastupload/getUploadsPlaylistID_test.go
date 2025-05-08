package lastupload_test

import (
	"testing"

	"internal/env"

	"github.com/ArkhamCookie/crabot/lib/lastupload"
)

func TestGetUploadsPlaylistID(t *testing.T) {
	want := "UUi_0J4Zm1qwiYVSGRsA0_Bg"

	YOUTUBE_KEY, err := env.GetEnvValue("YOUTUBE_KEY", "../../.env")
	if err != nil {
		t.Fatal(err)
	}

	output, err := lastupload.GetUploadsPlaylistID("UCi_0J4Zm1qwiYVSGRsA0_Bg", YOUTUBE_KEY)
	if err != nil {
		t.Fatal(err)
	}

	if want != output {
		t.Fatalf("`lastupload.GetUploadsPlaylistID` = %q, wanted %q", output, want)
	}
}
