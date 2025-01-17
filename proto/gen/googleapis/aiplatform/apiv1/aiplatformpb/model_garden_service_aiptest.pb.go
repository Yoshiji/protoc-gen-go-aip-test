// Code generated by protoc-gen-go-aip-test. DO NOT EDIT.

package aiplatformpb

import (
	context "context"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protocmp "google.golang.org/protobuf/testing/protocmp"
	assert "gotest.tools/v3/assert"
	strings "strings"
	testing "testing"
)

// ModelGardenServiceTestSuiteConfigProvider is the interface to implement to decide which resources
// that should be tested and how it's configured.
type ModelGardenServiceTestSuiteConfigProvider interface {
	// ModelGardenServicePublisherModel should return a config, or nil, which means that the tests will be skipped.
	ModelGardenServicePublisherModel(t *testing.T) *ModelGardenServicePublisherModelTestSuiteConfig
}

// testModelGardenService is the main entrypoint for starting the AIP tests.
func testModelGardenService(t *testing.T, s ModelGardenServiceTestSuiteConfigProvider) {
	testModelGardenServicePublisherModel(t, s)
}

func testModelGardenServicePublisherModel(t *testing.T, s ModelGardenServiceTestSuiteConfigProvider) {
	t.Run("PublisherModel", func(t *testing.T) {
		config := s.ModelGardenServicePublisherModel(t)
		if config == nil {
			t.Skip("Method ModelGardenServicePublisherModel not implemented")
		}
		if config.Service == nil {
			t.Skip("Method ModelGardenServicePublisherModel.Service() not implemented")
		}
		if config.Context == nil {
			config.Context = func() context.Context { return context.Background() }
		}
		config.test(t)
	})
}

type ModelGardenServiceTestSuite struct {
	T *testing.T
	// Server to test.
	Server ModelGardenServiceServer
}

func (fx ModelGardenServiceTestSuite) TestPublisherModel(ctx context.Context, options ModelGardenServicePublisherModelTestSuiteConfig) {
	fx.T.Run("PublisherModel", func(t *testing.T) {
		options.Context = func() context.Context { return ctx }
		options.Service = func() ModelGardenServiceServer { return fx.Server }
		options.test(t)
	})
}

type ModelGardenServicePublisherModelTestSuiteConfig struct {
	currParent int

	// Service should return the service that should be tested.
	// The service will be used for several tests.
	Service func() ModelGardenServiceServer
	// Context should return a new context.
	// The context will be used for several tests.
	Context func() context.Context
	// The parents to use when creating resources.
	// At least one parent needs to be set. Depending on methods available on the resource,
	// more may be required. If insufficient number of parents are
	// provided the test will fail.
	Parents []string
	// CreateResource should create a PublisherModel and return it.
	// If the field is not set, some tests will be skipped.
	//
	// This method is generated because service does not expose a Create
	// method (or it does not comply with AIP).
	CreateResource func(ctx context.Context, parent string) (*PublisherModel, error)
	// Patterns of tests to skip.
	// For example if a service has a Get method:
	// Skip: ["Get"] will skip all tests for Get.
	// Skip: ["Get/persisted"] will only skip the subtest called "persisted" of Get.
	Skip []string
}

func (fx *ModelGardenServicePublisherModelTestSuiteConfig) test(t *testing.T) {
	t.Run("Get", fx.testGet)
}

func (fx *ModelGardenServicePublisherModelTestSuiteConfig) testGet(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().GetPublisherModel(fx.Context(), &GetPublisherModelRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().GetPublisherModel(fx.Context(), &GetPublisherModelRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Resource should be returned without errors if it exists.
	t.Run("exists", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		msg, err := fx.Service().GetPublisherModel(fx.Context(), &GetPublisherModelRequest{
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
		_, err := fx.Service().GetPublisherModel(fx.Context(), &GetPublisherModelRequest{
			Name: created.Name + "notfound",
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name only contains wildcards ('-')
	t.Run("only wildcards", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().GetPublisherModel(fx.Context(), &GetPublisherModelRequest{
			Name: "publishers/-/models/-",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *ModelGardenServicePublisherModelTestSuiteConfig) nextParent(t *testing.T, pristine bool) string {
	if pristine {
		fx.currParent++
	}
	if fx.currParent >= len(fx.Parents) {
		t.Fatal("need at least", fx.currParent+1, "parents")
	}
	return fx.Parents[fx.currParent]
}

func (fx *ModelGardenServicePublisherModelTestSuiteConfig) peekNextParent(t *testing.T) string {
	next := fx.currParent + 1
	if next >= len(fx.Parents) {
		t.Fatal("need at least", next+1, "parents")
	}
	return fx.Parents[next]
}

func (fx *ModelGardenServicePublisherModelTestSuiteConfig) maybeSkip(t *testing.T) {
	for _, skip := range fx.Skip {
		if strings.Contains(t.Name(), skip) {
			t.Skip("skipped because of .Skip")
		}
	}
}

func (fx *ModelGardenServicePublisherModelTestSuiteConfig) create(t *testing.T, parent string) *PublisherModel {
	t.Helper()
	if fx.CreateResource == nil {
		t.Skip("Test skipped because CreateResource not specified on ModelGardenServicePublisherModelTestSuiteConfig")
	}
	created, err := fx.CreateResource(fx.Context(), parent)
	assert.NilError(t, err)
	return created
}
