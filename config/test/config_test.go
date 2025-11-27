package test

import (
	"os"
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/simonalong/gole/config"
	"github.com/simonalong/gole/logger"
	"github.com/simonalong/gole/util"
)

// key1:
//
//	key2:
//	  intdata: 12
//	  strdata: "data"
//	  booldata: true
//	  int64data: 12
//	  floatdata: 12.3
//	  objdata:
//	    field1: 1
//	    field2: "value2"
//	  arraydata:
//	    - 1
//	    - 2
//	    - 3
//	  arrayobjdata:
//	    - field1: 1
//	      field2: "name1"
//	    - field1: 2
//	      field2: "name2"
func TestLoadConfig(t *testing.T) {
	config.Load()

	assert.Equal(t, config.GetValueInt("key1.key2.intdata"), 12)
	assert.Equal(t, config.GetValueString("key1.key2.strdata"), "data")
	assert.Equal(t, config.GetValueBool("key1.key2.booldata"), true)
	assert.Equal(t, config.GetValueInt64("key1.key2.int64data"), util.ToInt64(12))
	assert.Equal(t, config.GetValueFloat32("key1.key2.floatdata"), util.ToFloat32(12.3))

	expectData := map[string]any{
		"field1": 1,
		"field2": "value2",
	}
	actData := map[string]any{}
	err := config.GetValueObject("key1.key2.objdata", &actData)
	if err != nil {
		return
	}
	assert.Equal(t, util.ToInt(actData["field1"]), util.ToInt(expectData["field1"]))
	assert.Equal(t, actData["field2"], expectData["field2"])
	//assert.Equal(t, config.GetValueArrayInt("key1.key2.arraydata"), []int{1, 2, 3})
}

func TestLoadConfig1(t *testing.T) {
	config.Load()
	assert.Equal(t, config.GetValueString("place.username"), "user")
}

func TestLoadConfigFrom(t *testing.T) {
	config.LoadYamlFile("./application.yaml")
	config.FinishLoad()
	assert.Equal(t, config.GetValueInt("key1.key2.intdata"), 12)
}

type Jing struct {
	Entity JingEntity
}

type JingEntity struct {
	Name string
}

func TestLoadConfig2(t *testing.T) {
	config.LoadFile("./application-jing.yaml")
	config.FinishLoad()

	jing := Jing{}
	config.GetValueObject("jing", &jing)
	assert.Equal(t, jing.Entity.Name, "#.key")
}

// 测试小驼峰
func TestSmall(t *testing.T) {
	config.LoadFile("./application.yaml")
	config.FinishLoad()

	entity1 := SmallEntity{}
	_ = config.GetValueObject("key1.ok1", &entity1)
	assert.Equal(t, entity1.NameAge, 32)
	assert.Equal(t, entity1.HaoDeOk, 12)

	entity2 := SmallEntity{}
	_ = config.GetValueObject("key1.ok2", &entity2)
	assert.Equal(t, entity2.NameAge, 32)
	assert.Equal(t, entity2.HaoDeOk, 12)

	entity3 := SmallEntity{}
	_ = config.GetValueObject("key1.ok3", &entity3)
	assert.Equal(t, entity3.NameAge, 32)
	assert.Equal(t, entity3.HaoDeOk, 12)

	entity4 := SmallEntity{}
	_ = config.GetValueObject("key1.ok4", &entity4)
	assert.Equal(t, entity4.NameAge, 32)
	assert.Equal(t, entity4.HaoDeOk, 12)
}

type SmallEntity struct {
	HaoDeOk int
	NameAge int
}

// 测试兼容标签：json、yaml
func TestJsonOrYamlTag(t *testing.T) {
	config.LoadFile("./application.yaml")
	config.FinishLoad()

	entity1 := SmallEntityJsonTag{}
	_ = config.GetValueObject("key1.json", &entity1)
	assert.Equal(t, entity1.NameAge, 32)
	assert.Equal(t, entity1.HaoDeOk, 12)

	entity1_1 := SmallEntityJsonTag2{}
	_ = config.GetValueObject("key1.json", &entity1_1)
	assert.Equal(t, entity1_1.NameAge, 32)
	assert.Equal(t, entity1_1.HaoDeOk, 12)

	entity2 := SmallEntityYamlTag{}
	_ = config.GetValueObject("key1.yaml", &entity2)
	assert.Equal(t, entity2.NameAge, 32)
	assert.Equal(t, entity2.HaoDeOk, 12)

	entity2_1 := SmallEntityYamlTag2{}
	_ = config.GetValueObject("key1.yaml", &entity2_1)
	assert.Equal(t, entity2_1.NameAge, 32)
	assert.Equal(t, entity2_1.HaoDeOk, 12)
}

