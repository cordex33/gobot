package scraper

import (
	"time"

	"github.com/chromedp/chromedp"
)

func Mondayfilter(list *[]string, link string) chromedp.Tasks {

	return chromedp.Tasks{
		chromedp.Sleep(500 * time.Millisecond),
		chromedp.WaitVisible(`.next-button.submit_button.button_843ba7af12.sizeMedium_c0b6e12605.kindPrimary_e3d5c4d9dc.colorPrimary_00e85329cd`, chromedp.ByQuery),
		chromedp.Click(`.next-button.submit_button.button_843ba7af12.sizeMedium_c0b6e12605.kindPrimary_e3d5c4d9dc.colorPrimary_00e85329cd`, chromedp.ByQuery),
		chromedp.WaitVisible(`#main`, chromedp.ByID),
		//Este filtro por ahora es exclusivo para mi xdd
		chromedp.Navigate(link),
		chromedp.WaitVisible(`.table-items`, chromedp.ByQuery),

		//aca esperamos que el filtro de las columnas este visible y lo clickeamos para quitar el 'ID ANTIGUO, muy volatil para colocarlo'
		chromedp.WaitVisible(`.board-filter-item-component.columns-filter-component`, chromedp.ByQuery),
		chromedp.Sleep(1500 * time.Millisecond),
		chromedp.Click(`.board-filter-item-component.columns-filter-component`, chromedp.ByQuery),
		chromedp.Sleep(500 * time.Millisecond),
		//acá comienzan los click a las columnas que queremos quitar
		//id antiguo
		chromedp.Click(`div.columns-section-module_columnsWrapper__CAkGU[aria-labelledby="columns"] g[filter="url(#filter0_d_43912_44061)"]`, chromedp.ByQuery),
		//Archivo
		chromedp.Click(`div.columns-section-module_columnsWrapper__CAkGU[aria-labelledby="columns"] g[clip-path="url(#clip0_43912_44049)"][filter="url(#filter0_d_43912_44049)"]`, chromedp.ByQuery),
		//Plantillas
		chromedp.Click(`//div[@class="columns-section-module_columnsWrapper__CAkGU" and @aria-labelledby="columns"]//div[text()="Plantillas"]`, chromedp.BySearch),
		//Queue time
		chromedp.Click(`//div[@class="columns-section-module_columnsWrapper__CAkGU" and @aria-labelledby="columns"]//div[text()="Queue Time"]`, chromedp.BySearch),
		//Cycle time
		chromedp.Click(`//div[@class="columns-section-module_columnsWrapper__CAkGU" and @aria-labelledby="columns"]//div[text()="Cycle Time"]`, chromedp.BySearch),
		//Eficiencia de flujo
		chromedp.Click(`//div[@class="columns-section-module_columnsWrapper__CAkGU" and @aria-labelledby="columns"]//div[text()="Eficiencia de Flujo"]`, chromedp.BySearch),
		//Lead Time
		chromedp.Click(`//div[@class="columns-section-module_columnsWrapper__CAkGU" and @aria-labelledby="columns"]//div[text()="Lead Time"]`, chromedp.BySearch),
		//Estado Activo
		chromedp.Click(`//div[@class="columns-section-module_columnsWrapper__CAkGU" and @aria-labelledby="columns"]//div[text()="Estado Activo"]`, chromedp.BySearch),
		//En desarrollo
		chromedp.Click(`//div[@class="columns-section-module_columnsWrapper__CAkGU" and @aria-labelledby="columns"]//div[text()="EnDesarrollo"]`, chromedp.BySearch),
		//Plazo
		chromedp.Click(`//div[@class="columns-section-module_columnsWrapper__CAkGU" and @aria-labelledby="columns"]//div[text()="Plazo"]`, chromedp.BySearch),
		//Nivel Cliente
		chromedp.Click(`//div[@class="columns-section-module_columnsWrapper__CAkGU" and @aria-labelledby="columns"]//div[text()="Nivel Cliente"]`, chromedp.BySearch),
		//Estado OC
		chromedp.Click(`//div[@class="columns-section-module_columnsWrapper__CAkGU" and @aria-labelledby="columns"]//div[text()="Estado OC"]`, chromedp.BySearch),
		//Jefe Proyecto (JP)
		chromedp.Click(`//div[@class="columns-section-module_columnsWrapper__CAkGU" and @aria-labelledby="columns"]//div[text()="Jefe Proyecto (JP)"]`, chromedp.BySearch),
		//Lider Tecnico
		chromedp.Click(`//div[@class="columns-section-module_columnsWrapper__CAkGU" and @aria-labelledby="columns"]//div[text()="Lider Tecnico"]`, chromedp.BySearch),
		//ac esperamos que cargue
		chromedp.WaitReady(`div[role="group"]`, chromedp.ByQuery),
		//quitamos el filtro de ocultar
		chromedp.Click(`.board-filter-item-component.columns-filter-component`, chromedp.ByQuery),
		//con esta podemos hacer que termine de cargar el final de la tabla dinamica donde están los datos para que los tome
		chromedp.WaitReady(`div[data-testid="group-footer"]`, chromedp.ByQuery),
		//esto lo hacemos para que se vea el total de la página y cargue toda o la gran mayoría de la tabla dinamica hedionda
		chromedp.Evaluate(`document.body.style.zoom = "30%"`, nil),
		chromedp.Sleep(2500 * time.Millisecond),

		//esté queryselector fue casi un 40% hecho por gpt
		chromedp.EvaluateAsDevTools(`
		Array.from(
			document.querySelectorAll('div[role="group"] > div:nth-child(n+2):nth-last-child(n+3)')
		)
			.map(div => {
			// Buscar todos los elementos de texto dentro de este div
			let texts = Array.from(div.querySelectorAll('[data-testid="text"], span.ds-text-component-content-text, div.formula-cell-component'))
							.map(el => el.textContent); // traer texto, vacío si no hay
			return texts;
			})
			.flat()
		`, list),
	}
}
