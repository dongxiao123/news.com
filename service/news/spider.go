package news

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego/orm"
	"github.com/chromedp/chromedp"
	"news.com/models"
	"news.com/utils"
	"strings"
	"time"
)

func Spider() (titleMap []*models.Title, err error) {
	o := orm.NewOrm()
	var titles []*models.Title
	num, err := o.QueryTable("title").Filter("has_spidered", "0").All(&titles)
	if err != nil {
		utils.Logs.Warning("QueryTable error:", err)
		return
	}
	if num < 1 {
		utils.Logs.Warning("no data ")
		return
	}
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("hide-scrollbars", false),
		chromedp.Flag("mute-audio", false),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36`),
	}
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)
	for _, t := range titles {
		t.HasSpidered = 1
		_, err := o.Update(t)
		if err != nil {
			utils.Logs.Warning("Update error :", err)
		}

		//创建chrome窗口
		allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), options...)
		defer cancel()
		ctx, cancel := chromedp.NewContext(allocCtx)
		defer cancel()
		var html string
		if err := chromedp.Run(ctx,
			chromedp.Navigate(t.Url),
			chromedp.WaitVisible(`div`, chromedp.ByQuery),
			chromedp.InnerHTML(`body`, &html, chromedp.NodeVisible, chromedp.ByQuery),
		)
			err != nil {
			utils.Logs.Warning(err.Error())
		}
		utils.Logs.Warning("html_html",html)
		dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
		if err != nil {
			utils.Logs.Warning(err.Error())
		}
		var titleMap []models.Title
		dom.Find("a").Each(func(i int, selection *goquery.Selection) {
			title := selection.Text()
			if !checkExcludeTitle(title) {
				return
			}
			url, _ := selection.Attr("href")
			if !checkExcludeUrl(url) {
				return
			}
			t1 := models.Title{
				Title:        title,
				Code:         Code,
				Url:          url,
				Md5CodeTitle: utils.Md5V(Code + title),
				CreatedAt:    time.Now().Format("2006-01-02 15:04:05"),
				UpdatedAt:    time.Now().Format("2006-01-02 15:04:05"),
			}
			titleMap = append(titleMap, t1)
		})
	}
	return
}
