package fileinfo

import (
	"testing"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		input    string
		expected MediaInfo
	}{
		{
			input: "Harry.Potter.and.the.Deathly.Hallows.Part.II.2011.1080p.BluRay.x264.DTS-WiKi.mkv",
			expected: MediaInfo{
				Title:      "Harry Potter and the Deathly Hallows Part II",
				Year:       2011,
				Source:     "BluRay",
				Resolution: "1080p",
				VideoCodec: "x264",
				AudioCodec: "DTS",
				Studio:     "WiKi",
				Type:       "mkv",
			},
		},
		{
			input: "Inception.2010.720p.BluRay.x264-SPARKS.mkv",
			expected: MediaInfo{
				Title:      "Inception",
				Year:       2010,
				Source:     "BluRay",
				Resolution: "720p",
				VideoCodec: "x264",
				AudioCodec: "",
				Studio:     "SPARKS",
				Type:       "mkv",
			},
		},
		{
			input: "Avengers.Endgame.2019.2160p.UHD.BluRay.x265-TERMiNAL.avi",
			expected: MediaInfo{
				Title:      "Avengers Endgame",
				Year:       2019,
				Source:     "BluRay",
				Resolution: "2160p",
				VideoCodec: "x265",
				AudioCodec: "",
				Studio:     "TERMiNAL",
				Type:       "avi",
			},
		},
		{
			input: "The.Mandalorian.S01E05.1080p.WEB-DL.DDP5.1.H.265-NTb.mp4",
			expected: MediaInfo{
				Title:      "The Mandalorian",
				Year:       0,
				Source:     "WEB-DL",
				Resolution: "1080p",
				VideoCodec: "H.265",
				AudioCodec: "DDP5.1",
				Studio:     "NTb",
				Type:       "mp4",
				Season:     "S01",
				Episode:    "E05",
			},
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := Parse(tc.input)
			if result.Title != tc.expected.Title {
				t.Errorf("Expected Title: %s, Got: %s", tc.expected.Title, result.Title)
			}
			if result.Year != tc.expected.Year {
				t.Errorf("Expected Year: %d, Got: %d", tc.expected.Year, result.Year)
			}
			if result.Source != tc.expected.Source {
				t.Errorf("Expected Source: %s, Got: %s", tc.expected.Source, result.Source)
			}
			if result.Resolution != tc.expected.Resolution {
				t.Errorf("Expected Resolution: %s, Got: %s", tc.expected.Resolution, result.Resolution)
			}
			if result.VideoCodec != tc.expected.VideoCodec {
				t.Errorf("Expected VideoCodec: %s, Got: %s", tc.expected.VideoCodec, result.VideoCodec)
			}
			if result.AudioCodec != tc.expected.AudioCodec {
				t.Errorf("Expected AudioCodec: %s, Got: %s", tc.expected.AudioCodec, result.AudioCodec)
			}
			if result.Type != tc.expected.Type {
				t.Errorf("Expected Type: %s, Got: %s", tc.expected.Type, result.Type)
			}
			// if result.Studio != tc.expected.Studio {
			// 	t.Errorf("Expected Studio: %s, Got: %s", tc.expected.Studio, result.Studio)
			// }
			if result.Channel != tc.expected.Channel {
				t.Errorf("Expected Channel: %s, Got: %s", tc.expected.Channel, result.Channel)
			}
			if result.Version != tc.expected.Version {
				t.Errorf("Expected Version: %s, Got: %s", tc.expected.Version, result.Version)
			}

			if result.Season != tc.expected.Season {
				t.Errorf("Expected Season: %s, Got: %s", tc.expected.Season, result.Season)
			}
			if result.Episode != tc.expected.Episode {
				t.Errorf("Expected Episode: %s, Got: %s", tc.expected.Episode, result.Episode)
			}

		})
	}
}
