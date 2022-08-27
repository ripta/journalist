// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/mrusme/journalist/ent/feed"
	"github.com/mrusme/journalist/ent/item"
	"github.com/mrusme/journalist/ent/read"
	"github.com/mrusme/journalist/ent/schema"
	"github.com/mrusme/journalist/ent/subscription"
	"github.com/mrusme/journalist/ent/token"
	"github.com/mrusme/journalist/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	feedFields := schema.Feed{}.Fields()
	_ = feedFields
	// feedDescURL is the schema descriptor for url field.
	feedDescURL := feedFields[1].Descriptor()
	// feed.URLValidator is a validator for the "url" field. It is called by the builders before save.
	feed.URLValidator = feedDescURL.Validators[0].(func(string) error)
	// feedDescUsername is the schema descriptor for username field.
	feedDescUsername := feedFields[2].Descriptor()
	// feed.DefaultUsername holds the default value on creation for the username field.
	feed.DefaultUsername = feedDescUsername.Default.(string)
	// feedDescPassword is the schema descriptor for password field.
	feedDescPassword := feedFields[3].Descriptor()
	// feed.DefaultPassword holds the default value on creation for the password field.
	feed.DefaultPassword = feedDescPassword.Default.(string)
	// feedDescCreatedAt is the schema descriptor for created_at field.
	feedDescCreatedAt := feedFields[18].Descriptor()
	// feed.DefaultCreatedAt holds the default value on creation for the created_at field.
	feed.DefaultCreatedAt = feedDescCreatedAt.Default.(func() time.Time)
	// feedDescUpdatedAt is the schema descriptor for updated_at field.
	feedDescUpdatedAt := feedFields[19].Descriptor()
	// feed.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	feed.DefaultUpdatedAt = feedDescUpdatedAt.Default.(func() time.Time)
	// feed.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	feed.UpdateDefaultUpdatedAt = feedDescUpdatedAt.UpdateDefault.(func() time.Time)
	// feedDescID is the schema descriptor for id field.
	feedDescID := feedFields[0].Descriptor()
	// feed.DefaultID holds the default value on creation for the id field.
	feed.DefaultID = feedDescID.Default.(func() uuid.UUID)
	itemFields := schema.Item{}.Fields()
	_ = itemFields
	// itemDescItemLink is the schema descriptor for item_link field.
	itemDescItemLink := itemFields[5].Descriptor()
	// item.ItemLinkValidator is a validator for the "item_link" field. It is called by the builders before save.
	item.ItemLinkValidator = itemDescItemLink.Validators[0].(func(string) error)
	// itemDescCreatedAt is the schema descriptor for created_at field.
	itemDescCreatedAt := itemFields[21].Descriptor()
	// item.DefaultCreatedAt holds the default value on creation for the created_at field.
	item.DefaultCreatedAt = itemDescCreatedAt.Default.(func() time.Time)
	// itemDescUpdatedAt is the schema descriptor for updated_at field.
	itemDescUpdatedAt := itemFields[22].Descriptor()
	// item.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	item.DefaultUpdatedAt = itemDescUpdatedAt.Default.(func() time.Time)
	// item.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	item.UpdateDefaultUpdatedAt = itemDescUpdatedAt.UpdateDefault.(func() time.Time)
	// itemDescID is the schema descriptor for id field.
	itemDescID := itemFields[0].Descriptor()
	// item.DefaultID holds the default value on creation for the id field.
	item.DefaultID = itemDescID.Default.(func() uuid.UUID)
	readFields := schema.Read{}.Fields()
	_ = readFields
	// readDescCreatedAt is the schema descriptor for created_at field.
	readDescCreatedAt := readFields[3].Descriptor()
	// read.DefaultCreatedAt holds the default value on creation for the created_at field.
	read.DefaultCreatedAt = readDescCreatedAt.Default.(func() time.Time)
	// readDescID is the schema descriptor for id field.
	readDescID := readFields[0].Descriptor()
	// read.DefaultID holds the default value on creation for the id field.
	read.DefaultID = readDescID.Default.(func() uuid.UUID)
	subscriptionFields := schema.Subscription{}.Fields()
	_ = subscriptionFields
	// subscriptionDescName is the schema descriptor for name field.
	subscriptionDescName := subscriptionFields[3].Descriptor()
	// subscription.NameValidator is a validator for the "name" field. It is called by the builders before save.
	subscription.NameValidator = subscriptionDescName.Validators[0].(func(string) error)
	// subscriptionDescGroup is the schema descriptor for group field.
	subscriptionDescGroup := subscriptionFields[4].Descriptor()
	// subscription.GroupValidator is a validator for the "group" field. It is called by the builders before save.
	subscription.GroupValidator = subscriptionDescGroup.Validators[0].(func(string) error)
	// subscriptionDescCreatedAt is the schema descriptor for created_at field.
	subscriptionDescCreatedAt := subscriptionFields[5].Descriptor()
	// subscription.DefaultCreatedAt holds the default value on creation for the created_at field.
	subscription.DefaultCreatedAt = subscriptionDescCreatedAt.Default.(func() time.Time)
	// subscriptionDescID is the schema descriptor for id field.
	subscriptionDescID := subscriptionFields[0].Descriptor()
	// subscription.DefaultID holds the default value on creation for the id field.
	subscription.DefaultID = subscriptionDescID.Default.(func() uuid.UUID)
	tokenFields := schema.Token{}.Fields()
	_ = tokenFields
	// tokenDescType is the schema descriptor for type field.
	tokenDescType := tokenFields[1].Descriptor()
	// token.DefaultType holds the default value on creation for the type field.
	token.DefaultType = tokenDescType.Default.(string)
	// token.TypeValidator is a validator for the "type" field. It is called by the builders before save.
	token.TypeValidator = tokenDescType.Validators[0].(func(string) error)
	// tokenDescName is the schema descriptor for name field.
	tokenDescName := tokenFields[2].Descriptor()
	// token.NameValidator is a validator for the "name" field. It is called by the builders before save.
	token.NameValidator = tokenDescName.Validators[0].(func(string) error)
	// tokenDescCreatedAt is the schema descriptor for created_at field.
	tokenDescCreatedAt := tokenFields[4].Descriptor()
	// token.DefaultCreatedAt holds the default value on creation for the created_at field.
	token.DefaultCreatedAt = tokenDescCreatedAt.Default.(func() time.Time)
	// tokenDescUpdatedAt is the schema descriptor for updated_at field.
	tokenDescUpdatedAt := tokenFields[5].Descriptor()
	// token.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	token.DefaultUpdatedAt = tokenDescUpdatedAt.Default.(func() time.Time)
	// token.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	token.UpdateDefaultUpdatedAt = tokenDescUpdatedAt.UpdateDefault.(func() time.Time)
	// tokenDescID is the schema descriptor for id field.
	tokenDescID := tokenFields[0].Descriptor()
	// token.DefaultID holds the default value on creation for the id field.
	token.DefaultID = tokenDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[1].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[2].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescRole is the schema descriptor for role field.
	userDescRole := userFields[3].Descriptor()
	// user.DefaultRole holds the default value on creation for the role field.
	user.DefaultRole = userDescRole.Default.(string)
	// user.RoleValidator is a validator for the "role" field. It is called by the builders before save.
	user.RoleValidator = userDescRole.Validators[0].(func(string) error)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[4].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[5].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}