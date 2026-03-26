# DevRyanBlog - Güvenli Kişisel Blog Platformu

DevRyanBlog, Go (Golang) ve SQLite kullanılarak defansif programlama (Secure Coding) prensipleriyle geliştirilmiş, siber güvenlik temalı kişisel bir blog sistemidir. "Terminal UI" ve "Cyberpunk" estetiğini barındıran bu proje, hem güvenli bir arka plan mimarisine hem de kullanıcı dostu bir yönetim paneline sahiptir.

![Ana Sayfa Görünümü] <img width="1807" height="846" alt="Ekran görüntüsü 2026-03-17 224624" src="https://github.com/user-attachments/assets/d1db843e-ad7b-48be-a8d2-30f2fcc26c21" />

<img width="1864" height="924" alt="Ekran görüntüsü 2026-03-17 224641" src="https://github.com/user-attachments/assets/ba2296b3-1518-4a5c-9a6c-ed828e10668c" />


##  Temel Özellikler

- **Güvenli Mimarisi:** XSS ve SQL Injection gibi temel web zafiyetlerine karşı Go'nun yerleşik güvenlik önlemleriyle donatılmıştır.
- **Siber Temalı Arayüz:** Zifiri siyah arka plan, neon kırmızı vurgular ve terminal esintili tipografi ile tam bir Blue Team konsepti.
- **Kapsamlı Admin Paneli:** Yeni blog yazıları yayınlama, kapak görselleri ekleme ve mevcut içerikleri yönetme.
- **Güvenli Yorum Sistemi:** Her yazıya özel, 250 karakter sınırıyla sınırlandırılmış ve admin tarafından yönetilebilen dinamik yorum altyapısı.
- **Brute-Force Koruması:** Yönetici paneline yönelik kaba kuvvet saldırılarını engellemek için 5 hatalı girişte sistemi 20 dakika boyunca kilitleyen mekanizma.

---

## 📸 Ekran Görüntüleri ve Modüller

### 1. Dinamik İçerik ve Yorum Sistemi
Kullanıcılar blog yazılarını okuyabilir ve iz bırakabilir (yorum yapabilir). Yorumlar 250 karakter ile sınırlandırılmıştır ve anlık olarak sisteme kaydedilir.

![Yorum Sistemi] <img width="1074" height="739" alt="Ekran görüntüsü 2026-03-17 224721" src="https://github.com/user-attachments/assets/929efb57-eab4-4a21-b632-e8bb61bc2fb1" />


### 2. Yönetim Paneli (Admin Dashboard)
Yöneticinin içeriklere müdahale ettiği, sildiği veya güncellediği güvenli alan.

- **İçerik Ekleme:** Başlık, metin ve kapak görseli (Max 5MB) yükleme desteği.
![Yeni Yazı Ekleme] <img width="1770" height="742" alt="Ekran görüntüsü 2026-03-17 224744" src="https://github.com/user-attachments/assets/fd6da53f-f69e-4f9c-ba01-8225923430b2" />


- **İçerik ve Yorum Yönetimi:** Sistemdeki tüm yazıları ve kullanıcı yorumlarını tek ekrandan kontrol etme.
![İçerik Yönetimi] <img width="1142" height="743" alt="Ekran görüntüsü 2026-03-17 224808" src="https://github.com/user-attachments/assets/0a5862a6-7e30-42d0-a2c4-eb753291793d" />


### 3. Güvenlik ve Kimlik Doğrulama
Sistemin kalbi olan admin girişi, yetkisiz erişimlere karşı katı kurallarla korunmaktadır.

- **Kullanıcı Adı:** `admin`
- **Parola:** `admin123`
- **Koruma Kalkanı:** 5 hatalı deneme sonrası IP/Oturum 20 dakika boyunca engellenir.

![Admin Login] <img width="790" height="674" alt="Ekran görüntüsü 2026-03-17 224904" src="https://github.com/user-attachments/assets/8e7e4b47-b5e6-4c11-9b86-ea8d188f8dbe" />


---

##  Kurulum ve Çalıştırma

Projeyi yerel ortamınızda çalıştırmak için aşağıdaki adımları izleyebilirsiniz.

1. **Projeyi Klonlayın:**
   ```bash
   git clone [https://github.com/devranbastemur/devryanblog.git](https://github.com/devranbastemur/devryanblog.git)
   cd devryanblog
