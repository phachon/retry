[![stable](https://img.shields.io/badge/stable-stable-green.svg)](https://github.com/phachon/retry/) 
[![build](https://img.shields.io/shippable/5444c5ecb904a4b21567b0ff.svg)](https://travis-ci.org/phachon/retry)
[![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/phachon/retry)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/phachon/retry/master/LICENSE)
[![go_Report](https://goreportcard.com/badge/github.com/phachon/retry)](https://goreportcard.com/report/github.com/phachon/retry)
[![platforms](https://img.shields.io/badge/platform-All-yellow.svg?style=flat)]()

一个重试工具包

[English](README.md)


## 功能
- 支持设置重试次数
- 支持重试间隔时间
- 支持重试超时时间

## 安装
```bash
go get code.byted.org/motor/gopkg/retry
```

## 使用

```go
import code.byted.org/motor/gopkg/retry
```

### 简单例子
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

### 重试间隔
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

### 重试超时
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

## 反馈

- 如果您喜欢该项目，请 [Star](https://github.com/phachon/wmqx/stargazers).
- 如果在使用过程中有任何问题， 请提交 [Issue](https://github.com/phachon/wmqx/issues).
- 如果您发现并解决了bug，请提交 [Pull Request](https://github.com/phachon/wmqx/pulls).
- 如果您想二次开发，欢迎 [Fork](https://github.com/phachon/wmqx/network/members).
- 如果你想交个朋友，欢迎发邮件给 [phachon@163.com](mailto:phachon@163.com).

## License

MIT

Thanks
---------
Create By phachon@163.com
