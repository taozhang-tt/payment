package ali

type Extra struct {
	SignType         string `json:"sign_type"` // 签名类型：RSA、RSA2
	IsSandbox        bool   `json:"is_sandbox"`
	SpecifiedChannel string `json:"specified_channel"`
}
