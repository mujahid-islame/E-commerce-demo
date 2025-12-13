package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hay Hi i am a hello Handler")
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hi i am mujahid islam and i am a student of polytecnic ")
}

type Product struct {
	Id          int
	Title       string
	Description string
	Price       float64
	ImgUrl      string
}

var productsList []Product

func getproducts(w http.ResponseWriter, r *http.Request) {

	encoder := json.NewEncoder(w)
	encoder.Encode(productsList)

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", helloHandler)
	mux.HandleFunc("GET /about", aboutHandler)
	mux.HandleFunc("GET /products", getproducts)
	fmt.Println("The Server is running frome :8080 port")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Server has error ", err)

	}
}

func init() {
	prd1 := Product{
		Id:          1,
		Title:       "Banana",
		Description: "Banana is a sweet fruit",
		Price:       30.5,
		ImgUrl:      "https://upload.wikimedia.org/wikipedia/commons/8/8a/Banana-Single.jpg",
	}

	prd2 := Product{
		Id:          2,
		Title:       "Apple",
		Description: "Apple is a healthy fruit",
		Price:       120.0,
		ImgUrl:      "https://upload.wikimedia.org/wikipedia/commons/1/15/Red_Apple.jpg",
	}

	prd3 := Product{
		Id:          3,
		Title:       "Orange",
		Description: "Orange is full of vitamin C",
		Price:       80.0,
		ImgUrl:      "https://upload.wikimedia.org/wikipedia/commons/c/c4/Orange-Fruit-Pieces.jpg",
	}

	prd4 := Product{
		Id:          4,
		Title:       "Mango",
		Description: "Mango is the king of fruits",
		Price:       150.0,
		ImgUrl:      "https://upload.wikimedia.org/wikipedia/commons/9/90/Hapus_Mango.jpg",
	}

	prd5 := Product{
		Id:          5,
		Title:       "Pineapple",
		Description: "Pineapple is juicy fruit",
		Price:       90.0,
		ImgUrl:      "https://upload.wikimedia.org/wikipedia/commons/c/cb/Pineapple_and_cross_section.jpg",
	}

	prd6 := Product{
		Id:          6,
		Title:       "Grapes",
		Description: "Grapes are small and sweet",
		Price:       110.0,
		ImgUrl:      "https://upload.wikimedia.org/wikipedia/commons/1/15/Red_grapes.jpg",
	}

	prd7 := Product{
		Id:          7,
		Title:       "Watermelon",
		Description: "Watermelon is refreshing",
		Price:       60.0,
		ImgUrl:      "https://upload.wikimedia.org/wikipedia/commons/e/ee/Watermelon_cross_BNC.jpg",
	}

	prd8 := Product{
		Id:          8,
		Title:       "Papaya",
		Description: "Papaya helps digestion",
		Price:       70.0,
		ImgUrl:      "https://upload.wikimedia.org/wikipedia/commons/6/6b/Papaya_cross_section_BNC.jpg",
	}

	prd9 := Product{
		Id:          9,
		Title:       "Guava",
		Description: "Guava is rich in fiber",
		Price:       50.0,
		ImgUrl:      "https://upload.wikimedia.org/wikipedia/commons/0/02/Guava_ID.jpg",
	}

	prd10 := Product{
		Id:          10,
		Title:       "Strawberry",
		Description: "Strawberry is very tasty",
		Price:       200.0,
		ImgUrl:      "https://upload.wikimedia.org/wikipedia/commons/2/29/PerfectStrawberry.jpg",
	}

	// productsList = append(
	// 	productsList,
	// 	prd1, prd2, prd3, prd4, prd5,
	// 	prd6, prd7, prd8, prd9, prd10,
	// )
	productsList = append(productsList, prd1)
	productsList = append(productsList, prd2)
	productsList = append(productsList, prd3)
	productsList = append(productsList, prd4)
	productsList = append(productsList, prd5)
	productsList = append(productsList, prd6)
	productsList = append(productsList, prd7)
	productsList = append(productsList, prd8)
	productsList = append(productsList, prd9)
	productsList = append(productsList, prd10)

}
