/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"context"
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	litellmv1alpha1 "github.com/bbdsoftware/litellm-operator/api/litellm/v1alpha1"
)

// nolint:unused
// log is for logging in this package.
var litellminstancelog = logf.Log.WithName("litellminstance-resource")

// SetupLiteLLMInstanceWebhookWithManager registers the webhook for LiteLLMInstance in the manager.
func SetupLiteLLMInstanceWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&litellmv1alpha1.LiteLLMInstance{}).
		WithValidator(&LiteLLMInstanceCustomValidator{}).
		Complete()
}

type LiteLLMInstanceCustomDefaulter struct {
	// TODO(user): Add more fields as needed for defaulting
}

// Default implements webhook.CustomDefaulter so a webhook will be registered for the Kind LiteLLMInstance.
func (d *LiteLLMInstanceCustomDefaulter) Default(_ context.Context, obj runtime.Object) error {
	litellminstance, ok := obj.(*litellmv1alpha1.LiteLLMInstance)

	if !ok {
		return fmt.Errorf("expected an LiteLLMInstance object but got %T", obj)
	}
	litellminstancelog.Info("Defaulting for LiteLLMInstance", "name", litellminstance.GetName())

	// Apply default image if not specified
	if litellminstance.Spec.Image == "" {
		litellminstance.Spec.Image = "ghcr.io/berriai/litellm-database:main-v1.74.9.rc.1"
		litellminstancelog.Info("Applied default image", "image", litellminstance.Spec.Image)
	}

	// Apply default values for Ingress if not specified
	if litellminstance.Spec.Ingress.Host == "" {
		litellminstance.Spec.Ingress.Enabled = false
		litellminstancelog.Info("Applied default ingress configuration", "enabled", litellminstance.Spec.Ingress.Enabled)
	}

	// Apply default values for Gateway if not specified
	if litellminstance.Spec.Gateway.Host == "" {
		litellminstance.Spec.Gateway.Enabled = false
		litellminstancelog.Info("Applied default gateway configuration", "enabled", litellminstance.Spec.Gateway.Enabled)
	}

	return nil
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// NOTE: The 'path' attribute must follow a specific pattern and should not be modified directly here.
// Modifying the path for an invalid path can cause API server errors; failing to locate the webhook.
// +kubebuilder:webhook:path=/validate-litellm-litellm-ai-v1alpha1-litellminstance,mutating=false,failurePolicy=fail,sideEffects=None,groups=litellm.litellm.ai,resources=litellminstances,verbs=create;update,versions=v1alpha1,name=vlitellminstance-v1alpha1.kb.io,admissionReviewVersions=v1

// LiteLLMInstanceCustomValidator struct is responsible for validating the LiteLLMInstance resource
// when it is created, updated, or deleted.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as this struct is used only for temporary operations and does not need to be deeply copied.
type LiteLLMInstanceCustomValidator struct {
	// TODO(user): Add more fields as needed for validation
}

var _ webhook.CustomValidator = &LiteLLMInstanceCustomValidator{}

// ValidateCreate implements webhook.CustomValidator so a webhook will be registered for the type LiteLLMInstance.
func (v *LiteLLMInstanceCustomValidator) ValidateCreate(_ context.Context, obj runtime.Object) (admission.Warnings, error) {
	litellminstance, ok := obj.(*litellmv1alpha1.LiteLLMInstance)
	if !ok {
		return nil, fmt.Errorf("expected a LiteLLMInstance object but got %T", obj)
	}
	litellminstancelog.Info("Validation for LiteLLMInstance upon creation", "name", litellminstance.GetName())

	return nil, v.validateLiteLLMInstance(litellminstance)
}

// ValidateUpdate implements webhook.CustomValidator so a webhook will be registered for the type LiteLLMInstance.
func (v *LiteLLMInstanceCustomValidator) ValidateUpdate(_ context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	litellminstance, ok := newObj.(*litellmv1alpha1.LiteLLMInstance)
	if !ok {
		return nil, fmt.Errorf("expected a LiteLLMInstance object for the newObj but got %T", newObj)
	}
	litellminstancelog.Info("Validation for LiteLLMInstance upon update", "name", litellminstance.GetName())

	return nil, v.validateLiteLLMInstance(litellminstance)
}

// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type LiteLLMInstance.
func (v *LiteLLMInstanceCustomValidator) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	litellminstance, ok := obj.(*litellmv1alpha1.LiteLLMInstance)
	if !ok {
		return nil, fmt.Errorf("expected a LiteLLMInstance object but got %T", obj)
	}
	litellminstancelog.Info("Validation for LiteLLMInstance upon deletion", "name", litellminstance.GetName())

	// TODO(user): fill in your validation logic upon object deletion.

	return nil, nil
}

