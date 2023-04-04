package gitignoreio

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type GitignoreConfig struct {
	Name     string
	Key      string
	Contents string
	FileName string
}

type GitignoreIo struct {
	baseUrl    string
	httpClient *http.Client
	list       []string
}

func NewClient() *GitignoreIo {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	return &GitignoreIo{
		baseUrl:    "https://www.toptal.com/developers/gitignore/api",
		httpClient: client,
	}
}

func (c *GitignoreIo) List() (resp []string, err error) {
	endpoint := c.baseUrl + "/list"
	res, err := c.do(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	ignoreList := parseIgnoreList(res.Body)

	if len(ignoreList) > 0 {
		return ignoreList, nil
	}
	return nil, errors.New("unable to parse list from gitignore.io")
}

func (c *GitignoreIo) FetchAll() (resp map[string]GitignoreConfig, err error) {
	endpoint := c.baseUrl + "/list?format=json"
	res, err := c.do(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var configs map[string]GitignoreConfig

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(body, &configs); err != nil {
		return nil, err
	}

	return configs, nil
}

func (c *GitignoreIo) FetchConfig(names []string) (resp string, err error) {
	endpoint := c.baseUrl + strings.Join(names, ",")
	res, err := c.do(http.MethodGet, endpoint, nil)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (c *GitignoreIo) do(method, endpoint string, params map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	q := req.URL.Query()
	for key, val := range params {
		q.Set(key, val)
	}
	req.URL.RawQuery = q.Encode()
	return c.httpClient.Do(req)
}

func parseIgnoreList(buf io.Reader) []string {
	var ignoreList []string

	scanner := bufio.NewScanner(buf)
	for scanner.Scan() {
		line := scanner.Text()
		names := strings.Split(line, ",")
		ignoreList = append(ignoreList, names...)
	}

	return ignoreList
}
