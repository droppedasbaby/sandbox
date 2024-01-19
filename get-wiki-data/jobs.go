package getwikidata

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"time"
)

type FileDetail struct {
	URL  string `json:"url"`
	MD5  string `json:"md5"`
	SHA1 string `json:"sha1"`
	Size int64  `json:"size"`
}

type DumpStatus struct {
	Files   map[string]FileDetail `json:"files"`
	Status  string                `json:"status"`
	Updated string                `json:"updated"`
}

type JobStatus struct {
	Jobs map[string]DumpStatus `json:"jobs"`
}

func jobStatusUrl(date string) string {
	return "https://dumps.wikimedia.org/enwiki/" + date + "/dumpstatus.json"
}

func newJobStatus(url string) (*JobStatus, error) {
	var jobStatus JobStatus

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	err = json.Unmarshal(body, &jobStatus)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	return &jobStatus, nil
}

func NewJobStatusWithDate(date string) (*JobStatus, error) {
	jobs, err := newJobStatus("https://dumps.wikimedia.org/enwiki/" + date + "/dumpstatus.json")
	return jobs, err
}

func NewLatestJobStatus() (*JobStatus, error) {
	now := time.Now()
	monthStart := fmt.Sprintf("%d%02d01", now.Year(), now.Month())
	return NewJobStatusWithDate(monthStart)
}

func (j *JobStatus) GetCompletedJobs() []string {
	completed := []string{}

	for job, status := range j.Jobs {
		if status.Status == "done" {
			completed = append(completed, job)
		}
	}
	sort.Strings(completed)
	return completed
}

func (j *JobStatus) GetFilesForJob(job string) []string {
	files := []string{}
	for file := range j.Jobs[job].Files {
		files = append(files, file)
	}
	sort.Strings(files)
	return files
}

func (j *JobStatus) DownloadFilesForJob(job string, path string) error {
	if j.Jobs[job].Status != "done" {
		return fmt.Errorf("job %s is not complete", job)
	}

	for fileName, fileDetail := range j.Jobs[job].Files {
		fmt.Printf("Downloading %s\n", fileName)

		err := downloadFile("https://dumps.wikimedia.org"+fileDetail.URL, path, fileName)
		if err != nil {
			fmt.Printf("Failed to download %s: %s\n", fileName, err)
			return err
		}
	}
	return nil
}
