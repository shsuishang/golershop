package phone

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/nyaruka/phonenumbers"
	"log"
	"strconv"
	"strings"
)

var (
	// 直辖市
	MUNICIPALITY = []string{"北京市", "天津市", "上海市", "重庆市"}
	// 自治区
	AUTONOMOUS_REGION = []string{"新疆", "内蒙古", "西藏", "宁夏", "广西"}
	// 中国大陆区区号
	COUNTRY_CODE_CHINA = 86
)

// CheckPhoneNumber 判断手机号是否有效（中国）
func CheckPhoneNumber(phoneNumber string) bool {
	return checkPhoneNumber(COUNTRY_CODE_CHINA, phoneNumber)
}

// CheckPhoneNumber 判断手机号是否有效（国际）
func checkPhoneNumber(countryCode int, phoneNumber string) bool {
	num, err := strconv.ParseInt(phoneNumber, 10, 64)
	if err != nil {
		log.Printf("Error parsing phone number: %v", err)
		return false
	}
	pn := &phonenumbers.PhoneNumber{
		CountryCode:    proto.Int32(int32(countryCode)),
		NationalNumber: proto.Uint64(uint64(num)),
	}
	return phonenumbers.IsValidNumber(pn)
}

// GetCarrier 根据手机号 判断手机运营商
func GetCarrier(phoneNumber string) string {
	num, err := strconv.ParseInt(phoneNumber, 10, 64)
	if err != nil {
		log.Printf("Error parsing phone number: %v", err)
		return ""
	}
	pn := &phonenumbers.PhoneNumber{
		CountryCode:    proto.Int32(int32(COUNTRY_CODE_CHINA)),
		NationalNumber: proto.Uint64(uint64(num)),
	}
	carrierEn, err := phonenumbers.GetCarrierForNumber(pn, "en")
	var carrierZh string
	switch carrierEn {
	case "China Mobile":
		carrierZh = "移动"
	case "China Unicom":
		carrierZh = "联通"
	case "China Telecom":
		carrierZh = "电信"
	default:
		carrierZh = ""
	}
	return carrierZh
}

// GetGeo 根据手机号 获取手机归属地
func GetGeo(phoneNumber string) string {
	return getGeo(COUNTRY_CODE_CHINA, phoneNumber)
}

// GetGeo 根据手机号 获取手机归属地
func getGeo(code int, phoneNumber string) string {
	num, err := strconv.ParseInt(phoneNumber, 10, 64)
	if err != nil {
		log.Printf("Error parsing phone number: %v", err)
		return ""
	}
	pn := &phonenumbers.PhoneNumber{
		CountryCode:    proto.Int32(int32(code)),
		NationalNumber: proto.Uint64(uint64(num)),
	}
	return phonenumbers.GetRegionCodeForNumber(pn)
}

// PhoneModel 包含手机信息的结构体
type PhoneModel struct {
	CountryCode    int
	CountryCodeStr string
	NationalNumber uint64
	ProvinceName   string
	CityName       string
	CountryName    string
}

// GetPhoneModelWithCountry 输入手机号码，返回手机号信息
func GetPhoneModelWithCountry(phoneNumber string) *PhoneModel {
	if !IsValidNumber(phoneNumber) {
		return nil
	}
	pn, err := phonenumbers.Parse(phoneNumber, "CH")
	if err != nil {
		log.Printf("Error parsing phone number: %v", err)
		return nil
	}
	countryCode := pn.GetCountryCode()
	nationalNumber := pn.GetNationalNumber()
	model := &PhoneModel{
		CountryCode:    int(countryCode),
		CountryCodeStr: fmt.Sprintf("+%d", countryCode),
		NationalNumber: nationalNumber,
	}
	geo := phonenumbers.GetRegionCodeForNumber(pn)
	// 直辖市
	for _, val := range MUNICIPALITY {
		if geo == val {
			model.ProvinceName = strings.TrimSuffix(val, "市")
			model.CityName = val
			return model
		}
	}
	// 自治区
	for _, val := range AUTONOMOUS_REGION {
		if strings.HasPrefix(geo, val) {
			model.ProvinceName = val
			model.CityName = strings.TrimPrefix(geo, val)
			return model
		}
	}
	// 其它
	if strings.Contains(geo, "省") {
		splitArr := strings.Split(geo, "省")
		if len(splitArr) == 2 {
			model.ProvinceName = splitArr[0]
			model.CityName = splitArr[1]
			return model
		}
	}
	model.CountryName = geo
	return model
}

// IsValidNumber 判定手机号是否可用
func IsValidNumber(phoneNumber string) bool {
	pn, err := phonenumbers.Parse(phoneNumber, "CH")
	if err != nil {
		log.Printf("Error parsing phone number: %v", err)
		return false
	}
	return phonenumbers.IsValidNumber(pn)
}
