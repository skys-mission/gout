// Copyright (c) 2024, https://github.com/skys-mission and SoyMilkWhisky

package iplocation

// baseIplocationURL 是获取IP位置信息的基本API地址。
const (
	baseIplocationURL = "https://api.iplocation.net"
	// ipv4IplocationURL 是专门用于获取IPv4地址位置信息的API地址。
	ipv4IplocationURL = "http://ipv4.iplocation.net"

	// getIPAddressSubURLPath 是获取IP地址的子URL路径。
	getIPAddressSubURLPath = "/?cmd=get-ip"
)

// ResponseCodeForIPLocator 定义了IP定位服务的响应码类型。
// 它使用字符串来表示不同的响应状态，以便于在处理IP定位请求时进行状态识别和错误处理。
type ResponseCodeForIPLocator string

// ResponseCodeForGetIPAddress 定义了获取IP地址服务的响应码类型。
// 它使用整数来表示不同的响应状态，使得在处理获取IP地址请求时可以量化地评估操作结果。
type ResponseCodeForGetIPAddress int64

// ResponseMessage 定义了与响应状态码相对应的简短描述类型
type ResponseMessage string

// ResponseIPVersion 定义了IP版本4or6
type ResponseIPVersion int8

// ISO3166V1Alpha2 定义了ISO 3166-1 Alpha-2标准
type ISO3166V1Alpha2 string

// ResponseCodeForIPLocator定义了IP定位服务响应状态码的常量
const (
	ResponseCodeOKForIPLocator         ResponseCodeForIPLocator = "200" // 表示请求成功
	ResponseCodeBadRequestForIPLocator ResponseCodeForIPLocator = "400" // 表示请求参数错误
	ResponseCodeNotFoundForIPLocator   ResponseCodeForIPLocator = "404" // 表示未找到对应的记录
)

// ResponseCodeForGetIPAddress定义了获取IP地址服务响应状态码的常量
const (
	ResponseCodeOKForGetIPAddress         ResponseCodeForGetIPAddress = 200 // 表示请求成功
	ResponseCodeBadRequestForGetIPAddress ResponseCodeForGetIPAddress = 400 // 表示请求参数错误
	ResponseCodeNotFoundForGetIPAddress   ResponseCodeForGetIPAddress = 404 // 表示未找到对应的记录
)

// 响应消息常量定义
const (
	// ResponseMessageOK 对应于状态码200的描述
	ResponseMessageOK ResponseMessage = "OK"
	// ResponseMessageBadRequest 对应于状态码400的描述
	ResponseMessageBadRequest ResponseMessage = "Bad Request"
	// ResponseMessageNotFound 对应于状态码404的描述
	ResponseMessageNotFound ResponseMessage = "Not Found"
)

// ResponseIPVersion 定义了IP版本的常量
const (
	// ResponseIPVersion4 代表IPv4版本的常量，值为4
	ResponseIPVersion4 ResponseIPVersion = 4
	// ResponseIPVersion6 代表IPv6版本的常量，值为6
	ResponseIPVersion6 ResponseIPVersion = 6
)

