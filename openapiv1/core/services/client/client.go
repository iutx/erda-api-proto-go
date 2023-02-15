// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: core_services.proto

package client

import (
	context "context"

	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/openapiv1/core/services/pb"
	grpc1 "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// Client provide all service clients.
type Client interface {
	// CoreServices core_services.proto
	CoreServices() pb.CoreServicesClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		coreServices: pb.NewCoreServicesClient(cc),
	}
}

type serviceClients struct {
	coreServices pb.CoreServicesClient
}

func (c *serviceClients) CoreServices() pb.CoreServicesClient {
	return c.coreServices
}

type coreServicesWrapper struct {
	client pb.CoreServicesClient
	opts   []grpc1.CallOption
}

func (s *coreServicesWrapper) CMDB_APPLICATIONS_LIST(ctx context.Context, req *pb.ApplicationListRequest) (*pb.ApplicationListResponse, error) {
	return s.client.CMDB_APPLICATIONS_LIST(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_APPLICATION_FETCH(ctx context.Context, req *pb.ApplicationFetchRequest) (*pb.ApplicationFetchResponse, error) {
	return s.client.CMDB_APPLICATION_FETCH(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_APPLICATION_PIN(ctx context.Context, req *pb.CMDB_APPLICATION_PIN_Request) (*emptypb.Empty, error) {
	return s.client.CMDB_APPLICATION_PIN(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_APPLICATION_UNPIN(ctx context.Context, req *pb.CMDB_APPLICATION_UNPIN_Request) (*emptypb.Empty, error) {
	return s.client.CMDB_APPLICATION_UNPIN(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_APPROVE_CREATE(ctx context.Context, req *pb.ApproveCreateRequest) (*pb.ApproveCreateResponse, error) {
	return s.client.CMDB_APPROVE_CREATE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_APP_LIST_TEMPLATES(ctx context.Context, req *pb.CMDB_APP_LIST_TEMPLATES_Request) (*emptypb.Empty, error) {
	return s.client.CMDB_APP_LIST_TEMPLATES(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_AUDITS_LIST_SET(ctx context.Context, req *pb.AuditListCleanCronRequest) (*pb.AuditListCleanCronResponse, error) {
	return s.client.CMDB_AUDITS_LIST_SET(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_AUDITS_SET(ctx context.Context, req *pb.AuditSetCleanCronRequest) (*pb.AuditSetCleanCronResponse, error) {
	return s.client.CMDB_AUDITS_SET(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_CLUSTER_DEREFERENCE(ctx context.Context, req *pb.DereferenceClusterRequest) (*pb.DereferenceClusterResponse, error) {
	return s.client.CMDB_CLUSTER_DEREFERENCE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_ERRORLOG_CREATE(ctx context.Context, req *pb.ErrorLogCreateRequest) (*pb.ErrorLogCreateResponse, error) {
	return s.client.CMDB_ERRORLOG_CREATE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_IMAGE_FETCH(ctx context.Context, req *pb.CMDB_IMAGE_FETCH_Request) (*emptypb.Empty, error) {
	return s.client.CMDB_IMAGE_FETCH(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_IMAGE_UPLOAD(ctx context.Context, req *pb.CMDB_IMAGE_UPLOAD_Request) (*emptypb.Empty, error) {
	return s.client.CMDB_IMAGE_UPLOAD(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_LABEL_CREATE(ctx context.Context, req *pb.ProjectLabelCreateRequest) (*pb.ProjectLabelCreateResponse, error) {
	return s.client.CMDB_LABEL_CREATE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_LABEL_DELETE(ctx context.Context, req *pb.CMDB_LABEL_DELETE_Request) (*emptypb.Empty, error) {
	return s.client.CMDB_LABEL_DELETE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_LABEL_LIST(ctx context.Context, req *pb.ProjectLabelListRequest) (*pb.ProjectLabelListResponse, error) {
	return s.client.CMDB_LABEL_LIST(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_LABEL_UPDATE(ctx context.Context, req *pb.ProjectLabelUpdateRequest) (*emptypb.Empty, error) {
	return s.client.CMDB_LABEL_UPDATE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_LICENSE(ctx context.Context, req *pb.CMDB_LICENSE_Request) (*pb.LicenseResponse, error) {
	return s.client.CMDB_LICENSE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_MANUAL_REVIEW_ADD(ctx context.Context, req *pb.CMDB_MANUAL_REVIEW_ADD_Request) (*emptypb.Empty, error) {
	return s.client.CMDB_MANUAL_REVIEW_ADD(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_MANUAL_REVIEW_ADDUSER(ctx context.Context, req *pb.CreateReviewUser) (*emptypb.Empty, error) {
	return s.client.CMDB_MANUAL_REVIEW_ADDUSER(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_MANUAL_REVIEW_AUTHORITY(ctx context.Context, req *pb.GetAuthorityByUserIdRequest) (*pb.GetAuthorityByUserIdResponse, error) {
	return s.client.CMDB_MANUAL_REVIEW_AUTHORITY(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_MANUAL_REVIEW_GET(ctx context.Context, req *pb.GetReviewByTaskIdIdRequest) (*pb.GetReviewByTaskIdIdResponse, error) {
	return s.client.CMDB_MANUAL_REVIEW_GET(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_MANUAL_REVIEW_LISTBYSPONSORID(ctx context.Context, req *pb.GetReviewsBySponsorIdRequest) (*pb.GetReviewsBySponsorIdResponse, error) {
	return s.client.CMDB_MANUAL_REVIEW_LISTBYSPONSORID(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_MANUAL_REVIEW_LISTBYUserID(ctx context.Context, req *pb.GetReviewsByUserIdRequest) (*pb.GetReviewsByUserIdResponse, error) {
	return s.client.CMDB_MANUAL_REVIEW_LISTBYUserID(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_MANUAL_REVIEW_UPDATE(ctx context.Context, req *pb.UpdateApproval) (*emptypb.Empty, error) {
	return s.client.CMDB_MANUAL_REVIEW_UPDATE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_MBOX_GET(ctx context.Context, req *pb.CMDB_MBOX_GET_Request) (*emptypb.Empty, error) {
	return s.client.CMDB_MBOX_GET(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_MBOX_LIST(ctx context.Context, req *pb.QueryMBoxRequest) (*pb.QueryMBoxResponse, error) {
	return s.client.CMDB_MBOX_LIST(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_MBOX_READ_ALL(ctx context.Context, req *pb.CMDB_MBOX_READ_ALL_Request) (*emptypb.Empty, error) {
	return s.client.CMDB_MBOX_READ_ALL(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_MBOX_STATS(ctx context.Context, req *pb.CMDB_MBOX_STATS_Request) (*emptypb.Empty, error) {
	return s.client.CMDB_MBOX_STATS(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_MEMBER_ADD(ctx context.Context, req *pb.MemberAddRequest) (*pb.MemberAddResponse, error) {
	return s.client.CMDB_MEMBER_ADD(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_MEMBER_ADD_BY_INVITECODE(ctx context.Context, req *pb.MemberAddByInviteCodeRequest) (*pb.MemberAddByInviteCodeResponse, error) {
	return s.client.CMDB_MEMBER_ADD_BY_INVITECODE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_MEMBER_LABEL_LIST(ctx context.Context, req *pb.CMDB_MEMBER_LABEL_LIST_Request) (*pb.MemberLabelListResponse, error) {
	return s.client.CMDB_MEMBER_LABEL_LIST(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_MEMBER_LIST(ctx context.Context, req *pb.MemberListRequest) (*pb.MemberListResponse, error) {
	return s.client.CMDB_MEMBER_LIST(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_MEMBER_REMOVE(ctx context.Context, req *pb.MemberRemoveRequest) (*pb.MemberRemoveResponse, error) {
	return s.client.CMDB_MEMBER_REMOVE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_MY_APPLICATIONS_LIST(ctx context.Context, req *pb.ApplicationListRequest) (*pb.ApplicationListResponse, error) {
	return s.client.CMDB_MY_APPLICATIONS_LIST(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_MY_PROJECTS_LIST(ctx context.Context, req *pb.ProjectListRequest) (*pb.ProjectListResponse, error) {
	return s.client.CMDB_MY_PROJECTS_LIST(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_NOTIFYITEM_QUERY(ctx context.Context, req *pb.QueryNotifyItemRequest) (*pb.QueryNotifyItemResponse, error) {
	return s.client.CMDB_NOTIFYITEM_QUERY(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_NOTIFYITEM_UPDATE(ctx context.Context, req *pb.UpdateNotifyItemRequest) (*pb.UpdateNotifyItemResponse, error) {
	return s.client.CMDB_NOTIFYITEM_UPDATE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_NOTIFY_CREATE(ctx context.Context, req *pb.CreateNotifyRequest) (*pb.CreateNotifyResponse, error) {
	return s.client.CMDB_NOTIFY_CREATE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_NOTIFY_DELETE(ctx context.Context, req *pb.CMDB_NOTIFY_DELETE_Request) (*pb.DeleteNotifyResponse, error) {
	return s.client.CMDB_NOTIFY_DELETE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_NOTIFY_DISABLE(ctx context.Context, req *pb.CMDB_NOTIFY_DISABLE_Request) (*emptypb.Empty, error) {
	return s.client.CMDB_NOTIFY_DISABLE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_NOTIFY_ENABLE(ctx context.Context, req *pb.CMDB_NOTIFY_ENABLE_Request) (*emptypb.Empty, error) {
	return s.client.CMDB_NOTIFY_ENABLE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_NOTIFY_FUZZY_QUERY(ctx context.Context, req *pb.CMDB_NOTIFY_FUZZY_QUERY_Request) (*emptypb.Empty, error) {
	return s.client.CMDB_NOTIFY_FUZZY_QUERY(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_NOTIFY_GET(ctx context.Context, req *pb.CMDB_NOTIFY_GET_Request) (*emptypb.Empty, error) {
	return s.client.CMDB_NOTIFY_GET(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_NOTIFY_HISTORIES(ctx context.Context, req *pb.QueryNotifyHistoryRequest) (*pb.QueryNotifyHistoryResponse, error) {
	return s.client.CMDB_NOTIFY_HISTORIES(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_NOTIFY_QUERY(ctx context.Context, req *pb.QueryNotifyRequest) (*pb.QueryNotifyResponse, error) {
	return s.client.CMDB_NOTIFY_QUERY(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_NOTIFY_UPDATE(ctx context.Context, req *pb.UpdateNotifyRequest) (*pb.UpdateNotifyResponse, error) {
	return s.client.CMDB_NOTIFY_UPDATE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_PERMISSION_ACCESS(ctx context.Context, req *pb.ScopeRoleAccessRequest) (*pb.PermissionListResponse, error) {
	return s.client.CMDB_PERMISSION_ACCESS(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_PERMISSION_CHECK(ctx context.Context, req *pb.PermissionCheckRequest) (*pb.PermissionCheckResponse, error) {
	return s.client.CMDB_PERMISSION_CHECK(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_PERMISSION_LIST(ctx context.Context, req *pb.CMDB_PERMISSION_LIST_Request) (*pb.ScopeRoleListResponse, error) {
	return s.client.CMDB_PERMISSION_LIST(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_PROJECT_FUNCTIONS_GET(ctx context.Context, req *pb.CMDB_PROJECT_FUNCTIONS_GET_Request) (*emptypb.Empty, error) {
	return s.client.CMDB_PROJECT_FUNCTIONS_GET(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_PROJECT_FUNCTIONS_SET(ctx context.Context, req *pb.ProjectFunctionSetRequest) (*pb.ProjectFunctionSetResponse, error) {
	return s.client.CMDB_PROJECT_FUNCTIONS_SET(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_PROJECT_GET_NSINFO(ctx context.Context, req *pb.CMDB_PROJECT_GET_NSINFO_Request) (*pb.ProjectNameSpaceInfoResponse, error) {
	return s.client.CMDB_PROJECT_GET_NSINFO(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_PROJECT_METRICS_HISTOGRAM(ctx context.Context, req *pb.CMDB_PROJECT_METRICS_HISTOGRAM_Request) (*emptypb.Empty, error) {
	return s.client.CMDB_PROJECT_METRICS_HISTOGRAM(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_PROJECT_UPDATE(ctx context.Context, req *pb.ProjectUpdateRequest) (*pb.ProjectUpdateResponse, error) {
	return s.client.CMDB_PROJECT_UPDATE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_PROJECT_WORKSPACE_ABILITIES_CREATE(ctx context.Context, req *pb.ProjectWorkSpaceAbility) (*emptypb.Empty, error) {
	return s.client.CMDB_PROJECT_WORKSPACE_ABILITIES_CREATE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_PROJECT_WORKSPACE_ABILITIES_DELETE(ctx context.Context, req *pb.CMDB_PROJECT_WORKSPACE_ABILITIES_DELETE_Request) (*emptypb.Empty, error) {
	return s.client.CMDB_PROJECT_WORKSPACE_ABILITIES_DELETE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_PROJECT_WORKSPACE_ABILITIES_GET(ctx context.Context, req *pb.CMDB_PROJECT_WORKSPACE_ABILITIES_GET_Request) (*pb.ProjectWorkSpaceAbilityResponse, error) {
	return s.client.CMDB_PROJECT_WORKSPACE_ABILITIES_GET(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_PROJECT_WORKSPACE_ABILITIES_UPDATE(ctx context.Context, req *pb.ProjectWorkSpaceAbility) (*emptypb.Empty, error) {
	return s.client.CMDB_PROJECT_WORKSPACE_ABILITIES_UPDATE(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) CMDB_ROLES_LIST_BY_USER(ctx context.Context, req *pb.ListMemberRolesByUserRequest) (*pb.ListMemberRolesByUserResponse, error) {
	return s.client.CMDB_ROLES_LIST_BY_USER(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) EVENTBOX_WEBSOCKET(ctx context.Context, req *pb.EVENTBOX_WEBSOCKET_Request) (*emptypb.Empty, error) {
	return s.client.EVENTBOX_WEBSOCKET(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) EVENTBOX_WEBSOCKET_INFO(ctx context.Context, req *pb.EVENTBOX_WEBSOCKET_INFO_Request) (*emptypb.Empty, error) {
	return s.client.EVENTBOX_WEBSOCKET_INFO(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) QUOTA_GET_PROJECT_NAMESPACES(ctx context.Context, req *pb.QUOTA_GET_PROJECT_NAMESPACES_Request) (*emptypb.Empty, error) {
	return s.client.QUOTA_GET_PROJECT_NAMESPACES(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *coreServicesWrapper) QUOTA_GET_PROJECT_QUOTA(ctx context.Context, req *pb.QUOTA_GET_PROJECT_QUOTA_Request) (*emptypb.Empty, error) {
	return s.client.QUOTA_GET_PROJECT_QUOTA(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}