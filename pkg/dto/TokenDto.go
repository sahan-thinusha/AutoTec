package dto

type TokenDto struct {
	Token      string `json:"token"`
	Token_type string `json:"grant_type"`
	Expiresin  int64  `json:"expiresin"`
	Error      string `json:"error"`
}
