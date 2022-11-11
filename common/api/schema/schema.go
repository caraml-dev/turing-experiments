// Package schema provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.1 DO NOT EDIT.
package schema

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/pkg/errors"
)

// Defines values for ExperimentField.
const (
	ExperimentFieldEndTime ExperimentField = "end_time"

	ExperimentFieldId ExperimentField = "id"

	ExperimentFieldName ExperimentField = "name"

	ExperimentFieldStartTime ExperimentField = "start_time"

	ExperimentFieldStatusFriendly ExperimentField = "status_friendly"

	ExperimentFieldTier ExperimentField = "tier"

	ExperimentFieldTreatments ExperimentField = "treatments"

	ExperimentFieldType ExperimentField = "type"

	ExperimentFieldUpdatedAt ExperimentField = "updated_at"
)

// Defines values for ExperimentStatus.
const (
	ExperimentStatusActive ExperimentStatus = "active"

	ExperimentStatusInactive ExperimentStatus = "inactive"
)

// Defines values for ExperimentStatusFriendly.
const (
	ExperimentStatusFriendlyCompleted ExperimentStatusFriendly = "completed"

	ExperimentStatusFriendlyDeactivated ExperimentStatusFriendly = "deactivated"

	ExperimentStatusFriendlyRunning ExperimentStatusFriendly = "running"

	ExperimentStatusFriendlyScheduled ExperimentStatusFriendly = "scheduled"
)

// Defines values for ExperimentTier.
const (
	ExperimentTierDefault ExperimentTier = "default"

	ExperimentTierOverride ExperimentTier = "override"
)

// Defines values for ExperimentType.
const (
	ExperimentTypeAB ExperimentType = "A/B"

	ExperimentTypeSwitchback ExperimentType = "Switchback"
)

// Defines values for SegmentField.
const (
	SegmentFieldId SegmentField = "id"

	SegmentFieldName SegmentField = "name"
)

// Defines values for SegmenterScope.
const (
	SegmenterScopeGlobal SegmenterScope = "global"

	SegmenterScopeProject SegmenterScope = "project"
)

// Defines values for SegmenterStatus.
const (
	SegmenterStatusActive SegmenterStatus = "active"

	SegmenterStatusInactive SegmenterStatus = "inactive"
)

// Defines values for SegmenterType.
const (
	SegmenterTypeBool SegmenterType = "bool"

	SegmenterTypeInteger SegmenterType = "integer"

	SegmenterTypeReal SegmenterType = "real"

	SegmenterTypeString SegmenterType = "string"
)

// Defines values for TreatmentField.
const (
	TreatmentFieldId TreatmentField = "id"

	TreatmentFieldName TreatmentField = "name"
)

// Constraint defines model for Constraint.
type Constraint struct {
	AllowedValues []SegmenterValues `json:"allowed_values"`
	Options       *SegmenterOptions `json:"options,omitempty"`
	PreRequisites []PreRequisite    `json:"pre_requisites"`
}

