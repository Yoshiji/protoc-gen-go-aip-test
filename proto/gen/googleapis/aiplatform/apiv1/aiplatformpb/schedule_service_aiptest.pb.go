// Code generated by protoc-gen-go-aip-test. DO NOT EDIT.

package aiplatformpb

import (
	context "context"
	cmpopts "github.com/google/go-cmp/cmp/cmpopts"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	proto "google.golang.org/protobuf/proto"
	protocmp "google.golang.org/protobuf/testing/protocmp"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	assert "gotest.tools/v3/assert"
	strings "strings"
	testing "testing"
	time "time"
)

type ScheduleServiceTestSuite struct {
	T *testing.T
	// Server to test.
	Server ScheduleServiceServer
}

func (fx ScheduleServiceTestSuite) TestSchedule(ctx context.Context, options ScheduleTestSuiteConfig) {
	fx.T.Run("Schedule", func(t *testing.T) {
		options.ctx = ctx
		options.service = fx.Server
		options.test(t)
	})
}

type ScheduleTestSuiteConfig struct {
	ctx        context.Context
	service    ScheduleServiceServer
	currParent int

	// The parents to use when creating resources.
	// At least one parent needs to be set. Depending on methods available on the resource,
	// more may be required. If insufficient number of parents are
	// provided the test will fail.
	Parents []string
	// Create should return a resource which is valid to create, i.e.
	// all required fields set.
	Create func(parent string) *Schedule
	// Update should return a resource which is valid to update, i.e.
	// all required fields set.
	Update func(parent string) *Schedule
	// Patterns of tests to skip.
	// For example if a service has a Get method:
	// Skip: ["Get"] will skip all tests for Get.
	// Skip: ["Get/persisted"] will only skip the subtest called "persisted" of Get.
	Skip []string
}

func (fx *ScheduleTestSuiteConfig) test(t *testing.T) {
	t.Run("Create", fx.testCreate)
	t.Run("Get", fx.testGet)
	t.Run("Update", fx.testUpdate)
	t.Run("List", fx.testList)
	t.Run("Delete", fx.testDelete)
}

