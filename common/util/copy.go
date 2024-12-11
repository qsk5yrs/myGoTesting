package util

import (
	"errors"
	"github.com/jinzhu/copier"
	"github.com/qsk5yrs/testing/common/enum"
	"regexp"
	"time"
)

// CopyProperties 将属性从src拷贝到dest,并完成time.Time和string类型的相互转换
// 参数传 指针类型
func CopyProperties(dst, src interface{}) error {
	err := copier.CopyWithOption(dst, src, copier.Option{
		IgnoreEmpty: true,
		DeepCopy:    true,
		Converters: []copier.TypeConverter{
			{
				// time.Time转换为字符换
				SrcType: time.Time{},
				DstType: copier.String,
				Fn: func(src interface{}) (dst interface{}, err error) {
					s, ok := src.(time.Time)
					if !ok {
						return nil, errors.New("src type is not time.Time")
					}
					return s.Format(enum.TimeFormatHyphenedYMDHIS), nil
				},
			},
			{
				// 字符换转换成time.Time
				SrcType: copier.String,
				DstType: time.Time{},
				Fn: func(src interface{}) (dst interface{}, err error) {
					s, ok := src.(string)
					if !ok {
						return nil, errors.New("src type is not time format string")
					}
					pattern := `^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}$` // YYYY-MM-DD HH:MM:SS
					matched, _ := regexp.MatchString(pattern, s)
					if matched {
						return time.Parse(enum.TimeFormatHyphenedYMDHIS, s)
					}
					return nil, errors.New("src type is not time format string")
				},
			},
		},
	})
	return err
}
