package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
)

const (
	host      = "db"
	user      = "postgres-dev"
	port      = 5432
	password  = "mysecretpassword"
	dbname    = "clrpix_dev"
	Migration = `CREATE TABLE IF NOT EXISTS colorpix(
	id serial PRIMARY KEY,
	pictype text NOT NULL,
	picowner text NOT NULL,
	created_at timestamp with time zone DEFAULT current_timestamp)
	`
)

type ColorPix struct {
	PicType   string    `json:"pictype" binding:"required"`
	PicOwner  string    `json:"picowner" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}

var db *sql.DB


func GetPix() ([]ColorPix, error) {
	const q = `SELECT pictype, picowner, created_at FROM colorpix`
	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	/*
	var count int
	row := db.QueryRow(`SELECT COUNT(*) FROM colorpix`)
	err = row.Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("count:",count)
	*/
	results := make([]ColorPix, 0)
	fmt.Println("Rows.Next:",rows.Next())
	for rows.Next() {
		var pictype string
		var picowner string
		var createdAt time.Time
		if err := rows.Scan(&pictype, &picowner, &createdAt); err != nil {
			return nil, err
		}
		results = append(results, ColorPix{pictype, picowner, createdAt,})
	}
	return results, nil

}

func AddPix(p ColorPix) error {
	const g = `INSERT INTO colorpix(pictype,picowner,created_at) VALUES ($1,$2,$3)`
	_, err := db.Exec(g, p.PicType, p.PicOwner, p.CreatedAt)

	return err

}

func main() {

	var err error

	r := gin.Default()
	r.GET("/colorpix", func(ctx *gin.Context) {
		results, err := GetPix()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "internal error:" + err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, results)
	})

	r.POST("/colorpix", func(ctx *gin.Context) {
		var p ColorPix
		if ctx.Bind(&p) == nil {
			p.CreatedAt = time.Now()
			if err = AddPix(p); err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"status": "internal error:" + err.Error()})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
		}

	})

	dbInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s", host, port, user, dbname, password)
	db, err = sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Query(Migration)
	if err != nil {
		log.Println("failed to run migrations", err.Error())
		return
	}
	log.Println("running")
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
