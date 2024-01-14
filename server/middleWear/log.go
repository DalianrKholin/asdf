package middleWear

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func OpenAppendOnlyFile(name string) (*os.File, error) {
	return os.OpenFile(name, os.O_APPEND, 0755)
}

func (s *ApiDbMiddleWear) SaveData(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
		go saveChanges(r)
	}
}

func saveChanges(r *http.Request) {
	name := r.Header.Get("name")
	fmt.Printf("--%v--", r.Header.Get("name"))
	if r.Header.Get("admin") == "true" {
		name = "admin"
	}
	fileName := "C:\\Users\\osado\\GolandProjects\\niceSite\\serverLog"
	file, err := OpenAppendOnlyFile(fileName)
	if err != nil {
		fmt.Printf("nie udało się zapisać danych zapytania")
	}
	toSaveString := time.Now().Format("2006-01-02 15:04:05") + " : " + r.Method + " logged User: " + name + "\n"
	file.Write([]byte(toSaveString))
	defer file.Close()

}
