package fetcher

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

// ExtractLinks fetches a URL and returns all found absolute links.
func ExtractLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Simple regex to find absolute URLs starting with http or https
	re := regexp.MustCompile(`href="(https?://[^"]+)"`)
	matches := re.FindAllStringSubmatch(string(body), -1)

	var links []string
	for _, m := range matches {
		links = append(links, m[1])
	}

	return links, nil
}
