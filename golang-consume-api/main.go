package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://instagram-scraper-api2.p.rapidapi.com/v1/likes?code_or_id_or_url=CxYQJO8xuC6"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	req.Header.Add("instagram-scraper-api2.p.rapidapi.com", "4bd94a6f59msh7a7afd5aa227628p191c40jsn5841142c8abf")


	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer res.Body.Close()

	body, errRes := ioutil.ReadAll(res.Body)
	if errRes != nil {
		fmt.Print(errRes.Error())
	}

	fmt.Print(string(body))

}
