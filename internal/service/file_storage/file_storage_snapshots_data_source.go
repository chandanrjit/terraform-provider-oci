// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package file_storage

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_file_storage "github.com/oracle/oci-go-sdk/v65/filestorage"
)

func FileStorageSnapshotsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFileStorageSnapshots,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"file_system_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"snapshots": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(FileStorageSnapshotResource()),
			},
		},
	}
}

func readFileStorageSnapshots(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageSnapshotsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()

	return tfresource.ReadResource(sync)
}

type FileStorageSnapshotsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_file_storage.FileStorageClient
	Res    *oci_file_storage.ListSnapshotsResponse
}

func (s *FileStorageSnapshotsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FileStorageSnapshotsDataSourceCrud) Get() error {
	request := oci_file_storage.ListSnapshotsRequest{}

	if fileSystemId, ok := s.D.GetOkExists("file_system_id"); ok {
		tmp := fileSystemId.(string)
		request.FileSystemId = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_file_storage.ListSnapshotsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "file_storage")

	response, err := s.Client.ListSnapshots(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSnapshots(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FileStorageSnapshotsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FileStorageSnapshotsDataSource-", FileStorageSnapshotsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		snapshot := map[string]interface{}{
			"file_system_id": *r.FileSystemId,
		}

		if r.DefinedTags != nil {
			snapshot["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		snapshot["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			snapshot["id"] = *r.Id
		}

		if r.IsCloneSource != nil {
			snapshot["is_clone_source"] = *r.IsCloneSource
		}

		if r.LifecycleDetails != nil {
			snapshot["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.Name != nil {
			snapshot["name"] = *r.Name
		}

		if r.ProvenanceId != nil {
			snapshot["provenance_id"] = *r.ProvenanceId
		}

		snapshot["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			snapshot["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, snapshot)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, FileStorageSnapshotsDataSource().Schema["snapshots"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("snapshots", resources); err != nil {
		return err
	}

	return nil
}
