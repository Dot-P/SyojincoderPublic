package controllers

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "github.com/go-sql-driver/mysql"
    "github.com/joho/godotenv"
    "log"
    "sort"
    "unicode"
    "strings"
    "time"
)

func splitWrong(wrong string) (string, string) {
    // 文字列を走査して数字が始まる場所を見つける
    var splitIndex int
    for i, r := range wrong {
        if unicode.IsDigit(r) {
            splitIndex = i
            break
        }
    }

    // 数字の部分とそれ以降を分割
    beforeDigits := wrong[:splitIndex]
    afterDigits := wrong[splitIndex:]

    // '_' を見つけて分割
    parts := strings.SplitN(afterDigits, "_", 2)

    // 分割した文字列を再結合
    result1 := beforeDigits // "abc"
    result2 := parts[1] + parts[0] // "d323"

    return result1, result2
}

func formatInfo(name, category string) string {
    // Category を分割（例: "d303" -> "303", "d"）
    var digitsPart, lettersPart string
    for i, r := range category {
        if unicode.IsDigit(r) {
            digitsPart = category[i:]
            lettersPart = strings.ToUpper(category[:i])
            break
        }
    }

    // Name を大文字にして、フォーマットされた Category と組み合わせる
    // 例: "abc" -> "ABC", "303", "D" -> "ABC303 - D"
    formattedName := strings.ToUpper(name)
    return formattedName + digitsPart + " - " + lettersPart
}

func dotProduct(a, b []float64) float64 {
    var sum float64
    for i := range a {
        sum += a[i] * b[i]
    }
    return sum
}

func connectDB() *sql.DB {

     err := godotenv.Load()
     if err != nil {
         log.Fatal("Error loading .env file")
     }
 
     username := os.Getenv("MYSQL_USER")
     password := os.Getenv("MYSQL_PASSWORD")

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal("LoadLocation is denied", err)
	}
	c := mysql.Config{
		DBName:    "problem",
		User:      username,
		Passwd:    password,
		Addr:      "localhost:3306",
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}

	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		log.Fatal("SQL Open is denied", err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}

func calEmbedding(wrongs []string) ([]float64, []string, []string) {

    var dotProductValues []float64
    var formattedStrings []string
    var processedWrongs []string

    // データベース接続
    db := connectDB()
    defer db.Close()

    err := db.Ping()
    if err != nil {
        log.Fatalf("Database connection is Failed: %v\n", err)
    }

    // トランザクション開始
    tx, err := db.Begin()
    if err != nil {
        log.Fatal("SQL Transaction is denied", err)
    }
    defer tx.Rollback() // エラーがあった場合にロールバック

    type EmbeddingInfo struct {
        Name      string
        Category  string
        Embedding []float64
    }

    // SQLクエリ
    query := `SELECT * FROM embedding WHERE name = ? AND category = ?`

    // データベースから全ての embedding を取得
    allEmbeddingsQuery := `SELECT name, category, embedding FROM embedding`
    allRows, err := db.Query(allEmbeddingsQuery)
    if err != nil {
        log.Fatal("Query execution is denied", err)
    }
    defer allRows.Close()

    var allEmbeddings []EmbeddingInfo
    for allRows.Next() {
        var name, category, embeddingJSON string
        if err := allRows.Scan(&name, &category, &embeddingJSON); err != nil {
            log.Fatal("Failed to scan row", err)
        }
        var embedding []float64
        if err := json.Unmarshal([]byte(embeddingJSON), &embedding); err != nil {
            log.Fatal("Failed to unmarshal JSON", err)
        }
        allEmbeddings = append(allEmbeddings, EmbeddingInfo{Name: name, Category: category, Embedding: embedding})
    }

    for _, wrong := range wrongs {
        w_cat, w_name := splitWrong(wrong)

        // SQLクエリを実行して特定の embedding を取得
        rows, err := tx.Query(query, w_name, w_cat)
        if err != nil {
            log.Fatal("Query execution is denied", err)
        }

        for rows.Next() {
            processedWrongs = append(processedWrongs, wrong)
            processedWrongs = append(processedWrongs, wrong)

            var category string
            var name string
            var embeddingJSON string

            if err := rows.Scan(&category, &name, &embeddingJSON); err != nil {
                log.Fatal("Failed to scan row", err)
            }

            var embedding []float64
            if err := json.Unmarshal([]byte(embeddingJSON), &embedding); err != nil {
                log.Fatal("Failed to unmarshal JSON", err)
            }

            type Result struct {
                Index int
                DotProduct float64
            }

            // 内積を計算して上位2つを選択
            var dotProducts []Result
            for i, embInfo := range allEmbeddings {
                dot := dotProduct(embedding, embInfo.Embedding) // emb2.Embedding を使用
                if 1 - dot > 0.0001 { // 内積が1でない場合のみ追加
                    dotProducts = append(dotProducts, Result{Index: i, DotProduct: dot})
                }
            }            

            // 内積の降順でソート
            sort.Slice(dotProducts, func(i, j int) bool {
                return dotProducts[i].DotProduct > dotProducts[j].DotProduct
            })

            // 上位2つを選択
            if len(dotProducts) > 2 {
                dotProducts = dotProducts[:2]
            }

            for _, dp := range dotProducts {
                embInfo := allEmbeddings[dp.Index]
                combinedInfo := formatInfo(embInfo.Category, embInfo.Name)
                dotProductValues = append(dotProductValues, dp.DotProduct)
                formattedStrings = append(formattedStrings, combinedInfo)
            }                                         
        }
    }

    // トランザクションコミット
    if err := tx.Commit(); err != nil {
        fmt.Println("Error:", err)
    }

    return dotProductValues, formattedStrings, processedWrongs
}