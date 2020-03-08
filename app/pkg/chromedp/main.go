package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"log"
	"time"
)

const host = "https://www.facebook.com"

func GetCookies() map[string]string {
	str := `{"datr":"uugYXfeMFyXAvGdPt7C7PsWF","_fbp":"fb.1.1583028294666.428290832","locale":"vi_VN","m_pixel_ratio":"2","c_user":"100000027929102","xs":"26%3AgAO8EiO0Aa8H7g%3A2%3A1583562998%3A8037%3A6176","spin":"r.1001807630_b.trunk_t.1583677611_s.1_v.2_","fr":"14htpvUUenz3SnLaq.AWXrBnh38jvvbxNp6_UKCn8HzjA.BePiAo.dd.F5i.0.0.BeZQ1x.AWU0E4Gl","presence":"EDvF3EtimeF1583682318EuserFA21BB27929102A2EstateFDt3F_5bDiFA2user_3a1B04608994306A2ErF1EoF1EfF4C_5dElm3FA2user_3a1B04608994306A2Eutc3F1583682317580G583682318352CEchF_7bCC","act":"1583683392195%2F8","wd":"323x616"}`
	var rs map[string]string
	err := json.Unmarshal([]byte(str), &rs)
	if err != nil {
		panic(err)
	}
	return rs
}

func GetUrl(url string) chromedp.Tasks {
	tasks := chromedp.Tasks{
		chromedp.ActionFunc(func(ctx context.Context) error {
			// create cookie expiration
			expr := cdp.TimeSinceEpoch(time.Now().Add(180 * 24 * time.Hour))
			// add cookies to chrome
			for k, v := range GetCookies() {
				success, err := network.SetCookie(k, v).
					WithExpires(&expr).
					WithDomain("facebook.com").
					WithHTTPOnly(false).
					Do(ctx)
				if err != nil {
					return err
				}
				if !success {
					return fmt.Errorf("could not set cookie %q to %q", k, v)
				}
			}

			return nil
		}),
		// navigate to site
		chromedp.Navigate(url),
	}

	return tasks
}

func main() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	err := chromedp.Run(ctx, GetUrl("https://www.facebook.com/groups/"))
	if err != nil {
		log.Fatal(err)
	}
	var nodes []*cdp.Node
	//if err := chromedp.Run(ctx, chromedp.WaitVisible("._2yau")); err != nil {
	//	log.Fatal(fmt.Errorf("could not get section: %v", err))
	//}
	if err := chromedp.Run(ctx, chromedp.Nodes("._2yau", &nodes)); err != nil {
		log.Fatal(fmt.Errorf("could not get section: %v", err))
	}
	fmt.Println(nodes)
}
