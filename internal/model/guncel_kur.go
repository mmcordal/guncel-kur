package model

import "time"

type GuncelKur struct {
	CreatedAt        time.Time `json:"created_at"`
	Dolar            float64   `json:"dolar"`
	Sterlin          float64   `json:"sterlin"`
	Euro             float64   `json:"euro"`
	CeyrekAltin      float64   `json:"ceyrek_altin"`
	YarimAltin       float64   `json:"yarim_altin"`
	TamAltin         float64   `json:"tam_altin"`
	CumhuriyetAltini float64   `json:"cumhuriyet_altini"`
	Bilezik22Ayar    float64   `json:"bilezik_22_ayar"`
	GramAltin14Ayar  float64   `json:"gram_altin_14_ayar"`
	GramAltin18Ayar  float64   `json:"gram_altin_18_ayar"`
	GramAltin22Ayar  float64   `json:"gram_altin_22_ayar"`
	GramAltin24Ayar  float64   `json:"gram_altin_24_ayar"`
	Gumus            float64   `json:"gumus"`
}
