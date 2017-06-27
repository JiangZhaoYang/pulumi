// *** WARNING: this file was generated by the Lumi IDL Compiler (LUMIDL). ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sns

import (
    "errors"

    pbempty "github.com/golang/protobuf/ptypes/empty"
    pbstruct "github.com/golang/protobuf/ptypes/struct"
    "golang.org/x/net/context"

    "github.com/pulumi/lumi/pkg/resource"
    "github.com/pulumi/lumi/pkg/resource/plugin"
    "github.com/pulumi/lumi/pkg/tokens"
    "github.com/pulumi/lumi/pkg/util/contract"
    "github.com/pulumi/lumi/pkg/util/mapper"
    "github.com/pulumi/lumi/sdk/go/pkg/lumirpc"
)

/* RPC stubs for Topic resource provider */

// TopicToken is the type token corresponding to the Topic package type.
const TopicToken = tokens.Type("aws:sns/topic:Topic")

// TopicProviderOps is a pluggable interface for Topic-related management functionality.
type TopicProviderOps interface {
    Check(ctx context.Context, obj *Topic, property string) error
    Create(ctx context.Context, obj *Topic) (resource.ID, error)
    Get(ctx context.Context, id resource.ID) (*Topic, error)
    InspectChange(ctx context.Context,
        id resource.ID, old *Topic, new *Topic, diff *resource.ObjectDiff) ([]string, error)
    Update(ctx context.Context,
        id resource.ID, old *Topic, new *Topic, diff *resource.ObjectDiff) error
    Delete(ctx context.Context, id resource.ID) error
}

// TopicProvider is a dynamic gRPC-based plugin for managing Topic resources.
type TopicProvider struct {
    ops TopicProviderOps
}

// NewTopicProvider allocates a resource provider that delegates to a ops instance.
func NewTopicProvider(ops TopicProviderOps) lumirpc.ResourceProviderServer {
    contract.Assert(ops != nil)
    return &TopicProvider{ops: ops}
}

func (p *TopicProvider) Check(
    ctx context.Context, req *lumirpc.CheckRequest) (*lumirpc.CheckResponse, error) {
    contract.Assert(req.GetType() == string(TopicToken))
    obj, _, err := p.Unmarshal(req.GetProperties())
    if err != nil {
        return plugin.NewCheckResponse(err), nil
    }
    var failures []error
    if failure := p.ops.Check(ctx, obj, ""); failure != nil {
        failures = append(failures, failure)
    }
    unks := req.GetUnknowns()
    if !unks["name"] {
        if failure := p.ops.Check(ctx, obj, "name"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("Topic", "name", failure))
        }
    }
    if !unks["topicName"] {
        if failure := p.ops.Check(ctx, obj, "topicName"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("Topic", "topicName", failure))
        }
    }
    if !unks["displayName"] {
        if failure := p.ops.Check(ctx, obj, "displayName"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("Topic", "displayName", failure))
        }
    }
    if !unks["subscription"] {
        if failure := p.ops.Check(ctx, obj, "subscription"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("Topic", "subscription", failure))
        }
    }
    if len(failures) > 0 {
        return plugin.NewCheckResponse(resource.NewErrors(failures)), nil
    }
    return plugin.NewCheckResponse(nil), nil
}

func (p *TopicProvider) Name(
    ctx context.Context, req *lumirpc.NameRequest) (*lumirpc.NameResponse, error) {
    contract.Assert(req.GetType() == string(TopicToken))
    obj, _, err := p.Unmarshal(req.GetProperties())
    if err != nil {
        return nil, err
    }
    if obj.Name == nil || *obj.Name == "" {
        if req.Unknowns[Topic_Name] {
            return nil, errors.New("Name property cannot be computed from unknown outputs")
        }
        return nil, errors.New("Name property cannot be empty")
    }
    return &lumirpc.NameResponse{Name: *obj.Name}, nil
}

func (p *TopicProvider) Create(
    ctx context.Context, req *lumirpc.CreateRequest) (*lumirpc.CreateResponse, error) {
    contract.Assert(req.GetType() == string(TopicToken))
    obj, _, err := p.Unmarshal(req.GetProperties())
    if err != nil {
        return nil, err
    }
    id, err := p.ops.Create(ctx, obj)
    if err != nil {
        return nil, err
    }
    return &lumirpc.CreateResponse{Id: string(id)}, nil
}

