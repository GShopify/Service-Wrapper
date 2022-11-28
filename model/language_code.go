package model

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

type LanguageCode string

const (
	LanguageCodeAf   LanguageCode = "AF"
	LanguageCodeAk   LanguageCode = "AK"
	LanguageCodeAm   LanguageCode = "AM"
	LanguageCodeAr   LanguageCode = "AR"
	LanguageCodeAs   LanguageCode = "AS"
	LanguageCodeAz   LanguageCode = "AZ"
	LanguageCodeBe   LanguageCode = "BE"
	LanguageCodeBg   LanguageCode = "BG"
	LanguageCodeBm   LanguageCode = "BM"
	LanguageCodeBn   LanguageCode = "BN"
	LanguageCodeBo   LanguageCode = "BO"
	LanguageCodeBr   LanguageCode = "BR"
	LanguageCodeBs   LanguageCode = "BS"
	LanguageCodeCa   LanguageCode = "CA"
	LanguageCodeCe   LanguageCode = "CE"
	LanguageCodeCs   LanguageCode = "CS"
	LanguageCodeCu   LanguageCode = "CU"
	LanguageCodeCy   LanguageCode = "CY"
	LanguageCodeDa   LanguageCode = "DA"
	LanguageCodeDe   LanguageCode = "DE"
	LanguageCodeDz   LanguageCode = "DZ"
	LanguageCodeEe   LanguageCode = "EE"
	LanguageCodeEl   LanguageCode = "EL"
	LanguageCodeEn   LanguageCode = "EN"
	LanguageCodeEo   LanguageCode = "EO"
	LanguageCodeEs   LanguageCode = "ES"
	LanguageCodeEt   LanguageCode = "ET"
	LanguageCodeEu   LanguageCode = "EU"
	LanguageCodeFa   LanguageCode = "FA"
	LanguageCodeFf   LanguageCode = "FF"
	LanguageCodeFi   LanguageCode = "FI"
	LanguageCodeFo   LanguageCode = "FO"
	LanguageCodeFr   LanguageCode = "FR"
	LanguageCodeFy   LanguageCode = "FY"
	LanguageCodeGa   LanguageCode = "GA"
	LanguageCodeGd   LanguageCode = "GD"
	LanguageCodeGl   LanguageCode = "GL"
	LanguageCodeGu   LanguageCode = "GU"
	LanguageCodeGv   LanguageCode = "GV"
	LanguageCodeHa   LanguageCode = "HA"
	LanguageCodeHe   LanguageCode = "HE"
	LanguageCodeHi   LanguageCode = "HI"
	LanguageCodeHr   LanguageCode = "HR"
	LanguageCodeHu   LanguageCode = "HU"
	LanguageCodeHy   LanguageCode = "HY"
	LanguageCodeIa   LanguageCode = "IA"
	LanguageCodeID   LanguageCode = "ID"
	LanguageCodeIg   LanguageCode = "IG"
	LanguageCodeIi   LanguageCode = "II"
	LanguageCodeIs   LanguageCode = "IS"
	LanguageCodeIt   LanguageCode = "IT"
	LanguageCodeJa   LanguageCode = "JA"
	LanguageCodeJv   LanguageCode = "JV"
	LanguageCodeKa   LanguageCode = "KA"
	LanguageCodeKi   LanguageCode = "KI"
	LanguageCodeKk   LanguageCode = "KK"
	LanguageCodeKl   LanguageCode = "KL"
	LanguageCodeKm   LanguageCode = "KM"
	LanguageCodeKn   LanguageCode = "KN"
	LanguageCodeKo   LanguageCode = "KO"
	LanguageCodeKs   LanguageCode = "KS"
	LanguageCodeKu   LanguageCode = "KU"
	LanguageCodeKw   LanguageCode = "KW"
	LanguageCodeKy   LanguageCode = "KY"
	LanguageCodeLb   LanguageCode = "LB"
	LanguageCodeLg   LanguageCode = "LG"
	LanguageCodeLn   LanguageCode = "LN"
	LanguageCodeLo   LanguageCode = "LO"
	LanguageCodeLt   LanguageCode = "LT"
	LanguageCodeLu   LanguageCode = "LU"
	LanguageCodeLv   LanguageCode = "LV"
	LanguageCodeMg   LanguageCode = "MG"
	LanguageCodeMi   LanguageCode = "MI"
	LanguageCodeMk   LanguageCode = "MK"
	LanguageCodeMl   LanguageCode = "ML"
	LanguageCodeMn   LanguageCode = "MN"
	LanguageCodeMr   LanguageCode = "MR"
	LanguageCodeMs   LanguageCode = "MS"
	LanguageCodeMt   LanguageCode = "MT"
	LanguageCodeMy   LanguageCode = "MY"
	LanguageCodeNb   LanguageCode = "NB"
	LanguageCodeNd   LanguageCode = "ND"
	LanguageCodeNe   LanguageCode = "NE"
	LanguageCodeNl   LanguageCode = "NL"
	LanguageCodeNn   LanguageCode = "NN"
	LanguageCodeNo   LanguageCode = "NO"
	LanguageCodeOm   LanguageCode = "OM"
	LanguageCodeOr   LanguageCode = "OR"
	LanguageCodeOs   LanguageCode = "OS"
	LanguageCodePa   LanguageCode = "PA"
	LanguageCodePl   LanguageCode = "PL"
	LanguageCodePs   LanguageCode = "PS"
	LanguageCodePtBr LanguageCode = "PT_BR"
	LanguageCodePtPt LanguageCode = "PT_PT"
	LanguageCodeQu   LanguageCode = "QU"
	LanguageCodeRm   LanguageCode = "RM"
	LanguageCodeRn   LanguageCode = "RN"
	LanguageCodeRo   LanguageCode = "RO"
	LanguageCodeRu   LanguageCode = "RU"
	LanguageCodeRw   LanguageCode = "RW"
	LanguageCodeSd   LanguageCode = "SD"
	LanguageCodeSe   LanguageCode = "SE"
	LanguageCodeSg   LanguageCode = "SG"
	LanguageCodeSi   LanguageCode = "SI"
	LanguageCodeSk   LanguageCode = "SK"
	LanguageCodeSl   LanguageCode = "SL"
	LanguageCodeSn   LanguageCode = "SN"
	LanguageCodeSo   LanguageCode = "SO"
	LanguageCodeSq   LanguageCode = "SQ"
	LanguageCodeSr   LanguageCode = "SR"
	LanguageCodeSu   LanguageCode = "SU"
	LanguageCodeSv   LanguageCode = "SV"
	LanguageCodeSw   LanguageCode = "SW"
	LanguageCodeTa   LanguageCode = "TA"
	LanguageCodeTe   LanguageCode = "TE"
	LanguageCodeTg   LanguageCode = "TG"
	LanguageCodeTh   LanguageCode = "TH"
	LanguageCodeTi   LanguageCode = "TI"
	LanguageCodeTk   LanguageCode = "TK"
	LanguageCodeTo   LanguageCode = "TO"
	LanguageCodeTr   LanguageCode = "TR"
	LanguageCodeTt   LanguageCode = "TT"
	LanguageCodeUg   LanguageCode = "UG"
	LanguageCodeUk   LanguageCode = "UK"
	LanguageCodeUr   LanguageCode = "UR"
	LanguageCodeUz   LanguageCode = "UZ"
	LanguageCodeVi   LanguageCode = "VI"
	LanguageCodeWo   LanguageCode = "WO"
	LanguageCodeXh   LanguageCode = "XH"
	LanguageCodeYi   LanguageCode = "YI"
	LanguageCodeYo   LanguageCode = "YO"
	LanguageCodeZhCn LanguageCode = "ZH_CN"
	LanguageCodeZhTw LanguageCode = "ZH_TW"
	LanguageCodeZu   LanguageCode = "ZU"
	LanguageCodeZh   LanguageCode = "ZH"
	LanguageCodePt   LanguageCode = "PT"
	LanguageCodeVo   LanguageCode = "VO"
)

