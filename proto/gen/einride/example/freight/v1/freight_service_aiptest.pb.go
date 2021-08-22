// Code generated by protoc-gen-go-aip-test. DO NOT EDIT.

package examplefreightv1

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

type FreightServiceTestSuite struct {
	T *testing.T
	// Server to test.
	Server FreightServiceServer
}

func (fx FreightServiceTestSuite) TestShipper(ctx context.Context, options ShipperTestSuiteConfig) {
	fx.T.Run("Shipper", func(t *testing.T) {
		options.ctx = ctx
		options.service = fx.Server
		options.test(t)
	})
}

func (fx FreightServiceTestSuite) TestSite(ctx context.Context, options SiteTestSuiteConfig) {
	fx.T.Run("Site", func(t *testing.T) {
		options.ctx = ctx
		options.service = fx.Server
		options.test(t)
	})
}

type ShipperTestSuiteConfig struct {
	ctx        context.Context
	service    FreightServiceServer
	currParent int

	// Create should return a resource which is valid to create, i.e.
	// all required fields set.
	Create func() *Shipper
	// Update should return a resource which is valid to update, i.e.
	// all required fields set.
	Update func() *Shipper
	// Patterns of tests to skip.
	// For example if a service has a Get method:
	// Skip: ["Get"] will skip all tests for Get.
	// Skip: ["Get/persisted"] will only skip the subtest called "persisted" of Get.
	Skip []string
}

func (fx *ShipperTestSuiteConfig) test(t *testing.T) {
	t.Run("Create", fx.testCreate)
	t.Run("Get", fx.testGet)
	t.Run("Update", fx.testUpdate)
	t.Run("List", fx.testList)
}

