package domain

// Order representa a entidade de pedido no domínio da aplicação
type Order struct {
	ID     int64   `json:"id"`      // ID único do pedido
	UserID int64   `json:"user_id"` // ID do usuário que fez o pedido
	Amount float64 `json:"amount"`  // Valor total do pedido
	Status string  `json:"status"`  // Status atual do pedido (ex: "pendente", "completo")
}
