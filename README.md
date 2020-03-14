LINE: LIFF template in Go
==============

 [![GoDoc](https://godoc.org/github.com/kkdai/liff-template-go.svg?status.svg)](https://godoc.org/github.com/kkdai/liff-template-go)[![Build Status](https://travis-ci.org/kkdai/liff-template-go.svg?branch=master)](https://travis-ci.org/kkdai/liff-template-go)[![goreportcard.com](https://goreportcard.com/badge/github.com/kkdai/liff-template-go)](https://goreportcard.com/report/github.com/kkdai/liff-template-go)

This is sample code to demostration LINE LIFF using Go HTML. For detail LIFF check [API reference](https://developers.line.biz/en/reference/liff/) here.



Installation
=============

- Create your LINE Developer account and create a developer trial account.  
  [![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

- Deploy this project to heroku
    - You need to fill following configuration during you create project.
        - YOUR_LIFF_ID
        - YOUR_REDIRECT_URL
- Done


FAQ
=============

1. How to get my LIFF app ID?

  - Remember to add LIFF App in LINE Login channel, refer to this [announcement](https://developers.line.biz/zh-hant/news/2020/02/05/liff-channel-type/).
  - Go [LINE Developer console](https://developers.line.biz/console/) create a LINE Login channel.

2. Why it show "Error: user doesn't grant required permissions yet" when I call `liff.sendMessages()`?

  - The root cause might be:
    - Your user login session timeout, please run `liff.login()` again.
    - Your channel is not using LINE Login, your user could not complete login process.

