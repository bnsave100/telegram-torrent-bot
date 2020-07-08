package qbittorrent

import "fmt"

type TorrentList []TorrentInfo

func (l TorrentList) ToString() (str string) {
	for _, t := range l {
		str = fmt.Sprintf("%s%s\n", str, t.ToString())
	}

	return
}

type TorrentInfo struct {
	Name     string  `json:"name"`
	Progress float64 `json:"progress"`
}

func (ti TorrentInfo) ToString() string {
	return fmt.Sprintf("%s - %.2f%%", ti.Name, ti.Progress*100)
}
