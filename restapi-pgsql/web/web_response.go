package web

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"` //kita masukkan sebagai interface kosong biar datanya apapun,
}