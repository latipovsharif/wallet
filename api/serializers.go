package api

type walletCreateSerializer struct {
	tableName struct{} `pg:"wallet"`
	Name      string   `json:"name" validate:"required"`
}

type moneyTransferSerializer struct {
	SrcWallet int32   `json:"src_wallet" validate:"required"`
	DstWallet int32   `json:"dst_wallet" validate:"required"`
	Amount    float32 `json:"amount" validate:"gt=0"`
	TrnID     string  `json:"transaction_id" validate:"uuid4"`
}

type depositSerializer struct {
	Wallet int32   `json:"wallet" validate:"required"`
	Amount float32 `json:"amount" validate:"required"`
	TrnID  string  `json:"transaction_id" validate:"uuid4"`
}

type withdrawSerializer struct {
	Wallet int32   `json:"wallet" validate:"required"`
	Amount float32 `json:"amount" validate:"required"`
	TrnID  string  `json:"transaction_id" validate:"uuid4"`
}
