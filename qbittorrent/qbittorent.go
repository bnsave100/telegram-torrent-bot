package qbittorrent

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

type qBittorrent struct {
	username string
	password string

	client  *http.Client
	baseUrl string
}

func (qb *qBittorrent) auth() {
	client := &http.Client{}

	resp, err := client.PostForm(qb.baseUrl+"auth/login", url.Values{
		"username": {qb.username},
		"password": {qb.password},
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	sid := between(resp.Header.Get("Set-Cookie"), "SID=", ";")
	sidCookie := &http.Cookie{
		Name:  "SID",
		Value: sid,
	}

	u, _ := url.Parse(qb.baseUrl)

	jar, _ := cookiejar.New(nil)
	jar.SetCookies(u, []*http.Cookie{sidCookie})

	client.Jar = jar

	qb.client = client
}

func NewQBittorrent(login string, password string, baseUrl string) *qBittorrent {
	qb := &qBittorrent{username: login, password: password}

	qb.baseUrl = baseUrl + "/api/v2/"
	qb.auth()

	return qb
}

func between(value string, a string, b string) string {
	posFirst := strings.Index(value, a)

	if posFirst == -1 {
		return ""
	}

	posLast := strings.Index(value, b)

	if posLast == -1 {
		return ""
	}

	posFirstAdjusted := posFirst + len(a)

	if posFirstAdjusted >= posLast {
		return ""
	}

	return value[posFirstAdjusted:posLast]
}
