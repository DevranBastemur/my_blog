# DevRyanBlog - Güvenli Kişisel Blog Platformu

DevRyanBlog, Go (Golang) ve SQLite kullanılarak defansif programlama (Secure Coding) prensipleriyle geliştirilmiş, siber güvenlik temalı kişisel bir blog sistemidir. "Terminal UI" ve "Cyberpunk" estetiğini barındıran bu proje, hem güvenli bir arka plan mimarisine hem de kullanıcı dostu bir yönetim paneline sahiptir.

<img width="1807" height="846" alt="Ekran görüntüsü 2026-03-17 224624" src="https://github.com/user-attachments/assets/b631c364-aafa-45b1-b2f4-4dc7621ea7ac" />


<img width="1864" height="924" alt="Ekran görüntüsü 2026-03-17 224641" src="https://github.com/user-attachments/assets/28bcedb6-11fc-4cbd-8f7a-5a8a07650aa7" />



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

<img width="1074" height="739" alt="Ekran görüntüsü 2026-03-17 224721" src="https://github.com/user-attachments/assets/546ad55b-38ec-44a2-b570-907dc7e227fd" />


### 2. Yönetim Paneli (Admin Dashboard)
Yöneticinin içeriklere müdahale ettiği, sildiği veya güncellediği güvenli alan.
<img width="1770" height="742" alt="Ekran görüntüsü 2026-03-17 224744" src="https://github.com/user-attachments/assets/0216b8f9-19dc-4b17-9294-7faedc9f3eab" />

- **İçerik ve Yorum Yönetimi:** Sistemdeki tüm yazıları ve kullanıcı yorumlarını tek ekrandan kontrol etme.
<img width="1142" height="743" alt="Ekran görüntüsü 2026-03-17 224808" src="https://github.com/user-attachments/assets/5d8bbb88-ab9a-4b77-8b25-05a974ab8277" />

### 3. Güvenlik ve Kimlik Doğrulama
Sistemin kalbi olan admin girişi, yetkisiz erişimlere karşı katı kurallarla korunmaktadır.

- **Kullanıcı Adı:** `admin`
- **Parola:** `admin123`
- **Koruma Kalkanı:** 5 hatalı deneme sonrası IP/Oturum 20 dakika boyunca engellenir.

<img width="790" height="674" alt="Ekran görüntüsü 2026-03-17 224904" src="https://github.com/user-attachments/assets/e09f8bf2-727b-4cfa-befa-ca3dfc0df290" />

---

## 🛠️ Kurulum ve Çalıştırma

Projeyi yerel ortamınızda çalıştırmak için aşağıdaki adımları izleyebilirsiniz.

1. **Projeyi Klonlayın:**
   ```bash
  git clone [https://github.com/kullaniciadi/devryanblog.git](https://github.com/kullaniciadi/devryanblog.git)
 cd devryanblog

