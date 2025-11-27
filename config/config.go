package config

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/simonalong/gole/file"
	"github.com/simonalong/gole/listener"
	"github.com/simonalong/gole/util"
	"gopkg.in/yaml.v2"
)

var appProperty *ApplicationProperty

var configExist = false
var loadLock sync.Mutex
var Loaded = false
var CurrentProfile = ""

func Load() {
	loadLock.Lock()
	defer loadLock.Unlock()
	if Loaded {
		return
	}
	LoadConfigFromRelativePath("")
	AppendEnvFromRelativePath("")

	FinishLoad()
}

func Clean() {
	appProperty = nil
}

// FinishLoad 配置加载完成
// 说明：对于一些不使用默认Load()函数的，需要自行加载的，请在加载完之后，调用该函数，用于后续的处理
func FinishLoad() {
	// 发布事件：配置加载完成
	go func() {
		time.Sleep(1 * time.Second)
		listener.PublishEvent(EventOfLoadFinish{})
	}()
}

// LoadConfigFromRelativePath 读取相对路径
func LoadConfigFromRelativePath(resourceAbsPath string) {
	dir, _ := os.Getwd()
	LoadConfigFromAbsPath(path.Join(dir+string(os.PathSeparator), "", resourceAbsPath))
}

func AppendEnvFromRelativePath(resourceAbsPath string) {
	dir, _ := os.Getwd()
	AppendEnvFromAbsPath(path.Join(dir+string(os.PathSeparator), "", resourceAbsPath))
}

func AppendEnvFromAbsPath(resourceAbsPath string) {
	if !strings.HasSuffix(resourceAbsPath, string(os.PathSeparator)) {
		resourceAbsPath += string(os.PathSeparator)
	}
	AppendEnvFile(resourceAbsPath + ".env")
}

// LoadConfigFromAbsPath 读取绝对路径
func LoadConfigFromAbsPath(resourceAbsPath string) {
	doLoadConfigFromAbsPath(resourceAbsPath)

	cmPath := os.Getenv("gole.config.additional-location")
	if cmPath == "" {
		cmPath = "./config/application-default.yml"
	}
	AppendConfigFromRelativePath(cmPath)

	if err := GetValueObject("gole", &GoleCfg); err != nil {
		log.Printf("加载 Base 配置失败(%v)", err)
	}
}

func AppendConfigFromRelativePath(fileName string) {
	dir, _ := os.Getwd()
	fileName = path.Join(dir+string(os.PathSeparator), "", fileName)
	extend := getFileExtension(fileName)
	extend = strings.ToLower(extend)
	switch extend {
	case "yaml":
		AppendYamlFile(fileName)
	case "yml":
		AppendYamlFile(fileName)
	case "properties":
		AppendPropertyFile(fileName)
	case "json":
		AppendJsonFile(fileName)
	}
}

func AppendConfigFromAbsPath(fileName string) {
	extend := getFileExtension(fileName)
	extend = strings.ToLower(extend)
	switch extend {
	case "yaml":
		AppendYamlFile(fileName)
	case "yml":
		AppendYamlFile(fileName)
	case "properties":
		AppendPropertyFile(fileName)
	case "json":
		AppendJsonFile(fileName)
	}
}

type EnvProperty struct {
	Key   string
	Value string
}

// ExistConfigFile
// deprecate 弃用，后续请使用Loaded
func ExistConfigFile() bool {
	return configExist
}

func GetConfigValues() interface{} {
	if nil != appProperty {
		return appProperty.ValueMap
	} else {
		return nil
	}
}

func GetConfigDeepValues() interface{} {
	if nil != appProperty {
		return appProperty.ValueDeepMap
	} else {
		return nil
	}
}

func GetConfigValue(key string) interface{} {
	if nil != appProperty {
		return GetValue(key)
	} else {
		return nil
	}
}

func UpdateConfig(key string, value interface{}) {
	SetValue(key, value)
}

