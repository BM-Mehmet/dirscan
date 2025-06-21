# own_fuff
/*
Bu program, bir hedef URL'deki "MSEC" kısmını verilen wordlist dosyasındaki kelimelerle değiştirerek HTTP istekleri gönderir.
Kullanıcıdan hedef URL (-u), wordlist dosya yolu (-w), eşzamanlı iş parçacığı sayısı (-t) ve eşleşecek HTTP durum kodu (-mc) parametrelerini alır.
Wordlist dosyasındaki her kelime için, URL'deki "MSEC" kısmı o kelimeyle değiştirilir ve HTTP GET isteği yapılır.
Eğer -mc parametresi verilmemişse (0 ise), tüm durum kodları ekrana yazdırılır. Eğer verilmişse, sadece o durum koduna sahip yanıtlar ekrana yazdırılır.
İşlemler, belirtilen sayıda eşzamanlı iş parçacığı (goroutine) ile gerçekleştirilir.
*/
