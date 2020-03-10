package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
	"log"
	"time"
)

const host = "https://www.facebook.com"

func GetCookies() map[string]string {
	str := `{"datr":"r7S_XZXAAhVOKS4rRnFe_cDA","c_user":"100000027929102","xs":"16%3AvrW_yOAMkLM-9A%3A2%3A1583136849%3A8037%3A6176","_fbp":"fb.1.1583399593262.720038549","spin":"r.1001807870_b.trunk_t.1583723572_s.1_v.2_","m_pixel_ratio":"1","x-referer":"eyJyIjoiL2hvbWUucGhwIiwiaCI6Ii9ob21lLnBocCIsInMiOiJtIn0%3D","wd":"1414x782","fr":"0JjvCGuFMFHambslt.AWVkxtC5mtc465fa_cimsl9DAbQ.Bdv6MD.VW.F5h.0.0.BeZu0G.AWXBCqfT","presence":"EDvF3EtimeF1583804325EuserFA21BB27929102A2EstateFDt3F_5bDiFA2user_3a1B04608994306A2EoF1EfF2C_5dEutc3F1583749830999G583749831362CEchF_7bCC","act":"1583804371104%2F2"}`
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
		chromedp.ActionFunc(func(context.Context) error {
			log.Println("Access URL: " + url)
			return nil
		}),
	}
	return tasks
}

func main() {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()
	c, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var nodes []*cdp.Node
	chromedp.Run(c, chromedp.Tasks{
		chromedp.ActionFunc(func(ctx context.Context) error {
			// create cookie expiration
			expr := cdp.TimeSinceEpoch(time.Now().Add(180 * 24 * time.Hour))
			// add cookies to chrome
			for k, v := range GetCookies() {
				success, err := network.SetCookie(k, v).
					WithExpires(&expr).
					WithDomain(".facebook.com").
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
		chromedp.Navigate("https://m.facebook.com/groups/"),
		chromedp.ActionFunc(func(context.Context) error {
			//log.Println("Access URL: " + url)
			return nil
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, exp, err := runtime.Evaluate(waitLoaded).Do(ctx)
			if err != nil {
				return err
			}
			if exp != nil {
				return exp
			}
			log.Println("Page Ready!")
			return nil
		}),
		chromedp.Nodes(`a._7hkg`, &nodes, chromedp.ByQueryAll),
		chromedp.ActionFunc(func(ctx context.Context) error {
			log.Println(nodes)
			return nil
		}),

	})

}

const waitLoaded  = `(function(){
    var tId = setInterval(function() {
        if (document.readyState == "complete") onComplete()
    }, 11);
    function onComplete(){
        clearInterval(tId);    
        alert("loaded!");    
    };
})()`