type SmallEntityJsonTag struct {
	HaoDeOk int `json:"test_haode"`
	NameAge int `json:"test_namehaha"`
}

type SmallEntityJsonTag2 struct {
	HaoDeOk int `json:"test_haode,omitempty"`
	NameAge int `json:"test_namehaha"`
}

type SmallEntityYamlTag struct {
	HaoDeOk int `yaml:"test_haode"`
	NameAge int `yaml:"test_namehaha"`
}

type SmallEntityYamlTag2 struct {
	HaoDeOk int `yaml:"test_haode,omitempty"`
	NameAge int `yaml:"test_namehaha,flow"`
}

// 测试读取某个文件
func TestRead(t *testing.T) {
	config.LoadFile("./application-local.yaml")
	config.FinishLoad()

	en := EntityTest{}
	err := config.GetValueObject("entity", &en)
	if err != nil {
		logger.Warn("转换告警")
		return
	}

	assert.Equal(t, en.Name, "name-change")
}

type EntityTest struct {
	Name string
}

// 测试append
func TestAppend(t *testing.T) {
	config.LoadFile("./application-append-original.yaml")
	config.AppendFile("./application-append.yaml")
	config.FinishLoad()

	assert.Equal(t, config.GetValueString("a.b.c"), "c-value-change")
	assert.Equal(t, config.GetValueString("a.b.d"), "d-value")
	assert.Equal(t, config.GetValueString("a.b.e.f"), "f-value")
}

// 测试cm文件位置定制化
func TestConfigInit(t *testing.T) {
	_ = os.Setenv("gole.config.additional-location", "./application-append.yaml")
	config.LoadConfigFromRelativePath("./application-append-original.yaml")
	config.FinishLoad()

	assert.Equal(t, config.GetValueString("a.b.c"), "c-value-change")
	assert.Equal(t, config.GetValueString("a.b.d"), "d-value")
	assert.Equal(t, config.GetValueString("a.b.e.f"), "f-value")
}

// 测试：yaml占位符的功能
func TestPlaceHolder1(t *testing.T) {
	config.LoadFile("./application-place1.yml")
	config.FinishLoad()

	assert.Equal(t, config.GetValueString("place.name"), "test")
	assert.Equal(t, config.GetValueString("test.name"), "test")
	assert.Equal(t, config.GetValueString("test.name2"), "test2")
}

// 测试：yaml占位符的功能
func TestPlaceHolder2(t *testing.T) {
	config.LoadFile("./application-place1.yaml")
	config.FinishLoad()

	assert.Equal(t, config.GetValueString("place.name"), "test")
	assert.Equal(t, config.GetValueString("test.name"), "test")
	assert.Equal(t, config.GetValueString("test.name2"), "test2")
	assert.Equal(t, config.GetValueString("test.name3"), "test2:oktest")
	assert.Equal(t, config.GetValueString("test.name4"), "pre_test2:oktest")
	assert.Equal(t, config.GetValueString("test.name5"), "pre_test2:test:oktest")
}

// 测试：yaml占位符的功能
func TestPlaceHolder3(t *testing.T) {
	config.LoadFile("./application-place1.json")
	config.FinishLoad()

	assert.Equal(t, config.GetValueString("place.name"), "test")
	assert.Equal(t, config.GetValueString("test.name"), "test")
	assert.Equal(t, config.GetValueString("test.name2"), "test2")
}

// 测试：yaml占位符的功能
func TestPlaceHolder4(t *testing.T) {
	config.LoadFile("./application-place1.properties")
	config.FinishLoad()

	assert.Equal(t, config.GetValueString("place.name"), "test")
	assert.Equal(t, config.GetValueString("test.name"), "test")
	assert.Equal(t, config.GetValueString("test.name2"), "test2")
}

// 测试：yaml占位符的功能
func TestPlaceHolder5(t *testing.T) {
	config.LoadFile("./application-place3.yaml")
	config.FinishLoad()
	type NameEntity struct {
		Name    string
		Age     int
		Address string
	}
	entity := NameEntity{}
	config.GetValueObject("data", &entity)

	assert.Equal(t, config.GetValueString("data.name"), "zhou")
	assert.Equal(t, entity.Name, "zhou")
	assert.Equal(t, entity.Age, 18)
	assert.Equal(t, entity.Address, "杭州市西湖区这边")
}

