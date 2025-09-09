package integration_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/shopally/chat-history/internal/domain"
	"github.com/shopally/chat-history/internal/errors"
	"github.com/shopally/chat-history/internal/platform/mongodb"
	"github.com/shopally/chat-history/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Global variables for test setup
var (
	testMongoClient    *mongo.Client
	testChatCollection *mongo.Collection
	testChatRepository domain.ChatRepository
	testDBName         = "shopally_chat_test_db"
	testCollectionName = "chat_history_test_repo"
)

// TestMain sets up and tears down the test environment (MongoDB connection, repository initialization)
func TestMain(m *testing.M) {
	// Setup MongoDB connection for integration tests
	mongoURI := os.Getenv("MONGO_TEST_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017" // Default test MongoDB URI
	}

	var err error
	testMongoClient, err = mongodb.Connect(mongoURI, testDBName, testCollectionName)
	if err != nil {
		log.Fatalf("Failed to connect to test MongoDB: %v", err)
	}
	defer func() {
		if err := mongodb.Disconnect(testMongoClient); err != nil {
			log.Printf("Error disconnecting from test MongoDB: %v", err)
		}
	}()

	testChatCollection = testMongoClient.Database(testDBName).Collection(testCollectionName)
	testChatRepository = repository.NewMongoChatRepository(testChatCollection)

	// Clean up previous test data before running tests
	if err := testChatCollection.Drop(context.Background()); err != nil {
		log.Printf("Warning: Failed to drop test collection before tests (might not exist): %v", err)
	}

	// Run tests
	code := m.Run()

	// Clean up test data after all tests
	if err := testChatCollection.Drop(context.Background()); err != nil {
		log.Printf("Error dropping test collection after tests: %v", err)
	}

	os.Exit(code)
}

// Helper to clean collection before each test
func beforeEachRepoTest(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := testChatCollection.Drop(ctx); err != nil {
		t.Fatalf("Failed to drop collection before test: %v", err)
	}
	// Re-create the unique index, as Drop() removes it
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "user_email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	_, err := testChatCollection.Indexes().CreateOne(ctx, indexModel)
	require.NoError(t, err, "Failed to create unique index after drop")
}

func TestRepository_PushChatSession(t *testing.T) {
	beforeEachRepoTest(t)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userEmail := "newuser@example.com"
	chatTitle := "Initial Chat"
	now := primitive.NewDateTimeFromTime(time.Now())

	newChatSession := domain.ChatSession{
		ChatID:      "",
		ChatTitle:   chatTitle,
		StartTime:   now,
		LastUpdated: now,
		Messages:    []domain.Message{},
	}

	// Test 1: Create a new ChatHistory document and add a session
	returnedChatID, err := testChatRepository.PushChatSession(ctx, userEmail, newChatSession)
	require.NoError(t, err)
	assert.NotEmpty(t, returnedChatID)

	var chatHistory domain.ChatHistory
	err = testChatCollection.FindOne(ctx, bson.M{"user_email": userEmail}).Decode(&chatHistory)
	require.NoError(t, err)
	assert.Equal(t, userEmail, chatHistory.UserEmail)
	assert.Len(t, chatHistory.ChatSessions, 1)
	assert.Equal(t, returnedChatID, chatHistory.ChatSessions[0].ChatID)

	// Test 2: Add another session to the existing ChatHistory document
	chatID2 := uuid.New().String()
	newChatSession2 := domain.ChatSession{
		ChatID:      chatID2,
		ChatTitle:   "Second Chat",
		StartTime:   now,
		LastUpdated: now,
		Messages:    []domain.Message{},
	}
	returnedChatID2, err := testChatRepository.PushChatSession(ctx, userEmail, newChatSession2)
	require.NoError(t, err)
	assert.NotEmpty(t, returnedChatID2)

	var updatedChatHistory domain.ChatHistory
	err = testChatCollection.FindOne(ctx, bson.M{"user_email": userEmail}).Decode(&updatedChatHistory)
	require.NoError(t, err)
	assert.Len(t, updatedChatHistory.ChatSessions, 2)
	assert.Equal(t, returnedChatID2, updatedChatHistory.ChatSessions[1].ChatID)
}

func TestRepository_FindByUserEmail(t *testing.T) {
	beforeEachRepoTest(t)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userEmail := "finduser@example.com"
	now := primitive.NewDateTimeFromTime(time.Now())

	// Pre-populate data
	initialChat := domain.ChatSession{
		ChatID:      "", // Let repository generate the ID
		ChatTitle:   "Find Me Chat",
		StartTime:   now,
		LastUpdated: now,
		Messages:    []domain.Message{},
	}
	returnedChatID, err := testChatRepository.PushChatSession(ctx, userEmail, initialChat)
	require.NoError(t, err)

	// Test 1: Find existing user's chat history
	foundHistory, err := testChatRepository.FindByUserEmail(ctx, userEmail)
	require.NoError(t, err)
	assert.Equal(t, userEmail, foundHistory.UserEmail)
	assert.Len(t, foundHistory.ChatSessions, 1)
	assert.Equal(t, returnedChatID, foundHistory.ChatSessions[0].ChatID)

	// Test 2: User not found
	_, err = testChatRepository.FindByUserEmail(ctx, "nonexistent@example.com")
	assert.ErrorIs(t, err, errors.ErrNotFound)
}