// 多种格式优先级：json > properties > yaml > yml
func doLoadConfigFromAbsPath(resourceAbsPath string) {
	if !strings.HasSuffix(resourceAbsPath, string(os.PathSeparator)) {
		resourceAbsPath += string(os.PathSeparator)
	}
	files, err := os.ReadDir(resourceAbsPath)
	if err != nil {
		return
	}

	if appProperty == nil {
		appProperty = &ApplicationProperty{}
		appProperty.ValueMap = make(map[string]interface{})
		appProperty.ValueDeepMap = make(map[string]interface{})
	} else if appProperty.ValueMap == nil {
		appProperty.ValueMap = make(map[string]interface{})
	} else if appProperty.ValueDeepMap == nil {
		appProperty.ValueDeepMap = make(map[string]interface{})
	}

	LoadYamlFile(resourceAbsPath + "application.yaml")
	LoadYamlFile(resourceAbsPath + "application.yml")
	LoadPropertyFile(resourceAbsPath + "application.properties")
	LoadJsonFile(resourceAbsPath + "application.json")

	for _, fileInfo := range files {
		if fileInfo.IsDir() {
			continue
		}

		fileName := fileInfo.Name()
		if !strings.HasPrefix(fileName, "application") {
			continue
		}

		// 默认配置
		if fileName == "application.yaml" {
			break
		} else if fileName == "application.yml" {
			break
		} else if fileName == "application.properties" {
			break
		} else if fileName == "application.json" {
			break
		}

		profile := getActiveProfile()
		if profile != "" {
			CurrentProfile = profile
			SetValue("gole.profiles.active", profile)
			currentProfile := getProfileFromFileName(fileName)
			if currentProfile == profile {
				AppendFile(resourceAbsPath + fileName)
			}
		}
	}
}

func LoadFile(filePath string) {
	extend := getFileExtension(filePath)
	extend = strings.ToLower(extend)
	if extend == "yaml" {
		LoadYamlFile(filePath)
	} else if extend == "yml" {
		LoadYamlFile(filePath)
	} else if extend == "properties" {
		LoadPropertyFile(filePath)
	} else if extend == "json" {
		LoadJsonFile(filePath)
	} else if extend == "env" {
		LoadEnvFile(filePath)
	}
}

func AppendFile(filePath string) {
	extend := getFileExtension(filePath)
	extend = strings.ToLower(extend)
	if extend == "yaml" {
		AppendYamlFile(filePath)
	} else if extend == "yml" {
		AppendYamlFile(filePath)
	} else if extend == "properties" {
		AppendPropertyFile(filePath)
	} else if extend == "json" {
		AppendJsonFile(filePath)
	} else if extend == "env" {
		AppendEnvFile(filePath)
	}
}

// 临时写死
// 优先级：环境变量 > 本地配置
func getActiveProfile() string {
	profile := os.Getenv("gole.profiles.active")
	if profile != "" {
		return profile
	}

	profile = GetValueString("gole.profiles.active")
	if profile != "" {
		return profile
	}
	return ""
}

func getProfileFromFileName(fileName string) string {
	if strings.HasPrefix(fileName, "application-") {
		words := strings.SplitN(fileName, ".", 2)
		appNames := words[0]

		appNameAndProfile := strings.SplitN(appNames, "-", 2)
		return appNameAndProfile[1]
	}
	return ""
}

func getFileExtension(fileName string) string {
	if strings.Contains(fileName, ".") {
		lastIndex := strings.LastIndex(fileName, ".")
		if lastIndex > -1 {
			return fileName[lastIndex+1:]
		}
	}
	return ""
}

func LoadYamlFile(filePath string) {
	if !file.FileExists(filePath) {
		return
	}
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("读取文件失败：%v", err)
		return
	}

	if appProperty == nil {
		appProperty = &ApplicationProperty{}
		appProperty.ValueMap = make(map[string]interface{})
		appProperty.ValueDeepMap = make(map[string]interface{})
	} else if appProperty.ValueMap == nil {
		appProperty.ValueMap = make(map[string]interface{})
	} else if appProperty.ValueDeepMap == nil {
		appProperty.ValueDeepMap = make(map[string]interface{})
	}

	property, err := util.YamlToProperties(string(content))
	if err != nil {
		log.Printf("YamlToProperties转换失败：%v", err)
		return
	}
	valueMap, _ := util.PropertiesToMap(property)
	appProperty.ValueMap = valueMap

	yamlMap, err := util.YamlToMap(string(content))
	if err != nil {
		log.Printf("YamlToMap转换失败：%v", err)
		return
	}
	appProperty.ValueDeepMap = yamlMap
	Loaded = true
}

