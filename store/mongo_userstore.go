package store

import (
	"context"
	"time"

	"github.com/blurbee/otpserver/api"
	"github.com/blurbee/otpserver/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Define a struct to map to your MongoDB document structure
type User struct {
	ID       string `bson:"_id"`
	Email    string `bson:"email"`
	Phone    string `bson:"phone"`
	Text     string `bson:"text"`
	Whatsapp string `bson:"whatsapp"`
}

type mongoUserStore struct {
	clientOptions *options.ClientOptions
	client        *mongo.Client
	config        util.MongoStoreConfig
}

var dbmap map[string]*mongoUserStore = map[string]*mongoUserStore{}

func (m *mongoUserStore) Close() (err api.StatusCode) {
	m.client.Disconnect(context.Background())
	return api.OK
}

func CloseAll() (err api.StatusCode) {
	for _, value := range dbmap {
		value.Close()
	}
	return api.OK
}

func InitMongo(cfg *util.Config) (err api.StatusCode) {
	mcfgs := cfg.GetMongoConfigs()
	if mcfgs == nil {
		return api.INVALID_INPUT
	}

	for _, mcfg := range *mcfgs {
		var m *mongoUserStore = new(mongoUserStore)
		var er error

		// get env variable
		connurl := cfg.GetSecret(mcfg.ConnectionUrlEnv)

		// Set up MongoDB connection
		m.clientOptions = options.Client().ApplyURI(connurl)
		m.client, er = mongo.Connect(context.Background(), m.clientOptions)
		if er != nil {
			util.Error("Unable to connect to Mongo:", er)
			err = api.CONN_FAILED
			return
		}

		// Check the connection
		er = m.client.Ping(context.Background(), nil)
		if er != nil {
			util.Error("Mongo connected but ping failed:", er)
			err = api.CONN_FAILED
			return
		}

		// copy mongo config to store
		m.config = mcfg

		util.Info("Connected to MongoDB:", mcfg.Id)
		err = api.OK
		dbmap[mcfg.Id] = m
	}
	return
}

/*
Get db given API
*/
func GetDB(id string) (db api.UserStore, err api.StatusCode) {
	db, er := dbmap[id]
	if er == false {
		return db, api.INVALID_INPUT
	}
	return db, api.OK
}

func (m *mongoUserStore) getAttribute(userid string, attribute string) (value string, err api.StatusCode) {

	// Specify the database and collection
	collection := m.client.Database(m.config.Database).Collection(m.config.Collection)

	// Create a context with a timeout (10 seconds in this example)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Search for the user by ID
	var user User
	er := collection.FindOne(ctx, bson.M{"_id": userid}).Decode(&user)
	if er != nil {
		util.Info("Error finding user: ", er)
		err = api.USER_NOT_FOUND
		return
	}

	// Print the user's email and phone number
	util.Debug("User ID: ", user.ID, " Email: ", user.Email, " Phone: %s ", user.Phone)
	err = api.OK
	switch attribute {
	case "Email":
		value = user.Email
	case "Phone":
		value = user.Phone
	case "Text":
		value = user.Text
	case "Whatsapp":
		value = user.Whatsapp
	}

	if value == "" {
		err = api.VALUE_NOT_FOUND
	}

	return
}

// Given user id, get the user's email address
func (m *mongoUserStore) GetEmail(userid string) (email string, err api.StatusCode) {
	return m.getAttribute(userid, "Email")
}

// Given user id, get the user's email address
func (m *mongoUserStore) GetPhone(userid string) (phone string, err api.StatusCode) {
	return m.getAttribute(userid, "Phone")
}

// Given user id, get the whatsapp number
func (m *mongoUserStore) GetWhatsapp(userid string) (whatsapp string, err api.StatusCode) {
	return m.getAttribute(userid, "Whatsapp")
}

// Given user id, get the text number
func (m *mongoUserStore) GetText(userid string) (text string, err api.StatusCode) {
	return m.getAttribute(userid, "Text")
}
