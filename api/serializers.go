package api

type walletCreateSerializer struct {
	tableName struct{} `pg:"wallet"`
	Name      string   `json:"name" validate:"required"`
}

type moneyTransferSerializer struct {
	SrcWallet int32
	DstWallet int32
	Amount    float32
	TrnID     string
}

type depositSerializer struct {
	Wallet int32
	Amount float32
	TrnID  string
}

type withdrawSerializer struct {
	Wallet int32
	Amount float32
	TrnID  string
}
