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
)

// ModelServiceTestSuiteConfigProvider is the interface to implement to decide which resources
// that should be tested and how it's configured.
type ModelServiceTestSuiteConfigProvider interface {
	// ModelServiceModel should return a config, or nil, which means that the tests will be skipped.
	ModelServiceModel(t *testing.T) *ModelServiceModelTestSuiteConfig
	// ModelServiceModelEvaluation should return a config, or nil, which means that the tests will be skipped.
	ModelServiceModelEvaluation(t *testing.T) *ModelServiceModelEvaluationTestSuiteConfig
	// ModelServiceModelEvaluationSlice should return a config, or nil, which means that the tests will be skipped.
	ModelServiceModelEvaluationSlice(t *testing.T) *ModelServiceModelEvaluationSliceTestSuiteConfig
}

// testModelService is the main entrypoint for starting the AIP tests.
func testModelService(t *testing.T, s ModelServiceTestSuiteConfigProvider) {
	testModelServiceModel(t, s)
	testModelServiceModelEvaluation(t, s)
	testModelServiceModelEvaluationSlice(t, s)
}

func testModelServiceModel(t *testing.T, s ModelServiceTestSuiteConfigProvider) {
	t.Run("Model", func(t *testing.T) {
		config := s.ModelServiceModel(t)
		if config == nil {
			t.Skip("Method ModelServiceModel not implemented")
		}
		if config.Service == nil {
			t.Skip("Method ModelServiceModel.Service() not implemented")
		}
		if config.Context == nil {
			config.Context = func() context.Context { return context.Background() }
		}
		config.test(t)
	})
}

func testModelServiceModelEvaluation(t *testing.T, s ModelServiceTestSuiteConfigProvider) {
	t.Run("ModelEvaluation", func(t *testing.T) {
		config := s.ModelServiceModelEvaluation(t)
		if config == nil {
			t.Skip("Method ModelServiceModelEvaluation not implemented")
		}
		if config.Service == nil {
			t.Skip("Method ModelServiceModelEvaluation.Service() not implemented")
		}
		if config.Context == nil {
			config.Context = func() context.Context { return context.Background() }
		}
		config.test(t)
	})
}

func testModelServiceModelEvaluationSlice(t *testing.T, s ModelServiceTestSuiteConfigProvider) {
	t.Run("ModelEvaluationSlice", func(t *testing.T) {
		config := s.ModelServiceModelEvaluationSlice(t)
		if config == nil {
			t.Skip("Method ModelServiceModelEvaluationSlice not implemented")
		}
		if config.Service == nil {
			t.Skip("Method ModelServiceModelEvaluationSlice.Service() not implemented")
		}
		if config.Context == nil {
			config.Context = func() context.Context { return context.Background() }
		}
		config.test(t)
	})
}

type ModelServiceTestSuite struct {
	T *testing.T
	// Server to test.
	Server ModelServiceServer
}

func (fx ModelServiceTestSuite) TestModel(ctx context.Context, options ModelServiceModelTestSuiteConfig) {
	fx.T.Run("Model", func(t *testing.T) {
		options.Context = func() context.Context { return ctx }
		options.Service = func() ModelServiceServer { return fx.Server }
		options.test(t)
	})
}

func (fx ModelServiceTestSuite) TestModelEvaluation(ctx context.Context, options ModelServiceModelEvaluationTestSuiteConfig) {
	fx.T.Run("ModelEvaluation", func(t *testing.T) {
		options.Context = func() context.Context { return ctx }
		options.Service = func() ModelServiceServer { return fx.Server }
		options.test(t)
	})
}

func (fx ModelServiceTestSuite) TestModelEvaluationSlice(ctx context.Context, options ModelServiceModelEvaluationSliceTestSuiteConfig) {
	fx.T.Run("ModelEvaluationSlice", func(t *testing.T) {
		options.Context = func() context.Context { return ctx }
		options.Service = func() ModelServiceServer { return fx.Server }
		options.test(t)
	})
}

