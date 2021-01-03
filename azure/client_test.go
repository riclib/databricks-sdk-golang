package azure_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	databricks "github.com/xinsnake/databricks-sdk-golang"
	. "github.com/xinsnake/databricks-sdk-golang/azure"
)

var _ = Describe("Client", func() {
	testUser := "test-user"
	option := databricks.DBClientOption{
		User: testUser,
	}

	var dbClient DBClient
	dbClient.Init(option)

	It("Should init with correct option", func() {
		Expect(dbClient.Option.User).To(Equal(testUser))
	})

	It("Should return initialised ClustersAPI", func() {
		Expect(dbClient.Clusters().Client.Option.User).To(Equal(testUser))
	})

	It("Should return initialised DbfsAPI", func() {
		Expect(dbClient.Dbfs().Client.Option.User).To(Equal(testUser))
	})

	It("Should return initialised GroupsAPI", func() {
		Expect(dbClient.Groups().Client.Option.User).To(Equal(testUser))
	})

	It("Should return initialised InstancePoolsAPI", func() {
		Expect(dbClient.InstancePools().Client.Option.User).To(Equal(testUser))
	})

	It("Should return initialised JobsAPI", func() {
		Expect(dbClient.Jobs().Client.Option.User).To(Equal(testUser))
	})

	It("Should return initialised LibrariesAPI", func() {
		Expect(dbClient.Libraries().Client.Option.User).To(Equal(testUser))
	})

	It("Should return initialised ScimAPI", func() {
		Expect(dbClient.Scim().Client.Option.User).To(Equal(testUser))
	})

	It("Should return initialised SecretsAPI", func() {
		Expect(dbClient.Secrets().Client.Option.User).To(Equal(testUser))
	})

	It("Should return initialised TokenAPI", func() {
		Expect(dbClient.Token().Client.Option.User).To(Equal(testUser))
	})

	It("Should return initialised WorkspaceAPI", func() {
		Expect(dbClient.Workspace().Client.Option.User).To(Equal(testUser))
	})
})