var AllLanguageCode = []LanguageCode{
	LanguageCodeAf,
	LanguageCodeAk,
	LanguageCodeAm,
	LanguageCodeAr,
	LanguageCodeAs,
	LanguageCodeAz,
	LanguageCodeBe,
	LanguageCodeBg,
	LanguageCodeBm,
	LanguageCodeBn,
	LanguageCodeBo,
	LanguageCodeBr,
	LanguageCodeBs,
	LanguageCodeCa,
	LanguageCodeCe,
	LanguageCodeCs,
	LanguageCodeCu,
	LanguageCodeCy,
	LanguageCodeDa,
	LanguageCodeDe,
	LanguageCodeDz,
	LanguageCodeEe,
	LanguageCodeEl,
	LanguageCodeEn,
	LanguageCodeEo,
	LanguageCodeEs,
	LanguageCodeEt,
	LanguageCodeEu,
	LanguageCodeFa,
	LanguageCodeFf,
	LanguageCodeFi,
	LanguageCodeFo,
	LanguageCodeFr,
	LanguageCodeFy,
	LanguageCodeGa,
	LanguageCodeGd,
	LanguageCodeGl,
	LanguageCodeGu,
	LanguageCodeGv,
	LanguageCodeHa,
	LanguageCodeHe,
	LanguageCodeHi,
	LanguageCodeHr,
	LanguageCodeHu,
	LanguageCodeHy,
	LanguageCodeIa,
	LanguageCodeID,
	LanguageCodeIg,
	LanguageCodeIi,
	LanguageCodeIs,
	LanguageCodeIt,
	LanguageCodeJa,
	LanguageCodeJv,
	LanguageCodeKa,
	LanguageCodeKi,
	LanguageCodeKk,
	LanguageCodeKl,
	LanguageCodeKm,
	LanguageCodeKn,
	LanguageCodeKo,
	LanguageCodeKs,
	LanguageCodeKu,
	LanguageCodeKw,
	LanguageCodeKy,
	LanguageCodeLb,
	LanguageCodeLg,
	LanguageCodeLn,
	LanguageCodeLo,
	LanguageCodeLt,
	LanguageCodeLu,
	LanguageCodeLv,
	LanguageCodeMg,
	LanguageCodeMi,
	LanguageCodeMk,
	LanguageCodeMl,
	LanguageCodeMn,
	LanguageCodeMr,
	LanguageCodeMs,
	LanguageCodeMt,
	LanguageCodeMy,
	LanguageCodeNb,
	LanguageCodeNd,
	LanguageCodeNe,
	LanguageCodeNl,
	LanguageCodeNn,
	LanguageCodeNo,
	LanguageCodeOm,
	LanguageCodeOr,
	LanguageCodeOs,
	LanguageCodePa,
	LanguageCodePl,
	LanguageCodePs,
	LanguageCodePtBr,
	LanguageCodePtPt,
	LanguageCodeQu,
	LanguageCodeRm,
	LanguageCodeRn,
	LanguageCodeRo,
	LanguageCodeRu,
	LanguageCodeRw,
	LanguageCodeSd,
	LanguageCodeSe,
	LanguageCodeSg,
	LanguageCodeSi,
	LanguageCodeSk,
	LanguageCodeSl,
	LanguageCodeSn,
	LanguageCodeSo,
	LanguageCodeSq,
	LanguageCodeSr,
	LanguageCodeSu,
	LanguageCodeSv,
	LanguageCodeSw,
	LanguageCodeTa,
	LanguageCodeTe,
	LanguageCodeTg,
	LanguageCodeTh,
	LanguageCodeTi,
	LanguageCodeTk,
	LanguageCodeTo,
	LanguageCodeTr,
	LanguageCodeTt,
	LanguageCodeUg,
	LanguageCodeUk,
	LanguageCodeUr,
	LanguageCodeUz,
	LanguageCodeVi,
	LanguageCodeWo,
	LanguageCodeXh,
	LanguageCodeYi,
	LanguageCodeYo,
	LanguageCodeZhCn,
	LanguageCodeZhTw,
	LanguageCodeZu,
	LanguageCodeZh,
	LanguageCodePt,
	LanguageCodeVo,
}