type ModelServiceModelTestSuiteConfig struct {
	currParent int

	// Service should return the service that should be tested.
	// The service will be used for several tests.
	Service func() ModelServiceServer
	// Context should return a new context.
	// The context will be used for several tests.
	Context func() context.Context
	// The parents to use when creating resources.
	// At least one parent needs to be set. Depending on methods available on the resource,
	// more may be required. If insufficient number of parents are
	// provided the test will fail.
	Parents []string
	// CreateResource should create a Model and return it.
	// If the field is not set, some tests will be skipped.
	//
	// This method is generated because service does not expose a Create
	// method (or it does not comply with AIP).
	CreateResource func(ctx context.Context, parent string) (*Model, error)
	// Update should return a resource which is valid to update, i.e.
	// all required fields set.
	Update func(parent string) *Model
	// Patterns of tests to skip.
	// For example if a service has a Get method:
	// Skip: ["Get"] will skip all tests for Get.
	// Skip: ["Get/persisted"] will only skip the subtest called "persisted" of Get.
	Skip []string
}

func (fx *ModelServiceModelTestSuiteConfig) test(t *testing.T) {
	t.Run("Get", fx.testGet)
	t.Run("Update", fx.testUpdate)
	t.Run("List", fx.testList)
	t.Run("Delete", fx.testDelete)
}

