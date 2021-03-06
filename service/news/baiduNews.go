package news

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"news.com/models"
	"news.com/utils"
	"strings"
	"time"
)

const Code = "BaiduNews"

type BaiduNews struct {
	Code string
}

func NewBaiduNews() BaiduNews {
	return BaiduNews{
		Code: Code,
	}
}

//获取百度标题数据
func (n BaiduNews) GetTitleData() []models.Title {
	//增加选项，允许chrome窗口显示出来
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("hide-scrollbars", false),
		chromedp.Flag("mute-audio", false),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36`),
	}
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)
	//创建chrome窗口
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), options...)
	defer cancel()
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()
	var html string
	if err := chromedp.Run(ctx,
		chromedp.Navigate(`http://news.baidu.com/`),
		chromedp.WaitVisible(`#pane-news`, chromedp.ByID),
		chromedp.InnerHTML(`#body`, &html, chromedp.NodeVisible, chromedp.ByID),
	)
		err != nil {
		utils.Logs.Warning(err.Error())
	}
	utils.Logs.Warning("html_html", html)
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		utils.Logs.Warning(err.Error())
	}
	var titleMap []models.Title
	dom.Find("a").Each(func(i int, selection *goquery.Selection) {

		title := strings.Trim(selection.Text(), "")
		if !checkExcludeTitle(title) {
			return
		}
		url, _ := selection.Attr("href")
		url = strings.Trim(url, "")
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
	return titleMap
}
func checkExcludeTitle(title string) bool {
	if len(title) < 2 {
		return false
	}
	return true
}
func checkExcludeUrl(url string) bool {
	if url == "javascript:void(0);" {
		return false
	}
	if len(url) < 6 {
		return false
	}
	return true
}