func (e LanguageCode) IsValid() bool {
	switch e {
	case LanguageCodeAf, LanguageCodeAk, LanguageCodeAm, LanguageCodeAr, LanguageCodeAs, LanguageCodeAz, LanguageCodeBe, LanguageCodeBg, LanguageCodeBm, LanguageCodeBn, LanguageCodeBo, LanguageCodeBr, LanguageCodeBs, LanguageCodeCa, LanguageCodeCe, LanguageCodeCs, LanguageCodeCu, LanguageCodeCy, LanguageCodeDa, LanguageCodeDe, LanguageCodeDz, LanguageCodeEe, LanguageCodeEl, LanguageCodeEn, LanguageCodeEo, LanguageCodeEs, LanguageCodeEt, LanguageCodeEu, LanguageCodeFa, LanguageCodeFf, LanguageCodeFi, LanguageCodeFo, LanguageCodeFr, LanguageCodeFy, LanguageCodeGa, LanguageCodeGd, LanguageCodeGl, LanguageCodeGu, LanguageCodeGv, LanguageCodeHa, LanguageCodeHe, LanguageCodeHi, LanguageCodeHr, LanguageCodeHu, LanguageCodeHy, LanguageCodeIa, LanguageCodeID, LanguageCodeIg, LanguageCodeIi, LanguageCodeIs, LanguageCodeIt, LanguageCodeJa, LanguageCodeJv, LanguageCodeKa, LanguageCodeKi, LanguageCodeKk, LanguageCodeKl, LanguageCodeKm, LanguageCodeKn, LanguageCodeKo, LanguageCodeKs, LanguageCodeKu, LanguageCodeKw, LanguageCodeKy, LanguageCodeLb, LanguageCodeLg, LanguageCodeLn, LanguageCodeLo, LanguageCodeLt, LanguageCodeLu, LanguageCodeLv, LanguageCodeMg, LanguageCodeMi, LanguageCodeMk, LanguageCodeMl, LanguageCodeMn, LanguageCodeMr, LanguageCodeMs, LanguageCodeMt, LanguageCodeMy, LanguageCodeNb, LanguageCodeNd, LanguageCodeNe, LanguageCodeNl, LanguageCodeNn, LanguageCodeNo, LanguageCodeOm, LanguageCodeOr, LanguageCodeOs, LanguageCodePa, LanguageCodePl, LanguageCodePs, LanguageCodePtBr, LanguageCodePtPt, LanguageCodeQu, LanguageCodeRm, LanguageCodeRn, LanguageCodeRo, LanguageCodeRu, LanguageCodeRw, LanguageCodeSd, LanguageCodeSe, LanguageCodeSg, LanguageCodeSi, LanguageCodeSk, LanguageCodeSl, LanguageCodeSn, LanguageCodeSo, LanguageCodeSq, LanguageCodeSr, LanguageCodeSu, LanguageCodeSv, LanguageCodeSw, LanguageCodeTa, LanguageCodeTe, LanguageCodeTg, LanguageCodeTh, LanguageCodeTi, LanguageCodeTk, LanguageCodeTo, LanguageCodeTr, LanguageCodeTt, LanguageCodeUg, LanguageCodeUk, LanguageCodeUr, LanguageCodeUz, LanguageCodeVi, LanguageCodeWo, LanguageCodeXh, LanguageCodeYi, LanguageCodeYo, LanguageCodeZhCn, LanguageCodeZhTw, LanguageCodeZu, LanguageCodeZh, LanguageCodePt, LanguageCodeVo:
		return true
	}
	return false
}

