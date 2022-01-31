package main

import "products/service"

/**func (p product) getProducts() {
	bs, err := ioutil.ReadFile("products.json")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return strings.Split(string(bs), ",")
}**/

func main() {
	//service.GetProducts("products.json")

	//ctx := context.Background()
	//service.Produce(ctx)

	service.Consume()
}
