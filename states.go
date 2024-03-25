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
          "company": "g2",
          "position": "developer",
          "years": 5
      },
      {
          "company":"cow ball"
          "position": "manager",
          "years": 2
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

// person{name="John Doe-Hopkins" age=25 phones["555-222" "777-52"]
// companies[
//  {company ="g2" position="developer" years=5}
//  {company ="cowbell" position="manager" years=2}
// ]
// isMarried=false family{mother{name="Jane Doe-Hopkins" age=57}
// father{name="Robert Doe-Hopkins" age=57}}

const (
	StateSearchKey = iota
	StateParseKey  = iota
	StateParseValue
)
