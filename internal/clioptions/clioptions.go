// Copyright Mia srl
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package clioptions

type CLIOptions struct {
	MiactlConfig string

	Endpoint string
	Insecure bool
	CAFile   string

	Context     string
	Auth        string
	ProjectID   string
	CompanyID   string
	Environment string

	Revision   string
	DeployType string
	NoSemVer   bool
	TriggerID  string

	IAMRole            string
	ProjectIAMRole     string
	EnvironmentIAMRole string
	EntityID           string

	UserEmail                 string
	KeepUserGroupMemeberships bool

	UserEmails []string
	UserIDs    []string

	ServiceAccountID string

	BasicClientID     string
	BasicClientSecret string
	JWTJsonPath       string
	OutputPath        string

	InputFilePath string

	MarketplaceResourcePaths []string
	// MarketplaceItemID is the itemId field of a Marketplace item
	MarketplaceItemID string
	// MarketplaceItemVersion is the version field of a Marketplace item
	MarketplaceItemVersion string
	// MarketplaceItemObjectID is the _id of a Marketplace item
	MarketplaceItemObjectID     string
	MarketplaceFetchPublicItems bool

	FromCronJob string

	FollowLogs bool

	// OutputFormat describes the output format of some commands. Can be json or yaml.
	OutputFormat string

	ShowUsers           bool
	ShowGroups          bool
	ShowServiceAccounts bool

	ResolveExtensionsDetails bool
}

// NewCLIOptions return a new CLIOptions instance
func NewCLIOptions() *CLIOptions {
	return &CLIOptions{}
}
