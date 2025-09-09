package repository

import (
	"context"
	"fmt" // Added for fmt.Printf
	"log"
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
	PushChatSession(ctx context.Context, userEmail string, chatSession domain.ChatSession) (string, error)
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
			return nil, errors.ErrNotFound
		}
		log.Printf("Repository: MongoDB error finding chat history for %s: %v", userEmail, err)
		return nil, fmt.Errorf("repository error finding chat history: %w", err)
	}
	return &chatHistory, nil
}

// PushChatSession adds a new chat session to a user's chat history document.
// It creates the top-level document if it doesn't exist (upsert: true).
func (r *mongoChatRepository) PushChatSession(ctx context.Context, userEmail string, chatSession domain.ChatSession) (string, error) {
	now := primitive.NewDateTimeFromTime(time.Now())

	filter := bson.M{"user_email": userEmail}
	update := bson.M{
		"$push": bson.M{"chat_sessions": chatSession},
		"$set":  bson.M{"last_activity": now},
		"$setOnInsert": bson.M{
			"user_email": userEmail,
		},
	}
	opts := options.Update().SetUpsert(true)

	_, err := r.collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		log.Printf("Repository: MongoDB error pushing chat session for %s: %v", userEmail, err)
		return "", fmt.Errorf("repository error pushing chat session: %w", err)
	}
	return chatSession.ChatID, nil
}

// PushMessageToSession adds a new message turn to a specific chat session within a user's history.
func (r *mongoChatRepository) PushMessageToSession(ctx context.Context, userEmail, chatID string, message domain.Message) error {
	now := primitive.NewDateTimeFromTime(time.Now())

	filter := bson.M{
		"user_email":            userEmail,
		"chat_sessions.chat_id": chatID,
	}
	update := bson.M{
		"$push": bson.M{"chat_sessions.$.messages": message},
		"$set": bson.M{
			"last_activity":                now,
			"chat_sessions.$.last_updated": now,
		},
	}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Printf("Repository: MongoDB error pushing message to chat %s for %s: %v", chatID, userEmail, err)
		return fmt.Errorf("repository error pushing message to session: %w", err)
	}
	if result.ModifiedCount == 0 {
		log.Printf("Repository: PushMessageToSession: Chat session %s not found for user %s or no modification made.", chatID, userEmail)
		return errors.ErrNotFound
	}
	return nil
}

// PullChatSession removes a specific chat session from a user's history.
func (r *mongoChatRepository) PullChatSession(ctx context.Context, userEmail, chatID string) error {
	now := primitive.NewDateTimeFromTime(time.Now())

	// Fixed: Filter should include both user_email AND the specific chat session
	filter := bson.M{
		"user_email":            userEmail,
		"chat_sessions.chat_id": chatID, // This ensures we only match documents that have this chat session
	}

	update := bson.M{
		"$pull": bson.M{"chat_sessions": bson.M{"chat_id": chatID}},
		"$set":  bson.M{"last_activity": now},
	}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println(ctx, "MongoDB error pulling chat session %s for %s: %v", chatID, userEmail, err)
		return fmt.Errorf("repository error pulling chat session: %w", err)
	}

	if result.ModifiedCount == 0 {
		log.Println(ctx, "Chat session %s not found for user %s", chatID, userEmail)
		return errors.ErrNotFound
	}

	log.Println(ctx, "Successfully removed chat session %s for user %s", chatID, userEmail)
	return nil
}