func (p *TopicProvider) Get(
    ctx context.Context, req *lumirpc.GetRequest) (*lumirpc.GetResponse, error) {
    contract.Assert(req.GetType() == string(TopicToken))
    id := resource.ID(req.GetId())
    obj, err := p.ops.Get(ctx, id)
    if err != nil {
        return nil, err
    }
    return &lumirpc.GetResponse{
        Properties: plugin.MarshalProperties(
            nil, resource.NewPropertyMap(obj), plugin.MarshalOptions{}),
    }, nil
}

func (p *TopicProvider) InspectChange(
    ctx context.Context, req *lumirpc.InspectChangeRequest) (*lumirpc.InspectChangeResponse, error) {
    contract.Assert(req.GetType() == string(TopicToken))
    id := resource.ID(req.GetId())
    old, oldprops, err := p.Unmarshal(req.GetOlds())
    if err != nil {
        return nil, err
    }
    new, newprops, err := p.Unmarshal(req.GetNews())
    if err != nil {
        return nil, err
    }
    var replaces []string
    diff := oldprops.Diff(newprops)
    if diff != nil {
        if diff.Changed("name") {
            replaces = append(replaces, "name")
        }
        if diff.Changed("topicName") {
            replaces = append(replaces, "topicName")
        }
    }
    more, err := p.ops.InspectChange(ctx, id, old, new, diff)
    if err != nil {
        return nil, err
    }
    return &lumirpc.InspectChangeResponse{
        Replaces: append(replaces, more...),
    }, err
}

func (p *TopicProvider) Update(
    ctx context.Context, req *lumirpc.UpdateRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(TopicToken))
    id := resource.ID(req.GetId())
    old, oldprops, err := p.Unmarshal(req.GetOlds())
    if err != nil {
        return nil, err
    }
    new, newprops, err := p.Unmarshal(req.GetNews())
    if err != nil {
        return nil, err
    }
    diff := oldprops.Diff(newprops)
    if err := p.ops.Update(ctx, id, old, new, diff); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *TopicProvider) Delete(
    ctx context.Context, req *lumirpc.DeleteRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(TopicToken))
    id := resource.ID(req.GetId())
    if err := p.ops.Delete(ctx, id); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *TopicProvider) Unmarshal(
    v *pbstruct.Struct) (*Topic, resource.PropertyMap, error) {
    var obj Topic
    props := plugin.UnmarshalProperties(nil, v, plugin.MarshalOptions{RawResources: true})
    return &obj, props, mapper.MapIU(props.Mappable(), &obj)
}

/* Marshalable Topic structure(s) */

// Topic is a marshalable representation of its corresponding IDL type.
type Topic struct {
    Name *string `lumi:"name,optional"`
    TopicName *string `lumi:"topicName,optional"`
    DisplayName *string `lumi:"displayName,optional"`
    Subscription *[]TopicSubscription `lumi:"subscription,optional"`
}

// Topic's properties have constants to make dealing with diffs and property bags easier.
const (
    Topic_Name = "name"
    Topic_TopicName = "topicName"
    Topic_DisplayName = "displayName"
    Topic_Subscription = "subscription"
)

/* Marshalable TopicSubscription structure(s) */

// TopicSubscription is a marshalable representation of its corresponding IDL type.
type TopicSubscription struct {
    Protocol TopicProtocol `lumi:"protocol"`
    Endpoint string `lumi:"endpoint"`
}

// TopicSubscription's properties have constants to make dealing with diffs and property bags easier.
const (
    TopicSubscription_Protocol = "protocol"
    TopicSubscription_Endpoint = "endpoint"
)

/* Typedefs */

type (
    TopicProtocol string
)

/* Constants */

const (
    ApplicationTopic TopicProtocol = "application"
    EmailJSONTopic TopicProtocol = "email-json"
    EmailTopic TopicProtocol = "email"
    HTTPSTopic TopicProtocol = "https"
    HTTPTopic TopicProtocol = "http"
    LambdaTopic TopicProtocol = "lambda"
    SMSTopic TopicProtocol = "sms"
    SQSTopic TopicProtocol = "sqs"
)


