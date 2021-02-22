package mongodb

import (
	"account-manager/merchant"
	"account-manager/util"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
	"sync"
)

const (
	DB         = "pace-merchant"
	Collection = "merchant"
)

var mongoOnce sync.Once

type MerchantMongoDB struct {
	Client *mongo.Client
}

// Client returns the mongo client
func (mdb MerchantMongoDB) GetClient() *mongo.Client {
	return mdb.Client
}

func (mdb MerchantMongoDB) GetMerchants() ([]merchant.Merchant, error) {
	var merchants []merchant.Merchant
	filter := bson.D{{}}

	collection := mdb.Client.Database(DB).Collection(Collection)

	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return merchants, util.NewCustomError("merchantdb: get merchants: finding merchant: " + err.Error())
	}

	for cur.Next(context.TODO()) {
		m := merchant.Merchant{}
		err = cur.Decode(&m)

		if err != nil {
			return merchants, util.NewCustomError("merchantdb: get merchants: decoding: " + err.Error())
		}

		merchants = append(merchants, m)
	}

	cur.Close(context.TODO())

	if len(merchants) == 0 {
		return merchants, util.NewCustomError("merchantdb: get merchants: " + mongo.ErrNoDocuments.Error())
	}

	return merchants, nil
}

func (mdb MerchantMongoDB) AddMember(merchantID string, member merchant.Member) error {
	var m merchant.Merchant
	filter := bson.D{primitive.E{Key: "_id", Value: merchantID}}

	collection :=  mdb.Client.Database(DB).Collection(Collection)

	err := collection.FindOne(context.TODO(), filter).Decode(&m)
	if err != nil {
		return util.NewCustomError("merchantdb: add member: finding merchant: " + err.Error())
	}

	key := strings.Split(member.Email, "@")[0]
	if v, ok := m.Members[key]; ok {
		return util.NewCustomError("merchantdb: add member: member already exist: " + v.Email)
	} else {
		m.Members[key] = member

		updater := bson.D{primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "members", Value: m.Members},
		}}}

		_, err = collection.UpdateOne(context.TODO(), filter, updater)
		if err != nil {
			return util.NewCustomError("merchantdb: add member: updating merchant: " + err.Error())
		}
		return nil
	}
}

func (mdb MerchantMongoDB) UpdateMember(merchantID string, member merchant.Member) error {
	var m merchant.Merchant
	filter := bson.D{primitive.E{Key: "_id", Value: merchantID}}

	collection :=  mdb.Client.Database(DB).Collection(Collection)

	err := collection.FindOne(context.TODO(), filter).Decode(&m)
	if err != nil {
		return util.NewCustomError("merchantdb: update member: finding merchant: " + err.Error())
	}

	key := strings.Split(member.Email, "@")[0]
	if v, ok := m.Members[key]; ok {
		v.Email = member.Email
		v.Name = member.Name
		m.Members[key] = v

		updater := bson.D{primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "members", Value: m.Members},
		}}}

		_, err = collection.UpdateOne(context.TODO(), filter, updater)
		if err != nil {
			return util.NewCustomError("merchantdb: update member: updating merchant: " + err.Error())
		}
		return nil

	} else {
		return util.NewCustomError("merchantdb: update member: member doesn't exist: " + member.Email)
	}
}

func (mdb MerchantMongoDB) DeleteMember(merchantID string, memberEmail string) error {
	var m merchant.Merchant
	filter := bson.D{primitive.E{Key: "_id", Value: merchantID}}

	collection :=  mdb.Client.Database(DB).Collection(Collection)

	err := collection.FindOne(context.TODO(), filter).Decode(&m)
	if err != nil {
		return util.NewCustomError("merchantdb: delete member: finding merchant: " + err.Error())
	}

	key := strings.Split(memberEmail, "@")[0]
	if _, ok := m.Members[key]; ok {
		delete(m.Members, key)

		updater := bson.D{primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "members", Value: m.Members},
		}}}

		_, err = collection.UpdateOne(context.TODO(), filter, updater)
		if err != nil {
			return util.NewCustomError("merchantdb: delete member: updating merchant: " + err.Error())
		}
		return nil

	} else {
		return util.NewCustomError("merchantdb: delete member: merchant doesn't exist: " + memberEmail)
	}
}

func (mdb MerchantMongoDB) GetMembers(merchantID string) ([]merchant.Member, error) {
	var m merchant.Merchant
	filter := bson.D{primitive.E{Key: "_id", Value: merchantID}}

	collection :=  mdb.Client.Database(DB).Collection(Collection)

	err := collection.FindOne(context.TODO(), filter).Decode(&m)
	if err != nil {
		return nil, util.NewCustomError("merchantdb: get members: finding merchant: " + err.Error())
	}
	var members []merchant.Member

	for _, v := range m.Members {
		members = append(members, v)
	}

	return members, nil
}