func AppendYamlFile(filePath string) {
	if !file.FileExists(filePath) {
		return
	}
	content, err := os.ReadFile(filePath)
	if err != nil {
		// log.Printf("读取文件失败(%v)", err)
		return
	}

	AppendYamlContent(string(content))
	Loaded = true
}

func AppendYamlContent(content string) {
	property, err := util.YamlToProperties(content)
	if err != nil {
		//logger.Errorf("配置Append异常 YamlToProperties：%v, content=%v", err, content)
		return
	}

	if appProperty == nil {
		appProperty = &ApplicationProperty{}
		appProperty.ValueMap = make(map[string]interface{})
		appProperty.ValueDeepMap = make(map[string]interface{})
	} else if appProperty.ValueMap == nil {
		appProperty.ValueMap = make(map[string]interface{})
	} else if appProperty.ValueDeepMap == nil {
		appProperty.ValueDeepMap = make(map[string]interface{})
	}
	AppendValue(property)
}

func LoadPropertyFile(filePath string) {
	if !file.FileExists(filePath) {
		return
	}
	content, err := os.ReadFile(filePath)
	if err != nil {
		// log.Printf("读取文件失败(%v)", err)
		return
	}

	if appProperty == nil {
		appProperty = &ApplicationProperty{}
		appProperty.ValueMap = make(map[string]interface{})
		appProperty.ValueDeepMap = make(map[string]interface{})
	} else if appProperty.ValueMap == nil {
		appProperty.ValueMap = make(map[string]interface{})
	} else if appProperty.ValueDeepMap == nil {
		appProperty.ValueDeepMap = make(map[string]interface{})
	}

	valueMap, _ := util.PropertiesToMap(string(content))
	appProperty.ValueMap = valueMap

	yamlStr, _ := util.PropertiesToYaml(string(content))
	yamlMap, _ := util.YamlToMap(yamlStr)
	appProperty.ValueDeepMap = yamlMap
	Loaded = true
}

func LoadEnvFile(filePath string) {
	if !file.FileExists(filePath) {
		return
	}
	content, err := os.ReadFile(filePath)
	if err != nil {
		return
	}

	if appProperty == nil {
		appProperty = &ApplicationProperty{}
		appProperty.ValueMap = make(map[string]interface{})
		appProperty.ValueDeepMap = make(map[string]interface{})
	} else if appProperty.ValueMap == nil {
		appProperty.ValueMap = make(map[string]interface{})
	} else if appProperty.ValueDeepMap == nil {
		appProperty.ValueDeepMap = make(map[string]interface{})
	}

	valueMap, _ := util.EnvToMap(string(content))
	appProperty.ValueMap = valueMap

	// 设置到os.Env里面去
	for k, v := range valueMap {
		_ = os.Setenv(k, util.ToString(v))
	}

	yamlStr, _ := util.PropertiesToYaml(string(content))
	yamlMap, _ := util.YamlToMap(yamlStr)
	appProperty.ValueDeepMap = yamlMap
	Loaded = true
}

func AppendPropertyFile(filePath string) {
	if !file.FileExists(filePath) {
		return
	}
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("读取文件失败(%v)", err)
		return
	}

	AppendPropertyContent(string(content))
}

func AppendPropertyContent(content string) {
	valueMap, err := util.PropertiesToMap(content)
	if err != nil {
		//logger.Errorf("配置Append异常 PropertiesToMap：%v, content=%v", err, content)
		return
	}
	propertiesValue, err := util.MapToProperties(valueMap)
	if err != nil {
		//logger.Errorf("配置Append异常 MapToProperties：%v, content=%v", err, content)
		return
	}
	if appProperty == nil {
		appProperty = &ApplicationProperty{}
		appProperty.ValueMap = make(map[string]interface{})
		appProperty.ValueDeepMap = make(map[string]interface{})
	} else if appProperty.ValueMap == nil {
		appProperty.ValueMap = make(map[string]interface{})
	} else if appProperty.ValueDeepMap == nil {
		appProperty.ValueDeepMap = make(map[string]interface{})
	}
	AppendValue(propertiesValue)
}