func (fx *ScheduleTestSuiteConfig) testCreate(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no parent is provided.
	t.Run("missing parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.CreateSchedule(fx.ctx, &CreateScheduleRequest{
			Parent:   "",
			Schedule: fx.Create(fx.nextParent(t, false)),
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.CreateSchedule(fx.ctx, &CreateScheduleRequest{
			Parent:   "invalid resource name",
			Schedule: fx.Create(fx.nextParent(t, false)),
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Field create_time should be populated when the resource is created.
	t.Run("create time", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg, err := fx.service.CreateSchedule(fx.ctx, &CreateScheduleRequest{
			Parent:   parent,
			Schedule: fx.Create(parent),
		})
		assert.NilError(t, err)
		assert.Check(t, msg.CreateTime != nil)
		assert.Check(t, !msg.CreateTime.AsTime().IsZero())
		assert.Check(t, !msg.CreateTime.AsTime().After(time.Now()))
	})

	// The created resource should be persisted and reachable with Get.
	t.Run("persisted", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg, err := fx.service.CreateSchedule(fx.ctx, &CreateScheduleRequest{
			Parent:   parent,
			Schedule: fx.Create(parent),
		})
		assert.NilError(t, err)
		persisted, err := fx.service.GetSchedule(fx.ctx, &GetScheduleRequest{
			Name: msg.Name,
		})
		assert.NilError(t, err)
		assert.DeepEqual(t, msg, persisted, protocmp.Transform())
	})

	// The method should fail with InvalidArgument if the resource has any
	// required fields and they are not provided.
	t.Run("required fields", func(t *testing.T) {
		fx.maybeSkip(t)
		t.Run(".create_pipeline_job_request.parent", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetCreatePipelineJobRequest()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("parent")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateSchedule(fx.ctx, &CreateScheduleRequest{
				Parent:   parent,
				Schedule: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".create_pipeline_job_request.pipeline_job", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetCreatePipelineJobRequest()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("pipeline_job")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateSchedule(fx.ctx, &CreateScheduleRequest{
				Parent:   parent,
				Schedule: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".create_pipeline_job_request.pipeline_job.runtime_config.gcs_output_directory", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetCreatePipelineJobRequest().GetPipelineJob().GetRuntimeConfig()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("gcs_output_directory")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateSchedule(fx.ctx, &CreateScheduleRequest{
				Parent:   parent,
				Schedule: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".create_pipeline_job_request.pipeline_job.encryption_spec.kms_key_name", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetCreatePipelineJobRequest().GetPipelineJob().GetEncryptionSpec()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("kms_key_name")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateSchedule(fx.ctx, &CreateScheduleRequest{
				Parent:   parent,
				Schedule: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".display_name", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("display_name")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateSchedule(fx.ctx, &CreateScheduleRequest{
				Parent:   parent,
				Schedule: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".max_concurrent_run_count", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("max_concurrent_run_count")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateSchedule(fx.ctx, &CreateScheduleRequest{
				Parent:   parent,
				Schedule: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
	})

	// The method should fail with InvalidArgument if the resource has any
	// resource references and they are invalid.
	t.Run("resource references", func(t *testing.T) {
		fx.maybeSkip(t)
		t.Run(".create_pipeline_job_request.parent", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetCreatePipelineJobRequest()
			if container == nil {
				t.Skip("not reachable")
			}
			container.Parent = "invalid resource name"
			_, err := fx.service.CreateSchedule(fx.ctx, &CreateScheduleRequest{
				Parent:   parent,
				Schedule: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".create_pipeline_job_request.pipeline_job.network", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetCreatePipelineJobRequest().GetPipelineJob()
			if container == nil {
				t.Skip("not reachable")
			}
			container.Network = "invalid resource name"
			_, err := fx.service.CreateSchedule(fx.ctx, &CreateScheduleRequest{
				Parent:   parent,
				Schedule: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
	})

}

func (fx *ScheduleTestSuiteConfig) testGet(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetSchedule(fx.ctx, &GetScheduleRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetSchedule(fx.ctx, &GetScheduleRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Resource should be returned without errors if it exists.
	t.Run("exists", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		msg, err := fx.service.GetSchedule(fx.ctx, &GetScheduleRequest{
			Name: created.Name,
		})
		assert.NilError(t, err)
		assert.DeepEqual(t, msg, created, protocmp.Transform())
	})

	// Method should fail with NotFound if the resource does not exist.
	t.Run("not found", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		_, err := fx.service.GetSchedule(fx.ctx, &GetScheduleRequest{
			Name: created.Name + "notfound",
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name only contains wildcards ('-')
	t.Run("only wildcards", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetSchedule(fx.ctx, &GetScheduleRequest{
			Name: "projects/-/locations/-/schedules/-",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *ScheduleTestSuiteConfig) testUpdate(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg := fx.Update(parent)
		msg.Name = ""
		_, err := fx.service.UpdateSchedule(fx.ctx, &UpdateScheduleRequest{
			Schedule: msg,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg := fx.Update(parent)
		msg.Name = "invalid resource name"
		_, err := fx.service.UpdateSchedule(fx.ctx, &UpdateScheduleRequest{
			Schedule: msg,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Field update_time should be updated when the resource is updated.
	t.Run("update time", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		updated, err := fx.service.UpdateSchedule(fx.ctx, &UpdateScheduleRequest{
			Schedule: created,
		})
		assert.NilError(t, err)
		assert.Check(t, updated.UpdateTime.AsTime().After(created.UpdateTime.AsTime()))
	})

	// The updated resource should be persisted and reachable with Get.
	t.Run("persisted", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		updated, err := fx.service.UpdateSchedule(fx.ctx, &UpdateScheduleRequest{
			Schedule: created,
		})
		assert.NilError(t, err)
		persisted, err := fx.service.GetSchedule(fx.ctx, &GetScheduleRequest{
			Name: updated.Name,
		})
		assert.NilError(t, err)
		assert.DeepEqual(t, updated, persisted, protocmp.Transform())
	})

	// The field create_time should be preserved when a '*'-update mask is used.
	t.Run("preserve create_time", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		originalCreateTime := created.CreateTime
		updated, err := fx.service.UpdateSchedule(fx.ctx, &UpdateScheduleRequest{
			Schedule: created,
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"*",
				},
			},
		})
		assert.NilError(t, err)
		assert.DeepEqual(t, originalCreateTime, updated.CreateTime, protocmp.Transform())
	})

	parent := fx.nextParent(t, false)
	created := fx.create(t, parent)
	// Method should fail with NotFound if the resource does not exist.
	t.Run("not found", func(t *testing.T) {
		fx.maybeSkip(t)
		msg := fx.Update(parent)
		msg.Name = created.Name + "notfound"
		_, err := fx.service.UpdateSchedule(fx.ctx, &UpdateScheduleRequest{
			Schedule: msg,
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// The method should fail with InvalidArgument if the update_mask is invalid.
	t.Run("invalid update mask", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.UpdateSchedule(fx.ctx, &UpdateScheduleRequest{
			Schedule: created,
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"invalid_field_xyz",
				},
			},
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if any required field is missing
	// when called with '*' update_mask.
	t.Run("required fields", func(t *testing.T) {
		fx.maybeSkip(t)
		t.Run(".create_pipeline_job_request.parent", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := proto.Clone(created).(*Schedule)
			container := msg.GetCreatePipelineJobRequest()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("parent")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.UpdateSchedule(fx.ctx, &UpdateScheduleRequest{
				Schedule: msg,
				UpdateMask: &fieldmaskpb.FieldMask{
					Paths: []string{
						"*",
					},
				},
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".create_pipeline_job_request.pipeline_job", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := proto.Clone(created).(*Schedule)
			container := msg.GetCreatePipelineJobRequest()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("pipeline_job")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.UpdateSchedule(fx.ctx, &UpdateScheduleRequest{
				Schedule: msg,
				UpdateMask: &fieldmaskpb.FieldMask{
					Paths: []string{
						"*",
					},
				},
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".create_pipeline_job_request.pipeline_job.runtime_config.gcs_output_directory", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := proto.Clone(created).(*Schedule)
			container := msg.GetCreatePipelineJobRequest().GetPipelineJob().GetRuntimeConfig()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("gcs_output_directory")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.UpdateSchedule(fx.ctx, &UpdateScheduleRequest{
				Schedule: msg,
				UpdateMask: &fieldmaskpb.FieldMask{
					Paths: []string{
						"*",
					},
				},
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".create_pipeline_job_request.pipeline_job.encryption_spec.kms_key_name", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := proto.Clone(created).(*Schedule)
			container := msg.GetCreatePipelineJobRequest().GetPipelineJob().GetEncryptionSpec()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("kms_key_name")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.UpdateSchedule(fx.ctx, &UpdateScheduleRequest{
				Schedule: msg,
				UpdateMask: &fieldmaskpb.FieldMask{
					Paths: []string{
						"*",
					},
				},
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".display_name", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := proto.Clone(created).(*Schedule)
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("display_name")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.UpdateSchedule(fx.ctx, &UpdateScheduleRequest{
				Schedule: msg,
				UpdateMask: &fieldmaskpb.FieldMask{
					Paths: []string{
						"*",
					},
				},
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".max_concurrent_run_count", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := proto.Clone(created).(*Schedule)
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("max_concurrent_run_count")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.UpdateSchedule(fx.ctx, &UpdateScheduleRequest{
				Schedule: msg,
				UpdateMask: &fieldmaskpb.FieldMask{
					Paths: []string{
						"*",
					},
				},
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
	})

}

func (fx *ScheduleTestSuiteConfig) testList(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.ListSchedules(fx.ctx, &ListSchedulesRequest{
			Parent: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page token is not valid.
	t.Run("invalid page token", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.service.ListSchedules(fx.ctx, &ListSchedulesRequest{
			Parent:    parent,
			PageToken: "invalid page token",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page size is negative.
	t.Run("negative page size", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.service.ListSchedules(fx.ctx, &ListSchedulesRequest{
			Parent:   parent,
			PageSize: -10,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	const resourcesCount = 15
	parent := fx.nextParent(t, true)
	parentMsgs := make([]*Schedule, resourcesCount)
	for i := 0; i < resourcesCount; i++ {
		parentMsgs[i] = fx.create(t, parent)
	}

	// If parent is provided the method must only return resources
	// under that parent.
	t.Run("isolation", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.service.ListSchedules(fx.ctx, &ListSchedulesRequest{
			Parent:   parent,
			PageSize: 999,
		})
		assert.NilError(t, err)
		assert.DeepEqual(
			t,
			parentMsgs,
			response.Schedules,
			cmpopts.SortSlices(func(a, b *Schedule) bool {
				return a.Name < b.Name
			}),
			protocmp.Transform(),
		)
	})

	// If there are no more resources, next_page_token should not be set.
	t.Run("last page", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.service.ListSchedules(fx.ctx, &ListSchedulesRequest{
			Parent:   parent,
			PageSize: resourcesCount,
		})
		assert.NilError(t, err)
		assert.Equal(t, "", response.NextPageToken)
	})

	// If there are more resources, next_page_token should be set.
	t.Run("more pages", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.service.ListSchedules(fx.ctx, &ListSchedulesRequest{
			Parent:   parent,
			PageSize: resourcesCount - 1,
		})
		assert.NilError(t, err)
		assert.Check(t, response.NextPageToken != "")
	})

	// Listing resource one by one should eventually return all resources.
	t.Run("one by one", func(t *testing.T) {
		fx.maybeSkip(t)
		msgs := make([]*Schedule, 0, resourcesCount)
		var nextPageToken string
		for {
			response, err := fx.service.ListSchedules(fx.ctx, &ListSchedulesRequest{
				Parent:    parent,
				PageSize:  1,
				PageToken: nextPageToken,
			})
			assert.NilError(t, err)
			assert.Equal(t, 1, len(response.Schedules))
			msgs = append(msgs, response.Schedules...)
			nextPageToken = response.NextPageToken
			if nextPageToken == "" {
				break
			}
		}
		assert.DeepEqual(
			t,
			parentMsgs,
			msgs,
			cmpopts.SortSlices(func(a, b *Schedule) bool {
				return a.Name < b.Name
			}),
			protocmp.Transform(),
		)
	})

	// Method should not return deleted resources.
	t.Run("deleted", func(t *testing.T) {
		fx.maybeSkip(t)
		const deleteCount = 5
		for i := 0; i < deleteCount; i++ {
			_, err := fx.service.DeleteSchedule(fx.ctx, &DeleteScheduleRequest{
				Name: parentMsgs[i].Name,
			})
			assert.NilError(t, err)
		}
		response, err := fx.service.ListSchedules(fx.ctx, &ListSchedulesRequest{
			Parent:   parent,
			PageSize: 9999,
		})
		assert.NilError(t, err)
		assert.DeepEqual(
			t,
			parentMsgs[deleteCount:],
			response.Schedules,
			cmpopts.SortSlices(func(a, b *Schedule) bool {
				return a.Name < b.Name
			}),
			protocmp.Transform(),
		)
	})

}

func (fx *ScheduleTestSuiteConfig) testDelete(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.DeleteSchedule(fx.ctx, &DeleteScheduleRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.DeleteSchedule(fx.ctx, &DeleteScheduleRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Resource should be deleted without errors if it exists.
	t.Run("exists", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		_, err := fx.service.DeleteSchedule(fx.ctx, &DeleteScheduleRequest{
			Name: created.Name,
		})
		assert.NilError(t, err)
	})

	// Method should fail with NotFound if the resource does not exist.
	t.Run("not found", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		_, err := fx.service.DeleteSchedule(fx.ctx, &DeleteScheduleRequest{
			Name: created.Name + "notfound",
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name only contains wildcards ('-')
	t.Run("only wildcards", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.DeleteSchedule(fx.ctx, &DeleteScheduleRequest{
			Name: "projects/-/locations/-/schedules/-",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *ScheduleTestSuiteConfig) nextParent(t *testing.T, pristine bool) string {
	if pristine {
		fx.currParent++
	}
	if fx.currParent >= len(fx.Parents) {
		t.Fatal("need at least", fx.currParent+1, "parents")
	}
	return fx.Parents[fx.currParent]
}

func (fx *ScheduleTestSuiteConfig) peekNextParent(t *testing.T) string {
	next := fx.currParent + 1
	if next >= len(fx.Parents) {
		t.Fatal("need at least", next+1, "parents")
	}
	return fx.Parents[next]
}

func (fx *ScheduleTestSuiteConfig) maybeSkip(t *testing.T) {
	for _, skip := range fx.Skip {
		if strings.Contains(t.Name(), skip) {
			t.Skip("skipped because of .Skip")
		}
	}
}

func (fx *ScheduleTestSuiteConfig) create(t *testing.T, parent string) *Schedule {
	t.Helper()
	created, err := fx.service.CreateSchedule(fx.ctx, &CreateScheduleRequest{
		Parent:   parent,
		Schedule: fx.Create(parent),
	})
	assert.NilError(t, err)
	return created
}
