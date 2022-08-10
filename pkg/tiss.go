package pkg

import (
	"bestprice/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	//"reflect"
	//"io"
	//"os"
)

func Tiss(article string) (map[string]float32, error) {
	fmt.Printf("Запрос к Tiss на %v\n", article)
	brands := brand(article)
	var targets models.Tiss
	rez := make(map[string]float32)
	for _, i := range brands {
		type Brand_Article_List struct {
			Brand   string `json:"Brand"`
			Article string `json:"Article"`
		}
		type zapros struct {
			Is_main_warehouse int                  `json:"is_main_warehouse"`
			BrandArticleList  []Brand_Article_List `json:"Brand_Article_List"`
		}
		var message zapros
		message.Is_main_warehouse = 0
		message.BrandArticleList = append(message.BrandArticleList, Brand_Article_List{i, article})

		//fmt.Printf("%v\n", message)
		bytesRepresentation, err := json.Marshal(message)
		//fmt.Printf("%s\n", bytesRepresentation)
		if err != nil {
			return nil, fmt.Errorf("Ошибка Masrshal %v", err)
		}
		u, err := url.Parse("http://api.tmparts.ru/api/StockByArticleList")
		if err != nil {
			return nil, fmt.Errorf("Ошибка url.Parse %v", err)
		}
		q := u.Query()
		q.Set("JSONparameter", string(bytesRepresentation))
		u.RawQuery = q.Encode()
		//fmt.Printf("%v\n", u.String())
		client := &http.Client{}
		r, _ := http.NewRequest(http.MethodGet, u.String(), nil) // URL-encoded payload
		r.Header.Add("Accept", `application/json`)
		r.Header.Add("Authorization", "Bearer loAM8nyBc_inqml98_6VuaPLkWZUPsIxclJTBN1xmszVKnCOd-NqNL7-aoyEuziA-kiJCR5r0sx8soB1Z4vhIg0LiUq_BgkI2eUK3cB1LMGVK1fvlN0wkdH5AaQA7UivDYBU7cYU_4iJe_sqWAeDuADw42nnVXLDgYuP__zk2uchwxj5qk7luo8Az2scKRV9F2sLRFmuYq9wFWSF0IQWaG2zeKnz2XbtbtAJ-P5mb0-Sqze4OyfGg7OwSpJbfXtO3BKH1HRfkxpeh5_jTCswkhJwogKaReXjjXqjb6CyRcY42yyzkZniCeY7Pm4u2djmETJQKaHVk7DPWDCquoMpmOiKvFt2z4WjQ5sCOTeqZYjgYKmMmACn504eLCIZKVO4xWtxsXDR03-e8UqWjBaiU4jq4W-GbdZ_AZPtxT4RkRKA4C0V86FBP1ACm734Rg8gS923_PSQf-CQsbchylr0X-3ZQdpQxdb20KGVTzE_8Hj8Y8UWXxyuqqmgObPGxRaA") // добавляем заголовок Accept
		resp, err := client.Do(r)
		if err != nil {
			return nil, fmt.Errorf("Ошибка client.DO %v", err)
		}
		defer resp.Body.Close()
		//io.Copy(os.Stdout, resp.Body)
		//fmt.Printf("\nresponse=%v\n", resp.Status)

		//var targets models.Tiss
		body, err := ioutil.ReadAll(resp.Body)
		//fmt.Printf("body=%v\n",string(body))
		err = json.Unmarshal(body, &targets)
		if err != nil {
			return nil, fmt.Errorf("Ошибка Unmarshal %v", err)
		}

		//fmt.Printf("%v\v", targets)
		if len(targets) != 0 {
			min := targets[0].Warehouse_offers[0]
			for _, j := range targets[0].Warehouse_offers {
				if j.Price < min.Price {
					min = j
				}
				//rez = append(rez, models.Result{j.Price, i})
				//fmt.Printf("rez append %v %v  %v\n", i, j, min)
			}
			//fmt.Printf("min=%v\n", min)
			rez[i] = min.Price
		}
		//fmt.Printf("\nДля номера %v\n", targets.Article)
		/*for _, i := range targets.Warehouse_offers {
			fmt.Printf("%v\n", i)
		}*/
	}
	//fmt.Printf("rez = %v\n", rez)
	if len(rez) == 0 {
		rez["нет ничего"] = 100500
		return rez, nil
	}
	/*min := rez[0]
	for _, i := range rez {
		if min.Price > i.Price {
			min = i
		}
	}*/
	return rez, nil
}
func brand(article string) []string {
	type Brand_Article_List struct {
		Article string `json:"Article"`
	}

	var message Brand_Article_List
	message.Article = article

	//fmt.Printf("%v\n", message)
	bytesRepresentation, err := json.Marshal(message)
	//fmt.Printf("%s\n", bytesRepresentation)
	if err != nil {
		log.Fatalln(err)
	}
	u, err := url.Parse("http://api.tmparts.ru/api/ArticleBrandList")
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	q.Set("JSONparameter", string(bytesRepresentation))
	u.RawQuery = q.Encode()
	//fmt.Printf("%v\n", u.String())
	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodGet, u.String(), nil) // URL-encoded payload
	r.Header.Add("Accept", `application/json`)
	r.Header.Add("Authorization", "Bearer loAM8nyBc_inqml98_6VuaPLkWZUPsIxclJTBN1xmszVKnCOd-NqNL7-aoyEuziA-kiJCR5r0sx8soB1Z4vhIg0LiUq_BgkI2eUK3cB1LMGVK1fvlN0wkdH5AaQA7UivDYBU7cYU_4iJe_sqWAeDuADw42nnVXLDgYuP__zk2uchwxj5qk7luo8Az2scKRV9F2sLRFmuYq9wFWSF0IQWaG2zeKnz2XbtbtAJ-P5mb0-Sqze4OyfGg7OwSpJbfXtO3BKH1HRfkxpeh5_jTCswkhJwogKaReXjjXqjb6CyRcY42yyzkZniCeY7Pm4u2djmETJQKaHVk7DPWDCquoMpmOiKvFt2z4WjQ5sCOTeqZYjgYKmMmACn504eLCIZKVO4xWtxsXDR03-e8UqWjBaiU4jq4W-GbdZ_AZPtxT4RkRKA4C0V86FBP1ACm734Rg8gS923_PSQf-CQsbchylr0X-3ZQdpQxdb20KGVTzE_8Hj8Y8UWXxyuqqmgObPGxRaA") // добавляем заголовок Accept
	resp, err := client.Do(r)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()
	//io.Copy(os.Stdout, resp.Body)
	//fmt.Printf("\nresponse=%v\n", resp.Status)

	var targets models.TissBrand
	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Printf("body=%v\n",string(body))
	err = json.Unmarshal(body, &targets)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	rez := []string{}
	for _, i := range targets.BrandList {
		rez = append(rez, i.BrandName)
	}
	//fmt.Printf("Бренда Тисса %v\n", rez)
	return rez

}
