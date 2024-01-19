package getwikidata

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type PageViewFile struct {
	Filename string
	URL      string
}

type PageViewDump struct {
	Files []PageViewFile
}

func newPageViewDump(url string) (*PageViewDump, error) {
	var pageViewDump PageViewDump

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %w", err)
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "<a href=\"pageviews") {
			start := strings.Index(line, "\"") + 1
			end := strings.Index(line[start:], "\"") + start
			href := line[start:end]
			pageViewDump.Files = append(pageViewDump.Files, PageViewFile{
				Filename: href,
				URL:      url + href,
			})
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	return &pageViewDump, nil
}

func NewPageViewDumpWithYearMonth(date string) (*PageViewDump, error) {
	year := date[:4]
	url := fmt.Sprintf("https://dumps.wikimedia.org/other/pageview_complete/%s/%s-%s/", year, year, date[4:6])
	return newPageViewDump(url)
}

func NewLatestPageViewDump() (*PageViewDump, error) {
	now := time.Now()
	yearMonth := fmt.Sprintf("%d%02d", now.Year(), now.Month())

	return NewPageViewDumpWithYearMonth(yearMonth)
}

func (p *PageViewDump) DownloadFiles(path string) error {
	for _, file := range p.Files {
		err := downloadFile(file.URL, path, file.Filename)
		if err != nil {
			return err
		}
	}
	return nil
}
