package command_test

import (
	"github.com/cloudfoundry/cloud-service-broker/pkg/providers/tf/command"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Context("Terraform Commands", func() {
	Context("initCommand", func() {
		It("calls init with the plugin directory", func() {
			initCommand := command.NewInit("plugindir")
			Expect(initCommand.Command()).To(Equal([]string{"init", "-plugin-dir=plugindir", "-no-color"}))
		})
	})

	Context("init012", func() {
		It("calls init with the plugin directory", func() {
			initCommand := command.NewInit012Command("plugindir")
			Expect(initCommand.Command()).To(Equal([]string{"init", "-plugin-dir=plugindir", "-get-plugins=false", "-no-color"}))
		})
	})
})