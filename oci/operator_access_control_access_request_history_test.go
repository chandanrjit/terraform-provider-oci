// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	accessReqId = getEnvSettingWithBlankDefault("test_access_req_id")

	accessRequestHistoryDataSourceRepresentation = map[string]interface{}{
		"access_request_id": Representation{RepType: Required, Create: accessReqId},
	}
	accessRequestHistorySingularDataSourceRepresentation = map[string]interface{}{
		"access_request_id": Representation{RepType: Required, Create: accessReqId},
	}

	AccessRequestHistoryResourceConfig = GenerateDataSourceFromRepresentationMap("oci_operator_access_control_access_requests", "test_access_requests", Required, Create, accessRequestDataSourceRepresentation)
)

// issue-routing-tag: operator_access_control/default
func TestOperatorAccessControlAccessRequestHistoryResource_basic(t *testing.T) {
	t.Skip("Access Requests are created outside customer api. Access requests may not be available all the time")
	httpreplay.SetScenario("TestOperatorAccessControlAccessRequestHistoryResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_operator_access_control_access_request_history.test_access_request_history"
	singularDatasourceName := "data.oci_operator_access_control_access_request_history.test_access_request_history"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_operator_access_control_access_request_history", "test_access_request_history", Required, Create, accessRequestHistoryDataSourceRepresentation) +
				compartmentIdVariableStr + AccessRequestHistoryResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "access_request_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "items.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_operator_access_control_access_request_history", "test_access_request_history", Required, Create, accessRequestHistorySingularDataSourceRepresentation) +
				compartmentIdVariableStr + AccessRequestHistoryResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "access_request_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "items.#"),
			),
		},
	})
}
