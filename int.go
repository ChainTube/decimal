package decimal

import "math"

// The number of decimals used to store:
// BIGINT allows up to 9223372036854775807, so with
// 4 digits before comma + 15 digits after comma we end at double precision
// https://www.postgresql.org/docs/9.1/datatype-numeric.html
// We use a global variable because this should not be changed (within the same application)
// or else database serialization & deserialization will give wrong results.
var shiftStorageDecimals = 8

// When using BIGINT in Postgres as a 64bit signed integer,
// this is the max number that we can store.
var maxDecimalValue = computeMaxSafeValue()

// SetStorageDecimals sets the default decimals to shift when creating
// new Decimals. The value can also be adjusted on each decimal
// object - but be sure to always use the same one for type of number!
func SetStorageDecimals(decimals int) {
	shiftStorageDecimals = decimals
	maxDecimalValue = computeMaxSafeValue()
}

// GetStorageDecimals returns the number of decimals used to store in database.
func GetStorageDecimals() int {
	return shiftStorageDecimals
}

// GetMaxSafeValue returns the maximal value that can be stored as Postgres BIGINT using
// our current shiftStorageDecimals setting.
// If in doubt, you should compare myDecimal.LessThanOrEqual(decimal.GetMaxSafeValue())
func GetMaxSafeValue() Decimal {
	return maxDecimalValue
}

// SetMaxSafeValue can be used to overwrite the
// automatically computed max safe value.
func SetMaxSafeValue(max Decimal) {
	maxDecimalValue = max
}

func computeMaxSafeValue() Decimal {
	nrStr := "9"
	for i := 0; i < shiftStorageDecimals; i++ {
		nrStr += "9"
	}
	nr, err := NewFromString(nrStr)
	if err != nil {
		panic(err)
	}

	maxNr := NewFromInt(math.MaxInt32) // MaxInt64 causes an overflow with Postgres BIGINT (although both 8 bytes)
	if maxNr.GreaterThan(nr) {
		return maxNr
	}
	return nr
}