// validateLiteLLMInstance performs comprehensive validation of the LiteLLMInstance
func (v *LiteLLMInstanceCustomValidator) validateLiteLLMInstance(instance *litellmv1alpha1.LiteLLMInstance) error {
	// Validate image
	if err := v.validateImage(instance); err != nil {
		return err
	}

	// Validate masterKey
	if err := v.validateMasterKey(instance); err != nil {
		return err
	}

	// Validate database secret reference
	if err := v.validateDatabaseSecretRef(instance); err != nil {
		return err
	}

	// Validate redis secret reference
	if err := v.validateRedisSecretRef(instance); err != nil {
		return err
	}

	// Validate ingress configuration
	if err := v.validateIngress(instance); err != nil {
		return err
	}

	// Validate gateway configuration
	if err := v.validateGateway(instance); err != nil {
		return err
	}

	// Validate cross-field dependencies
	if err := v.validateCrossFieldDependencies(instance); err != nil {
		return err
	}

	return nil
}

// validateImage validates the image field
func (v *LiteLLMInstanceCustomValidator) validateImage(instance *litellmv1alpha1.LiteLLMInstance) error {
	if instance.Spec.Image == "" {
		return fmt.Errorf("spec.image: image is required")
	}

	// Validate image format (basic check)
	if !strings.Contains(instance.Spec.Image, "/") {
		return fmt.Errorf("spec.image: image must be in format 'registry/repository:tag'")
	}

	// Validate image tag is present
	if !strings.Contains(instance.Spec.Image, ":") {
		return fmt.Errorf("spec.image: image must include a tag")
	}

	return nil
}

// validateMasterKey validates the masterKey field
func (v *LiteLLMInstanceCustomValidator) validateMasterKey(instance *litellmv1alpha1.LiteLLMInstance) error {
	if instance.Spec.MasterKey == "" {
		return fmt.Errorf("spec.masterKey: masterKey is required")
	}

	if len(instance.Spec.MasterKey) < 8 {
		return fmt.Errorf("spec.masterKey: masterKey must be at least 8 characters long")
	}

	// Check for common weak patterns
	if strings.ToLower(instance.Spec.MasterKey) == "masterkey" ||
		strings.ToLower(instance.Spec.MasterKey) == "password" ||
		strings.ToLower(instance.Spec.MasterKey) == "secret" {
		return fmt.Errorf("spec.masterKey: masterKey cannot be a common weak value")
	}

	return nil
}

// validateDatabaseSecretRef validates the databaseSecretRef field
func (v *LiteLLMInstanceCustomValidator) validateDatabaseSecretRef(instance *litellmv1alpha1.LiteLLMInstance) error {
	if instance.Spec.DatabaseSecretRef.NameRef == "" {
		return fmt.Errorf("spec.databaseSecretRef.nameRef: database secret reference is required")
	}

	// Validate secret name format
	if errs := validation.IsDNS1123Subdomain(instance.Spec.DatabaseSecretRef.NameRef); len(errs) > 0 {
		return fmt.Errorf("spec.databaseSecretRef.nameRef: invalid secret name format: %v", errs)
	}

	// Validate required keys
	if instance.Spec.DatabaseSecretRef.Keys.HostSecret == "" {
		return fmt.Errorf("spec.databaseSecretRef.keys.hostSecret: host secret key is required")
	}
	if instance.Spec.DatabaseSecretRef.Keys.PasswordSecret == "" {
		return fmt.Errorf("spec.databaseSecretRef.keys.passwordSecret: password secret key is required")
	}
	if instance.Spec.DatabaseSecretRef.Keys.UsernameSecret == "" {
		return fmt.Errorf("spec.databaseSecretRef.keys.usernameSecret: username secret key is required")
	}
	if instance.Spec.DatabaseSecretRef.Keys.DbnameSecret == "" {
		return fmt.Errorf("spec.databaseSecretRef.keys.dbnameSecret: database name secret key is required")
	}

	return nil
}

