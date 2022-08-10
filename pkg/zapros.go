package pkg

import (
	"bestprice/models"
	"fmt"
)

func Zapros(article string) {
	models.Rez.Firms = nil
	models.Rez.Firms = make(map[string][]models.Str)
	myChan := make(chan models.Rezult)
	var rezult models.Rezult
	//fmt.Printf("Запрос на артикул %v\n", article)
	go Tiss(article, myChan)
	go Froza(article, myChan)
	go Mikado(article, myChan)
	go Ivers(article, myChan)
	go Forum(article, myChan)
	for i := 0; i < 5; i++ {
		rezult = <-myChan

		//fmt.Printf("Tiss: \n")
		for key, value := range rezult.Price {
			models.Rez.Firms[key] = append(models.Rez.Firms[key], models.Str{rezult.Firm, value})
			fmt.Printf("%v   %-16s  \t%v\n", rezult.Firm, key, value)
		}
	}
	/*froza, err := Froza(article)
	if err != nil {
		fmt.Printf("error %v", err)
	}
	//fmt.Printf("Froza: \n")
	for key, value := range froza {
		models.Rez.Firms[key] = append(models.Rez.Firms[key], models.Str{"Froza", value})
		//fmt.Printf("   %-16s  \t%v\n", key, value)
	}
	mikado, err := Mikado(article)
	if err != nil {
		fmt.Printf("error %v", err)
	}
	//fmt.Printf("Mikado: \n")
	for key, value := range mikado {
		models.Rez.Firms[key] = append(models.Rez.Firms[key], models.Str{"Mikado", value})
		//fmt.Printf("   %-16s  \t%v\n", key, value)
	}
	ivers, err := Ivers(article)
	if err != nil {
		fmt.Printf("error %v", err)
	}
	//fmt.Printf("Ivers: \n")
	for key, value := range ivers {
		models.Rez.Firms[key] = append(models.Rez.Firms[key], models.Str{"Ivers", value})
		//fmt.Printf("   %-16s  \t%v\n", key, value)
	}
	forum, err := Forum(article)
	if err != nil {
		fmt.Printf("error %v", err)
	}
	//fmt.Printf("Forum: \n")
	for key, value := range forum {
		models.Rez.Firms[key] = append(models.Rez.Firms[key], models.Str{"Forum", value})
		//fmt.Printf("   %-16s  \t%v\n", key, value)
	}*/
	//models.Rez.Keys = nil
	//models.Rez.Keys = make([]string)
	models.Rez.Viev = nil
	for key, value := range models.Rez.Firms {
		models.Rez.Viev = append(models.Rez.Viev, models.Str2{key, value})
		//models.Rez.Keys = append(models.Rez.Keys, key)
	}

}
