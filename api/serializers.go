package api

type walletCreateSerializer struct {
	tableName struct{} `pg:"wallets"`
	Name      string   `json:"name" binding:"required"`
}

type moneyTransferSerializer struct {
	SrcWallet int32   `json:"src_wallet" binding:"required"`
	DstWallet int32   `json:"dst_wallet" binding:"required"`
	Amount    float32 `json:"amount" binding:"gt=0"`
	TrnID     string  `json:"transaction_id" binding:"uuid4"`
}

type depositSerializer struct {
	Wallet int32   `json:"wallet" binding:"required"`
	Amount float32 `json:"amount" binding:"required"`
	TrnID  string  `json:"transaction_id" binding:"uuid4"`
}

type withdrawSerializer struct {
	Wallet int32   `json:"wallet" binding:"required"`
	Amount float32 `json:"amount" binding:"required"`
	TrnID  string  `json:"transaction_id" binding:"uuid4"`
}
