package main

import (
	"encoding/json"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"sync"
)

type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var (
	posts   = []Post{}
	nextID  = 1
	postsMu sync.Mutex
)

func main() {
	http.HandleFunc("/posts", postsHandler)
	http.HandleFunc("/posts/", postHandler)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func postsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getPosts(w, r)
	case http.MethodPost:
		createPost(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(filepath.Base(r.URL.Path))
	if err != nil {
		log.Printf("Invalid ID: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		getPost(w, r, id)
	case http.MethodPut:
		updatePost(w, r, id)
	case http.MethodDelete:
		deletePost(w, r, id)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	postsMu.Lock()
	defer postsMu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	var post Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		log.Printf("Failed to decode post: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	postsMu.Lock()
	defer postsMu.Unlock()

	post.ID = nextID
	nextID++
	posts = append(posts, post)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

func getPost(w http.ResponseWriter, r *http.Request, id int) {
	postsMu.Lock()
	defer postsMu.Unlock()

	for _, post := range posts {
		if post.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(post)
			return
		}
	}

	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func updatePost(w http.ResponseWriter, r *http.Request, id int) {
	var updatedPost Post
	if err := json.NewDecoder(r.Body).Decode(&updatedPost); err != nil {
		log.Printf("Failed to decode updated post: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	postsMu.Lock()
	defer postsMu.Unlock()

	for i, post := range posts {
		if post.ID == id {
			posts[i].Title = updatedPost.Title
			posts[i].Content = updatedPost.Content
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(posts[i])
			return
		}
	}

	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func deletePost(w http.ResponseWriter, r *http.Request, id int) {
	postsMu.Lock()
	defer postsMu.Unlock()

	for i, post := range posts {
		if post.ID == id {
			posts = append(posts[:i], posts[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}
