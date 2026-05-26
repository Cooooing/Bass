package enum

import (
	v1 "common/api/gen/user/v1"
	"common/pkg/enum"
)

type MBTI string

const (
	MBTIIntj MBTI = "intj"
	MBTIIntp MBTI = "intp"
	MBTIEntj MBTI = "entj"
	MBTIEntp MBTI = "entp"
	MBTIInfj MBTI = "infj"
	MBTIInfp MBTI = "infp"
	MBTIEnfj MBTI = "enfj"
	MBTIEnfp MBTI = "enfp"
	MBTIIstj MBTI = "istj"
	MBTIIsfj MBTI = "isfj"
	MBTIEstj MBTI = "estj"
	MBTIEsfj MBTI = "esfj"
	MBTIIstp MBTI = "istp"
	MBTIIsfp MBTI = "isfp"
	MBTIEstp MBTI = "estp"
	MBTIEsfp MBTI = "esfp"
)

var MBTIMap = enum.NewMapping[MBTI, v1.MBTI](map[MBTI]enum.Entry[MBTI, v1.MBTI]{
	MBTIIntj: {Proto: v1.MBTI_MBTI_INTJ},
	MBTIIntp: {Proto: v1.MBTI_MBTI_INTP},
	MBTIEntj: {Proto: v1.MBTI_MBTI_ENTJ},
	MBTIEntp: {Proto: v1.MBTI_MBTI_ENTP},
	MBTIInfj: {Proto: v1.MBTI_MBTI_INFJ},
	MBTIInfp: {Proto: v1.MBTI_MBTI_INFP},
	MBTIEnfj: {Proto: v1.MBTI_MBTI_ENFJ},
	MBTIEnfp: {Proto: v1.MBTI_MBTI_ENFP},
	MBTIIstj: {Proto: v1.MBTI_MBTI_ISTJ},
	MBTIIsfj: {Proto: v1.MBTI_MBTI_ISFJ},
	MBTIEstj: {Proto: v1.MBTI_MBTI_ESTJ},
	MBTIEsfj: {Proto: v1.MBTI_MBTI_ESFJ},
	MBTIIstp: {Proto: v1.MBTI_MBTI_ISTP},
	MBTIIsfp: {Proto: v1.MBTI_MBTI_ISFP},
	MBTIEstp: {Proto: v1.MBTI_MBTI_ESTP},
	MBTIEsfp: {Proto: v1.MBTI_MBTI_ESFP},
})
