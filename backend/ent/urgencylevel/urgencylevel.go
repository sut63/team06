// Code generated by entc, DO NOT EDIT.

package urgencylevel

const (
	// Label holds the string label denoting the urgencylevel type in the database.
	Label = "urgency_level"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUrgencyName holds the string denoting the urgencyname field in the database.
	FieldUrgencyName = "urgency_name"

	// EdgeUrgencyLevelToTriageResult holds the string denoting the urgencyleveltotriageresult edge name in mutations.
	EdgeUrgencyLevelToTriageResult = "urgencyLevelToTriageResult"

	// Table holds the table name of the urgencylevel in the database.
	Table = "urgency_levels"
	// UrgencyLevelToTriageResultTable is the table the holds the urgencyLevelToTriageResult relation/edge.
	UrgencyLevelToTriageResultTable = "triage_results"
	// UrgencyLevelToTriageResultInverseTable is the table name for the TriageResult entity.
	// It exists in this package in order to avoid circular dependency with the "triageresult" package.
	UrgencyLevelToTriageResultInverseTable = "triage_results"
	// UrgencyLevelToTriageResultColumn is the table column denoting the urgencyLevelToTriageResult relation/edge.
	UrgencyLevelToTriageResultColumn = "urgency_level_urgency_level_to_triage_result"
)

// Columns holds all SQL columns for urgencylevel fields.
var Columns = []string{
	FieldID,
	FieldUrgencyName,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// UrgencyNameValidator is a validator for the "urgencyName" field. It is called by the builders before save.
	UrgencyNameValidator func(string) error
)