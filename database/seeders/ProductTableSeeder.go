package seeders

import (
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v3"
)

func (s DatabaseSeeder) ProductSeedTableSeeder() {
	for i := 0; i < 100; i++ {
		stmt, _ := s.db.Prepare(`INSERT INTO products(user_id, name, description, price, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`)
		_, err := stmt.Exec((rand.Intn(100-1) + 1), faker.Word(), faker.Paragraph(), (rand.Intn(1000000-10000) + 10000), time.Now(), time.Now())
		if err != nil {
			panic(err)
		}
	}
}
