package gakujo

import (
	"bytes"
	"io"
	"log"
	"net/url"
	"time"

	"test/model"

	"github.com/PuerkitoBio/goquery"
	"github.com/szpp-dev-team/gakujo-api/util"
)

func (c *Client) getTask() []model.TaskRow {

	datas := make(url.Values)
	datas.Set("headTitle", "ホーム")
	datas.Set("menuCode", "Z07") // TODO: menucode を定数化(まとめてやる)
	datas.Set("nextPath", "/home/home/initialize")

	urll := "https://gakujo.shizuoka.ac.jp/portal/common/generalPurpose/"

	resp, err := c.getPage(urll, datas)

	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp)

	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(io.NopCloser(bytes.NewBuffer(body)))

	if err != nil {
		log.Fatal(err)
	}

	var taskRows []model.TaskRow

	//taskRows = make([]TaskRow, 0)
	doc.Find("#tbl_submission > tbody > tr").EachWithBreak(func(i int, selection *goquery.Selection) bool {
		var inerr error
		//taskType := model.ToTasktype(selection.Find("td.arart > span > span").Text())
		taskType := selection.Find("td.arart > span > span").Text()
		deadlineText := selection.Find("td.daytime").Text()
		var deadline time.Time
		if deadlineText != "" {
			deadline, inerr = util.Parse2400("2006/01/02 15:04", deadlineText)
			if inerr != nil {
				err = inerr
				return false
			}
		}
		data := model.TaskRow{
			Type:     taskType,
			Deadline: deadline,
			Name:     selection.Find("td:nth-child(3) > a").Text(),
		}
		taskRows = append(taskRows, data)
		return true
	})

	return taskRows
}
