// Copyright (c) 2011-2015, HL7, Inc & The MITRE Corporation
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without modification,
// are permitted provided that the following conditions are met:
//
//     * Redistributions of source code must retain the above copyright notice, this
//       list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above copyright notice,
//       this list of conditions and the following disclaimer in the documentation
//       and/or other materials provided with the distribution.
//     * Neither the name of HL7 nor the names of its contributors may be used to
//       endorse or promote products derived from this software without specific
//       prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
// IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT,
// INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT
// NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
// PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
// WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

type DeviceUseRequest struct {
	DomainResource          `bson:",inline"`
	BodySiteCodeableConcept *CodeableConcept  `bson:"bodySiteCodeableConcept,omitempty" json:"bodySiteCodeableConcept,omitempty"`
	BodySiteReference       *Reference        `bson:"bodySiteReference,omitempty" json:"bodySiteReference,omitempty"`
	Status                  string            `bson:"status,omitempty" json:"status,omitempty"`
	Device                  *Reference        `bson:"device,omitempty" json:"device,omitempty"`
	Encounter               *Reference        `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Identifier              []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Indication              []CodeableConcept `bson:"indication,omitempty" json:"indication,omitempty"`
	Notes                   []string          `bson:"notes,omitempty" json:"notes,omitempty"`
	PrnReason               []CodeableConcept `bson:"prnReason,omitempty" json:"prnReason,omitempty"`
	OrderedOn               *FHIRDateTime     `bson:"orderedOn,omitempty" json:"orderedOn,omitempty"`
	RecordedOn              *FHIRDateTime     `bson:"recordedOn,omitempty" json:"recordedOn,omitempty"`
	Subject                 *Reference        `bson:"subject,omitempty" json:"subject,omitempty"`
	TimingTiming            *Timing           `bson:"timingTiming,omitempty" json:"timingTiming,omitempty"`
	TimingPeriod            *Period           `bson:"timingPeriod,omitempty" json:"timingPeriod,omitempty"`
	TimingDateTime          *FHIRDateTime     `bson:"timingDateTime,omitempty" json:"timingDateTime,omitempty"`
	Priority                string            `bson:"priority,omitempty" json:"priority,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *DeviceUseRequest) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "DeviceUseRequest"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to DeviceUseRequest), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *DeviceUseRequest) GetBSON() (interface{}, error) {
	x.ResourceType = "DeviceUseRequest"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "deviceUseRequest" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type deviceUseRequest DeviceUseRequest

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *DeviceUseRequest) UnmarshalJSON(data []byte) (err error) {
	x2 := deviceUseRequest{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = DeviceUseRequest(x2)
		return x.checkResourceType()
	}
	return
}

func (x *DeviceUseRequest) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "DeviceUseRequest"
	} else if x.ResourceType != "DeviceUseRequest" {
		return errors.New(fmt.Sprintf("Expected resourceType to be DeviceUseRequest, instead received %s", x.ResourceType))
	}
	return nil
}

type DeviceUseRequestPlus struct {
	DeviceUseRequest             `bson:",inline"`
	DeviceUseRequestPlusIncludes `bson:",inline"`
}

type DeviceUseRequestPlusIncludes struct {
	IncludedSubjectResources *[]Patient `bson:"_includedSubjectResources,omitempty"`
	IncludedPatientResources *[]Patient `bson:"_includedPatientResources,omitempty"`
	IncludedDeviceResources  *[]Device  `bson:"_includedDeviceResources,omitempty"`
}

func (d *DeviceUseRequestPlusIncludes) GetIncludedSubjectResource() (patient *Patient, err error) {
	if d.IncludedSubjectResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*d.IncludedSubjectResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*d.IncludedSubjectResources))
	} else if len(*d.IncludedSubjectResources) == 1 {
		patient = &(*d.IncludedSubjectResources)[0]
	}
	return
}

func (d *DeviceUseRequestPlusIncludes) GetIncludedPatientResource() (patient *Patient, err error) {
	if d.IncludedPatientResources == nil {
		err = errors.New("Included patients not requested")
	} else if len(*d.IncludedPatientResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*d.IncludedPatientResources))
	} else if len(*d.IncludedPatientResources) == 1 {
		patient = &(*d.IncludedPatientResources)[0]
	}
	return
}

func (d *DeviceUseRequestPlusIncludes) GetIncludedDeviceResource() (device *Device, err error) {
	if d.IncludedDeviceResources == nil {
		err = errors.New("Included devices not requested")
	} else if len(*d.IncludedDeviceResources) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*d.IncludedDeviceResources))
	} else if len(*d.IncludedDeviceResources) == 1 {
		device = &(*d.IncludedDeviceResources)[0]
	}
	return
}

func (d *DeviceUseRequestPlusIncludes) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if d.IncludedSubjectResources != nil {
		for _, r := range *d.IncludedSubjectResources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedPatientResources != nil {
		for _, r := range *d.IncludedPatientResources {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedDeviceResources != nil {
		for _, r := range *d.IncludedDeviceResources {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
