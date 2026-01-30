# guncel-kur
## Güncel Kur API (Go & Redis)
Clean Architecture prensiplerine uygun olarak Go ile geliştirilmiş bu servis; Türkiye Diyanet Vakfı (TDV) kaynağından döviz ve kıymetli maden kur verilerini toplayan, Redis ile önbelleğe alan ve REST API üzerinden sunan bir backend uygulamasıdır.

### Bu projenin amacı:
- Harici bir kaynaktan veri çekme (web scraping)
- Elde edilen veriyi işleyerek domain modeline dönüştürme
- Redis kullanarak performans odaklı bir cache mekanizması kurma
- Katmanlı mimari (Clean Architecture) ile sürdürülebilir bir backend yapısı oluşturma

### Ana Özellikler:
- TDV'den döviz ve altın kuru verisi kazıma
- Redis tabanlı önbellekleme mekanizması
- Zamana dayalı önbellek geçersiz kılma stratejisi
- RESTful API tasarımı
- Clean Architecture ile uyumlu katmanlı mimari
- Merkezi hata işleme yapısı

### Mimari Genel Bakış: 
```pgsql
cmd/                    → Uygulama giriş noktası
internal/
   ├── service/         → İş mantığı katmanı
   ├── model/           → Domain modelleri
   ├── viewmodel/       → API response modelleri
   └── infrastructure/
           └── cache/   → Redis entegrasyonu
```

### Mimari Prensibi:
Bağımlılıklar içeri doğru akar. İş mantığı, çerçevelerden ve harici sistemlerden bağımsız kalır.

### Kullanılan Teknolojiler:
- Go
- Fiber (HTTP framework)
- Redis
- Goquery (HTML parsing)
- Viper (configuration management)

### Uygulama Akışı:
1. Bir istemci API'ye bir istek gönderir.
2. Hizmet, önbelleğe alınmış hız verileri için Redis'i kontrol eder.
3. Geçerli önbelleğe alınmış veriler varsa, doğrudan döndürülür.
4. Aksi takdirde:
  - TDV kaynağı istenir.
  - HTML tablo ayrıştırılır.
  - Alan modeli oluşturulur.
  - Veriler Redis'e depolanır.
  - Yanıt istemciye döndürülür.

## Kurulum ve Çalıştırma 

### Gereksinimler
- Go 1.20+
- Redis sunucusu

### Uygulamayı çalıştırma
```bash
go mod tidy
go run cmd/main.go
```

### Varsayılan Yapılandırma:
```go
redis.host = localhost
redis.port = 6379
server.port = 3000
```
> **Önemli:** Gerekirse bu değerler ortam değişkenleri aracılığıyla geçersiz kılınabilir.

### Örnek Endpoint:
```bash
GET /guncel-kur
```

### Örnek response:
```json
{
  "dolar": 32.45,
  "euro": 34.12,
  "sterlin": 40.22,
  "gramAltin24Ayar": 2450.15
}
```

> **Notlar:**
- Uygulama, TDV sayfasının HTML yapısına bağlıdır. Yapısal değişiklikler ayrıştırıcı güncellemeleri gerektirebilir.
- Önbellekleme stratejisi harici istekleri azaltır ve performansı artırır.


# EN

# Current Rate
## Current Exchange Rate API (Go & Redis)
This project is a Go-based backend service built following Clean Architecture principles. It retrieves exchange and precious metal rate data from TDV, caches it using Redis, and exposes it through a REST API.

### The purpose of this project
- Extract data from an external source (web scraping)
- Process the obtained data and transform it into a domain model
- Establish a performance-focused caching mechanism using Redis
- Build a sustainable backend structure with layered architecture (Clean Architecture)

### Key Features
- Exchange and gold rate data scraping from TDV
- Redis-based caching mechanism
- Time-based cache invalidation strategy
- RESTful API design
- Layered architecture aligned with Clean Architecture
- Centralized error handling structure

### Architecture Overview:
```pgsql
cmd/                    → Application entry point
internal/
   ├── service/         → Business logic layer
   ├── model/           → Domain models
   ├── viewmodel/       → API response models
   └── infrastructure/
           └── cache/   → Redis integration
```
           
### Architectural Principle:
Dependencies flow inward. Business logic remains independent from frameworks and external systems.

### Technologies Used:
- Go
- Fiber (HTTP framework)
- Redis
- Goquery (HTML parsing)
- Viper (configuration management)

### Application Flow:
1. A client sends a request to the API.
2. The service checks Redis for cached rate data.
3. If valid cached data exists, it is returned directly.
4. Otherwise:
    - The TDV source is requested.
    - The HTML table is parsed.
    - Domain model is constructed.
    - Data is stored in Redis.
    - Response is returned to the client.
  
## Setup & Run

### Requirements:
- Go 1.20+
- Redis server

### Run the application:
```bash
go mod tidy
go run cmd/main.go
```

### Default Configuration:
```go
redis.host = localhost
redis.port = 6379
server.port = 3000
```
> **Önemli:** These values can be overridden via environment variables if needed.

### Example Endpoint:
```bash
GET /guncel-kur
```

### Example response:
```json
{
  "dolar": 32.45,
  "euro": 34.12,
  "sterlin": 40.22,
  "gramAltin24Ayar": 2450.15
}
```

> **Notes:**
- The application depends on the HTML structure of the TDV page. Structural changes may require parser updates.
- The caching strategy reduces external requests and improves performance.






