package scraper

import (
	"context"
	model "go-discord-bot/types"
	tools "go-discord-bot/utilitys"

	"log"
	"os"
	"time"

	"github.com/chromedp/chromedp"
)

func InicScraper(link string) []model.Developer {
	//user y pass
	user := os.Getenv("USER")
	pass := os.Getenv("PASS")

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true), //permite ver lo que está haciendo el navegador
		chromedp.WindowSize(1920, 1080),
		chromedp.Flag("disable-gpu", false),
	)
	// Crear allocator con las opciones
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	// Crear contexto de Chrome usando el allocator
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()
	// Timeout opcional
	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	var list []string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://masterkey-hq.monday.com/auth/login_monday/email_password`),
		chromedp.WaitVisible(`#user_email`, chromedp.ByID),
		chromedp.WaitVisible(`#user_password`, chromedp.ByID),
		chromedp.SendKeys(`#user_email`, user, chromedp.ByID),
		chromedp.SendKeys(`#user_password`, pass, chromedp.ByID),
		Mondayfilter(&list, link),
	)
	if err != nil {
		log.Fatal(err)
	}

	//instancia de modelo
	joblist := []model.Developer{}
	//acá separamos la lista para dejarla más ordenada
	listoflist := tools.Separator(list)
	//acá vamos creando las distintas instancias
	for _, jobs := range listoflist {
		dev := model.NewDeveloper(jobs, link)
		joblist = append(joblist, dev)
	}

	return joblist
}
