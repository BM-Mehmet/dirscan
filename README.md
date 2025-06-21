# dirscan

Bu program, bir hedef URL'deki `"MSEC"` kısmını verilen wordlist dosyasındaki kelimelerle değiştirerek HTTP istekleri gönderen basit bir dizin tarayıcıdır (fuzzer).

---

## Özellikler

- Hedef URL içindeki `"MSEC"` kelimesini, verilen wordlist dosyasındaki kelimelerle sırayla değiştirir.
- HTTP GET istekleri gönderir.
- Eşzamanlı (concurrent) olarak belirtilen sayıda iş parçacığı (goroutine) ile tarama yapar.
- HTTP durum kodlarına göre filtreleme yapabilir (örneğin sadece 200 dönenleri gösterme).
- Parametreler sayesinde esnek kullanım imkanı sağlar.
- Subdomain taraması için de kullanılabilir `"MSEC"` yerini değiştirmek yeterlidir.

---

## Kullanım

```bash
# Dizin/endpoint taraması
./dirscan -u http://example.com/MSEC -w wordlist.txt -t 50 -mc 200

# Subdomain taraması
./dirscan -u http://MSEC.example.com -w subdomains.txt -t 30

