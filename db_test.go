//go:build persistent
// +build persistent

package main

import (
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func setup() {
	os.Remove("test.db")
	InitDB("test.db")
}

func teardown() {
	os.Remove("test.db")
}

func TestInitDB(t *testing.T) {
	setup()
	defer teardown()

	_, err := os.Stat("test.db")
	if os.IsNotExist(err) {
		t.Fatalf("database file not created")
	}
}

func TestInsertPost(t *testing.T) {
	setup()
	defer teardown()

	post := &Post{Title: "Test Title", Content: "Test Content"}
	err := InsertPost(post)
	if err != nil {
		t.Fatalf("failed to insert post: %v", err)
	}

	if post.ID == 0 {
		t.Fatalf("expected post ID to be set, got %v", post.ID)
	}
}

func TestGetAllPosts(t *testing.T) {
	setup()
	defer teardown()

	post := &Post{Title: "Test Title", Content: "Test Content"}
	InsertPost(post)

	posts, err := GetAllPosts()
	if err != nil {
		t.Fatalf("failed to get posts: %v", err)
	}

	if len(posts) != 1 {
		t.Fatalf("expected 1 post, got %v", len(posts))
	}
}

func TestGetPostByID(t *testing.T) {
	setup()
	defer teardown()

	post := &Post{Title: "Test Title", Content: "Test Content"}
	InsertPost(post)

	fetchedPost, err := GetPostByID(post.ID)
	if err != nil {
		t.Fatalf("failed to get post: %v", err)
	}

	if fetchedPost == nil || fetchedPost.ID != post.ID {
		t.Fatalf("expected to fetch post with ID %v, got %v", post.ID, fetchedPost)
	}
}

func TestUpdatePostByID(t *testing.T) {
	setup()
	defer teardown()

	post := &Post{Title: "Test Title", Content: "Test Content"}
	InsertPost(post)

	updatedPost := &Post{Title: "Updated Title", Content: "Updated Content"}
	err := UpdatePostByID(post.ID, updatedPost)
	if err != nil {
		t.Fatalf("failed to update post: %v", err)
	}

	fetchedPost, err := GetPostByID(post.ID)
	if err != nil {
		t.Fatalf("failed to get post: %v", err)
	}

	if fetchedPost.Title != updatedPost.Title || fetchedPost.Content != updatedPost.Content {
		t.Fatalf("expected post to be updated to %v, got %v", updatedPost, fetchedPost)
	}
}
