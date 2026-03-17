package models

import (
	"database/sql"
	"time"
)

// BlogPost, veritabanındaki tek bir blog yazısını temsil eder
type BlogPost struct {
	ID        int
	Title     string
	Content   string
	CreatedAt time.Time
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

// ID'sine göre blog yazısını silme
func (m *BlogModel) Delete(id int) error {
	stmt := `DELETE FROM blogs WHERE id = ?`
	_, err := m.DB.Exec(stmt, id)
	return err
}
