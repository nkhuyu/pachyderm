package scraper

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/pachyderm/pachyderm/src/client/pkg/require"
)

func TestScraper(t *testing.T) {
	wd, err := os.Getwd()
	require.NoError(t, err)
	fmt.Printf("wd: %s", wd)
	require.NoError(t, exec.Command("pachctl", "create-repo", "urls").Run())
	require.NoError(t, exec.Command("pachctl", "start-commit", "urls", "master").Run())
	putFileCmd := exec.Command("pachctl", "put-file", "urls", "master", "urls")
	urls, err := os.Open("urls")
	require.NoError(t, err)
	putFileCmd.Stdin = urls
	require.NoError(t, putFileCmd.Run())
	require.NoError(t, exec.Command("pachctl", "finish-commit", "urls", "master").Run())
	require.NoError(t, exec.Command("pachctl", "create-pipeline", "-f", "scraper.json").Run())
	require.NoError(t, exec.Command("pachctl", "flush-commit", "urls/master").Run())
}