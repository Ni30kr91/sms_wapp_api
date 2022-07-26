package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func sendWp(quote string) {

	apiurl := "https://api.twilio.com/2010-04-01/Accounts/AC71b7dd8470ec0632e19c655e1ee43745/Messages.json"
	method := "POST"

	data := url.Values{}
	data.Set("To", "whatsapp:+917205767391")
	data.Set("From", "whatsapp:+14155238886")
	//data.Set("MessagingServiceSid", "iR3zdSPkQHfgR9QyeRQRWvkaZeknPjmQ")
	data.Set("Body", quote)
	client := &http.Client{}
	req, err := http.NewRequest(method, apiurl, strings.NewReader(data.Encode()))

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Basic QUM3MWI3ZGQ4NDcwZWMwNjMyZTE5YzY1NWUxZWU0Mzc0NTo5M2YyMmYyZjg3NTZlN2E4NzZhOGU4YmUwZTZjZGQwZQ==")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
