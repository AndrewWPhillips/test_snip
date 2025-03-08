package __

// tests for blog on unique used to store discount codes

import (
	crand "crypto/rand"
	"math/rand"
	"testing"
	"unique"
)

type (
	productID int
	Currency  int // price in cents
	product   struct {
		ID       productID // each product has a unique ID
		Price    Currency
		Discount []string // valid discount codes for the product
	}
	product2 struct {
		ID       productID // each product has a unique ID
		Price    Currency
		Discount []unique.Handle[string]
	}
	product3 struct {
		ID       productID // each product has a unique ID
		Price    Currency
		Discount map[string]struct{}
	}
	product4 struct {
		ID       productID // each product has a unique ID
		Price    Currency
		Discount map[unique.Handle[string]]struct{}
	}

	customer3 struct {
		Name     string
		Vouchers []string
	}
	customer4 struct {
		Name     string
		Vouchers []unique.Handle[string]
	}
)

// ---  BENCHMARKS ---

func BenchmarkVoucher3(b *testing.B) { // 17 to 70 ns/op
	products := GenerateProducts3(1000, 10)

	cust := GenerateCustomer3(5)
	prod := products[RandomInt(len(products))] // pick one at random

	for b.Loop() {
		for _, voucher := range cust.Vouchers {
			if _, found := prod.Discount[voucher]; found {
				break
			}
		}
	}
}

func BenchmarkVoucher4(b *testing.B) { // 5 to 25 ns/op
	products := GenerateProducts4(1000, 10)

	cust := GenerateCustomer4(5)
	prod := products[RandomInt(len(products))] // pick one at random

	for b.Loop() {
		for _, voucher := range cust.Vouchers {
			if _, found := prod.Discount[voucher]; found {
				break
			}
		}
	}
}

func BenchmarkDiscountCodeString(b *testing.B) {
	//const maxProducts, maxCodes = 100, 10  // 2683 ns/op
	const maxProducts, maxCodes = 100, 100 // 22191 ns/op
	products := make([]product, 0, maxProducts)

	for i := range maxProducts {
		codes := make([]string, RandomInt(maxCodes))
		for j := range codes {
			codes[j] = GenerateCode()
		}

		products = append(products, product{
			ID:       productID(i),
			Discount: codes,
		})
	}

	found, code := false, GenerateCode()
	for b.Loop() {
		for i := range products {
			// linear search comparing strings
			for j := range products[i].Discount {
				found = products[i].Discount[j] == code
			}
		}
	}
	println(found)
}

func BenchmarkDiscountCodeUnique(b *testing.B) { // 3110 ns/op
	//const maxProducts, maxCodes = 100, 10
	const maxProducts, maxCodes = 100, 100 // 3110 ns (about 7 times faster than BenchmarkDiscountCodeString)
	products := make([]product2, 0, maxProducts)

	for i := range maxProducts {
		codes := make([]unique.Handle[string], RandomInt(maxCodes))
		for j := range codes {
			codes[j] = unique.Make(GenerateCode())
		}

		products = append(products, product2{
			ID:       productID(i),
			Discount: codes,
		})
	}

	found, code := false, GenerateCode()
	for b.Loop() {
		h := unique.Make(code)
		for i := range products {
			// linear search comparing handles
			for j := range products[i].Discount {
				found = products[i].Discount[j] == h
			}
		}
	}
	println(found)
}

func BenchmarkDiscountCodeStringMap(b *testing.B) { // 1700 ns/op
	products := GenerateProducts3(100, 10)

	found, code := false, GenerateCode()
	for b.Loop() {
		for i := range products {
			// hash table lookup
			_, found = products[i].Discount[code]
		}
	}
	println(found)
}

func BenchmarkDiscountCodeUniqueMap(b *testing.B) { // 764
	products := GenerateProducts4(100, 10)

	found, code := false, GenerateCode()
	for b.Loop() {
		h := unique.Make(code)
		for i := range products {
			// hash table lookup
			_, found = products[i].Discount[h]
		}
	}
	println(found)
}

// --- HELPERS ---

func GenerateCustomer3(maxVouchers int) customer3 {
	codes := make([]string, RandomInt(maxVouchers))
	for j := range codes {
		codes[j] = GenerateCode()
	}
	return customer3{
		Vouchers: codes,
	}
}

func GenerateCustomer4(maxVouchers int) customer4 {
	codes := make([]unique.Handle[string], RandomInt(maxVouchers))
	for j := range codes {
		codes[j] = unique.Make(GenerateCode())
	}
	return customer4{
		Vouchers: codes,
	}
}

func GenerateProducts3(maxProducts, maxCodes int) (r []product3) {
	r = make([]product3, 0, maxProducts)

	for i := range maxProducts {
		codes := make(map[string]struct{})
		for range RandomInt(maxCodes) {
			codes[GenerateCode()] = struct{}{}
		}

		r = append(r, product3{
			ID:       productID(i),
			Discount: codes,
		})
	}
	return
}

func GenerateProducts4(maxProducts, maxCodes int) (r []product4) {
	r = make([]product4, 0, maxProducts)

	for i := range maxProducts {
		codes := make(map[unique.Handle[string]]struct{})
		for range RandomInt(maxCodes) {
			codes[unique.Make(GenerateCode())] = struct{}{}
		}

		r = append(r, product4{
			ID:       productID(i),
			Discount: codes,
		})
	}
	return
}

func RandomInt(n int) int {
	return rand.Intn(n)
	//r, err := crand.Int(crand.Reader, big.NewInt(int64(n)))
	//if err != nil {
	//	panic(err)
	//}
	//return int(r.Int64())
}

func GenerateCode() string {
	return crand.Text()
}
