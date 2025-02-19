package zhttp

import (
	"io"
	"net/http"
	"strings"
)

type HttpConfig struct {
	Payload  *strings.Reader
	Headers  map[string]string
	Cookies  map[string]string
	Redirect bool
}

func HttpGetBody(url string, httpConfig HttpConfig) ([]byte, error) {
	res, err := HttpGet(url, httpConfig)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func HttpPostBody(url string, httpConfig HttpConfig) ([]byte, error) {
	res, err := HttpPost(url, httpConfig)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func HttpPatchBody(url string, httpConfig HttpConfig) ([]byte, error) {
	res, err := HttpPatch(url, httpConfig)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// HttpGet 发起Get请求
func HttpGet(url string, httpConfig HttpConfig) (*http.Response, error) {
	method := "GET"

	client := &http.Client{}

	// 设置 CheckRedirect 回调，禁止自动重定向
	if !httpConfig.Redirect {
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			// 返回 http.ErrUseLastResponse 表示停止重定向并返回最后的响应
			return http.ErrUseLastResponse
		}
	}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	headers := httpConfig.Headers
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}

	cookies := httpConfig.Cookies
	if cookies != nil {
		for key, val := range cookies {
			req.AddCookie(&http.Cookie{Name: key, Value: val})
		}
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// HttpPost 发起Post请求
func HttpPost(url string, httpConfig HttpConfig) (*http.Response, error) {
	method := "POST"

	client := &http.Client{}

	// 设置 CheckRedirect 回调，禁止自动重定向
	if httpConfig.Redirect {
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			// 返回 http.ErrUseLastResponse 表示停止重定向并返回最后的响应
			return http.ErrUseLastResponse
		}
	}

	payload := httpConfig.Payload
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}

	headers := httpConfig.Headers
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}

	cookies := httpConfig.Cookies
	if cookies != nil {
		for key, val := range cookies {
			req.AddCookie(&http.Cookie{Name: key, Value: val})
		}
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// HttpPatch 发起Patch请求
func HttpPatch(url string, httpConfig HttpConfig) (*http.Response, error) {
	method := "PATCH"

	client := &http.Client{}

	// 设置 CheckRedirect 回调，禁止自动重定向
	if httpConfig.Redirect {
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			// 返回 http.ErrUseLastResponse 表示停止重定向并返回最后的响应
			return http.ErrUseLastResponse
		}
	}

	payload := httpConfig.Payload
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}

	headers := httpConfig.Headers
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}

	cookies := httpConfig.Cookies
	if cookies != nil {
		for key, val := range cookies {
			req.AddCookie(&http.Cookie{Name: key, Value: val})
		}
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
