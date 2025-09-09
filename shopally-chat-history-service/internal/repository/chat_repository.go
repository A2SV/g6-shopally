package repository

import (
	"context"
	"fmt"
	"log" // Using standard log for error logging as no custom logger is injected into the struct
	"time"

	"github.com/shopally/chat-history/internal/domain"
	"github.com/shopally/chat-history/internal/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ChatRepository defines the interface for database operations related to chat history.
type ChatRepository interface {
	FindByUserEmail(ctx context.Context, userEmail string) (*domain.ChatHistory, error)
	PushChatSession(ctx context.Context, userEmail string, chatSession domain.ChatSession) error
	PushMessageToSession(ctx context.Context, userEmail, chatID string, message domain.Message) error
	PullChatSession(ctx context.Context, userEmail, chatID string) error
}

// mongoChatRepository implements ChatRepository for MongoDB.
type mongoChatRepository struct {
	collection *mongo.Collection
}

// NewMongoChatRepository creates a new MongoDB implementation of ChatRepository.
func NewMongoChatRepository(collection *mongo.Collection) ChatRepository {
	return &mongoChatRepository{
		collection: collection,
	}
}

// FindByUserEmail retrieves the entire ChatHistory document for a given user email.
func (r *mongoChatRepository) FindByUserEmail(ctx context.Context, userEmail string) (*domain.ChatHistory, error) {
	var chatHistory domain.ChatHistory
	filter := bson.M{"user_email": userEmail}

	err := r.collection.FindOne(ctx, filter).Decode(&chatHistory)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("Repository: No chat history found for user email %s", userEmail)
			return nil, errors.ErrNotFound // Translate MongoDB's "no documents" error
		}
		log.Printf("Repository: MongoDB error finding chat history for %s: %v", userEmail, err)
		return nil, fmt.Errorf("repository error finding chat history: %w", err)
	}
	return &chatHistory, nil
}

// PushChatSession adds a new chat session to a user's chat history document.
// It creates the top-level document if it doesn't exist (upsert: true).
func (r *mongoChatRepository) PushChatSession(ctx context.Context, userEmail string, chatSession domain.ChatSession) error {
	now := primitive.NewDateTimeFromTime(time.Now())

	filter := bson.M{"user_email": userEmail}
	update := bson.M{
		"$push": bson.M{"chat_sessions": chatSession}, // Add the new chat session to the array
		"$set":  bson.M{"last_activity": now},         // Update last activity timestamp
		"$setOnInsert": bson.M{ // These fields are set only if a new document is inserted
			"user_email": userEmail,
		},
	}
	opts := options.Update().SetUpsert(true) // Crucial: create document if it doesn't exist

	_, err := r.collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		log.Printf("Repository: MongoDB error pushing chat session for %s: %v", userEmail, err)
		return fmt.Errorf("repository error pushing chat session: %w", err)
	}
	return nil
}

// PushMessageToSession adds a new message turn to a specific chat session within a user's history.
func (r *mongoChatRepository) PushMessageToSession(ctx context.Context, userEmail, chatID string, message domain.Message) error {
	now := primitive.NewDateTimeFromTime(time.Now())

	filter := bson.M{
		"user_email":            userEmail,
		"chat_sessions.chat_id": chatID, // Find the document and the specific chat session within it
	}
	update := bson.M{
		"$push": bson.M{"chat_sessions.$.messages": message}, // Add the new message to the messages array of the matched session
		"$set": bson.M{
			"last_activity":                now, // Update top-level activity
			"chat_sessions.$.last_updated": now, // Update specific chat session's last updated timestamp
		},
	}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Printf("Repository: MongoDB error pushing message to chat %s for %s: %v", chatID, userEmail, err)
		return fmt.Errorf("repository error pushing message to session: %w", err)
	}
	if result.ModifiedCount == 0 {
		// This implies either user_email or chat_id was not found
		log.Printf("Repository: PushMessageToSession: Chat session %s not found for user %s or no modification made.", chatID, userEmail)
		return errors.ErrNotFound
	}
	return nil
}

// PullChatSession removes a specific chat session from a user's history.
func (r *mongoChatRepository) PullChatSession(ctx context.Context, userEmail, chatID string) error {
	now := primitive.NewDateTimeFromTime(time.Now())

	filter := bson.M{"user_email": userEmail}
	update := bson.M{
		"$pull": bson.M{"chat_sessions": bson.M{"chat_id": chatID}}, // Remove the chat session with matching chat_id
		"$set":  bson.M{"last_activity": now},                       // Update top-level activity
	}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Printf("Repository: MongoDB error pulling chat session %s for %s: %v", chatID, userEmail, err)
		return fmt.Errorf("repository error pulling chat session: %w", err)
	}
	if result.ModifiedCount == 0 {
		// This implies either user_email or chat_id was not found
		log.Printf("Repository: PullChatSession: Chat session %s not found for user %s or no modification made.", chatID, userEmail)
		return errors.ErrNotFound
	}
	return nil
}