func AppendEnvFile(filePath string) {
	if !file.FileExists(filePath) {
		return
	}
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("读取文件失败(%v)", err)
		return
	}

	if appProperty == nil {
		appProperty = &ApplicationProperty{}
		appProperty.ValueMap = make(map[string]interface{})
		appProperty.ValueDeepMap = make(map[string]interface{})
	} else if appProperty.ValueMap == nil {
		appProperty.ValueMap = make(map[string]interface{})
	} else if appProperty.ValueDeepMap == nil {
		appProperty.ValueDeepMap = make(map[string]interface{})
	}

	valueMap, err := util.EnvToMap(string(content))
	if err != nil {
		return
	}

	// 设置到os.Env里面去
	for k, v := range valueMap {
		_ = os.Setenv(k, util.ToString(v))
	}

	propertiesValue, err := util.MapToProperties(valueMap)
	if err != nil {
		return
	}

	AppendValue(propertiesValue)
}

func LoadJsonFile(filePath string) {
	if !file.FileExists(filePath) {
		return
	}
	content, err := os.ReadFile(filePath)
	if err != nil {
		// log.Printf("读取文件失败(%v)", err)
		return
	}

	if appProperty == nil {
		appProperty = &ApplicationProperty{}
		appProperty.ValueMap = make(map[string]interface{})
		appProperty.ValueDeepMap = make(map[string]interface{})
	} else if appProperty.ValueMap == nil {
		appProperty.ValueMap = make(map[string]interface{})
	} else if appProperty.ValueDeepMap == nil {
		appProperty.ValueDeepMap = make(map[string]interface{})
	}

	yamlStr, _ := util.JsonToYaml(string(content))
	property, _ := util.YamlToProperties(yamlStr)
	valueMap, _ := util.PropertiesToMap(property)
	appProperty.ValueMap = valueMap

	yamlMap, _ := util.YamlToMap(yamlStr)
	appProperty.ValueDeepMap = yamlMap
	Loaded = true
}

func AppendJsonFile(filePath string) {
	if !file.FileExists(filePath) {
		return
	}
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("读取文件失败(%v)", err)
		return
	}

	AppendJsonContent(string(content))
}

func AppendJsonContent(content string) {
	yamlStr, err := util.JsonToYaml(content)
	if err != nil {
		//logger.Errorf("配置Append异常 JsonToYaml：%v, content=%v", err, content)
		return
	}
	property, err := util.YamlToProperties(yamlStr)
	if err != nil {
		//logger.Errorf("配置Append异常 YamlToProperties：%v, content=%v", err, content)
		return
	}
	if appProperty == nil {
		appProperty = &ApplicationProperty{}
		appProperty.ValueMap = make(map[string]interface{})
		appProperty.ValueDeepMap = make(map[string]interface{})
	} else if appProperty.ValueMap == nil {
		appProperty.ValueMap = make(map[string]interface{})
	} else if appProperty.ValueDeepMap == nil {
		appProperty.ValueDeepMap = make(map[string]interface{})
	}
	AppendValue(property)
}

func AppendValue(content string) {
	pMap, err := util.PropertiesToMap(content)
	for k, v := range pMap {
		appProperty.ValueMap[k] = v
	}

	propertiesValueOfOriginal, err := util.MapToProperties(appProperty.ValueMap)
	if err != nil {
		//logger.Errorf("配置Append异常 MapToProperties：%v, content=%v", err, content)
		return
	}

	resultYaml, err := util.PropertiesToYaml(propertiesValueOfOriginal)
	if err != nil {
		//logger.Errorf("配置Append异常 PropertiesToYaml：%v, content=%v", err, content)
		return
	}
	resultDeepMap, err := util.YamlToMap(resultYaml)
	if err != nil {
		//logger.Errorf("配置Append异常 YamlToMap ：%v, content=%v", err, content)
		return
	}
	appProperty.ValueDeepMap = resultDeepMap
	Loaded = true
}

