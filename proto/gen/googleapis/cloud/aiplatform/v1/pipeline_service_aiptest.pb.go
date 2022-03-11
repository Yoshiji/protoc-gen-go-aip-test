// Code generated by protoc-gen-go-aip-test. DO NOT EDIT.

package aiplatform

import (
	context "context"
	cmpopts "github.com/google/go-cmp/cmp/cmpopts"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protocmp "google.golang.org/protobuf/testing/protocmp"
	assert "gotest.tools/v3/assert"
	strings "strings"
	testing "testing"
	time "time"
)

type PipelineServiceTestSuite struct {
	T *testing.T
	// Server to test.
	Server PipelineServiceServer
}

func (fx PipelineServiceTestSuite) TestPipelineJob(ctx context.Context, options PipelineJobTestSuiteConfig) {
	fx.T.Run("PipelineJob", func(t *testing.T) {
		options.ctx = ctx
		options.service = fx.Server
		options.test(t)
	})
}

func (fx PipelineServiceTestSuite) TestTrainingPipeline(ctx context.Context, options TrainingPipelineTestSuiteConfig) {
	fx.T.Run("TrainingPipeline", func(t *testing.T) {
		options.ctx = ctx
		options.service = fx.Server
		options.test(t)
	})
}

type PipelineJobTestSuiteConfig struct {
	ctx        context.Context
	service    PipelineServiceServer
	currParent int

	// The parents to use when creating resources.
	// At least one parent needs to be set. Depending on methods available on the resource,
	// more may be required. If insufficient number of parents are
	// provided the test will fail.
	Parents []string
	// Create should return a resource which is valid to create, i.e.
	// all required fields set.
	Create func(parent string) *PipelineJob
	// Patterns of tests to skip.
	// For example if a service has a Get method:
	// Skip: ["Get"] will skip all tests for Get.
	// Skip: ["Get/persisted"] will only skip the subtest called "persisted" of Get.
	Skip []string
}

func (fx *PipelineJobTestSuiteConfig) test(t *testing.T) {
	t.Run("Create", fx.testCreate)
	t.Run("Get", fx.testGet)
	t.Run("List", fx.testList)
}

