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

func Ivers(article string) (map[string]float32, error) {
	fmt.Printf("Запрос к Ivers на %v\n", article)
	cookie, err := Login()

	if err != nil {
		return nil, err
	}
	u, err := url.Parse("https://order.ivers.ru/api/v1/product/search")
	if err != nil {
		return nil, fmt.Errorf("Ошибка url.Parse %v", err)
	}
	q := u.Query()

	q.Set("text", article)
	//q.Set("type", "ivers-assortment")
	q.Set("typesOriginal[]", "ivers-assortment")
	q.Add("typesOriginal[]", "inomarket")
	q.Add("typesOriginal[]", "ivers-no-assortment")
	q.Add("typesOriginal[]", "alternative")
	u.RawQuery = q.Encode()
	//fmt.Printf("%v\n", u.String())
	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, u.String(), nil) // URL-encoded payload
	r.Header.Add("Accept", `application/json`)
	r.Header.Add("Cookie", cookie)
	resp, err := client.Do(r)
	if err != nil {
		return nil, fmt.Errorf("Ошибка client.DO %v", err)
	}
	defer resp.Body.Close()
	//io.Copy(os.Stdout, resp.Body)
	//fmt.Printf("\nresponse=%v\n", resp.Status)

	var targets models.Ivers
	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Printf("body=%v\n",string(body))
	err = json.Unmarshal(body, &targets)
	if err != nil {
		return nil, fmt.Errorf("Ошибка Unmarshal %v", err)
	}

	//fmt.Printf("%v\v", targets)
	if !targets.Success {
		return nil, fmt.Errorf("Ошибка запроса к Иверс %v", err)
	}
	if !targets.Result.Params.Success {
		return nil, fmt.Errorf("Ошибка запроса к Иверс %v", err)
	}
	rez := make(map[string]float32)
	for _, i := range targets.Result.Params.Objects.Original {
		//fmt.Printf("%v\n\n", i)
		if rez[i.Brand] == 0 || rez[i.Brand] > i.Prices[0].Price {
			rez[i.Brand] = i.Prices[0].Price
		}
	}

	/*if len(targets) != 0 {
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


	//fmt.Printf("rez = %v\n", rez)
	if len(rez) == 0 {
		rez["нет ничего"] = 100500
		return rez, nil
	}

	min := rez[0]
	for _, i := range rez {
		if min.Price > i.Price {
			min = i
		}
	}*/
	Logout(cookie)
	return rez, nil
}
func Login() (string, error) {
	u, err := url.Parse("https://order.ivers.ru/api/v1/login")
	if err != nil {

		return "", fmt.Errorf("Ошибка url.Parse %v", err)
	}
	q := u.Query()
	q.Set("login", models.Login)
	q.Set("password", models.Pass)
	u.RawQuery = q.Encode()

	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, u.String(), nil) // URL-encoded payload
	r.Header.Add("Accept", `application/json`)
	resp, err := client.Do(r)
	if err != nil {
		return "", fmt.Errorf("Ошибка DO %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		//fmt.Printf("Ошибка при авторизации.. Неверный логин/пароль? \n")
		return "", fmt.Errorf("Ошибка при авторизации.. Неверный логин/пароль?")
	}
	//fmt.Println(resp.Status)
	targets := models.LoginType{}
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &targets)
	if err != nil {
		log.Fatal(err)
	}
	if !targets.Success {
		//fmt.Printf("Ошибка авторизации:Ошибка запроса\n")
		return "", fmt.Errorf("Ошибка авторизации:Ошибка запроса")
	}
	if !targets.Result.Params.Success {
		//fmt.Printf("Ошибка авторизации:Неправильный логин/пароль?\n")
		return "", fmt.Errorf("Ошибка авторизации:Неправильный логин/пароль?")
	}
	return targets.Result.Params.Token, nil
}
func Logout(token string) error {
	u, err := url.Parse("https://order.ivers.ru/api/v1/logout")
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	u.RawQuery = q.Encode()
	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, u.String(), nil) // URL-encoded payload
	r.Header.Add("Accept", `application/json`)
	r.Header.Add("Cookie", token)
	resp, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		fmt.Printf("Ошибка при выходе.  \n")
		return fmt.Errorf("Ошибка при выходе")
	}
	//fmt.Println(resp.Status)
	targets := models.Logout{}
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &targets)
	if err != nil {
		log.Fatal(err)
	}
	if !targets.Success {
		fmt.Printf("Ошибка выхода \n")
		return fmt.Errorf("Ошибка выхода")
	}
	return nil
}