func SetValue(key string, value any) {
	if nil == value {
		return
	}
	if appProperty == nil {
		appProperty = &ApplicationProperty{}
		appProperty.ValueMap = make(map[string]interface{})
		appProperty.ValueDeepMap = make(map[string]interface{})
	} else if appProperty.ValueMap == nil {
		appProperty.ValueMap = make(map[string]interface{})
	} else if appProperty.ValueDeepMap == nil {
		appProperty.ValueDeepMap = make(map[string]interface{})
	}

	if oldValue, exist := appProperty.ValueMap[key]; exist {
		if !util.IsBaseType(reflect.TypeOf(oldValue)) {
			if reflect.TypeOf(oldValue) != reflect.TypeOf(value) {
				return
			}
		}
	}
	propertiesValueOfOriginal, err := util.MapToProperties(appProperty.ValueDeepMap)
	if err != nil {
		return
	}
	resultMap, err := util.PropertiesToMap(propertiesValueOfOriginal)
	if err != nil {
		return
	}

	resultMap, err = parseProperties(key, value, resultMap)

	appProperty.ValueMap = resultMap

	mapProperties, err := util.MapToProperties(resultMap)
	if err != nil {
		return
	}
	mapYaml, err := util.PropertiesToYaml(mapProperties)
	if err != nil {
		return
	}
	resultDeepMap, err := util.YamlToMap(mapYaml)
	if err != nil {
		return
	}
	appProperty.ValueDeepMap = resultDeepMap

	// 发布配置变更事件
	listener.PublishEvent(EventOfChange{Key: key, Value: util.ToString(util.ObjectToData(value))})
}

func parseProperties(key string, value any, resultMap map[string]any) (map[string]any, error) {
	if reflect.ValueOf(value).Kind() == reflect.Map || reflect.ValueOf(value).Kind() == reflect.Struct {
		valueMap, err := util.JsonToMap(util.ObjectToJson(value))
		if err != nil {
			return resultMap, err
		}
		for k, v := range valueMap {
			resultMap, err = parseProperties(key+"."+k, v, resultMap)
		}
	} else if reflect.ValueOf(value).Kind() == reflect.Slice || reflect.ValueOf(value).Kind() == reflect.Array {
		values := []any{}
		_, err := util.DataToEntity(util.ObjectToJson(value), &values)
		if err != nil {
			return resultMap, err
		}
		for i, v := range values {
			resultMap[key+"["+util.ToString(i)+"]"] = v
			resultMap, err = parseProperties(key+"["+util.ToString(i)+"]", v, resultMap)
		}
	} else {
		if reflect.ValueOf(value).Kind() == reflect.String && util.ToString(value) != "" {
			resultMap[key] = value
		} else if value != nil {
			resultMap[key] = value
		}
	}
	return resultMap, nil
}

func GetValueString(key string) string {
	if nil == appProperty {
		return ""
	}
	if value, exist := appProperty.ValueMap[key]; exist {
		result := util.ToString(parseVariableOfString(util.ToString(value)))
		if strings.HasPrefix(result, "'") {
			result = result[1 : len(result)-1]
		}

		if strings.HasSuffix(result, "'") {
			result = result[0 : len(result)-1]
		}
		return result
	}
	return ""
}

func GetValueInt(key string) int {
	if nil == appProperty {
		return 0
	}
	if value, exist := appProperty.ValueMap[key]; exist {
		return util.ToInt(parseVariableOfString(util.ToString(value)))
	}
	return 0
}

func GetValueInt8(key string) int8 {
	if nil == appProperty {
		return 0
	}
	if value, exist := appProperty.ValueMap[key]; exist {
		return util.ToInt8(parseVariableOfString(util.ToString(value)))
	}
	return 0
}

func GetValueInt16(key string) int16 {
	if nil == appProperty {
		return 0
	}
	if value, exist := appProperty.ValueMap[key]; exist {
		return util.ToInt16(parseVariableOfString(util.ToString(value)))
	}
	return 0
}

func GetValueInt32(key string) int32 {
	if nil == appProperty {
		return 0
	}
	if value, exist := appProperty.ValueMap[key]; exist {
		return util.ToInt32(parseVariableOfString(util.ToString(value)))
	}
	return 0
}

func GetValueInt64(key string) int64 {
	if nil == appProperty {
		return 0
	}
	if value, exist := appProperty.ValueMap[key]; exist {
		return util.ToInt64(parseVariableOfString(util.ToString(value)))
	}
	return 0
}

func GetValueUInt(key string) uint {
	if nil == appProperty {
		return 0
	}
	if value, exist := appProperty.ValueMap[key]; exist {
		return util.ToUInt(parseVariableOfString(util.ToString(value)))
	}
	return 0
}

func GetValueUInt8(key string) uint8 {
	if nil == appProperty {
		return 0
	}
	if value, exist := appProperty.ValueMap[key]; exist {
		return util.ToUInt8(parseVariableOfString(util.ToString(value)))
	}
	return 0
}

