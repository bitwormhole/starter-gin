package security

import (
	"encoding/base64"
	"errors"
	"sort"
	"strings"

	"github.com/bitwormhole/starter/collection"
	"github.com/gin-gonic/gin"
)

////////////////////////////////////////////////////////////////////////////////

type TokenSessionFactory struct {
}

func (inst *TokenSessionFactory) _impl_() SessionFactory {
	return inst
}

func (inst *TokenSessionFactory) Create(c *gin.Context) Session {
	session := &innerTokenSession{context: c, factory: inst}
	return session.init()
}

func (inst *TokenSessionFactory) setToken(session *innerTokenSession, token string) error {
	session.context.Writer.Header().Set("x-set-token", token)
	session.plainTextOld = token
	session.plainTextNew = token
	return nil
}

func (inst *TokenSessionFactory) getToken(session *innerTokenSession) (string, error) {
	token := session.context.Request.Header.Get("x-token")
	if token == "" {
		return "", errors.New("no http header named 'x-token'.")
	}
	session.plainTextOld = token
	return token, nil
}

func (inst *TokenSessionFactory) parseProperties(text string, props collection.Properties) error {
	_, err := collection.ParseProperties(text, props)
	return err
}

func (inst *TokenSessionFactory) stringifyProperties(props collection.Properties) (string, error) {
	table := props.Export(nil)
	keys := []string{}
	for key := range table {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	builder := &strings.Builder{}
	for index := range keys {
		key := keys[index]
		val := table[key]
		if val == "" {
			continue
		}
		builder.WriteString(key)
		builder.WriteString("=")
		builder.WriteString(val)
		builder.WriteString("\n")
	}
	str := builder.String()
	return str, nil
}

func (inst *TokenSessionFactory) base64encode(data []byte) (string, error) {
	str := base64.StdEncoding.EncodeToString(data)
	return str, nil
}

func (inst *TokenSessionFactory) base64decode(text string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(text)
}

func (inst *TokenSessionFactory) encrypt(input []byte) ([]byte, error) {
	return input, nil
}

func (inst *TokenSessionFactory) decrypt(input []byte) ([]byte, error) {
	return input, nil
}

////////////////////////////////////////////////////////////////////////////////

type innerTokenSession struct {
	factory *TokenSessionFactory
	context *gin.Context

	getter       SessionGetter
	setter       SessionSetter
	properties   collection.Properties
	plainTextOld string
	plainTextNew string
}

func (inst *innerTokenSession) init() Session {
	inst.properties = collection.CreateProperties()
	inst.getter = nil
	inst.setter = nil
	return inst
}

func (inst *innerTokenSession) Properties() collection.Properties {
	return inst.properties
}

func (inst *innerTokenSession) Getter() SessionGetter {
	return inst.getter
}

func (inst *innerTokenSession) Setter() SessionSetter {
	return inst.setter
}

func (inst *innerTokenSession) Open() error {

	factory := inst.factory
	props := inst.properties

	token, err := factory.getToken(inst)
	if err != nil {
		return err
	}

	data, err := factory.base64decode(token)
	if err != nil {
		return err
	}

	plain, err := factory.decrypt(data)
	if err != nil {
		return err
	}

	err = factory.parseProperties(string(plain), props)
	if err != nil {
		return err
	}

	return nil
}

func (inst *innerTokenSession) Close() error {
	return inst.Commit()
}

func (inst *innerTokenSession) Commit() error {

	factory := inst.factory
	props := inst.properties

	plain, err := factory.stringifyProperties(props)
	if err != nil {
		return err
	}

	data, err := factory.encrypt([]byte(plain))
	if err != nil {
		return err
	}

	token, err := factory.base64encode(data)
	if err != nil {
		return err
	}

	if token == inst.plainTextOld {
		// skip
		return nil
	}

	err = factory.setToken(inst, token)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////
