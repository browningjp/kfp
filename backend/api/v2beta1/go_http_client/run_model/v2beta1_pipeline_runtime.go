// Code generated by go-swagger; DO NOT EDIT.

package run_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V2beta1PipelineRuntime v2beta1 pipeline runtime
// swagger:model v2beta1PipelineRuntime
type V2beta1PipelineRuntime struct {

	// Output. The runtime JSON manifest of the pipeline, including the status
	// of pipeline steps and fields need for UI visualization etc.
	PipelineManifest string `json:"pipeline_manifest,omitempty"`

	// Output. The runtime JSON manifest of the argo workflow.
	// This is deprecated after pipeline_runtime_manifest is in use.
	WorkflowManifest string `json:"workflow_manifest,omitempty"`
}

// Validate validates this v2beta1 pipeline runtime
func (m *V2beta1PipelineRuntime) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V2beta1PipelineRuntime) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2beta1PipelineRuntime) UnmarshalBinary(b []byte) error {
	var res V2beta1PipelineRuntime
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}