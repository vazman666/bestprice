package pkg

import (
	"bestprice/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func Forum(article string) (map[string]float32, error) {
	fmt.Printf("Запрос к Forum на %v\n", article)
	u, err := url.Parse("https://api.forum-auto.ru/v2/listGoods")
	//u, err := url.Parse("http://localhost:8080")
	if err != nil {
		return nil, fmt.Errorf("Ошибка url.Parse %v", err)
	}
	q := u.Query()
	q.Set("login", "460469_taygantsevIP")
	q.Set("pass", "4698251G7eMp")
	q.Set("art", article)
	//q.Set("_", "1653230127174")
	u.RawQuery = q.Encode()
	client := &http.Client{}
	r, err := http.NewRequest(http.MethodGet, u.String(), nil) // URL-encoded payload
	if err != nil {
		return nil, fmt.Errorf("Ошибка создания запроса NewRequest %v", err)
	}
	r.Header.Add("Accept", `application/xml`)
	r.Header.Add("Cookie", "auto_user_id=11897; auto_user_login=TTG5; auto_user_pass=FRGqKY0iyYLSc; auto_user_language=ru; user=jonoh8519pmdkr4h6asu9nbvti")
	//r.Header.Add("Authorization", "Bearer loAM8nyBc_inqml98_6VuaPLkWZUPsIxclJTBN1xmszVKnCOd-NqNL7-aoyEuziA-kiJCR5r0sx8soB1Z4vhIg0LiUq_BgkI2eUK3cB1LMGVK1fvlN0wkdH5AaQA7UivDYBU7cYU_4iJe_sqWAeDuADw42nnVXLDgYuP__zk2uchwxj5qk7luo8Az2scKRV9F2sLRFmuYq9wFWSF0IQWaG2zeKnz2XbtbtAJ-P5mb0-Sqze4OyfGg7OwSpJbfXtO3BKH1HRfkxpeh5_jTCswkhJwogKaReXjjXqjb6CyRcY42yyzkZniCeY7Pm4u2djmETJQKaHVk7DPWDCquoMpmOiKvFt2z4WjQ5sCOTeqZYjgYKmMmACn504eLCIZKVO4xWtxsXDR03-e8UqWjBaiU4jq4W-GbdZ_AZPtxT4RkRKA4C0V86FBP1ACm734Rg8gS923_PSQf-CQsbchylr0X-3ZQdpQxdb20KGVTzE_8Hj8Y8UWXxyuqqmgObPGxRaA") // добавляем заголовок Accept
	resp, err := client.Do(r)
	if err != nil {
		return nil, fmt.Errorf("Ошибка выполнения запроса DO %v", err)
	}
	defer resp.Body.Close()
	//io.Copy(os.Stdout, resp.Body)
	rez := make(map[string]float32)
	var targets []models.Forum
	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Printf("body=%v\n",string(body))
	err = json.Unmarshal(body, &targets)
	if err != nil {
		//fmt.Println(err)
		rez["Не нашлось нифига"] = 100500
		//return nil, fmt.Errorf("Ошибка unmarshal %v", err)
		return rez, nil
	}
	//fmt.Printf("%v\n", targets)
	if len(targets) == 0 {
		rez["нет ничего"] = 100500
		return rez, nil
	}
	rez[targets[0].Brand] = targets[0].Price
	for _, value := range targets {
		if rez[value.Brand] > value.Price {
			rez[value.Brand] = value.Price
		}
	}

	return rez, nil
}