// 测试：yaml占位符的功能
func TestGetObjectArray(t *testing.T) {
	config.LoadFile("./application-array.yaml")
	config.FinishLoad()

	type NameEntity struct {
		Name string
	}

	type DemoEntity struct {
		Key2 []NameEntity
	}

	var demoEntity DemoEntity
	config.GetValueObject("key1", &demoEntity)

	assert.Equal(t, "name1", demoEntity.Key2[0].Name)
	assert.Equal(t, "name2", demoEntity.Key2[1].Name)
}

func TestGetKeys(t *testing.T) {
	config.LoadFile("./application.yaml")
	config.FinishLoad()

	data := config.GetValue("key1")
	dataMap := util.ToMap(data)
	var keys []string
	for k, _ := range dataMap {
		keys = append(keys, util.ToString(k))
	}
	//fmt.Println(keys)
}

func TestEnv(t *testing.T) {
	//config.Load()
	//config.LoadEnvFile("./.env")
	config.LoadFile("./.env")
	config.FinishLoad()

	assert.Equal(t, config.GetValueString("test"), "33")
	assert.Equal(t, config.GetValueString("cloud"), "http://cloud.com")
	assert.Equal(t, config.GetValueString("host"), "127.0.0.1")
	assert.Equal(t, config.GetValueString("username"), "user")
	assert.Equal(t, os.Getenv("username"), "user")
}

func TestPlaceAndEnv(t *testing.T) {
	config.Clean()
	config.LoadFile("./application-place2.yaml")
	config.AppendFile("./.env")

	assert.Equal(t, config.GetValueString("cloud"), "http://cloud.com")
	assert.Equal(t, config.GetValueString("test.name"), "pre_http://cloud.com:8080/test")
}

func TestPlaceAndEnvDefault(t *testing.T) {
	config.LoadFile("./application-env-default.yaml")
	config.AppendFile("./.env")
	config.FinishLoad()

	assert.Equal(t, config.GetValueString("cloud"), "http://cloud.com")
	assert.Equal(t, config.GetValueString("demo.name"), "pre_http://cloud.com:8080/test")
	assert.Equal(t, config.GetValueString("demo.name1"), "pre_test-url:8080/test")
}

func TestGetRsaKey(t *testing.T) {
	config.LoadFile("./application-rsa.yaml")
	config.FinishLoad()

	privateKey := "-----BEGIN RSA PRIVATE KEY-----\nMIICWwIBAAKBgQCy4A/aQ1v6+IG8RxaiYYJ9ueICvhntvGM5HKaxo6gkq5Uyk4MJ\ne9OzsZd+fy5icAJQRhUNug8K8BLt7VvjsHiPx7cyIGZ4Ms7RTn6LdKR/rLOVc9rh\nE5Yc1VoPv7UGRcosAhrsjtaKp1GO1hhneks6pZWHWF09yNWEo7XxdBjagQIDAQAB\nAoGAJO8W1t5ps5x0TUfwaH7xzrv+6soN2IS5iCVeVfeQ1GGJYPQMbnze7Y+R1FC2\nZyTxlVmjJz5vtLZ1ciM8gfsCKXVjZQKiOqpjsRRok5LaKyBpgptpx8N91rggd1Wt\nhubLzvsevCTFRpKeQK4Id1DxRfY1MrRbQrHv4pSYGuh+W3ECQQDH0QGgNn2H0E4l\nqKXob+Hmos0lIVY8aHOOW63K7gIuXtcFqISnBuJFKGrvg4uacYQHmU+PV9SjJT0L\n+PcP2gBvAkEA5Suv5RK+uUBu4JzXB4TKrKCnhhwVX28m20BhFUzeoKNLxjuiCAVS\na+gmnphj1MqGHvJwW90/IFe7qk2FnhxsDwJAN/y1IuoBtFtGekDN89ndhx0YtA2q\nNxThRAMmKBUWYV3Li9dTC+Xe4pfXlrLaG/UwlFx9sWFfwDK/7yncOAHSWwJAIffR\nwJCAuJC2XpCgxrqGGARQEG7FNDoTdlgai7+zF/hcWOup3qp7RwdIAiXwVjAWpSum\nP9eRbcfTRzDqZz8rPQJAcVz4Yz1mOUQ519rsY0hmX0vijT0bnczP8xp2G90bJvr0\nHPipwLd4qJB7WhB4u7S7W1YfxCk/13ZNR8eccw2Stg==\n-----END RSA PRIVATE KEY-----\n"
	assert.Equal(t, config.GetValueString("rsa.private-key"), privateKey)
}