func (e LanguageCode) String() string {
	return string(e)
}

func (e LanguageCode) SqlFieldSelection(f, namespace, wrap string) string {
	field := func(f, namespace string) string {
		builder := strings.Builder{}
		builder.WriteString(namespace)

		if builder.Len() > 0 {
			builder.WriteString(".")
		}

		builder.WriteString(f)
		return builder.String()
	}(f, namespace)

	if LanguageCodeEn == e {
		return fmt.Sprintf("%[4]s(%[1]s['%[2]s']) as %[3]s", field, LanguageCodeEn, f, wrap)
	}

	return fmt.Sprintf("if(notEmpty(%[5]s(%[1]s['%[2]s'])), %[5]s(%[1]s['%[2]s']), %[5]s(%[1]s['%[3]s'])) as %[4]s", field, e, LanguageCodeEn, f, wrap)
}

func (e LanguageCode) SqlArraySelection(f string) string {
	if LanguageCodeEn == e {
		return fmt.Sprintf("arrayMap(m -> (m['%[2]s']), %[1]s) as %[1]s", f, LanguageCodeEn)
	}

	return fmt.Sprintf("arrayMap(m -> (if(notEmpty(m['%[2]s']), m['%[2]s'], m['%[3]s'])), %[1]s) as %[1]s", f, e, LanguageCodeEn)
}

func (e *LanguageCode) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = LanguageCode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid LanguageCode", str)
	}
	return nil
}

func (e LanguageCode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