// validateRedisSecretRef validates the redisSecretRef field
func (v *LiteLLMInstanceCustomValidator) validateRedisSecretRef(instance *litellmv1alpha1.LiteLLMInstance) error {
	if instance.Spec.RedisSecretRef.NameRef == "" {
		return fmt.Errorf("spec.redisSecretRef.nameRef: redis secret reference is required")
	}

	// Validate secret name format
	if errs := validation.IsDNS1123Subdomain(instance.Spec.RedisSecretRef.NameRef); len(errs) > 0 {
		return fmt.Errorf("spec.redisSecretRef.nameRef: invalid secret name format: %v", errs)
	}

	// Validate required keys
	if instance.Spec.RedisSecretRef.Keys.HostSecret == "" {
		return fmt.Errorf("spec.redisSecretRef.keys.hostSecret: host secret key is required")
	}
	if instance.Spec.RedisSecretRef.Keys.PasswordSecret == "" {
		return fmt.Errorf("spec.redisSecretRef.keys.passwordSecret: password secret key is required")
	}
	if instance.Spec.RedisSecretRef.Keys.PortSecret <= 0 || instance.Spec.RedisSecretRef.Keys.PortSecret > 65535 {
		return fmt.Errorf("spec.redisSecretRef.keys.portSecret: port must be between 1 and 65535")
	}

	return nil
}

// validateIngress validates the ingress configuration
func (v *LiteLLMInstanceCustomValidator) validateIngress(instance *litellmv1alpha1.LiteLLMInstance) error {
	if !instance.Spec.Ingress.Enabled {
		return nil // No validation needed if ingress is disabled
	}

	if instance.Spec.Ingress.Host == "" {
		return fmt.Errorf("spec.ingress.host: host is required when ingress is enabled")
	}

	// Validate host format
	if errs := validation.IsDNS1123Subdomain(instance.Spec.Ingress.Host); len(errs) > 0 {
		return fmt.Errorf("spec.ingress.host: invalid host format: %v", errs)
	}

	// Validate host is not localhost or internal IP
	if strings.Contains(instance.Spec.Ingress.Host, "localhost") ||
		strings.Contains(instance.Spec.Ingress.Host, "127.0.0.1") ||
		strings.Contains(instance.Spec.Ingress.Host, "::1") {
		return fmt.Errorf("spec.ingress.host: host cannot be localhost or internal IP")
	}

	return nil
}

// validateGateway validates the gateway configuration
func (v *LiteLLMInstanceCustomValidator) validateGateway(instance *litellmv1alpha1.LiteLLMInstance) error {
	if !instance.Spec.Gateway.Enabled {
		return nil // No validation needed if gateway is disabled
	}

	if instance.Spec.Gateway.Host == "" {
		return fmt.Errorf("spec.gateway.host: host is required when gateway is enabled")
	}

	// Validate host format
	if errs := validation.IsDNS1123Subdomain(instance.Spec.Gateway.Host); len(errs) > 0 {
		return fmt.Errorf("spec.gateway.host: invalid host format: %v", errs)
	}

	// Validate host is not localhost or internal IP
	if strings.Contains(instance.Spec.Gateway.Host, "localhost") ||
		strings.Contains(instance.Spec.Gateway.Host, "127.0.0.1") ||
		strings.Contains(instance.Spec.Gateway.Host, "::1") {
		return fmt.Errorf("spec.gateway.host: host cannot be localhost or internal IP")
	}

	return nil
}

// validateCrossFieldDependencies validates cross-field dependencies
func (v *LiteLLMInstanceCustomValidator) validateCrossFieldDependencies(instance *litellmv1alpha1.LiteLLMInstance) error {
	// If both ingress and gateway are enabled, they should have different hosts
	if instance.Spec.Ingress.Enabled && instance.Spec.Gateway.Enabled {
		if instance.Spec.Ingress.Host == instance.Spec.Gateway.Host {
			return fmt.Errorf("spec.ingress.host and spec.gateway.host: cannot be the same when both are enabled")
		}
	}

	// Validate that database and redis secret refs are different
	if instance.Spec.DatabaseSecretRef.NameRef == instance.Spec.RedisSecretRef.NameRef {
		return fmt.Errorf("spec.databaseSecretRef.nameRef and spec.redisSecretRef.nameRef: cannot be the same")
	}

	return nil
}
