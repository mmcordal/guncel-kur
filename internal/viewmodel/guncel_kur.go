package viewmodel

import "guncel-kur/internal/model"

type GuncelKurVM struct {
	CreatedAt        string  `json:"created_at"`
	Dolar            float64 `json:"dolar"`
	Sterlin          float64 `json:"sterlin"`
	Euro             float64 `json:"euro"`
	CeyrekAltin      float64 `json:"ceyrek_altin"`
	YarimAltin       float64 `json:"yarim_altin"`
	TamAltin         float64 `json:"tam_altin"`
	CumhuriyetAltini float64 `json:"cumhuriyet_altini"`
	Bilezik22Ayar    float64 `json:"bilezik_22_ayar"`
	GramAltin14Ayar  float64 `json:"gram_altin_14_ayar"`
	GramAltin18Ayar  float64 `json:"gram_altin_18_ayar"`
	GramAltin22Ayar  float64 `json:"gram_altin_22_ayar"`
	GramAltin24Ayar  float64 `json:"gram_altin_24_ayar"`
	Gumus            float64 `json:"gumus"`
}

func ToGuncelKurVM(k *model.GuncelKur) *GuncelKurVM {
	return &GuncelKurVM{
		CreatedAt:        k.CreatedAt.Format("2006-01-02 15:04:05"),
		Dolar:            k.Dolar,
		Sterlin:          k.Sterlin,
		Euro:             k.Euro,
		CeyrekAltin:      k.CeyrekAltin,
		YarimAltin:       k.YarimAltin,
		TamAltin:         k.TamAltin,
		CumhuriyetAltini: k.CumhuriyetAltini,
		Bilezik22Ayar:    k.Bilezik22Ayar,
		GramAltin14Ayar:  k.GramAltin14Ayar,
		GramAltin18Ayar:  k.GramAltin18Ayar,
		GramAltin22Ayar:  k.GramAltin22Ayar,
		GramAltin24Ayar:  k.GramAltin24Ayar,
		Gumus:            k.Gumus,
	}
}
