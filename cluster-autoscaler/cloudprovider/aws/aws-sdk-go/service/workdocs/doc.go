// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package workdocs provides the client and types for making API
// requests to Amazon WorkDocs.
//
// The Amazon WorkDocs API is designed for the following use cases:
//
//   - File Migration: File migration applications are supported for users
//     who want to migrate their files from an on-premises or off-premises file
//     system or service. Users can insert files into a user directory structure,
//     as well as allow for basic metadata changes, such as modifications to
//     the permissions of files.
//
//   - Security: Support security applications are supported for users who
//     have additional security needs, such as antivirus or data loss prevention.
//     The API actions, along with CloudTrail, allow these applications to detect
//     when changes occur in Amazon WorkDocs. Then, the application can take
//     the necessary actions and replace the target file. If the target file
//     violates the policy, the application can also choose to email the user.
//
//   - eDiscovery/Analytics: General administrative applications are supported,
//     such as eDiscovery and analytics. These applications can choose to mimic
//     or record the actions in an Amazon WorkDocs site, along with CloudTrail,
//     to replicate data for eDiscovery, backup, or analytical applications.
//
// All Amazon WorkDocs API actions are Amazon authenticated and certificate-signed.
// They not only require the use of the Amazon Web Services SDK, but also allow
// for the exclusive use of IAM users and roles to help facilitate access, trust,
// and permission policies. By creating a role and allowing an IAM user to access
// the Amazon WorkDocs site, the IAM user gains full administrative visibility
// into the entire Amazon WorkDocs site (or as set in the IAM policy). This
// includes, but is not limited to, the ability to modify file permissions and
// upload any file to any user. This allows developers to perform the three
// use cases above, as well as give users the ability to grant access on a selective
// basis using the IAM model.
//
// The pricing for Amazon WorkDocs APIs varies depending on the API call type
// for these actions:
//
//   - READ (Get*)
//
//   - WRITE (Activate*, Add*, Create*, Deactivate*, Initiate*, Update*)
//
//   - LIST (Describe*)
//
//   - DELETE*, CANCEL
//
// For information about Amazon WorkDocs API pricing, see Amazon WorkDocs Pricing
// (https://aws.amazon.com/workdocs/pricing/).
//
// See https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01 for more information on this service.
//
// See workdocs package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/workdocs/
//
// # Using the Client
//
// To contact Amazon WorkDocs with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the Amazon WorkDocs client WorkDocs for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/workdocs/#New
package workdocs
