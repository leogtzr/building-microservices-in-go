package shi

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type shi struct {
	locale             string
	pluralsCardinal    []locales.PluralRule
	pluralsOrdinal     []locales.PluralRule
	pluralsRange       []locales.PluralRule
	decimal            string
	group              string
	minus              string
	percent            string
	perMille           string
	timeSeparator      string
	inifinity          string
	currencies         []string // idx = enum of currency code
	monthsAbbreviated  []string
	monthsNarrow       []string
	monthsWide         []string
	daysAbbreviated    []string
	daysNarrow         []string
	daysShort          []string
	daysWide           []string
	periodsAbbreviated []string
	periodsNarrow      []string
	periodsShort       []string
	periodsWide        []string
	erasAbbreviated    []string
	erasNarrow         []string
	erasWide           []string
	timezones          map[string]string
}

// New returns a new instance of translator for the 'shi' locale
func New() locales.Translator {
	return &shi{
		locale:             "shi",
		pluralsCardinal:    []locales.PluralRule{2, 4, 6},
		pluralsOrdinal:     nil,
		pluralsRange:       nil,
		decimal:            ",",
		group:              " ",
		timeSeparator:      ":",
		currencies:         []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UZS", "VEB", "VEF", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		monthsAbbreviated:  []string{"", "ⵉⵏⵏ", "ⴱⵕⴰ", "ⵎⴰⵕ", "ⵉⴱⵔ", "ⵎⴰⵢ", "ⵢⵓⵏ", "ⵢⵓⵍ", "ⵖⵓⵛ", "ⵛⵓⵜ", "ⴽⵜⵓ", "ⵏⵓⵡ", "ⴷⵓⵊ"},
		monthsNarrow:       []string{"", "ⵉ", "ⴱ", "ⵎ", "ⵉ", "ⵎ", "ⵢ", "ⵢ", "ⵖ", "ⵛ", "ⴽ", "ⵏ", "ⴷ"},
		monthsWide:         []string{"", "ⵉⵏⵏⴰⵢⵔ", "ⴱⵕⴰⵢⵕ", "ⵎⴰⵕⵚ", "ⵉⴱⵔⵉⵔ", "ⵎⴰⵢⵢⵓ", "ⵢⵓⵏⵢⵓ", "ⵢⵓⵍⵢⵓⵣ", "ⵖⵓⵛⵜ", "ⵛⵓⵜⴰⵏⴱⵉⵔ", "ⴽⵜⵓⴱⵔ", "ⵏⵓⵡⴰⵏⴱⵉⵔ", "ⴷⵓⵊⴰⵏⴱⵉⵔ"},
		daysAbbreviated:    []string{"ⴰⵙⴰ", "ⴰⵢⵏ", "ⴰⵙⵉ", "ⴰⴽⵕ", "ⴰⴽⵡ", "ⴰⵙⵉⵎ", "ⴰⵙⵉⴹ"},
		daysWide:           []string{"ⴰⵙⴰⵎⴰⵙ", "ⴰⵢⵏⴰⵙ", "ⴰⵙⵉⵏⴰⵙ", "ⴰⴽⵕⴰⵙ", "ⴰⴽⵡⴰⵙ", "ⵙⵉⵎⵡⴰⵙ", "ⴰⵙⵉⴹⵢⴰⵙ"},
		periodsAbbreviated: []string{"ⵜⵉⴼⴰⵡⵜ", "ⵜⴰⴷⴳⴳⵯⴰⵜ"},
		periodsWide:        []string{"ⵜⵉⴼⴰⵡⵜ", "ⵜⴰⴷⴳⴳⵯⴰⵜ"},
		erasAbbreviated:    []string{"ⴷⴰⵄ", "ⴷⴼⵄ"},
		erasNarrow:         []string{"", ""},
		erasWide:           []string{"ⴷⴰⵜ ⵏ ⵄⵉⵙⴰ", "ⴷⴼⴼⵉⵔ ⵏ ⵄⵉⵙⴰ"},
		timezones:          map[string]string{"AKDT": "AKDT", "EDT": "EDT", "CLST": "CLST", "WIT": "WIT", "LHST": "LHST", "UYT": "UYT", "WEZ": "WEZ", "CHADT": "CHADT", "MST": "MST", "AWST": "AWST", "EAT": "EAT", "WAST": "WAST", "WESZ": "WESZ", "GFT": "GFT", "SAST": "SAST", "MDT": "MDT", "BOT": "BOT", "WIB": "WIB", "GYT": "GYT", "NZST": "NZST", "LHDT": "LHDT", "OEZ": "OEZ", "CHAST": "CHAST", "NZDT": "NZDT", "AEDT": "AEDT", "ARST": "ARST", "WARST": "WARST", "TMT": "TMT", "AWDT": "AWDT", "GMT": "GMT", "SRT": "SRT", "BT": "BT", "OESZ": "OESZ", "HAST": "HAST", "WAT": "WAT", "EST": "EST", "CLT": "CLT", "MEZ": "MEZ", "ACST": "ACST", "VET": "VET", "HADT": "HADT", "ART": "ART", "IST": "IST", "CDT": "CDT", "HAT": "HAT", "MESZ": "MESZ", "ACWST": "ACWST", "PST": "PST", "WITA": "WITA", "JDT": "JDT", "ACDT": "ACDT", "SGT": "SGT", "AKST": "AKST", "CST": "CST", "JST": "JST", "AST": "AST", "MYT": "MYT", "ECT": "ECT", "PDT": "PDT", "TMST": "TMST", "HKT": "HKT", "COT": "COT", "COST": "COST", "ADT": "ADT", "CAT": "CAT", "HNT": "HNT", "UYST": "UYST", "ACWDT": "ACWDT", "AEST": "AEST", "WART": "WART", "∅∅∅": "∅∅∅", "ChST": "ChST", "HKST": "HKST"},
	}
}

