package repository

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/imorti/crash/rest-api/entity"
	"google.golang.org/api/iterator"
)

// PostRepository - interface
type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type repo struct{}

// NewPostRepository - method for adding repository
func NewPostRepository() PostRepository {
	return &repo{}
}

const (
	projectID      string = "pragmatic-crash"
	collectionName string = "posts"
)

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)

	if err != nil {
		log.Fatalf("Failed to create new firestore client: %v", err)
		return nil, err
	}
	defer client.Close()
	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})
	if err != nil {
		log.Fatalf("Failed to add new post %v", err)
		return nil, err
	}
	return post, nil

}

func (*repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)

	if err != nil {
		log.Fatalf("Failed to create new firestore client: %v", err)
		return nil, err
	}
	defer client.Close()
	var posts []entity.Post
	it := client.Collection(collectionName).Documents(ctx)

	for {

		doc, err := it.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Fatalf("Failed to iterate through collection: %v", err)
			return nil, err
		}
		post := entity.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}
	return posts, nil
}
