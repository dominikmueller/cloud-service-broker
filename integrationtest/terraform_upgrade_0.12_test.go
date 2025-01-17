package integrationtest_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudfoundry/cloud-service-broker/db_service/models"

	"github.com/cloudfoundry/cloud-service-broker/integrationtest/helper"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
	"github.com/pborman/uuid"
	"github.com/pivotal-cf/brokerapi/v8/domain"
)

var lastOperationPollingFrequency = time.Second * 1

func pollLastOperation(testHelper *helper.TestHelper, serviceInstanceGUID string) func() domain.LastOperationState {
	return func() domain.LastOperationState {
		lastOperationResponse := testHelper.Client().LastOperation(serviceInstanceGUID, requestID())
		Expect(lastOperationResponse.Error).NotTo(HaveOccurred())
		Expect(lastOperationResponse.StatusCode).To(Or(Equal(http.StatusOK), Equal(http.StatusGone)))
		var receiver domain.LastOperation
		err := json.Unmarshal(lastOperationResponse.ResponseBody, &receiver)
		Expect(err).NotTo(HaveOccurred())
		return receiver.State
	}
}

var _ = Describe("Terraform 0.12 Upgrade", func() {
	const serviceOfferingGUID = "df2c1512-3013-11ec-8704-2fbfa9c8a802"
	const servicePlanGUID = "e59773ce-3013-11ec-9bbb-9376b4f72d14"

	var (
		testHelper *helper.TestHelper
		session    *Session
	)

	BeforeEach(func() {
		testHelper = helper.New(csb)
		testHelper.BuildBrokerpak(testHelper.OriginalDir, "fixtures", "brokerpak-terraform-0.12")

		session = testHelper.StartBroker()
	})

	AfterEach(func() {
		session.Terminate()
		testHelper.Restore()
	})

	terraformStateVersion := func(serviceInstanceGUID string) string {
		var tfDeploymentReceiver models.TerraformDeployment
		Expect(testHelper.DBConn().Where("id = ?", fmt.Sprintf("tf:%s:", serviceInstanceGUID)).First(&tfDeploymentReceiver).Error).NotTo(HaveOccurred())
		var workspaceReceiver struct {
			State []byte `json:"tfstate"`
		}
		Expect(json.Unmarshal(tfDeploymentReceiver.Workspace, &workspaceReceiver)).NotTo(HaveOccurred())
		var stateReceiver struct {
			Version string `json:"terraform_version"`
		}
		Expect(json.Unmarshal(workspaceReceiver.State, &stateReceiver)).NotTo(HaveOccurred())
		return stateReceiver.Version
	}

	Context("TF Upgrades are enabled", func() {
		It("runs 'terraform apply' at each version in the upgrade path", func() {
			By("provisioning a service instance at 0.12")
			serviceInstanceGUID := uuid.New()
			provisionResponse := testHelper.Client().Provision(serviceInstanceGUID, serviceOfferingGUID, servicePlanGUID, requestID(), nil)
			Expect(provisionResponse.Error).NotTo(HaveOccurred())
			Expect(provisionResponse.StatusCode).To(Equal(http.StatusAccepted))

			Eventually(pollLastOperation(testHelper, serviceInstanceGUID), time.Minute*2, lastOperationPollingFrequency).ShouldNot(Equal(domain.InProgress))
			Expect(pollLastOperation(testHelper, serviceInstanceGUID)()).To(Equal(domain.Succeeded))
			Expect(terraformStateVersion(serviceInstanceGUID)).To(Equal("0.12.21"))

			By("updating the brokerpak and restarting the broker")
			session.Terminate()
			testHelper.BuildBrokerpak(testHelper.OriginalDir, "fixtures", "brokerpak-terraform-upgrade")

			session = testHelper.StartBroker("TERRAFORM_UPGRADES_ENABLED=true")

			By("running 'cf update-service'")
			updateResponse := testHelper.Client().Update(serviceInstanceGUID, serviceOfferingGUID, servicePlanGUID, requestID(), nil)
			Expect(updateResponse.Error).NotTo(HaveOccurred())
			Expect(updateResponse.StatusCode).To(Equal(http.StatusAccepted))

			Eventually(pollLastOperation(testHelper, serviceInstanceGUID), time.Minute*2, lastOperationPollingFrequency).ShouldNot(Equal(domain.InProgress))
			Expect(pollLastOperation(testHelper, serviceInstanceGUID)()).To(Equal(domain.Succeeded))

			By("observing that the TF state file has been updated to the latest version")
			Expect(terraformStateVersion(serviceInstanceGUID)).To(Equal("1.1.6"))
		})
	})

	Context("TF Upgrades are disabled", func() {
		It("does not upgrade the instance", func() {
			By("provisioning a service instance at 0.12")
			serviceInstanceGUID := uuid.New()
			provisionResponse := testHelper.Client().Provision(serviceInstanceGUID, serviceOfferingGUID, servicePlanGUID, requestID(), nil)
			Expect(provisionResponse.Error).NotTo(HaveOccurred())
			Expect(provisionResponse.StatusCode).To(Equal(http.StatusAccepted))

			Eventually(pollLastOperation(testHelper, serviceInstanceGUID), time.Minute*2, lastOperationPollingFrequency).Should(Equal(domain.Succeeded))
			Expect(terraformStateVersion(serviceInstanceGUID)).To(Equal("0.12.21"))

			By("updating the brokerpak and restarting the broker")
			session.Terminate()
			testHelper.BuildBrokerpak(testHelper.OriginalDir, "fixtures", "brokerpak-terraform-upgrade")
			session = testHelper.StartBroker()

			By("running 'cf update-service'")
			updateResponse := testHelper.Client().Update(serviceInstanceGUID, serviceOfferingGUID, servicePlanGUID, requestID(), nil)
			Expect(updateResponse.Error).NotTo(HaveOccurred())
			Expect(updateResponse.StatusCode).To(Equal(http.StatusAccepted))

			Eventually(pollLastOperation(testHelper, serviceInstanceGUID), time.Minute*2, lastOperationPollingFrequency).ShouldNot(Equal(domain.InProgress))
			Eventually(pollLastOperation(testHelper, serviceInstanceGUID), time.Minute*2, lastOperationPollingFrequency).Should(Equal(domain.Failed))

			By("observing that the TF version remains the same in the state file")
			Expect(terraformStateVersion(serviceInstanceGUID)).To(Equal("0.12.21"))
		})
	})

})
