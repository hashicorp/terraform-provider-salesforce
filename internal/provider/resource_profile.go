package provider

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nimajalali/go-force/force"
)

type profileType struct {
}

func (profileType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Type:     types.StringType,
				Computed: true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					staticComputed{},
				},
			},
			"name": {
				Type:     types.StringType,
				Required: true,
				Validators: []tfsdk.AttributeValidator{
					emptyString{},
				},
			},
			"description": {
				Type:     types.StringType,
				Optional: true,
			},
			"user_license_id": {
				Type:     types.StringType,
				Required: true,
				Validators: []tfsdk.AttributeValidator{
					emptyString{},
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					NormalizeId{},
					tfsdk.RequiresReplace(),
				},
			},
			// permissions
			"permissions_email_single":                       {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_email_mass":                         {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_edit_task":                          {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_edit_event":                         {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_export_report":                      {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_import_personal":                    {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_data_export":                        {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_users":                       {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_edit_public_filters":                {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_edit_public_templates":              {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_modify_all_data":                    {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_cases":                       {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_mass_inline_edit":                   {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_edit_knowledge":                     {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_knowledge":                   {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_solutions":                   {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_customize_application":              {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_edit_readonly_fields":               {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_run_reports":                        {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_setup":                         {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_transfer_any_entity":                {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_new_report_builder":                 {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_activate_contract":                  {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_activate_order":                     {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_import_leads":                       {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_leads":                       {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_transfer_any_lead":                  {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_all_data":                      {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_edit_public_documents":              {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_encrypted_data":                {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_edit_brand_templates":               {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_edit_html_templates":                {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_chatter_internal_user":              {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_encryption_keys":             {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_delete_activated_contract":          {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_chatter_invite_external_users":      {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_send_sit_requests":                  {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_remote_access":               {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_can_use_new_dashboard_builder":      {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_categories":                  {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_convert_leads":                      {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_password_never_expires":             {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_use_team_reassign_wizards":          {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_edit_activated_orders":              {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_install_multiforce":                 {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_publish_multiforce":                 {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_chatter_own_groups":                 {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_edit_opp_line_item_unit_price":      {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_create_multiforce":                  {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_bulk_api_hard_delete":               {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_solution_import":                    {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_call_centers":                {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_synonyms":                    {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_content":                       {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_email_client_config":         {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_enable_notifications":               {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_data_integrations":           {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_distribute_from_pers_wksp":          {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_data_categories":               {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_data_categories":             {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_author_apex":                        {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_mobile":                      {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_api_enabled":                        {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_custom_report_types":         {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_edit_case_comments":                 {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_transfer_any_case":                  {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_content_administrator":              {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_create_workspaces":                  {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_content_permissions":         {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_content_properties":          {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_content_types":               {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_exchange_config":             {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_analytic_snapshots":          {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_schedule_reports":                   {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_business_hour_holidays":      {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_entitlements":                {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_dynamic_dashboards":          {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_custom_sidebar_on_all_pages":        {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_interaction":                 {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_my_teams_dashboards":           {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_moderate_chatter":                   {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_reset_passwords":                    {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_flow_ufl_required":                  {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_can_insert_feed_system_fields":      {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_activities_access":                  {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_knowledge_import_export":     {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_email_template_management":          {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_email_administration":               {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_chatter_messages":            {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_allow_email_ic":                     {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_chatter_file_link":                  {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_force_two_factor":                   {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_event_log_files":               {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_networks":                    {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_auth_providers":              {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_run_flow":                           {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_create_customize_dashboards":        {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_create_dashboard_folders":           {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_public_dashboards":             {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_dashbds_in_pub_folders":      {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_create_customize_reports":           {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_create_report_folders":              {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_public_reports":                {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_reports_in_pub_folders":      {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_edit_my_dashboards":                 {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_edit_my_reports":                    {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_all_users":                     {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_allow_universal_search":             {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_connect_org_to_environment_hub":     {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_work_calibration_user":              {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_create_customize_filters":           {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_work_dot_com_user_perm":             {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_content_hub_user":                   {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_govern_networks":                    {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_sales_console":                      {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_two_factor_api":                     {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_delete_topics":                      {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_edit_topics":                        {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_create_topics":                      {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_assign_topics":                      {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_identity_enabled":                   {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_identity_connect":                   {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_allow_view_knowledge":               {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_content_workspaces":                 {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_create_work_badge_definition":       {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_search_promotion_rules":      {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_custom_mobile_apps_access":          {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_help_link":                     {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_profiles_permissionsets":     {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_assign_permission_sets":             {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_roles":                       {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_ip_addresses":                {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_sharing":                     {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_internal_users":              {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_password_policies":           {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_login_access_policies":       {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_platform_events":               {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_custom_permissions":          {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_can_verify_comment":                 {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_unlisted_groups":             {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_std_automatic_activity_capture":     {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_insights_app_dashboard_editor":      {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_two_factor":                  {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_insights_app_user":                  {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_insights_app_admin":                 {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_insights_app_elt_editor":            {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_insights_app_upload_user":           {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_insights_create_application":        {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_lightning_experience_user":          {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_data_leakage_events":           {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_config_custom_recs":                 {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_submit_macros_allowed":              {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_bulk_macros_allowed":                {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_share_internal_articles":            {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_session_permission_sets":     {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_templated_app":               {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_use_templated_app":                  {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_send_announcement_emails":           {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_chatter_edit_own_post":              {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_chatter_edit_own_record_post":       {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_sales_analytics_user":               {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_service_analytics_user":             {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_wave_tabular_download":              {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_automatic_activity_capture":         {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_import_custom_objects":              {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_delegated_two_factor":               {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_chatter_compose_ui_codesnippet":     {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_select_files_from_salesforce":       {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_moderate_network_users":             {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_merge_topics":                       {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_subscribe_to_lightning_reports":     {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_pvt_rpts_and_dashbds":        {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_allow_lightning_login":              {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_campaign_influence2":                {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_data_assessment":               {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_remove_direct_message_members":      {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_can_approve_feed_post":              {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_add_direct_message_members":         {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_allow_view_edit_converted_leads":    {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_show_company_name_as_user_badge":    {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_access_cmc":                         {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_health_check":                  {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_health_check":                {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_packaging2":                         {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_certificates":                {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_create_report_in_lightning":         {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_prevent_classic_experience":         {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_hide_read_by_list":                  {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_use_smart_data_discovery":           {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_get_smart_data_discovery":           {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_create_update_sdd_dataset":          {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_create_update_sdd_story":            {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_smart_data_discovery":        {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_share_smart_data_discovery_story":   {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_smart_data_discovery_model":  {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_list_email_send":                    {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_feed_pinning":                       {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_change_dashboard_colors":            {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_iot_user":                           {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_recommendation_strategies":   {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_propositions":                {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_close_conversations":                {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_subscribe_report_roles_grps":        {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_subscribe_dashboard_roles_grps":     {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_use_web_link":                       {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_has_unlimited_nba_executions":       {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_only_embedded_app_user":        {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_adoption_analytics_user":            {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_all_activities":                {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_subscribe_report_to_other_users":    {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_lightning_console_allowed_for_user": {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_subscribe_reports_run_as_user":      {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_subscribe_to_lightning_dashboards":  {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_subscribe_dashboard_to_other_users": {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_create_ltng_temp_in_pub":            {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_transactional_email_send":           {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_private_static_resources":      {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_create_ltng_temp_folder":            {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_apex_rest_services":                 {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_enable_community_app_launcher":      {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_give_recognition_badge":             {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_can_run_analysis":                   {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_use_my_search":                      {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_ltng_promo_reserved01_user_perm":    {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_subscriptions":               {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_wave_manage_private_assets_user":    {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_can_edit_data_prep_recipe":          {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_add_analytics_remote_connections":   {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_surveys":                     {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_use_assistant_dialog":               {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_use_query_suggestions":              {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_record_visibility_api":              {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_roles":                         {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_smart_data_discovery_for_community": {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_can_manage_maps":                    {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_story_on_ds_with_predicate":         {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_lm_outbound_messaging_user_perm":    {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_modify_data_classification":         {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_privacy_data_access":                {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_query_all_files":                    {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_modify_metadata":                    {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_cms":                         {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_sandbox_testing_in_community_app":   {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_can_edit_prompts":                   {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_user_pii":                      {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_hub_connections":             {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_b_2_b_marketing_analytics_user":     {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_trace_xds_queries":                  {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_security_command_center":       {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_security_command_center":     {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_all_custom_settings":           {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_all_foreign_key_names":         {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_allow_survey_advanced_features":     {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_add_wave_notification_recipients":   {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_headless_cms_access":                {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_lm_end_messaging_session_user_perm": {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_consent_api_update":                 {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_payments_api_user":                  {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_access_content_builder":             {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_account_switcher_user":              {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_anomaly_events":                {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_c_360_a_connections":         {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_release_updates":             {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_all_profiles":                  {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_skip_identity_confirmation":         {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_learning_manager":                   {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_send_custom_notifications":          {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_packaging2_delete":                  {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_autonomous_analytics_privacy":       {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_sonic_consumer":                     {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_bot_manage_bots":                    {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_bot_manage_bots_training_data":      {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_learning_reporting":          {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_isotope_c_to_c_user":                {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_isotope_access":                     {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_isotope_lex":                        {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_get_smart_data_discovery_external":  {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_quip_metrics_access":                {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_quip_user_engagement_metrics":       {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_manage_external_connections":        {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_use_subscription_emails":            {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_ai_view_insight_objects":            {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_ai_create_insight_objects":          {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_lifecycle_management_api_user":      {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_native_webview_scrolling":           {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
			"permissions_view_developer_name":                {Type: types.BoolType, Optional: true, Computed: true, PlanModifiers: tfsdk.AttributePlanModifiers{optionalComputed{}}},
		},
	}, nil
}

func (p profileType) NewResource(_ context.Context, prov tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	provider, ok := prov.(*provider)
	if !ok {
		return nil, diag.Diagnostics{errorConvertingProvider(p)}
	}
	return &profileResource{
		Resource: Resource{
			Client:              provider.client,
			Data:                &profileResourceData{},
			NeedsGetAfterUpsert: true,
		},
	}, nil
}

type profileResource struct {
	Resource
}

type profileResourceData struct {
	Name          string       `tfsdk:"name" force:",omitempty"`
	Description   *string      `tfsdk:"description" force:",omitempty"`
	UserLicenseId string       `tfsdk:"user_license_id" force:",omitempty"`
	Id            types.String `tfsdk:"id" force:"-"`
	// permissions
	PermissionsEmailSingle                    types.Bool `tfsdk:"permissions_email_single" force:",omitempty"`
	PermissionsEmailMass                      types.Bool `tfsdk:"permissions_email_mass" force:",omitempty"`
	PermissionsEditTask                       types.Bool `tfsdk:"permissions_edit_task" force:",omitempty"`
	PermissionsEditEvent                      types.Bool `tfsdk:"permissions_edit_event" force:",omitempty"`
	PermissionsExportReport                   types.Bool `tfsdk:"permissions_export_report" force:",omitempty"`
	PermissionsImportPersonal                 types.Bool `tfsdk:"permissions_import_personal" force:",omitempty"`
	PermissionsDataExport                     types.Bool `tfsdk:"permissions_data_export" force:",omitempty"`
	PermissionsManageUsers                    types.Bool `tfsdk:"permissions_manage_users" force:",omitempty"`
	PermissionsEditPublicFilters              types.Bool `tfsdk:"permissions_edit_public_filters" force:",omitempty"`
	PermissionsEditPublicTemplates            types.Bool `tfsdk:"permissions_edit_public_templates" force:",omitempty"`
	PermissionsModifyAllData                  types.Bool `tfsdk:"permissions_modify_all_data" force:",omitempty"`
	PermissionsManageCases                    types.Bool `tfsdk:"permissions_manage_cases" force:",omitempty"`
	PermissionsMassInlineEdit                 types.Bool `tfsdk:"permissions_mass_inline_edit" force:",omitempty"`
	PermissionsEditKnowledge                  types.Bool `tfsdk:"permissions_edit_knowledge" force:",omitempty"`
	PermissionsManageKnowledge                types.Bool `tfsdk:"permissions_manage_knowledge" force:",omitempty"`
	PermissionsManageSolutions                types.Bool `tfsdk:"permissions_manage_solutions" force:",omitempty"`
	PermissionsCustomizeApplication           types.Bool `tfsdk:"permissions_customize_application" force:",omitempty"`
	PermissionsEditReadonlyFields             types.Bool `tfsdk:"permissions_edit_readonly_fields" force:",omitempty"`
	PermissionsRunReports                     types.Bool `tfsdk:"permissions_run_reports" force:",omitempty"`
	PermissionsViewSetup                      types.Bool `tfsdk:"permissions_view_setup" force:",omitempty"`
	PermissionsTransferAnyEntity              types.Bool `tfsdk:"permissions_transfer_any_entity" force:",omitempty"`
	PermissionsNewReportBuilder               types.Bool `tfsdk:"permissions_new_report_builder" force:",omitempty"`
	PermissionsActivateContract               types.Bool `tfsdk:"permissions_activate_contract" force:",omitempty"`
	PermissionsActivateOrder                  types.Bool `tfsdk:"permissions_activate_order" force:",omitempty"`
	PermissionsImportLeads                    types.Bool `tfsdk:"permissions_import_leads" force:",omitempty"`
	PermissionsManageLeads                    types.Bool `tfsdk:"permissions_manage_leads" force:",omitempty"`
	PermissionsTransferAnyLead                types.Bool `tfsdk:"permissions_transfer_any_lead" force:",omitempty"`
	PermissionsViewAllData                    types.Bool `tfsdk:"permissions_view_all_data" force:",omitempty"`
	PermissionsEditPublicDocuments            types.Bool `tfsdk:"permissions_edit_public_documents" force:",omitempty"`
	PermissionsViewEncryptedData              types.Bool `tfsdk:"permissions_view_encrypted_data" force:",omitempty"`
	PermissionsEditBrandTemplates             types.Bool `tfsdk:"permissions_edit_brand_templates" force:",omitempty"`
	PermissionsEditHtmlTemplates              types.Bool `tfsdk:"permissions_edit_html_templates" force:",omitempty"`
	PermissionsChatterInternalUser            types.Bool `tfsdk:"permissions_chatter_internal_user" force:",omitempty"`
	PermissionsManageEncryptionKeys           types.Bool `tfsdk:"permissions_manage_encryption_keys" force:",omitempty"`
	PermissionsDeleteActivatedContract        types.Bool `tfsdk:"permissions_delete_activated_contract" force:",omitempty"`
	PermissionsChatterInviteExternalUsers     types.Bool `tfsdk:"permissions_chatter_invite_external_users" force:",omitempty"`
	PermissionsSendSitRequests                types.Bool `tfsdk:"permissions_send_sit_requests" force:",omitempty"`
	PermissionsManageRemoteAccess             types.Bool `tfsdk:"permissions_manage_remote_access" force:",omitempty"`
	PermissionsCanUseNewDashboardBuilder      types.Bool `tfsdk:"permissions_can_use_new_dashboard_builder" force:",omitempty"`
	PermissionsManageCategories               types.Bool `tfsdk:"permissions_manage_categories" force:",omitempty"`
	PermissionsConvertLeads                   types.Bool `tfsdk:"permissions_convert_leads" force:",omitempty"`
	PermissionsPasswordNeverExpires           types.Bool `tfsdk:"permissions_password_never_expires" force:",omitempty"`
	PermissionsUseTeamReassignWizards         types.Bool `tfsdk:"permissions_use_team_reassign_wizards" force:",omitempty"`
	PermissionsEditActivatedOrders            types.Bool `tfsdk:"permissions_edit_activated_orders" force:",omitempty"`
	PermissionsInstallMultiforce              types.Bool `tfsdk:"permissions_install_multiforce" force:",omitempty"`
	PermissionsPublishMultiforce              types.Bool `tfsdk:"permissions_publish_multiforce" force:",omitempty"`
	PermissionsChatterOwnGroups               types.Bool `tfsdk:"permissions_chatter_own_groups" force:",omitempty"`
	PermissionsEditOppLineItemUnitPrice       types.Bool `tfsdk:"permissions_edit_opp_line_item_unit_price" force:",omitempty"`
	PermissionsCreateMultiforce               types.Bool `tfsdk:"permissions_create_multiforce" force:",omitempty"`
	PermissionsBulkApiHardDelete              types.Bool `tfsdk:"permissions_bulk_api_hard_delete" force:",omitempty"`
	PermissionsSolutionImport                 types.Bool `tfsdk:"permissions_solution_import" force:",omitempty"`
	PermissionsManageCallCenters              types.Bool `tfsdk:"permissions_manage_call_centers" force:",omitempty"`
	PermissionsManageSynonyms                 types.Bool `tfsdk:"permissions_manage_synonyms" force:",omitempty"`
	PermissionsViewContent                    types.Bool `tfsdk:"permissions_view_content" force:",omitempty"`
	PermissionsManageEmailClientConfig        types.Bool `tfsdk:"permissions_manage_email_client_config" force:",omitempty"`
	PermissionsEnableNotifications            types.Bool `tfsdk:"permissions_enable_notifications" force:",omitempty"`
	PermissionsManageDataIntegrations         types.Bool `tfsdk:"permissions_manage_data_integrations" force:",omitempty"`
	PermissionsDistributeFromPersWksp         types.Bool `tfsdk:"permissions_distribute_from_pers_wksp" force:",omitempty"`
	PermissionsViewDataCategories             types.Bool `tfsdk:"permissions_view_data_categories" force:",omitempty"`
	PermissionsManageDataCategories           types.Bool `tfsdk:"permissions_manage_data_categories" force:",omitempty"`
	PermissionsAuthorApex                     types.Bool `tfsdk:"permissions_author_apex" force:",omitempty"`
	PermissionsManageMobile                   types.Bool `tfsdk:"permissions_manage_mobile" force:",omitempty"`
	PermissionsApiEnabled                     types.Bool `tfsdk:"permissions_api_enabled" force:",omitempty"`
	PermissionsManageCustomReportTypes        types.Bool `tfsdk:"permissions_manage_custom_report_types" force:",omitempty"`
	PermissionsEditCaseComments               types.Bool `tfsdk:"permissions_edit_case_comments" force:",omitempty"`
	PermissionsTransferAnyCase                types.Bool `tfsdk:"permissions_transfer_any_case" force:",omitempty"`
	PermissionsContentAdministrator           types.Bool `tfsdk:"permissions_content_administrator" force:",omitempty"`
	PermissionsCreateWorkspaces               types.Bool `tfsdk:"permissions_create_workspaces" force:",omitempty"`
	PermissionsManageContentPermissions       types.Bool `tfsdk:"permissions_manage_content_permissions" force:",omitempty"`
	PermissionsManageContentProperties        types.Bool `tfsdk:"permissions_manage_content_properties" force:",omitempty"`
	PermissionsManageContentTypes             types.Bool `tfsdk:"permissions_manage_content_types" force:",omitempty"`
	PermissionsManageExchangeConfig           types.Bool `tfsdk:"permissions_manage_exchange_config" force:",omitempty"`
	PermissionsManageAnalyticSnapshots        types.Bool `tfsdk:"permissions_manage_analytic_snapshots" force:",omitempty"`
	PermissionsScheduleReports                types.Bool `tfsdk:"permissions_schedule_reports" force:",omitempty"`
	PermissionsManageBusinessHourHolidays     types.Bool `tfsdk:"permissions_manage_business_hour_holidays" force:",omitempty"`
	PermissionsManageEntitlements             types.Bool `tfsdk:"permissions_manage_entitlements" force:",omitempty"`
	PermissionsManageDynamicDashboards        types.Bool `tfsdk:"permissions_manage_dynamic_dashboards" force:",omitempty"`
	PermissionsCustomSidebarOnAllPages        types.Bool `tfsdk:"permissions_custom_sidebar_on_all_pages" force:",omitempty"`
	PermissionsManageInteraction              types.Bool `tfsdk:"permissions_manage_interaction" force:",omitempty"`
	PermissionsViewMyTeamsDashboards          types.Bool `tfsdk:"permissions_view_my_teams_dashboards" force:",omitempty"`
	PermissionsModerateChatter                types.Bool `tfsdk:"permissions_moderate_chatter" force:",omitempty"`
	PermissionsResetPasswords                 types.Bool `tfsdk:"permissions_reset_passwords" force:",omitempty"`
	PermissionsFlowUFLRequired                types.Bool `tfsdk:"permissions_flow_ufl_required" force:",omitempty"`
	PermissionsCanInsertFeedSystemFields      types.Bool `tfsdk:"permissions_can_insert_feed_system_fields" force:",omitempty"`
	PermissionsActivitiesAccess               types.Bool `tfsdk:"permissions_activities_access" force:",omitempty"`
	PermissionsManageKnowledgeImportExport    types.Bool `tfsdk:"permissions_manage_knowledge_import_export" force:",omitempty"`
	PermissionsEmailTemplateManagement        types.Bool `tfsdk:"permissions_email_template_management" force:",omitempty"`
	PermissionsEmailAdministration            types.Bool `tfsdk:"permissions_email_administration" force:",omitempty"`
	PermissionsManageChatterMessages          types.Bool `tfsdk:"permissions_manage_chatter_messages" force:",omitempty"`
	PermissionsAllowEmailIC                   types.Bool `tfsdk:"permissions_allow_email_ic" force:",omitempty"`
	PermissionsChatterFileLink                types.Bool `tfsdk:"permissions_chatter_file_link" force:",omitempty"`
	PermissionsForceTwoFactor                 types.Bool `tfsdk:"permissions_force_two_factor" force:",omitempty"`
	PermissionsViewEventLogFiles              types.Bool `tfsdk:"permissions_view_event_log_files" force:",omitempty"`
	PermissionsManageNetworks                 types.Bool `tfsdk:"permissions_manage_networks" force:",omitempty"`
	PermissionsManageAuthProviders            types.Bool `tfsdk:"permissions_manage_auth_providers" force:",omitempty"`
	PermissionsRunFlow                        types.Bool `tfsdk:"permissions_run_flow" force:",omitempty"`
	PermissionsCreateCustomizeDashboards      types.Bool `tfsdk:"permissions_create_customize_dashboards" force:",omitempty"`
	PermissionsCreateDashboardFolders         types.Bool `tfsdk:"permissions_create_dashboard_folders" force:",omitempty"`
	PermissionsViewPublicDashboards           types.Bool `tfsdk:"permissions_view_public_dashboards" force:",omitempty"`
	PermissionsManageDashbdsInPubFolders      types.Bool `tfsdk:"permissions_manage_dashbds_in_pub_folders" force:",omitempty"`
	PermissionsCreateCustomizeReports         types.Bool `tfsdk:"permissions_create_customize_reports" force:",omitempty"`
	PermissionsCreateReportFolders            types.Bool `tfsdk:"permissions_create_report_folders" force:",omitempty"`
	PermissionsViewPublicReports              types.Bool `tfsdk:"permissions_view_public_reports" force:",omitempty"`
	PermissionsManageReportsInPubFolders      types.Bool `tfsdk:"permissions_manage_reports_in_pub_folders" force:",omitempty"`
	PermissionsEditMyDashboards               types.Bool `tfsdk:"permissions_edit_my_dashboards" force:",omitempty"`
	PermissionsEditMyReports                  types.Bool `tfsdk:"permissions_edit_my_reports" force:",omitempty"`
	PermissionsViewAllUsers                   types.Bool `tfsdk:"permissions_view_all_users" force:",omitempty"`
	PermissionsAllowUniversalSearch           types.Bool `tfsdk:"permissions_allow_universal_search" force:",omitempty"`
	PermissionsConnectOrgToEnvironmentHub     types.Bool `tfsdk:"permissions_connect_org_to_environment_hub" force:",omitempty"`
	PermissionsWorkCalibrationUser            types.Bool `tfsdk:"permissions_work_calibration_user" force:",omitempty"`
	PermissionsCreateCustomizeFilters         types.Bool `tfsdk:"permissions_create_customize_filters" force:",omitempty"`
	PermissionsWorkDotComUserPerm             types.Bool `tfsdk:"permissions_work_dot_com_user_perm" force:",omitempty"`
	PermissionsContentHubUser                 types.Bool `tfsdk:"permissions_content_hub_user" force:",omitempty"`
	PermissionsGovernNetworks                 types.Bool `tfsdk:"permissions_govern_networks" force:",omitempty"`
	PermissionsSalesConsole                   types.Bool `tfsdk:"permissions_sales_console" force:",omitempty"`
	PermissionsTwoFactorApi                   types.Bool `tfsdk:"permissions_two_factor_api" force:",omitempty"`
	PermissionsDeleteTopics                   types.Bool `tfsdk:"permissions_delete_topics" force:",omitempty"`
	PermissionsEditTopics                     types.Bool `tfsdk:"permissions_edit_topics" force:",omitempty"`
	PermissionsCreateTopics                   types.Bool `tfsdk:"permissions_create_topics" force:",omitempty"`
	PermissionsAssignTopics                   types.Bool `tfsdk:"permissions_assign_topics" force:",omitempty"`
	PermissionsIdentityEnabled                types.Bool `tfsdk:"permissions_identity_enabled" force:",omitempty"`
	PermissionsIdentityConnect                types.Bool `tfsdk:"permissions_identity_connect" force:",omitempty"`
	PermissionsAllowViewKnowledge             types.Bool `tfsdk:"permissions_allow_view_knowledge" force:",omitempty"`
	PermissionsContentWorkspaces              types.Bool `tfsdk:"permissions_content_workspaces" force:",omitempty"`
	PermissionsCreateWorkBadgeDefinition      types.Bool `tfsdk:"permissions_create_work_badge_definition" force:",omitempty"`
	PermissionsManageSearchPromotionRules     types.Bool `tfsdk:"permissions_manage_search_promotion_rules" force:",omitempty"`
	PermissionsCustomMobileAppsAccess         types.Bool `tfsdk:"permissions_custom_mobile_apps_access" force:",omitempty"`
	PermissionsViewHelpLink                   types.Bool `tfsdk:"permissions_view_help_link" force:",omitempty"`
	PermissionsManageProfilesPermissionsets   types.Bool `tfsdk:"permissions_manage_profiles_permissionsets" force:",omitempty"`
	PermissionsAssignPermissionSets           types.Bool `tfsdk:"permissions_assign_permission_sets" force:",omitempty"`
	PermissionsManageRoles                    types.Bool `tfsdk:"permissions_manage_roles" force:",omitempty"`
	PermissionsManageIpAddresses              types.Bool `tfsdk:"permissions_manage_ip_addresses" force:",omitempty"`
	PermissionsManageSharing                  types.Bool `tfsdk:"permissions_manage_sharing" force:",omitempty"`
	PermissionsManageInternalUsers            types.Bool `tfsdk:"permissions_manage_internal_users" force:",omitempty"`
	PermissionsManagePasswordPolicies         types.Bool `tfsdk:"permissions_manage_password_policies" force:",omitempty"`
	PermissionsManageLoginAccessPolicies      types.Bool `tfsdk:"permissions_manage_login_access_policies" force:",omitempty"`
	PermissionsViewPlatformEvents             types.Bool `tfsdk:"permissions_view_platform_events" force:",omitempty"`
	PermissionsManageCustomPermissions        types.Bool `tfsdk:"permissions_manage_custom_permissions" force:",omitempty"`
	PermissionsCanVerifyComment               types.Bool `tfsdk:"permissions_can_verify_comment" force:",omitempty"`
	PermissionsManageUnlistedGroups           types.Bool `tfsdk:"permissions_manage_unlisted_groups" force:",omitempty"`
	PermissionsStdAutomaticActivityCapture    types.Bool `tfsdk:"permissions_std_automatic_activity_capture" force:",omitempty"`
	PermissionsInsightsAppDashboardEditor     types.Bool `tfsdk:"permissions_insights_app_dashboard_editor" force:",omitempty"`
	PermissionsManageTwoFactor                types.Bool `tfsdk:"permissions_manage_two_factor" force:",omitempty"`
	PermissionsInsightsAppUser                types.Bool `tfsdk:"permissions_insights_app_user" force:",omitempty"`
	PermissionsInsightsAppAdmin               types.Bool `tfsdk:"permissions_insights_app_admin" force:",omitempty"`
	PermissionsInsightsAppEltEditor           types.Bool `tfsdk:"permissions_insights_app_elt_editor" force:",omitempty"`
	PermissionsInsightsAppUploadUser          types.Bool `tfsdk:"permissions_insights_app_upload_user" force:",omitempty"`
	PermissionsInsightsCreateApplication      types.Bool `tfsdk:"permissions_insights_create_application" force:",omitempty"`
	PermissionsLightningExperienceUser        types.Bool `tfsdk:"permissions_lightning_experience_user" force:",omitempty"`
	PermissionsViewDataLeakageEvents          types.Bool `tfsdk:"permissions_view_data_leakage_events" force:",omitempty"`
	PermissionsConfigCustomRecs               types.Bool `tfsdk:"permissions_config_custom_recs" force:",omitempty"`
	PermissionsSubmitMacrosAllowed            types.Bool `tfsdk:"permissions_submit_macros_allowed" force:",omitempty"`
	PermissionsBulkMacrosAllowed              types.Bool `tfsdk:"permissions_bulk_macros_allowed" force:",omitempty"`
	PermissionsShareInternalArticles          types.Bool `tfsdk:"permissions_share_internal_articles" force:",omitempty"`
	PermissionsManageSessionPermissionSets    types.Bool `tfsdk:"permissions_manage_session_permission_sets" force:",omitempty"`
	PermissionsManageTemplatedApp             types.Bool `tfsdk:"permissions_manage_templated_app" force:",omitempty"`
	PermissionsUseTemplatedApp                types.Bool `tfsdk:"permissions_use_templated_app" force:",omitempty"`
	PermissionsSendAnnouncementEmails         types.Bool `tfsdk:"permissions_send_announcement_emails" force:",omitempty"`
	PermissionsChatterEditOwnPost             types.Bool `tfsdk:"permissions_chatter_edit_own_post" force:",omitempty"`
	PermissionsChatterEditOwnRecordPost       types.Bool `tfsdk:"permissions_chatter_edit_own_record_post" force:",omitempty"`
	PermissionsSalesAnalyticsUser             types.Bool `tfsdk:"permissions_sales_analytics_user" force:",omitempty"`
	PermissionsServiceAnalyticsUser           types.Bool `tfsdk:"permissions_service_analytics_user" force:",omitempty"`
	PermissionsWaveTabularDownload            types.Bool `tfsdk:"permissions_wave_tabular_download" force:",omitempty"`
	PermissionsAutomaticActivityCapture       types.Bool `tfsdk:"permissions_automatic_activity_capture" force:",omitempty"`
	PermissionsImportCustomObjects            types.Bool `tfsdk:"permissions_import_custom_objects" force:",omitempty"`
	PermissionsDelegatedTwoFactor             types.Bool `tfsdk:"permissions_delegated_two_factor" force:",omitempty"`
	PermissionsChatterComposeUiCodesnippet    types.Bool `tfsdk:"permissions_chatter_compose_ui_codesnippet" force:",omitempty"`
	PermissionsSelectFilesFromSalesforce      types.Bool `tfsdk:"permissions_select_files_from_salesforce" force:",omitempty"`
	PermissionsModerateNetworkUsers           types.Bool `tfsdk:"permissions_moderate_network_users" force:",omitempty"`
	PermissionsMergeTopics                    types.Bool `tfsdk:"permissions_merge_topics" force:",omitempty"`
	PermissionsSubscribeToLightningReports    types.Bool `tfsdk:"permissions_subscribe_to_lightning_reports" force:",omitempty"`
	PermissionsManagePvtRptsAndDashbds        types.Bool `tfsdk:"permissions_manage_pvt_rpts_and_dashbds" force:",omitempty"`
	PermissionsAllowLightningLogin            types.Bool `tfsdk:"permissions_allow_lightning_login" force:",omitempty"`
	PermissionsCampaignInfluence2             types.Bool `tfsdk:"permissions_campaign_influence2" force:",omitempty"`
	PermissionsViewDataAssessment             types.Bool `tfsdk:"permissions_view_data_assessment" force:",omitempty"`
	PermissionsRemoveDirectMessageMembers     types.Bool `tfsdk:"permissions_remove_direct_message_members" force:",omitempty"`
	PermissionsCanApproveFeedPost             types.Bool `tfsdk:"permissions_can_approve_feed_post" force:",omitempty"`
	PermissionsAddDirectMessageMembers        types.Bool `tfsdk:"permissions_add_direct_message_members" force:",omitempty"`
	PermissionsAllowViewEditConvertedLeads    types.Bool `tfsdk:"permissions_allow_view_edit_converted_leads" force:",omitempty"`
	PermissionsShowCompanyNameAsUserBadge     types.Bool `tfsdk:"permissions_show_company_name_as_user_badge" force:",omitempty"`
	PermissionsAccessCMC                      types.Bool `tfsdk:"permissions_access_cmc" force:",omitempty"`
	PermissionsViewHealthCheck                types.Bool `tfsdk:"permissions_view_health_check" force:",omitempty"`
	PermissionsManageHealthCheck              types.Bool `tfsdk:"permissions_manage_health_check" force:",omitempty"`
	PermissionsPackaging2                     types.Bool `tfsdk:"permissions_packaging2" force:",omitempty"`
	PermissionsManageCertificates             types.Bool `tfsdk:"permissions_manage_certificates" force:",omitempty"`
	PermissionsCreateReportInLightning        types.Bool `tfsdk:"permissions_create_report_in_lightning" force:",omitempty"`
	PermissionsPreventClassicExperience       types.Bool `tfsdk:"permissions_prevent_classic_experience" force:",omitempty"`
	PermissionsHideReadByList                 types.Bool `tfsdk:"permissions_hide_read_by_list" force:",omitempty"`
	PermissionsUseSmartDataDiscovery          types.Bool `tfsdk:"permissions_use_smart_data_discovery" force:",omitempty"`
	PermissionsGetSmartDataDiscovery          types.Bool `tfsdk:"permissions_get_smart_data_discovery" force:",omitempty"`
	PermissionsCreateUpdateSDDDataset         types.Bool `tfsdk:"permissions_create_update_sdd_dataset" force:",omitempty"`
	PermissionsCreateUpdateSDDStory           types.Bool `tfsdk:"permissions_create_update_sdd_story" force:",omitempty"`
	PermissionsManageSmartDataDiscovery       types.Bool `tfsdk:"permissions_manage_smart_data_discovery" force:",omitempty"`
	PermissionsShareSmartDataDiscoveryStory   types.Bool `tfsdk:"permissions_share_smart_data_discovery_story" force:",omitempty"`
	PermissionsManageSmartDataDiscoveryModel  types.Bool `tfsdk:"permissions_manage_smart_data_discovery_model" force:",omitempty"`
	PermissionsListEmailSend                  types.Bool `tfsdk:"permissions_list_email_send" force:",omitempty"`
	PermissionsFeedPinning                    types.Bool `tfsdk:"permissions_feed_pinning" force:",omitempty"`
	PermissionsChangeDashboardColors          types.Bool `tfsdk:"permissions_change_dashboard_colors" force:",omitempty"`
	PermissionsIotUser                        types.Bool `tfsdk:"permissions_iot_user" force:",omitempty"`
	PermissionsManageRecommendationStrategies types.Bool `tfsdk:"permissions_manage_recommendation_strategies" force:",omitempty"`
	PermissionsManagePropositions             types.Bool `tfsdk:"permissions_manage_propositions" force:",omitempty"`
	PermissionsCloseConversations             types.Bool `tfsdk:"permissions_close_conversations" force:",omitempty"`
	PermissionsSubscribeReportRolesGrps       types.Bool `tfsdk:"permissions_subscribe_report_roles_grps" force:",omitempty"`
	PermissionsSubscribeDashboardRolesGrps    types.Bool `tfsdk:"permissions_subscribe_dashboard_roles_grps" force:",omitempty"`
	PermissionsUseWebLink                     types.Bool `tfsdk:"permissions_use_web_link" force:",omitempty"`
	PermissionsHasUnlimitedNBAExecutions      types.Bool `tfsdk:"permissions_has_unlimited_nba_executions" force:",omitempty"`
	PermissionsViewOnlyEmbeddedAppUser        types.Bool `tfsdk:"permissions_view_only_embedded_app_user" force:",omitempty"`
	PermissionsAdoptionAnalyticsUser          types.Bool `tfsdk:"permissions_adoption_analytics_user" force:",omitempty"`
	PermissionsViewAllActivities              types.Bool `tfsdk:"permissions_view_all_activities" force:",omitempty"`
	PermissionsSubscribeReportToOtherUsers    types.Bool `tfsdk:"permissions_subscribe_report_to_other_users" force:",omitempty"`
	PermissionsLightningConsoleAllowedForUser types.Bool `tfsdk:"permissions_lightning_console_allowed_for_user" force:",omitempty"`
	PermissionsSubscribeReportsRunAsUser      types.Bool `tfsdk:"permissions_subscribe_reports_run_as_user" force:",omitempty"`
	PermissionsSubscribeToLightningDashboards types.Bool `tfsdk:"permissions_subscribe_to_lightning_dashboards" force:",omitempty"`
	PermissionsSubscribeDashboardToOtherUsers types.Bool `tfsdk:"permissions_subscribe_dashboard_to_other_users" force:",omitempty"`
	PermissionsCreateLtngTempInPub            types.Bool `tfsdk:"permissions_create_ltng_temp_in_pub" force:",omitempty"`
	PermissionsTransactionalEmailSend         types.Bool `tfsdk:"permissions_transactional_email_send" force:",omitempty"`
	PermissionsViewPrivateStaticResources     types.Bool `tfsdk:"permissions_view_private_static_resources" force:",omitempty"`
	PermissionsCreateLtngTempFolder           types.Bool `tfsdk:"permissions_create_ltng_temp_folder" force:",omitempty"`
	PermissionsApexRestServices               types.Bool `tfsdk:"permissions_apex_rest_services" force:",omitempty"`
	PermissionsEnableCommunityAppLauncher     types.Bool `tfsdk:"permissions_enable_community_app_launcher" force:",omitempty"`
	PermissionsGiveRecognitionBadge           types.Bool `tfsdk:"permissions_give_recognition_badge" force:",omitempty"`
	PermissionsCanRunAnalysis                 types.Bool `tfsdk:"permissions_can_run_analysis" force:",omitempty"`
	PermissionsUseMySearch                    types.Bool `tfsdk:"permissions_use_my_search" force:",omitempty"`
	PermissionsLtngPromoReserved01UserPerm    types.Bool `tfsdk:"permissions_ltng_promo_reserved01_user_perm" force:",omitempty"`
	PermissionsManageSubscriptions            types.Bool `tfsdk:"permissions_manage_subscriptions" force:",omitempty"`
	PermissionsWaveManagePrivateAssetsUser    types.Bool `tfsdk:"permissions_wave_manage_private_assets_user" force:",omitempty"`
	PermissionsCanEditDataPrepRecipe          types.Bool `tfsdk:"permissions_can_edit_data_prep_recipe" force:",omitempty"`
	PermissionsAddAnalyticsRemoteConnections  types.Bool `tfsdk:"permissions_add_analytics_remote_connections" force:",omitempty"`
	PermissionsManageSurveys                  types.Bool `tfsdk:"permissions_manage_surveys" force:",omitempty"`
	PermissionsUseAssistantDialog             types.Bool `tfsdk:"permissions_use_assistant_dialog" force:",omitempty"`
	PermissionsUseQuerySuggestions            types.Bool `tfsdk:"permissions_use_query_suggestions" force:",omitempty"`
	PermissionsRecordVisibilityAPI            types.Bool `tfsdk:"permissions_record_visibility_api" force:",omitempty"`
	PermissionsViewRoles                      types.Bool `tfsdk:"permissions_view_roles" force:",omitempty"`
	PermissionsSmartDataDiscoveryForCommunity types.Bool `tfsdk:"permissions_smart_data_discovery_for_community" force:",omitempty"`
	PermissionsCanManageMaps                  types.Bool `tfsdk:"permissions_can_manage_maps" force:",omitempty"`
	PermissionsStoryOnDSWithPredicate         types.Bool `tfsdk:"permissions_story_on_ds_with_predicate" force:",omitempty"`
	PermissionsLMOutboundMessagingUserPerm    types.Bool `tfsdk:"permissions_lm_outbound_messaging_user_perm" force:",omitempty"`
	PermissionsModifyDataClassification       types.Bool `tfsdk:"permissions_modify_data_classification" force:",omitempty"`
	PermissionsPrivacyDataAccess              types.Bool `tfsdk:"permissions_privacy_data_access" force:",omitempty"`
	PermissionsQueryAllFiles                  types.Bool `tfsdk:"permissions_query_all_files" force:",omitempty"`
	PermissionsModifyMetadata                 types.Bool `tfsdk:"permissions_modify_metadata" force:",omitempty"`
	PermissionsManageCMS                      types.Bool `tfsdk:"permissions_manage_cms" force:",omitempty"`
	PermissionsSandboxTestingInCommunityApp   types.Bool `tfsdk:"permissions_sandbox_testing_in_community_app" force:",omitempty"`
	PermissionsCanEditPrompts                 types.Bool `tfsdk:"permissions_can_edit_prompts" force:",omitempty"`
	PermissionsViewUserPII                    types.Bool `tfsdk:"permissions_view_user_pii" force:",omitempty"`
	PermissionsManageHubConnections           types.Bool `tfsdk:"permissions_manage_hub_connections" force:",omitempty"`
	PermissionsB2BMarketingAnalyticsUser      types.Bool `tfsdk:"permissions_b_2_b_marketing_analytics_user" force:",omitempty"`
	PermissionsTraceXdsQueries                types.Bool `tfsdk:"permissions_trace_xds_queries" force:",omitempty"`
	PermissionsViewSecurityCommandCenter      types.Bool `tfsdk:"permissions_view_security_command_center" force:",omitempty"`
	PermissionsManageSecurityCommandCenter    types.Bool `tfsdk:"permissions_manage_security_command_center" force:",omitempty"`
	PermissionsViewAllCustomSettings          types.Bool `tfsdk:"permissions_view_all_custom_settings" force:",omitempty"`
	PermissionsViewAllForeignKeyNames         types.Bool `tfsdk:"permissions_view_all_foreign_key_names" force:",omitempty"`
	PermissionsAllowSurveyAdvancedFeatures    types.Bool `tfsdk:"permissions_allow_survey_advanced_features" force:",omitempty"`
	PermissionsAddWaveNotificationRecipients  types.Bool `tfsdk:"permissions_add_wave_notification_recipients" force:",omitempty"`
	PermissionsHeadlessCMSAccess              types.Bool `tfsdk:"permissions_headless_cms_access" force:",omitempty"`
	PermissionsLMEndMessagingSessionUserPerm  types.Bool `tfsdk:"permissions_lm_end_messaging_session_user_perm" force:",omitempty"`
	PermissionsConsentApiUpdate               types.Bool `tfsdk:"permissions_consent_api_update" force:",omitempty"`
	PermissionsPaymentsAPIUser                types.Bool `tfsdk:"permissions_payments_api_user" force:",omitempty"`
	PermissionsAccessContentBuilder           types.Bool `tfsdk:"permissions_access_content_builder" force:",omitempty"`
	PermissionsAccountSwitcherUser            types.Bool `tfsdk:"permissions_account_switcher_user" force:",omitempty"`
	PermissionsViewAnomalyEvents              types.Bool `tfsdk:"permissions_view_anomaly_events" force:",omitempty"`
	PermissionsManageC360AConnections         types.Bool `tfsdk:"permissions_manage_c_360_a_connections" force:",omitempty"`
	PermissionsManageReleaseUpdates           types.Bool `tfsdk:"permissions_manage_release_updates" force:",omitempty"`
	PermissionsViewAllProfiles                types.Bool `tfsdk:"permissions_view_all_profiles" force:",omitempty"`
	PermissionsSkipIdentityConfirmation       types.Bool `tfsdk:"permissions_skip_identity_confirmation" force:",omitempty"`
	PermissionsLearningManager                types.Bool `tfsdk:"permissions_learning_manager" force:",omitempty"`
	PermissionsSendCustomNotifications        types.Bool `tfsdk:"permissions_send_custom_notifications" force:",omitempty"`
	PermissionsPackaging2Delete               types.Bool `tfsdk:"permissions_packaging2_delete" force:",omitempty"`
	PermissionsAutonomousAnalyticsPrivacy     types.Bool `tfsdk:"permissions_autonomous_analytics_privacy" force:",omitempty"`
	PermissionsSonicConsumer                  types.Bool `tfsdk:"permissions_sonic_consumer" force:",omitempty"`
	PermissionsBotManageBots                  types.Bool `tfsdk:"permissions_bot_manage_bots" force:",omitempty"`
	PermissionsBotManageBotsTrainingData      types.Bool `tfsdk:"permissions_bot_manage_bots_training_data" force:",omitempty"`
	PermissionsManageLearningReporting        types.Bool `tfsdk:"permissions_manage_learning_reporting" force:",omitempty"`
	PermissionsIsotopeCToCUser                types.Bool `tfsdk:"permissions_isotope_c_to_c_user" force:",omitempty"`
	PermissionsIsotopeAccess                  types.Bool `tfsdk:"permissions_isotope_access" force:",omitempty"`
	PermissionsIsotopeLEX                     types.Bool `tfsdk:"permissions_isotope_lex" force:",omitempty"`
	PermissionsGetSmartDataDiscoveryExternal  types.Bool `tfsdk:"permissions_get_smart_data_discovery_external" force:",omitempty"`
	PermissionsQuipMetricsAccess              types.Bool `tfsdk:"permissions_quip_metrics_access" force:",omitempty"`
	PermissionsQuipUserEngagementMetrics      types.Bool `tfsdk:"permissions_quip_user_engagement_metrics" force:",omitempty"`
	PermissionsManageExternalConnections      types.Bool `tfsdk:"permissions_manage_external_connections" force:",omitempty"`
	PermissionsUseSubscriptionEmails          types.Bool `tfsdk:"permissions_use_subscription_emails" force:",omitempty"`
	PermissionsAIViewInsightObjects           types.Bool `tfsdk:"permissions_ai_view_insight_objects" force:",omitempty"`
	PermissionsAICreateInsightObjects         types.Bool `tfsdk:"permissions_ai_create_insight_objects" force:",omitempty"`
	PermissionsLifecycleManagementAPIUser     types.Bool `tfsdk:"permissions_lifecycle_management_api_user" force:",omitempty"`
	PermissionsNativeWebviewScrolling         types.Bool `tfsdk:"permissions_native_webview_scrolling" force:",omitempty"`
	PermissionsViewDeveloperName              types.Bool `tfsdk:"permissions_view_developer_name" force:",omitempty"`
}

// In order to support json omitempty for boolean fields, a new type has to be copied into
type profileResourceDataJSON struct {
	Name          string  `force:",omitempty"`
	Description   *string `force:",omitempty"`
	UserLicenseId string  `force:",omitempty"`
	// permissions
	PermissionsEmailSingle                    *bool `force:",omitempty"`
	PermissionsEmailMass                      *bool `force:",omitempty"`
	PermissionsEditTask                       *bool `force:",omitempty"`
	PermissionsEditEvent                      *bool `force:",omitempty"`
	PermissionsExportReport                   *bool `force:",omitempty"`
	PermissionsImportPersonal                 *bool `force:",omitempty"`
	PermissionsDataExport                     *bool `force:",omitempty"`
	PermissionsManageUsers                    *bool `force:",omitempty"`
	PermissionsEditPublicFilters              *bool `force:",omitempty"`
	PermissionsEditPublicTemplates            *bool `force:",omitempty"`
	PermissionsModifyAllData                  *bool `force:",omitempty"`
	PermissionsManageCases                    *bool `force:",omitempty"`
	PermissionsMassInlineEdit                 *bool `force:",omitempty"`
	PermissionsEditKnowledge                  *bool `force:",omitempty"`
	PermissionsManageKnowledge                *bool `force:",omitempty"`
	PermissionsManageSolutions                *bool `force:",omitempty"`
	PermissionsCustomizeApplication           *bool `force:",omitempty"`
	PermissionsEditReadonlyFields             *bool `force:",omitempty"`
	PermissionsRunReports                     *bool `force:",omitempty"`
	PermissionsViewSetup                      *bool `force:",omitempty"`
	PermissionsTransferAnyEntity              *bool `force:",omitempty"`
	PermissionsNewReportBuilder               *bool `force:",omitempty"`
	PermissionsActivateContract               *bool `force:",omitempty"`
	PermissionsActivateOrder                  *bool `force:",omitempty"`
	PermissionsImportLeads                    *bool `force:",omitempty"`
	PermissionsManageLeads                    *bool `force:",omitempty"`
	PermissionsTransferAnyLead                *bool `force:",omitempty"`
	PermissionsViewAllData                    *bool `force:",omitempty"`
	PermissionsEditPublicDocuments            *bool `force:",omitempty"`
	PermissionsViewEncryptedData              *bool `force:",omitempty"`
	PermissionsEditBrandTemplates             *bool `force:",omitempty"`
	PermissionsEditHtmlTemplates              *bool `force:",omitempty"`
	PermissionsChatterInternalUser            *bool `force:",omitempty"`
	PermissionsManageEncryptionKeys           *bool `force:",omitempty"`
	PermissionsDeleteActivatedContract        *bool `force:",omitempty"`
	PermissionsChatterInviteExternalUsers     *bool `force:",omitempty"`
	PermissionsSendSitRequests                *bool `force:",omitempty"`
	PermissionsManageRemoteAccess             *bool `force:",omitempty"`
	PermissionsCanUseNewDashboardBuilder      *bool `force:",omitempty"`
	PermissionsManageCategories               *bool `force:",omitempty"`
	PermissionsConvertLeads                   *bool `force:",omitempty"`
	PermissionsPasswordNeverExpires           *bool `force:",omitempty"`
	PermissionsUseTeamReassignWizards         *bool `force:",omitempty"`
	PermissionsEditActivatedOrders            *bool `force:",omitempty"`
	PermissionsInstallMultiforce              *bool `force:",omitempty"`
	PermissionsPublishMultiforce              *bool `force:",omitempty"`
	PermissionsChatterOwnGroups               *bool `force:",omitempty"`
	PermissionsEditOppLineItemUnitPrice       *bool `force:",omitempty"`
	PermissionsCreateMultiforce               *bool `force:",omitempty"`
	PermissionsBulkApiHardDelete              *bool `force:",omitempty"`
	PermissionsSolutionImport                 *bool `force:",omitempty"`
	PermissionsManageCallCenters              *bool `force:",omitempty"`
	PermissionsManageSynonyms                 *bool `force:",omitempty"`
	PermissionsViewContent                    *bool `force:",omitempty"`
	PermissionsManageEmailClientConfig        *bool `force:",omitempty"`
	PermissionsEnableNotifications            *bool `force:",omitempty"`
	PermissionsManageDataIntegrations         *bool `force:",omitempty"`
	PermissionsDistributeFromPersWksp         *bool `force:",omitempty"`
	PermissionsViewDataCategories             *bool `force:",omitempty"`
	PermissionsManageDataCategories           *bool `force:",omitempty"`
	PermissionsAuthorApex                     *bool `force:",omitempty"`
	PermissionsManageMobile                   *bool `force:",omitempty"`
	PermissionsApiEnabled                     *bool `force:",omitempty"`
	PermissionsManageCustomReportTypes        *bool `force:",omitempty"`
	PermissionsEditCaseComments               *bool `force:",omitempty"`
	PermissionsTransferAnyCase                *bool `force:",omitempty"`
	PermissionsContentAdministrator           *bool `force:",omitempty"`
	PermissionsCreateWorkspaces               *bool `force:",omitempty"`
	PermissionsManageContentPermissions       *bool `force:",omitempty"`
	PermissionsManageContentProperties        *bool `force:",omitempty"`
	PermissionsManageContentTypes             *bool `force:",omitempty"`
	PermissionsManageExchangeConfig           *bool `force:",omitempty"`
	PermissionsManageAnalyticSnapshots        *bool `force:",omitempty"`
	PermissionsScheduleReports                *bool `force:",omitempty"`
	PermissionsManageBusinessHourHolidays     *bool `force:",omitempty"`
	PermissionsManageEntitlements             *bool `force:",omitempty"`
	PermissionsManageDynamicDashboards        *bool `force:",omitempty"`
	PermissionsCustomSidebarOnAllPages        *bool `force:",omitempty"`
	PermissionsManageInteraction              *bool `force:",omitempty"`
	PermissionsViewMyTeamsDashboards          *bool `force:",omitempty"`
	PermissionsModerateChatter                *bool `force:",omitempty"`
	PermissionsResetPasswords                 *bool `force:",omitempty"`
	PermissionsFlowUFLRequired                *bool `force:",omitempty"`
	PermissionsCanInsertFeedSystemFields      *bool `force:",omitempty"`
	PermissionsActivitiesAccess               *bool `force:",omitempty"`
	PermissionsManageKnowledgeImportExport    *bool `force:",omitempty"`
	PermissionsEmailTemplateManagement        *bool `force:",omitempty"`
	PermissionsEmailAdministration            *bool `force:",omitempty"`
	PermissionsManageChatterMessages          *bool `force:",omitempty"`
	PermissionsAllowEmailIC                   *bool `force:",omitempty"`
	PermissionsChatterFileLink                *bool `force:",omitempty"`
	PermissionsForceTwoFactor                 *bool `force:",omitempty"`
	PermissionsViewEventLogFiles              *bool `force:",omitempty"`
	PermissionsManageNetworks                 *bool `force:",omitempty"`
	PermissionsManageAuthProviders            *bool `force:",omitempty"`
	PermissionsRunFlow                        *bool `force:",omitempty"`
	PermissionsCreateCustomizeDashboards      *bool `force:",omitempty"`
	PermissionsCreateDashboardFolders         *bool `force:",omitempty"`
	PermissionsViewPublicDashboards           *bool `force:",omitempty"`
	PermissionsManageDashbdsInPubFolders      *bool `force:",omitempty"`
	PermissionsCreateCustomizeReports         *bool `force:",omitempty"`
	PermissionsCreateReportFolders            *bool `force:",omitempty"`
	PermissionsViewPublicReports              *bool `force:",omitempty"`
	PermissionsManageReportsInPubFolders      *bool `force:",omitempty"`
	PermissionsEditMyDashboards               *bool `force:",omitempty"`
	PermissionsEditMyReports                  *bool `force:",omitempty"`
	PermissionsViewAllUsers                   *bool `force:",omitempty"`
	PermissionsAllowUniversalSearch           *bool `force:",omitempty"`
	PermissionsConnectOrgToEnvironmentHub     *bool `force:",omitempty"`
	PermissionsWorkCalibrationUser            *bool `force:",omitempty"`
	PermissionsCreateCustomizeFilters         *bool `force:",omitempty"`
	PermissionsWorkDotComUserPerm             *bool `force:",omitempty"`
	PermissionsContentHubUser                 *bool `force:",omitempty"`
	PermissionsGovernNetworks                 *bool `force:",omitempty"`
	PermissionsSalesConsole                   *bool `force:",omitempty"`
	PermissionsTwoFactorApi                   *bool `force:",omitempty"`
	PermissionsDeleteTopics                   *bool `force:",omitempty"`
	PermissionsEditTopics                     *bool `force:",omitempty"`
	PermissionsCreateTopics                   *bool `force:",omitempty"`
	PermissionsAssignTopics                   *bool `force:",omitempty"`
	PermissionsIdentityEnabled                *bool `force:",omitempty"`
	PermissionsIdentityConnect                *bool `force:",omitempty"`
	PermissionsAllowViewKnowledge             *bool `force:",omitempty"`
	PermissionsContentWorkspaces              *bool `force:",omitempty"`
	PermissionsCreateWorkBadgeDefinition      *bool `force:",omitempty"`
	PermissionsManageSearchPromotionRules     *bool `force:",omitempty"`
	PermissionsCustomMobileAppsAccess         *bool `force:",omitempty"`
	PermissionsViewHelpLink                   *bool `force:",omitempty"`
	PermissionsManageProfilesPermissionsets   *bool `force:",omitempty"`
	PermissionsAssignPermissionSets           *bool `force:",omitempty"`
	PermissionsManageRoles                    *bool `force:",omitempty"`
	PermissionsManageIpAddresses              *bool `force:",omitempty"`
	PermissionsManageSharing                  *bool `force:",omitempty"`
	PermissionsManageInternalUsers            *bool `force:",omitempty"`
	PermissionsManagePasswordPolicies         *bool `force:",omitempty"`
	PermissionsManageLoginAccessPolicies      *bool `force:",omitempty"`
	PermissionsViewPlatformEvents             *bool `force:",omitempty"`
	PermissionsManageCustomPermissions        *bool `force:",omitempty"`
	PermissionsCanVerifyComment               *bool `force:",omitempty"`
	PermissionsManageUnlistedGroups           *bool `force:",omitempty"`
	PermissionsStdAutomaticActivityCapture    *bool `force:",omitempty"`
	PermissionsInsightsAppDashboardEditor     *bool `force:",omitempty"`
	PermissionsManageTwoFactor                *bool `force:",omitempty"`
	PermissionsInsightsAppUser                *bool `force:",omitempty"`
	PermissionsInsightsAppAdmin               *bool `force:",omitempty"`
	PermissionsInsightsAppEltEditor           *bool `force:",omitempty"`
	PermissionsInsightsAppUploadUser          *bool `force:",omitempty"`
	PermissionsInsightsCreateApplication      *bool `force:",omitempty"`
	PermissionsLightningExperienceUser        *bool `force:",omitempty"`
	PermissionsViewDataLeakageEvents          *bool `force:",omitempty"`
	PermissionsConfigCustomRecs               *bool `force:",omitempty"`
	PermissionsSubmitMacrosAllowed            *bool `force:",omitempty"`
	PermissionsBulkMacrosAllowed              *bool `force:",omitempty"`
	PermissionsShareInternalArticles          *bool `force:",omitempty"`
	PermissionsManageSessionPermissionSets    *bool `force:",omitempty"`
	PermissionsManageTemplatedApp             *bool `force:",omitempty"`
	PermissionsUseTemplatedApp                *bool `force:",omitempty"`
	PermissionsSendAnnouncementEmails         *bool `force:",omitempty"`
	PermissionsChatterEditOwnPost             *bool `force:",omitempty"`
	PermissionsChatterEditOwnRecordPost       *bool `force:",omitempty"`
	PermissionsSalesAnalyticsUser             *bool `force:",omitempty"`
	PermissionsServiceAnalyticsUser           *bool `force:",omitempty"`
	PermissionsWaveTabularDownload            *bool `force:",omitempty"`
	PermissionsAutomaticActivityCapture       *bool `force:",omitempty"`
	PermissionsImportCustomObjects            *bool `force:",omitempty"`
	PermissionsDelegatedTwoFactor             *bool `force:",omitempty"`
	PermissionsChatterComposeUiCodesnippet    *bool `force:",omitempty"`
	PermissionsSelectFilesFromSalesforce      *bool `force:",omitempty"`
	PermissionsModerateNetworkUsers           *bool `force:",omitempty"`
	PermissionsMergeTopics                    *bool `force:",omitempty"`
	PermissionsSubscribeToLightningReports    *bool `force:",omitempty"`
	PermissionsManagePvtRptsAndDashbds        *bool `force:",omitempty"`
	PermissionsAllowLightningLogin            *bool `force:",omitempty"`
	PermissionsCampaignInfluence2             *bool `force:",omitempty"`
	PermissionsViewDataAssessment             *bool `force:",omitempty"`
	PermissionsRemoveDirectMessageMembers     *bool `force:",omitempty"`
	PermissionsCanApproveFeedPost             *bool `force:",omitempty"`
	PermissionsAddDirectMessageMembers        *bool `force:",omitempty"`
	PermissionsAllowViewEditConvertedLeads    *bool `force:",omitempty"`
	PermissionsShowCompanyNameAsUserBadge     *bool `force:",omitempty"`
	PermissionsAccessCMC                      *bool `force:",omitempty"`
	PermissionsViewHealthCheck                *bool `force:",omitempty"`
	PermissionsManageHealthCheck              *bool `force:",omitempty"`
	PermissionsPackaging2                     *bool `force:",omitempty"`
	PermissionsManageCertificates             *bool `force:",omitempty"`
	PermissionsCreateReportInLightning        *bool `force:",omitempty"`
	PermissionsPreventClassicExperience       *bool `force:",omitempty"`
	PermissionsHideReadByList                 *bool `force:",omitempty"`
	PermissionsUseSmartDataDiscovery          *bool `force:",omitempty"`
	PermissionsGetSmartDataDiscovery          *bool `force:",omitempty"`
	PermissionsCreateUpdateSDDDataset         *bool `force:",omitempty"`
	PermissionsCreateUpdateSDDStory           *bool `force:",omitempty"`
	PermissionsManageSmartDataDiscovery       *bool `force:",omitempty"`
	PermissionsShareSmartDataDiscoveryStory   *bool `force:",omitempty"`
	PermissionsManageSmartDataDiscoveryModel  *bool `force:",omitempty"`
	PermissionsListEmailSend                  *bool `force:",omitempty"`
	PermissionsFeedPinning                    *bool `force:",omitempty"`
	PermissionsChangeDashboardColors          *bool `force:",omitempty"`
	PermissionsIotUser                        *bool `force:",omitempty"`
	PermissionsManageRecommendationStrategies *bool `force:",omitempty"`
	PermissionsManagePropositions             *bool `force:",omitempty"`
	PermissionsCloseConversations             *bool `force:",omitempty"`
	PermissionsSubscribeReportRolesGrps       *bool `force:",omitempty"`
	PermissionsSubscribeDashboardRolesGrps    *bool `force:",omitempty"`
	PermissionsUseWebLink                     *bool `force:",omitempty"`
	PermissionsHasUnlimitedNBAExecutions      *bool `force:",omitempty"`
	PermissionsViewOnlyEmbeddedAppUser        *bool `force:",omitempty"`
	PermissionsAdoptionAnalyticsUser          *bool `force:",omitempty"`
	PermissionsViewAllActivities              *bool `force:",omitempty"`
	PermissionsSubscribeReportToOtherUsers    *bool `force:",omitempty"`
	PermissionsLightningConsoleAllowedForUser *bool `force:",omitempty"`
	PermissionsSubscribeReportsRunAsUser      *bool `force:",omitempty"`
	PermissionsSubscribeToLightningDashboards *bool `force:",omitempty"`
	PermissionsSubscribeDashboardToOtherUsers *bool `force:",omitempty"`
	PermissionsCreateLtngTempInPub            *bool `force:",omitempty"`
	PermissionsTransactionalEmailSend         *bool `force:",omitempty"`
	PermissionsViewPrivateStaticResources     *bool `force:",omitempty"`
	PermissionsCreateLtngTempFolder           *bool `force:",omitempty"`
	PermissionsApexRestServices               *bool `force:",omitempty"`
	PermissionsEnableCommunityAppLauncher     *bool `force:",omitempty"`
	PermissionsGiveRecognitionBadge           *bool `force:",omitempty"`
	PermissionsCanRunAnalysis                 *bool `force:",omitempty"`
	PermissionsUseMySearch                    *bool `force:",omitempty"`
	PermissionsLtngPromoReserved01UserPerm    *bool `force:",omitempty"`
	PermissionsManageSubscriptions            *bool `force:",omitempty"`
	PermissionsWaveManagePrivateAssetsUser    *bool `force:",omitempty"`
	PermissionsCanEditDataPrepRecipe          *bool `force:",omitempty"`
	PermissionsAddAnalyticsRemoteConnections  *bool `force:",omitempty"`
	PermissionsManageSurveys                  *bool `force:",omitempty"`
	PermissionsUseAssistantDialog             *bool `force:",omitempty"`
	PermissionsUseQuerySuggestions            *bool `force:",omitempty"`
	PermissionsRecordVisibilityAPI            *bool `force:",omitempty"`
	PermissionsViewRoles                      *bool `force:",omitempty"`
	PermissionsSmartDataDiscoveryForCommunity *bool `force:",omitempty"`
	PermissionsCanManageMaps                  *bool `force:",omitempty"`
	PermissionsStoryOnDSWithPredicate         *bool `force:",omitempty"`
	PermissionsLMOutboundMessagingUserPerm    *bool `force:",omitempty"`
	PermissionsModifyDataClassification       *bool `force:",omitempty"`
	PermissionsPrivacyDataAccess              *bool `force:",omitempty"`
	PermissionsQueryAllFiles                  *bool `force:",omitempty"`
	PermissionsModifyMetadata                 *bool `force:",omitempty"`
	PermissionsManageCMS                      *bool `force:",omitempty"`
	PermissionsSandboxTestingInCommunityApp   *bool `force:",omitempty"`
	PermissionsCanEditPrompts                 *bool `force:",omitempty"`
	PermissionsViewUserPII                    *bool `force:",omitempty"`
	PermissionsManageHubConnections           *bool `force:",omitempty"`
	PermissionsB2BMarketingAnalyticsUser      *bool `force:",omitempty"`
	PermissionsTraceXdsQueries                *bool `force:",omitempty"`
	PermissionsViewSecurityCommandCenter      *bool `force:",omitempty"`
	PermissionsManageSecurityCommandCenter    *bool `force:",omitempty"`
	PermissionsViewAllCustomSettings          *bool `force:",omitempty"`
	PermissionsViewAllForeignKeyNames         *bool `force:",omitempty"`
	PermissionsAllowSurveyAdvancedFeatures    *bool `force:",omitempty"`
	PermissionsAddWaveNotificationRecipients  *bool `force:",omitempty"`
	PermissionsHeadlessCMSAccess              *bool `force:",omitempty"`
	PermissionsLMEndMessagingSessionUserPerm  *bool `force:",omitempty"`
	PermissionsConsentApiUpdate               *bool `force:",omitempty"`
	PermissionsPaymentsAPIUser                *bool `force:",omitempty"`
	PermissionsAccessContentBuilder           *bool `force:",omitempty"`
	PermissionsAccountSwitcherUser            *bool `force:",omitempty"`
	PermissionsViewAnomalyEvents              *bool `force:",omitempty"`
	PermissionsManageC360AConnections         *bool `force:",omitempty"`
	PermissionsManageReleaseUpdates           *bool `force:",omitempty"`
	PermissionsViewAllProfiles                *bool `force:",omitempty"`
	PermissionsSkipIdentityConfirmation       *bool `force:",omitempty"`
	PermissionsLearningManager                *bool `force:",omitempty"`
	PermissionsSendCustomNotifications        *bool `force:",omitempty"`
	PermissionsPackaging2Delete               *bool `force:",omitempty"`
	PermissionsAutonomousAnalyticsPrivacy     *bool `force:",omitempty"`
	PermissionsSonicConsumer                  *bool `force:",omitempty"`
	PermissionsBotManageBots                  *bool `force:",omitempty"`
	PermissionsBotManageBotsTrainingData      *bool `force:",omitempty"`
	PermissionsManageLearningReporting        *bool `force:",omitempty"`
	PermissionsIsotopeCToCUser                *bool `force:",omitempty"`
	PermissionsIsotopeAccess                  *bool `force:",omitempty"`
	PermissionsIsotopeLEX                     *bool `force:",omitempty"`
	PermissionsGetSmartDataDiscoveryExternal  *bool `force:",omitempty"`
	PermissionsQuipMetricsAccess              *bool `force:",omitempty"`
	PermissionsQuipUserEngagementMetrics      *bool `force:",omitempty"`
	PermissionsManageExternalConnections      *bool `force:",omitempty"`
	PermissionsUseSubscriptionEmails          *bool `force:",omitempty"`
	PermissionsAIViewInsightObjects           *bool `force:",omitempty"`
	PermissionsAICreateInsightObjects         *bool `force:",omitempty"`
	PermissionsLifecycleManagementAPIUser     *bool `force:",omitempty"`
	PermissionsNativeWebviewScrolling         *bool `force:",omitempty"`
	PermissionsViewDeveloperName              *bool `force:",omitempty"`
}

func (profileResourceDataJSON) ApiName() string {
	return "Profile"
}

func (profileResourceDataJSON) ExternalIdApiName() string {
	return ""
}

func (profileResourceData) ApiName() string {
	return "Profile"
}

func (profileResourceData) ExternalIdApiName() string {
	return ""
}

func (p *profileResourceData) Instance() force.SObject {
	return p
}

func (p *profileResourceData) toJSONData() profileResourceDataJSON {
	jsonBytes, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	var data profileResourceDataJSON
	err = json.Unmarshal(jsonBytes, &data)
	if err != nil {
		panic(err)
	}
	return data
}

func (p *profileResourceData) Insertable() force.SObject {
	return p.toJSONData()
}

func (p *profileResourceData) Updatable() force.SObject {
	updatable := p.toJSONData()
	updatable.UserLicenseId = ""
	return updatable
}

func (p *profileResourceData) GetId() string {
	return p.Id.Value
}

func (p *profileResourceData) SetId(id string) {
	p.Id = types.String{Value: id}
}