func GetValueUInt16(key string) uint16 {
	if nil == appProperty {
		return 0
	}
	if value, exist := appProperty.ValueMap[key]; exist {
		return util.ToUInt16(parseVariableOfString(util.ToString(value)))
	}
	return 0
}

func GetValueUInt32(key string) uint32 {
	if nil == appProperty {
		return 0
	}
	if value, exist := appProperty.ValueMap[key]; exist {
		return util.ToUInt32(parseVariableOfString(util.ToString(value)))
	}
	return 0
}

func GetValueUInt64(key string) uint64 {
	if nil == appProperty {
		return 0
	}
	if value, exist := appProperty.ValueMap[key]; exist {
		return util.ToUInt64(parseVariableOfString(util.ToString(value)))
	}
	return 0
}

func GetValueFloat32(key string) float32 {
	if nil == appProperty {
		return 0
	}
	if value, exist := appProperty.ValueMap[key]; exist {
		return util.ToFloat32(parseVariableOfString(util.ToString(value)))
	}
	return 0
}

func GetValueFloat64(key string) float64 {
	if nil == appProperty {
		return 0
	}
	if value, exist := appProperty.ValueMap[key]; exist {
		return util.ToFloat64(parseVariableOfString(util.ToString(value)))
	}
	return 0
}

func GetValueBool(key string) bool {
	if nil == appProperty {
		return false
	}
	if value, exist := appProperty.ValueMap[key]; exist {
		return util.ToBool(parseVariableOfString(util.ToString(value)))
	}
	return false
}

func GetValueStringDefault(key, defaultValue string) string {
	if nil == appProperty {
		return defaultValue
	}
	if value, exist := appProperty.ValueMap[key]; exist {
		return util.ToString(parseVariableOfString(util.ToString(value)))
	}
	return defaultValue
}

func GetValueIntDefault(key string, defaultValue int) int {
	if nil == appProperty {
		return defaultValue
	}
	if value, exist := appProperty.ValueMap[key]; exist {
		return util.ToInt(parseVariableOfString(util.ToString(value)))
	}
	return defaultValue
}

func GetValueInt8Default(key string, defaultValue int8) int8 {
	if nil == appProperty {
		return defaultValue
	}
	if value, exist := appProperty.ValueMap[key]; exist {
		return util.ToInt8(parseVariableOfString(util.ToString(value)))
	}
	return defaultValue
}

func GetValueInt16Default(key string, defaultValue int16) int16 {
	if nil == appProperty {
		return defaultValue
	}
	if value, exist := appProperty.ValueMap[key]; exist {
		return util.ToInt16(parseVariableOfString(util.ToString(value)))
	}
	return defaultValue
}

func GetValueInt32Default(key string, defaultValue int32) int32 {
	if nil == appProperty {
		return defaultValue
	}
	if value, exist := appProperty.ValueMap[key]; exist {
		return util.ToInt32(parseVariableOfString(util.ToString(value)))
	}
	return defaultValue
}

func GetValueInt64Default(key string, defaultValue int64) int64 {
	if nil == appProperty {
		return defaultValue
	}
	if value, exist := appProperty.ValueMap[key]; exist {
		return util.ToInt64(parseVariableOfString(util.ToString(value)))
	}
	return defaultValue
}

func GetValueUIntDefault(key string, defaultValue uint) uint {
	if nil == appProperty {
		return defaultValue
	}
	if value, exist := appProperty.ValueMap[key]; exist {
		return util.ToUInt(parseVariableOfString(util.ToString(value)))
	}
	return defaultValue
}

func GetValueUInt8Default(key string, defaultValue uint8) uint8 {
	if nil == appProperty {
		return defaultValue
	}
	if value, exist := appProperty.ValueMap[key]; exist {
		return util.ToUInt8(parseVariableOfString(util.ToString(value)))
	}
	return defaultValue
}

func GetValueUInt16Default(key string, defaultValue uint16) uint16 {
	if nil == appProperty {
		return defaultValue
	}
	if value, exist := appProperty.ValueMap[key]; exist {
		return util.ToUInt16(parseVariableOfString(util.ToString(value)))
	}
	return defaultValue
}

func GetValueUInt32Default(key string, defaultValue uint32) uint32 {
	if nil == appProperty {
		return defaultValue
	}
	if value, exist := appProperty.ValueMap[key]; exist {
		return util.ToUInt32(parseVariableOfString(util.ToString(value)))
	}
	return defaultValue
}

