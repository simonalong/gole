package tools

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

/**
 *  1.yaml <---> properties
 *  2.yaml <---> json
 *  3.yaml <---> map
 *  4.yaml <---> list
 *  5.yaml <---> kvList
 */

// NewLine 换行符
var NewLine = "\n"

// RemarkPre 注释标识
var RemarkPre = "# "

// IndentBlanks 缩进空格
var IndentBlanks = "  "

// SignSemicolon 分号连接符
var SignSemicolon = ":"

// SignEqual 等号连接符
var SignEqual = "="

// Dot 点
var Dot = "."

// ArrayBlanks 数组缩进
var ArrayBlanks = "- "

// NewLineDom yaml的value换行符
var YamlNewLineDom = "|\n"

var rangePattern = regexp.MustCompile("^(.*)\\[(\\d*)\\]$")

type KeyValue struct {
	Key   string
	Value interface{}
}

func YamlToProperties(contentOfYaml string) (string, error) {
	// yaml 到 map
	dataMap, err := YamlToMap(contentOfYaml)
	if err != nil {
		log.Fatalf("YamlToPropertiesStr error: %v", err)
		return "", err
	}

	return MapToProperties(dataMap)
}

func PropertiesStrToYaml(contentOfProperties string) string {
	var yamlLineList []string
	var yamlNodes []YamlNode
	propertiesLineWordList := getPropertiesItemLineList(contentOfProperties)
	for _, line := range propertiesLineWordList {
		line = strings.TrimSpace(line)
		if line != "" {
			// 注释数据不要
			if strings.HasPrefix(line, "#") {
				continue
			}

			index := strings.Index(line, "=")
			if index > -1 {
				key := line[:index]
				value := line[index+1:]

				if strings.Contains(value, "\n") {
					value = YamlNewLineDom + value
				}

				lineWordList := strings.Split(key, "\\.")
				lineWordList, yamlNodes = wordToNode(lineWordList, yamlNodes, nil, false, -1, appendSpaceForArrayValue(value));
			}
		}
	}
	formatPropertiesToYaml(yamlLineList, yamlNodes, false, "");
	return strings.Join(yamlLineList, "\n") + "\n"
}

func YamlToMap(contentOfYaml string) (map[string]interface{}, error) {
	resultMap := make(map[string]interface{})
	err := yaml.Unmarshal([]byte(contentOfYaml), &resultMap)
	if err != nil {
		log.Fatalf("YamlToMap error: %v", err)
		return nil, err
	}

	return resultMap, nil
}

func MapToYaml(dataMap map[string]interface{}) string {
	bytes2, err := yaml.Marshal(dataMap)
	if err != nil {
		log.Fatalf("MapToYaml error: %v", err)
		return ""
	}
	return string(bytes2)
}

//
//func YamlToJson(contentOfYaml string) string {
//
//}
//
//func JsonToYaml(contentOfJson string) string {
//
//}
//
//func YamlToList(contentOfYaml string) []interface{} {
//
//}
//
//func KvListToYaml(kvList []KeyValue) string {
//
//}

// 进行深层嵌套的map数据处理
func MapToProperties(dataMap map[string]interface{}) (string, error) {
	propertyStrList := []string{}
	for key, value := range dataMap {
		valueKind := reflect.TypeOf(value).Kind()
		switch valueKind {
		case reflect.Map:
			{
				propertyStrList = doMapToProperties(propertyStrList, value, prefixWithDOT("")+key)
			}
		case reflect.Array, reflect.Slice:
			{
				objectValue := reflect.ValueOf(value)
				for index := 0; index < objectValue.Len(); index++ {
					propertyStrList = doMapToProperties(propertyStrList, objectValue.Index(index), prefixWithDOT("")+key+"["+strconv.Itoa(index)+"]")
				}
			}
		case reflect.String:
			objectValue := reflect.ValueOf(value)
			objectValueStr := strings.ReplaceAll(objectValue.String(), "\n", "\\\n")
			propertyStrList = append(propertyStrList, prefixWithDOT("")+key+"="+objectValueStr)
		default:
			propertyStrList = append(propertyStrList, prefixWithDOT("")+key+"="+fmt.Sprintf("%v", value))
		}
	}
	resultStr := ""
	for _, propertyStr := range propertyStrList {
		resultStr += propertyStr + "\n"
	}

	return resultStr, nil
}

func getPropertiesItemLineList(content string) []string {
	if "" == content {
		return []string{}
	}

	lineList := strings.Split(content, NewLine)
	var itemLineList []string
	var stringAppender string
	for _, line := range lineList {
		if strings.HasSuffix(content, "\\") {
			stringAppender += line + "\n"
		} else {
			stringAppender += line
			itemLineList = append(itemLineList, stringAppender)
			stringAppender = ""
		}
	}
	return itemLineList;
}

