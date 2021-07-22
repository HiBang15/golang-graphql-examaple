package main

import (
	"github.com/HiBang15/golang-graphql-examaple.git/api/graphql"
	"github.com/HiBang15/golang-graphql-examaple.git/model"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"net/url"
)

func giaiThua(n int) int {
	if (n == 1) {
		return n
	}
	return n * giaiThua(n-1)
}

func soHoanHao(n int) bool  {
	sum := 0
	for i :=0; i<n; i++ {
		if sum % i == 0 {
			sum += i
		}
	}
	if sum == n {
		return true
	}
	return false
}

////xoa ki tu cua 1 chuoi
//func xoaKiTu(a string, c string) string {
//	strLen := len(a)
//	for i := 0; i< strLen; i++ {
//		for j := i; j<strLen; j ++ {
//			if a[i] == c[0] {
//				a[j] = a[j+1]
//				a[strLen] = '\0'
//				i--
//			}
//		}
//	}
//}
//
//func xoaKiTu(a string, c string) string {
//	//mapp := make(map[string]string)
//	//for k, v := range mapp {
//	//
//	//}
//	//xoa ki tu c trong chuoi a
//	return strings.Replace(a, c, "", -1)
//}
//
//func getNumber(s string) (string, string) {
//	//xoa tung ki tu lan luot trong chuoi, luu duoi dang map
//	//tim ra value nho nhat trong map
//	//tu value -> key
//	var strMap = make(map[string]string)
//	for i := 0; i < len(s); i ++ {
//		var strNew = xoaKiTu(s, string(s[i]))
//		strMap[string(s[i])] = strNew
//	}
//
//	min,_ := strconv.Atoi(strMap[string(s[0])])
//	minChar := string (s[0])
//	for key, value := range strMap {
//		num, _ := strconv.Atoi(value)
//		if num < min {
//			min = num
//			minChar = string(key)
//		}
//	}
//	return fmt.Sprintf("%d", min), minChar
//}


//func chuoiNhoNhat(a string) string {
//	lengthStr := len(a)
//	arrNumDeleted := make([]string, 100000000)
//	for i := 0; i < lengthStr; i++ {
//		if a[i] == a[i+1] {
//
//		}
//
//	}
//}

func main() {
	routes := chi.NewRouter()
	r := graphql.RegisterRoutes(routes)
	log.Println("Server ready at 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func init() {
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	env := model.Environment{}
	env.SetEnvironment()
	env.LoadConfig()
	env.InitMongoDB()
}
