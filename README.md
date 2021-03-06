[![stable](https://img.shields.io/badge/stable-stable-green.svg)](https://github.com/phachon/retry/) 
[![build](https://img.shields.io/shippable/5444c5ecb904a4b21567b0ff.svg)](https://travis-ci.org/phachon/retry)
[![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/phachon/retry)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/phachon/retry/master/LICENSE)
[![go_Report](https://goreportcard.com/badge/github.com/phachon/retry)](https://goreportcard.com/report/github.com/phachon/retry)
[![platforms](https://img.shields.io/badge/platform-All-yellow.svg?style=flat)]()

Have a good retry.

[中文文档](README_CN.md)

## Feature
- Support setting retry times
- Support setting retry interval
- Support setting retry timeout

## Install
```bash
go get github.com/phachon/retry
```

## Used
```go
import github.com/phachon/retry
```

### Simple
```go
url := "https://api.github.com/"
req, _ := http.NewRequest("GET", url, nil)
var res *http.Response

err := retry.Retry(func() (err error) {
    res, err = http.DefaultClient.Do(req)
    if err != nil {
        return err
    }
    return nil
}, retry.SetRetryCount(2))
if err != nil {
    painc(err)
}
body, _ := ioutil.ReadAll(res.Body)
defer res.Body.Close()

fmt.Println(string(body))
```

### Retry interval
```go
url := "https://api.github.com/"
req, _ := http.NewRequest("GET", url, nil)
var res *http.Response

err := retry.Retry(func() (err error) {
    res, err = http.DefaultClient.Do(req)
    if err != nil {
        return err
    }
    return nil
}, retry.SetRetryCount(3), retry.SetIntervalTime(3 * time.Second))
if err != nil {
    panic(err)
}

body, _ := ioutil.ReadAll(res.Body)
defer res.Body.Close()

fmt.Println(string(body))
```

### Retry timeout
```go
url := "https://api.github.com/"
req, _ := http.NewRequest("GET", url, nil)
var res *http.Response

err := retry.Retry(func() (err error) {
    res, err = http.DefaultClient.Do(req)
    if err != nil {
        return err
    }
    return nil
}, retry.SetRetryCount(3), retry.SetRetryTimeout(800 * time.Millisecond))
if err != nil {
    panic(err)
    return
}

body, _ := ioutil.ReadAll(res.Body)
defer res.Body.Close()

fmt.Println(string(body))
```

## Feedback

- If you like the project, please [Star](https://github.com/phachon/retry/stargazers).
- If you have any problems in the process of use, welcome submit [Issue](https://github.com/phachon/retry/issues).
- If you find and solve bug, welcome submit [Pull Request](https://github.com/phachon/retry/pulls).
- If you want to redevelop, welcome [Fork](https://github.com/phachon/retry/network/members).
- If you want to make a friend, welcome send email to [phachon@163.com](mailto:phachon@163.com).

## License

MIT

Thanks
---------
Create By phachon@163.com