// Error defines model for Error.
type Error struct {
	Code    string `json:"code"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

// Experiment defines model for Experiment.
type Experiment struct {
	CreatedAt   *time.Time         `json:"created_at,omitempty"`
	Description *string            `json:"description"`
	EndTime     *time.Time         `json:"end_time,omitempty"`
	Id          *int64             `json:"id,omitempty"`
	Interval    *int32             `json:"interval"`
	Name        *string            `json:"name,omitempty"`
	ProjectId   *int64             `json:"project_id,omitempty"`
	Segment     *ExperimentSegment `json:"segment,omitempty"`
	StartTime   *time.Time         `json:"start_time,omitempty"`
	Status      *ExperimentStatus  `json:"status,omitempty"`

	// The user-friendly classification of experiment statuses. The categories are
	// self-explanatory. Note that the current time plays a role in the definition
	// of some of these statuses.
	StatusFriendly *ExperimentStatusFriendly `json:"status_friendly,omitempty"`
	Tier           *ExperimentTier           `json:"tier,omitempty"`
	Treatments     *[]ExperimentTreatment    `json:"treatments,omitempty"`
	Type           *ExperimentType           `json:"type,omitempty"`
	UpdatedAt      *time.Time                `json:"updated_at,omitempty"`
	UpdatedBy      *string                   `json:"updated_by,omitempty"`
	Version        *int64                    `json:"version,omitempty"`
}

// ExperimentField defines model for ExperimentField.
type ExperimentField string

// ExperimentHistory defines model for ExperimentHistory.
type ExperimentHistory struct {
	CreatedAt    time.Time             `json:"created_at"`
	Description  *string               `json:"description"`
	EndTime      time.Time             `json:"end_time"`
	ExperimentId int64                 `json:"experiment_id"`
	Id           int64                 `json:"id"`
	Interval     *int32                `json:"interval"`
	Name         string                `json:"name"`
	Segment      ExperimentSegment     `json:"segment"`
	StartTime    time.Time             `json:"start_time"`
	Status       ExperimentStatus      `json:"status"`
	Tier         ExperimentTier        `json:"tier"`
	Treatments   []ExperimentTreatment `json:"treatments"`
	Type         ExperimentType        `json:"type"`
	UpdatedAt    time.Time             `json:"updated_at"`
	UpdatedBy    string                `json:"updated_by"`
	Version      int64                 `json:"version"`
}

// ExperimentSegment defines model for ExperimentSegment.
type ExperimentSegment map[string]interface{}

// ExperimentStatus defines model for ExperimentStatus.
type ExperimentStatus string

// The user-friendly classification of experiment statuses. The categories are
// self-explanatory. Note that the current time plays a role in the definition
// of some of these statuses.
type ExperimentStatusFriendly string

// ExperimentTier defines model for ExperimentTier.
type ExperimentTier string

// ExperimentTreatment defines model for ExperimentTreatment.
type ExperimentTreatment struct {

	// Configuration associated with the given treatment
	Configuration map[string]interface{} `json:"configuration"`

	// Name of the treatment
	Name string `json:"name"`

	// When the experiment is matched, the % traffic to be directed to the treatment.
	// Optional for Switchback Experiments.
	Traffic *int32 `json:"traffic,omitempty"`
}

// ExperimentType defines model for ExperimentType.
type ExperimentType string

// Paging defines model for Paging.
type Paging struct {

	// Number of the current page
	Page int32 `json:"page"`

	// Total number of pages
	Pages int32 `json:"pages"`

	// Total number of results matching the query criteria
	Total int32 `json:"total"`
}

// PreRequisite defines model for PreRequisite.
type PreRequisite struct {
	SegmenterName   string            `json:"segmenter_name"`
	SegmenterValues []SegmenterValues `json:"segmenter_values"`
}

// Project defines model for Project.
type Project struct {
	CreatedAt        time.Time `json:"created_at"`
	Id               int64     `json:"id"`
	RandomizationKey string    `json:"randomization_key"`
	Segmenters       []string  `json:"segmenters"`
	UpdatedAt        time.Time `json:"updated_at"`
	Username         string    `json:"username"`
}

// ProjectSegmenters defines model for ProjectSegmenters.
type ProjectSegmenters struct {

	// List of segmenters name within Project, in priority order.
	Names []string `json:"names"`

	// Mapping of segmenter to the configured experiment variables
	Variables ProjectSegmenters_Variables `json:"variables"`
}

// Mapping of segmenter to the configured experiment variables
type ProjectSegmenters_Variables struct {
	AdditionalProperties map[string][]string `json:"-"`
}

// ProjectSettings defines model for ProjectSettings.
type ProjectSettings struct {
	CreatedAt            time.Time         `json:"created_at"`
	EnableS2idClustering bool              `json:"enable_s2id_clustering"`
	Passkey              string            `json:"passkey"`
	ProjectId            int64             `json:"project_id"`
	RandomizationKey     string            `json:"randomization_key"`
	Segmenters           ProjectSegmenters `json:"segmenters"`

	// Object containing information to define a valid treatment schema
	TreatmentSchema *TreatmentSchema `json:"treatment_schema,omitempty"`
	UpdatedAt       time.Time        `json:"updated_at"`
	Username        string           `json:"username"`
	ValidationUrl   *string          `json:"validation_url,omitempty"`
}

// PubSub defines model for PubSub.
type PubSub struct {

	// Project name of the PubSub subscription
	Project *string `json:"project,omitempty"`

	// Topic name of the PubSub subscription
	TopicName *string `json:"topic_name,omitempty"`
}

// A rule that forms part of a definition of a valid treatment schema
type Rule struct {
	Name string `json:"name"`

	// A Go template expression that must return a boolean value
	Predicate string `json:"predicate"`
}

// List of rules that define a valid treatment schema
type Rules []Rule

// Segment defines model for Segment.
type Segment struct {
	CreatedAt *time.Time         `json:"created_at,omitempty"`
	Id        *int64             `json:"id,omitempty"`
	Name      *string            `json:"name,omitempty"`
	ProjectId *int64             `json:"project_id,omitempty"`
	Segment   *ExperimentSegment `json:"segment,omitempty"`
	UpdatedAt *time.Time         `json:"updated_at,omitempty"`
	UpdatedBy *string            `json:"updated_by,omitempty"`
}

// SegmentField defines model for SegmentField.
type SegmentField string

// SegmentHistory defines model for SegmentHistory.
type SegmentHistory struct {
	CreatedAt time.Time         `json:"created_at"`
	Id        int64             `json:"id"`
	Name      string            `json:"name"`
	Segment   ExperimentSegment `json:"segment"`
	SegmentId int64             `json:"segment_id"`
	UpdatedAt time.Time         `json:"updated_at"`
	UpdatedBy string            `json:"updated_by"`
	Version   int64             `json:"version"`
}

// Segmenter defines model for Segmenter.
type Segmenter struct {
	Constraints []Constraint     `json:"constraints"`
	CreatedAt   *time.Time       `json:"created_at,omitempty"`
	Description *string          `json:"description,omitempty"`
	MultiValued bool             `json:"multi_valued"`
	Name        string           `json:"name"`
	Options     SegmenterOptions `json:"options"`
	Required    bool             `json:"required"`
	Scope       *SegmenterScope  `json:"scope,omitempty"`
	Status      *SegmenterStatus `json:"status,omitempty"`

	// List of varying combination of variables in which this segmenter is can be derived from
	TreatmentRequestFields [][]string    `json:"treatment_request_fields"`
	Type                   SegmenterType `json:"type"`
	UpdatedAt              *time.Time    `json:"updated_at,omitempty"`
}

// SegmenterConfig defines model for SegmenterConfig.
type SegmenterConfig map[string]interface{}

// SegmenterOptions defines model for SegmenterOptions.
type SegmenterOptions struct {
	AdditionalProperties map[string]interface{} `json:"-"`
}

// SegmenterScope defines model for SegmenterScope.
type SegmenterScope string

// SegmenterStatus defines model for SegmenterStatus.
type SegmenterStatus string

// SegmenterType defines model for SegmenterType.
type SegmenterType string

// SegmenterValues defines model for SegmenterValues.
type SegmenterValues interface{}

// SelectedTreatment defines model for SelectedTreatment.
type SelectedTreatment struct {
	ExperimentId   int64                     `json:"experiment_id"`
	ExperimentName string                    `json:"experiment_name"`
	Metadata       SelectedTreatmentMetadata `json:"metadata"`
	Treatment      SelectedTreatmentData     `json:"treatment"`
}

// SelectedTreatmentData defines model for SelectedTreatmentData.
type SelectedTreatmentData struct {

	// Custom configuration associated with the given treatment
	Configuration map[string]interface{} `json:"configuration"`
	Id            *int64                 `json:"id,omitempty"`

	// Name of the treatment
	Name string `json:"name"`

	// When the experiment is matched, the % traffic to be directed to the treatment.
	// Optional for Switchback Experiments.
	Traffic *int32 `json:"traffic,omitempty"`
}

// SelectedTreatmentMetadata defines model for SelectedTreatmentMetadata.
type SelectedTreatmentMetadata struct {
	ExperimentType    ExperimentType `json:"experiment_type"`
	ExperimentVersion int64          `json:"experiment_version"`

	// The window id since the beginning of the current version of the Switchback experiment.
	// This field will only be set for Switchback experiments and the window id starts at 0.
	SwitchbackWindowId *int64 `json:"switchback_window_id,omitempty"`
}

// Treatment defines model for Treatment.
type Treatment struct {
	Configuration *map[string]interface{} `json:"configuration,omitempty"`
	CreatedAt     *time.Time              `json:"created_at,omitempty"`
	Id            *int64                  `json:"id,omitempty"`
	Name          *string                 `json:"name,omitempty"`
	ProjectId     *int64                  `json:"project_id,omitempty"`
	UpdatedAt     *time.Time              `json:"updated_at,omitempty"`
	UpdatedBy     *string                 `json:"updated_by,omitempty"`
}

// TreatmentField defines model for TreatmentField.
type TreatmentField string

// TreatmentHistory defines model for TreatmentHistory.
type TreatmentHistory struct {
	Configuration map[string]interface{} `json:"configuration"`
	CreatedAt     time.Time              `json:"created_at"`
	Id            int64                  `json:"id"`
	Name          string                 `json:"name"`
	TreatmentId   int64                  `json:"treatment_id"`
	UpdatedAt     time.Time              `json:"updated_at"`
	UpdatedBy     string                 `json:"updated_by"`
	Version       int64                  `json:"version"`
}

// Object containing information to define a valid treatment schema
type TreatmentSchema struct {

	// List of rules that define a valid treatment schema
	Rules Rules `json:"rules"`
}

// TreatmentServiceConfig defines model for TreatmentServiceConfig.
type TreatmentServiceConfig struct {
	PubSub          *PubSub          `json:"pub_sub,omitempty"`
	SegmenterConfig *SegmenterConfig `json:"segmenter_config,omitempty"`
}

// Getter for additional properties for ProjectSegmenters_Variables. Returns the specified
// element and whether it was found
func (a ProjectSegmenters_Variables) Get(fieldName string) (value []string, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for ProjectSegmenters_Variables
func (a *ProjectSegmenters_Variables) Set(fieldName string, value []string) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string][]string)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for ProjectSegmenters_Variables to handle AdditionalProperties
func (a *ProjectSegmenters_Variables) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string][]string)
		for fieldName, fieldBuf := range object {
			var fieldVal []string
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error unmarshaling field %s", fieldName))
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for ProjectSegmenters_Variables to handle AdditionalProperties
func (a ProjectSegmenters_Variables) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling '%s'", fieldName))
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for SegmenterOptions. Returns the specified
// element and whether it was found
func (a SegmenterOptions) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for SegmenterOptions
func (a *SegmenterOptions) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for SegmenterOptions to handle AdditionalProperties
func (a *SegmenterOptions) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error unmarshaling field %s", fieldName))
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for SegmenterOptions to handle AdditionalProperties
func (a SegmenterOptions) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling '%s'", fieldName))
		}
	}
	return json.Marshal(object)
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xaW6/jthH+K4TavmnPFmnRB7+ladM+NNlFfJA+5CwMWhzZTChSGVJ23IX/e8GL7rQu",
	"XiPZA+TpyNLMcDg3znw8H5NMFaWSII1ONh8TnR2hoO7xKyW1Qcqlsb9KVCWg4eC+USHUGdjuREXl33AD",
	"hXv4I0KebJI/vG0Fvw1S327hUIA0gN97vmuamEsJySahiPRif6vScCWXS3oX6K9pUiLsEH6uuOZmhVLv",
	"Eb6rucYaXdPEyURgyeaH4Rrp0BIfGn61/xEyYwX+E1Hh2IaZYmD/BnptkMuDpYeafvSlAK3pIcY1UNPJ",
	"bulrmVHtfikBuTVmREUEaoDtqPuWKyzsU8KogTeGF1bwSEcGOkPuvGKZZCUE3QtINgYriNCDZDsna/EK",
	"nPVouTR/+2tLx6WBA6AjtAFyomJI/pcvkvSWYh12SYu4g0pU1nq7xYpoH61zkdi6IoS34zUUzUoLaUNN",
	"pVcs5+kbzl2OHCQTl7Uivq75bB5xwOX8z9ybytiQK+pytCiBO0Jq5lhl8b8Xi7LU1zSpSrY6BWqe/SUa",
	"PidAHbJjNnaukxn7NQfhYhBkVdi85ywJcRv4xh4NjukFVicLezvuueNDZKetKv/m2ii8vJYaAo3iy7P4",
	"N687r6eM/J77j8n97pneD9lWVD9deqns6JpobCpDHUeDGhDc3RSIjjuaatLJ5kGl6Gx8us3YtlE8RdXE",
	"XlPbJM0MP1ktwsN0RRqcSJuP/bKSPB+BVBrwTV0aSSao1jznGbUkROWktTnx1gH9RCxjRg0cFHLQhCK8",
	"SA0ifwO/lIJKauvgE/lWGSDmSA0xlr5CtFKsqUkp6EUTSlAJIFw6AgY5l9yu+yJVTrQqwCpgjqChXfvF",
	"O9gbBCsp7a5T17WzSoB1t41sAcY9M3CWsn6ZMdZzSFgGOa2Ei/Lw1K7XvlEnQORszgNtSkaaX5nzQ4W0",
	"rvF933zV/Uyo1irjdhfkzM3R2evATyBJE6JJJOTqOtoX/S1tLBtjb/dhkOY5z8YS/nsE77NOdHBNCmqs",
	"G1L36U8ksBOjyB4I4wiZ3YBR/ZWfXqQfYagguUKyPXOTHfc0+4m0hgyOH50lMxWjb+RgkOnkfA6Fsvb5",
	"l2//nqRJq1TU4+/pwT6NnFyGWWXggKrYA9YuqBOj9HPK7BZTJ1VH0lkZKohshHuyRRKNZZ2XiKArYYKj",
	"uTw4/X+uAC8kQ24AOb3DSX5xv62k3l3MSb05dWRrXQ/Eu7n2AfDhY/tgSwNdIivH9+emqsd0kIt7NaSS",
	"qYL/z+XI7ie4TJuub7RxzRj0HXd1EBrwhg8HdnbH+8SJXAuK7bK3pwl3bHs77zvGCo9k4n+4NjZf2gWI",
	"pXS1m0sSBKf25CuRK+TmQhQywCfbrSy27Ykit420B6QY476Kvu+pGNesYbU6nI8882eKBuGLdKO5reu2",
	"ytel21ZyQH4CRnJUxSp9+6p8Q8vS1pCunerDoa7bwLpHTLvfkbcGceH90rXQpION4fKgH5N3IO2CO/0F",
	"Z7tMVNpWRX80BNK9UgKo9IVc61sJtxpiuSePp1HBYfh3B5SdJ5sT0jRAW0/++IpgnSw487uuUMwXjY5l",
	"FxaP2k+zZeSm+6PhV+231T7SNLTHQD9jgkd8LQndgxdCdLXvTkDjVFQlz3bxfvDZflsvNAbMfFeJyAJf",
	"EqxEGAWsvzUpKboyRDtdv//tnNk2hySEWRopvDfSBpgdX6Jq/EsRA0UpqHGtK4K286NXrKi0IQimQkko",
	"CUlK3Gkd3fuw3CTdtT/csM1ERbYm0l4VZxOYMsaipsU5I1KGO+Pnr9hmfB5Y8sMBjVgWhPUmsMnYDBG4",
	"Hgojfrp3Pglx84/LPft5wU0d9QOQ1AJOIxzpblioOV+jEEG4f1w+qXTuLCOp/wAoenwlVwnD/VjD4m3O",
	"zeD6hKvO1lGxFXWm5uHORuzWUS+Ge1u+Fu1t2iKrF2izy23yTzfgF9v9ZqrYc9nAbdG+nOt+P55ROdWH",
	"xxeM9dEpOR9BkkoDs+tlSv5YycwypsNF+lqsavvvgaIbG9+PRMfP6IDi1pE3CN8JT6a9dOzInkxqD+NF",
	"kd5RVN8c4nrXJBEB2zra64PmINTeQyqhlZw4b5ow7vA3MHODOE8KGEJmgSR1GRnw94MLGgQqpmV93wAz",
	"SsK7PNn8MI6wSMY3rzxYlVw/OKF+mp1AYe+5Bevw3KxsBRjKqKHzcT5Q8ZuasVtVVkv5h5Mwc30y3Ed3",
	"wc4O4vEdW3A1xl1powqSPQLqXt3p/I6Jz2Hit2NzKo3uu2jsCFjTsaWJbiyzO3PJ1Dnk8fimy38mnBHN",
	"ZQbO4Hs4cHeFNATjgxL16479W02fXuSzPRXdAUHOXAiipLhYx2owQ7+1fJpQyZzYjkqGov1gyJ/HXl15",
	"N9p2qUO3xLy85opqxPxKJsZfZeprDLly7mv4bk9+n6kf2l7plU54vQ10x7vufxEN6+Xdk94QCx1VqXeO",
	"1J6HhnJXlbj0W3IolVoADPUDB2vIaQ4m0iPTeNbpbQCeeAZtizuAMav9Tnt8cxJm9iho73Yua0QumhGC",
	"BpGsvLr/RcpVspGVELbtB0lLnmwSh+qao/Zfrv8PAAD//7mpDQoGLAAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