// Locale returns the current translators string locale
func (shi *shi) Locale() string {
	return shi.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'shi'
func (shi *shi) PluralsCardinal() []locales.PluralRule {
	return shi.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'shi'
func (shi *shi) PluralsOrdinal() []locales.PluralRule {
	return shi.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'shi'
func (shi *shi) PluralsRange() []locales.PluralRule {
	return shi.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'shi'
func (shi *shi) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if (i == 0) || (n == 1) {
		return locales.PluralRuleOne
	} else if n >= 2 && n <= 10 {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'shi'
func (shi *shi) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'shi'
func (shi *shi) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (shi *shi) MonthAbbreviated(month time.Month) string {
	return shi.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (shi *shi) MonthsAbbreviated() []string {
	return shi.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (shi *shi) MonthNarrow(month time.Month) string {
	return shi.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (shi *shi) MonthsNarrow() []string {
	return shi.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (shi *shi) MonthWide(month time.Month) string {
	return shi.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (shi *shi) MonthsWide() []string {
	return shi.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (shi *shi) WeekdayAbbreviated(weekday time.Weekday) string {
	return shi.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (shi *shi) WeekdaysAbbreviated() []string {
	return shi.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (shi *shi) WeekdayNarrow(weekday time.Weekday) string {
	return shi.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (shi *shi) WeekdaysNarrow() []string {
	return shi.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (shi *shi) WeekdayShort(weekday time.Weekday) string {
	return shi.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (shi *shi) WeekdaysShort() []string {
	return shi.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (shi *shi) WeekdayWide(weekday time.Weekday) string {
	return shi.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (shi *shi) WeekdaysWide() []string {
	return shi.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'shi' and handles both Whole and Real numbers based on 'v'
func (shi *shi) FmtNumber(num float64, v uint64) string {

	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'shi' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (shi *shi) FmtPercent(num float64, v uint64) string {
	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'shi'
func (shi *shi) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := shi.currencies[currency]
	l := len(s) + len(symbol) + 1 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, shi.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(shi.group) - 1; j >= 0; j-- {
					b = append(b, shi.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, shi.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, shi.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'shi'
// in accounting notation.
func (shi *shi) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := shi.currencies[currency]
	l := len(s) + len(symbol) + 1 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, shi.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(shi.group) - 1; j >= 0; j-- {
					b = append(b, shi.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, shi.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, shi.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, symbol...)
	} else {

		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'shi'
func (shi *shi) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'shi'
func (shi *shi) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, shi.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'shi'
func (shi *shi) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, shi.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'shi'
func (shi *shi) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, shi.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, shi.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'shi'
func (shi *shi) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'shi'
func (shi *shi) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'shi'
func (shi *shi) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'shi'
func (shi *shi) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}