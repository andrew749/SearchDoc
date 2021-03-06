package main

import (
	"flag"
	"fmt"
	"searchdoc/src/core"
	"searchdoc/src/ui"
)

func main() {

	// get the arguments
	var (
		query    string
		language string
	)

	flag.StringVar(&query, "query", "", "The query to search")
	flag.StringVar(&language, "language", "", "The query to search")

	download_list := flag.Bool("download_list", false, "Indicate if you want to list the downloadable packages.")
	installed_list := flag.Bool("list", false, "List all installed packages.")
	package_to_download := flag.String("download", "", "Download the specified package.")

	flag.Parse()

	if *download_list {
		for _, x := range core.ListDownloadableDocsets() {
			fmt.Println(x)
		}
		return
	} else if *installed_list {
		for _, x := range core.ListInstalledDocsets() {
			fmt.Println(x)
		}
		return
	} else if *package_to_download != "" {
		core.DownloadDocset(*package_to_download)
		return
	}

	fmt.Printf("language: %s\nquery: %s\n", language, query)
	// process the command
	// TODO: replace with connection to ui
	results, err := core.Query(query, language)
	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		return
	}

	resultStrs := make([]string, len(results))
	for i, result := range results {
		resultStrs[i] = result.Name
	}
	ui.SetQueryResults(resultStrs)

	if data, err := core.GetDocContent(results[0], language); err == nil {
		ui.SetHtmlContent(data)
	}

	ui.Init()

	//fmt.Print(data)
}
