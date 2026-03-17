package models

import (
	"database/sql"
	"time"
)

// Comment, veritabanındaki tek bir yorumu temsil eder
type Comment struct {
	ID        int
	BlogID    int
	BlogTitle string // Admin panelinde hangi bloğa ait olduğunu görmek için
	Content   string
	CreatedAt time.Time
}

// BlogPost, veritabanındaki tek bir blog yazısını temsil eder
type BlogPost struct {
	ID        int
	Title     string
	Content   string
	CreatedAt time.Time
	Comments  []*Comment
}

// BlogModel, veritabanı işlemlerini yürütecek yapıdır
type BlogModel struct {
	DB *sql.DB
}

// Yeni bir blog yazısı ekleme (İleride admin panelinde kullanacağız)
func (m *BlogModel) Insert(title, content string) (int, error) {
	// Parametreli sorgu (?) kullanarak SQL Injection saldırılarını kesin olarak engelliyoruz
	stmt := `INSERT INTO blogs (title, content, created_at) VALUES (?, ?, ?)`

	result, err := m.DB.Exec(stmt, title, content, time.Now())
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// ID'ye göre tek bir blog yazısını detaylarıyla çekme (Yeni detay sayfası için)
func (m *BlogModel) Get(id int) (*BlogPost, error) {
	stmt := `SELECT id, title, content, created_at FROM blogs WHERE id = ?`
	b := &BlogPost{}
	err := m.DB.QueryRow(stmt, id).Scan(&b.ID, &b.Title, &b.Content, &b.CreatedAt)
	if err != nil {
		return nil, err
	}
	b.Comments = m.GetCommentsForBlog(b.ID)
	return b, nil
}

// Tüm blog yazılarını çekme (Ana sayfada kullanacağız)
func (m *BlogModel) Latest() ([]*BlogPost, error) {
	stmt := `SELECT id, title, content, created_at FROM blogs ORDER BY created_at DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var blogs []*BlogPost

	for rows.Next() {
		b := &BlogPost{}
		err = rows.Scan(&b.ID, &b.Title, &b.Content, &b.CreatedAt)
		if err != nil {
			return nil, err
		}
		blogs = append(blogs, b)
	}

	return blogs, nil
}

// Tüm blog yazılarını çekme (Admin panelinde tablo olarak listelemek için)
func (m *BlogModel) All() ([]*BlogPost, error) {
	stmt := `SELECT id, title, content, created_at FROM blogs ORDER BY created_at DESC`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var blogs []*BlogPost
	for rows.Next() {
		b := &BlogPost{}
		err = rows.Scan(&b.ID, &b.Title, &b.Content, &b.CreatedAt)
		if err != nil {
			return nil, err
		}
		blogs = append(blogs, b)
	}
	return blogs, nil
}

// ID'ye göre blog yazısını güncelleme
func (m *BlogModel) Update(id int, title, content string) error {
	stmt := `UPDATE blogs SET title = ?, content = ? WHERE id = ?`
	_, err := m.DB.Exec(stmt, title, content, id)
	return err
}

// ID'sine göre blog yazısını silme
func (m *BlogModel) Delete(id int) error {
	// Önce bloğa ait yorumları silelim
	_, _ = m.DB.Exec(`DELETE FROM comments WHERE blog_id = ?`, id)

	stmt := `DELETE FROM blogs WHERE id = ?`
	_, err := m.DB.Exec(stmt, id)
	return err
}

// --- YORUM (COMMENT) FONKSİYONLARI ---

func (m *BlogModel) InsertComment(blogID int, content string) error {
	stmt := `INSERT INTO comments (blog_id, content, created_at) VALUES (?, ?, ?)`
	_, err := m.DB.Exec(stmt, blogID, content, time.Now())
	return err
}

func (m *BlogModel) DeleteComment(id int) error {
	stmt := `DELETE FROM comments WHERE id = ?`
	_, err := m.DB.Exec(stmt, id)
	return err
}

func (m *BlogModel) GetCommentsForBlog(blogID int) []*Comment {
	stmt := `SELECT id, blog_id, content, created_at FROM comments WHERE blog_id = ? ORDER BY created_at ASC`
	rows, err := m.DB.Query(stmt, blogID)
	var comments []*Comment
	if err != nil {
		return comments
	}
	defer rows.Close()
	for rows.Next() {
		c := &Comment{}
		rows.Scan(&c.ID, &c.BlogID, &c.Content, &c.CreatedAt)
		comments = append(comments, c)
	}
	return comments
}

func (m *BlogModel) GetAllComments() ([]*Comment, error) {
	stmt := `SELECT c.id, c.blog_id, c.content, c.created_at, b.title 
			 FROM comments c 
			 JOIN blogs b ON c.blog_id = b.id 
			 ORDER BY c.created_at DESC`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*Comment
	for rows.Next() {
		c := &Comment{}
		err = rows.Scan(&c.ID, &c.BlogID, &c.Content, &c.CreatedAt, &c.BlogTitle)
		if err == nil {
			comments = append(comments, c)
		}
	}
	return comments, nil
}
