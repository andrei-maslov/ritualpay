package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/andrei-maslov/ritualpay/internal/domain"
	"github.com/andrei-maslov/ritualpay/internal/parser"
	"github.com/andrei-maslov/ritualpay/internal/report"
	"github.com/andrei-maslov/ritualpay/internal/utils"
)

func main() {
	fmt.Println("RitaulPay запущен!")

	err := utils.ClearReportDir()
	if err != nil {
		fmt.Println("ERR: Не удалось очистить папку со старыми отчетами")
		fmt.Println(err)
	}
	err = utils.CreateReportDir()
	if err != nil {
		fmt.Println("ERR: Не удалось создать папку для отчетов")
		fmt.Println(err)
	}

	files, _ := utils.GetOrderFiles()
	var orders []*domain.Order

	fmt.Printf("Найденные файлы с ответами:\n\t%s", strings.Join(files, "\n\t"))

	for _, file := range files {
		order, err := parser.Parse(file)
		if err != nil {
			fmt.Println("\n*** ERR: Не удалось прочитать заказ в файле: " + file)
			fmt.Println(err)
			continue
		}
		orders = append(orders, order)
	}
	fmt.Printf("\nПрочтено %d заказ(а)/(ов)\n", len(orders))

	// Отчет по суммам для исполнителей
	fmt.Println("Формирование сводного отчета")
	summaryReport := report.PerformerSummaryReport{}
	summaryReportTest := summaryReport.Generate(orders)
	fmt.Println("Cводного отчет готов!\n")
	err = os.WriteFile(utils.ReportDir()+"/Сводный отчет.txt", []byte(summaryReportTest), 0644)
	if err != nil {
		fmt.Println("ERR: Не удалось сохранить сводный отчет")
		fmt.Println(err)
	}

	// Отчет по конкретному исполнителю
	fmt.Println("Полный список исполнителей: " + strings.Join(summaryReport.GetPerformers(), ","))
	performerReports := make(map[string]string)
	for _, performer := range summaryReport.GetPerformers() {
		fmt.Printf("Формирование подробного отчета для %s\n", performer)
		detailsReport := report.PerformerDetailsReport{PerformerName: performer}
		performerReports[performer] = detailsReport.Generate(orders)

		err = os.WriteFile(utils.ReportDir()+"/Отчет по исполнителю - "+performer+".txt", []byte(performerReports[performer]), 0644)
		if err != nil {
			fmt.Println("ERR: Не удалось сохранить подробный отчет по " + performer)
			fmt.Println(err)
		}
	}

}
