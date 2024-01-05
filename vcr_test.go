package rockset_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/seborama/govcr/v14"
	"github.com/seborama/govcr/v14/cassette/track"
	"github.com/stretchr/testify/require"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/internal/test"
)

func vcrTestClient(t *testing.T, name string) (*rockset.RockClient, func(string) string) {
	rc, fn, err := vcrClient(name)
	if err != nil {
		require.NoError(t, err)
	}

	return rc, fn
}

func vcrClient(name string) (*rockset.RockClient, func(string) string, error) {
	return vcrClientWrapper(strings.ToLower(os.Getenv("VCR_MODE")), name)
}

func vcrClientWrapper(mode, name string) (*rockset.RockClient, func(string) string, error) {
	randFn := test.RandomName
	options := []rockset.RockOption{
		rockset.WithUserAgent("rockset-go-integration-tests"),
		rockset.WithCustomHeader("x-rockset-test", "go-client"),
	}
	var settings []govcr.Setting
	path := fmt.Sprintf("./testing_assets/cassettes/%s/%s.cassette.gz", rockset.Version, name)

	switch mode {
	case "record":
		randFn = func(pfx string) string {
			return fmt.Sprintf("go_%s", pfx)
		}
		settings = vcrSettings(false)
	case "offline", "": // quick run using a recorded response
		randFn = func(pfx string) string {
			return fmt.Sprintf("go_%s", pfx)
		}
		settings = vcrSettings(true)
		options = append(options, rockset.WithAPIKey("fake"),
			rockset.WithAPIServer("fake"), rockset.WithRetry(&test.Retrier{}))
	case "online": // for running everything live
		settings = append(vcrSettings(false), govcr.WithLiveOnlyMode(), govcr.WithReadOnlyMode())
	case "disabled":
		rc, err := rockset.NewClient(options...)
		return rc, randFn, err
	default:
		return nil, nil, fmt.Errorf("unknown VCR_MODE: %s", mode)
	}

	vcr := govcr.NewVCR(govcr.NewCassetteLoader(path), settings...)
	options = append(options, rockset.WithHTTPClient(vcr.HTTPClient()),
		rockset.WithCustomHeader("x-rockset-test", "go-client"))

	rc, err := rockset.NewClient(options...)
	return rc, randFn, err
}

// stripPatchVersion removes the patch version from a SemVer string
func stripPatchVersion(v string) string {
	fields := strings.Split(v, ".")
	if len(fields) != 3 {
		panic("malformed version string: " + v)
	}

	return strings.Join(fields[:2], ".")
}

// VCR settings that exclude the HTTP header Authorization and ignores the patch version
func vcrSettings(offline bool) []govcr.Setting {
	const authHeader = "Authorization"
	settings := []govcr.Setting{
		govcr.WithRequestMatchers(
			func(httpRequest, trackRequest *track.Request) bool {
				httpRequest.Header.Del(authHeader)
				trackRequest.Header.Del(authHeader)

				v := httpRequest.Header.Get(rockset.HeaderVersionName)
				v = stripPatchVersion(v)
				httpRequest.Header.Set(rockset.HeaderVersionName, v)

				v = trackRequest.Header.Get(rockset.HeaderVersionName)
				v = stripPatchVersion(v)
				trackRequest.Header.Set(rockset.HeaderVersionName, v)

				return govcr.DefaultHeaderMatcher(httpRequest, trackRequest)
			},
		),
		govcr.WithTrackRecordingMutators(track.TrackRequestDeleteHeaderKeys(authHeader)),
		govcr.WithTrackRecordingMutators(track.ResponseDeleteTLS()),
		govcr.WithTrackReplayingMutators(track.TrackRequestDeleteHeaderKeys(authHeader)),
		govcr.WithTrackReplayingMutators(track.ResponseDeleteTLS()),
	}
	if offline {
		settings = append(settings, govcr.WithOfflineMode())
	}

	return settings
}
