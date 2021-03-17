package models_test

import (
	"encoding/json"
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/xerardoo/sapip/models"
	_ "github.com/xerardoo/sapip/tester"
)

func TestIncident(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Incident Test Suite")
}

var _ = Describe("Incident", func() {

	incident := new(models.Incident)
	incident.Type = "rutina"
	incident.Description = "se hizo alto preventivo"
	incident.UserID = 1
	incident.LocationID = 1

	// Add new
	incidentAdded, err := incident.Add()
	It("cant be added", func() {
		Expect(err).NotTo(HaveOccurred())
		Expect(incidentAdded.Description).Should(Equal("se hizo alto preventivo"))
	})

	var count int64
	err = models.DB.Model(models.Incident{}).Count(&count).Error
	It("cant be added", func() {
		Expect(err).NotTo(HaveOccurred())
		fmt.Println("COUNT", count)
	})

	location, err := incidentAdded.GetLocation()
	It("cant be added", func() {
		Expect(err).NotTo(HaveOccurred())
		res, _ := json.Marshal(location)
		fmt.Println("ROWS", string(res))
	})
	user, err := incidentAdded.GetUser()
	It("cant be added", func() {
		Expect(err).NotTo(HaveOccurred())
		Expect(user.FirstName).Should(Equal("Juan"))
	})

	var incidents []models.Incident
	err = models.DB.Find(&incidents).Error
	It("cant be added", func() {
		Expect(err).NotTo(HaveOccurred())
	})
})
