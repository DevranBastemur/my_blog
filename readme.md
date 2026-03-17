# DevRyanBlog - Güvenli Kişisel Blog Platformu

DevRyanBlog, Go (Golang) ve SQLite kullanılarak defansif programlama (Secure Coding) prensipleriyle geliştirilmiş, siber güvenlik temalı kişisel bir blog sistemidir. "Terminal UI" ve "Cyberpunk" estetiğini barındıran bu proje, hem güvenli bir arka plan mimarisine hem de kullanıcı dostu bir yönetim paneline sahiptir.

![Ana Sayfa Görünümü](Ekran görüntüsü 2026-03-17 224641.jpg)
*(Yukarıdaki alana ana sayfa blog kartlarını gösteren görseli ekleyebilirsin)*

## 🚀 Temel Özellikler

- **Güvenli Mimarisi:** XSS ve SQL Injection gibi temel web zafiyetlerine karşı Go'nun yerleşik güvenlik önlemleriyle donatılmıştır.
- **Siber Temalı Arayüz:** Zifiri siyah arka plan, neon kırmızı vurgular ve terminal esintili tipografi ile tam bir Blue Team konsepti.
- **Kapsamlı Admin Paneli:** Yeni blog yazıları yayınlama, kapak görselleri ekleme ve mevcut içerikleri yönetme.
- **Güvenli Yorum Sistemi:** Her yazıya özel, 250 karakter sınırıyla sınırlandırılmış ve admin tarafından yönetilebilen dinamik yorum altyapısı.
- **Brute-Force Koruması:** Yönetici paneline yönelik kaba kuvvet saldırılarını engellemek için 5 hatalı girişte sistemi 20 dakika boyunca kilitleyen mekanizma.

---

## 📸 Ekran Görüntüleri ve Modüller

### 1. Dinamik İçerik ve Yorum Sistemi
Kullanıcılar blog yazılarını okuyabilir ve iz bırakabilir (yorum yapabilir). Yorumlar 250 karakter ile sınırlandırılmıştır ve anlık olarak sisteme kaydedilir.

![Yorum Sistemi](Ekran görüntüsü 2026-03-17 224721.png)
*(Yukarıdaki alana yorum satırını ve "Ağ iz bırak" kısmını gösteren görseli ekleyebilirsin)*

### 2. Yönetim Paneli (Admin Dashboard)
Yöneticinin içeriklere müdahale ettiği, sildiği veya güncellediği güvenli alan.

- **İçerik Ekleme:** Başlık, metin ve kapak görseli (Max 5MB) yükleme desteği.
![Yeni Yazı Ekleme](Ekran görüntüsü 2026-03-17 224757.png)

- **İçerik ve Yorum Yönetimi:** Sistemdeki tüm yazıları ve kullanıcı yorumlarını tek ekrandan kontrol etme.
![İçerik Yönetimi](Ekran görüntüsü 2026-03-17 224808.png)

### 3. Güvenlik ve Kimlik Doğrulama
Sistemin kalbi olan admin girişi, yetkisiz erişimlere karşı katı kurallarla korunmaktadır.

- **Kullanıcı Adı:** `admin`
- **Parola:** `admin123`
- **Koruma Kalkanı:** 5 hatalı deneme sonrası IP/Oturum 20 dakika boyunca engellenir.

