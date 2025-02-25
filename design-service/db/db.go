package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

const (
	host     = "design-service-db"
	port     = 5432
	user     = "postgres"
	password = "secret"
	dbname   = "design_service"
)

func InitDB() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err.Error())
	}

	DB = db

	createTables()
	seedDatabase()
	fmt.Println("Tables created successfully")

}

func createTables() {
	createCategoriesTable := `
	CREATE TABLE IF NOT EXISTS categories (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL UNIQUE,
		slug TEXT NOT NULL UNIQUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := DB.Exec(createCategoriesTable)
	if err != nil {
		panic("Could not create categories table.")
	}

	createDesignsTable := `
	CREATE TABLE IF NOT EXISTS designs (
		id SERIAL PRIMARY KEY,
		user_id INTEGER NOT NULL,
		title TEXT NOT NULL,
		slug TEXT UNIQUE NOT NULL,
		description TEXT,
		category_id INTEGER,
		tags TEXT,
		visibility TEXT CHECK(visibility IN ('public', 'private', 'unlisted')) DEFAULT 'public',
		status TEXT CHECK(status IN ('draft', 'published')) DEFAULT 'draft',
		views_count INTEGER DEFAULT 0,
		likes_count INTEGER DEFAULT 0,
		comments_count INTEGER DEFAULT 0,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(category_id) REFERENCES categories(id) ON DELETE SET NULL
	);
	`

	_, err = DB.Exec(createDesignsTable)
	if err != nil {
		panic(err.Error())
	}

	createDesignAssetsTable := `
	CREATE TABLE IF NOT EXISTS design_assets (
		id SERIAL PRIMARY KEY,
		design_id INTEGER NOT NULL,
		file_url TEXT NOT NULL,
		file_type TEXT CHECK(file_type IN ('image', 'video', 'pdf')),
		order_index INTEGER DEFAULT 0,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(design_id) REFERENCES designs(id) ON DELETE CASCADE
	);
	`

	_, err = DB.Exec(createDesignAssetsTable)
	if err != nil {
		panic("Could not create design_assets table.")
	}
}

func seedDatabase() {
	// Сидируем категории
	insertCategories := `
	INSERT INTO categories (name, slug) VALUES 
		('Graphics', 'graphics'),
		('Illustrations', 'illustrations'),
		('UI/UX', 'ui-ux')
	ON CONFLICT (slug) DO NOTHING;
	`

	_, err := DB.Exec(insertCategories)
	if err != nil {
		panic("Could not seed categories: " + err.Error())
	}

	// Сидируем дизайны
	insertDesigns := `
	INSERT INTO designs (user_id, title, slug, description, category_id, visibility, status)
	SELECT 1, 'Modern UI Design', 'modern-ui', 'A modern UI design concept', id, 'public', 'published' FROM categories WHERE slug = 'ui-ux'
	ON CONFLICT (slug) DO NOTHING;
	`

	_, err = DB.Exec(insertDesigns)
	if err != nil {
		panic("Could not seed designs: " + err.Error())
	}

	// Сидируем assets
	insertAssets := `
	INSERT INTO design_assets (design_id, file_url, file_type, order_index)
	SELECT id, 'https://example.com/image1.png', 'image', 1 FROM designs WHERE slug = 'modern-ui'
	ON CONFLICT DO NOTHING;
	`

	_, err = DB.Exec(insertAssets)
	if err != nil {
		panic("Could not seed design_assets: " + err.Error())
	}

	fmt.Println("Seeding completed successfully.")
}