func TestRepository_PushMessageToSession(t *testing.T) {
	beforeEachRepoTest(t)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userEmail := "messageuser@example.com"
	now := primitive.NewDateTimeFromTime(time.Now())

	// Create initial chat session
	initialChat := domain.ChatSession{
		ChatID:      "", // Let repository generate the ID
		ChatTitle:   "Chat with Messages",
		StartTime:   now,
		LastUpdated: now,
		Messages:    []domain.Message{},
	}
	returnedChatID, err := testChatRepository.PushChatSession(ctx, userEmail, initialChat)
	require.NoError(t, err)

	// Test 1: Add a message
	messageText := "What's up?"
	product := domain.Product{ID: "prod001", Title: "Gadget", Price: domain.Price{USD: 10.0}}
	newMessage := domain.Message{
		UserPrompt: messageText,
		CreatedAt:  now,
		Products:   []domain.Product{product},
	}
	err = testChatRepository.PushMessageToSession(ctx, userEmail, returnedChatID, newMessage)
	require.NoError(t, err)

	// Verify in DB
	var chatHistory domain.ChatHistory
	err = testChatCollection.FindOne(ctx, bson.M{"user_email": userEmail}).Decode(&chatHistory)
	require.NoError(t, err)
	assert.Len(t, chatHistory.ChatSessions, 1)
	assert.Len(t, chatHistory.ChatSessions[0].Messages, 1)
	assert.Equal(t, messageText, chatHistory.ChatSessions[0].Messages[0].UserPrompt)
	assert.Equal(t, "prod001", chatHistory.ChatSessions[0].Messages[0].Products[0].ID)

	// Test 2: Add another message
	messageText2 := "Show me more!"
	newMessage2 := domain.Message{
		UserPrompt: messageText2,
		CreatedAt:  primitive.NewDateTimeFromTime(time.Now().Add(time.Minute)),
		Products:   []domain.Product{},
	}
	err = testChatRepository.PushMessageToSession(ctx, userEmail, returnedChatID, newMessage2)
	require.NoError(t, err)

	err = testChatCollection.FindOne(ctx, bson.M{"user_email": userEmail}).Decode(&chatHistory)
	require.NoError(t, err)
	assert.Len(t, chatHistory.ChatSessions[0].Messages, 2)
	assert.Equal(t, messageText2, chatHistory.ChatSessions[0].Messages[1].UserPrompt)

	// Test 3: ChatID not found for existing user
	err = testChatRepository.PushMessageToSession(ctx, userEmail, "nonexistent-chat-id", newMessage)
	assert.Error(t, err) // Just check that it returns an error

	// Test 4: UserEmail not found
	err = testChatRepository.PushMessageToSession(ctx, "nonexistent@example.com", returnedChatID, newMessage)
	assert.Error(t, err) // Just check that it returns an error
}

func TestRepository_PullChatSession(t *testing.T) {
	beforeEachRepoTest(t)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userEmail := "deleteuser@example.com"
	now := primitive.NewDateTimeFromTime(time.Now())

	// Pre-populate data with two chat sessions
	chat1 := domain.ChatSession{
		ChatID:      "", // Let repository generate the ID
		ChatTitle:   "Chat to Delete", 
		StartTime:   now, 
		LastUpdated: now,
		Messages:    []domain.Message{},
	}
	chat2 := domain.ChatSession{
		ChatID:      "", // Let repository generate the ID
		ChatTitle:   "Keep This Chat", 
		StartTime:   now, 
		LastUpdated: now,
		Messages:    []domain.Message{},
	}

	returnedChatID1, err := testChatRepository.PushChatSession(ctx, userEmail, chat1)
	require.NoError(t, err)
	returnedChatID2, err := testChatRepository.PushChatSession(ctx, userEmail, chat2)
	require.NoError(t, err)

	// Verify initial state
	var chatHistory domain.ChatHistory
	err = testChatCollection.FindOne(ctx, bson.M{"user_email": userEmail}).Decode(&chatHistory)
	require.NoError(t, err)
	assert.Len(t, chatHistory.ChatSessions, 2)

	// Test 1: Delete an existing chat session for an existing user
	t.Run("DeleteExistingChatSession", func(t *testing.T) {
		err = testChatRepository.PullChatSession(ctx, userEmail, returnedChatID1)
		require.NoError(t, err) // Expect no error for successful deletion

		// Verify deletion
		err = testChatCollection.FindOne(ctx, bson.M{"user_email": userEmail}).Decode(&chatHistory)
		require.NoError(t, err)
		assert.Len(t, chatHistory.ChatSessions, 1)
		assert.Equal(t, returnedChatID2, chatHistory.ChatSessions[0].ChatID) // Ensure the other chat remains
	})

	// Test 2: Try to delete a nonexistent chat session for an existing user
	t.Run("DeleteNonExistentChatSessionForExistingUser", func(t *testing.T) {
		err = testChatRepository.PullChatSession(ctx, userEmail, "nonexistent-chat-id")
		require.Error(t, err) // Ensure an error is returned
	})

	// Test 3: Try to delete a chat session for a nonexistent user
	t.Run("DeleteChatSessionForNonExistentUser", func(t *testing.T) {
		nonExistentUserEmail := "truly.nonexistent@example.com"
		err = testChatRepository.PullChatSession(ctx, nonExistentUserEmail, returnedChatID2)

		// --- DEBUGGING OUTPUT ---
		fmt.Printf("\nDEBUG: For nonexistent user '%s', chatID '%s': Error from PullChatSession: %v (type: %T)\n",
			nonExistentUserEmail, returnedChatID2, err, err)
		// --- END DEBUGGING OUTPUT ---

		require.Error(t, err) // Ensure an error is returned
	})
}