func (fx *ShipperTestSuiteConfig) testCreate(t *testing.T) {

	// Field create_time should be populated when the resource is created.
	t.Run("create time", func(t *testing.T) {
		fx.maybeSkip(t)
		msg, err := fx.service.CreateShipper(fx.ctx, &CreateShipperRequest{
			Shipper: fx.Create(),
		})
		assert.NilError(t, err)
		assert.Check(t, time.Since(msg.CreateTime.AsTime()) < time.Second)
	})

	// The created resource should be persisted and reachable with Get.
	t.Run("persisted", func(t *testing.T) {
		fx.maybeSkip(t)
		msg, err := fx.service.CreateShipper(fx.ctx, &CreateShipperRequest{
			Shipper: fx.Create(),
		})
		assert.NilError(t, err)
		persisted, err := fx.service.GetShipper(fx.ctx, &GetShipperRequest{
			Name: msg.Name,
		})
		assert.NilError(t, err)
		assert.DeepEqual(t, msg, persisted, protocmp.Transform())
	})

	// If method support user settable IDs, when set the resource should
	// be returned with the provided ID.
	t.Run("user settable id", func(t *testing.T) {
		fx.maybeSkip(t)
		msg, err := fx.service.CreateShipper(fx.ctx, &CreateShipperRequest{
			Shipper:   fx.Create(),
			ShipperId: "usersetid",
		})
		assert.NilError(t, err)
		assert.Check(t, strings.HasSuffix(msg.GetName(), "usersetid"))
	})

	// If method support user settable IDs and the same ID is reused
	// the method should return AlreadyExists.
	t.Run("already exists", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.CreateShipper(fx.ctx, &CreateShipperRequest{
			Shipper:   fx.Create(),
			ShipperId: "alreadyexists",
		})
		assert.NilError(t, err)
		_, err = fx.service.CreateShipper(fx.ctx, &CreateShipperRequest{
			Shipper:   fx.Create(),
			ShipperId: "alreadyexists",
		})
		assert.Equal(t, codes.AlreadyExists, status.Code(err), err)
	})

	// The method should fail with InvalidArgument if the resource has any
	// required fields and they are not provided.
	t.Run("required fields", func(t *testing.T) {
		fx.maybeSkip(t)
		t.Run(".display_name", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := fx.Create()
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("display_name")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateShipper(fx.ctx, &CreateShipperRequest{
				Shipper: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".billing_account", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := fx.Create()
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("billing_account")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateShipper(fx.ctx, &CreateShipperRequest{
				Shipper: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
	})

	// The method should fail with InvalidArgument if the resource has any
	// resource references and they are invalid.
	t.Run("resource references", func(t *testing.T) {
		fx.maybeSkip(t)
		t.Run(".billing_account", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := fx.Create()
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			container.BillingAccount = "invalid resource name"
			_, err := fx.service.CreateShipper(fx.ctx, &CreateShipperRequest{
				Shipper: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
	})

}

func (fx *ShipperTestSuiteConfig) testGet(t *testing.T) {
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetShipper(fx.ctx, &GetShipperRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetShipper(fx.ctx, &GetShipperRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Resource should be returned without errors if it exists.
	t.Run("exists", func(t *testing.T) {
		fx.maybeSkip(t)
		created, err := fx.service.CreateShipper(fx.ctx, &CreateShipperRequest{
			Shipper: fx.Create(),
		})
		assert.NilError(t, err)
		msg, err := fx.service.GetShipper(fx.ctx, &GetShipperRequest{
			Name: created.Name,
		})
		assert.NilError(t, err)
		assert.DeepEqual(t, msg, created, protocmp.Transform())
	})

	// Method should fail with NotFound if the resource does not exist.
	t.Run("not found", func(t *testing.T) {
		fx.maybeSkip(t)
		created, err := fx.service.CreateShipper(fx.ctx, &CreateShipperRequest{
			Shipper: fx.Create(),
		})
		assert.NilError(t, err)
		_, err = fx.service.GetShipper(fx.ctx, &GetShipperRequest{
			Name: created.Name + "notfound",
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

}

func (fx *ShipperTestSuiteConfig) testUpdate(t *testing.T) {
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		msg := fx.Update()
		msg.Name = ""
		_, err := fx.service.UpdateShipper(fx.ctx, &UpdateShipperRequest{
			Shipper: msg,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		msg := fx.Update()
		msg.Name = "invalid resource name"
		_, err := fx.service.UpdateShipper(fx.ctx, &UpdateShipperRequest{
			Shipper: msg,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Field update_time should be updated when the resource is updated.
	t.Run("update time", func(t *testing.T) {
		fx.maybeSkip(t)
		created, err := fx.service.CreateShipper(fx.ctx, &CreateShipperRequest{
			Shipper: fx.Create(),
		})
		assert.NilError(t, err)
		updated, err := fx.service.UpdateShipper(fx.ctx, &UpdateShipperRequest{
			Shipper: created,
		})
		assert.NilError(t, err)
		assert.Check(t, updated.UpdateTime.AsTime().After(created.UpdateTime.AsTime()))
	})

	// The updated resource should be persisted and reachable with Get.
	t.Run("persisted", func(t *testing.T) {
		fx.maybeSkip(t)
		created, err := fx.service.CreateShipper(fx.ctx, &CreateShipperRequest{
			Shipper: fx.Create(),
		})
		assert.NilError(t, err)
		updated, err := fx.service.UpdateShipper(fx.ctx, &UpdateShipperRequest{
			Shipper: created,
		})
		assert.NilError(t, err)
		persisted, err := fx.service.GetShipper(fx.ctx, &GetShipperRequest{
			Name: updated.Name,
		})
		assert.NilError(t, err)
		assert.DeepEqual(t, updated, persisted, protocmp.Transform())
	})

	created, err := fx.service.CreateShipper(fx.ctx, &CreateShipperRequest{
		Shipper: fx.Create(),
	})
	assert.NilError(t, err)
	// Method should fail with NotFound if the resource does not exist.
	t.Run("not found", func(t *testing.T) {
		fx.maybeSkip(t)
		msg := fx.Update()
		msg.Name = created.Name + "notfound"
		_, err := fx.service.UpdateShipper(fx.ctx, &UpdateShipperRequest{
			Shipper: msg,
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// The method should fail with InvalidArgument if the update_mask is invalid.
	t.Run("invalid update mask", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.UpdateShipper(fx.ctx, &UpdateShipperRequest{
			Shipper: created,
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
			msg := proto.Clone(created).(*Shipper)
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("display_name")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.UpdateShipper(fx.ctx, &UpdateShipperRequest{
				Shipper: msg,
				UpdateMask: &fieldmaskpb.FieldMask{
					Paths: []string{
						"*",
					},
				},
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".billing_account", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := proto.Clone(created).(*Shipper)
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("billing_account")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.UpdateShipper(fx.ctx, &UpdateShipperRequest{
				Shipper: msg,
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

func (fx *ShipperTestSuiteConfig) testList(t *testing.T) {

	// Method should fail with InvalidArgument is provided page token is not valid.
	t.Run("invalid page token", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.ListShippers(fx.ctx, &ListShippersRequest{
			PageToken: "invalid page token",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page size is negative.
	t.Run("negative page size", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.ListShippers(fx.ctx, &ListShippersRequest{
			PageSize: -10,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *ShipperTestSuiteConfig) maybeSkip(t *testing.T) {
	for _, skip := range fx.Skip {
		if strings.Contains(t.Name(), skip) {
			t.Skip("skipped because of .Skip")
		}
	}
}

type SiteTestSuiteConfig struct {
	ctx        context.Context
	service    FreightServiceServer
	currParent int

	// The parents to use when creating resources.
	// At least one parent needs to be set. Depending on methods available on the resource,
	// more may be required. If insufficient number of parents are
	// provided the test will fail.
	Parents []string
	// Create should return a resource which is valid to create, i.e.
	// all required fields set.
	Create func(parent string) *Site
	// Update should return a resource which is valid to update, i.e.
	// all required fields set.
	Update func(parent string) *Site
	// Patterns of tests to skip.
	// For example if a service has a Get method:
	// Skip: ["Get"] will skip all tests for Get.
	// Skip: ["Get/persisted"] will only skip the subtest called "persisted" of Get.
	Skip []string
}

func (fx *SiteTestSuiteConfig) test(t *testing.T) {
	t.Run("Create", fx.testCreate)
	t.Run("Get", fx.testGet)
	t.Run("BatchGet", fx.testBatchGet)
	t.Run("Update", fx.testUpdate)
	t.Run("List", fx.testList)
}

func (fx *SiteTestSuiteConfig) testCreate(t *testing.T) {
	// Method should fail with InvalidArgument if no parent is provided.
	t.Run("missing parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.CreateSite(fx.ctx, &CreateSiteRequest{
			Parent: "",
			Site:   fx.Create(""),
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.CreateSite(fx.ctx, &CreateSiteRequest{
			Parent: "invalid resource name",
			Site:   fx.Create("invalid resource name"),
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Field create_time should be populated when the resource is created.
	t.Run("create time", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg, err := fx.service.CreateSite(fx.ctx, &CreateSiteRequest{
			Parent: parent,
			Site:   fx.Create(parent),
		})
		assert.NilError(t, err)
		assert.Check(t, time.Since(msg.CreateTime.AsTime()) < time.Second)
	})

	// The created resource should be persisted and reachable with Get.
	t.Run("persisted", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg, err := fx.service.CreateSite(fx.ctx, &CreateSiteRequest{
			Parent: parent,
			Site:   fx.Create(parent),
		})
		assert.NilError(t, err)
		persisted, err := fx.service.GetSite(fx.ctx, &GetSiteRequest{
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
			_, err := fx.service.CreateSite(fx.ctx, &CreateSiteRequest{
				Parent: parent,
				Site:   msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
	})

	// The method should fail with InvalidArgument if the resource has any
	// resource references and they are invalid.
	t.Run("resource references", func(t *testing.T) {
		fx.maybeSkip(t)
		t.Run(".billing.billing_account", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetBilling()
			if container == nil {
				t.Skip("not reachable")
			}
			container.BillingAccount = "invalid resource name"
			_, err := fx.service.CreateSite(fx.ctx, &CreateSiteRequest{
				Parent: parent,
				Site:   msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
	})

}

func (fx *SiteTestSuiteConfig) testGet(t *testing.T) {
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetSite(fx.ctx, &GetSiteRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetSite(fx.ctx, &GetSiteRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Resource should be returned without errors if it exists.
	t.Run("exists", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created, err := fx.service.CreateSite(fx.ctx, &CreateSiteRequest{
			Parent: parent,
			Site:   fx.Create(parent),
		})
		assert.NilError(t, err)
		msg, err := fx.service.GetSite(fx.ctx, &GetSiteRequest{
			Name: created.Name,
		})
		assert.NilError(t, err)
		assert.DeepEqual(t, msg, created, protocmp.Transform())
	})

	// Method should fail with NotFound if the resource does not exist.
	t.Run("not found", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created, err := fx.service.CreateSite(fx.ctx, &CreateSiteRequest{
			Parent: parent,
			Site:   fx.Create(parent),
		})
		assert.NilError(t, err)
		_, err = fx.service.GetSite(fx.ctx, &GetSiteRequest{
			Name: created.Name + "notfound",
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

}

func (fx *SiteTestSuiteConfig) testBatchGet(t *testing.T) {
	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.BatchGetSites(fx.ctx, &BatchGetSitesRequest{
			Parent: "invalid resource name",
			Names:  []string{},
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if no names are provided.
	t.Run("names missing", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.service.BatchGetSites(fx.ctx, &BatchGetSitesRequest{
			Parent: parent,
			Names:  []string{},
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if a provided name is not valid.
	t.Run("names missing", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.service.BatchGetSites(fx.ctx, &BatchGetSitesRequest{
			Parent: parent,
			Names: []string{
				"invalid resource name",
			},
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	parent := fx.nextParent(t, false)
	created00, err := fx.service.CreateSite(fx.ctx, &CreateSiteRequest{
		Parent: parent,
		Site:   fx.Create(parent),
	})
	assert.NilError(t, err)
	created01, err := fx.service.CreateSite(fx.ctx, &CreateSiteRequest{
		Parent: parent,
		Site:   fx.Create(parent),
	})
	assert.NilError(t, err)
	created02, err := fx.service.CreateSite(fx.ctx, &CreateSiteRequest{
		Parent: parent,
		Site:   fx.Create(parent),
	})
	assert.NilError(t, err)
	// Resources should be returned without errors if they exist.
	t.Run("all exists", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.service.BatchGetSites(fx.ctx, &BatchGetSitesRequest{
			Parent: parent,
			Names: []string{
				created00.Name,
				created01.Name,
				created02.Name,
			},
		})
		assert.NilError(t, err)
		assert.DeepEqual(
			t,
			[]*Site{
				created00,
				created01,
				created02,
			},
			response.Sites,
			protocmp.Transform(),
		)
	})

	// The method must be atomic; it must fail for all resources
	// or succeed for all resources (no partial success).
	t.Run("atomic", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.BatchGetSites(fx.ctx, &BatchGetSitesRequest{
			Parent: parent,
			Names: []string{
				created00.Name,
				created01.Name + "notfound",
				created02.Name,
			},
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// If a caller sets the "parent", and the parent collection in the name of any resource
	// being retrieved does not match, the request must fail.
	t.Run("parent mismatch", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.BatchGetSites(fx.ctx, &BatchGetSitesRequest{
			Parent: fx.peekNextParent(t),
			Names: []string{
				created00.Name,
			},
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// The order of resources in the response must be the same as the names in the request.
	t.Run("ordered", func(t *testing.T) {
		fx.maybeSkip(t)
		for _, order := range [][]*Site{
			{created00, created01, created02},
			{created01, created00, created02},
			{created02, created01, created00},
		} {
			response, err := fx.service.BatchGetSites(fx.ctx, &BatchGetSitesRequest{
				Parent: parent,
				Names: []string{
					order[0].GetName(),
					order[1].GetName(),
					order[2].GetName(),
				},
			})
			assert.NilError(t, err)
			assert.DeepEqual(t, order, response.Sites, protocmp.Transform())
		}
	})

	// If a caller provides duplicate names, the service should return
	// duplicate resources.
	t.Run("duplicate names", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.service.BatchGetSites(fx.ctx, &BatchGetSitesRequest{
			Parent: parent,
			Names: []string{
				created00.Name,
				created00.Name,
			},
		})
		assert.NilError(t, err)
		assert.DeepEqual(
			t,
			[]*Site{
				created00,
				created00,
			},
			response.Sites,
			protocmp.Transform(),
		)
	})

}

func (fx *SiteTestSuiteConfig) testUpdate(t *testing.T) {
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg := fx.Update(parent)
		msg.Name = ""
		_, err := fx.service.UpdateSite(fx.ctx, &UpdateSiteRequest{
			Site: msg,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg := fx.Update(parent)
		msg.Name = "invalid resource name"
		_, err := fx.service.UpdateSite(fx.ctx, &UpdateSiteRequest{
			Site: msg,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Field update_time should be updated when the resource is updated.
	t.Run("update time", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created, err := fx.service.CreateSite(fx.ctx, &CreateSiteRequest{
			Parent: parent,
			Site:   fx.Create(parent),
		})
		assert.NilError(t, err)
		updated, err := fx.service.UpdateSite(fx.ctx, &UpdateSiteRequest{
			Site: created,
		})
		assert.NilError(t, err)
		assert.Check(t, updated.UpdateTime.AsTime().After(created.UpdateTime.AsTime()))
	})

	// The updated resource should be persisted and reachable with Get.
	t.Run("persisted", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created, err := fx.service.CreateSite(fx.ctx, &CreateSiteRequest{
			Parent: parent,
			Site:   fx.Create(parent),
		})
		assert.NilError(t, err)
		updated, err := fx.service.UpdateSite(fx.ctx, &UpdateSiteRequest{
			Site: created,
		})
		assert.NilError(t, err)
		persisted, err := fx.service.GetSite(fx.ctx, &GetSiteRequest{
			Name: updated.Name,
		})
		assert.NilError(t, err)
		assert.DeepEqual(t, updated, persisted, protocmp.Transform())
	})

	parent := fx.nextParent(t, false)
	created, err := fx.service.CreateSite(fx.ctx, &CreateSiteRequest{
		Parent: parent,
		Site:   fx.Create(parent),
	})
	assert.NilError(t, err)
	// Method should fail with NotFound if the resource does not exist.
	t.Run("not found", func(t *testing.T) {
		fx.maybeSkip(t)
		msg := fx.Update(parent)
		msg.Name = created.Name + "notfound"
		_, err := fx.service.UpdateSite(fx.ctx, &UpdateSiteRequest{
			Site: msg,
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// The method should fail with InvalidArgument if the update_mask is invalid.
	t.Run("invalid update mask", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.UpdateSite(fx.ctx, &UpdateSiteRequest{
			Site: created,
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
			msg := proto.Clone(created).(*Site)
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("display_name")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.UpdateSite(fx.ctx, &UpdateSiteRequest{
				Site: msg,
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

func (fx *SiteTestSuiteConfig) testList(t *testing.T) {
	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.ListSites(fx.ctx, &ListSitesRequest{
			Parent: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page token is not valid.
	t.Run("invalid page token", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.service.ListSites(fx.ctx, &ListSitesRequest{
			Parent:    parent,
			PageToken: "invalid page token",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page size is negative.
	t.Run("negative page size", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.service.ListSites(fx.ctx, &ListSitesRequest{
			Parent:   parent,
			PageSize: -10,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	const resourcesCount = 15
	parent := fx.nextParent(t, true)
	parentMsgs := make([]*Site, resourcesCount)
	for i := 0; i < resourcesCount; i++ {
		msg, err := fx.service.CreateSite(fx.ctx, &CreateSiteRequest{
			Parent: parent,
			Site:   fx.Create(parent),
		})
		assert.NilError(t, err)
		parentMsgs[i] = msg
	}

	// If parent is provided the method must only return resources
	// under that parent.
	t.Run("isolation", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.service.ListSites(fx.ctx, &ListSitesRequest{
			Parent:   parent,
			PageSize: 999,
		})
		assert.NilError(t, err)
		assert.DeepEqual(
			t,
			parentMsgs,
			response.Sites,
			cmpopts.SortSlices(func(a, b *Site) bool {
				return a.Name < b.Name
			}),
			protocmp.Transform(),
		)
	})

	// If there are no more resources, next_page_token should not be set.
	t.Run("last page", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.service.ListSites(fx.ctx, &ListSitesRequest{
			Parent:   parent,
			PageSize: resourcesCount,
		})
		assert.NilError(t, err)
		assert.Equal(t, "", response.NextPageToken)
	})

	// If there are more resources, next_page_token should be set.
	t.Run("more pages", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.service.ListSites(fx.ctx, &ListSitesRequest{
			Parent:   parent,
			PageSize: resourcesCount - 1,
		})
		assert.NilError(t, err)
		assert.Check(t, response.NextPageToken != "")
	})

	// Listing resource one by one should eventually return all resources.
	t.Run("one by one", func(t *testing.T) {
		fx.maybeSkip(t)
		msgs := make([]*Site, 0, resourcesCount)
		var nextPageToken string
		for {
			response, err := fx.service.ListSites(fx.ctx, &ListSitesRequest{
				Parent:    parent,
				PageSize:  1,
				PageToken: nextPageToken,
			})
			assert.NilError(t, err)
			assert.Equal(t, 1, len(response.Sites))
			msgs = append(msgs, response.Sites...)
			nextPageToken = response.NextPageToken
			if nextPageToken == "" {
				break
			}
		}
		assert.DeepEqual(
			t,
			parentMsgs,
			msgs,
			cmpopts.SortSlices(func(a, b *Site) bool {
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
			_, err := fx.service.DeleteSite(fx.ctx, &DeleteSiteRequest{
				Name: parentMsgs[i].Name,
			})
			assert.NilError(t, err)
		}
		response, err := fx.service.ListSites(fx.ctx, &ListSitesRequest{
			Parent:   parent,
			PageSize: 9999,
		})
		assert.NilError(t, err)
		assert.DeepEqual(
			t,
			parentMsgs[deleteCount:],
			response.Sites,
			cmpopts.SortSlices(func(a, b *Site) bool {
				return a.Name < b.Name
			}),
			protocmp.Transform(),
		)
	})

}

func (fx *SiteTestSuiteConfig) nextParent(t *testing.T, pristine bool) string {
	if pristine {
		fx.currParent++
	}
	if fx.currParent >= len(fx.Parents) {
		t.Fatal("need at least", fx.currParent+1, "parents")
	}
	return fx.Parents[fx.currParent]
}

func (fx *SiteTestSuiteConfig) peekNextParent(t *testing.T) string {
	next := fx.currParent + 1
	if next >= len(fx.Parents) {
		t.Fatal("need at least", next+1, "parents")
	}
	return fx.Parents[next]
}

func (fx *SiteTestSuiteConfig) maybeSkip(t *testing.T) {
	for _, skip := range fx.Skip {
		if strings.Contains(t.Name(), skip) {
			t.Skip("skipped because of .Skip")
		}
	}
}