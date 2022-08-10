package pkg

import (
	"bestprice/models"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func Mikado(article string) (map[string]float32, error) {
	article = strings.ToLower(article)
	u, err := url.Parse("http://www.mikado-parts.ru/ws1/service.asmx/Code_Search")
	//u, err := url.Parse("http://localhost:8080")
	if err != nil {
		return nil, fmt.Errorf("Ошибка url.Parse %v", err)
	}
	q := u.Query()
	q.Set("Search_Code", article)
	q.Set("ClientID", "33939")
	q.Set("Password", "slava116")
	q.Set("FromStockOnly", "FromStockAndByOrder")
	u.RawQuery = q.Encode()
	client := &http.Client{}
	r, err := http.NewRequest(http.MethodGet, u.String(), nil) // URL-encoded payload
	if err != nil {
		return nil, fmt.Errorf("Ошибка создания запроса NewRequest %v", err)
	}
	r.Header.Add("Accept", `application/xml`)
	//r.Header.Add("Cookie", "auto_user_id=11897; auto_user_login=TTG5; auto_user_pass=FRGqKY0iyYLSc; auto_user_language=ru; user=jonoh8519pmdkr4h6asu9nbvti")
	//r.Header.Add("Authorization", "Bearer loAM8nyBc_inqml98_6VuaPLkWZUPsIxclJTBN1xmszVKnCOd-NqNL7-aoyEuziA-kiJCR5r0sx8soB1Z4vhIg0LiUq_BgkI2eUK3cB1LMGVK1fvlN0wkdH5AaQA7UivDYBU7cYU_4iJe_sqWAeDuADw42nnVXLDgYuP__zk2uchwxj5qk7luo8Az2scKRV9F2sLRFmuYq9wFWSF0IQWaG2zeKnz2XbtbtAJ-P5mb0-Sqze4OyfGg7OwSpJbfXtO3BKH1HRfkxpeh5_jTCswkhJwogKaReXjjXqjb6CyRcY42yyzkZniCeY7Pm4u2djmETJQKaHVk7DPWDCquoMpmOiKvFt2z4WjQ5sCOTeqZYjgYKmMmACn504eLCIZKVO4xWtxsXDR03-e8UqWjBaiU4jq4W-GbdZ_AZPtxT4RkRKA4C0V86FBP1ACm734Rg8gS923_PSQf-CQsbchylr0X-3ZQdpQxdb20KGVTzE_8Hj8Y8UWXxyuqqmgObPGxRaA") // добавляем заголовок Accept
	resp, err := client.Do(r)
	if err != nil {
		return nil, fmt.Errorf("Ошибка выполнения запроса DO %v", err)
	}
	defer resp.Body.Close()
	//io.Copy(os.Stdout, resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Printf("%v", body)
	var res models.Mikado
	err = xml.Unmarshal(body, &res)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%v\n", res)
	rez := make(map[string]float32)
	//rez[article] = 100500

	for _, value := range res.List[0].Code_List_Row {
		//fmt.Printf(" val %v\n", value)
		if value.ProducerCode == article {
			//fmt.Printf("Это наше\n")
			if s, err := strconv.ParseFloat(value.PriceRUR, 32); err == nil {
				if rez[value.Brand] == 0 || rez[value.Brand] > float32(s) {
					rez[value.Brand] = float32(s)
				}
			}
		}
	}
	//fmt.Printf("rez=%v\n", rez)
	return rez, nil
}
