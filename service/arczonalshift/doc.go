// Code generated by smithy-go-codegen DO NOT EDIT.

// Package arczonalshift provides the API client, operations, and parameter types
// for AWS ARC - Zonal Shift.
//
// This is the API Reference Guide for the zonal shift feature of Amazon Route 53
// Application Recovery Controller. This guide is for developers who need detailed
// information about zonal shift API actions, data types, and errors. Zonal shift
// is in preview release for Amazon Route 53 Application Recovery Controller and is
// subject to change. Zonal shift in Route 53 ARC enables you to move traffic for a
// load balancer resource away from an Availability Zone. Starting a zonal shift
// helps your application recover immediately, for example, from a developer's bad
// code deployment or from an AWS infrastructure failure in a single Availability
// Zone, reducing the impact and time lost from an issue in one zone. Supported AWS
// resources are automatically registered with Route 53 ARC. Resources that are
// registered for zonal shifts in Route 53 ARC are managed resources in Route 53
// ARC. You can start a zonal shift for any managed resource in your account in a
// Region. At this time, you can only start a zonal shift for Network Load
// Balancers and Application Load Balancers with cross-zone load balancing turned
// off. Zonal shifts are temporary. You must specify an expiration when you start a
// zonal shift, of up to three days initially. If you want to still keep traffic
// away from an Availability Zone, you can update the zonal shift and set a new
// expiration. You can also cancel a zonal shift, before it expires, for example,
// if you're ready to restore traffic to the Availability Zone. For more
// information about using zonal shift, see the Amazon Route 53 Application
// Recovery Controller Developer Guide
// (https://docs.aws.amazon.com/r53recovery/latest/dg/what-is-route53-recovery.html).
package arczonalshift