func GetValueUInt64Default(key string, defaultValue uint64) uint64 {
	if nil == appProperty {
		return defaultValue
	}
	if value, exist := appProperty.ValueMap[key]; exist {
		return util.ToUInt64(parseVariableOfString(util.ToString(value)))
	}
	return defaultValue
}

func GetValueFloat32Default(key string, defaultValue float32) float32 {
	if nil == appProperty {
		return defaultValue
	}
	if value, exist := appProperty.ValueMap[key]; exist {
		return util.ToFloat32(parseVariableOfString(util.ToString(value)))
	}
	return defaultValue
}

func GetValueFloat64Default(key string, defaultValue float64) float64 {
	if nil == appProperty {
		return defaultValue
	}
	if value, exist := appProperty.ValueMap[key]; exist {
		return util.ToFloat64(parseVariableOfString(util.ToString(value)))
	}
	return defaultValue
}

func GetValueBoolDefault(key string, defaultValue bool) bool {
	if nil == appProperty {
		return defaultValue
	}
	if value, exist := appProperty.ValueMap[key]; exist {
		return util.ToBool(parseVariableOfString(util.ToString(value)))
	}
	return defaultValue
}

func GetValueObject(key string, targetPtrObj any) error {
	if nil == appProperty {
		return nil
	}
	data := doGetValue(appProperty.ValueDeepMap, key)
	_, err := util.DataToEntity(data, targetPtrObj)
	if err != nil {
		return err
	}
	return nil
}

func GetValueArray(key string) []any {
	if nil == appProperty {
		return nil
	}

	var arrayResult []any
	data := doGetValue(appProperty.ValueDeepMap, key)
	_, err := util.DataToEntity(data, &arrayResult)
	if err != nil {
		return arrayResult
	}
	return arrayResult
}

func GetValueArrayInt(key string) []int {
	if nil == appProperty {
		return nil
	}

	var arrayResult []int
	data := doGetValue(appProperty.ValueDeepMap, key)
	_, err := util.DataToEntity(data, &arrayResult)
	if err != nil {
		return arrayResult
	}
	return arrayResult
}

func GetValueArrayString(key string) []string {
	if nil == appProperty {
		return nil
	}

	var arrayResult []string
	data := doGetValue(appProperty.ValueDeepMap, key)
	_, err := util.DataToEntity(data, &arrayResult)
	if err != nil {
		return arrayResult
	}
	return arrayResult
}

func GetValue(key string) any {
	if nil == appProperty {
		return nil
	}
	return doGetValue(appProperty.ValueDeepMap, key)
}

func doGetValue(parentValue any, key string) any {
	if key == "" {
		return parseVariableOfObj(parentValue)
	}
	parentValueKind := reflect.ValueOf(parentValue).Kind()
	if parentValueKind == reflect.Map {
		keys := strings.SplitN(key, ".", 2)
		v1 := reflect.ValueOf(parentValue).MapIndex(reflect.ValueOf(keys[0]))
		emptyValue := reflect.Value{}
		if v1 == emptyValue {
			return nil
		}
		if len(keys) == 1 {
			return doGetValue(v1.Interface(), "")
		} else {
			return doGetValue(v1.Interface(), fmt.Sprintf("%v", keys[1]))
		}
	}
	return nil
}

func parseVariableOfObj(value any) any {
	valType := reflect.TypeOf(value)
	if valType.Kind() == reflect.String {
		return parseVariableOfString(value.(string))
	} else if valType.Kind() == reflect.Map {
		valueValue := reflect.ValueOf(value)
		dataMap := map[string]interface{}{}
		for mapR := valueValue.MapRange(); mapR.Next(); {
			dataMap[mapR.Key().Interface().(string)] = parseVariableOfObj(mapR.Value().Interface())
		}
		return dataMap
	} else {
		return value
	}
}

func parseVariableOfString(value string) string {
	if !strings.Contains(value, "${") || !strings.Contains(value, "}") {
		return value
	}

	keys, defaultKeyValueMap, formatValue := parseAllVariable(value)
	var values []any
	for _, key := range keys {
		val := GetValue(key)
		if val == nil {
			val = defaultKeyValueMap[key]
		}
		values = append(values, val)
	}
	return fmt.Sprintf(formatValue, values...)
}