func formatPropertiesToYaml(yamlLineList []string, yamlNodes []YamlNode, lastNodeArrayFlag bool, blanks string) []string {
	var beforeNodeIndex int
	var equalSign string

	for _, yamlNode := range yamlNodes {
		value := yamlNode.value

		equalSign = SignSemicolon
		if "" == value {
			value = ""
		} else {
			equalSign = SignSemicolon + " "
		}
	}


	Integer beforeNodeIndex = null;
	String equalSign;
	for (YamlNode YamlNode : YamlNodes) {
		String value = YamlNode.getValue();
		String remark = YamlNode.getRemark();

		equalSign = SIGN_SEMICOLON;
		if (null == value || "".equals(value)) {
			value = "";
		} else {
			equalSign = SIGN_SEMICOLON + " ";
		}
		YamlNode.resortValue();

		String name = YamlNode.getName();
		if (lastNodeArrayFlag) {
			if (null == name) {
				yamlLineList.add(blanks + ARRAY_BLANKS + stringValueWrap(value));
			} else {
				if (null != beforeNodeIndex && beforeNodeIndex.equals(YamlNode.getLastNodeIndex())) {
					yamlLineList.add(blanks + INDENT_BLANKS + name + equalSign + stringValueWrap(value));
				} else {
					yamlLineList.add(blanks + ARRAY_BLANKS + name + equalSign + stringValueWrap(value));
				}
			}
			beforeNodeIndex = YamlNode.getLastNodeIndex();
		} else {
			// 父节点为空，表示，当前为顶层
			if (null == YamlNode.getParent()) {
				String remarkTem = getRemarkProject(YamlNode.getProjectRemark());
				if (!"".equals(remarkTem)) {
					yamlLineList.add(blanks + getRemarkProject(YamlNode.getProjectRemark()));
				}
			}

			// 自己节点为数组，则添加对应的注释
			if (YamlNode.getArrayFlag()) {
				if (null != remark && !"".equals(remark)) {
					yamlLineList.add(blanks + remark);
				}
			}
			yamlLineList.add(blanks + name + equalSign + stringValueWrap(value, remark));
		}

		if (YamlNode.getArrayFlag()) {
			if (lastNodeArrayFlag) {
				formatPropertiesToYaml(yamlLineList, YamlNode.getValueList(), true, INDENT_BLANKS + INDENT_BLANKS + blanks);
			} else {
				formatPropertiesToYaml(yamlLineList, YamlNode.getValueList(), true, INDENT_BLANKS + blanks);
			}
		} else {
			if (lastNodeArrayFlag) {
				formatPropertiesToYaml(yamlLineList, YamlNode.getChildren(), false, INDENT_BLANKS + INDENT_BLANKS + blanks);
			} else {
				formatPropertiesToYaml(yamlLineList, YamlNode.getChildren(), false, INDENT_BLANKS + blanks);
			}
		}
	}
}

func wordToNode(lineWordList []string, nodeList []YamlNode, parentNode *YamlNode, lastNodeArrayFlag bool, index int, value string) ([]string, []YamlNode) {
	if len(lineWordList) == 0 {
		if lastNodeArrayFlag {
			node := YamlNode{value: value}
			nodeList = append(nodeList, node)
		}
	} else {
		nodeName := lineWordList[0]
		nodeName, nextIndex := peelArray(nodeName);

		var node YamlNode
		if nil != parentNode {
			node = YamlNode{name: nodeName, parent: parentNode, lastNodeIndex: index}
		} else {
			node = YamlNode{name: nodeName, lastNodeIndex: index}
		}
		lineWordList = lineWordList[1:]

		//如果节点下面的子节点数量为0，则为终端节点，也就是赋值节点
		if len(lineWordList) == 0 {
			if -1 == nextIndex {
				node.value = value
			}
		}

		// nextIndex 不空，表示当前节点为数组，则之后的数据为他的节点数据
		if nextIndex != -1 {
			node.arrayFlag = true
			var hasEqualsName = false

			//遍历查询节点是否存在
			for _, yamlNode := range nodeList {
				//如果节点名称已存在，则递归添加剩下的数据节点
				if nodeName == yamlNode.name && yamlNode.arrayFlag {
					yamlNodeIndex := yamlNode.lastNodeIndex
					if 0 == yamlNodeIndex || index == yamlNodeIndex {
						hasEqualsName = true
						lineWordList, yamlNode.valueList = wordToNode(lineWordList, yamlNode.valueList, node.parent, true, nextIndex, appendSpaceForArrayValue(value))
					}
				}
			}

			//如果遍历结果为节点名称不存在，则递归添加剩下的数据节点，并把新节点添加到上级yamlTree的子节点中
			if !hasEqualsName {
				lineWordList, node.valueList = wordToNode(lineWordList, node.valueList, node.parent, true, nextIndex, appendSpaceForArrayValue(value))
			}
		} else {
			var hasEqualsName = false
			for _, yamlNode := range nodeList {
				if !lastNodeArrayFlag {
					//如果节点名称已存在，则递归添加剩下的数据节点
					if nodeName == yamlNode.name {
						hasEqualsName = true
						lineWordList, yamlNode.children = wordToNode(lineWordList, yamlNode.children, &yamlNode, false, nextIndex, appendSpaceForArrayValue(value))
					}
				} else {
					//如果节点名称已存在，则递归添加剩下的数据节点
					if nodeName == yamlNode.name {
						yamlNodeIndex := yamlNode.lastNodeIndex
						if -1 == yamlNodeIndex {
							hasEqualsName = true
							lineWordList, yamlNode.children = wordToNode(lineWordList, yamlNode.children, &yamlNode, true, nextIndex, appendSpaceForArrayValue(value))
						}
					}
				}
			}

			//如果遍历结果为节点名称不存在，则递归添加剩下的数据节点，并把新节点添加到上级yamlTree的子节点中
			if !hasEqualsName {
				lineWordList, node.children = wordToNode(lineWordList, node.children, &node, false, nextIndex, appendSpaceForArrayValue(value))
				nodeList = append(nodeList, node)
			}
		}
	}
	return lineWordList, nodeList
}

