package models

type Pagamento struct {
	Cliente    string            `json:"cliente"`
	Troco      float64           `json:"troco"`
	Credito    float64           `json:"credito"`
	Debito     float64           `json:"debito"`
	Dinheiro   float64           `json:"dinheiro"`
	Picpay     float64           `json:"picpay"`
	Pix        float64           `json:"pix"`
	Persycoins float64           `json:"persycoins"`
	Data       string            `json:"data"`
	Pedidos    map[string]Pedido `json:"pedidos"`
}

type Pedido struct {
	Nome       string  `json:"nome"`
	Preco      float64 `json:"preco"`
	Quantidade int     `json:"quantidade"`
}
