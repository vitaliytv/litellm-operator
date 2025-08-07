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

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	litellmv1alpha1 "github.com/bbdsoftware/litellm-operator/api/litellm/v1alpha1"
)

var _ = Describe("LiteLLMInstance Webhook", func() {
	var (
		obj       *litellmv1alpha1.LiteLLMInstance
		oldObj    *litellmv1alpha1.LiteLLMInstance
		validator LiteLLMInstanceCustomValidator
		defaulter LiteLLMInstanceCustomDefaulter
		ctx       context.Context
	)

	BeforeEach(func() {
		obj = &litellmv1alpha1.LiteLLMInstance{}
		oldObj = &litellmv1alpha1.LiteLLMInstance{}
		validator = LiteLLMInstanceCustomValidator{}
		defaulter = LiteLLMInstanceCustomDefaulter{}
		ctx = context.Background()

		// Set up a valid base configuration
		obj.Spec = litellmv1alpha1.LiteLLMInstanceSpec{
			Image:     "ghcr.io/berriai/litellm-database:main-v1.74.9.rc.1",
			MasterKey: "valid-master-key-123",
			DatabaseSecretRef: litellmv1alpha1.DatabaseSecretRef{
				NameRef: "database-secret",
				Keys: litellmv1alpha1.DatabaseSecretKeys{
					HostSecret:     "host",
					PasswordSecret: "password",
					UsernameSecret: "username",
					DbnameSecret:   "dbname",
				},
			},
			RedisSecretRef: litellmv1alpha1.RedisSecretRef{
				NameRef: "redis-secret",
				Keys: litellmv1alpha1.RedisSecretKeys{
					HostSecret:     "host",
					PasswordSecret: "password",
					PortSecret:     6379,
				},
			},
			Ingress: litellmv1alpha1.Ingress{
				Enabled: false,
			},
			Gateway: litellmv1alpha1.Gateway{
				Enabled: false,
			},
		}
	})

	Context("When creating LiteLLMInstance under Defaulting Webhook", func() {
		It("Should apply defaults when image is empty", func() {
			By("setting image to empty")
			obj.Spec.Image = ""

			By("calling the Default method to apply defaults")
			err := defaulter.Default(ctx, obj)

			By("checking that no error occurred")
			Expect(err).To(BeNil())

			By("checking that the default image is set")
			Expect(obj.Spec.Image).To(Equal("ghcr.io/berriai/litellm-database:main-v1.74.9.rc.1"))
		})

		It("Should apply defaults for ingress and gateway when not specified", func() {
			By("ensuring ingress and gateway are not set")
			obj.Spec.Ingress.Enabled = false
			obj.Spec.Ingress.Host = ""
			obj.Spec.Gateway.Enabled = false
			obj.Spec.Gateway.Host = ""

			By("calling the Default method to apply defaults")
			err := defaulter.Default(ctx, obj)

			By("checking that no error occurred")
			Expect(err).To(BeNil())

			By("checking that ingress and gateway defaults are applied")
			Expect(obj.Spec.Ingress.Enabled).To(BeFalse())
			Expect(obj.Spec.Gateway.Enabled).To(BeFalse())
		})
	})

	Context("When creating LiteLLMInstance under Validating Webhook", func() {
		Describe("Image validation", func() {
			It("Should deny creation if image is empty", func() {
				By("setting image to empty")
				obj.Spec.Image = ""

				By("validating the creation")
				warnings, err := validator.ValidateCreate(ctx, obj)

				By("checking that validation failed")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("spec.image: image is required"))
				Expect(warnings).To(BeNil())
			})

			It("Should deny creation if image format is invalid", func() {
				By("setting invalid image format")
				obj.Spec.Image = "invalid-image"

				By("validating the creation")
				warnings, err := validator.ValidateCreate(ctx, obj)

				By("checking that validation failed")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("spec.image: image must be in format 'registry/repository:tag'"))
				Expect(warnings).To(BeNil())
			})

			It("Should deny creation if image has no tag", func() {
				By("setting image without tag")
				obj.Spec.Image = "ghcr.io/berriai/litellm-database"

				By("validating the creation")
				warnings, err := validator.ValidateCreate(ctx, obj)

				By("checking that validation failed")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("spec.image: image must include a tag"))
				Expect(warnings).To(BeNil())
			})

			It("Should allow creation with valid image", func() {
				By("validating the creation")
				warnings, err := validator.ValidateCreate(ctx, obj)

				By("checking that validation passed")
				Expect(err).To(BeNil())
				Expect(warnings).To(BeNil())
			})
		})

		Describe("MasterKey validation", func() {
			It("Should deny creation if masterKey is empty", func() {
				By("setting masterKey to empty")
				obj.Spec.MasterKey = ""

				By("validating the creation")
				warnings, err := validator.ValidateCreate(ctx, obj)

				By("checking that validation failed")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("spec.masterKey: masterKey is required"))
				Expect(warnings).To(BeNil())
			})

			It("Should deny creation if masterKey is too short", func() {
				By("setting masterKey to short value")
				obj.Spec.MasterKey = "short"

				By("validating the creation")
				warnings, err := validator.ValidateCreate(ctx, obj)

				By("checking that validation failed")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("spec.masterKey: masterKey must be at least 8 characters long"))
				Expect(warnings).To(BeNil())
			})

			It("Should deny creation if masterKey is a weak value", func() {
				By("setting masterKey to weak value")
				obj.Spec.MasterKey = "masterkey"

				By("validating the creation")
				warnings, err := validator.ValidateCreate(ctx, obj)

				By("checking that validation failed")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("spec.masterKey: masterKey cannot be a common weak value"))
				Expect(warnings).To(BeNil())
			})

			It("Should allow creation with valid masterKey", func() {
				By("validating the creation")
				warnings, err := validator.ValidateCreate(ctx, obj)

				By("checking that validation passed")
				Expect(err).To(BeNil())
				Expect(warnings).To(BeNil())
			})
		})

		Describe("DatabaseSecretRef validation", func() {
			It("Should deny creation if database secret name is empty", func() {
				By("setting database secret name to empty")
				obj.Spec.DatabaseSecretRef.NameRef = ""

				By("validating the creation")
				warnings, err := validator.ValidateCreate(ctx, obj)

				By("checking that validation failed")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("spec.databaseSecretRef.nameRef: database secret reference is required"))
				Expect(warnings).To(BeNil())
			})

			It("Should deny creation if database secret name is invalid", func() {
				By("setting invalid database secret name")
				obj.Spec.DatabaseSecretRef.NameRef = "invalid-name!"

				By("validating the creation")
				warnings, err := validator.ValidateCreate(ctx, obj)

				By("checking that validation failed")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("spec.databaseSecretRef.nameRef: invalid secret name format"))
				Expect(warnings).To(BeNil())
			})

			It("Should deny creation if database secret keys are missing", func() {
				By("setting empty database secret keys")
				obj.Spec.DatabaseSecretRef.Keys.HostSecret = ""

				By("validating the creation")
				warnings, err := validator.ValidateCreate(ctx, obj)

				By("checking that validation failed")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("spec.databaseSecretRef.keys.hostSecret: host secret key is required"))
				Expect(warnings).To(BeNil())
			})

			It("Should allow creation with valid database secret ref", func() {
				By("validating the creation")
				warnings, err := validator.ValidateCreate(ctx, obj)

				By("checking that validation passed")
				Expect(err).To(BeNil())
				Expect(warnings).To(BeNil())
			})
		})

		Describe("RedisSecretRef validation", func() {
			It("Should deny creation if redis secret name is empty", func() {
				By("setting redis secret name to empty")
				obj.Spec.RedisSecretRef.NameRef = ""

				By("validating the creation")
				warnings, err := validator.ValidateCreate(ctx, obj)

				By("checking that validation failed")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("spec.redisSecretRef.nameRef: redis secret reference is required"))
				Expect(warnings).To(BeNil())
			})

			It("Should deny creation if redis secret name is invalid", func() {
				By("setting invalid redis secret name")
				obj.Spec.RedisSecretRef.NameRef = "invalid-name!"

				By("validating the creation")
				warnings, err := validator.ValidateCreate(ctx, obj)

				By("checking that validation failed")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("spec.redisSecretRef.nameRef: invalid secret name format"))
				Expect(warnings).To(BeNil())
			})

			It("Should deny creation if redis port is invalid", func() {
				By("setting invalid redis port")
				obj.Spec.RedisSecretRef.Keys.PortSecret = 0

				By("validating the creation")
				warnings, err := validator.ValidateCreate(ctx, obj)

				By("checking that validation failed")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("spec.redisSecretRef.keys.portSecret: port must be between 1 and 65535"))
				Expect(warnings).To(BeNil())
			})

			It("Should allow creation with valid redis secret ref", func() {
				By("validating the creation")
				warnings, err := validator.ValidateCreate(ctx, obj)

				By("checking that validation passed")
				Expect(err).To(BeNil())
				Expect(warnings).To(BeNil())
			})
		})

		Describe("Ingress validation", func() {
			It("Should allow creation when ingress is disabled", func() {
				By("validating the creation")
				warnings, err := validator.ValidateCreate(ctx, obj)

				By("checking that validation passed")
				Expect(err).To(BeNil())
				Expect(warnings).To(BeNil())
			})

			It("Should deny creation if ingress is enabled but host is empty", func() {
				By("enabling ingress without host")
				obj.Spec.Ingress.Enabled = true
				obj.Spec.Ingress.Host = ""

				By("validating the creation")
				warnings, err := validator.ValidateCreate(ctx, obj)

				By("checking that validation failed")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("spec.ingress.host: host is required when ingress is enabled"))
				Expect(warnings).To(BeNil())
			})

			It("Should deny creation if ingress host is invalid", func() {
				By("enabling ingress with invalid host")
				obj.Spec.Ingress.Enabled = true
				obj.Spec.Ingress.Host = "invalid-host!"

				By("validating the creation")
				warnings, err := validator.ValidateCreate(ctx, obj)

				By("checking that validation failed")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("spec.ingress.host: invalid host format"))
				Expect(warnings).To(BeNil())
			})

			It("Should deny creation if ingress host is localhost", func() {
				By("enabling ingress with localhost")
				obj.Spec.Ingress.Enabled = true
				obj.Spec.Ingress.Host = "localhost"

				By("validating the creation")
				warnings, err := validator.ValidateCreate(ctx, obj)

				By("checking that validation failed")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("spec.ingress.host: host cannot be localhost or internal IP"))
				Expect(warnings).To(BeNil())
			})

			It("Should allow creation with valid ingress configuration", func() {
				By("enabling ingress with valid host")
				obj.Spec.Ingress.Enabled = true
				obj.Spec.Ingress.Host = "litellm.example.com"

				By("validating the creation")
				warnings, err := validator.ValidateCreate(ctx, obj)

				By("checking that validation passed")
				Expect(err).To(BeNil())
				Expect(warnings).To(BeNil())
			})
		})

		Describe("Gateway validation", func() {
			It("Should allow creation when gateway is disabled", func() {
				By("validating the creation")
				warnings, err := validator.ValidateCreate(ctx, obj)

				By("checking that validation passed")
				Expect(err).To(BeNil())
				Expect(warnings).To(BeNil())
			})

			It("Should deny creation if gateway is enabled but host is empty", func() {
				By("enabling gateway without host")
				obj.Spec.Gateway.Enabled = true
				obj.Spec.Gateway.Host = ""

				By("validating the creation")
				warnings, err := validator.ValidateCreate(ctx, obj)

				By("checking that validation failed")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("spec.gateway.host: host is required when gateway is enabled"))
				Expect(warnings).To(BeNil())
			})

			It("Should allow creation with valid gateway configuration", func() {
				By("enabling gateway with valid host")
				obj.Spec.Gateway.Enabled = true
				obj.Spec.Gateway.Host = "gateway.example.com"

				By("validating the creation")
				warnings, err := validator.ValidateCreate(ctx, obj)

				By("checking that validation passed")
				Expect(err).To(BeNil())
				Expect(warnings).To(BeNil())
			})
		})

		Describe("Cross-field validation", func() {
			It("Should deny creation if ingress and gateway have same host", func() {
				By("enabling both ingress and gateway with same host")
				obj.Spec.Ingress.Enabled = true
				obj.Spec.Ingress.Host = "example.com"
				obj.Spec.Gateway.Enabled = true
				obj.Spec.Gateway.Host = "example.com"

				By("validating the creation")
				warnings, err := validator.ValidateCreate(ctx, obj)

				By("checking that validation failed")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("spec.ingress.host and spec.gateway.host: cannot be the same when both are enabled"))
				Expect(warnings).To(BeNil())
			})

			It("Should deny creation if database and redis secrets are the same", func() {
				By("setting same secret name for database and redis")
				obj.Spec.DatabaseSecretRef.NameRef = "same-secret"
				obj.Spec.RedisSecretRef.NameRef = "same-secret"

				By("validating the creation")
				warnings, err := validator.ValidateCreate(ctx, obj)

				By("checking that validation failed")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("spec.databaseSecretRef.nameRef and spec.redisSecretRef.nameRef: cannot be the same"))
				Expect(warnings).To(BeNil())
			})

			It("Should allow creation with different hosts for ingress and gateway", func() {
				By("enabling both ingress and gateway with different hosts")
				obj.Spec.Ingress.Enabled = true
				obj.Spec.Ingress.Host = "ingress.example.com"
				obj.Spec.Gateway.Enabled = true
				obj.Spec.Gateway.Host = "gateway.example.com"

				By("validating the creation")
				warnings, err := validator.ValidateCreate(ctx, obj)

				By("checking that validation passed")
				Expect(err).To(BeNil())
				Expect(warnings).To(BeNil())
			})
		})

		It("Should allow creation with valid configuration", func() {
			By("validating the creation")
			warnings, err := validator.ValidateCreate(ctx, obj)

			By("checking that validation passed")
			Expect(err).To(BeNil())
			Expect(warnings).To(BeNil())
		})
	})

	Context("When updating LiteLLMInstance under Validating Webhook", func() {
		BeforeEach(func() {
			// Set up old object with same valid configuration
			oldObj.Spec = obj.Spec
		})

		It("Should allow update with valid configuration", func() {
			By("validating the update")
			warnings, err := validator.ValidateUpdate(ctx, oldObj, obj)

			By("checking that validation passed")
			Expect(err).To(BeNil())
			Expect(warnings).To(BeNil())
		})

		It("Should deny update with invalid image", func() {
			By("setting invalid image in new object")
			obj.Spec.Image = ""

			By("validating the update")
			warnings, err := validator.ValidateUpdate(ctx, oldObj, obj)

			By("checking that validation failed")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("spec.image: image is required"))
			Expect(warnings).To(BeNil())
		})

		It("Should deny update with invalid masterKey", func() {
			By("setting invalid masterKey in new object")
			obj.Spec.MasterKey = "short"

			By("validating the update")
			warnings, err := validator.ValidateUpdate(ctx, oldObj, obj)

			By("checking that validation failed")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("spec.masterKey: masterKey must be at least 8 characters long"))
			Expect(warnings).To(BeNil())
		})
	})

	Context("When deleting LiteLLMInstance under Validating Webhook", func() {
		It("Should allow deletion", func() {
			By("validating the deletion")
			warnings, err := validator.ValidateDelete(ctx, obj)

			By("checking that validation passed")
			Expect(err).To(BeNil())
			Expect(warnings).To(BeNil())
		})
	})
})
