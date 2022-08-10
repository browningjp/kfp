// Code generated by go-swagger; DO NOT EDIT.

package job_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V2beta1Job v2beta1 job
// swagger:model v2beta1Job
type V2beta1Job struct {

	// Output. The time this job is created.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// Optional input field. Describing the purpose of the job
	Description string `json:"description,omitempty"`

	// Input. Whether the job is enabled or not.
	Enabled bool `json:"enabled,omitempty"`

	// In case any error happens retrieving a job field, only job ID
	// and the error message is returned. Client has the flexibility of choosing
	// how to handle error. This is especially useful during listing call.
	Error string `json:"error,omitempty"`

	// Output. Unique run ID. Generated by API server.
	ID string `json:"id,omitempty"`

	// Required input field.
	// Specify how many runs can be executed concurrently. Rage [1-10]
	MaxConcurrency int64 `json:"max_concurrency,omitempty,string"`

	// mode
	Mode JobMode `json:"mode,omitempty"`

	// Required input field. Job name provided by user. Not unique.
	Name string `json:"name,omitempty"`

	// Optional input field. Whether the job should catch up if behind schedule.
	// If true, the job will only schedule the latest interval if behind schedule.
	// If false, the job will catch up on each past interval.
	NoCatchup bool `json:"no_catchup,omitempty"`

	// Required input field.
	// Describing what the pipeline manifest and parameters to use
	// for the scheduled job.
	PipelineSpec *V2beta1PipelineSpec `json:"pipeline_spec,omitempty"`

	// Optional input field. Specify which resource this job belongs to.
	ResourceReferences []*V2beta1ResourceReference `json:"resource_references"`

	// Optional input field. Specify which Kubernetes service account this job uses.
	ServiceAccount string `json:"service_account,omitempty"`

	// Output. The status of the job.
	// One of [Enable, Disable, Error]
	Status string `json:"status,omitempty"`

	// Required input field.
	// Specify how a run is triggered. Support cron mode or periodic mode.
	Trigger *V2beta1Trigger `json:"trigger,omitempty"`

	// Output. The last time this job is updated.
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this v2beta1 job
func (m *V2beta1Job) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePipelineSpec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceReferences(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrigger(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V2beta1Job) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V2beta1Job) validateMode(formats strfmt.Registry) error {

	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	if err := m.Mode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("mode")
		}
		return err
	}

	return nil
}

func (m *V2beta1Job) validatePipelineSpec(formats strfmt.Registry) error {

	if swag.IsZero(m.PipelineSpec) { // not required
		return nil
	}

	if m.PipelineSpec != nil {
		if err := m.PipelineSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pipeline_spec")
			}
			return err
		}
	}

	return nil
}

func (m *V2beta1Job) validateResourceReferences(formats strfmt.Registry) error {

	if swag.IsZero(m.ResourceReferences) { // not required
		return nil
	}

	for i := 0; i < len(m.ResourceReferences); i++ {
		if swag.IsZero(m.ResourceReferences[i]) { // not required
			continue
		}

		if m.ResourceReferences[i] != nil {
			if err := m.ResourceReferences[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resource_references" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V2beta1Job) validateTrigger(formats strfmt.Registry) error {

	if swag.IsZero(m.Trigger) { // not required
		return nil
	}

	if m.Trigger != nil {
		if err := m.Trigger.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trigger")
			}
			return err
		}
	}

	return nil
}

func (m *V2beta1Job) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V2beta1Job) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2beta1Job) UnmarshalBinary(b []byte) error {
	var res V2beta1Job
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}