package main

import (
	"fmt"
	"net/http"
	"os"

	zone "zone/handlers"
)

func main() {
	// جلب الفنانين
	var err error
	zone.AllArtists, err = zone.FetchArtists()
	if err != nil {
		fmt.Println("Error fetching artists:", err)
		return
	}

	// إعداد المسارات
	http.HandleFunc("/static/", zone.HandleStatic)
	http.HandleFunc("/", zone.HandlerHome)
	http.HandleFunc("/artist/", zone.HandlerArtist)

	// استخدام PORT من البيئة (مهم للنشر على منصات مثل Render أو Replit)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // الوضع المحلي
	}

	fmt.Printf("Server running on http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