func (fx *ModelServiceModelTestSuiteConfig) testGet(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().GetModel(fx.Context(), &GetModelRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().GetModel(fx.Context(), &GetModelRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Resource should be returned without errors if it exists.
	t.Run("exists", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		msg, err := fx.Service().GetModel(fx.Context(), &GetModelRequest{
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
		_, err := fx.Service().GetModel(fx.Context(), &GetModelRequest{
			Name: created.Name + "notfound",
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name only contains wildcards ('-')
	t.Run("only wildcards", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().GetModel(fx.Context(), &GetModelRequest{
			Name: "projects/-/locations/-/models/-",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *ModelServiceModelTestSuiteConfig) testUpdate(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg := fx.Update(parent)
		msg.Name = ""
		_, err := fx.Service().UpdateModel(fx.Context(), &UpdateModelRequest{
			Model: msg,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg := fx.Update(parent)
		msg.Name = "invalid resource name"
		_, err := fx.Service().UpdateModel(fx.Context(), &UpdateModelRequest{
			Model: msg,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// The updated resource should be persisted and reachable with Get.
	t.Run("persisted", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		updated, err := fx.Service().UpdateModel(fx.Context(), &UpdateModelRequest{
			Model: created,
		})
		assert.NilError(t, err)
		persisted, err := fx.Service().GetModel(fx.Context(), &GetModelRequest{
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
		updated, err := fx.Service().UpdateModel(fx.Context(), &UpdateModelRequest{
			Model: created,
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
		_, err := fx.Service().UpdateModel(fx.Context(), &UpdateModelRequest{
			Model: msg,
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// The method should fail with InvalidArgument if the update_mask is invalid.
	t.Run("invalid update mask", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().UpdateModel(fx.Context(), &UpdateModelRequest{
			Model: created,
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
		t.Run(".display_name", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := proto.Clone(created).(*Model)
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("display_name")
			container.ProtoReflect().Clear(fd)
			_, err := fx.Service().UpdateModel(fx.Context(), &UpdateModelRequest{
				Model: msg,
				UpdateMask: &fieldmaskpb.FieldMask{
					Paths: []string{
						"*",
					},
				},
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".explanation_spec.parameters", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := proto.Clone(created).(*Model)
			container := msg.GetExplanationSpec()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("parameters")
			container.ProtoReflect().Clear(fd)
			_, err := fx.Service().UpdateModel(fx.Context(), &UpdateModelRequest{
				Model: msg,
				UpdateMask: &fieldmaskpb.FieldMask{
					Paths: []string{
						"*",
					},
				},
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".explanation_spec.parameters.sampled_shapley_attribution.path_count", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := proto.Clone(created).(*Model)
			container := msg.GetExplanationSpec().GetParameters().GetSampledShapleyAttribution()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("path_count")
			container.ProtoReflect().Clear(fd)
			_, err := fx.Service().UpdateModel(fx.Context(), &UpdateModelRequest{
				Model: msg,
				UpdateMask: &fieldmaskpb.FieldMask{
					Paths: []string{
						"*",
					},
				},
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".explanation_spec.parameters.integrated_gradients_attribution.step_count", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := proto.Clone(created).(*Model)
			container := msg.GetExplanationSpec().GetParameters().GetIntegratedGradientsAttribution()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("step_count")
			container.ProtoReflect().Clear(fd)
			_, err := fx.Service().UpdateModel(fx.Context(), &UpdateModelRequest{
				Model: msg,
				UpdateMask: &fieldmaskpb.FieldMask{
					Paths: []string{
						"*",
					},
				},
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".explanation_spec.parameters.xrai_attribution.step_count", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := proto.Clone(created).(*Model)
			container := msg.GetExplanationSpec().GetParameters().GetXraiAttribution()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("step_count")
			container.ProtoReflect().Clear(fd)
			_, err := fx.Service().UpdateModel(fx.Context(), &UpdateModelRequest{
				Model: msg,
				UpdateMask: &fieldmaskpb.FieldMask{
					Paths: []string{
						"*",
					},
				},
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".explanation_spec.parameters.examples.example_gcs_source.gcs_source.uris", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := proto.Clone(created).(*Model)
			container := msg.GetExplanationSpec().GetParameters().GetExamples().GetExampleGcsSource().GetGcsSource()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("uris")
			container.ProtoReflect().Clear(fd)
			_, err := fx.Service().UpdateModel(fx.Context(), &UpdateModelRequest{
				Model: msg,
				UpdateMask: &fieldmaskpb.FieldMask{
					Paths: []string{
						"*",
					},
				},
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".explanation_spec.metadata.inputs", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := proto.Clone(created).(*Model)
			container := msg.GetExplanationSpec().GetMetadata()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("inputs")
			container.ProtoReflect().Clear(fd)
			_, err := fx.Service().UpdateModel(fx.Context(), &UpdateModelRequest{
				Model: msg,
				UpdateMask: &fieldmaskpb.FieldMask{
					Paths: []string{
						"*",
					},
				},
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".explanation_spec.metadata.outputs", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := proto.Clone(created).(*Model)
			container := msg.GetExplanationSpec().GetMetadata()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("outputs")
			container.ProtoReflect().Clear(fd)
			_, err := fx.Service().UpdateModel(fx.Context(), &UpdateModelRequest{
				Model: msg,
				UpdateMask: &fieldmaskpb.FieldMask{
					Paths: []string{
						"*",
					},
				},
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".encryption_spec.kms_key_name", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := proto.Clone(created).(*Model)
			container := msg.GetEncryptionSpec()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("kms_key_name")
			container.ProtoReflect().Clear(fd)
			_, err := fx.Service().UpdateModel(fx.Context(), &UpdateModelRequest{
				Model: msg,
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

func (fx *ModelServiceModelTestSuiteConfig) testList(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().ListModels(fx.Context(), &ListModelsRequest{
			Parent: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page token is not valid.
	t.Run("invalid page token", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.Service().ListModels(fx.Context(), &ListModelsRequest{
			Parent:    parent,
			PageToken: "invalid page token",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page size is negative.
	t.Run("negative page size", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.Service().ListModels(fx.Context(), &ListModelsRequest{
			Parent:   parent,
			PageSize: -10,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	const resourcesCount = 15
	parent := fx.nextParent(t, true)
	parentMsgs := make([]*Model, resourcesCount)
	for i := 0; i < resourcesCount; i++ {
		parentMsgs[i] = fx.create(t, parent)
	}

	// If parent is provided the method must only return resources
	// under that parent.
	t.Run("isolation", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.Service().ListModels(fx.Context(), &ListModelsRequest{
			Parent:   parent,
			PageSize: 999,
		})
		assert.NilError(t, err)
		assert.DeepEqual(
			t,
			parentMsgs,
			response.Models,
			cmpopts.SortSlices(func(a, b *Model) bool {
				return a.Name < b.Name
			}),
			protocmp.Transform(),
		)
	})

	// If there are no more resources, next_page_token should not be set.
	t.Run("last page", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.Service().ListModels(fx.Context(), &ListModelsRequest{
			Parent:   parent,
			PageSize: resourcesCount,
		})
		assert.NilError(t, err)
		assert.Equal(t, "", response.NextPageToken)
	})

	// If there are more resources, next_page_token should be set.
	t.Run("more pages", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.Service().ListModels(fx.Context(), &ListModelsRequest{
			Parent:   parent,
			PageSize: resourcesCount - 1,
		})
		assert.NilError(t, err)
		assert.Check(t, response.NextPageToken != "")
	})

	// Listing resource one by one should eventually return all resources.
	t.Run("one by one", func(t *testing.T) {
		fx.maybeSkip(t)
		msgs := make([]*Model, 0, resourcesCount)
		var nextPageToken string
		for {
			response, err := fx.Service().ListModels(fx.Context(), &ListModelsRequest{
				Parent:    parent,
				PageSize:  1,
				PageToken: nextPageToken,
			})
			assert.NilError(t, err)
			assert.Equal(t, 1, len(response.Models))
			msgs = append(msgs, response.Models...)
			nextPageToken = response.NextPageToken
			if nextPageToken == "" {
				break
			}
		}
		assert.DeepEqual(
			t,
			parentMsgs,
			msgs,
			cmpopts.SortSlices(func(a, b *Model) bool {
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
			_, err := fx.Service().DeleteModel(fx.Context(), &DeleteModelRequest{
				Name: parentMsgs[i].Name,
			})
			assert.NilError(t, err)
		}
		response, err := fx.Service().ListModels(fx.Context(), &ListModelsRequest{
			Parent:   parent,
			PageSize: 9999,
		})
		assert.NilError(t, err)
		assert.DeepEqual(
			t,
			parentMsgs[deleteCount:],
			response.Models,
			cmpopts.SortSlices(func(a, b *Model) bool {
				return a.Name < b.Name
			}),
			protocmp.Transform(),
		)
	})

}

func (fx *ModelServiceModelTestSuiteConfig) testDelete(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().DeleteModel(fx.Context(), &DeleteModelRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().DeleteModel(fx.Context(), &DeleteModelRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Resource should be deleted without errors if it exists.
	t.Run("exists", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		_, err := fx.Service().DeleteModel(fx.Context(), &DeleteModelRequest{
			Name: created.Name,
		})
		assert.NilError(t, err)
	})

	// Method should fail with NotFound if the resource does not exist.
	t.Run("not found", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		_, err := fx.Service().DeleteModel(fx.Context(), &DeleteModelRequest{
			Name: created.Name + "notfound",
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with NotFound if the resource was already deleted. This also applies to soft-deletion.
	t.Run("already deleted", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		deleted, err := fx.Service().DeleteModel(fx.Context(), &DeleteModelRequest{
			Name: created.Name,
		})
		assert.NilError(t, err)
		_ = deleted
		_, err = fx.Service().DeleteModel(fx.Context(), &DeleteModelRequest{
			Name: created.Name,
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name only contains wildcards ('-')
	t.Run("only wildcards", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().DeleteModel(fx.Context(), &DeleteModelRequest{
			Name: "projects/-/locations/-/models/-",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *ModelServiceModelTestSuiteConfig) nextParent(t *testing.T, pristine bool) string {
	if pristine {
		fx.currParent++
	}
	if fx.currParent >= len(fx.Parents) {
		t.Fatal("need at least", fx.currParent+1, "parents")
	}
	return fx.Parents[fx.currParent]
}

func (fx *ModelServiceModelTestSuiteConfig) peekNextParent(t *testing.T) string {
	next := fx.currParent + 1
	if next >= len(fx.Parents) {
		t.Fatal("need at least", next+1, "parents")
	}
	return fx.Parents[next]
}

func (fx *ModelServiceModelTestSuiteConfig) maybeSkip(t *testing.T) {
	for _, skip := range fx.Skip {
		if strings.Contains(t.Name(), skip) {
			t.Skip("skipped because of .Skip")
		}
	}
}

func (fx *ModelServiceModelTestSuiteConfig) create(t *testing.T, parent string) *Model {
	t.Helper()
	if fx.CreateResource == nil {
		t.Skip("Test skipped because CreateResource not specified on ModelServiceModelTestSuiteConfig")
	}
	created, err := fx.CreateResource(fx.Context(), parent)
	assert.NilError(t, err)
	return created
}

type ModelServiceModelEvaluationTestSuiteConfig struct {
	currParent int

	// Service should return the service that should be tested.
	// The service will be used for several tests.
	Service func() ModelServiceServer
	// Context should return a new context.
	// The context will be used for several tests.
	Context func() context.Context
	// The parents to use when creating resources.
	// At least one parent needs to be set. Depending on methods available on the resource,
	// more may be required. If insufficient number of parents are
	// provided the test will fail.
	Parents []string
	// CreateResource should create a ModelEvaluation and return it.
	// If the field is not set, some tests will be skipped.
	//
	// This method is generated because service does not expose a Create
	// method (or it does not comply with AIP).
	CreateResource func(ctx context.Context, parent string) (*ModelEvaluation, error)
	// Patterns of tests to skip.
	// For example if a service has a Get method:
	// Skip: ["Get"] will skip all tests for Get.
	// Skip: ["Get/persisted"] will only skip the subtest called "persisted" of Get.
	Skip []string
}

func (fx *ModelServiceModelEvaluationTestSuiteConfig) test(t *testing.T) {
	t.Run("Get", fx.testGet)
	t.Run("List", fx.testList)
}

func (fx *ModelServiceModelEvaluationTestSuiteConfig) testGet(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().GetModelEvaluation(fx.Context(), &GetModelEvaluationRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().GetModelEvaluation(fx.Context(), &GetModelEvaluationRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Resource should be returned without errors if it exists.
	t.Run("exists", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		msg, err := fx.Service().GetModelEvaluation(fx.Context(), &GetModelEvaluationRequest{
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
		_, err := fx.Service().GetModelEvaluation(fx.Context(), &GetModelEvaluationRequest{
			Name: created.Name + "notfound",
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name only contains wildcards ('-')
	t.Run("only wildcards", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().GetModelEvaluation(fx.Context(), &GetModelEvaluationRequest{
			Name: "projects/-/locations/-/models/-/evaluations/-",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *ModelServiceModelEvaluationTestSuiteConfig) testList(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().ListModelEvaluations(fx.Context(), &ListModelEvaluationsRequest{
			Parent: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page token is not valid.
	t.Run("invalid page token", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.Service().ListModelEvaluations(fx.Context(), &ListModelEvaluationsRequest{
			Parent:    parent,
			PageToken: "invalid page token",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page size is negative.
	t.Run("negative page size", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.Service().ListModelEvaluations(fx.Context(), &ListModelEvaluationsRequest{
			Parent:   parent,
			PageSize: -10,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	const resourcesCount = 15
	parent := fx.nextParent(t, true)
	parentMsgs := make([]*ModelEvaluation, resourcesCount)
	for i := 0; i < resourcesCount; i++ {
		parentMsgs[i] = fx.create(t, parent)
	}

	// If parent is provided the method must only return resources
	// under that parent.
	t.Run("isolation", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.Service().ListModelEvaluations(fx.Context(), &ListModelEvaluationsRequest{
			Parent:   parent,
			PageSize: 999,
		})
		assert.NilError(t, err)
		assert.DeepEqual(
			t,
			parentMsgs,
			response.ModelEvaluations,
			cmpopts.SortSlices(func(a, b *ModelEvaluation) bool {
				return a.Name < b.Name
			}),
			protocmp.Transform(),
		)
	})

	// If there are no more resources, next_page_token should not be set.
	t.Run("last page", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.Service().ListModelEvaluations(fx.Context(), &ListModelEvaluationsRequest{
			Parent:   parent,
			PageSize: resourcesCount,
		})
		assert.NilError(t, err)
		assert.Equal(t, "", response.NextPageToken)
	})

	// If there are more resources, next_page_token should be set.
	t.Run("more pages", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.Service().ListModelEvaluations(fx.Context(), &ListModelEvaluationsRequest{
			Parent:   parent,
			PageSize: resourcesCount - 1,
		})
		assert.NilError(t, err)
		assert.Check(t, response.NextPageToken != "")
	})

	// Listing resource one by one should eventually return all resources.
	t.Run("one by one", func(t *testing.T) {
		fx.maybeSkip(t)
		msgs := make([]*ModelEvaluation, 0, resourcesCount)
		var nextPageToken string
		for {
			response, err := fx.Service().ListModelEvaluations(fx.Context(), &ListModelEvaluationsRequest{
				Parent:    parent,
				PageSize:  1,
				PageToken: nextPageToken,
			})
			assert.NilError(t, err)
			assert.Equal(t, 1, len(response.ModelEvaluations))
			msgs = append(msgs, response.ModelEvaluations...)
			nextPageToken = response.NextPageToken
			if nextPageToken == "" {
				break
			}
		}
		assert.DeepEqual(
			t,
			parentMsgs,
			msgs,
			cmpopts.SortSlices(func(a, b *ModelEvaluation) bool {
				return a.Name < b.Name
			}),
			protocmp.Transform(),
		)
	})

}

func (fx *ModelServiceModelEvaluationTestSuiteConfig) nextParent(t *testing.T, pristine bool) string {
	if pristine {
		fx.currParent++
	}
	if fx.currParent >= len(fx.Parents) {
		t.Fatal("need at least", fx.currParent+1, "parents")
	}
	return fx.Parents[fx.currParent]
}

func (fx *ModelServiceModelEvaluationTestSuiteConfig) peekNextParent(t *testing.T) string {
	next := fx.currParent + 1
	if next >= len(fx.Parents) {
		t.Fatal("need at least", next+1, "parents")
	}
	return fx.Parents[next]
}

func (fx *ModelServiceModelEvaluationTestSuiteConfig) maybeSkip(t *testing.T) {
	for _, skip := range fx.Skip {
		if strings.Contains(t.Name(), skip) {
			t.Skip("skipped because of .Skip")
		}
	}
}

func (fx *ModelServiceModelEvaluationTestSuiteConfig) create(t *testing.T, parent string) *ModelEvaluation {
	t.Helper()
	if fx.CreateResource == nil {
		t.Skip("Test skipped because CreateResource not specified on ModelServiceModelEvaluationTestSuiteConfig")
	}
	created, err := fx.CreateResource(fx.Context(), parent)
	assert.NilError(t, err)
	return created
}

type ModelServiceModelEvaluationSliceTestSuiteConfig struct {
	currParent int

	// Service should return the service that should be tested.
	// The service will be used for several tests.
	Service func() ModelServiceServer
	// Context should return a new context.
	// The context will be used for several tests.
	Context func() context.Context
	// The parents to use when creating resources.
	// At least one parent needs to be set. Depending on methods available on the resource,
	// more may be required. If insufficient number of parents are
	// provided the test will fail.
	Parents []string
	// CreateResource should create a ModelEvaluationSlice and return it.
	// If the field is not set, some tests will be skipped.
	//
	// This method is generated because service does not expose a Create
	// method (or it does not comply with AIP).
	CreateResource func(ctx context.Context, parent string) (*ModelEvaluationSlice, error)
	// Patterns of tests to skip.
	// For example if a service has a Get method:
	// Skip: ["Get"] will skip all tests for Get.
	// Skip: ["Get/persisted"] will only skip the subtest called "persisted" of Get.
	Skip []string
}

func (fx *ModelServiceModelEvaluationSliceTestSuiteConfig) test(t *testing.T) {
	t.Run("Get", fx.testGet)
	t.Run("List", fx.testList)
}

func (fx *ModelServiceModelEvaluationSliceTestSuiteConfig) testGet(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().GetModelEvaluationSlice(fx.Context(), &GetModelEvaluationSliceRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().GetModelEvaluationSlice(fx.Context(), &GetModelEvaluationSliceRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Resource should be returned without errors if it exists.
	t.Run("exists", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		msg, err := fx.Service().GetModelEvaluationSlice(fx.Context(), &GetModelEvaluationSliceRequest{
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
		_, err := fx.Service().GetModelEvaluationSlice(fx.Context(), &GetModelEvaluationSliceRequest{
			Name: created.Name + "notfound",
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name only contains wildcards ('-')
	t.Run("only wildcards", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().GetModelEvaluationSlice(fx.Context(), &GetModelEvaluationSliceRequest{
			Name: "projects/-/locations/-/models/-/evaluations/-/slices/-",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *ModelServiceModelEvaluationSliceTestSuiteConfig) testList(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().ListModelEvaluationSlices(fx.Context(), &ListModelEvaluationSlicesRequest{
			Parent: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page token is not valid.
	t.Run("invalid page token", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.Service().ListModelEvaluationSlices(fx.Context(), &ListModelEvaluationSlicesRequest{
			Parent:    parent,
			PageToken: "invalid page token",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page size is negative.
	t.Run("negative page size", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.Service().ListModelEvaluationSlices(fx.Context(), &ListModelEvaluationSlicesRequest{
			Parent:   parent,
			PageSize: -10,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	const resourcesCount = 15
	parent := fx.nextParent(t, true)
	parentMsgs := make([]*ModelEvaluationSlice, resourcesCount)
	for i := 0; i < resourcesCount; i++ {
		parentMsgs[i] = fx.create(t, parent)
	}

	// If parent is provided the method must only return resources
	// under that parent.
	t.Run("isolation", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.Service().ListModelEvaluationSlices(fx.Context(), &ListModelEvaluationSlicesRequest{
			Parent:   parent,
			PageSize: 999,
		})
		assert.NilError(t, err)
		assert.DeepEqual(
			t,
			parentMsgs,
			response.ModelEvaluationSlices,
			cmpopts.SortSlices(func(a, b *ModelEvaluationSlice) bool {
				return a.Name < b.Name
			}),
			protocmp.Transform(),
		)
	})

	// If there are no more resources, next_page_token should not be set.
	t.Run("last page", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.Service().ListModelEvaluationSlices(fx.Context(), &ListModelEvaluationSlicesRequest{
			Parent:   parent,
			PageSize: resourcesCount,
		})
		assert.NilError(t, err)
		assert.Equal(t, "", response.NextPageToken)
	})

	// If there are more resources, next_page_token should be set.
	t.Run("more pages", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.Service().ListModelEvaluationSlices(fx.Context(), &ListModelEvaluationSlicesRequest{
			Parent:   parent,
			PageSize: resourcesCount - 1,
		})
		assert.NilError(t, err)
		assert.Check(t, response.NextPageToken != "")
	})

	// Listing resource one by one should eventually return all resources.
	t.Run("one by one", func(t *testing.T) {
		fx.maybeSkip(t)
		msgs := make([]*ModelEvaluationSlice, 0, resourcesCount)
		var nextPageToken string
		for {
			response, err := fx.Service().ListModelEvaluationSlices(fx.Context(), &ListModelEvaluationSlicesRequest{
				Parent:    parent,
				PageSize:  1,
				PageToken: nextPageToken,
			})
			assert.NilError(t, err)
			assert.Equal(t, 1, len(response.ModelEvaluationSlices))
			msgs = append(msgs, response.ModelEvaluationSlices...)
			nextPageToken = response.NextPageToken
			if nextPageToken == "" {
				break
			}
		}
		assert.DeepEqual(
			t,
			parentMsgs,
			msgs,
			cmpopts.SortSlices(func(a, b *ModelEvaluationSlice) bool {
				return a.Name < b.Name
			}),
			protocmp.Transform(),
		)
	})

}

func (fx *ModelServiceModelEvaluationSliceTestSuiteConfig) nextParent(t *testing.T, pristine bool) string {
	if pristine {
		fx.currParent++
	}
	if fx.currParent >= len(fx.Parents) {
		t.Fatal("need at least", fx.currParent+1, "parents")
	}
	return fx.Parents[fx.currParent]
}

func (fx *ModelServiceModelEvaluationSliceTestSuiteConfig) peekNextParent(t *testing.T) string {
	next := fx.currParent + 1
	if next >= len(fx.Parents) {
		t.Fatal("need at least", next+1, "parents")
	}
	return fx.Parents[next]
}

func (fx *ModelServiceModelEvaluationSliceTestSuiteConfig) maybeSkip(t *testing.T) {
	for _, skip := range fx.Skip {
		if strings.Contains(t.Name(), skip) {
			t.Skip("skipped because of .Skip")
		}
	}
}

func (fx *ModelServiceModelEvaluationSliceTestSuiteConfig) create(t *testing.T, parent string) *ModelEvaluationSlice {
	t.Helper()
	if fx.CreateResource == nil {
		t.Skip("Test skipped because CreateResource not specified on ModelServiceModelEvaluationSliceTestSuiteConfig")
	}
	created, err := fx.CreateResource(fx.Context(), parent)
	assert.NilError(t, err)
	return created
}
