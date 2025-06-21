# own_fuff

Bu program, bir hedef URL'deki `"MSEC"` kısmını verilen wordlist dosyasındaki kelimelerle değiştirerek HTTP istekleri gönderen basit bir dizin tarayıcıdır (fuzzer).

---

## Özellikler

- Hedef URL içindeki `"MSEC"` kelimesini, verilen wordlist dosyasındaki kelimelerle sırayla değiştirir.
- HTTP GET istekleri gönderir.
- Eşzamanlı (concurrent) olarak belirtilen sayıda iş parçacığı (goroutine) ile tarama yapar.
- HTTP durum kodlarına göre filtreleme yapabilir (örneğin sadece 200 dönenleri gösterme).
- Parametreler sayesinde esnek kullanım imkanı sağlar.

---

## Kullanım

```bash
./dirscan -u http://example.com/MSEC -w wordlist.txt -t 20 -mc 200


