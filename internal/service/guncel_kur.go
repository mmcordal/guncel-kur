package service

import (
	"context"
	"encoding/json"
	"fmt"
	"guncel-kur/internal/infrastructure/cache"
	"guncel-kur/internal/model"
	"guncel-kur/internal/viewmodel"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const kurCacheKey = "guncel_kur_data"

type cachePayload struct {
	Data *viewmodel.GuncelKurVM `json:"data"`
}

type KurService interface {
	FetchFromTDV() (*viewmodel.GuncelKurVM, error)
}

type kurService struct {
	cache *cache.RedisClient
}

func NewKurService(c *cache.RedisClient) KurService {
	return &kurService{cache: c}
}

func parsePrice(s string) float64 {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, ".", "")
	s = strings.ReplaceAll(s, ",", ".")
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Println("price parse error:", err)
		return 0
	}
	return f
}

func (s *kurService) FetchFromTDV() (*viewmodel.GuncelKurVM, error) {
	cachedStr, err := s.cache.Get(context.Background(), kurCacheKey) // burda bunu direkt çekmeden önce saat kontrolü yapıp eğer sonraysa öyle çekebiliriz sanırım
	if err == nil && cachedStr != "" {
		var payload cachePayload
		if json.Unmarshal([]byte(cachedStr), &payload) == nil {
			now := time.Now()
			today17 := time.Date(now.Year(), now.Month(), now.Day(), 17, 0, 0, 0, now.Location())

			if now.Before(today17) {
				log.Println("Redis cache kullanıldı (17:00 öncesi)")
				return payload.Data, nil
			}

			log.Println("Saat 17:00 sonrası, cache güncellenecek")
		}
	} else {
		log.Println("Redis'te cache yok veya okunamadı")
	}

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Get("https://zekathesapla.tdv.org/Kur-Bilgisi")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("tdv response not OK: %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	kur := &model.GuncelKur{}

	doc.Find("table.table-bordered tbody tr").Each(func(i int, s *goquery.Selection) {
		tds := s.Find("td")
		if tds.Length() < 4 {
			return
		}

		name := strings.TrimSpace(tds.Eq(0).Text())
		sell := parsePrice(tds.Eq(2).Text())

		if name == "" {
			return
		}

		switch name {
		case "Amerikan Doları (USD)":
			kur.Dolar = sell
		case "Euro (EUR)":
			kur.Euro = sell
		case "İngiliz Sterlini (GBP)":
			kur.Sterlin = sell
		case "Çeyrek Altın (C)":
			kur.CeyrekAltin = sell
		case "Yarım Altın (Y)":
			kur.YarimAltin = sell
		case "Tam Altın (TAM)":
			kur.TamAltin = sell
		case "Cumhuriyet Altını (CMHT)":
			kur.CumhuriyetAltini = sell
		case "22 Ayar Bilezik (B)":
			kur.Bilezik22Ayar = sell
		case "14 Ayar Gram Altın (14)":
			kur.GramAltin14Ayar = sell
		case "18 Ayar Gram Altın (18)":
			kur.GramAltin18Ayar = sell
		case "22 Ayar Gram Altın (GA22)":
			kur.GramAltin22Ayar = sell
		case "24 Ayar Gram Altın (GA)":
			kur.GramAltin24Ayar = sell
		case "Gümüş (AG_T)":
			kur.Gumus = sell
		}
	})

	kur.CreatedAt = time.Now()
	vm := viewmodel.ToGuncelKurVM(kur)

	payload := cachePayload{Data: vm}
	bytes, _ := json.Marshal(payload)
	_ = s.cache.Set(context.Background(), kurCacheKey, string(bytes), 24*time.Hour)

	log.Println("Güncel kur verileri başarıyla alındı:", fmt.Sprintf("%+v", kur))

	return vm, nil
}
