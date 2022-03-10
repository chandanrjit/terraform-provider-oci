// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
//

package core

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v61/common"
	"strings"
)

// VcnDrgAttachmentNetworkUpdateDetails Specifies the update details for the VCN attachment.
type VcnDrgAttachmentNetworkUpdateDetails struct {

	// This is the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the route table that is used to route the traffic as it enters a VCN through this attachment.
	// For information about why you would associate a route table with a DRG attachment, see:
	//   * Transit Routing: Access to Multiple VCNs in Same Region (https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/transitrouting.htm)
	//   * Transit Routing: Private Access to Oracle Services (https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/transitroutingoracleservices.htm)
	RouteTableId *string `mandatory:"false" json:"routeTableId"`

	// Indicates whether the VCN CIDR(s) or the individual Subnet CIDR(s) are imported from the attachment.
	// Routes from the VCN Ingress Route Table are always imported.
	VcnRouteType VcnDrgAttachmentNetworkDetailsVcnRouteTypeEnum `mandatory:"false" json:"vcnRouteType,omitempty"`
}

func (m VcnDrgAttachmentNetworkUpdateDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VcnDrgAttachmentNetworkUpdateDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingVcnDrgAttachmentNetworkDetailsVcnRouteTypeEnum(string(m.VcnRouteType)); !ok && m.VcnRouteType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VcnRouteType: %s. Supported values are: %s.", m.VcnRouteType, strings.Join(GetVcnDrgAttachmentNetworkDetailsVcnRouteTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m VcnDrgAttachmentNetworkUpdateDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeVcnDrgAttachmentNetworkUpdateDetails VcnDrgAttachmentNetworkUpdateDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeVcnDrgAttachmentNetworkUpdateDetails
	}{
		"VCN",
		(MarshalTypeVcnDrgAttachmentNetworkUpdateDetails)(m),
	}

	return json.Marshal(&s)
}