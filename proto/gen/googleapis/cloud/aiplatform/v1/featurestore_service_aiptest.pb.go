// Code generated by protoc-gen-go-aip-test. DO NOT EDIT.

package aiplatform

import (
	context "context"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	assert "gotest.tools/v3/assert"
	strings "strings"
	testing "testing"
)

type FeaturestoreServiceTestSuite struct {
	T *testing.T
	// Server to test.
	Server FeaturestoreServiceServer
}

func (fx FeaturestoreServiceTestSuite) TestEntityType(ctx context.Context, options EntityTypeTestSuiteConfig) {
	fx.T.Run("EntityType", func(t *testing.T) {
		options.ctx = ctx
		options.service = fx.Server
		options.test(t)
	})
}

func (fx FeaturestoreServiceTestSuite) TestFeature(ctx context.Context, options FeatureTestSuiteConfig) {
	fx.T.Run("Feature", func(t *testing.T) {
		options.ctx = ctx
		options.service = fx.Server
		options.test(t)
	})
}

func (fx FeaturestoreServiceTestSuite) TestFeaturestore(ctx context.Context, options FeaturestoreTestSuiteConfig) {
	fx.T.Run("Featurestore", func(t *testing.T) {
		options.ctx = ctx
		options.service = fx.Server
		options.test(t)
	})
}

type EntityTypeTestSuiteConfig struct {
	ctx        context.Context
	service    FeaturestoreServiceServer
	currParent int

	// The parents to use when creating resources.
	// At least one parent needs to be set. Depending on methods available on the resource,
	// more may be required. If insufficient number of parents are
	// provided the test will fail.
	Parents []string
	// Create should return a resource which is valid to create, i.e.
	// all required fields set.
	Create func(parent string) *EntityType
	// Update should return a resource which is valid to update, i.e.
	// all required fields set.
	Update func(parent string) *EntityType
	// Patterns of tests to skip.
	// For example if a service has a Get method:
	// Skip: ["Get"] will skip all tests for Get.
	// Skip: ["Get/persisted"] will only skip the subtest called "persisted" of Get.
	Skip []string
}

func (fx *EntityTypeTestSuiteConfig) test(t *testing.T) {
	t.Run("Create", fx.testCreate)
	t.Run("Get", fx.testGet)
	t.Run("Update", fx.testUpdate)
	t.Run("List", fx.testList)
}

