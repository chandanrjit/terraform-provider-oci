// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	categoryDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Optional, create: `${var.compartment_id}`},
	}

	CategoryResourceConfig = ""
)

func TestMarketplaceCategoryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMarketplaceCategoryResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_marketplace_categories.test_categories"

	saveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_marketplace_categories", "test_categories", Required, Create, categoryDataSourceRepresentation) +
					compartmentIdVariableStr + CategoryResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(

					resource.TestCheckResourceAttrSet(datasourceName, "categories.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "categories.0.name"),
				),
			},
		},
	})
}
