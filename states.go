package cliargstojson

const (
	TokenTypeUnknown = iota
	TokenTypeNumber
	TokenTypeString
	TokenTypeBoolean
	TokenTypeArray
	TokenTypeObject
)

/**
  "person": {
    "name": "John Doe-Hopkins",
    "age": 25,
    "phones": [
      "+1 234 567 8901",
      "+1 234 567 8902"
    ],
    "companies": [
      {
        "g2": {
          "position": "developer",
          "years": 5
        }
      },
      {
        "cow ball": {
          "position": "manager",
          "years": 2
        }
      }
    ],
    "isMarried": false,
    "family": {
      "mother": {
        "name": "Jane Doe-Hopkins",
        "age": 55
      },
      "father": {
        "name": "Robert Doe-Hopkins",
        "age": 57
      }
    }
  }
*/
// -person. name="John Doe-Hopkins" age=30
// phones=["+1 234 567 8901", "+1 234 567 8902"]
// family. mother. name="Jane Doe-Hopkins" age=55
// father. name="Robert Doe-Hopkins" age=57
// -person. isMarried=false
// companies=[g2. position="developer" years=5, cow ball. position="manager" years=2]
//
const (
	StateSearchKey = iota
	StateParseKey  = iota
	StateParseValue
)