func (fx *EntityTypeTestSuiteConfig) testCreate(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no parent is provided.
	t.Run("missing parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.CreateEntityType(fx.ctx, &CreateEntityTypeRequest{
			Parent:     "",
			EntityType: fx.Create(fx.nextParent(t, false)),
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.CreateEntityType(fx.ctx, &CreateEntityTypeRequest{
			Parent:     "invalid resource name",
			EntityType: fx.Create(fx.nextParent(t, false)),
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *EntityTypeTestSuiteConfig) testGet(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetEntityType(fx.ctx, &GetEntityTypeRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetEntityType(fx.ctx, &GetEntityTypeRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *EntityTypeTestSuiteConfig) testUpdate(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg := fx.Update(parent)
		msg.Name = ""
		_, err := fx.service.UpdateEntityType(fx.ctx, &UpdateEntityTypeRequest{
			EntityType: msg,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg := fx.Update(parent)
		msg.Name = "invalid resource name"
		_, err := fx.service.UpdateEntityType(fx.ctx, &UpdateEntityTypeRequest{
			EntityType: msg,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *EntityTypeTestSuiteConfig) testList(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.ListEntityTypes(fx.ctx, &ListEntityTypesRequest{
			Parent: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page token is not valid.
	t.Run("invalid page token", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.service.ListEntityTypes(fx.ctx, &ListEntityTypesRequest{
			Parent:    parent,
			PageToken: "invalid page token",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page size is negative.
	t.Run("negative page size", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.service.ListEntityTypes(fx.ctx, &ListEntityTypesRequest{
			Parent:   parent,
			PageSize: -10,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *EntityTypeTestSuiteConfig) nextParent(t *testing.T, pristine bool) string {
	if pristine {
		fx.currParent++
	}
	if fx.currParent >= len(fx.Parents) {
		t.Fatal("need at least", fx.currParent+1, "parents")
	}
	return fx.Parents[fx.currParent]
}

func (fx *EntityTypeTestSuiteConfig) peekNextParent(t *testing.T) string {
	next := fx.currParent + 1
	if next >= len(fx.Parents) {
		t.Fatal("need at least", next+1, "parents")
	}
	return fx.Parents[next]
}

func (fx *EntityTypeTestSuiteConfig) maybeSkip(t *testing.T) {
	for _, skip := range fx.Skip {
		if strings.Contains(t.Name(), skip) {
			t.Skip("skipped because of .Skip")
		}
	}
}

type FeatureTestSuiteConfig struct {
	ctx        context.Context
	service    FeaturestoreServiceServer
	currParent int

	// The parents to use when creating resources.
	// At least one parent needs to be set. Depending on methods available on the resource,
	// more may be required. If insufficient number of parents are
	// provided the test will fail.
	Parents []string
	// Create should return a resource which is valid to create, i.e.
	// all required fields set.
	Create func(parent string) *Feature
	// Update should return a resource which is valid to update, i.e.
	// all required fields set.
	Update func(parent string) *Feature
	// Patterns of tests to skip.
	// For example if a service has a Get method:
	// Skip: ["Get"] will skip all tests for Get.
	// Skip: ["Get/persisted"] will only skip the subtest called "persisted" of Get.
	Skip []string
}

func (fx *FeatureTestSuiteConfig) test(t *testing.T) {
	t.Run("Create", fx.testCreate)
	t.Run("Get", fx.testGet)
	t.Run("Update", fx.testUpdate)
	t.Run("List", fx.testList)
}

func (fx *FeatureTestSuiteConfig) testCreate(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no parent is provided.
	t.Run("missing parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.CreateFeature(fx.ctx, &CreateFeatureRequest{
			Parent:  "",
			Feature: fx.Create(fx.nextParent(t, false)),
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.CreateFeature(fx.ctx, &CreateFeatureRequest{
			Parent:  "invalid resource name",
			Feature: fx.Create(fx.nextParent(t, false)),
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// The method should fail with InvalidArgument if the resource has any
	// required fields and they are not provided.
	t.Run("required fields", func(t *testing.T) {
		fx.maybeSkip(t)
		t.Run(".value_type", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("value_type")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateFeature(fx.ctx, &CreateFeatureRequest{
				Parent:  parent,
				Feature: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
	})

}

func (fx *FeatureTestSuiteConfig) testGet(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetFeature(fx.ctx, &GetFeatureRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetFeature(fx.ctx, &GetFeatureRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *FeatureTestSuiteConfig) testUpdate(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg := fx.Update(parent)
		msg.Name = ""
		_, err := fx.service.UpdateFeature(fx.ctx, &UpdateFeatureRequest{
			Feature: msg,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg := fx.Update(parent)
		msg.Name = "invalid resource name"
		_, err := fx.service.UpdateFeature(fx.ctx, &UpdateFeatureRequest{
			Feature: msg,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *FeatureTestSuiteConfig) testList(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.ListFeatures(fx.ctx, &ListFeaturesRequest{
			Parent: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page token is not valid.
	t.Run("invalid page token", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.service.ListFeatures(fx.ctx, &ListFeaturesRequest{
			Parent:    parent,
			PageToken: "invalid page token",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page size is negative.
	t.Run("negative page size", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.service.ListFeatures(fx.ctx, &ListFeaturesRequest{
			Parent:   parent,
			PageSize: -10,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *FeatureTestSuiteConfig) nextParent(t *testing.T, pristine bool) string {
	if pristine {
		fx.currParent++
	}
	if fx.currParent >= len(fx.Parents) {
		t.Fatal("need at least", fx.currParent+1, "parents")
	}
	return fx.Parents[fx.currParent]
}

func (fx *FeatureTestSuiteConfig) peekNextParent(t *testing.T) string {
	next := fx.currParent + 1
	if next >= len(fx.Parents) {
		t.Fatal("need at least", next+1, "parents")
	}
	return fx.Parents[next]
}

func (fx *FeatureTestSuiteConfig) maybeSkip(t *testing.T) {
	for _, skip := range fx.Skip {
		if strings.Contains(t.Name(), skip) {
			t.Skip("skipped because of .Skip")
		}
	}
}

type FeaturestoreTestSuiteConfig struct {
	ctx        context.Context
	service    FeaturestoreServiceServer
	currParent int

	// The parents to use when creating resources.
	// At least one parent needs to be set. Depending on methods available on the resource,
	// more may be required. If insufficient number of parents are
	// provided the test will fail.
	Parents []string
	// Create should return a resource which is valid to create, i.e.
	// all required fields set.
	Create func(parent string) *Featurestore
	// Update should return a resource which is valid to update, i.e.
	// all required fields set.
	Update func(parent string) *Featurestore
	// Patterns of tests to skip.
	// For example if a service has a Get method:
	// Skip: ["Get"] will skip all tests for Get.
	// Skip: ["Get/persisted"] will only skip the subtest called "persisted" of Get.
	Skip []string
}

func (fx *FeaturestoreTestSuiteConfig) test(t *testing.T) {
	t.Run("Create", fx.testCreate)
	t.Run("Get", fx.testGet)
	t.Run("Update", fx.testUpdate)
	t.Run("List", fx.testList)
}

func (fx *FeaturestoreTestSuiteConfig) testCreate(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no parent is provided.
	t.Run("missing parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.CreateFeaturestore(fx.ctx, &CreateFeaturestoreRequest{
			Parent:       "",
			Featurestore: fx.Create(fx.nextParent(t, false)),
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.CreateFeaturestore(fx.ctx, &CreateFeaturestoreRequest{
			Parent:       "invalid resource name",
			Featurestore: fx.Create(fx.nextParent(t, false)),
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// The method should fail with InvalidArgument if the resource has any
	// required fields and they are not provided.
	t.Run("required fields", func(t *testing.T) {
		fx.maybeSkip(t)
		t.Run(".online_serving_config", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("online_serving_config")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateFeaturestore(fx.ctx, &CreateFeaturestoreRequest{
				Parent:       parent,
				Featurestore: msg,
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
			_, err := fx.service.CreateFeaturestore(fx.ctx, &CreateFeaturestoreRequest{
				Parent:       parent,
				Featurestore: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
	})

}

func (fx *FeaturestoreTestSuiteConfig) testGet(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetFeaturestore(fx.ctx, &GetFeaturestoreRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetFeaturestore(fx.ctx, &GetFeaturestoreRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *FeaturestoreTestSuiteConfig) testUpdate(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg := fx.Update(parent)
		msg.Name = ""
		_, err := fx.service.UpdateFeaturestore(fx.ctx, &UpdateFeaturestoreRequest{
			Featurestore: msg,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg := fx.Update(parent)
		msg.Name = "invalid resource name"
		_, err := fx.service.UpdateFeaturestore(fx.ctx, &UpdateFeaturestoreRequest{
			Featurestore: msg,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *FeaturestoreTestSuiteConfig) testList(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.ListFeaturestores(fx.ctx, &ListFeaturestoresRequest{
			Parent: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page token is not valid.
	t.Run("invalid page token", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.service.ListFeaturestores(fx.ctx, &ListFeaturestoresRequest{
			Parent:    parent,
			PageToken: "invalid page token",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page size is negative.
	t.Run("negative page size", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.service.ListFeaturestores(fx.ctx, &ListFeaturestoresRequest{
			Parent:   parent,
			PageSize: -10,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *FeaturestoreTestSuiteConfig) nextParent(t *testing.T, pristine bool) string {
	if pristine {
		fx.currParent++
	}
	if fx.currParent >= len(fx.Parents) {
		t.Fatal("need at least", fx.currParent+1, "parents")
	}
	return fx.Parents[fx.currParent]
}

func (fx *FeaturestoreTestSuiteConfig) peekNextParent(t *testing.T) string {
	next := fx.currParent + 1
	if next >= len(fx.Parents) {
		t.Fatal("need at least", next+1, "parents")
	}
	return fx.Parents[next]
}

func (fx *FeaturestoreTestSuiteConfig) maybeSkip(t *testing.T) {
	for _, skip := range fx.Skip {
		if strings.Contains(t.Name(), skip) {
			t.Skip("skipped because of .Skip")
		}
	}
}