![Admin Login](# DevRyanBlog - Güvenli Kişisel Blog Platformu

DevRyanBlog, Go (Golang) ve SQLite kullanılarak defansif programlama (Secure Coding) prensipleriyle geliştirilmiş, siber güvenlik temalı kişisel bir blog sistemidir. "Terminal UI" ve "Cyberpunk" estetiğini barındıran bu proje, hem güvenli bir arka plan mimarisine hem de kullanıcı dostu bir yönetim paneline sahiptir.

![Ana Sayfa Görünümü](Ekran görüntüsü 2026-03-17 224641.jpg)
*(Yukarıdaki alana ana sayfa blog kartlarını gösteren görseli ekleyebilirsin)*

## 🚀 Temel Özellikler

- **Güvenli Mimarisi:** XSS ve SQL Injection gibi temel web zafiyetlerine karşı Go'nun yerleşik güvenlik önlemleriyle donatılmıştır.
- **Siber Temalı Arayüz:** Zifiri siyah arka plan, neon kırmızı vurgular ve terminal esintili tipografi ile tam bir Blue Team konsepti.
- **Kapsamlı Admin Paneli:** Yeni blog yazıları yayınlama, kapak görselleri ekleme ve mevcut içerikleri yönetme.
- **Güvenli Yorum Sistemi:** Her yazıya özel, 250 karakter sınırıyla sınırlandırılmış ve admin tarafından yönetilebilen dinamik yorum altyapısı.
- **Brute-Force Koruması:** Yönetici paneline yönelik kaba kuvvet saldırılarını engellemek için 5 hatalı girişte sistemi 20 dakika boyunca kilitleyen mekanizma.

---

## 📸 Ekran Görüntüleri ve Modüller

### 1. Dinamik İçerik ve Yorum Sistemi
Kullanıcılar blog yazılarını okuyabilir ve iz bırakabilir (yorum yapabilir). Yorumlar 250 karakter ile sınırlandırılmıştır ve anlık olarak sisteme kaydedilir.

![Yorum Sistemi](Ekran görüntüsü 2026-03-17 224721.png)
*(Yukarıdaki alana yorum satırını ve "Ağ iz bırak" kısmını gösteren görseli ekleyebilirsin)*

### 2. Yönetim Paneli (Admin Dashboard)
Yöneticinin içeriklere müdahale ettiği, sildiği veya güncellediği güvenli alan.

- **İçerik Ekleme:** Başlık, metin ve kapak görseli (Max 5MB) yükleme desteği.
![Yeni Yazı Ekleme](Ekran görüntüsü 2026-03-17 224757.png)

- **İçerik ve Yorum Yönetimi:** Sistemdeki tüm yazıları ve kullanıcı yorumlarını tek ekrandan kontrol etme.
![İçerik Yönetimi](Ekran görüntüsü 2026-03-17 224808.png)

### 3. Güvenlik ve Kimlik Doğrulama
Sistemin kalbi olan admin girişi, yetkisiz erişimlere karşı katı kurallarla korunmaktadır.

- **Kullanıcı Adı:** `admin`
- **Parola:** `admin123`
- **Koruma Kalkanı:** 5 hatalı deneme sonrası IP/Oturum 20 dakika boyunca engellenir.

![Admin Login] <img width="790" height="674" alt="Ekran görüntüsü 2026-03-17 224904" src="https://github.com/user-attachments/assets/2d50f06a-096f-4b8e-992c-3a80e1c12c50" />
 
*(Yukarıdaki alana "Hatalı giriş! Kalan deneme hakkı: 4" uyarısını gösteren login ekranını ekleyebilirsin)*

---

## 🛠️ Kurulum ve Çalıştırma

Projeyi yerel ortamınızda çalıştırmak için aşağıdaki adımları izleyebilirsiniz.

1. **Projeyi Klonlayın:**
   ```bash
   git clone [https://github.com/kullaniciadi/devryanblog.git](https://github.com/kullaniciadi/devryanblog.git)
   cd devryanblog)
*(Yukarıdaki alana "Hatalı giriş! Kalan deneme hakkı: 4" uyarısını gösteren login ekranını ekleyebilirsin)*

---

## 🛠️ Kurulum ve Çalıştırma

Projeyi yerel ortamınızda çalıştırmak için aşağıdaki adımları izleyebilirsiniz.

1. **Projeyi Klonlayın:**
   ```bash
   git clone [https://github.com/kullaniciadi/devryanblog.git](https://github.com/kullaniciadi/devryanblog.git)
   cd devryanblog
