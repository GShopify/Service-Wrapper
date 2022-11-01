package model

import (
	"fmt"
	"io"
	"strconv"
)

type CountryCode string

const (
	CountryCodeAf CountryCode = "AF"
	CountryCodeAx CountryCode = "AX"
	CountryCodeAl CountryCode = "AL"
	CountryCodeDz CountryCode = "DZ"
	CountryCodeAd CountryCode = "AD"
	CountryCodeAo CountryCode = "AO"
	CountryCodeAi CountryCode = "AI"
	CountryCodeAg CountryCode = "AG"
	CountryCodeAr CountryCode = "AR"
	CountryCodeAm CountryCode = "AM"
	CountryCodeAw CountryCode = "AW"
	CountryCodeAc CountryCode = "AC"
	CountryCodeAu CountryCode = "AU"
	CountryCodeAt CountryCode = "AT"
	CountryCodeAz CountryCode = "AZ"
	CountryCodeBs CountryCode = "BS"
	CountryCodeBh CountryCode = "BH"
	CountryCodeBd CountryCode = "BD"
	CountryCodeBb CountryCode = "BB"
	CountryCodeBy CountryCode = "BY"
	CountryCodeBe CountryCode = "BE"
	CountryCodeBz CountryCode = "BZ"
	CountryCodeBj CountryCode = "BJ"
	CountryCodeBm CountryCode = "BM"
	CountryCodeBt CountryCode = "BT"
	CountryCodeBo CountryCode = "BO"
	CountryCodeBa CountryCode = "BA"
	CountryCodeBw CountryCode = "BW"
	CountryCodeBv CountryCode = "BV"
	CountryCodeBr CountryCode = "BR"
	CountryCodeIo CountryCode = "IO"
	CountryCodeBn CountryCode = "BN"
	CountryCodeBg CountryCode = "BG"
	CountryCodeBf CountryCode = "BF"
	CountryCodeBi CountryCode = "BI"
	CountryCodeKh CountryCode = "KH"
	CountryCodeCa CountryCode = "CA"
	CountryCodeCv CountryCode = "CV"
	CountryCodeBq CountryCode = "BQ"
	CountryCodeKy CountryCode = "KY"
	CountryCodeCf CountryCode = "CF"
	CountryCodeTd CountryCode = "TD"
	CountryCodeCl CountryCode = "CL"
	CountryCodeCn CountryCode = "CN"
	CountryCodeCx CountryCode = "CX"
	CountryCodeCc CountryCode = "CC"
	CountryCodeCo CountryCode = "CO"
	CountryCodeKm CountryCode = "KM"
	CountryCodeCg CountryCode = "CG"
	CountryCodeCd CountryCode = "CD"
	CountryCodeCk CountryCode = "CK"
	CountryCodeCr CountryCode = "CR"
	CountryCodeHr CountryCode = "HR"
	CountryCodeCu CountryCode = "CU"
	CountryCodeCw CountryCode = "CW"
	CountryCodeCy CountryCode = "CY"
	CountryCodeCz CountryCode = "CZ"
	CountryCodeCi CountryCode = "CI"
	CountryCodeDk CountryCode = "DK"
	CountryCodeDj CountryCode = "DJ"
	CountryCodeDm CountryCode = "DM"
	CountryCodeDo CountryCode = "DO"
	CountryCodeEc CountryCode = "EC"
	CountryCodeEg CountryCode = "EG"
	CountryCodeSv CountryCode = "SV"
	CountryCodeGq CountryCode = "GQ"
	CountryCodeEr CountryCode = "ER"
	CountryCodeEe CountryCode = "EE"
	CountryCodeSz CountryCode = "SZ"
	CountryCodeEt CountryCode = "ET"
	CountryCodeFk CountryCode = "FK"
	CountryCodeFo CountryCode = "FO"
	CountryCodeFj CountryCode = "FJ"
	CountryCodeFi CountryCode = "FI"
	CountryCodeFr CountryCode = "FR"
	CountryCodeGf CountryCode = "GF"
	CountryCodePf CountryCode = "PF"
	CountryCodeTf CountryCode = "TF"
	CountryCodeGa CountryCode = "GA"
	CountryCodeGm CountryCode = "GM"
	CountryCodeGe CountryCode = "GE"
	CountryCodeDe CountryCode = "DE"
	CountryCodeGh CountryCode = "GH"
	CountryCodeGi CountryCode = "GI"
	CountryCodeGr CountryCode = "GR"
	CountryCodeGl CountryCode = "GL"
	CountryCodeGd CountryCode = "GD"
	CountryCodeGp CountryCode = "GP"
	CountryCodeGt CountryCode = "GT"
	CountryCodeGg CountryCode = "GG"
	CountryCodeGn CountryCode = "GN"
	CountryCodeGw CountryCode = "GW"
	CountryCodeGy CountryCode = "GY"
	CountryCodeHt CountryCode = "HT"
	CountryCodeHm CountryCode = "HM"
	CountryCodeVa CountryCode = "VA"
	CountryCodeHn CountryCode = "HN"
	CountryCodeHk CountryCode = "HK"
	CountryCodeHu CountryCode = "HU"
	CountryCodeIs CountryCode = "IS"
	CountryCodeIn CountryCode = "IN"
	CountryCodeID CountryCode = "ID"
	CountryCodeIr CountryCode = "IR"
	CountryCodeIq CountryCode = "IQ"
	CountryCodeIe CountryCode = "IE"
	CountryCodeIm CountryCode = "IM"
	CountryCodeIl CountryCode = "IL"
	CountryCodeIt CountryCode = "IT"
	CountryCodeJm CountryCode = "JM"
	CountryCodeJp CountryCode = "JP"
	CountryCodeJe CountryCode = "JE"
	CountryCodeJo CountryCode = "JO"
	CountryCodeKz CountryCode = "KZ"
	CountryCodeKe CountryCode = "KE"
	CountryCodeKi CountryCode = "KI"
	CountryCodeKp CountryCode = "KP"
	CountryCodeXk CountryCode = "XK"
	CountryCodeKw CountryCode = "KW"
	CountryCodeKg CountryCode = "KG"
	CountryCodeLa CountryCode = "LA"
	CountryCodeLv CountryCode = "LV"
	CountryCodeLb CountryCode = "LB"
	CountryCodeLs CountryCode = "LS"
	CountryCodeLr CountryCode = "LR"
	CountryCodeLy CountryCode = "LY"
	CountryCodeLi CountryCode = "LI"
	CountryCodeLt CountryCode = "LT"
	CountryCodeLu CountryCode = "LU"
	CountryCodeMo CountryCode = "MO"
	CountryCodeMg CountryCode = "MG"
	CountryCodeMw CountryCode = "MW"
	CountryCodeMy CountryCode = "MY"
	CountryCodeMv CountryCode = "MV"
	CountryCodeMl CountryCode = "ML"
	CountryCodeMt CountryCode = "MT"
	CountryCodeMq CountryCode = "MQ"
	CountryCodeMr CountryCode = "MR"
	CountryCodeMu CountryCode = "MU"
	CountryCodeYt CountryCode = "YT"
	CountryCodeMx CountryCode = "MX"
	CountryCodeMd CountryCode = "MD"
	CountryCodeMc CountryCode = "MC"
	CountryCodeMn CountryCode = "MN"
	CountryCodeMe CountryCode = "ME"
	CountryCodeMs CountryCode = "MS"
	CountryCodeMa CountryCode = "MA"
	CountryCodeMz CountryCode = "MZ"
	CountryCodeMm CountryCode = "MM"
	CountryCodeNa CountryCode = "NA"
	CountryCodeNr CountryCode = "NR"
	CountryCodeNp CountryCode = "NP"
	CountryCodeNl CountryCode = "NL"
	CountryCodeAn CountryCode = "AN"
	CountryCodeNc CountryCode = "NC"
	CountryCodeNz CountryCode = "NZ"
	CountryCodeNi CountryCode = "NI"
	CountryCodeNe CountryCode = "NE"
	CountryCodeNg CountryCode = "NG"
	CountryCodeNu CountryCode = "NU"
	CountryCodeNf CountryCode = "NF"
	CountryCodeMk CountryCode = "MK"
	CountryCodeNo CountryCode = "NO"
	CountryCodeOm CountryCode = "OM"
	CountryCodePk CountryCode = "PK"
	CountryCodePs CountryCode = "PS"
	CountryCodePa CountryCode = "PA"
	CountryCodePg CountryCode = "PG"
	CountryCodePy CountryCode = "PY"
	CountryCodePe CountryCode = "PE"
	CountryCodePh CountryCode = "PH"
	CountryCodePn CountryCode = "PN"
	CountryCodePl CountryCode = "PL"
	CountryCodePt CountryCode = "PT"
	CountryCodeQa CountryCode = "QA"
	CountryCodeCm CountryCode = "CM"
	CountryCodeRe CountryCode = "RE"
	CountryCodeRo CountryCode = "RO"
	CountryCodeRu CountryCode = "RU"
	CountryCodeRw CountryCode = "RW"
	CountryCodeBl CountryCode = "BL"
	CountryCodeSh CountryCode = "SH"
	CountryCodeKn CountryCode = "KN"
	CountryCodeLc CountryCode = "LC"
	CountryCodeMf CountryCode = "MF"
	CountryCodePm CountryCode = "PM"
	CountryCodeWs CountryCode = "WS"
	CountryCodeSm CountryCode = "SM"
	CountryCodeSt CountryCode = "ST"
	CountryCodeSa CountryCode = "SA"
	CountryCodeSn CountryCode = "SN"
	CountryCodeRs CountryCode = "RS"
	CountryCodeSc CountryCode = "SC"
	CountryCodeSl CountryCode = "SL"
	CountryCodeSg CountryCode = "SG"
	CountryCodeSx CountryCode = "SX"
	CountryCodeSk CountryCode = "SK"
	CountryCodeSi CountryCode = "SI"
	CountryCodeSb CountryCode = "SB"
	CountryCodeSo CountryCode = "SO"
	CountryCodeZa CountryCode = "ZA"
	CountryCodeGs CountryCode = "GS"
	CountryCodeKr CountryCode = "KR"
	CountryCodeSs CountryCode = "SS"
	CountryCodeEs CountryCode = "ES"
	CountryCodeLk CountryCode = "LK"
	CountryCodeVc CountryCode = "VC"
	CountryCodeSd CountryCode = "SD"
	CountryCodeSr CountryCode = "SR"
	CountryCodeSj CountryCode = "SJ"
	CountryCodeSe CountryCode = "SE"
	CountryCodeCh CountryCode = "CH"
	CountryCodeSy CountryCode = "SY"
	CountryCodeTw CountryCode = "TW"
	CountryCodeTj CountryCode = "TJ"
	CountryCodeTz CountryCode = "TZ"
	CountryCodeTh CountryCode = "TH"
	CountryCodeTl CountryCode = "TL"
	CountryCodeTg CountryCode = "TG"
	CountryCodeTk CountryCode = "TK"
	CountryCodeTo CountryCode = "TO"
	CountryCodeTt CountryCode = "TT"
	CountryCodeTa CountryCode = "TA"
	CountryCodeTn CountryCode = "TN"
	CountryCodeTr CountryCode = "TR"
	CountryCodeTm CountryCode = "TM"
	CountryCodeTc CountryCode = "TC"
	CountryCodeTv CountryCode = "TV"
	CountryCodeUg CountryCode = "UG"
	CountryCodeUa CountryCode = "UA"
	CountryCodeAe CountryCode = "AE"
	CountryCodeGb CountryCode = "GB"
	CountryCodeUs CountryCode = "US"
	CountryCodeUm CountryCode = "UM"
	CountryCodeUy CountryCode = "UY"
	CountryCodeUz CountryCode = "UZ"
	CountryCodeVu CountryCode = "VU"
	CountryCodeVe CountryCode = "VE"
	CountryCodeVn CountryCode = "VN"
	CountryCodeVg CountryCode = "VG"
	CountryCodeWf CountryCode = "WF"
	CountryCodeEh CountryCode = "EH"
	CountryCodeYe CountryCode = "YE"
	CountryCodeZm CountryCode = "ZM"
	CountryCodeZw CountryCode = "ZW"
	CountryCodeZz CountryCode = "ZZ"
)

