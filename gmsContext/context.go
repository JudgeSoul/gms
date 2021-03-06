package gmsContext

import (
	"context"
	"log"

	"github.com/gmsorg/gms/serialize"
	"github.com/gmsorg/gms/protocol"
)

type Context struct {
	context.Context // todo context 功能待完善（参考gin的context实现）
	message         protocol.Imessage
	resultData      []byte
}

// func (c *Context) SetParam(b []byte) error {
// 	panic("implement me")
// }

func NewContext() *Context {
	return &Context{}
}

func (c *Context) SetMessage(message protocol.Imessage) error {
	c.message = message
	return nil
}

/**
把请求中的信息反序列化成用户指定的对象
*/
func (c *Context) Param(param interface{}) error {

	// 获取指定的序列化器
	serialize := serialize.GetSerialize(c.message.GetSerializeType())
	err := serialize.UnSerialize(c.message.GetData(), param)
	if err != nil {
		log.Println("[Param] error", err)
		return err
	}
	return nil
}

func (c *Context) Result(result interface{}) error {
	serialize := serialize.GetSerialize(c.message.GetSerializeType())
	r, err := serialize.Serialize(result)
	if err != nil {
		log.Println(err)
	}
	// log.Println(len(r))
	c.resultData = r
	return nil
}

func (c *Context) GetResult() ([]byte, error) {
	return c.resultData, nil
}