func (fx *PipelineJobTestSuiteConfig) testCreate(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no parent is provided.
	t.Run("missing parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.CreatePipelineJob(fx.ctx, &CreatePipelineJobRequest{
			Parent:      "",
			PipelineJob: fx.Create(fx.nextParent(t, false)),
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.CreatePipelineJob(fx.ctx, &CreatePipelineJobRequest{
			Parent:      "invalid resource name",
			PipelineJob: fx.Create(fx.nextParent(t, false)),
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Field create_time should be populated when the resource is created.
	t.Run("create time", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg, err := fx.service.CreatePipelineJob(fx.ctx, &CreatePipelineJobRequest{
			Parent:      parent,
			PipelineJob: fx.Create(parent),
		})
		assert.NilError(t, err)
		assert.Check(t, time.Since(msg.CreateTime.AsTime()) < time.Second)
	})

	// The created resource should be persisted and reachable with Get.
	t.Run("persisted", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg, err := fx.service.CreatePipelineJob(fx.ctx, &CreatePipelineJobRequest{
			Parent:      parent,
			PipelineJob: fx.Create(parent),
		})
		assert.NilError(t, err)
		persisted, err := fx.service.GetPipelineJob(fx.ctx, &GetPipelineJobRequest{
			Name: msg.Name,
		})
		assert.NilError(t, err)
		assert.DeepEqual(t, msg, persisted, protocmp.Transform())
	})

	// The method should fail with InvalidArgument if the resource has any
	// required fields and they are not provided.
	t.Run("required fields", func(t *testing.T) {
		fx.maybeSkip(t)
		t.Run(".pipeline_spec", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("pipeline_spec")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreatePipelineJob(fx.ctx, &CreatePipelineJobRequest{
				Parent:      parent,
				PipelineJob: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".runtime_config.gcs_output_directory", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetRuntimeConfig()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("gcs_output_directory")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreatePipelineJob(fx.ctx, &CreatePipelineJobRequest{
				Parent:      parent,
				PipelineJob: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".encryption_spec.kms_key_name", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetEncryptionSpec()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("kms_key_name")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreatePipelineJob(fx.ctx, &CreatePipelineJobRequest{
				Parent:      parent,
				PipelineJob: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
	})

	// The method should fail with InvalidArgument if the resource has any
	// resource references and they are invalid.
	t.Run("resource references", func(t *testing.T) {
		fx.maybeSkip(t)
		t.Run(".network", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			container.Network = "invalid resource name"
			_, err := fx.service.CreatePipelineJob(fx.ctx, &CreatePipelineJobRequest{
				Parent:      parent,
				PipelineJob: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
	})

}

func (fx *PipelineJobTestSuiteConfig) testGet(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetPipelineJob(fx.ctx, &GetPipelineJobRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetPipelineJob(fx.ctx, &GetPipelineJobRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Resource should be returned without errors if it exists.
	t.Run("exists", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created, err := fx.service.CreatePipelineJob(fx.ctx, &CreatePipelineJobRequest{
			Parent:      parent,
			PipelineJob: fx.Create(parent),
		})
		assert.NilError(t, err)
		msg, err := fx.service.GetPipelineJob(fx.ctx, &GetPipelineJobRequest{
			Name: created.Name,
		})
		assert.NilError(t, err)
		assert.DeepEqual(t, msg, created, protocmp.Transform())
	})

	// Method should fail with NotFound if the resource does not exist.
	t.Run("not found", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created, err := fx.service.CreatePipelineJob(fx.ctx, &CreatePipelineJobRequest{
			Parent:      parent,
			PipelineJob: fx.Create(parent),
		})
		assert.NilError(t, err)
		_, err = fx.service.GetPipelineJob(fx.ctx, &GetPipelineJobRequest{
			Name: created.Name + "notfound",
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

}

func (fx *PipelineJobTestSuiteConfig) testList(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.ListPipelineJobs(fx.ctx, &ListPipelineJobsRequest{
			Parent: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page token is not valid.
	t.Run("invalid page token", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.service.ListPipelineJobs(fx.ctx, &ListPipelineJobsRequest{
			Parent:    parent,
			PageToken: "invalid page token",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page size is negative.
	t.Run("negative page size", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.service.ListPipelineJobs(fx.ctx, &ListPipelineJobsRequest{
			Parent:   parent,
			PageSize: -10,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	const resourcesCount = 15
	parent := fx.nextParent(t, true)
	parentMsgs := make([]*PipelineJob, resourcesCount)
	for i := 0; i < resourcesCount; i++ {
		msg, err := fx.service.CreatePipelineJob(fx.ctx, &CreatePipelineJobRequest{
			Parent:      parent,
			PipelineJob: fx.Create(parent),
		})
		assert.NilError(t, err)
		parentMsgs[i] = msg
	}

	// If parent is provided the method must only return resources
	// under that parent.
	t.Run("isolation", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.service.ListPipelineJobs(fx.ctx, &ListPipelineJobsRequest{
			Parent:   parent,
			PageSize: 999,
		})
		assert.NilError(t, err)
		assert.DeepEqual(
			t,
			parentMsgs,
			response.PipelineJobs,
			cmpopts.SortSlices(func(a, b *PipelineJob) bool {
				return a.Name < b.Name
			}),
			protocmp.Transform(),
		)
	})

	// If there are no more resources, next_page_token should not be set.
	t.Run("last page", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.service.ListPipelineJobs(fx.ctx, &ListPipelineJobsRequest{
			Parent:   parent,
			PageSize: resourcesCount,
		})
		assert.NilError(t, err)
		assert.Equal(t, "", response.NextPageToken)
	})

	// If there are more resources, next_page_token should be set.
	t.Run("more pages", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.service.ListPipelineJobs(fx.ctx, &ListPipelineJobsRequest{
			Parent:   parent,
			PageSize: resourcesCount - 1,
		})
		assert.NilError(t, err)
		assert.Check(t, response.NextPageToken != "")
	})

	// Listing resource one by one should eventually return all resources.
	t.Run("one by one", func(t *testing.T) {
		fx.maybeSkip(t)
		msgs := make([]*PipelineJob, 0, resourcesCount)
		var nextPageToken string
		for {
			response, err := fx.service.ListPipelineJobs(fx.ctx, &ListPipelineJobsRequest{
				Parent:    parent,
				PageSize:  1,
				PageToken: nextPageToken,
			})
			assert.NilError(t, err)
			assert.Equal(t, 1, len(response.PipelineJobs))
			msgs = append(msgs, response.PipelineJobs...)
			nextPageToken = response.NextPageToken
			if nextPageToken == "" {
				break
			}
		}
		assert.DeepEqual(
			t,
			parentMsgs,
			msgs,
			cmpopts.SortSlices(func(a, b *PipelineJob) bool {
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
			_, err := fx.service.DeletePipelineJob(fx.ctx, &DeletePipelineJobRequest{
				Name: parentMsgs[i].Name,
			})
			assert.NilError(t, err)
		}
		response, err := fx.service.ListPipelineJobs(fx.ctx, &ListPipelineJobsRequest{
			Parent:   parent,
			PageSize: 9999,
		})
		assert.NilError(t, err)
		assert.DeepEqual(
			t,
			parentMsgs[deleteCount:],
			response.PipelineJobs,
			cmpopts.SortSlices(func(a, b *PipelineJob) bool {
				return a.Name < b.Name
			}),
			protocmp.Transform(),
		)
	})

}

func (fx *PipelineJobTestSuiteConfig) nextParent(t *testing.T, pristine bool) string {
	if pristine {
		fx.currParent++
	}
	if fx.currParent >= len(fx.Parents) {
		t.Fatal("need at least", fx.currParent+1, "parents")
	}
	return fx.Parents[fx.currParent]
}

func (fx *PipelineJobTestSuiteConfig) peekNextParent(t *testing.T) string {
	next := fx.currParent + 1
	if next >= len(fx.Parents) {
		t.Fatal("need at least", next+1, "parents")
	}
	return fx.Parents[next]
}

func (fx *PipelineJobTestSuiteConfig) maybeSkip(t *testing.T) {
	for _, skip := range fx.Skip {
		if strings.Contains(t.Name(), skip) {
			t.Skip("skipped because of .Skip")
		}
	}
}

type TrainingPipelineTestSuiteConfig struct {
	ctx        context.Context
	service    PipelineServiceServer
	currParent int

	// The parents to use when creating resources.
	// At least one parent needs to be set. Depending on methods available on the resource,
	// more may be required. If insufficient number of parents are
	// provided the test will fail.
	Parents []string
	// Create should return a resource which is valid to create, i.e.
	// all required fields set.
	Create func(parent string) *TrainingPipeline
	// Patterns of tests to skip.
	// For example if a service has a Get method:
	// Skip: ["Get"] will skip all tests for Get.
	// Skip: ["Get/persisted"] will only skip the subtest called "persisted" of Get.
	Skip []string
}

func (fx *TrainingPipelineTestSuiteConfig) test(t *testing.T) {
	t.Run("Create", fx.testCreate)
	t.Run("Get", fx.testGet)
	t.Run("List", fx.testList)
}

func (fx *TrainingPipelineTestSuiteConfig) testCreate(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no parent is provided.
	t.Run("missing parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.CreateTrainingPipeline(fx.ctx, &CreateTrainingPipelineRequest{
			Parent:           "",
			TrainingPipeline: fx.Create(fx.nextParent(t, false)),
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.CreateTrainingPipeline(fx.ctx, &CreateTrainingPipelineRequest{
			Parent:           "invalid resource name",
			TrainingPipeline: fx.Create(fx.nextParent(t, false)),
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Field create_time should be populated when the resource is created.
	t.Run("create time", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg, err := fx.service.CreateTrainingPipeline(fx.ctx, &CreateTrainingPipelineRequest{
			Parent:           parent,
			TrainingPipeline: fx.Create(parent),
		})
		assert.NilError(t, err)
		assert.Check(t, time.Since(msg.CreateTime.AsTime()) < time.Second)
	})

	// The created resource should be persisted and reachable with Get.
	t.Run("persisted", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg, err := fx.service.CreateTrainingPipeline(fx.ctx, &CreateTrainingPipelineRequest{
			Parent:           parent,
			TrainingPipeline: fx.Create(parent),
		})
		assert.NilError(t, err)
		persisted, err := fx.service.GetTrainingPipeline(fx.ctx, &GetTrainingPipelineRequest{
			Name: msg.Name,
		})
		assert.NilError(t, err)
		assert.DeepEqual(t, msg, persisted, protocmp.Transform())
	})

	// The method should fail with InvalidArgument if the resource has any
	// required fields and they are not provided.
	t.Run("required fields", func(t *testing.T) {
		fx.maybeSkip(t)
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
			_, err := fx.service.CreateTrainingPipeline(fx.ctx, &CreateTrainingPipelineRequest{
				Parent:           parent,
				TrainingPipeline: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".input_data_config.filter_split.training_filter", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetInputDataConfig().GetFilterSplit()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("training_filter")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateTrainingPipeline(fx.ctx, &CreateTrainingPipelineRequest{
				Parent:           parent,
				TrainingPipeline: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".input_data_config.filter_split.validation_filter", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetInputDataConfig().GetFilterSplit()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("validation_filter")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateTrainingPipeline(fx.ctx, &CreateTrainingPipelineRequest{
				Parent:           parent,
				TrainingPipeline: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".input_data_config.filter_split.test_filter", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetInputDataConfig().GetFilterSplit()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("test_filter")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateTrainingPipeline(fx.ctx, &CreateTrainingPipelineRequest{
				Parent:           parent,
				TrainingPipeline: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".input_data_config.predefined_split.key", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetInputDataConfig().GetPredefinedSplit()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("key")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateTrainingPipeline(fx.ctx, &CreateTrainingPipelineRequest{
				Parent:           parent,
				TrainingPipeline: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".input_data_config.timestamp_split.key", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetInputDataConfig().GetTimestampSplit()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("key")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateTrainingPipeline(fx.ctx, &CreateTrainingPipelineRequest{
				Parent:           parent,
				TrainingPipeline: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".input_data_config.stratified_split.key", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetInputDataConfig().GetStratifiedSplit()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("key")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateTrainingPipeline(fx.ctx, &CreateTrainingPipelineRequest{
				Parent:           parent,
				TrainingPipeline: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".input_data_config.gcs_destination.output_uri_prefix", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetInputDataConfig().GetGcsDestination()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("output_uri_prefix")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateTrainingPipeline(fx.ctx, &CreateTrainingPipelineRequest{
				Parent:           parent,
				TrainingPipeline: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".input_data_config.bigquery_destination.output_uri", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetInputDataConfig().GetBigqueryDestination()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("output_uri")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateTrainingPipeline(fx.ctx, &CreateTrainingPipelineRequest{
				Parent:           parent,
				TrainingPipeline: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".input_data_config.dataset_id", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetInputDataConfig()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("dataset_id")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateTrainingPipeline(fx.ctx, &CreateTrainingPipelineRequest{
				Parent:           parent,
				TrainingPipeline: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".training_task_definition", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("training_task_definition")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateTrainingPipeline(fx.ctx, &CreateTrainingPipelineRequest{
				Parent:           parent,
				TrainingPipeline: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".training_task_inputs", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("training_task_inputs")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateTrainingPipeline(fx.ctx, &CreateTrainingPipelineRequest{
				Parent:           parent,
				TrainingPipeline: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".model_to_upload.display_name", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetModelToUpload()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("display_name")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateTrainingPipeline(fx.ctx, &CreateTrainingPipelineRequest{
				Parent:           parent,
				TrainingPipeline: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".model_to_upload.container_spec.image_uri", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetModelToUpload().GetContainerSpec()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("image_uri")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateTrainingPipeline(fx.ctx, &CreateTrainingPipelineRequest{
				Parent:           parent,
				TrainingPipeline: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".model_to_upload.explanation_spec.parameters", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetModelToUpload().GetExplanationSpec()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("parameters")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateTrainingPipeline(fx.ctx, &CreateTrainingPipelineRequest{
				Parent:           parent,
				TrainingPipeline: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".model_to_upload.explanation_spec.parameters.sampled_shapley_attribution.path_count", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetModelToUpload().GetExplanationSpec().GetParameters().GetSampledShapleyAttribution()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("path_count")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateTrainingPipeline(fx.ctx, &CreateTrainingPipelineRequest{
				Parent:           parent,
				TrainingPipeline: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".model_to_upload.explanation_spec.parameters.integrated_gradients_attribution.step_count", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetModelToUpload().GetExplanationSpec().GetParameters().GetIntegratedGradientsAttribution()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("step_count")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateTrainingPipeline(fx.ctx, &CreateTrainingPipelineRequest{
				Parent:           parent,
				TrainingPipeline: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".model_to_upload.explanation_spec.parameters.xrai_attribution.step_count", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetModelToUpload().GetExplanationSpec().GetParameters().GetXraiAttribution()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("step_count")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateTrainingPipeline(fx.ctx, &CreateTrainingPipelineRequest{
				Parent:           parent,
				TrainingPipeline: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".model_to_upload.explanation_spec.metadata", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetModelToUpload().GetExplanationSpec()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("metadata")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateTrainingPipeline(fx.ctx, &CreateTrainingPipelineRequest{
				Parent:           parent,
				TrainingPipeline: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".model_to_upload.explanation_spec.metadata.inputs", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetModelToUpload().GetExplanationSpec().GetMetadata()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("inputs")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateTrainingPipeline(fx.ctx, &CreateTrainingPipelineRequest{
				Parent:           parent,
				TrainingPipeline: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".model_to_upload.explanation_spec.metadata.outputs", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetModelToUpload().GetExplanationSpec().GetMetadata()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("outputs")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateTrainingPipeline(fx.ctx, &CreateTrainingPipelineRequest{
				Parent:           parent,
				TrainingPipeline: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".model_to_upload.encryption_spec.kms_key_name", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetModelToUpload().GetEncryptionSpec()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("kms_key_name")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateTrainingPipeline(fx.ctx, &CreateTrainingPipelineRequest{
				Parent:           parent,
				TrainingPipeline: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
	})

}

func (fx *TrainingPipelineTestSuiteConfig) testGet(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetTrainingPipeline(fx.ctx, &GetTrainingPipelineRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetTrainingPipeline(fx.ctx, &GetTrainingPipelineRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Resource should be returned without errors if it exists.
	t.Run("exists", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created, err := fx.service.CreateTrainingPipeline(fx.ctx, &CreateTrainingPipelineRequest{
			Parent:           parent,
			TrainingPipeline: fx.Create(parent),
		})
		assert.NilError(t, err)
		msg, err := fx.service.GetTrainingPipeline(fx.ctx, &GetTrainingPipelineRequest{
			Name: created.Name,
		})
		assert.NilError(t, err)
		assert.DeepEqual(t, msg, created, protocmp.Transform())
	})

	// Method should fail with NotFound if the resource does not exist.
	t.Run("not found", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created, err := fx.service.CreateTrainingPipeline(fx.ctx, &CreateTrainingPipelineRequest{
			Parent:           parent,
			TrainingPipeline: fx.Create(parent),
		})
		assert.NilError(t, err)
		_, err = fx.service.GetTrainingPipeline(fx.ctx, &GetTrainingPipelineRequest{
			Name: created.Name + "notfound",
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

}

func (fx *TrainingPipelineTestSuiteConfig) testList(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.ListTrainingPipelines(fx.ctx, &ListTrainingPipelinesRequest{
			Parent: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page token is not valid.
	t.Run("invalid page token", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.service.ListTrainingPipelines(fx.ctx, &ListTrainingPipelinesRequest{
			Parent:    parent,
			PageToken: "invalid page token",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page size is negative.
	t.Run("negative page size", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.service.ListTrainingPipelines(fx.ctx, &ListTrainingPipelinesRequest{
			Parent:   parent,
			PageSize: -10,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	const resourcesCount = 15
	parent := fx.nextParent(t, true)
	parentMsgs := make([]*TrainingPipeline, resourcesCount)
	for i := 0; i < resourcesCount; i++ {
		msg, err := fx.service.CreateTrainingPipeline(fx.ctx, &CreateTrainingPipelineRequest{
			Parent:           parent,
			TrainingPipeline: fx.Create(parent),
		})
		assert.NilError(t, err)
		parentMsgs[i] = msg
	}

	// If parent is provided the method must only return resources
	// under that parent.
	t.Run("isolation", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.service.ListTrainingPipelines(fx.ctx, &ListTrainingPipelinesRequest{
			Parent:   parent,
			PageSize: 999,
		})
		assert.NilError(t, err)
		assert.DeepEqual(
			t,
			parentMsgs,
			response.TrainingPipelines,
			cmpopts.SortSlices(func(a, b *TrainingPipeline) bool {
				return a.Name < b.Name
			}),
			protocmp.Transform(),
		)
	})

	// If there are no more resources, next_page_token should not be set.
	t.Run("last page", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.service.ListTrainingPipelines(fx.ctx, &ListTrainingPipelinesRequest{
			Parent:   parent,
			PageSize: resourcesCount,
		})
		assert.NilError(t, err)
		assert.Equal(t, "", response.NextPageToken)
	})

	// If there are more resources, next_page_token should be set.
	t.Run("more pages", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.service.ListTrainingPipelines(fx.ctx, &ListTrainingPipelinesRequest{
			Parent:   parent,
			PageSize: resourcesCount - 1,
		})
		assert.NilError(t, err)
		assert.Check(t, response.NextPageToken != "")
	})

	// Listing resource one by one should eventually return all resources.
	t.Run("one by one", func(t *testing.T) {
		fx.maybeSkip(t)
		msgs := make([]*TrainingPipeline, 0, resourcesCount)
		var nextPageToken string
		for {
			response, err := fx.service.ListTrainingPipelines(fx.ctx, &ListTrainingPipelinesRequest{
				Parent:    parent,
				PageSize:  1,
				PageToken: nextPageToken,
			})
			assert.NilError(t, err)
			assert.Equal(t, 1, len(response.TrainingPipelines))
			msgs = append(msgs, response.TrainingPipelines...)
			nextPageToken = response.NextPageToken
			if nextPageToken == "" {
				break
			}
		}
		assert.DeepEqual(
			t,
			parentMsgs,
			msgs,
			cmpopts.SortSlices(func(a, b *TrainingPipeline) bool {
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
			_, err := fx.service.DeleteTrainingPipeline(fx.ctx, &DeleteTrainingPipelineRequest{
				Name: parentMsgs[i].Name,
			})
			assert.NilError(t, err)
		}
		response, err := fx.service.ListTrainingPipelines(fx.ctx, &ListTrainingPipelinesRequest{
			Parent:   parent,
			PageSize: 9999,
		})
		assert.NilError(t, err)
		assert.DeepEqual(
			t,
			parentMsgs[deleteCount:],
			response.TrainingPipelines,
			cmpopts.SortSlices(func(a, b *TrainingPipeline) bool {
				return a.Name < b.Name
			}),
			protocmp.Transform(),
		)
	})

}

func (fx *TrainingPipelineTestSuiteConfig) nextParent(t *testing.T, pristine bool) string {
	if pristine {
		fx.currParent++
	}
	if fx.currParent >= len(fx.Parents) {
		t.Fatal("need at least", fx.currParent+1, "parents")
	}
	return fx.Parents[fx.currParent]
}

func (fx *TrainingPipelineTestSuiteConfig) peekNextParent(t *testing.T) string {
	next := fx.currParent + 1
	if next >= len(fx.Parents) {
		t.Fatal("need at least", next+1, "parents")
	}
	return fx.Parents[next]
}

func (fx *TrainingPipelineTestSuiteConfig) maybeSkip(t *testing.T) {
	for _, skip := range fx.Skip {
		if strings.Contains(t.Name(), skip) {
			t.Skip("skipped because of .Skip")
		}
	}
}