func doMapToProperties(propertyStrList []string, value interface{}, prefix string) []string {
	valueKind := reflect.TypeOf(value).Kind()
	switch valueKind {
	case reflect.Map:
		{
			// map结构
			if reflect.ValueOf(value).Len() == 0 {
				return []string{}
			}

			for mapR := reflect.ValueOf(value).MapRange(); mapR.Next(); {
				mapKey := mapR.Key().Interface()
				mapValue := mapR.Value().Interface()
				propertyStrList = doMapToProperties(propertyStrList, mapValue, prefixWithDOT(prefix)+fmt.Sprintf("%v", mapKey))
			}
		}
	case reflect.Array, reflect.Slice:
		{
			objectValue := reflect.ValueOf(value)
			for index := 0; index < objectValue.Len(); index++ {
				propertyStrList = doMapToProperties(propertyStrList, objectValue.Index(index), prefix+"["+strconv.Itoa(index)+"]")
			}
		}
	case reflect.String:
		objectValue := reflect.ValueOf(value)
		objectValueStr := strings.ReplaceAll(objectValue.String(), "\n", "\\\n")
		propertyStrList = append(propertyStrList, prefix+"="+objectValueStr)
	default:
		objectValue := fmt.Sprintf("%v", reflect.ValueOf(value))
		propertyStrList = append(propertyStrList, prefix+"="+objectValue)
	}
	return propertyStrList
}

func prefixWithDOT(prefix string) string {
	if "" == prefix {
		return ""
	}
	return prefix + "."
}

func peelArray(nodeName string) (string, int) {
	var index = -1
	var name string
	var err error

	subData := rangePattern.FindAllStringSubmatch(nodeName, -1)
	if len(subData) > 0 {
		name = subData[0][1]
		indexStr := subData[0][2]
		if "" != indexStr {
			index, err = strconv.Atoi(indexStr)
			if err != nil {
				log.Fatalf("解析错误, nodeName=" + nodeName)
				return "", -1
			}
		}
	}
	return name, index
}

//
// 将yaml对应的这种value进行添加前缀空格，其中value为key1对应的value
// test:
//   key1: |
//     value1
//     value2
//     value3
// 对应的值
// {@code
// |
// value1
// value2
// value3
// }
//
// @param value 待转换的值比如{@code
//              test:
//              key1: |
//              value1
//              value2
//              value3
//              }
// @return 添加前缀空格之后的处理
// {@code
// |
// value1
// value2
// value3
// }
//
func appendSpaceForArrayValue(value string) string {
	if !strings.HasPrefix(value, YamlNewLineDom) {
		return value
	}

	value = value[len(YamlNewLineDom):]
	valueTems := strings.Split(value, "\\n")

	strs := []string{}
	for _, element := range valueTems {
		tem := element
		if strings.HasSuffix(element, "\\") {
			tem = element[:len(element)-1]
		}
		strs = append(strs, IndentBlanks + tem)
	}
	return YamlNewLineDom + strings.Join(strs, "\n")
}

type YamlNode struct {

	// 父节点
	parent *YamlNode

	// 只有parent为null时候，该值才可能有值
	projectRemark string

	// name
	name string

	// value
	value string

	// 子节点
	children []YamlNode

	// 数组标示
	arrayFlag bool

	// 存储的数组中的前一个节点的下标
	lastNodeIndex int

	// 只有数组标示为true，下面的value才有值
	valueList []YamlNode
}

func (yamlNode *YamlNode) resortValue()  {
	// todo
	//
	if (!arrayFlag || valueList.isEmpty()) {
		return;
	}

	// 升序
	valueList.sort((a, b) -> {
		if (null == a.getLastNodeIndex() || null == b.getLastNodeIndex()) {
			return 0;
		}

		return a.getLastNodeIndex() - b.getLastNodeIndex();
	});

	// 是数组的节点也循环下
	valueList.forEach(YamlNode::resortValue);
}