var AllCountryCode = []CountryCode{
	CountryCodeAf,
	CountryCodeAx,
	CountryCodeAl,
	CountryCodeDz,
	CountryCodeAd,
	CountryCodeAo,
	CountryCodeAi,
	CountryCodeAg,
	CountryCodeAr,
	CountryCodeAm,
	CountryCodeAw,
	CountryCodeAc,
	CountryCodeAu,
	CountryCodeAt,
	CountryCodeAz,
	CountryCodeBs,
	CountryCodeBh,
	CountryCodeBd,
	CountryCodeBb,
	CountryCodeBy,
	CountryCodeBe,
	CountryCodeBz,
	CountryCodeBj,
	CountryCodeBm,
	CountryCodeBt,
	CountryCodeBo,
	CountryCodeBa,
	CountryCodeBw,
	CountryCodeBv,
	CountryCodeBr,
	CountryCodeIo,
	CountryCodeBn,
	CountryCodeBg,
	CountryCodeBf,
	CountryCodeBi,
	CountryCodeKh,
	CountryCodeCa,
	CountryCodeCv,
	CountryCodeBq,
	CountryCodeKy,
	CountryCodeCf,
	CountryCodeTd,
	CountryCodeCl,
	CountryCodeCn,
	CountryCodeCx,
	CountryCodeCc,
	CountryCodeCo,
	CountryCodeKm,
	CountryCodeCg,
	CountryCodeCd,
	CountryCodeCk,
	CountryCodeCr,
	CountryCodeHr,
	CountryCodeCu,
	CountryCodeCw,
	CountryCodeCy,
	CountryCodeCz,
	CountryCodeCi,
	CountryCodeDk,
	CountryCodeDj,
	CountryCodeDm,
	CountryCodeDo,
	CountryCodeEc,
	CountryCodeEg,
	CountryCodeSv,
	CountryCodeGq,
	CountryCodeEr,
	CountryCodeEe,
	CountryCodeSz,
	CountryCodeEt,
	CountryCodeFk,
	CountryCodeFo,
	CountryCodeFj,
	CountryCodeFi,
	CountryCodeFr,
	CountryCodeGf,
	CountryCodePf,
	CountryCodeTf,
	CountryCodeGa,
	CountryCodeGm,
	CountryCodeGe,
	CountryCodeDe,
	CountryCodeGh,
	CountryCodeGi,
	CountryCodeGr,
	CountryCodeGl,
	CountryCodeGd,
	CountryCodeGp,
	CountryCodeGt,
	CountryCodeGg,
	CountryCodeGn,
	CountryCodeGw,
	CountryCodeGy,
	CountryCodeHt,
	CountryCodeHm,
	CountryCodeVa,
	CountryCodeHn,
	CountryCodeHk,
	CountryCodeHu,
	CountryCodeIs,
	CountryCodeIn,
	CountryCodeID,
	CountryCodeIr,
	CountryCodeIq,
	CountryCodeIe,
	CountryCodeIm,
	CountryCodeIl,
	CountryCodeIt,
	CountryCodeJm,
	CountryCodeJp,
	CountryCodeJe,
	CountryCodeJo,
	CountryCodeKz,
	CountryCodeKe,
	CountryCodeKi,
	CountryCodeKp,
	CountryCodeXk,
	CountryCodeKw,
	CountryCodeKg,
	CountryCodeLa,
	CountryCodeLv,
	CountryCodeLb,
	CountryCodeLs,
	CountryCodeLr,
	CountryCodeLy,
	CountryCodeLi,
	CountryCodeLt,
	CountryCodeLu,
	CountryCodeMo,
	CountryCodeMg,
	CountryCodeMw,
	CountryCodeMy,
	CountryCodeMv,
	CountryCodeMl,
	CountryCodeMt,
	CountryCodeMq,
	CountryCodeMr,
	CountryCodeMu,
	CountryCodeYt,
	CountryCodeMx,
	CountryCodeMd,
	CountryCodeMc,
	CountryCodeMn,
	CountryCodeMe,
	CountryCodeMs,
	CountryCodeMa,
	CountryCodeMz,
	CountryCodeMm,
	CountryCodeNa,
	CountryCodeNr,
	CountryCodeNp,
	CountryCodeNl,
	CountryCodeAn,
	CountryCodeNc,
	CountryCodeNz,
	CountryCodeNi,
	CountryCodeNe,
	CountryCodeNg,
	CountryCodeNu,
	CountryCodeNf,
	CountryCodeMk,
	CountryCodeNo,
	CountryCodeOm,
	CountryCodePk,
	CountryCodePs,
	CountryCodePa,
	CountryCodePg,
	CountryCodePy,
	CountryCodePe,
	CountryCodePh,
	CountryCodePn,
	CountryCodePl,
	CountryCodePt,
	CountryCodeQa,
	CountryCodeCm,
	CountryCodeRe,
	CountryCodeRo,
	CountryCodeRu,
	CountryCodeRw,
	CountryCodeBl,
	CountryCodeSh,
	CountryCodeKn,
	CountryCodeLc,
	CountryCodeMf,
	CountryCodePm,
	CountryCodeWs,
	CountryCodeSm,
	CountryCodeSt,
	CountryCodeSa,
	CountryCodeSn,
	CountryCodeRs,
	CountryCodeSc,
	CountryCodeSl,
	CountryCodeSg,
	CountryCodeSx,
	CountryCodeSk,
	CountryCodeSi,
	CountryCodeSb,
	CountryCodeSo,
	CountryCodeZa,
	CountryCodeGs,
	CountryCodeKr,
	CountryCodeSs,
	CountryCodeEs,
	CountryCodeLk,
	CountryCodeVc,
	CountryCodeSd,
	CountryCodeSr,
	CountryCodeSj,
	CountryCodeSe,
	CountryCodeCh,
	CountryCodeSy,
	CountryCodeTw,
	CountryCodeTj,
	CountryCodeTz,
	CountryCodeTh,
	CountryCodeTl,
	CountryCodeTg,
	CountryCodeTk,
	CountryCodeTo,
	CountryCodeTt,
	CountryCodeTa,
	CountryCodeTn,
	CountryCodeTr,
	CountryCodeTm,
	CountryCodeTc,
	CountryCodeTv,
	CountryCodeUg,
	CountryCodeUa,
	CountryCodeAe,
	CountryCodeGb,
	CountryCodeUs,
	CountryCodeUm,
	CountryCodeUy,
	CountryCodeUz,
	CountryCodeVu,
	CountryCodeVe,
	CountryCodeVn,
	CountryCodeVg,
	CountryCodeWf,
	CountryCodeEh,
	CountryCodeYe,
	CountryCodeZm,
	CountryCodeZw,
	CountryCodeZz,
}

