## Defining a dynamic provider
Providers have historically been built using the Terraform Plugin SDK. The SDK publishes an API centered around declared schema and CRUD operations for resource types. This has been a succesful model in that it's flexible enough to smooth over heterogenous upstream APIs, delivering a consistent experience to the practicioner.

Often providers end up following a formulaic, sometimes copypaste development workflow. The desire to declare a schema with dynamic types has existed for sometime with providers that don't communicate with many endpoints (or at the very least, the endpoint and order of operations is largely identical across most resources). With the release of the terraform-plugin-go library, it is now possible to leverage Terraform's `DynamicPseudoType`, albeit with caveats.

The first provider using this technique put into production is the kubernetes-alpha provider. There is also a proof of concept sql provider. Deciding to develop a provider using the low level terraform-plugin-go library isn't recommended in general due to the heavy lifting you will need to perform, you are essentially re-implementing much of the logic of the existing Terraform Plugin SDK. If you choose to attempt to build a dynamic provider using terraform-plugin-go, consider the following:

1) Is the Create, Read, Update, Delete code consistent across resources? The URL and payload formation should be predictable and consistent.
2) Is validation of input available at runtime? Or could it be encapsulated statically into the provider codebase in a way that isn't as much work as writing a provider the old fashioned way?
3) Is a schema of sorts available at runtime, either via an endpoint such as DescribeX, or perhaps an OpenAPI spec is available and could be embedded?

### 1: API/library Consistency
Interacting with the upstream service needs to be consistent across resources. You may be working with an API that has a single endpoint that accepts a single document describing the entire deployment, or perhaps the resources are controlled by a predictable URL scheme and a consistent payload (such was the case for Salesforce sObjects).

It's realistic that some exceptions exist to handling of the CRUD operations, an example being the `User` object in Salesforce cannot be deleted, but rather only deactivated (via update). A provider may need to create a framework for overriding the code for specific operations. Again, it should be kept in mind that if the amount of variations becomes extremely large, it may not be worth the trouble that comes with terraform-plugin-go.

### 2: Validation
The Terraform Plugin SDK provides means for defining validation logic for the schema beyond the basics of fields being required. A dynamic provider will need to provide validation at runtime, either via checking the proposed state against an embedded schema of some sort, against an upstream validation service, or perhaps the ability to dry-run the changes.

Without a solution as described, the practitioner will only receive errors at apply time, which is a degraded experience from a traditional provider. Developing a framework for statically defining validation could be a solution but may be an indicator that terraform-plugin-go isn't a good fit.

### 3: Schema
`DynamicPseudoType` allows for the schema to be defined for a field at first plan. Meaning the provider needs to have means for translating a schema of some form, whether it be OpenAPI spec, or a series of describe endpoints, but it needs to define the `DynamicPseudoType` in concrete Terraform terms at first plan.

Essentially the `DynamicPseudoType` becomes an `Object`, and at first plan, the provider will need to fill in all the possible optional attributes and types so that subsequent operations don't produce an inconsistent result error. Terraform considers the types of all attributes (and nested blocks) as the unique signature of an `Object` and that signature can't change between operations.

## Final notes
In general a dynamic provider is a good fit if the provider has a large number of resources that follow a consistent code pattern, and the schema is available in some format at runtime. You will pay for work upfront, but as more resources come online for the upstream service, if they remain consistent with the other resource types, they will work immediately (or require small customization to enable).