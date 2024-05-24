package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func clearData() {
	posts = []Post{}
	nextID = 1
}

func TestGetPosts(t *testing.T) {
	persistent = false // Use in-memory storage for testing
	clearData()        // Clear any existing data

	req, err := http.NewRequest("GET", "/posts", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(postsHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := []Post{}
	var actual []Post
	if err := json.NewDecoder(rr.Body).Decode(&actual); err != nil {
		t.Fatalf("could not decode response: %v", err)
	}

	if len(actual) != len(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestCreatePost(t *testing.T) {
	persistent = false // Use in-memory storage for testing
	clearData()        // Clear any existing data

	post := &Post{Title: "Test Title", Content: "Test Content"}
	body, err := json.Marshal(post)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/posts", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(postsHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	var createdPost Post
	if err := json.NewDecoder(rr.Body).Decode(&createdPost); err != nil {
		t.Fatal(err)
	}

	if createdPost.Title != post.Title || createdPost.Content != post.Content {
		t.Errorf("handler returned unexpected body: got %v want %v", createdPost, post)
	}

	if createdPost.ID != 1 {
		t.Errorf("expected post ID to be 1, got %v", createdPost.ID)
	}
}

func TestGetPost(t *testing.T) {
	persistent = false // Use in-memory storage for testing
	clearData()        // Clear any existing data

	// First, create a post to retrieve
	post := &Post{Title: "Test Title", Content: "Test Content"}
	post.ID = nextID
	nextID++
	posts = append(posts, *post)

	req, err := http.NewRequest("GET", "/posts/"+strconv.Itoa(post.ID), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(postHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var fetchedPost Post
	if err := json.NewDecoder(rr.Body).Decode(&fetchedPost); err != nil {
		t.Fatal(err)
	}

	if fetchedPost.Title != post.Title || fetchedPost.Content != post.Content {
		t.Errorf("handler returned unexpected body: got %v want %v", fetchedPost, post)
	}
}

func TestUpdatePost(t *testing.T) {
	persistent = false // Use in-memory storage for testing
	clearData()        // Clear any existing data

	// First, create a post to update
	post := &Post{Title: "Test Title", Content: "Test Content"}
	post.ID = nextID
	nextID++
	posts = append(posts, *post)

	updatedPost := &Post{Title: "Updated Title", Content: "Updated Content"}
	body, err := json.Marshal(updatedPost)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", "/posts/"+strconv.Itoa(post.ID), bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(postHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var fetchedPost Post
	if err := json.NewDecoder(rr.Body).Decode(&fetchedPost); err != nil {
		t.Fatal(err)
	}

	if fetchedPost.Title != updatedPost.Title || fetchedPost.Content != updatedPost.Content {
		t.Errorf("handler returned unexpected body: got %v want %v", fetchedPost, updatedPost)
	}
}

func TestDeletePost(t *testing.T) {
	persistent = false // Use in-memory storage for testing
	clearData()        // Clear any existing data

	// First, create a post to delete
	post := &Post{Title: "Test Title", Content: "Test Content"}
	post.ID = nextID
	nextID++
	posts = append(posts, *post)

	req, err := http.NewRequest("DELETE", "/posts/"+strconv.Itoa(post.ID), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(postHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
	}

	// Ensure the post was deleted
	for _, p := range posts {
		if p.ID == post.ID {
			t.Errorf("post with ID %v was not deleted", post.ID)
		}
	}
}
