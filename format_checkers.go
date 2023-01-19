package gojsonschemaformatcheckers

import (
	"regexp"
	"strconv"

	"github.com/xeipuuv/gojsonschema"
)

var (
	rxPhone    = regexp.MustCompile(`/^1[3456789]\d{9}$/`)
	rxIdCard   = regexp.MustCompile(`/^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$/`)
	rxIdCard1  = regexp.MustCompile(`/^[1-9]\d{5}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}$/`)
	rxPostCode = regexp.MustCompile(`/^[1-9]{1}(\d+){5}$/`)
)

// 加载当前包,即注册
func init() {
	gojsonschema.FormatCheckers.Add("number", NumberFormatChecker{})     // 数字格式验证
	gojsonschema.FormatCheckers.Add("phone", PhoneFormatChecker{})       // 手机号格式验证
	gojsonschema.FormatCheckers.Add("idCard", IDCardFormatChecker{})     // 身份证号格式验证
	gojsonschema.FormatCheckers.Add("postCode", PostCodeFormatChecker{}) // 邮编格式验证
}

type NumberFormatChecker struct{}

// IsFormat checks if input is a correctly formatted number string
func (f NumberFormatChecker) IsFormat(input interface{}) bool {
	asString, ok := input.(string)
	if !ok {
		return false
	}
	_, err := strconv.ParseFloat(asString, 64)
	return err == nil
}

type PhoneFormatChecker struct{}

// IsFormat checks if input is a correctly formatted phone string
func (f PhoneFormatChecker) IsFormat(input interface{}) bool {
	asString, ok := input.(string)
	if !ok {
		return false
	}
	out := rxPhone.MatchString(asString)
	return out
}

type IDCardFormatChecker struct{}

// IsFormat checks if input is a correctly formatted IDCard string
func (f IDCardFormatChecker) IsFormat(input interface{}) bool {
	asString, ok := input.(string)
	if !ok {
		return false
	}
	out := rxIdCard.MatchString(asString) || rxIdCard1.MatchString(asString)
	return out
}

type PostCodeFormatChecker struct{}

// IsFormat checks if input is a correctly formatted postcode string
func (f PostCodeFormatChecker) IsFormat(input interface{}) bool {
	asString, ok := input.(string)
	if !ok {
		return false
	}
	out := rxPostCode.MatchString(asString)
	return out
}
