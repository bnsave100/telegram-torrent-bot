package qbittorrent

import (
	"errors"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

type QBittorrent struct {
	username string
	password string

	client  *http.Client
	baseUrl string
}

func (qb *QBittorrent) auth() {
	client := &http.Client{}

	resp, err := client.PostForm(qb.baseUrl+"auth/login", url.Values{
		"username": {qb.username},
		"password": {qb.password},
	})

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	sid := between(resp.Header.Get("Set-Cookie"), "SID=", ";")

	if sid == "" {
		log.Fatal("qBittorrent auth error")
	}

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

func (qb *QBittorrent) Add(urls []string) (err error) {
	v := url.Values{}

	for _, u := range urls {
		if u == "" {
			continue
		}
		v.Add("urls", u)
	}

	res, err := qb.client.PostForm(qb.baseUrl+"torrents/add", v)

	if err != nil {
		return
	}

	if res.StatusCode != 200 {
		return errors.New("something went wrong")
	}

	return
}

func NewQBittorrent(login string, password string, baseUrl string) *QBittorrent {
	qb := &QBittorrent{
		username: login,
		password: password,
		baseUrl:  baseUrl + "/api/v2/",
	}

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