func (e CountryCode) IsValid() bool {
	switch e {
	case CountryCodeAf, CountryCodeAx, CountryCodeAl, CountryCodeDz, CountryCodeAd, CountryCodeAo, CountryCodeAi, CountryCodeAg, CountryCodeAr, CountryCodeAm, CountryCodeAw, CountryCodeAc, CountryCodeAu, CountryCodeAt, CountryCodeAz, CountryCodeBs, CountryCodeBh, CountryCodeBd, CountryCodeBb, CountryCodeBy, CountryCodeBe, CountryCodeBz, CountryCodeBj, CountryCodeBm, CountryCodeBt, CountryCodeBo, CountryCodeBa, CountryCodeBw, CountryCodeBv, CountryCodeBr, CountryCodeIo, CountryCodeBn, CountryCodeBg, CountryCodeBf, CountryCodeBi, CountryCodeKh, CountryCodeCa, CountryCodeCv, CountryCodeBq, CountryCodeKy, CountryCodeCf, CountryCodeTd, CountryCodeCl, CountryCodeCn, CountryCodeCx, CountryCodeCc, CountryCodeCo, CountryCodeKm, CountryCodeCg, CountryCodeCd, CountryCodeCk, CountryCodeCr, CountryCodeHr, CountryCodeCu, CountryCodeCw, CountryCodeCy, CountryCodeCz, CountryCodeCi, CountryCodeDk, CountryCodeDj, CountryCodeDm, CountryCodeDo, CountryCodeEc, CountryCodeEg, CountryCodeSv, CountryCodeGq, CountryCodeEr, CountryCodeEe, CountryCodeSz, CountryCodeEt, CountryCodeFk, CountryCodeFo, CountryCodeFj, CountryCodeFi, CountryCodeFr, CountryCodeGf, CountryCodePf, CountryCodeTf, CountryCodeGa, CountryCodeGm, CountryCodeGe, CountryCodeDe, CountryCodeGh, CountryCodeGi, CountryCodeGr, CountryCodeGl, CountryCodeGd, CountryCodeGp, CountryCodeGt, CountryCodeGg, CountryCodeGn, CountryCodeGw, CountryCodeGy, CountryCodeHt, CountryCodeHm, CountryCodeVa, CountryCodeHn, CountryCodeHk, CountryCodeHu, CountryCodeIs, CountryCodeIn, CountryCodeID, CountryCodeIr, CountryCodeIq, CountryCodeIe, CountryCodeIm, CountryCodeIl, CountryCodeIt, CountryCodeJm, CountryCodeJp, CountryCodeJe, CountryCodeJo, CountryCodeKz, CountryCodeKe, CountryCodeKi, CountryCodeKp, CountryCodeXk, CountryCodeKw, CountryCodeKg, CountryCodeLa, CountryCodeLv, CountryCodeLb, CountryCodeLs, CountryCodeLr, CountryCodeLy, CountryCodeLi, CountryCodeLt, CountryCodeLu, CountryCodeMo, CountryCodeMg, CountryCodeMw, CountryCodeMy, CountryCodeMv, CountryCodeMl, CountryCodeMt, CountryCodeMq, CountryCodeMr, CountryCodeMu, CountryCodeYt, CountryCodeMx, CountryCodeMd, CountryCodeMc, CountryCodeMn, CountryCodeMe, CountryCodeMs, CountryCodeMa, CountryCodeMz, CountryCodeMm, CountryCodeNa, CountryCodeNr, CountryCodeNp, CountryCodeNl, CountryCodeAn, CountryCodeNc, CountryCodeNz, CountryCodeNi, CountryCodeNe, CountryCodeNg, CountryCodeNu, CountryCodeNf, CountryCodeMk, CountryCodeNo, CountryCodeOm, CountryCodePk, CountryCodePs, CountryCodePa, CountryCodePg, CountryCodePy, CountryCodePe, CountryCodePh, CountryCodePn, CountryCodePl, CountryCodePt, CountryCodeQa, CountryCodeCm, CountryCodeRe, CountryCodeRo, CountryCodeRu, CountryCodeRw, CountryCodeBl, CountryCodeSh, CountryCodeKn, CountryCodeLc, CountryCodeMf, CountryCodePm, CountryCodeWs, CountryCodeSm, CountryCodeSt, CountryCodeSa, CountryCodeSn, CountryCodeRs, CountryCodeSc, CountryCodeSl, CountryCodeSg, CountryCodeSx, CountryCodeSk, CountryCodeSi, CountryCodeSb, CountryCodeSo, CountryCodeZa, CountryCodeGs, CountryCodeKr, CountryCodeSs, CountryCodeEs, CountryCodeLk, CountryCodeVc, CountryCodeSd, CountryCodeSr, CountryCodeSj, CountryCodeSe, CountryCodeCh, CountryCodeSy, CountryCodeTw, CountryCodeTj, CountryCodeTz, CountryCodeTh, CountryCodeTl, CountryCodeTg, CountryCodeTk, CountryCodeTo, CountryCodeTt, CountryCodeTa, CountryCodeTn, CountryCodeTr, CountryCodeTm, CountryCodeTc, CountryCodeTv, CountryCodeUg, CountryCodeUa, CountryCodeAe, CountryCodeGb, CountryCodeUs, CountryCodeUm, CountryCodeUy, CountryCodeUz, CountryCodeVu, CountryCodeVe, CountryCodeVn, CountryCodeVg, CountryCodeWf, CountryCodeEh, CountryCodeYe, CountryCodeZm, CountryCodeZw, CountryCodeZz:
		return true
	}
	return false
}

func (e CountryCode) String() string {
	return string(e)
}

func (e *CountryCode) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CountryCode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CountryCode", str)
	}
	return nil
}

func (e CountryCode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