func parseAllVariable(value string) ([]string, map[string]any, string) {
	if !strings.Contains(value, "${") || !strings.Contains(value, "}") {
		return nil, nil, value
	}
	var keys []string
	beginIndex := strings.Index(value, "${")
	if beginIndex == -1 {
		return keys, nil, value
	}
	endIndex := strings.Index(value, "}")
	if endIndex == -1 || endIndex == 0 {
		return keys, nil, value
	}

	defaultKeyValueMap := map[string]any{}
	keyName := value[beginIndex+2 : endIndex]
	keyName = strings.TrimSpace(keyName)
	if keyName != "" {
		relKeyName, defaultValue := parseKeyForDefaultValue(keyName)
		keys = append(keys, relKeyName)
		defaultKeyValueMap[relKeyName] = defaultValue
		nextKeys, nextDefaultKeyValueMap, endStr := parseAllVariable(value[endIndex+1:])

		if nextDefaultKeyValueMap != nil {
			for k, v := range nextDefaultKeyValueMap {
				defaultKeyValueMap[k] = v
			}
		}

		if nextKeys != nil && len(nextKeys) > 0 {
			keys = append(keys, nextKeys...)
		}
		return keys, defaultKeyValueMap, value[:beginIndex] + "%v" + endStr
	} else {
		nextVars, nextDefaultKeyValueMap, endStr := parseAllVariable(value[endIndex+1:])

		if nextDefaultKeyValueMap != nil {
			for k, v := range nextDefaultKeyValueMap {
				defaultKeyValueMap[k] = v
			}
		}

		if nextVars != nil && len(nextVars) > 0 {
			keys = append(keys, nextVars...)
		}
		return keys, defaultKeyValueMap, value[:endIndex+1] + endStr
	}
}

func parseKeyForDefaultValue(keyName string) (string, string) {
	if !strings.Contains(keyName, ":") {
		return keyName, ""
	}
	index := strings.Index(keyName, ":")
	if index == -1 {
		return keyName, ""
	}
	return strings.TrimSpace(keyName[:index]), strings.TrimSpace(keyName[index+1:])
}

type ApplicationProperty struct {
	ValueMap     map[string]any
	ValueDeepMap map[string]any
}

// LoadYamlConfig read fileName from private path fileName,eg:application.yml, and transform it to AConfig
// note: AConfig must be a pointer
func LoadYamlConfig(fileName string, AConfig any, handler func(data []byte, AConfig any) error) error {
	pwd, _ := os.Getwd()
	fp := filepath.Join(pwd, fileName)
	return LoadYamlConfigByAbsolutPath(fp, AConfig, handler)
}

// LoadYamlConfigByAbsolutPath read fileName from absolute path fileName,eg:/home/base/application.yml, and transform it to AConfig
// note: AConfig must be a pointer
func LoadYamlConfigByAbsolutPath(path string, AConfig any, handler func(data []byte, AConfig any) error) error {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Printf("读取文件异常(%v)", err)
	}
	return handler(data, AConfig)
}

//LoadSpringConfig read fileName from current dictionary and fileName is application.yml,eg:/home/base/application.yml, and transform it to AConfig
//note: AConfig must be a pointer
//note: if it has Spring.Profiles.Active,eg: Spring.Profiles.Active=dev,will load config from /home/base/application-dev.yml,and same key
//will write in the last one.

func LoadSpringConfig(AConfig any) {
	_ = LoadYamlConfig("application.yml", AConfig, func(data []byte, AConfig any) error {
		err := yaml.Unmarshal(data, AConfig)
		if err != nil {
			log.Printf("读取 application.yml 异常(%v)", err)
			return err
		}
		v1 := reflect.ValueOf(AConfig).Elem()
		o1 := v1.FieldByName("Spring").Interface()
		v2 := reflect.ValueOf(o1)
		o2 := v2.FieldByName("Profiles").Interface()
		v3 := reflect.ValueOf(o2)
		act := v3.FieldByName("Active").String()
		if act != "" && act != "default" {
			yamlAdditional, err := os.ReadFile(fmt.Sprintf("./application-%s.yml", act))
			if err != nil {
				log.Printf("读取 application-%s.yml 失败", act)
				return err
			} else {
				err = yaml.Unmarshal(yamlAdditional, AConfig)
				if err != nil {
					log.Printf("读取 application-%s.yml 异常", act)
					return err
				}
			}
		}
		return nil
	})
}
