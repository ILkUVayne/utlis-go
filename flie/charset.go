package flie

import (
	"bufio"
	"github.com/saintfish/chardet"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

var encodingMaps = map[string]encoding.Encoding{
	"UTF-8":        encoding.Nop,
	"GBK":          simplifiedchinese.GBK,
	"GB-18030":     simplifiedchinese.GB18030,
	"HZ-GB-2312":   simplifiedchinese.HZGB2312,
	"Big5":         traditionalchinese.Big5,
	"UTF-16LE":     unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM),
	"UTF-16BE":     unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM),
	"EUC-JP":       japanese.EUCJP,
	"Shift_JIS":    japanese.ShiftJIS,
	"ISO-2022-JP":  japanese.ISO2022JP,
	"ISO-8859-1":   charmap.ISO8859_1,
	"ISO-8859-2":   charmap.ISO8859_2,
	"ISO-8859-3":   charmap.ISO8859_3,
	"ISO-8859-4":   charmap.ISO8859_4,
	"ISO-8859-5":   charmap.ISO8859_5,
	"ISO-8859-6":   charmap.ISO8859_6,
	"ISO-8859-7":   charmap.ISO8859_7,
	"ISO-8859-8":   charmap.ISO8859_8,
	"ISO-8859-9":   charmap.ISO8859_9,
	"ISO-8859-10":  charmap.ISO8859_10,
	"ISO-8859-13":  charmap.ISO8859_13,
	"ISO-8859-14":  charmap.ISO8859_14,
	"ISO-8859-15":  charmap.ISO8859_15,
	"ISO-8859-16":  charmap.ISO8859_16,
	"Windows-1250": charmap.Windows1250,
	"Windows-1251": charmap.Windows1251,
	"Windows-1252": charmap.Windows1252,
	"Windows-1253": charmap.Windows1253,
	"Windows-1254": charmap.Windows1254,
	"Windows-1255": charmap.Windows1255,
	"Windows-1256": charmap.Windows1256,
	"Windows-1257": charmap.Windows1257,
	"Windows-1258": charmap.Windows1258,
	"KOI8-R":       charmap.KOI8R,
	"KOI8-U":       charmap.KOI8U,
}

// DecodeEncoding 将对应格式文本转换成utf-8
func DecodeEncoding(r *bufio.Reader) (*transform.Reader, error) {
	de, err := DetermineEncoding(r)
	if err != nil {
		return nil, err
	}
	return transform.NewReader(r, de.NewDecoder()), nil
}

// DetermineEncoding 判断传输来的文本的字符集格式是什么
func DetermineEncoding(r *bufio.Reader) (encoding.Encoding, error) {
	d := chardet.NewTextDetector()
	peek, err := r.Peek(1024)
	if err != nil {
		return nil, err
	}
	res, _ := d.DetectBest(peek)
	return getEncoder(res.Charset), nil
}

func getEncoder(encodingName string) encoding.Encoding {
	if e, ok := encodingMaps[encodingName]; ok {
		return e
	}
	return nil
}