// ISO3166V1Alpha2 定义了ISO 3166-1 Alpha-2标准，排名不分先后，部分代码由AI提示完成
const (
	LocalCodeChinaMainland ISO3166V1Alpha2 = "CN"
	LocalCodeChinaTaiwan   ISO3166V1Alpha2 = "TW"
	LocalCodeChinaHongKong ISO3166V1Alpha2 = "HK"
	LocalCodeChinaMacao    ISO3166V1Alpha2 = "MO"
	LocalCodeUS            ISO3166V1Alpha2 = "US"
	LocalCodeRussian       ISO3166V1Alpha2 = "RU"
	LocalCodeJapan         ISO3166V1Alpha2 = "JP"
	LocalCodeSpain         ISO3166V1Alpha2 = "ES"
	LocalCodeGermany       ISO3166V1Alpha2 = "DE"
	LocalCodeItaly         ISO3166V1Alpha2 = "IT"
	LocalCodeFrance        ISO3166V1Alpha2 = "FR"
	LocalCodeUK            ISO3166V1Alpha2 = "GB"
	LocalCodeCanada        ISO3166V1Alpha2 = "CA"
	LocalCodeAustralia     ISO3166V1Alpha2 = "AU"
	LocalCodeIndia         ISO3166V1Alpha2 = "IN"
	LocalCodeSaudiArabia   ISO3166V1Alpha2 = "SA"
	LocalCodeBrazil        ISO3166V1Alpha2 = "BR"
	LocalCodeChile         ISO3166V1Alpha2 = "CL"
	LocalCodeMexico        ISO3166V1Alpha2 = "MX"
	LocalCodeArgentina     ISO3166V1Alpha2 = "AR"
	LocalCodeColombia      ISO3166V1Alpha2 = "CO"
	LocalCodePeru          ISO3166V1Alpha2 = "PE"
	LocalCodeVenezuela     ISO3166V1Alpha2 = "VE"
	LocalCodeThailand      ISO3166V1Alpha2 = "TH"
	LocalCodeIndonesia     ISO3166V1Alpha2 = "ID"
	LocalCodeMalaysia      ISO3166V1Alpha2 = "MY"
	LocalCodePhilippines   ISO3166V1Alpha2 = "PH"
	LocalCodeNetherlands   ISO3166V1Alpha2 = "NL"
	LocalCodeVietName      ISO3166V1Alpha2 = "VN"
	LocalCodeSingapore     ISO3166V1Alpha2 = "SG"
	LocalCodeSouthKorea    ISO3166V1Alpha2 = "KR"
	LocalCodeNorthKorea    ISO3166V1Alpha2 = "KP" // 下方Code由AI生成，所以存在注释
	LocalCodeSouthAfrica   ISO3166V1Alpha2 = "ZA" // 南非
	LocalCodeNewZealand    ISO3166V1Alpha2 = "NZ" // 新西兰
	LocalCodeSweden        ISO3166V1Alpha2 = "SE" // 瑞典
	LocalCodeNorway        ISO3166V1Alpha2 = "NO" // 挪威
	LocalCodeDenmark       ISO3166V1Alpha2 = "DK" // 丹麦
	LocalCodeFinland       ISO3166V1Alpha2 = "FI" // 芬兰
	LocalCodeSwitzerland   ISO3166V1Alpha2 = "CH" // 瑞士
	LocalCodeAustria       ISO3166V1Alpha2 = "AT" // 奥地利
	LocalCodeBelgium       ISO3166V1Alpha2 = "BE" // 比利时
	LocalCodePortugal      ISO3166V1Alpha2 = "PT" // 葡萄牙
	LocalCodeGreece        ISO3166V1Alpha2 = "GR" // 希腊
	LocalCodeTurkey        ISO3166V1Alpha2 = "TR" // 土耳其
	LocalCodeEgypt         ISO3166V1Alpha2 = "EG" // 埃及
	LocalCodePakistan      ISO3166V1Alpha2 = "PK" // 巴基斯坦
	LocalCodeBangladesh    ISO3166V1Alpha2 = "BD" // 孟加拉国
	LocalCodeIsrael        ISO3166V1Alpha2 = "IL" // 以色列
	LocalCodeUkraine       ISO3166V1Alpha2 = "UA" // 乌克兰
	LocalCodePoland        ISO3166V1Alpha2 = "PL" // 波兰
	LocalCodeCzechia       ISO3166V1Alpha2 = "CZ" // 捷克
	LocalCodeHungary       ISO3166V1Alpha2 = "HU" // 匈牙利
	LocalCodeUAE           ISO3166V1Alpha2 = "AE" // 阿联酋
	LocalCodeQatar         ISO3166V1Alpha2 = "QA" // 卡塔尔
	LocalCodeKuwait        ISO3166V1Alpha2 = "KW" // 科威特
	LocalCodeIran          ISO3166V1Alpha2 = "IR" // 伊朗
	LocalCodeIraq          ISO3166V1Alpha2 = "IQ" // 伊拉克
	LocalCodeMorocco       ISO3166V1Alpha2 = "MA" // 摩洛哥
	LocalCodeNigeria       ISO3166V1Alpha2 = "NG" // 尼日利亚
	LocalCodeKenya         ISO3166V1Alpha2 = "KE" // 肯尼亚
	LocalCodeParaguay      ISO3166V1Alpha2 = "PY" // 巴拉圭
	LocalCodeUruguay       ISO3166V1Alpha2 = "UY" // 乌拉圭
	LocalCodeEcuador       ISO3166V1Alpha2 = "EC" // 厄瓜多尔
	LocalCodeBolivia       ISO3166V1Alpha2 = "BO" // 玻利维亚
)
